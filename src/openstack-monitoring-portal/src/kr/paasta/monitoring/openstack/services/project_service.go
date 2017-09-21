package services

import (client "github.com/influxdata/influxdb/client/v2"
	"kr/paasta/monitoring/openstack/models"
	"kr/paasta/monitoring/openstack/integration"
	"kr/paasta/monitoring/openstack/dao"
	"kr/paasta/monitoring/openstack/utils"
	"github.com/rackspace/gophercloud"
	"fmt"
	"sync"
	"strings"
	"encoding/json"
	"strconv"
)

type ProjectService struct {
	openstackProvider models.OpenstackProvider
	provider          *gophercloud.ProviderClient
	influxClient 	client.Client
}

func GetProjectService(openstackProvider models.OpenstackProvider, provider *gophercloud.ProviderClient,influxClient client.Client) *ProjectService {
	return &ProjectService{
		openstackProvider: openstackProvider,
		provider: provider,
		influxClient: 	influxClient,
	}
}



//CPU 사용률
func (n ProjectService) GetInstanceCpuUsageList(request models.DetailReq)(result []map[string]interface{}, _ models.ErrMessage){

	cpuUsageResp, err := dao.GetInstanceDao(n.influxClient).GetInstanceCpuUsageList(request)
	if err != nil {
		models.MonitLogger.Error(err)
		return result, err
	}else {
		cpuUsage, _ := utils.GetResponseConverter().InfluxConverterList(cpuUsageResp, models.METRIC_NAME_CPU_USAGE)

		datamap := cpuUsage[models.RESULT_DATA_NAME].([]map[string]interface{})
		for _, data := range datamap{

			swapFree := data["usage"].(json.Number)
			convertData, _ := strconv.ParseFloat(swapFree.String(),64)
			//swap 사용률을 구한다. ( 100 -  freeUsage)
			data["usage"] = utils.RoundFloatDigit2(convertData)
		}

		result = append(result,cpuUsage )
		return result, nil
	}
}


//Instance Memory Usage
func (s ProjectService) GetInstanceMemoryUsageList(request models.DetailReq)(result []map[string]interface{}, _ models.ErrMessage){

	memoryResp, err := dao.GetInstanceDao(s.influxClient).GetInstanceMemoryUsageList(request)
	if err != nil {
		models.MonitLogger.Error(err)
		return result, err
	}else {
		memoryUsage, _ := utils.GetResponseConverter().InfluxConverter4Usage(memoryResp, models.METRIC_NAME_MEMORY_USAGE)

		datamap := memoryUsage[models.RESULT_DATA_NAME].([]map[string]interface{})
		for _, data := range datamap{

			swapFree := data["usage"].(float64)
			data["usage"] = utils.RoundFloatDigit2(swapFree)
		}

		result = append(result, memoryUsage )
		return result, nil
	}
}


//Disk IO Read Kbyte
func (s ProjectService) GetInstanceDiskIoKbyteList(request models.DetailReq, gubun string)(result []map[string]interface{}, _ models.ErrMessage){

	memoryResp, err := dao.GetInstanceDao(s.influxClient).GetInstanceDiskIoKbyte(request, gubun)
	if err != nil {
		models.MonitLogger.Error(err)
		return result, err
	}else {
		var resultName string

		if gubun == "read"{
			resultName = models.METRIC_NAME_DISK_READ_KBYTE
		}else{
			resultName = models.METRIC_NAME_DISK_WRITE_KBYTE
		}
		byte, _ := utils.GetResponseConverter().InfluxConverterList(memoryResp, resultName)

		datamap := byte[models.RESULT_DATA_NAME].([]map[string]interface{})
		fmt.Println(datamap)
		for _, data := range datamap{
			usage := utils.TypeChecker_float64(data["usage"]).(float64)
			data["usage"] = utils.RoundFloatDigit2(usage)
		}

		result = append(result, byte )
		return result, nil
	}
}

