package controller_test

import (
	"github.com/stretchr/testify/assert"
	. "github.com/onsi/ginkgo"
	"net/http"
	"kr/paasta/monitoring/openstack/models"
	"encoding/json"
)


var _ = Describe("ProjectController", func() {

	Describe("Project Summary", func() {
		var instanceId string
		Context("project Info", func() {
			var projectId string
			It("Project Summry", func() {
				res, err := DoGet(testUrl + "/v1/openstack/projects/summary")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
				assert.NotEmpty(t, res.Content)

				projectInfoList := &[]models.ProjectSummaryInfo{}
				json.Unmarshal([]byte(res.Content), projectInfoList)
				assert.True(t, len((*projectInfoList)) > 0)
				projectId = (*projectInfoList)[0].Id
			})

			It("Get Instance From Project", func() {
				res, err := DoGet(testUrl + "/v1/openstack/projects/" + projectId + "/instances?limit=1")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
				assert.NotEmpty(t, res.Content)
				//Response Data형변환
				var msgMapTemplate interface{}
				json.Unmarshal([]byte(res.Content), &msgMapTemplate)
				msgMap := msgMapTemplate.(map[string]interface{})
				tmp := msgMap["metric"].([]interface{})
				data := tmp[0].(map[string]interface{})

				assert.True(t, data["instance_id"] != "")
				instanceId = data["instance_id"].(string)
			})
		})
		Context("Instance Metric Detail", func() {
			It("Get Instance Cpu Usage", func() {
				res, err := DoGet(testUrl + "/v1/openstack/projects/" + instanceId + "/cpuUsage?defaultTimeRange=10m&groupBy=1m")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
				assert.NotEmpty(t, res.Content)
			})

			It("Get Instance Memory Usage", func() {
				res, err := DoGet(testUrl + "/v1/openstack/projects/" + instanceId + "/memUsage?defaultTimeRange=10m&groupBy=1m")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
				assert.NotEmpty(t, res.Content)
			})

			It("Get Instance diskRead", func() {
				res, err := DoGet(testUrl + "/v1/openstack/projects/" + instanceId + "/diskRead?defaultTimeRange=10m&groupBy=1m")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
				assert.NotEmpty(t, res.Content)
			})

			It("Get Instance diskWrite", func() {
				res, err := DoGet(testUrl + "/v1/openstack/projects/" + instanceId + "/diskWrite?defaultTimeRange=10m&groupBy=1m")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
				assert.NotEmpty(t, res.Content)
			})

			It("Get Instance networkIo", func() {
				res, err := DoGet(testUrl + "/v1/openstack/projects/" + instanceId + "/networkIo?defaultTimeRange=10m&groupBy=1m")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
				assert.NotEmpty(t, res.Content)
			})

			It("Get Instance networkPackets", func() {
				res, err := DoGet(testUrl + "/v1/openstack/projects/" + instanceId + "/networkPackets?defaultTimeRange=10m&groupBy=1m")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
				assert.NotEmpty(t, res.Content)
			})

		})


	})

})

