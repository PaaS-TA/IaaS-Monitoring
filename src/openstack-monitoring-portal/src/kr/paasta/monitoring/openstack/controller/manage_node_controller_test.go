package controller_test

import (
	"github.com/stretchr/testify/assert"
	. "github.com/onsi/ginkgo"
	"net/http"
)

var _ = Describe("ManageNodeController", func() {

	Describe("ManageNode", func() {
		Context("GET", func() {
			It("Manage Node Summary", func() {
				res, err := DoGet(testUrl + "/v1/openstack/manageNode/summary")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Cpu Top Process", func() {
				res, err := DoGet(testUrl + "/v1/openstack/manageNode/controller/topProcessCpu")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("Memory Top Process", func() {
				res, err := DoGet(testUrl + "/v1/openstack/manageNode/controller/topProcessMem")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

			It("RabbitMq Status", func() {
				res, err := DoGet(testUrl + "/v1/openstack/manageNode/rabbitMqSummary")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})
		})

	})

})
