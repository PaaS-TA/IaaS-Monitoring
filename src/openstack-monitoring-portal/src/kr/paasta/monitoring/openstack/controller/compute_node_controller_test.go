package controller_test

import (
	"github.com/stretchr/testify/assert"
	. "github.com/onsi/ginkgo"
	"net/http"
)

var _ = Describe("ComputeNodeController", func() {

		Describe("GetNode List", func() {

			It("Compute Node List", func() {
				res, err := DoGet(testUrl + "/v1/openstack/node")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

		})

		Describe("ComputeNode", func() {

			It("Compute Node Summary", func() {
				res, err := DoGet(testUrl + "/v1/openstack/computeNode/summary")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})


		})

		Describe("Node Usage", func() {
			Context("Cpu", func() {
				It("Node Cpu Usage", func() {
					res, err := DoGet(testUrl + "/v1/openstack/node/controller/cpuUsage?defaultTimeRange=10m&groupBy=1m")
					assert.Nil(t, err)
					assert.Equal(t, http.StatusOK, res.Code)
				})

				It("Node Cpu Load", func() {
					res, err := DoGet(testUrl + "/v1/openstack/node/controller/cpuLoad?defaultTimeRange=10m&groupBy=1m")
					assert.Nil(t, err)
					assert.Equal(t, http.StatusOK, res.Code)
				})
			})

			Context("Memory", func() {
				It("Node Memory Usage", func() {
					res, err := DoGet(testUrl + "/v1/openstack/node/controller/memUsage?defaultTimeRange=10m&groupBy=1m")
					assert.Nil(t, err)
					assert.Equal(t, http.StatusOK, res.Code)
				})

				It("Node Swap", func() {
					res, err := DoGet(testUrl + "/v1/openstack/node/controller/swapUsage?defaultTimeRange=10m&groupBy=1m")
					assert.Nil(t, err)
					assert.Equal(t, http.StatusOK, res.Code)
				})
			})

			Context("Disk", func() {
				It("Node Disk Usage", func() {
					res, err := DoGet(testUrl + "/v1/openstack/node/controller/diskUsage?defaultTimeRange=10m&groupBy=1m")
					assert.Nil(t, err)
					assert.Equal(t, http.StatusOK, res.Code)
				})

				It("Node Disk Read", func() {
					res, err := DoGet(testUrl + "/v1/openstack/node/controller/diskRead?defaultTimeRange=10m&groupBy=1m")
					assert.Nil(t, err)
					assert.Equal(t, http.StatusOK, res.Code)
				})


				It("Node Disk Write", func() {
					res, err := DoGet(testUrl + "/v1/openstack/node/controller/diskWrite?defaultTimeRange=10m&groupBy=1m")
					assert.Nil(t, err)
					assert.Equal(t, http.StatusOK, res.Code)
				})
			})

			Context("Network", func() {
				It("Node Network IO", func() {
					res, err := DoGet(testUrl + "/v1/openstack/node/controller/networkIo?defaultTimeRange=10m&groupBy=1m")
					assert.Nil(t, err)
					assert.Equal(t, http.StatusOK, res.Code)
				})
			})

			Context("Network", func() {
				It("Node Network Err", func() {
					res, err := DoGet(testUrl + "/v1/openstack/node/controller/networkError?defaultTimeRange=10m&groupBy=1m")
					assert.Nil(t, err)
					assert.Equal(t, http.StatusOK, res.Code)
				})
			})

			Context("Network", func() {
				It("Node Network DropPacket", func() {
					res, err := DoGet(testUrl + "/v1/openstack/node/controller/networkDropPacket?defaultTimeRange=10m&groupBy=1m")
					assert.Nil(t, err)
					assert.Equal(t, http.StatusOK, res.Code)
				})
			})

		})


})

