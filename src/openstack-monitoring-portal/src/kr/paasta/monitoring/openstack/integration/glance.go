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

type Glance struct {
	OpenstackProvider		models.OpenstackProvider
}

func NewGlance(openstack_provider models.OpenstackProvider) *Glance {
	return &Glance{
		OpenstackProvider: 	openstack_provider,
	}
}

/**
Description : Get Image information
 */
func (n *Glance) GetImageInfo(image_id, project_id string) (image_info models.ImageInfo, err error) {
	var data interface{}

	provider, err := utils.GetAdminToken(n.OpenstackProvider)
	if err != nil {
		return image_info, err
	}

	//client for Compute API operation
	client, err := openstack.NewImageServiceV2(provider, gophercloud.EndpointOpts{
		Region: n.OpenstackProvider.Region,
	})

	response, err := client.Get(fmt.Sprintf("%s/%s/%s/images/%s", models.NovaUrl, models.NovaVersion, project_id, image_id), nil, nil)
	if err != nil {
		return image_info, err
	}
	rawdata, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return image_info, err
	}
	json.Unmarshal(rawdata, &data)
	msg := data.(map[string]interface{})
	i_info := msg["image"].(map[string]interface{})

	if len(i_info) > 0 {
		resources := i_info["metadata"].(map[string]interface{})
		if len(resources) > 0 {
			image_info.Id = i_info["id"].(string)
			image_info.Name = i_info["name"].(string)
			image_info.OsKind = resources["os_distro"].(string)
			image_info.OsType = resources["os_type"].(string)
			image_info.HypervisorType = resources["hypervisor_type"].(string)

			return image_info, err
		}
	}
	return image_info, err
}