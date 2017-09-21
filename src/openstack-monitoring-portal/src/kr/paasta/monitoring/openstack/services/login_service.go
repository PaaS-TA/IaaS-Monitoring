package services

import (
	"kr/paasta/monitoring/openstack/models"
	"kr/paasta/monitoring/openstack/utils"
	"github.com/rackspace/gophercloud"
	"github.com/rackspace/gophercloud/openstack/identity/v3/tokens"
	"kr/paasta/monitoring/openstack/integration"
)


type LoginService struct {
	openstackProvider models.OpenstackProvider
}

func GetLoginService(openstackProvider models.OpenstackProvider) *LoginService {
	return &LoginService{
		openstackProvider: openstackProvider,
	}
}


func (n LoginService) LogOut(provider *gophercloud.ProviderClient)(tokens.RevokeResult) {

	result := integration.GetKeystone(n.openstackProvider, provider).RevokeToken()
	return result
}

func (n LoginService) Login(req models.User)(userInfo models.User, provider  *gophercloud.ProviderClient, err error) {

	n.openstackProvider.Username = req.Username
	n.openstackProvider.Password = req.Password
	provider, err = utils.GetAdminToken(n.openstackProvider)

	if err != nil {
		return req,  provider, err
	}

	//var userInfo models.User
	userInfo.Username = req.Username
	userInfo.Password = req.Password
	userInfo.Token = provider.TokenID

	return userInfo, provider, err
}