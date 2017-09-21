package integration

import (
	"github.com/rackspace/gophercloud/openstack"
	"kr/paasta/monitoring/openstack/utils"
	"kr/paasta/monitoring/openstack/models"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/identity/v3/tokens"
	"fmt"
)

type Keystone struct {
	OpenstackProvider 	models.OpenstackProvider
	provider        	*gophercloud.ProviderClient
}

func GetKeystone(openstack_provider models.OpenstackProvider, provider *gophercloud.ProviderClient) *Keystone {
	return &Keystone{
		OpenstackProvider: openstack_provider,
		provider: provider,
	}
}

/**
Description : Get openstack's registered project lists
 */
func (k *Keystone) GetProjectList() (project_lists []models.ProjectInfo, err error){

	//client for Keystone API operation
	client := openstack.NewIdentityV3(k.provider)
	response, err := client.Get(fmt.Sprintf("%s/%s/projects", models.KeystoneUrl, models.KeystoneVersion), nil, nil)

	msg , err := utils.ResponseUnmarshal(response, err)

	for _, v :=range msg["projects"].([]interface{}){
		var project_info models.ProjectInfo
		v_detail := v.(map[string]interface{})
		project_info.Id 	= utils.TypeChecker_string(v_detail["id"]).(string)
		project_info.Name	= utils.TypeChecker_string(v_detail["name"]).(string)
		project_info.IsDomain	= v_detail["is_domain"].(bool)
		project_info.Description	= utils.TypeChecker_string(v_detail["description"]).(string)
		project_info.Enabled	= v_detail["enabled"].(bool)
		project_info.ParentId	= utils.TypeChecker_string(v_detail["parent_id"]).(string)
		project_info.Links	= v_detail["links"].(map[string]interface{})

		project_lists = append(project_lists, project_info)
	}
	return project_lists, nil
}

func (k *Keystone) GetUserIdByName(userName string)(userId string, err error) {


	/*provider, err := utils.GetAdminToken(k.OpenstackProvider)
	if err != nil {
		return userId, err
	}*/
	//client for Keystone API operation
	client := openstack.NewIdentityV3(k.provider)

	response, err := client.Get(fmt.Sprintf("%s/%s/users?name=%s", models.KeystoneUrl, models.KeystoneVersion, userName ), nil, nil)
	msg , err := utils.ResponseUnmarshal(response, err)

	for _, v :=range msg["users"].([]interface{}){
		v_detail := v.(map[string]interface{})
		userId = utils.TypeChecker_string(v_detail["id"]).(string)
	}
	return userId, err
}

func (k *Keystone) GetUserProjectList(userId string) (project_lists []models.ProjectInfo, err error){

	client := openstack.NewIdentityV3(k.provider)

	response, err := client.Get(fmt.Sprintf("%s/%s/users/%s/projects", models.KeystoneUrl, models.KeystoneVersion, userId ), nil, nil)
	msg , err := utils.ResponseUnmarshal(response, err)

	for _, v :=range msg["projects"].([]interface{}){
		var project_info models.ProjectInfo
		v_detail := v.(map[string]interface{})
		project_info.Id 	= utils.TypeChecker_string(v_detail["id"]).(string)
		project_info.Name	= utils.TypeChecker_string(v_detail["name"]).(string)
		project_info.IsDomain	= v_detail["is_domain"].(bool)
		project_info.Description	= utils.TypeChecker_string(v_detail["description"]).(string)
		project_info.Enabled	= v_detail["enabled"].(bool)
		project_info.ParentId	= utils.TypeChecker_string(v_detail["parent_id"]).(string)
		project_info.Links	= v_detail["links"].(map[string]interface{})

		project_lists = append(project_lists, project_info)
	}
	return project_lists, nil
}

func (k *Keystone) RevokeToken()(tokens.RevokeResult){

	client := openstack.NewIdentityV3(k.provider)
	result := tokens.Revoke(client, client.TokenID)

	return result
}