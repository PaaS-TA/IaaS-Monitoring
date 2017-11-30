package utils

import (
	"strconv"
	"fmt"
	"net/http"
	"encoding/json"
	"log"
	"io/ioutil"
	"kr/paasta/monitoring/openstack/models"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	client "github.com/influxdata/influxdb/client/v2"
	"math"
	"strings"
	"encoding/base64"
	"crypto/rand"
	"github.com/monasca/golang-monascaclient/monascaclient"
)




type errorMessage struct{
	models.ErrMessage
}

func GetError() *errorMessage{
	return &errorMessage{}
}

func (e errorMessage) GetCheckErrorMessage(err error) (models.ErrMessage) {

	if err != nil{
		errMessage := models.ErrMessage{
			"Message": err.Error() ,
		}
		return  errMessage
	}else{
		return nil
	}
}



func Round(num float64) int {
	return int(num + math.Copysign(0.5, num))
}

func RoundFloat(num float64, precision int) float64 {
	output := math.Pow(10, float64(precision))
	return float64(Round(num * output)) / output
}

func RoundFloatDigit2(num float64) float64 {
	return RoundFloat(num , 2)
}

func FloattostrDigit2(fv float64) string {
	return strconv.FormatFloat(RoundFloatDigit2(fv), 'f', 2, 64)
}

func GetConnectionString(host, port, user, pass , dbname string) string {

	return fmt.Sprintf("%s:%s@%s([%s]:%s)/%s%s",
		user, pass, "tcp", host, port, dbname, "")

}

func StringArrayDistinct(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}




//오류 체크 모듈로 오류 발생시 오류 메시지 리턴
func (e errorMessage) CheckError(resp client.Response, err error) (client.Response, models.ErrMessage) {


	if err != nil {
		errMessage := models.ErrMessage{
			"Message": err.Error() ,
		}
		return resp , errMessage

	}else if resp.Error() != nil {
		errMessage := models.ErrMessage{
			"Message": resp.Err ,
		}
		return resp, errMessage
	}else {

		return resp, nil
	}
}


func ResponseUnmarshal(response *http.Response, resErr error) (map[string]interface{}, error ){

	if resErr != nil{
		return nil, resErr
	}
	var data interface{}
	rawdata, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(rawdata, &data)
	msg := data.(map[string]interface{})

	return msg, nil

}

func GetMonascaClient(r *http.Request, client monascaclient.Client) (monascaclient.Client,  error) {

	var err error
	session := models.SessionManager.Load(r)
	userSession := new(models.UserSession)
	reqToken := r.Header.Get(models.CSRF_TOKEN_NAME)
	testToken  := r.Header.Get(models.TEST_TOKEN_NAME)

	if testToken == ""{
		err = session.GetObject(models.USER_SESSION_NAME + reqToken, userSession)
	}else{

		userSession.MonAuth.Username = models.TestUserName
		userSession.MonAuth.Password = models.TestPassword
		userSession.MonAuth.DomainName = models.TestDomainName
		userSession.MonAuth.TenantID = models.TestTenantID
		userSession.MonAuth.IdentityEndpoint = models.TestIdentityEndpoint
	}
	if err != nil{
		return client,  err
	}

	fmt.Println("authInfo=========+>", userSession.MonAuth)
	client.SetKeystoneConfig(&userSession.MonAuth)

	return client, err
}

func GetOpenstackProvider(r *http.Request) (provider *gophercloud.ProviderClient, username string ,  err error) {

	session := models.SessionManager.Load(r)
	userSession := new(models.UserSession)
	reqToken   := r.Header.Get(models.CSRF_TOKEN_NAME)
	testToken  := r.Header.Get(models.TEST_TOKEN_NAME)

	if testToken == ""{
		err = session.GetObject(models.USER_SESSION_NAME + reqToken, userSession)
	}else{
		//Test Case인경우
		err = session.GetObject(models.USER_SESSION_NAME + testToken, userSession)
		userSession.Username = r.Header.Get("username")
	}

	if err != nil{
		return provider, username, err
	}

	for _, data := range models.OpenStackClient{
		if data[reqToken] != nil{
			provider = data[userSession.CsrfToken]
			break
		}
	}


	return provider, userSession.Username, err
}


