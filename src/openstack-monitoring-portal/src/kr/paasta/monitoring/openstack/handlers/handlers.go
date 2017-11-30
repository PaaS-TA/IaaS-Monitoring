package handlers

import (
	"net/http"
	"github.com/tedsuo/rata"
	"github.com/monasca/golang-monascaclient/monascaclient"
	"kr/paasta/monitoring/openstack/routes"
	"kr/paasta/monitoring/openstack/models"
	client "github.com/influxdata/influxdb/client/v2"
	"kr/paasta/monitoring/openstack/controller"
	"kr/paasta/monitoring/openstack/utils"
	"github.com/gophercloud/gophercloud"
	"github.com/jinzhu/gorm"
	"gopkg.in/olivere/elastic.v3"
	"io"
	"time"
	"fmt"
	"strings"
)

func NewHandler(openstack_provider models.OpenstackProvider, influxClient client.Client,  txn *gorm.DB,
elasticClient *elastic.Client, monsClient monascaclient.Client, auth gophercloud.AuthOptions) http.Handler {

	//Controller선언
	loginController         := controller.NewLoginController(openstack_provider, monsClient, auth)
	mainController 		:= controller.NewMainController(openstack_provider, influxClient)
	computeController 	:= controller.NewComputeController(openstack_provider, influxClient)

	manageNodeController 	:= controller.NewManageNodeController(openstack_provider, influxClient)
	projectController 	:= controller.NewOpenstackProjectController(openstack_provider, influxClient)
	notificationController 	:= controller.NewNotificationController(monsClient, influxClient)
	definitionController 	:= controller.NewAlarmDefinitionController(monsClient, influxClient)

	stautsController 	:= controller.NewAlarmStatusController(monsClient, influxClient, txn)
	logController 	:= controller.NewLogController(openstack_provider, influxClient, elasticClient)

	actions := rata.Handlers{

		routes.PING: route(loginController.Ping),
		routes.LOGIN: route(loginController.Login),
		routes.LOGOUT: route(loginController.Logout),

		//Integrated with routes
		routes.OPENSTACK_RESOURCE_SUMMARY: route(mainController.OpenstackSummary),
		routes.NODE_RESOURCE_SUMMARY: route(computeController.NodeSummary),
		routes.NODE_LIST: route(manageNodeController.GetNodeList),

		routes.NODE_CPU_USAGE_LIST: route(computeController.GetCpuUsageList),
		routes.NODE_CPU_LOAD_LIST: route(computeController.GetCpuLoadList),
		routes.NODE_MEMORY_SWAP_LIST: route(computeController.GetMemorySwapList),
		routes.NODE_MEMORY_USAGE_LIST: route(computeController.GetMemoryUsageList),
		routes.NODE_DISK_USAGE_LIST: route(computeController.GetDiskUsageList),
		routes.NODE_DISK_IO_READ_LIST: route(computeController.GetDiskIoReadList),
		routes.NODE_DISK_IO_WRITE_LIST: route(computeController.GetDiskIoWriteList),
		routes.NODE_NETWORK_IO_KBYTE_LIST: route(computeController.GetNetworkInOutKByteList),
		routes.NODE_NETWORK_ERROR_LIST: route(computeController.GetNetworkInOutErrorList),
		routes.NODE_NETWORK_DROPPED_PACKET_LIST: route(computeController.GetNetworkDroppedPacketList),

		routes.MANAGE_RESOURCE_SUMMARY: route(manageNodeController.ManageNodeSummary),
		routes.MANAGE_RABBIT_MQ_SUMMARY: route(manageNodeController.ManageRabbitMqSummary),
		routes.MANAGE_TOP_PROCESS_BY_CPU: route(manageNodeController.GetTopProcessByCpu),
		routes.MANAGE_TOP_PROCESS_BY_MEMORY: route(manageNodeController.GetTopProcessByMemory),

		routes.ALL_PROJECT_RESOURCE_SUMMARY: route(projectController.ProjectSummary),
		routes.PROJECT_INSTANCE_LIST: route(projectController.GetProjectInstanceList),
		routes.PROJECT_INSTANCE_CPU_USAGE_LIST: route(projectController.GetInstanceCpuUsageList),
		routes.PROJECT_INSTANCE_MEMORY_USAGE_LIST: route(projectController.GetInstanceMemoryUsageList),
		routes.PROJECT_INSTANCE_DISK_READ_LIST: route(projectController.GetInstanceDiskReadList),
		routes.PROJECT_INSTANCE_DISK_WRITE_LIST: route(projectController.GetInstanceDiskWriteList),
		routes.PROJECT_INSTANCE_NETWORK_IO_LIST: route(projectController.GetInstanceNetworkIoList),
		routes.PROJECT_INSTANCE_NETWORK_PACKET_LIST: route(projectController.GetInstanceNetworkPacketsList),

		routes.GET_DEFAULT_RECENT_LOG: route(logController.GetDefaultRecentLog),
		routes.GET_SPECIFIC_TIME_RANGE_LOG: route(logController.GetSpecificTimeRangeLog),

		routes.GET_NOTIFICATION_LIST: route(notificationController.GetAlarmNotificationList),
		routes.CREATE_NOTIFICATION: route(notificationController.CreateAlarmNotification),
		routes.UPDATE_NOTIFICATION: route(notificationController.UpdateAlarmNotification),
		routes.DELETE_NOTIFICATION: route(notificationController.DeleteAlarmNotification),

		routes.GET_ALARM_DEFINITION_LIST: route(definitionController.GetAlarmDefinitionList),
		routes.GET_ALARM_DEFINITION: route(definitionController.GetAlarmDefinition),
		routes.CREATE_ALARM_DEFINITION: route(definitionController.CreateAlarmDefinition),
		routes.PATCH_ALARM_DEFINITION: route(definitionController.UpdateAlarmDefinition),
		routes.DELETE_ALARM_DEFINITION: route(definitionController.DeleteAlarmDefinition),

		routes.ALARM_STATUS_LIST: route(stautsController.GetAlarmStatusList),
		routes.ALARM_STATUS: route(stautsController.GetAlarmStatus),
		routes.ALARM_HISTORY_LIST: route(stautsController.GetAlarmHistoryList),
		routes.ALARM_STATUS_COUNT: route(stautsController.GetAlarmStatusCount),

		routes.GET_ALARM_ACTION_LIST: route(stautsController.GetAlarmHistoryActionList),
		routes.CREATE_ALARM_ACTION: route(stautsController.CreateAlarmHistoryAction),
		routes.UPDATE_ALARM_ACTION: route(stautsController.UpdateAlarmHistoryAction),
		routes.DELETE_ALARM_ACTION: route(stautsController.DeleteAlarmHistoryAction),

		// Html
		routes.Main: route(mainController.Main),
		routes.Static: route(StaticHandler),
	}

	handler, err := rata.NewRouter(routes.Routes, actions)
	if err != nil {
		panic("unable to create router: " + err.Error())
	}
	fmt.Println("Monit Application Started")
	return HttpWrap(handler)
}


