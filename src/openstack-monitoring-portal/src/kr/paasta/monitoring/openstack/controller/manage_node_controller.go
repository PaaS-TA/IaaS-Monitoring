package controller

import (
	"kr/paasta/monitoring/openstack/models"
	"kr/paasta/monitoring/openstack/utils"
	"kr/paasta/monitoring/openstack/services"
	client "github.com/influxdata/influxdb/client/v2"
	"net/http"
)

//Main Page Controller
type OpenstackManageNode struct{
	OpenstackProvider models.OpenstackProvider
	influxClient 	client.Client
}

func NewManageNodeController(openstackProvider models.OpenstackProvider, influxClient client.Client) *OpenstackManageNode {
	return &OpenstackManageNode{
		OpenstackProvider: openstackProvider,
		influxClient: influxClient,
	}
}

func (s *OpenstackManageNode)ManageNodeSummary(w http.ResponseWriter, r *http.Request) {

	var apiRequest models.NodeReq
	apiRequest.HostName = r.URL.Query().Get("hostname")
	provider, _, _ := utils.GetOpenstackProvider(r)
	manageNodeSummary, err := services.GetManageNodeService(s.OpenstackProvider, provider, s.influxClient).GetManageNodeSummary(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	} else {
		utils.RenderJsonResponse(manageNodeSummary, w)
	}
}


func (s *OpenstackManageNode)GetNodeList(w http.ResponseWriter, r *http.Request) {

	provider, _, err := utils.GetOpenstackProvider(r)

	nodeList, err := services.GetManageNodeService(s.OpenstackProvider, provider, s.influxClient).GetNodeList()
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	} else {
		utils.RenderJsonResponse(nodeList, w)
	}
}

func (s *OpenstackManageNode)ManageRabbitMqSummary(w http.ResponseWriter, r *http.Request) {

	provider, _, _ := utils.GetOpenstackProvider(r)
	manageNodeSummary, err := services.GetManageNodeService(s.OpenstackProvider, provider, s.influxClient).GetRabbitMqSummary()
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	} else {
		utils.RenderJsonResponse(manageNodeSummary, w)
	}
}


func (s *OpenstackManageNode)GetTopProcessByCpu(w http.ResponseWriter, r *http.Request) {

	var apiRequest models.DetailReq
	apiRequest.HostName = r.FormValue(":hostname")
	provider, _, _ := utils.GetOpenstackProvider(r)

	topProcess, err := services.GetManageNodeService(s.OpenstackProvider, provider, s.influxClient).GetTopProcessListByCpu(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	} else {
		utils.RenderJsonResponse(topProcess, w)
	}
}

func (s *OpenstackManageNode)GetTopProcessByMemory(w http.ResponseWriter, r *http.Request) {

	var apiRequest models.DetailReq
	apiRequest.HostName = r.FormValue(":hostname")
	provider, _, _ := utils.GetOpenstackProvider(r)
	topProcess, err := services.GetManageNodeService(s.OpenstackProvider, provider, s.influxClient).GetTopProcessListByMemory(apiRequest)
	if err != nil {
		utils.ErrRenderJsonResponse(err, w)
	} else {
		utils.RenderJsonResponse(topProcess, w)
	}
}