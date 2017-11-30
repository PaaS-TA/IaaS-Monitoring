package controller_test

import (
	. "github.com/onsi/ginkgo"
	"testing"
	"net/http"
	"bufio"
	"strings"
	"io"
	"io/ioutil"
	"fmt"
	"net/http/httptest"
	"github.com/jinzhu/gorm"
	"os"
	"encoding/json"
	"log"
	"github.com/cihub/seelog"
	"strconv"
	"kr/paasta/monitoring/openstack/handlers"
	"kr/paasta/monitoring/openstack/utils"
	"kr/paasta/monitoring/openstack/models"
	"gopkg.in/olivere/elastic.v3"
	"github.com/monasca/golang-monascaclient/monascaclient"
	"github.com/gophercloud/gophercloud"
	"github.com/influxdata/influxdb/client/v2"
	"github.com/alexedwards/scs"
	"github.com/stretchr/testify/assert"
	_ "github.com/go-sql-driver/mysql"
)

type DBConfig struct {
	DbType string
	UserName string
	UserPassword string
	Host string
	Port string
	DbName string
}

type Config map[string]string
var TestToken string
type Response struct {
	Token   string
	Content string
	Code    int
}

type PingResponse struct {
	Token   string
	Code    int
}

type LoginRequestBody struct {
	UserId	 string   `json:"username"`
	Password string   `json:"password"`
}


var Logger seelog.LoggerInterface

var _ = Describe("LoginController", func() {

	Describe("Login Controller", func() {
		Context("Login & Logout ", func() {

			It("Login", func() {
				var query LoginRequestBody

				query.UserId = "admin"
				query.Password = "cfmonit"

				data, _ := json.Marshal(query)

				res, err := DoPost(testUrl + "/v1/login", TestToken, strings.NewReader(string(data)))

				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Logout", func() {
				DoLogout(testUrl + "/v1/logout", TestToken)
			})

		})
	})

	BeforeSuite(func() {

		models.SessionManager = *scs.NewCookieManager("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")

		config, err := readConfig(`../test_config.ini`)
		if err != nil {
			fmt.Errorf("read config file error: %s", err)
			os.Exit(0)
		}

		xmlFile, err := ReadXmlConfig(`../log_config.xml`)
		if err != nil {
			log.Println(err)
			os.Exit(-1)
		}
		logger, err := seelog.LoggerFromConfigAsBytes([]byte(xmlFile))

		if err != nil {
			fmt.Println(err)
			return
		}
		models.MonitLogger = logger
		UseLogger(logger)

		timeGap, _ := strconv.Atoi(config["gmt.time.gap"])
		models.GmtTimeGap = timeGap

		configDbCon := new(DBConfig)
		configDbCon.DbType        = config["monitoring.db.type"]
		configDbCon.DbName        = config["monitoring.db.dbname"]
		configDbCon.UserName      = config["monitoring.db.username"]
		configDbCon.UserPassword  = config["monitoring.db.password"]
		configDbCon.Host          = config["monitoring.db.host"]
		configDbCon.Port          = config["monitoring.db.port"]
		connectionString := utils.GetConnectionString(configDbCon.Host , configDbCon.Port, configDbCon.UserName, configDbCon.UserPassword, configDbCon.DbName )

		fmt.Println("String:",connectionString)
		dbAccessObj, dbErr := gorm.Open(configDbCon.DbType, connectionString + "?charset=utf8&parseTime=true")
		if dbErr != nil{
			fmt.Println("err::",dbErr)
		}

		//Alarm 처리 내역 정보 Table 생성
		dbAccessObj.Debug().AutoMigrate(&models.AlarmActionHistory{})

		//InfluxDB Info
		url     ,  _ := config["metric.db.url"]
		userName,  _ := config["metric.db.username"]
		password,  _ := config["metric.db.password"]

		InfluxServerClient, _ := client.NewHTTPClient(client.HTTPConfig{
			Addr: url,
			Username: userName,
			Password: password,
		})

		// ElasticSearch
		elasticUrl, _ := config["elastic.url"]
		elasticClient, err := elastic.NewClient(
			elastic.SetURL(fmt.Sprintf("http://%s", elasticUrl)),
			elastic.SetSniff(false),
		)
		var openstackProvider models.OpenstackProvider
		openstackProvider.Region, _ 	= config["default.region"]
		openstackProvider.Username, _ 	= config["default.username"]
		openstackProvider.Password, _ 	= config["default.password"]
		openstackProvider.Domain, _ 	= config["default.domain"]
		openstackProvider.TenantName, _ 	= config["default.tenant_name"]
		openstackProvider.AdminProjectId, _ 	= config["default.project_id"]
		openstackProvider.KeystoneUrl, _ 	= config["keystone.url"]
		openstackProvider.IdentityEndpoint, _ 	= config["identity.endpoint"]
		openstackProvider.RabbitmqUser, _ 	= config["rabbitmq.user"]
		openstackProvider.RabbitmqPass, _	= config["rabbitmq.pass"]
		openstackProvider.RabbitmqTargetNode, _ = config["rabbitmq.target.node"]

		models.MetricDBName, _ 		= config["metric.db.name"]
		models.NovaUrl, _ 		= config["nova.target.url"]
		models.NovaVersion, _ 		= config["nova.target.version"]
		models.NeutronUrl, _ 		= config["neutron.target.url"]
		models.NeutronVersion, _ 	= config["neutron.target.version"]
		models.KeystoneUrl, _ 		= config["keystone.target.url"]
		models.KeystoneVersion, _ 	= config["keystone.target.version"]
		models.CinderUrl, _ 		= config["cinder.target.url"]
		models.CinderVersion, _ 	= config["cinder.target.version"]
		models.GlanceUrl, _ 		= config["glance.target.url"]
		models.GlanceVersion,_ 		= config["glance.target.version"]
		models.DefaultProjectId, _	= config["default.project_id"]
		models.RabbitMqIp, _ 		= config["rabbitmq.ip"]
		models.RabbitMqPort, _ 	        = config["rabbitmq.port"]
		models.GMTTimeGap, _ 	        = strconv.ParseInt(config["gmt.time.gap"], 10, 64)

		monClient := monascaclient.New()
		monClient.SetBaseURL(config["monasca.url"])
		timeOut, _ := strconv.Atoi(config["monasca.connect.timeout"])
		monClient.SetTimeout(timeOut)

		tls, _ := strconv.ParseBool(config["monasca.secure.tls"])
		monClient.SetInsecure(tls)

		auth := gophercloud.AuthOptions{
			DomainName : config["default.domain"],
			IdentityEndpoint : config["keystone.url"],
			Username : config["default.username"],
			Password : config["default.password"],
			TenantID : config["default.project_id"],
		}
		models.TestUserName = auth.Username
		models.TestPassword = auth.Password
		models.TestTenantID = auth.TenantID
		models.TestDomainName = auth.DomainName
		models.TestIdentityEndpoint = auth.IdentityEndpoint

		var handler http.Handler
		handler = handlers.NewHandler(openstackProvider, InfluxServerClient,  dbAccessObj, elasticClient, *monClient, auth)
		server = httptest.NewServer(handler);
		testUrl = server.URL

		//testUrl = config["server.url"]
		res, err := DoGetPing(testUrl + "/v1/ping")

		var user models.User
		user.Username = "admin"
		user.Password = "cfmonit"

		TestToken = res.Token

		data, _ := json.Marshal(user)
		DoPost(testUrl + "/v1/login", res.Token, strings.NewReader(string(data)))
		//fmt.Println(loginRes)
	})

	AfterSuite(func() {
		logoutRes, _ := DoPost(testUrl + "/v1/logout", TestToken, nil)
		assert.Equal(t, http.StatusCreated, logoutRes.Code)
	})
})