func HttpWrap(handler http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if origin := r.Header.Get("Origin"); origin != "" {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, PATCH, DELETE")
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, X-XSRF-TOKEN, Accept-Encoding, Authorization")
			w.Header().Set("Access-Control-Allow-Credentials", "true")
			w.Header().Set("Access-Control-Expose-Headers", "X-XSRF-TOKEN")
		}
		// Stop here if its Preflighted OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}

		if r.RequestURI != "/v1/login" && r.RequestURI != "/v1/ping" && r.RequestURI != "/" && !strings.Contains(r.RequestURI, "/public/") {

			session := models.SessionManager.Load(r)

			//fmt.Println( "=======+>>>>>>>>>", session.Exists())
			reqToken := r.Header.Get(models.CSRF_TOKEN_NAME)
			testToken    := r.Header.Get(models.TEST_TOKEN_NAME)

			userSession := new(models.UserSession)
			//fmt.Println("Handler TOken:", reqToken)
			err := session.GetObject(models.USER_SESSION_NAME + reqToken, userSession)

			if err != nil {
				models.MonitLogger.Debug("LoginSession Obj Err==>", err)
			}

			/*fmt.Println("userSession.CsrfToken::", userSession.CsrfToken)
			fmt.Println("reqToken::", reqToken)*/

			if userSession.CsrfToken != reqToken || (reqToken == "" && testToken == "") {
				errMessage := models.ErrMessage{
					"Message": "UnAuthrized",
				}
				utils.RenderJsonUnAuthResponse(errMessage, http.StatusUnauthorized, w)
			} else {
				handler.ServeHTTP(w, r)
			}
		}else{

			handler.ServeHTTP(w, r)
		}

	}
}


func route(f func(w http.ResponseWriter, r *http.Request)) http.Handler {
	return http.HandlerFunc(f)
}

const STATIC_URL string = "/public/"
const STATIC_ROOT string = "public/"

func StaticHandler(w http.ResponseWriter, req *http.Request) {
	static_file := req.URL.Path[len(STATIC_URL):]
	if len(static_file) != 0 {
		f, err := http.Dir(STATIC_ROOT).Open(static_file)
		if err == nil {
			content := io.ReadSeeker(f)
			http.ServeContent(w, req, static_file, time.Now(), content)
			return
		}
	}
	http.NotFound(w, req)
}