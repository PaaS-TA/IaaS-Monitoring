package controller_test

import (
	"github.com/stretchr/testify/assert"
	. "github.com/onsi/ginkgo"
	"net/http"
)


var _ = Describe("LogController", func() {

	Describe("Log Info Get", func() {

		Context("Get", func() {
			It("LogRecent", func() {
				res, err := DoGet(testUrl + "/v1/openstack/log/recent?hostname=controller&logType=log&pageIndex=1&pageItems=10&period=5m")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
				assert.NotEmpty(t, res.Content)
				//logList := &models.LogMessage{}
				//json.Unmarshal([]byte(res.Content), logList)
				//fmt.Println("LogInfo=============>",*logList)
				//assert.True(t, len(*logList) > 0)
			})

			It("LogSpecific", func() {
				res, err := DoGet(testUrl + "/v1/openstack/log/specific?endTime=17:50:00&hostname=controller&logType=log&pageIndex=1&pageItems=50&period=5m&startTime=05:13:00&targetDate=2017-09-26")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
				assert.NotEmpty(t, res.Content)
				//logList := &models.LogMessage{}
				//json.Unmarshal([]byte(res.Content), logList)
				//fmt.Println("LogInfo=============>",*logList)
				//assert.True(t, len(*logList) > 0)
			})
		})

	})
})