//Network IO Kbyte
func (s ProjectService) GetInstanceNetworkIoKbyteList(request models.DetailReq)(result []map[string]interface{}, _ models.ErrMessage){

	networkInResp, err := dao.GetInstanceDao(s.influxClient).GetInstanceNetworkKbyte(request, "in")
	networkOutResp, err := dao.GetInstanceDao(s.influxClient).GetInstanceNetworkKbyte(request, "out")

	if err != nil {
		models.MonitLogger.Error(err)
		return result, err
	}else {


		inData, _ := utils.GetResponseConverter().InfluxConverterList(networkInResp, models.METRIC_NAME_NETWORK_IN)
		outData, _ := utils.GetResponseConverter().InfluxConverterList(networkOutResp, models.METRIC_NAME_NETWORK_OUT)

		inDatamap := inData[models.RESULT_DATA_NAME].([]map[string]interface{})
		for _, data := range inDatamap{
			usage := utils.TypeChecker_float64(data["usage"]).(float64)
			data["usage"] = utils.RoundFloatDigit2(usage)
		}

		outDatamap := outData[models.RESULT_DATA_NAME].([]map[string]interface{})
		for _, data := range outDatamap{
			usage := utils.TypeChecker_float64(data["usage"]).(float64)
			data["usage"] = utils.RoundFloatDigit2(usage)
		}

		result = append(result, inData )
		result = append(result, outData )
		return result, nil
	}
}

//Network Packets
func (s ProjectService) GetInstanceNetworkPacketsList(request models.DetailReq)(result []map[string]interface{}, _ models.ErrMessage){

	networkInResp,  err := dao.GetInstanceDao(s.influxClient).GetInstanceNetworkPackets(request, "in")
	networkOutResp, err := dao.GetInstanceDao(s.influxClient).GetInstanceNetworkPackets(request, "out")

	if err != nil {
		models.MonitLogger.Error(err)
		return result, err
	}else {


		inData, _ := utils.GetResponseConverter().InfluxConverterList(networkInResp, models.METRIC_NAME_NETWORK_IN)
		outData, _ := utils.GetResponseConverter().InfluxConverterList(networkOutResp, models.METRIC_NAME_NETWORK_OUT)

		inDatamap := inData[models.RESULT_DATA_NAME].([]map[string]interface{})

		for _, data := range inDatamap{
			usage := utils.TypeChecker_float64(data["usage"]).(float64)
			data["usage"] = utils.RoundFloatDigit2(usage)
		}

		outDatamap := outData[models.RESULT_DATA_NAME].([]map[string]interface{})

		for _, data := range outDatamap{
			usage := utils.TypeChecker_float64(data["usage"]).(float64)
			data["usage"] = utils.RoundFloatDigit2(usage)
		}

		result = append(result, inData )
		result = append(result, outData )
		return result, nil
	}
}

func (n ProjectService) GetProjectSummary(apiRequest models.ProjectReq)(result []models.ProjectSummaryInfo, err error){

	userId, err := integration.GetKeystone(n.openstackProvider, n.provider).GetUserIdByName(n.openstackProvider.Username)
	if err != nil {
		fmt.Println("Get UserId Error :", err)
		return result, err
	}

	//Get Project List by User Own
	projectLists, err := integration.GetKeystone(n.openstackProvider,  n.provider).GetUserProjectList(userId)

	var searchProjectList []models.ProjectInfo

	for _, project := range projectLists{
		if apiRequest.ProjectName != "" && strings.Contains(project.Name, apiRequest.ProjectName){
			searchProjectList = append(searchProjectList, project)
		}else if apiRequest.ProjectName == ""{
			searchProjectList = append(searchProjectList, project)
		}


	}

	if err != nil {
		fmt.Println("Get nodes resources error :", err)
	}

	var projectSummaryInfos []models.ProjectSummaryInfo

	for _, project := range searchProjectList{

		var projectSummaryInfo models.ProjectSummaryInfo

		projectInstances,  projectResourceLimit, projectNetworkLimit, projectFloatingIps, projectSecurityGroups, projectStorageResources,_ := getProjectSummary_Sub(n.openstackProvider, n.provider,  project.Id, project.Name)

		var total_vcpus, total_memory float64
		for _, instance :=range projectInstances{

			total_vcpus = total_vcpus + instance.Vcpus
			total_memory = total_memory + instance.MemoryMb
			//total_disk = total_disk + instance.Disk_gb
		}

		projectSummaryInfo.Name = project.Name
		projectSummaryInfo.Id = project.Id
		projectSummaryInfo.Enabled = project.Enabled
		projectSummaryInfo.InstancesUsed = len(projectInstances)
		projectSummaryInfo.MemoryMbUsed = total_memory
		projectSummaryInfo.VcpusUsed = total_vcpus

		projectSummaryInfo.MemoryMbLimit = projectResourceLimit.MemoryMbLimit
		projectSummaryInfo.InstancesLimit = projectResourceLimit.InstancesLimit
		projectSummaryInfo.VcpusLimit  = projectResourceLimit.CoresLimit

		projectSummaryInfo.SecurityGroupsLimit = projectNetworkLimit.SecurityGroupLimit
		projectSummaryInfo.FloatingIpsLimit    = projectNetworkLimit.FloatingIpsLimit

		projectSummaryInfo.FloatingIpsUsed     = len(projectFloatingIps)
		projectSummaryInfo.SecurityGroupsUsed  = projectSecurityGroups

		projectSummaryInfo.VolumeStorageLimit   = projectStorageResources.VolumesLimit
		projectSummaryInfo.VolumeStorageUsed    = projectStorageResources.Volumes
		projectSummaryInfo.VolumeStorageLimitGb = projectStorageResources.VolumeLimitGb
		projectSummaryInfo.VolumeStorageUsedGb    = projectStorageResources.VolumeGb

		projectSummaryInfos = append(projectSummaryInfos, projectSummaryInfo)
	}

	return projectSummaryInfos, nil

}

