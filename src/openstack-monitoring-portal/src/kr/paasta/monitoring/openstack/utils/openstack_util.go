package utils

import (
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack"
	"kr/paasta/monitoring/openstack/models"
)

func GetComputeClient(provider *gophercloud.ProviderClient, region string) (*gophercloud.ServiceClient, error) {

	client, err := openstack.NewComputeV2(provider, gophercloud.EndpointOpts{
		Region: region,
	})
	if err != nil {
		models.MonitLogger.Error("GetComputeClient::", err)
		return client, err
	}
	return client, nil
}

func GetKeystoneClient(provider *gophercloud.ProviderClient) (*gophercloud.ServiceClient) {

	client := openstack.NewIdentityV3(provider)

	return client
}
