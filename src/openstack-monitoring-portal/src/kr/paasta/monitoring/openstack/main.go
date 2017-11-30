package main

import (
	"bufio"
	"io"
	"os"
	"log"
	"strings"
	"strconv"
	"net/http"
	"github.com/influxdata/influxdb/client/v2"
	_ "github.com/go-sql-driver/mysql"
	"kr/paasta/monitoring/openstack/models"
	"kr/paasta/monitoring/openstack/handlers"
	"kr/paasta/monitoring/openstack/utils"
	"fmt"
	"github.com/jinzhu/gorm"
	"io/ioutil"
	"github.com/cihub/seelog"
	"gopkg.in/olivere/elastic.v3"
	"github.com/monasca/golang-monascaclient/monascaclient"
	"github.com/gophercloud/gophercloud"
	"github.com/alexedwards/scs"
	"time"
)

type Config map[string]string

type DBConfig struct {
	DbType string
	UserName string
	UserPassword string
	Host string
	Port string
	DbName string
}

var Logger seelog.LoggerInterface


func main() {

	sessionCookie, _ := utils.GenerateRandomString(32)
	models.SessionManager = *scs.NewCookieManager(sessionCookie) //("u46IpCV9y5Vlur8YvODJEhgOY8m9JVE4")
	models.SessionManager.Lifetime(time.Minute * 30)

	//models.SessionManager.Secure()
	//============================================
	// 기본적인 프로퍼티 설정 정보 읽어오기
	config, err := ReadConfig(`config.ini`)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	xmlFile, err := ReadXmlConfig(`log_config.xml`)
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

	//============================================
	//Server Port
	apiPort, _ := strconv.Atoi(config["server.port"])

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
	/*openstackProvider.Username, _ 	= config["default.username"]
	openstackProvider.Password, _ 	= config["default.password"]*/
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
		/*Username : config["default.username"],
		Password : config["default.password"],*/
		TenantID : config["default.project_id"],
	}
	//monClient.SetKeystoneConfig(&auth)


	//monascaclient.SetKeystoneToken()
	// Route Path 정보와 처리 서비스 연결
	handlers := handlers.NewHandler(openstackProvider, InfluxServerClient,  dbAccessObj, elasticClient, *monClient, auth)
	if err := http.ListenAndServe(fmt.Sprintf(":%v", apiPort), handlers); err != nil {
		log.Fatalln(err)
	}

}


/*func (manager *models.Manager) SessionStart(w http.ResponseWriter, r *http.Request) (session models.Session) {

	manager.lock.Lock()
	defer manager.lock.Unlock()
	cookie, err := r.Cookie(manager.cookieName)
	if err != nil || cookie.Value == "" {
		sid := manager.sessionId()
		session, _ = manager.provider.SessionInit(sid)
		cookie := http.Cookie{Name: manager.cookieName, Value: url.QueryEscape(sid), Path: "/", HttpOnly: true, MaxAge: int(manager.maxlifetime)}
		http.SetCookie(w, &cookie)
	} else {
		sid, _ := url.QueryUnescape(cookie.Value)
		session, _ = manager.provider.SessionRead(sid)
	}
	return
}*/

func UseLogger(newLogger seelog.LoggerInterface) {
	Logger = newLogger
}

func ReadXmlConfig (filename string) (string, error) {
	xmlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return "", err
	}

	return string(xmlFile),  nil

}

// Config 파일 읽어 오기
func ReadConfig(filename string) (Config, error) {
	// init with some bogus data
	config := Config{
		"server.ip":     "127.0.0.1",
		"server.port":   "8888",
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
