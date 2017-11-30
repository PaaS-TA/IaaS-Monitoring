package controller

import (
	"kr/paasta/monitoring/openstack/models"
	"kr/paasta/monitoring/openstack/services"
	"kr/paasta/monitoring/openstack/utils"
	"github.com/rackspace/gophercloud"
	monascagopher "github.com/gophercloud/gophercloud"
	"net/http"
	"errors"
	"encoding/json"
	"github.com/monasca/golang-monascaclient/monascaclient"
	"fmt"
)

//Compute Node Controller
type LoginController struct{
	OpenstackProvider models.OpenstackProvider
	MonAuth     	  monascagopher.AuthOptions
	MonClient         monascaclient.Client
}

func NewLoginController(openstackProvider models.OpenstackProvider, monsClient monascaclient.Client, auth monascagopher.AuthOptions) *LoginController {
	return &LoginController{
		OpenstackProvider: openstackProvider,
		MonAuth: auth,
		MonClient: monsClient,
	}

}

func (s *LoginController)Ping(w http.ResponseWriter, r *http.Request) {

	token, _ := utils.GenerateRandomString(32)
	session := models.SessionManager.Load(r)

	testToken := r.Header.Get(models.TEST_TOKEN_NAME)
	if testToken != ""{
		w.Header().Add(models.TEST_TOKEN_NAME, token)
	}else{
		fmt.Println("pint Token::::", token)
		session.PutString(w, token, token)
		w.Header().Add(models.CSRF_TOKEN_NAME, token)
	}

	utils.RenderJsonResponse(nil, w)
}



func (s *LoginController)Login(w http.ResponseWriter, r *http.Request) {

	session := models.SessionManager.Load(r)

	reqCsrfToken := r.Header.Get(models.CSRF_TOKEN_NAME)
	testToken    := r.Header.Get(models.TEST_TOKEN_NAME)
	sessionCsrfToken, _ := session.GetString(reqCsrfToken)



	if sessionCsrfToken == "" && testToken == ""{
		utils.RenderJsonForbiddenResponse("Forbidden login", w)
	}else{

		fmt.Println("Login Test Token:", testToken)
		var apiRequest models.User

		err := json.NewDecoder(r.Body).Decode(&apiRequest)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			return
		}else {
			err := loginValidate(apiRequest)
			if err != nil{
				loginErr := utils.GetError().GetCheckErrorMessage(err)
				utils.ErrRenderJsonResponse(loginErr, w)
				return
			}

			userInfo, provider,  err := services.GetLoginService(s.OpenstackProvider).Login(apiRequest)
			loginErr := utils.GetError().GetCheckErrorMessage(err)

			if loginErr != nil {

				utils.ErrRenderJsonResponse(loginErr, w)
				return
			} else {

				//Monasca Client 설정

				var userSession models.UserSession
				if testToken != ""{
					userSession.CsrfToken = testToken
				}else{
					userSession.CsrfToken = reqCsrfToken
				}

				userSession.Username = userInfo.Username
				//userSession.OpenstackToken    = userInfo.Token
				s.MonAuth.Username  = apiRequest.Username
				s.MonAuth.Password  = apiRequest.Password
				userSession.MonAuth = s.MonAuth

				//Ping Token은 login후 삭제 하고 Session에 저장한다.
				session.Remove(w, sessionCsrfToken)

				if testToken != ""{
					session.PutObject(w, models.USER_SESSION_NAME + testToken, userSession)
				}else{
					//fmt.Println("Login==========>", models.USER_SESSION_NAME + reqCsrfToken)
					session.PutObject(w, models.USER_SESSION_NAME + reqCsrfToken, userSession)
				}

				clientProvider := map[string]*gophercloud.ProviderClient{
					sessionCsrfToken: provider,
				}

				clientProvider[sessionCsrfToken] = provider
				models.OpenStackClient = append(models.OpenStackClient, clientProvider)

				var result models.User
				result.Username = userInfo.Username
				utils.RenderJsonResponse(result, w)

				return
			}
		}
	}
}


func (s *LoginController)Logout(w http.ResponseWriter, r *http.Request) {

	session := models.SessionManager.Load(r)
	reqCsrfToken := r.Header.Get(models.CSRF_TOKEN_NAME)
	testToken    := r.Header.Get(models.TEST_TOKEN_NAME)

	sessionCsrfToken, _ := session.GetString(models.USER_SESSION_NAME  + reqCsrfToken)


	if sessionCsrfToken == "" && testToken == ""{
		utils.RenderJsonForbiddenResponse("Forbidden logout", w)
	}else{
		provider, _, _ := utils.GetOpenstackProvider(r)
		services.GetLoginService(s.OpenstackProvider).Logout(provider)


		session.Remove(w, sessionCsrfToken)
		var newClient []map[string]*gophercloud.ProviderClient
		for _, data := range models.OpenStackClient{
			if testToken != ""{
				if data[reqCsrfToken] == nil{
					newClient = append(newClient, data)
				}
			}else{
				if data[testToken] == nil{
					newClient = append(newClient, data)
				}
			}

		}
		models.OpenStackClient = newClient
		utils.RenderJsonLogoutResponse(nil, w)
	}

}

func loginValidate(apiRequest models.User) error {

	if apiRequest.Username == ""{
		return errors.New("Required input value does not exist. [username]");
	}

	if apiRequest.Password == ""{
		return errors.New("Required input value does not exist. [password]");
	}

	return nil
}
