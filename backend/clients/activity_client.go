package clients

import (
	"github.com/maxabella/appgym/dao"
	"github.com/maxabella/appgym/domain"
)

type Activity interface {
	GetActivities() (domain.Activities, error)
	GetActivityById(int) (dao.Activity, error)
}

type ActivityClient struct {
	DbClient *Mysql_Client
}

func (s *ActivityClient) GetActivities() (dao.Activities, error) {
	var activitiesDao dao.Activities
	result := s.DbClient.db.Preload("Coach").Preload("ActivityType").Preload("ActivityHours").Find(&activitiesDao)
	if result.Error != nil {
		return dao.Activities{}, result.Error
	}
	return activitiesDao, nil
}

func (s *ActivityClient) GetActivityByKeyword(keyword string) (dao.Activities, error) {
	var activitiesDao dao.Activities
	searchKeyword := "%" + keyword + "%"
	result := s.DbClient.db.Preload("Coach").Preload("ActivityType").Preload("ActivityHours").Joins("JOIN activity_types ON activities.activity_type_id = activity_types.id").Where("activities.name LIKE ? OR activities.description LIKE ? OR activity_types.name LIKE ?", searchKeyword, searchKeyword, searchKeyword).Find(&activitiesDao)
	if result.Error != nil {
		return dao.Activities{}, result.Error
	}
	return activitiesDao, nil
}

func (s *ActivityClient) GetActivityById(id int) (dao.Activity, error) {
	var activityDao dao.Activity
	result := s.DbClient.db.Preload("Coach").Preload("ActivityType").Preload("ActivityHours").Where("id = ?", id).First(&activityDao)
	if result.Error != nil {
		return dao.Activity{}, result.Error
	}
	return activityDao, nil
}
