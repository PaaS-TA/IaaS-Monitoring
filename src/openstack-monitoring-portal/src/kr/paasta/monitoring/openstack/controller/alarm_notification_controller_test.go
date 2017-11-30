package controller_test

import (
	"net/http"
	"github.com/stretchr/testify/assert"
	//"github.com/monasca/golang-monascaclient/monascaclient/models"
	. "github.com/onsi/ginkgo"
	"encoding/json"
	"strings"
)

type NotificationRequestBody struct {
	Id    	string `json:"id,omitempty"`
	Name    string `json:"name,omitempty"`
	Period  int    `json:"period,omitempty"`
	Type    string `json:"type,omitempty"`
	Address string `json:"address,omitempty"`
}

var notification_id = ""

var _ = Describe("AlarmNotificationController", func() {
	Describe("Alarm Notification", func() {
		Context("project Info", func() {

			It("Notification List", func() {
				res, err := DoGet(testUrl + "/v1/alarm/notification?offset=0&limit=10")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Notification Create Info", func() {
				var query NotificationRequestBody
				query.Name = "testname"
				query.Address = "testmail@gmail.com"

				data, _ := json.Marshal(query)

				res, err := DoPost(testUrl + "/v1/alarm/notification", TestToken, strings.NewReader(string(data)))
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
				notification_id = strings.Replace(res.Content, "\"", "", -1)
			})

			It("Notification Update Info", func() {
				var query NotificationRequestBody
				query.Name = "testname2"
				query.Address = "testmail2@gmail.com"

				data, _ := json.Marshal(query)

				res, err := DoUpdate(testUrl + "/v1/alarm/notification/" + notification_id, TestToken, strings.NewReader(string(data)))

				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Notification Delete Info", func() {
				res, err := DoDelete(testUrl + "/v1/alarm/notification/" + notification_id, TestToken, strings.NewReader(string("")))

				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

		})


	})
})