func  getProjectSummary_Sub(opts models.OpenstackProvider, provider *gophercloud.ProviderClient, projectId ,projectName string)(projectInstances []models.InstanceInfo,
			projectResourcesLimit models.ProjectResourcesLimit, projectNetworkLimit models.ProjectNetworkLimit,
			projectFloatingIps []models.FloatingIPInfo, projectSecurityGroups int,  projectStorageResources models.ProjectStorageResources,
			_ models.ErrMessage) {

	var errs []models.ErrMessage
	var err error
	var wg sync.WaitGroup
	wg.Add(6)
	for i := 0; i < 6; i++ {
		go func(wg *sync.WaitGroup, index int) {
			switch index {
			case 0 :
				projectInstances, err  = integration.GetNova(opts, provider).GetProjectInstances(projectId)
				if err != nil {
					//errs = append(errs, err)
				}
			case 1 :
				projectResourcesLimit, err = integration.GetNova(opts, provider).GetProjectResourcesLimit(projectId)
				if err != nil {
					//errs = append(errs, err)
				}
			case 2 :
				projectNetworkLimit, err = integration.GetNeutron(opts, provider).GetProjectNetworkLimit(projectId)
				if err != nil {
					//errs = append(errs, err)
				}
			case 3 :
				projectFloatingIps, err = integration.GetNeutron(opts, provider).GetProjectFloatingIps(projectId)
				if err != nil {
					//errs = append(errs, err)
				}
			case 4 :
				projectSecurityGroups, err = integration.GetNeutron(opts, provider).GetProjectSecurityGroups(projectId)
				if err != nil {
					//errs = append(errs, err)
				}
			case 5 :
				projectStorageResources, err = integration.GetCinder(opts, provider).GetProjectStorageResources(projectId, projectName)
				if err != nil {
					//errs = append(errs, err)
				}
			}

			wg.Done()
		}(&wg, i)
	}
	wg.Wait()

	//==========================================================================
	// Error가 여러건일 경우 대해 고려해야함.
	if len(errs) > 0 {
		var returnErrMessage string
		for _, err := range errs{
			returnErrMessage = returnErrMessage + " " + err["Message"].(string)
		}
		errMessage := models.ErrMessage{
			"Message": returnErrMessage ,
		}
		return projectInstances, projectResourcesLimit, projectNetworkLimit, projectFloatingIps, projectSecurityGroups, projectStorageResources, errMessage
	}
	//==========================================================================
	models.MonitLogger.Debug("projectStorageResources::", projectStorageResources)
	return projectInstances, projectResourcesLimit, projectNetworkLimit, projectFloatingIps, projectSecurityGroups, projectStorageResources, nil
}

