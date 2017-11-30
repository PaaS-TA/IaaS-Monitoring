package openstack_monitoring_portal_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

func TestOpenstackMonitoringPortal(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "OpenstackMonitoringPortal Suite")
}
