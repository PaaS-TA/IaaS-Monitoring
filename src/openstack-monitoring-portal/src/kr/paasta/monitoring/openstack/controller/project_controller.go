package controller

import (
	"kr/paasta/monitoring/openstack/models"
	client "github.com/influxdata/influxdb/client/v2"
	"kr/paasta/monitoring/openstack/utils"
	"kr/paasta/monitoring/openstack/services"
	"net/http"
)
//Project Controller
type OpenstackProject struct{
	openstackProvider models.OpenstackProvider
	influxClient 	client.Client
}

func NewOpenstackProjectController(openstackProvider models.OpenstackProvider, influxClient client.Client) *OpenstackProject {
	return &OpenstackProject{
		openstackProvider: openstackProvider,
		influxClient: influxClient,
	}
}

func (s *OpenstackProject)ProjectSummary(w http.ResponseWriter, r *http.Request){

	//projectName은 조회조건 (Optional)
	var apiRequest models.ProjectReq
	apiRequest.ProjectName   = r.URL.Query().Get("projectName")

	provider, username, _ := utils.GetOpenstackProvider(r)
	s.openstackProvider.Username = username
	projectSummary, err := services.GetProjectService(s.openstackProvider, provider, s.influxClient).GetProjectSummary(apiRequest)

	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(projectSummary, w)
	}
}

func (s *OpenstackProject)GetProjectInstanceList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.ProjectReq
	apiRequest.ProjectId  = r.FormValue(":projectId")
	//hostname은 조회조건 (Optional)
	apiRequest.HostName   = r.URL.Query().Get("hostname")
	//Paging Size (Optional)
	apiRequest.Limit      = r.URL.Query().Get("limit")
	//Paging 처리시 현재 Page Limit의 마지막 Instance Id를 요청 받으면 다음 Page를 조회 할 수 있다.
	//Limit과 같이 사용되어야 함 (Optional)
	apiRequest.Marker     = r.URL.Query().Get("marker")

	validation := apiRequest.ProjectInstanceRequestValidate(apiRequest)

	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}

	provider, _, _ := utils.GetOpenstackProvider(r)
	projectSummary, err := services.GetProjectService(s.openstackProvider, provider, s.influxClient).GetProjectInstanceList(apiRequest)

	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(projectSummary, w)
	}
}


func (s *OpenstackProject)GetInstanceCpuUsageList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.InstanceId = r.FormValue(":instanceId")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.InstanceMetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _, _ := utils.GetOpenstackProvider(r)
	cpuUsageList, err := services.GetProjectService(s.openstackProvider, provider, s.influxClient).GetInstanceCpuUsageList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(cpuUsageList, w)
	}
}


func (s *OpenstackProject)GetInstanceMemoryUsageList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.InstanceId = r.FormValue(":instanceId")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.InstanceMetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _, _ := utils.GetOpenstackProvider(r)
	cpuUsageList, err := services.GetProjectService(s.openstackProvider, provider, s.influxClient).GetInstanceMemoryUsageList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(cpuUsageList, w)
	}
}


func (s *OpenstackProject)GetInstanceDiskReadList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.InstanceId = r.FormValue(":instanceId")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.InstanceMetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}

	provider, _, _ := utils.GetOpenstackProvider(r)
	cpuUsageList, err := services.GetProjectService(s.openstackProvider, provider,  s.influxClient).GetInstanceDiskIoKbyteList(apiRequest, "read")
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(cpuUsageList, w)
	}
}

func (s *OpenstackProject)GetInstanceDiskWriteList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.InstanceId = r.FormValue(":instanceId")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.InstanceMetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _, _ := utils.GetOpenstackProvider(r)
	cpuUsageList, err := services.GetProjectService(s.openstackProvider, provider, s.influxClient).GetInstanceDiskIoKbyteList(apiRequest, "write")
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(cpuUsageList, w)
	}
}


func (s *OpenstackProject)GetInstanceNetworkIoList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.InstanceId = r.FormValue(":instanceId")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.InstanceMetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _, _ := utils.GetOpenstackProvider(r)
	cpuUsageList, err := services.GetProjectService(s.openstackProvider, provider, s.influxClient).GetInstanceNetworkIoKbyteList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(cpuUsageList, w)
	}
}

func (s *OpenstackProject)GetInstanceNetworkPacketsList(w http.ResponseWriter, r *http.Request){

	var apiRequest models.DetailReq
	apiRequest.InstanceId = r.FormValue(":instanceId")
	apiRequest.DefaultTimeRange = r.URL.Query().Get("defaultTimeRange")
	apiRequest.TimeRangeFrom = r.URL.Query().Get("timeRangeFrom")
	apiRequest.TimeRangeTo   = r.URL.Query().Get("timeRangeTo")
	apiRequest.GroupBy       = r.URL.Query().Get("groupBy")


	validation :=  apiRequest.InstanceMetricRequestValidate(apiRequest )
	if validation != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(validation.Error()))
		return
	}
	provider, _, _ := utils.GetOpenstackProvider(r)
	cpuUsageList, err := services.GetProjectService(s.openstackProvider, provider, s.influxClient).GetInstanceNetworkPacketsList(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(cpuUsageList, w)
	}
}