func ReadXmlConfig (filename string) (string, error) {
	xmlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}

	return string(xmlFile),  nil

}

func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

func DoGetPing(url string) (*PingResponse, error) {

	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(models.TEST_TOKEN_NAME, models.TEST_TOKEN_VALUE)
	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	token := response.Header.Get(models.TEST_TOKEN_NAME)

	return &PingResponse{Token: string(token), Code: response.StatusCode}, nil
}

func DoGet(url string) (*Response, error) {

	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(models.TEST_TOKEN_NAME, TestToken)
	req.Header.Add("username", "admin")
	req.Header.Add("password", "cfmonit")


	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println("======>", response)
	return &Response{Content: string(contents), Code: response.StatusCode}, nil
}

func DoPost(url, token string, body io.Reader) (*Response, error) {

	client := &http.Client{}

	req, _ := http.NewRequest("POST", url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add(models.TEST_TOKEN_NAME, token)


	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &Response{Content: string(contents), Code: response.StatusCode}, nil
}

func DoLogout(url, testToken string) (*Response, error) {

	client := &http.Client{}

	token, _ := utils.GenerateRandomString(32)

	req, _ := http.NewRequest("POST", url, nil)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add(models.TEST_TOKEN_NAME, token)
	req.Header.Add(models.CSRF_TOKEN_NAME, token)

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &Response{Content: string(contents), Code: response.StatusCode}, nil
}

func DoUpdate(url, token string, body io.Reader) (*Response, error) {

	client := &http.Client{}

	req, _ := http.NewRequest("PUT", url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add(models.TEST_TOKEN_NAME, token)

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &Response{Content: string(contents), Code: response.StatusCode}, nil
}

func DoPatch(url, token string, body io.Reader) (*Response, error) {

	client := &http.Client{}

	req, _ := http.NewRequest("PATCH", url, body)
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Accept", "application/json")
	req.Header.Add(models.TEST_TOKEN_NAME, token)

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &Response{Content: string(contents), Code: response.StatusCode}, nil
}

func DoDelete(url, token string, body io.Reader) (*Response, error) {

	client := &http.Client{}

	req, _ := http.NewRequest("DELETE", url, body)
	req.Header.Add("Accept", "application/json")
	req.Header.Add(models.TEST_TOKEN_NAME, token)

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer response.Body.Close()
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return &Response{Content: string(contents), Code: response.StatusCode}, nil
}

func DoDetail(url, token string, body io.Reader) (*Response, error) {

	client := &http.Client{}

	req, _ := http.NewRequest("GET", url, body)

	req.Header.Add("Accept", "application/json")
	req.Header.Add(models.TEST_TOKEN_NAME, token)

	response, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	//fmt.Println("======>", response)
	return &Response{Content: string(contents), Code: response.StatusCode}, nil
}


var (
	server  *httptest.Server
	testUrl string
	t *testing.T
)

func readConfig(filename string) (Config, error) {
	// init with some bogus data
	config := Config{
		"server.port": "9999",
	}

	if len(filename) == 0 {
		return config, nil
	}
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		// check if the line has = sign
		// and process the line. Ignore the rest.
		if equal := strings.Index(line, "="); equal >= 0 {
			if key := strings.TrimSpace(line[:equal]); len(key) > 0 {
				value := ""
				if len(line) > equal {
					value = strings.TrimSpace(line[equal+1:])
				}
				// assign the config map
				config[key] = value
			}
		}
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}
	}
	return config, nil
}

