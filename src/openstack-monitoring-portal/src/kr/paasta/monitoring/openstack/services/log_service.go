package services

import (
	"kr/paasta/monitoring/openstack/models"
	"kr/paasta/monitoring/openstack/dao"
	"gopkg.in/olivere/elastic.v3"
)

type LogServiceStruct struct {
	elasticClient *elastic.Client
}

func GetLogService(elasticClient *elastic.Client) *LogServiceStruct{
	return &LogServiceStruct{
		elasticClient: 	elasticClient,
	}
}

func (log LogServiceStruct) GetDefaultRecentLog(request models.LogMessage, paging bool) (models.LogMessage, models.ErrMessage) {

	//최근 로그 조회
	return dao.GetLogDao(log.elasticClient).GetDefaultRecentLog(request, paging)
}

func (log LogServiceStruct) GetSpecificTimeRangeLog(request models.LogMessage, paging bool) (models.LogMessage, models.ErrMessage) {

	//특정 시간대 로그 조회
	return dao.GetLogDao(log.elasticClient).GetSpecificTimeRangeLog(request, paging)
}
