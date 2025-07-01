package clients

import (
	"fmt"

	"github.com/maxabella/appgym/dao"
	"github.com/maxabella/appgym/domain"
)

type Activity interface {
	GetActivities() (domain.Activities, error)
	GetActivityById(int) (dao.Activity, error)
	GetActivityByKeyword(keyword string) (dao.Activities, error)
	CreateActivity(dao.Activity) error
	GetCategories() (dao.ActivityTypes, error)
	CreateActivityHour(dao.ActivityHour) error
	EditActivity(dao.Activity) error
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

func (s *ActivityClient) CreateActivity(ActivityInfoDAO dao.Activity) error {
	err := s.DbClient.db.Create(&ActivityInfoDAO)
	if err.Error != nil {
		panic(err.Error)
	}
	return nil
}

func (s *ActivityClient) GetCategories() (dao.ActivityTypes, error) {
	var CategoriesDAO dao.ActivityTypes
	err := s.DbClient.db.Find(&CategoriesDAO)
	if err.Error != nil {
		return dao.ActivityTypes{}, err.Error
	}
	fmt.Println(CategoriesDAO)
	return CategoriesDAO, nil
}

func (s *ActivityClient) CreateActivityHour(ActivityInfo dao.ActivityHour) error {
	err := s.DbClient.db.Create(&ActivityInfo)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (s *ActivityClient) CreateCategory(CategoryInfo dao.ActivityType) error {
	err := s.DbClient.db.Create(&CategoryInfo)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (s *ActivityClient) EditActivity(ActivityInfo dao.Activity) error {
	err := s.DbClient.db.Save(&ActivityInfo)
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (s *ActivityClient) EditActivityHour(HourInfo dao.ActivityHour) error {
	fmt.Println(HourInfo)
	err := s.DbClient.db.Model(dao.ActivityHour{}).Where("id = ?", HourInfo.ID).Updates(map[string]interface{}{
		"day":           HourInfo.Day,
		"starting_hour": HourInfo.Starting_Hour,
		"finish_hour":   HourInfo.Finish_hour,
	})
	if err.Error != nil {
		return err.Error
	}
	return nil
}

func (s *ActivityClient) UpdateInscriptions(InscriptionInfo dao.Inscription) error {
	fmt.Println(InscriptionInfo)
	err := s.DbClient.db.Model(dao.Inscription{}).Where("activity_id = ?", InscriptionInfo.ActivityID).Updates(map[string]interface{}{
		"day":           InscriptionInfo.Day,
		"starting_hour": InscriptionInfo.Starting_Hour,
		"finish_hour":   InscriptionInfo.Finish_hour,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *ActivityClient) DeleteActivity(id uint) error {
	// DELETE INSCRIPTIONS 1

	err := s.DbClient.db.Where("activity_id = ?", id).Delete(dao.Inscription{})
	if err.Error != nil {
		return err.Error
	}
	// DELETE HOURS 2

	err = s.DbClient.db.Where("activity_id = ?", id).Delete(dao.ActivityHour{})
	if err.Error != nil {
		return err.Error
	}

	// DELETE Activity 3
	err = s.DbClient.db.Where("id = ?", id).Delete(dao.Activities{})
	if err.Error != nil {
		return err.Error
	}

	return nil
}
