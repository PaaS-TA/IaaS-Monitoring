package controller_test

import (
	. "github.com/onsi/ginkgo"
	"github.com/stretchr/testify/assert"
	"net/http"
)

/*type Config map[string]string
var TestToken string
type Response struct {
	Token   string
	Content string
	Code    int
}

type PingResponse struct {
	Token   string
	Code    int
}*/


var _ = Describe("MainController", func() {

	Describe("Main & Openstack Summary", func() {
		Context("Openstack ", func() {

			It("Main", func() {
				DoGet(testUrl + "/")
			})

			It("Openstack Summary", func() {
				res, err := DoGet(testUrl + "/v1/openstack/summary")
				assert.Nil(t, err)
				assert.Equal(t, http.StatusOK, res.Code)
			})

		})
	})



})
