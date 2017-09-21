package services

import (
	"github.com/monasca/golang-monascaclient/monascaclient"
	client "github.com/influxdata/influxdb/client/v2"
	"kr/paasta/monitoring/openstack/integration"
	mod "github.com/monasca/golang-monascaclient/monascaclient/models"
	"kr/paasta/monitoring/openstack/models"
	"github.com/jinzhu/gorm"
	"kr/paasta/monitoring/openstack/dao"
	"kr/paasta/monitoring/openstack/utils"
	"reflect"
	"time"
)

type AlarmStatusService struct {
	monClient     monascaclient.Client
	influxClient 	client.Client
	txn           *gorm.DB
}

func GetAlarmStatusService( monClient    monascaclient.Client, influxClient client.Client, txn *gorm.DB) *AlarmStatusService {
	return &AlarmStatusService{
		monClient: monClient,
		influxClient: 	influxClient,
		txn:   txn,
	}
}

func (a *AlarmStatusService)GetAlarmStatusCount(query mod.AlarmQuery) (map[string]interface{}, error){

	totalData, errCnt := integration.GetMonasca(a.monClient).GetAlarmCount(query)

	if errCnt != nil{
		return nil, errCnt
	}
	resultData := map[string]interface{}{
		models.RESULT_CNT: totalData[0][0],
	}
	return resultData, nil
}

func (a *AlarmStatusService)GetAlarmStatusList(query mod.AlarmQuery) (map[string]interface{}, error){

	var allQuery mod.AlarmQuery

	allQuery.Severity = query.Severity
	allQuery.State = query.State


	totalData, errCnt := integration.GetMonasca(a.monClient).GetAlarmList(allQuery)

	if errCnt != nil{
		return nil, errCnt
	}


	alarmStatusList, err := integration.GetMonasca(a.monClient).GetAlarmList(query)

	if err != nil{
		return nil, err
	}

	var result []models.AlarmStatus

	for _, data := range alarmStatusList{

		alarmDefinition, err := dao.GetMonascaDbDao(a.txn).GetAlarmsDefinition(data.Id)
		if err != nil{
			return nil, err
		}
		alarmStatus := data
		alarmStatus.AlarmDefinitionId   = alarmDefinition.AlarmDefinitionId
		alarmStatus.AlarmDefinitionName = alarmDefinition.Name
		alarmStatus.Expression 		= alarmDefinition.Expression
		alarmStatus.Severity		= alarmDefinition.Severity

		result = append(result, alarmStatus)
	}
	resultData := map[string]interface{}{
		models.RESULT_CNT: len(totalData),
		models.RESULT_DATA: result,
	}

	return resultData, err
}

func (a *AlarmStatusService)GetAlarmHistoryList(alarmReq models.AlarmReq) (result []models.AlarmHistory, err models.ErrMessage){

	alarmHistoryResp, err := dao.GetAlarmDao(a.influxClient, a.txn).GetAlarmHistoryList(alarmReq)
	if err != nil{
		models.MonitLogger.Error("Error==>", err)
		return result, err
	}

	alarmHistoryList, err := utils.GetResponseConverter().InfluxConverterToMap(alarmHistoryResp)

	if err != nil{
		return result, err
	}

	for _, data := range alarmHistoryList{
		var alarmHistory models.AlarmHistory

		alarmHistory.Id       = alarmReq.AlarmId
		occurDate := time.Unix(reflect.ValueOf(data["time"]).Int(), 0 )
		alarmHistory.Time     = occurDate.Format("2006-01-02 15:04:05")
		alarmHistory.NewState = reflect.ValueOf(data["new_state"]).String()
		alarmHistory.OldState = reflect.ValueOf(data["old_state"]).String()
		alarmHistory.Reason   = reflect.ValueOf(data["reason"]).String()
		result = append(result, alarmHistory)
	}

	return result, err

}

func (a *AlarmStatusService)GetAlarmStatus(alarmId string) (result models.AlarmStatus, err error){

	result, err = integration.GetMonasca(a.monClient).GetAlarm(alarmId)

	if err != nil{
		return result, err
	}

	alarmDefinition, err := dao.GetMonascaDbDao(a.txn).GetAlarmsDefinition(result.Id)
	if err != nil{
		return result, err
	}


	result.AlarmDefinitionId   = alarmDefinition.AlarmDefinitionId
	result.AlarmDefinitionName = alarmDefinition.Name
	result.Expression 	   = alarmDefinition.Expression
	result.Severity		   = alarmDefinition.Severity

	return result, nil
}

func (a *AlarmStatusService)GetAlarmHistoryActionList(alarmId string) (result []models.AlarmActionResponse, err error){

	var alarmRequest models.AlarmActionRequest
	alarmRequest.AlarmId = alarmId

	alarmActionList , err  := dao.GetAlarmDao(a.influxClient, a.txn).GetAlarmsActionHistoryList(alarmRequest)


	for _, data := range alarmActionList{
		var alarmAction models.AlarmActionResponse
		alarmAction.Id = data.Id
		alarmAction.AlarmId = data.AlarmId
		alarmAction.AlarmActionDesc = data.AlarmActionDesc
		alarmAction.RegDate  = data.RegDate.Add(time.Duration(models.GmtTimeGap) * time.Hour).Format("2006-01-02 15:04:05")
		alarmAction.RegUser  = data.RegUser
		alarmAction.RegDate  = data.RegDate.Add(time.Duration(models.GmtTimeGap) * time.Hour).Format("2006-01-02 15:04:05")
		alarmAction.ModiUser  = data.ModiUser

		result = append(result, alarmAction)
	}
	if err != nil {
		return result , err
	}

	return result, nil
}




func (h *AlarmStatusService) CreateAlarmHistoryAction(request models.AlarmActionRequest) error {

	dbErr := dao.GetAlarmDao(h.influxClient, h.txn).CreateAlarmAction(request)
	return dbErr
}

func (h *AlarmStatusService) UpdateAlarmAction(request models.AlarmActionRequest) error {

	dbErr := dao.GetAlarmDao(h.influxClient, h.txn).UpdateAlarmAction(request)
	return dbErr
}

func (h *AlarmStatusService) DeleteAlarmAction(request models.AlarmActionRequest) error {

	dbErr := dao.GetAlarmDao(h.influxClient, h.txn).DeleteAlarmAction(request)
	return dbErr
}
