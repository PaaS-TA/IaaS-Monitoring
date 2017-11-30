package controller_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"net/http"
	"strings"
	"encoding/json"
)

type StatusRequestBody struct {
	Id                uint
	AlarmId           string
	AlarmActionDesc   string
}
var alramId = "32"
var alramDefinitionId = "f62a847c-cc62-40c6-b31c-6291a9480033"

var _ = Describe("AlarmStatusController", func() {
	Describe("Alarm Status", func() {
		Context("Status", func() {

			It("Alarm Status List", func() {
				res, err := DoGet(testUrl + "/v1/alarm/status?severity=HIGH&state=OK&offset=0&limit=10")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Alarm Status List Count", func() {
				res, err := DoGet(testUrl + "/v1/alarm/status/count?state=OK")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Alarm Status Detail", func() {
				res, err := DoGet(testUrl + "/v1/alarm/"+alramDefinitionId+"/status")

				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Alarm Status History List", func() {
				res, err := DoGet(testUrl + "/v1/alarm/"+alramDefinitionId+"/history")

				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Alarm Status Action List", func() {
				res, err := DoGet(testUrl + "/v1/alarm/"+alramDefinitionId+"/action")

				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Alarm Status Create Info", func() {
				var query StatusRequestBody
				query.AlarmId = alramDefinitionId
				query.AlarmActionDesc = "Test Alram Action Desc!!"

				data, _ := json.Marshal(query)

				res, err := DoPost(testUrl + "/v1/alarm/action", TestToken, strings.NewReader(string(data)))

				assert.Nil(t, err)
				assert.Equal(t, http.StatusCreated, res.Code)
			})

			It("Alarm Status Update Info", func() {
				var query StatusRequestBody
				query.Id = 12
				query.AlarmActionDesc = "Update Test Alram Action Desc!! : " + alramId

				data, _ := json.Marshal(query)

				res, err := DoUpdate(testUrl + "/v1/alarm/action/" + alramId, TestToken, strings.NewReader(string(data)))

				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Notification Delete Info", func() {
				res, err := DoDelete(testUrl + "/v1/alarm/action/" + alramId, TestToken, strings.NewReader(string("")))

				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

		})
	})
})
