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
	session.PutString(w, token, token)
	w.Header().Add(models.CSRF_TOKEN_NAME, token)
	utils.RenderJsonResponse(nil, w)
}

func (s *LoginController)LogOut(w http.ResponseWriter, r *http.Request) {

	session := models.SessionManager.Load(r)


	reqCsrfToken := r.Header.Get(models.CSRF_TOKEN_NAME)
	sessionCsrfToken, _ := session.GetString(models.USER_SESSION_NAME  + reqCsrfToken)


	if sessionCsrfToken == "" {
		utils.RenderJsonForbiddenResponse("Forbidden logout", w)
	}else{
		provider, _, _ := utils.GetOpenstackProvider(r)
		services.GetLoginService(s.OpenstackProvider).LogOut(provider)
		utils.RenderJsonResponse(nil, w)

		session.Remove(w, sessionCsrfToken)
		var newClient []map[string]*gophercloud.ProviderClient
		for _, data := range models.OpenStackClient{
			if data[reqCsrfToken] == nil{
				newClient = append(newClient, data)
			}
		}
		models.OpenStackClient = newClient

		/*userSession := new(models.UserSession)
		session.GetObject(models.USER_SESSION_NAME + reqCsrfToken, userSession)
		var result models.User
		result.Username = userSession.Username*/

		utils.RenderJsonResponse("logout", w)
		return
	}

}

func (s *LoginController)Login(w http.ResponseWriter, r *http.Request) {

	session := models.SessionManager.Load(r)

	reqCsrfToken := r.Header.Get(models.CSRF_TOKEN_NAME)
	sessionCsrfToken, _ := session.GetString(reqCsrfToken)

	if sessionCsrfToken == "" {
		utils.RenderJsonForbiddenResponse("Forbidden login", w)
	}else{
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
				userSession.CsrfToken = reqCsrfToken
				userSession.Username = userInfo.Username
				//userSession.OpenstackToken    = userInfo.Token
				s.MonAuth.Username  = apiRequest.Username
				s.MonAuth.Password  = apiRequest.Password
				userSession.MonAuth = s.MonAuth

				fmt.Println("MonAuth===>",userSession.MonAuth)
				//Ping Token은 login후 삭제 하고 Session에 저장한다.
				session.Remove(w, sessionCsrfToken)
				err := session.PutObject(w, models.USER_SESSION_NAME + reqCsrfToken, userSession)
				fmt.Println("ERR:::", err)

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


func loginValidate(apiRequest models.User) error {

	if apiRequest.Username == ""{
		return errors.New("Required input value does not exist. [username]");
	}

	if apiRequest.Password == ""{
		return errors.New("Required input value does not exist. [password]");
	}

	return nil
}
