package controller

import (
	//"kr/paasta/monitoring/openstack/services"
	"kr/paasta/monitoring/openstack/models"
	"kr/paasta/monitoring/openstack/utils"
	client "github.com/influxdata/influxdb/client/v2"
	"net/http"
	"kr/paasta/monitoring/openstack/services"
)

//Compute Node Controller
type OpenstackComputeNode struct{
	OpenstackProvider models.OpenstackProvider
	influxClient 	client.Client
}

func NewComputeController(openstackProvider models.OpenstackProvider, influxClient client.Client) *OpenstackComputeNode {
	return &OpenstackComputeNode{
		OpenstackProvider: openstackProvider,
		influxClient: influxClient,
	}
}

func (s *OpenstackComputeNode)NodeSummary(w http.ResponseWriter, r *http.Request){

	var apiRequest models.NodeReq
	apiRequest.HostName = r.URL.Query().Get("hostname")
	provider, _, _ := utils.GetOpenstackProvider(r)

	computeNodeSummary, err := services.GetComputeNodeService(s.OpenstackProvider, provider,  s.influxClient).GetComputeNodeSummary(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(computeNodeSummary, w)
	}
}



func (s *OpenstackComputeNode)GetCpuUsageList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.HostName = r.FormValue(":hostname")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.MetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _, _ := utils.GetOpenstackProvider(r)
	cpuUsageList, err := services.GetComputeNodeService(s.OpenstackProvider, provider,  s.influxClient).GetComputeNodeCpuUsageList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(cpuUsageList, w)
	}
}


func (s *OpenstackComputeNode)GetCpuLoadList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.HostName = r.FormValue(":hostname")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.MetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _,  _ := utils.GetOpenstackProvider(r)
	cpuUsageList, err := services.GetComputeNodeService(s.OpenstackProvider, provider, s.influxClient).GetComputeNodeCpuLoad1mList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(cpuUsageList, w)
	}
}

//Memory 사용률
func (s *OpenstackComputeNode)GetMemoryUsageList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.HostName = r.FormValue(":hostname")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.MetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _,  _ := utils.GetOpenstackProvider(r)
	cpuUsageList, err := services.GetComputeNodeService(s.OpenstackProvider, provider, s.influxClient).GetComputeNodeMemoryUsageList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(cpuUsageList, w)
	}
}

//Memory Swap 사용률
func (s *OpenstackComputeNode)GetMemorySwapList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.HostName = r.FormValue(":hostname")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.MetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _, _ := utils.GetOpenstackProvider(r)
	cpuUsageList, err := services.GetComputeNodeService(s.OpenstackProvider, provider, s.influxClient).GetComputeNodeSwapUsageList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(cpuUsageList, w)
	}
}


//Disk 사용률(Mountpoint)
func (s *OpenstackComputeNode)GetDiskUsageList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.HostName = r.FormValue(":hostname")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.MetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _, _ := utils.GetOpenstackProvider(r)
	diskUsageList, err := services.GetComputeNodeService(s.OpenstackProvider, provider, s.influxClient).GetNodeDiskUsageList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(diskUsageList, w)
	}
}



//Disk IO Raad Kbyte(Mountpoint)
func (s *OpenstackComputeNode)GetDiskIoReadList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.HostName = r.FormValue(":hostname")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.MetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _, _ := utils.GetOpenstackProvider(r)

	diskUsageList, err := services.GetComputeNodeService(s.OpenstackProvider, provider,  s.influxClient).GetNodeDiskIoReadList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(diskUsageList, w)
	}
}

//Disk IO Write Kbyte(Mountpoint)
func (s *OpenstackComputeNode)GetDiskIoWriteList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.HostName = r.FormValue(":hostname")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.MetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}

	provider, _, _ := utils.GetOpenstackProvider(r)
	diskUsageList, err := services.GetComputeNodeService(s.OpenstackProvider, provider,  s.influxClient).GetNodeDiskIoWriteList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(diskUsageList, w)
	}
}


//Disk IO Write Kbyte(Mountpoint)
func (s *OpenstackComputeNode)GetNetworkInOutKByteList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.HostName = r.FormValue(":hostname")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.MetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _, _ := utils.GetOpenstackProvider(r)
	networkUsageList, err := services.GetComputeNodeService(s.OpenstackProvider, provider, s.influxClient).GetNodeNetworkInOutKByteList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(networkUsageList, w)
	}
}

//Disk IO Write Kbyte(Mountpoint)
func (s *OpenstackComputeNode)GetNetworkInOutErrorList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.HostName = r.FormValue(":hostname")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.MetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _, _ := utils.GetOpenstackProvider(r)

	networkUsageList, err := services.GetComputeNodeService(s.OpenstackProvider, provider, s.influxClient).GetNodeNetworkInOutErrorList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(networkUsageList, w)
	}
}


//Disk IO Write Kbyte(Mountpoint)
func (s *OpenstackComputeNode)GetNetworkDroppedPacketList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.HostName = r.FormValue(":hostname")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.MetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _, _ := utils.GetOpenstackProvider(r)
	networkUsageList, err := services.GetComputeNodeService(s.OpenstackProvider, provider, s.influxClient).GetNodeNetworkDropPacketList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(networkUsageList, w)
	}
}