//Get Openstack Admin Token - based on Default Domain & Admin tenant
func GetAdminToken(openstack_provider models.OpenstackProvider) (*gophercloud.ProviderClient, error) {

	// Option 1: Pass in the values yourself
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: openstack_provider.IdentityEndpoint,
		Username: openstack_provider.Username,
		UserID: openstack_provider.UserId,
		Password: openstack_provider.Password,
		TenantName: openstack_provider.TenantName,
		DomainName: openstack_provider.Domain,
	}

	//Provider is the top-level client that all of your OpenStack services
	provider, err := openstack.AuthenticatedClient(opts)

	return provider, err
}


func GetAdminToken2(openstack_provider models.OpenstackProvider, tokenId string) (*gophercloud.ProviderClient, error) {

	// Option 1: Pass in the values yourself
	opts := gophercloud.AuthOptions{
		IdentityEndpoint: openstack_provider.IdentityEndpoint,
		Username: openstack_provider.Username,
		UserID: openstack_provider.UserId,
		Password: openstack_provider.Password,
		TenantName: openstack_provider.TenantName,
		DomainName: openstack_provider.Domain,
		TokenID: tokenId,
	}

	//Provider is the top-level client that all of your OpenStack services
	provider, err := openstack.AuthenticatedClient(opts)

	fmt.Println("TokenId==========+>>",provider.TokenID)
	return provider, err
}

func TypeChecker_int(target interface{}) interface{}{
	switch target.(type) {
	case int:
		// v is an int here, so e.g. v + 1 is possible.
		return target.(int)
	case float64:
		// v is a float64 here, so e.g. v + 1.0 is possible.
		return int(target.(float64))
	case string:
		// v is a string here, so e.g. v + " Yeah!" is possible.
		i, _ := strconv.ParseInt(target.(string), 10, 0)
		return i
	case json.Number:
		jsonValue := target.(json.Number)
		f, _ := strconv.ParseInt(jsonValue.String(), 10, 0)

		//f, _ := strconv.ParseFloat(jsonValue.String(),64)
		return f
	case nil:
		// v is a string here, so e.g. v + " Yeah!" is possible.
		return int(0)
	default:
		// And here I'm feeling dumb. ;)
		return int(0)
	}
}

func TypeChecker_float64(target interface{}) interface{}{

	switch target.(type) {
	case int:
		// v is an int here, so e.g. v + 1 is possible.
		return float64(target.(int))
	case float64:
		// v is a float64 here, so e.g. v + 1.0 is possible.
		return target.(float64)
	case string:
		// v is a string here, so e.g. v + " Yeah!" is possible.
		f, _ := strconv.ParseFloat(target.(string), 64)
		return f
	case nil:
		// v is a string here, so e.g. v + " Yeah!" is possible.
		return float64(0)
	case json.Number:
		jsonValue := target.(json.Number)
		f, _ := strconv.ParseFloat(jsonValue.String(),64)
		return f

	default:
		// And here I'm feeling dumb. ;)
		return float64(0)
	}
}

func TypeChecker_string(target interface{}) interface{}{
	switch target.(type) {
	case int:
		// v is an int here, so e.g. v + 1 is possible.
		return fmt.Sprintf("%d", target)
	case float64:
		// v is a float64 here, so e.g. v + 1.0 is possible.
		return fmt.Sprintf("%f", target)
	case string:
		// v is a string here, so e.g. v + " Yeah!" is possible.
		return target.(string)
	case nil:
		// v is a string here, so e.g. v + " Yeah!" is possible.
		return ""
	default:
		// And here I'm feeling dumb. ;)
		return ""
	}
}

