package controller

import (
	client "github.com/influxdata/influxdb/client/v2"
	"net/http"
	"kr/paasta/monitoring/openstack/models"
	"kr/paasta/monitoring/openstack/utils"
	"kr/paasta/monitoring/openstack/services"
	"fmt"
)

//Main Page Controller
type OpenstackServices struct{
	OpenstackProvider models.OpenstackProvider
	influxClient 	client.Client
}


func NewMainController(openstackProvider models.OpenstackProvider, influxClient client.Client) *OpenstackServices {
	return &OpenstackServices{
		OpenstackProvider: openstackProvider,
		influxClient: influxClient,
	}
}

func (h *OpenstackServices) Main(w http.ResponseWriter, r *http.Request) {
	models.MonitLogger.Debug("Main API Called")

	url := "/public/dist/index.html"
	http.Redirect(w, r, url, 302)
}

func (s *OpenstackServices)OpenstackSummary(w http.ResponseWriter, r *http.Request){

	fmt.Println("------------")
	provider, username, err := utils.GetOpenstackProvider(r)
	projectResourceSummary, err := services.GetMainService(s.OpenstackProvider, provider, s.influxClient).GetOpenstackSummary(username)

	if err != nil {
		models.MonitLogger.Error("GetOpenstackResources error :", err)
		utils.ErrRenderJsonResponse(err, w)
	}else{
		utils.RenderJsonResponse(projectResourceSummary, w)
	}

}
