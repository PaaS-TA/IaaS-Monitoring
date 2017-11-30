package integration

import (
	"fmt"
	"io/ioutil"
	"encoding/json"

	"kr/paasta/monitoring/openstack/models"
	"kr/paasta/monitoring/openstack/utils"

	"github.com/rackspace/gophercloud/openstack"
	"github.com/rackspace/gophercloud"
)

type Cinder struct {
	OpenstackProvider	models.OpenstackProvider
	Provider 		*gophercloud.ProviderClient
}

func GetCinder(openstack_provider models.OpenstackProvider, provider *gophercloud.ProviderClient) *Cinder {
	return &Cinder{
		OpenstackProvider: 	openstack_provider,
		Provider: 		provider,
	}
}

/**
Description : Get project Storage Max & Used information
 */
func (n *Cinder) GetProjectStorageResources(project_id, project_name string) (result models.ProjectStorageResources, err error) {
	var data interface{}


	n.OpenstackProvider.TenantName = project_name


	//client for Cinder API operation
	client, err := openstack.NewBlockStorageV2(n.Provider, gophercloud.EndpointOpts{
		Region: n.OpenstackProvider.Region,
	})

	response, err := client.Get(fmt.Sprintf("%s/%s/%s/limits", models.CinderUrl, models.CinderVersion, project_id), nil, nil)
	if err != nil {
		return result, err
	}
	rawdata, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return result, err
	}
	json.Unmarshal(rawdata, &data)
	msg := data.(map[string]interface{})

	sub_msg := msg["limits"].(map[string]interface{})

	if len(sub_msg) < 1 {
		return result, nil
	}

	resources := sub_msg["absolute"].(map[string]interface{})

	if len(resources) > 0 {
		result.VolumeLimitGb  = utils.TypeChecker_int(resources["maxTotalVolumeGigabytes"]).(int)
		result.VolumeGb       = utils.TypeChecker_int(resources["totalGigabytesUsed"]).(int)
		result.VolumesLimit   = utils.TypeChecker_int(resources["maxTotalVolumes"]).(int)
		result.Volumes        = utils.TypeChecker_int(resources["totalVolumesUsed"]).(int)
		result.SnapshotsLimit = utils.TypeChecker_int(resources["maxTotalSnapshots"]).(int)
		result.Snapshots      = utils.TypeChecker_int(resources["totalSnapshotsUsed"]).(int)
		result.BackupsLimit   = utils.TypeChecker_int(resources["maxTotalBackups"]).(int)
		result.Backups        = utils.TypeChecker_int(resources["totalBackupsUsed"]).(int)
	}
	return result, nil
}