func GetVmStatusCount(noStatusList, runningList, idleList, pausedList, shutDownList, shutOffList, crashedList, powerOffList []string) []models.VmState{

	var vmStatusList []models.VmState

	//if len(noStatusList) != 0 {
	var vmStatusNo models.VmState
	vmStatusNo.VmStateName = models.VM_STATUS_NO
	vmStatusNo.VmCnt = len(noStatusList)
	vmStatusList = append(vmStatusList, vmStatusNo)
	//}

	//if len(runningList) != 0 {
	var vmStatusRunning models.VmState
	vmStatusRunning.VmStateName = models.VM_STATUS_RUNNING
	vmStatusRunning.VmCnt = len(runningList)
	vmStatusList = append(vmStatusList, vmStatusRunning)
	//}

	//if len(idleList) != 0 {
	var vmStatusIdle models.VmState
	vmStatusIdle.VmStateName = models.VM_STATUS_IDLE
	vmStatusIdle.VmCnt = len(idleList)
	vmStatusList = append(vmStatusList, vmStatusIdle)
	//}

	//if len(pausedList) != 0 {
	var vmStatusPaused models.VmState
	vmStatusPaused.VmStateName = models.VM_STATUS_PAUSED
	vmStatusPaused.VmCnt = len(pausedList)
	vmStatusList = append(vmStatusList, vmStatusPaused)
	//}

	//if len(shutDownList) != 0 {
	var vmStatusShutDown models.VmState
	vmStatusShutDown.VmStateName = models.VM_STATUS_SHUTDOWN
	vmStatusShutDown.VmCnt = len(shutDownList)
	vmStatusList = append(vmStatusList, vmStatusShutDown)
	//}

	//if len(shutOffList) != 0 {
	var vmStatusShutOff models.VmState
	vmStatusShutOff.VmStateName = models.VM_STATUS_SHUTOFF
	vmStatusShutOff.VmCnt = len(shutOffList)
	vmStatusList = append(vmStatusList, vmStatusShutOff)
	//}

	//if len(crashedList) != 0 {
	var vmStatusCrash models.VmState
	vmStatusCrash.VmStateName = models.VM_STATUS_CRASHED
	vmStatusCrash.VmCnt = len(crashedList)
	vmStatusList = append(vmStatusList, vmStatusCrash)
	//}

	//if len(powerOffList) != 0 {
	var vmStatusPowerOff models.VmState
	vmStatusPowerOff.VmStateName = models.VM_STATUS_POEWR_OFF
	vmStatusPowerOff.VmCnt = len(powerOffList)
	vmStatusList = append(vmStatusList, vmStatusPowerOff)
	//}

	return vmStatusList
}

func ErrRenderJsonResponse(data interface{}, w http.ResponseWriter) {

	var errorCode float64
	var errorStruct models.ErrorMessageStruct

	errData := data.(models.ErrMessage)
	errorMessage := errData["Message"]

	//Openstack Error 인경우
	if strings.Contains(errorMessage.(string), "instead") {
		//Error Message 정보가 instead 이후에 json data로 되어 있음.
		errorJson := strings.SplitAfter(errorMessage.(string), "instead")

		var errorString interface{}
		json.Unmarshal([]byte(errorJson[1]), &errorString)
		errorMsgJson := errorString.(map[string]interface{})
		for _, v := range errorMsgJson {
			errorDetail := v.(map[string]interface{})
			errorStruct.HttpStatus = int(errorDetail["code"].(float64))
			errorStruct.Message =  errorDetail["message"].(string)
		}
	} else {
		errorStruct.Message =  errorMessage.(string)
		if errData["HttpStatus"] != nil {
			errorStruct.HttpStatus = errData["HttpStatus"].(int)
			errorCode = float64(errData["HttpStatus"].(int))
		} else {
			errorStruct.HttpStatus = 500
			errorCode = float64(500)
		}
	}

	fmt.Println("===>",errorCode)
	js, err := json.Marshal(errorStruct)

	if err != nil {
		log.Fatalln("Error writing JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(errorStruct.HttpStatus)
	w.Write(js)
	return
}

func RenderJsonUnAuthResponse(data interface{}, status int, w http.ResponseWriter) {

	var errorCode float64
	var errorStruct models.ErrorMessageStruct

	errData := data.(models.ErrMessage)
	errorMessage := errData["Message"]
	errorStruct.Message =  errorMessage.(string)
	if errData["HttpStatus"] != nil {
		errorStruct.HttpStatus = errData["HttpStatus"].(int)
		errorCode = float64(errData["HttpStatus"].(int))
	} else {
		errorStruct.HttpStatus = status
		errorCode = float64(status)
	}

	js, err := json.Marshal(errorStruct)

	if err != nil {
		log.Fatalln("Error writing JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(int(errorCode))
	w.Write(js)
	return
}

func RenderJsonResponse(data interface{}, w http.ResponseWriter) {

	js, err := json.Marshal(data)
	if err != nil {
		log.Fatalln("Error writing JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(js)
	return
}

func RenderJsonLogoutResponse(data interface{}, w http.ResponseWriter) {

	js, err := json.Marshal(data)
	if err != nil {
		log.Fatalln("Error writing JSON:", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(js)

	return
}


func RenderJsonForbiddenResponse(data interface{}, w http.ResponseWriter) {

	js, err := json.Marshal(data)
	if err != nil {
		log.Fatalln("Error writing JSON:", err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusForbidden)
	w.Write(js)
	return
}





func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

// GenerateRandomString returns a URL-safe, base64 encoded
// securely generated random string.
// It will return an error if the system's secure random
// number generator fails to function correctly, in which
// case the caller should not continue.
func GenerateRandomString(s int) (string, error) {
	b, err := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b), err
}