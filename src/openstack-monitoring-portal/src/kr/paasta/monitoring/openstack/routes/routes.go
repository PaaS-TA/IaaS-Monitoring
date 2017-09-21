package routes

import "github.com/tedsuo/rata"

const (
	PING = "PING"
	LOGIN = "LOGIN API"
	LOGOUT = "LOGIN OUT"

	OPENSTACK_RESOURCE_SUMMARY = "OPENSTACK RESOURCE SUMMARY API"
	NODE_RESOURCE_SUMMARY = "OPENSTACK COMPUTE NODE SUMMARY API"
	NODE_LIST = "OPENSTACK NODE_LIST API"

	NODE_CPU_USAGE_LIST = "OPENSTACK NODE CPU USAGE LIST API"
	NODE_CPU_LOAD_LIST = "OPENSTACK NODE CPU Load 1Minute LIST API"
	NODE_MEMORY_SWAP_LIST = "OPENSTACK NODE MEMORY SWAP LIST API"
	NODE_MEMORY_USAGE_LIST = "OPENSTACK NODE MEMORY USAGE LIST API"
	NODE_DISK_USAGE_LIST = "OPENSTACK NODE Disk USAGE LIST API"
	NODE_DISK_IO_READ_LIST = "OPENSTACK NODE DISK IO READ KBYTE LIST API"
	NODE_DISK_IO_WRITE_LIST = "OPENSTACK NODE DISK IO WRITE KBYTE LIST API"
	NODE_NETWORK_IO_KBYTE_LIST = "OPENSTACK NODE NETWORK IO KBYTE LIST API"
	NODE_NETWORK_ERROR_LIST = "OPENSTACK NODE NETWORK IO ERROR LIST API"
	NODE_NETWORK_DROPPED_PACKET_LIST = "OPENSTACK NODE NETWORK IN DROP PACKET LIST API"
	MANAGE_RABBIT_MQ_SUMMARY = "RABBIT MQ RESOURCE SUMMARY API"
	MANAGE_RESOURCE_SUMMARY = "OPENSTACK MANAGE NODE SUMMARY API"
	MANAGE_TOP_PROCESS_BY_CPU = "OPENSTACK TOP PROCESS BY CPU API"
	MANAGE_TOP_PROCESS_BY_MEMORY = "OPENSTACK TOP PROCESS BY MEMORY API"

	ALL_PROJECT_RESOURCE_SUMMARY = "OPENSTACK ALL PROJECT RESOURCE SUMMARY"
	PROJECT_INSTANCE_LIST = "GET OPENSTACK INSTANCE LIST BY PROJECT ID"

	PROJECT_INSTANCE_CPU_USAGE_LIST   = "GET OPENSTACK INSTANCE CPU USAGE LIST"
	PROJECT_INSTANCE_MEMORY_USAGE_LIST = "GET OPENSTACK INSTANCE MEMORY USAGE LIST"
	PROJECT_INSTANCE_DISK_READ_LIST   = "GET OPENSTACK INSTANCE DISK READ LIST"
	PROJECT_INSTANCE_DISK_WRITE_LIST  = "GET OPENSTACK INSTANCE DISK WRITE LIST"
	PROJECT_INSTANCE_NETWORK_IO_LIST  = "GET OPENSTACK INSTANCE NETWORK IO BYTE LIST"
	PROJECT_INSTANCE_NETWORK_PACKET_LIST = "GET OPENSTACK INSTANCE NETWORK PACKET LIST"

	GET_DEFAULT_RECENT_LOG = "GET_DEFAULT_RECENT_LOG"
	GET_SPECIFIC_TIME_RANGE_LOG = "GET_SPECIFIC_TIME_RANGE_LOG"

	GET_NOTIFICATION_LIST = "GET_NOTIFICATION_LIST"
	CREATE_NOTIFICATION = "CREATE_NOTIFICATION"
	DELETE_NOTIFICATION = "DELETE_NOTIFICATION"
	UPDATE_NOTIFICATION = "UPDATE_NOTIFICATION"

	GET_ALARM_DEFINITION_LIST = "GET ALARM DEFINITION LIST"
	GET_ALARM_DEFINITION = "GET ALARM DEFINITION"
	CREATE_ALARM_DEFINITION = "CREATE ALARM DEFINITION"
	PATCH_ALARM_DEFINITION = "PATCH ALARM DEFINITION"
	DELETE_ALARM_DEFINITION = "DELETE ALARM DEFINITION"

	ALARM_STATUS_COUNT = "ALARM_STATUS_COUNT"
	ALARM_STATUS_LIST = "ALARM_STATUS_LIST"
	ALARM_STATUS = "ALARM_STATUS"
	ALARM_HISTORY_LIST = "ALARM_HISTORY_LIST"

	GET_ALARM_ACTION_LIST = "GET_ALARM_ACTION_LIST"
	CREATE_ALARM_ACTION = "CREATE_ALARM_ACTION"
	UPDATE_ALARM_ACTION = "UPDATE_ALARM_ACTION"
	DELETE_ALARM_ACTION = "DELETE_ALARM_ACTION"


	// Web Resource
	Main = "Main"
	Static = "Static"
)

