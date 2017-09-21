package dao

import (
	"github.com/jinzhu/gorm"
	"kr/paasta/monitoring/openstack/models"
)

type MonascaDao struct {
	txn   *gorm.DB
}

func GetMonascaDbDao(txn *gorm.DB) *MonascaDao {
	return &MonascaDao{
		txn:   txn,
	}
}

func (m *MonascaDao) GetAlarmsDefinition(alarmId string) (models.Alarm, error) {

	var alarm models.Alarm

	status := m.txn.Debug().Table("alarm").
		Select("alarm.id, alarm.alarm_definition_id, alarm_definition.name, alarm_definition.expression, alarm_definition.severity").
		Joins("inner join alarm_definition on alarm_definition.id = alarm.alarm_definition_id ").
		Where("alarm.id = ?", alarmId).
		Find(&alarm)
	if status.Error != nil{
		return alarm, status.Error
	}
	return alarm, status.Error
}

