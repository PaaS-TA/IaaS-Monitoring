package controller_test

import (
	"net/http"
	"github.com/stretchr/testify/assert"
	//"github.com/monasca/golang-monascaclient/monascaclient/models"
	. "github.com/onsi/ginkgo"
	"strings"
	"encoding/json"
)

type DefinitionRequestBody struct {
	Id	                string   `json:"id,omitempty"`
	Name                string   `json:"name,omitempty"`
	Description         string   `json:"description,omitempty"`
	Severity            string   `json:"severity,omitempty"`
	MatchBy             []string `json:"match_by,omitempty"`
	Expression          string   `json:"expression,omitempty"`
	ActionsEnabled      bool     `json:"actions_enabled,omitempty"`
	AlarmActions        []string `json:"alarm_actions,omitempty"`
	OkActions           []string `json:"ok_actions,omitempty"`
	UndeterminedActions []string `json:"undetermined_actions,omitempty"`
}

var definition_id = ""

var _ = Describe("AlarmDefinitionController", func() {
	Describe("Alarm Definition", func() {
		Context("project Info", func() {

			It("Definition List", func() {
				res, err := DoGet(testUrl + "/v1/alarm/definition?name=monasca&severity=HIGH&offset=0&limit=10")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Definition Create Info", func() {
				var query DefinitionRequestBody

				query.Name = "monasca_cpu_row"
				query.Severity = "LOW"
				query.MatchBy = []string{"hostname"}
				query.Expression = "max(cpu.utilization_norm_perc{hostname=monasca}) > 60"
//				query.ActionsEnabled = true
				query.AlarmActions = []string{"76efea3e-3ac3-46cb-905c-fc0161dcb48c"}
				query.Description = "Test Description"

				data, _ := json.Marshal(query)

				res, err := DoPost(testUrl + "/v1/alarm/definition", TestToken, strings.NewReader(string(data)))
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Definition Detail Info", func() {
				definition_id = "b6bf8d9f-2bdf-4886-b412-5df0586a8901"

				res, err := DoDetail(testUrl + "/v1/alarm/definition/"+definition_id, TestToken, strings.NewReader(string("")))

				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Definition Patch Info", func() {
				definition_id = "dcb45301-e871-4d9f-a6c4-633ee050094a"

				var query DefinitionRequestBody
				query.Expression = "max(cpu.utilization_norm_perc{hostname=monasca}) > 80"
				query.Description = "Test Description Test Update _ " + definition_id

				data, _ := json.Marshal(query)

				res, err := DoPatch(testUrl + "/v1/alarm/definition/" + definition_id, TestToken, strings.NewReader(string(data)))

				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Definition Delete Info", func() {
				definition_id = "10e43f5d-2863-4c6d-882d-a00f5da847b1"
				res, err := DoDelete(testUrl + "/v1/alarm/definition/" + definition_id, TestToken, strings.NewReader(string("")))

				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

		})


	})
})