var Routes = rata.Routes{

	{Path: "/v1/ping", Method: "GET", Name: PING},
	{Path: "/v1/login", Method: "POST", Name: LOGIN},
	{Path: "/v1/logout", Method: "POST", Name: LOGOUT},

	{Path: "/v1/openstack/summary", Method: "GET", Name: OPENSTACK_RESOURCE_SUMMARY},
	{Path: "/v1/openstack/computeNode/summary", Method: "GET", Name: NODE_RESOURCE_SUMMARY},

	{Path: "/v1/openstack/node", Method: "GET", Name: NODE_LIST},

	{Path: "/v1/openstack/node/:hostname/cpuUsage", Method: "GET", Name: NODE_CPU_USAGE_LIST},
	{Path: "/v1/openstack/node/:hostname/cpuLoad", Method: "GET", Name: NODE_CPU_LOAD_LIST},
	{Path: "/v1/openstack/node/:hostname/swapUsage", Method: "GET", Name: NODE_MEMORY_SWAP_LIST},
	{Path: "/v1/openstack/node/:hostname/memUsage", Method: "GET", Name: NODE_MEMORY_USAGE_LIST},
	{Path: "/v1/openstack/node/:hostname/diskUsage", Method: "GET", Name: NODE_DISK_USAGE_LIST},
	{Path: "/v1/openstack/node/:hostname/diskRead", Method: "GET", Name: NODE_DISK_IO_READ_LIST},
	{Path: "/v1/openstack/node/:hostname/diskWrite", Method: "GET", Name: NODE_DISK_IO_WRITE_LIST},
	{Path: "/v1/openstack/node/:hostname/networkIo", Method: "GET", Name: NODE_NETWORK_IO_KBYTE_LIST},
	{Path: "/v1/openstack/node/:hostname/networkError", Method: "GET", Name: NODE_NETWORK_ERROR_LIST},
	{Path: "/v1/openstack/node/:hostname/networkDropPacket", Method: "GET", Name: NODE_NETWORK_DROPPED_PACKET_LIST},

	//Manage Node Api
	{Path: "/v1/openstack/manageNode/summary", Method: "GET", Name: MANAGE_RESOURCE_SUMMARY},
	{Path: "/v1/openstack/manageNode/rabbitMqSummary", Method: "GET", Name: MANAGE_RABBIT_MQ_SUMMARY},
	{Path: "/v1/openstack/manageNode/:hostname/topProcessCpu", Method: "GET", Name: MANAGE_TOP_PROCESS_BY_CPU},
	{Path: "/v1/openstack/manageNode/:hostname/topProcessMem", Method: "GET", Name: MANAGE_TOP_PROCESS_BY_MEMORY},


	//Project Api
	{Path: "/v1/openstack/projects/summary", Method: "GET", Name: ALL_PROJECT_RESOURCE_SUMMARY},
	{Path: "/v1/openstack/projects/:projectId/instances", Method: "GET", Name: PROJECT_INSTANCE_LIST},
	{Path: "/v1/openstack/projects/:instanceId/cpuUsage", Method: "GET", Name: PROJECT_INSTANCE_CPU_USAGE_LIST},
	{Path: "/v1/openstack/projects/:instanceId/memUsage", Method: "GET", Name: PROJECT_INSTANCE_MEMORY_USAGE_LIST},
	{Path: "/v1/openstack/projects/:instanceId/diskRead", Method: "GET", Name: PROJECT_INSTANCE_DISK_READ_LIST},
	{Path: "/v1/openstack/projects/:instanceId/diskWrite", Method: "GET", Name: PROJECT_INSTANCE_DISK_WRITE_LIST},
	{Path: "/v1/openstack/projects/:instanceId/networkIo", Method: "GET", Name: PROJECT_INSTANCE_NETWORK_IO_LIST},
	{Path: "/v1/openstack/projects/:instanceId/networkPackets", Method: "GET", Name: PROJECT_INSTANCE_NETWORK_PACKET_LIST},

	// Log Api
	{Path: "/v1/openstack/log/recent", Method: "GET", Name: GET_DEFAULT_RECENT_LOG},
	{Path: "/v1/openstack/log/specific", Method: "GET", Name: GET_SPECIFIC_TIME_RANGE_LOG},


	{Path: "/v1/alarm/notification", Method: "GET", Name: GET_NOTIFICATION_LIST},
	{Path: "/v1/alarm/notification", Method: "POST", Name: CREATE_NOTIFICATION},
	{Path: "/v1/alarm/notification/:id", Method: "PUT", Name: UPDATE_NOTIFICATION},
	{Path: "/v1/alarm/notification/:id", Method: "DELETE", Name: DELETE_NOTIFICATION},

	{Path: "/v1/alarm/definition", Method: "GET", Name: GET_ALARM_DEFINITION_LIST},
	{Path: "/v1/alarm/definition/:id", Method: "GET", Name: GET_ALARM_DEFINITION},
	{Path: "/v1/alarm/definition", Method: "POST", Name: CREATE_ALARM_DEFINITION},
	{Path: "/v1/alarm/definition/:id", Method: "PATCH", Name: PATCH_ALARM_DEFINITION},
	{Path: "/v1/alarm/definition/:id", Method: "DELETE", Name: DELETE_ALARM_DEFINITION},

	{Path: "/v1/alarm/status/count", Method: "GET", Name: ALARM_STATUS_COUNT},
	{Path: "/v1/alarm/status", Method: "GET", Name: ALARM_STATUS_LIST},
	{Path: "/v1/alarm/:alarmId/status", Method: "GET", Name: ALARM_STATUS},
	{Path: "/v1/alarm/:alarmId/history", Method: "GET", Name: ALARM_HISTORY_LIST},

	{Path: "/v1/alarm/:alarmId/action", Method: "GET", Name: GET_ALARM_ACTION_LIST},
	{Path: "/v1/alarm/action", Method: "POST", Name: CREATE_ALARM_ACTION},
	{Path: "/v1/alarm/action/:id", Method: "PUT", Name: UPDATE_ALARM_ACTION},
	{Path: "/v1/alarm/action/:id", Method: "DELETE", Name: DELETE_ALARM_ACTION},

	// Web Resource
	{Path: "/", Method: "GET", Name: Main},
	{Path: "/public/", Method: "GET", Name: Static},
}