func (n ProjectService) GetProjectInstanceList(apiRequest models.ProjectReq)(resultArr map[string]interface{}, err error){

	var result []models.InstanceInfo
	instanceMainList , _ := integration.GetNova(n.openstackProvider, n.provider).GetProjectInstancesList(apiRequest)
	instanceSubInfo , _ :=  integration.GetNova(n.openstackProvider, n.provider).GetProjectInstances(apiRequest.ProjectId)

	totInstance := 0
	for _, instance := range instanceSubInfo{
		if apiRequest.HostName != "" && strings.Contains(instance.Name, apiRequest.HostName){
			totInstance += 1
		}else if apiRequest.HostName == ""{
			totInstance += 1
		}
	}

	var errs []models.ErrMessage

	for _, mainInstance := range instanceMainList{
		for _, subInstance := range instanceSubInfo{
			if mainInstance.InstanceId == subInstance.InstanceId{
				mainInstance.Flavor    = subInstance.Flavor
				mainInstance.TenantId  = subInstance.TenantId
				mainInstance.Vcpus     = subInstance.Vcpus
				mainInstance.MemoryMb  = subInstance.MemoryMb
				mainInstance.DiskGb    = subInstance.DiskGb
				mainInstance.StartedAt = subInstance.StartedAt
				mainInstance.Uptime    = subInstance.Uptime
			}
		}
		var req models.InstanceReq
		req.InstanceId = mainInstance.InstanceId
		cpuData, memTotData, memFreeData, err := getProjectInstanceStatus_Sub(req, n.influxClient)

		if err != nil {
			errs = append(errs, err)
		}

		cpuUsage  := utils.GetDataFloatFromInterfaceSingle(cpuData)
		memTot    := utils.GetDataFloatFromInterfaceSingle(memTotData)
		memFree   := utils.GetDataFloatFromInterfaceSingle(memFreeData)

		mainInstance.CpuUsage = cpuUsage
		if int(memTot) > 0 && int(memFree) > 0{
			mainInstance.MemoryUsage = utils.RoundFloatDigit2(100 - ((memFree/memTot)*100))
		}

		result = append(result, mainInstance)
	}

	resultArr = map[string]interface{}{
		models.RESULT_CNT: totInstance,
		models.RESULT_PROJECT_ID: apiRequest.ProjectId,
		models.RESULT_DATA_NAME: result,
	}

	return resultArr, nil

}

func  getProjectInstanceStatus_Sub( request models.InstanceReq, f client.Client )(map[string]interface{}, map[string]interface{}, map[string]interface{}, models.ErrMessage) {

	var errs []models.ErrMessage
	var err models.ErrMessage
	var wg sync.WaitGroup
	var cpuResp, memTotalResp, memFreeResp client.Response

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go func(wg *sync.WaitGroup, index int) {
			switch index {
			case 0 :
				cpuResp, err  = dao.GetInstanceDao(f).GetInstanceCpuUsage(request)
				if err != nil {
					errs = append(errs, err)
				}
			case 1 :
				memTotalResp, err = dao.GetInstanceDao(f).GetInstanceTotalMemoryUsage(request)
				if err != nil {
					errs = append(errs, err)
				}
			case 2 :
				memFreeResp, err = dao.GetInstanceDao(f).GetInstanceFreeMemoryUsage(request)
				if err != nil {
					errs = append(errs, err)
				}
			}

			wg.Done()
		}(&wg, i)
	}
	wg.Wait()

	//==========================================================================
	// Error가 여러건일 경우 대해 고려해야함.
	if len(errs) > 0 {
		var returnErrMessage string
		for _, err := range errs{
			returnErrMessage = returnErrMessage + " " + err["Message"].(string)
		}
		errMessage := models.ErrMessage{
			"Message": returnErrMessage ,
		}
		return nil, nil, nil  , errMessage
	}
	//==========================================================================

	cpuUsage, _   := utils.GetResponseConverter().InfluxConverter(cpuResp)
	memTotal, _   := utils.GetResponseConverter().InfluxConverter(memTotalResp)
	memFree,  _   := utils.GetResponseConverter().InfluxConverter(memFreeResp)


	return cpuUsage, memTotal , memFree, nil
}