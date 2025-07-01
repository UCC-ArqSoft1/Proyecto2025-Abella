package services

import (
	"fmt"

	"github.com/maxabella/appgym/clients"
	"github.com/maxabella/appgym/dao"
	"github.com/maxabella/appgym/domain"
)

/*
 */

type Activity interface {
	GetActivities() (domain.Activities, error)
	GetActivityByKeyword(keyword string) (domain.Activities, error)
	GetActivityById(int) (domain.Activity, error)
	CreateActivity(domain.NewActivity) error
	GetCategories() (domain.ActivityTypes, error)
	CreateActivityHour(domain.NewHour) error
}

type ActivityService struct {
	ActivityClient *clients.ActivityClient
}

func (s *ActivityService) GetCategories() (domain.ActivityTypes, error) {
	var CategoriesDTO domain.ActivityTypes
	CategoriesDAO, err := s.ActivityClient.GetCategories()
	if err != nil {
		return domain.ActivityTypes{}, err
	}
	fmt.Println(CategoriesDTO)
	for _, Category := range CategoriesDAO {
		CategoriesDTO = append(CategoriesDTO, domain.ActivityType{
			Id:   Category.ID,
			Name: Category.Name,
		})
	}
	return CategoriesDTO, nil
}

func (s *ActivityService) GetActivities() (domain.Activities, error) {

	activitiesDao, err := s.ActivityClient.GetActivities()
	if err != nil {
		return domain.Activities{}, err
	}

	var Activities domain.Activities
	for _, activity := range activitiesDao {
		var activitieshours []domain.ActivityHours
		for _, hour := range activity.ActivityHours {
			activitieshours = append(activitieshours, domain.ActivityHours{
				Id:          hour.ID,
				Day:         hour.Day,
				Hour_start:  hour.Starting_Hour,
				Hour_finish: hour.Finish_hour,
			})
		}
		Activities = append(Activities, domain.Activity{
			ID:            activity.ID,
			Name:          activity.Name,
			Description:   activity.Description,
			Duration:      activity.Duration,
			CoachName:     activity.Coach.Name,
			ActivityType:  activity.ActivityType.Name,
			ActivityHours: activitieshours,
		})
	}
	return Activities, nil
}

func (s *ActivityService) GetActivityByKeyword(keyword string) (domain.Activities, error) {
	activitiesDao, err := s.ActivityClient.GetActivityByKeyword(keyword)
	if err != nil {
		return domain.Activities{}, err
	}

	var Activities domain.Activities
	for _, activity := range activitiesDao {
		var activitieshours []domain.ActivityHours
		for _, hour := range activity.ActivityHours {
			activitieshours = append(activitieshours, domain.ActivityHours{
				Id:          hour.ID,
				Day:         hour.Day,
				Hour_start:  hour.Starting_Hour,
				Hour_finish: hour.Finish_hour,
			})
		}
		Activities = append(Activities, domain.Activity{
			ID:            activity.ID,
			Name:          activity.Name,
			Description:   activity.Description,
			Duration:      activity.Duration,
			CoachName:     activity.Coach.Name,
			ActivityType:  activity.ActivityType.Name,
			ActivityHours: activitieshours,
		})
	}
	return Activities, nil
}

func (s *ActivityService) GetActivityById(id int) (domain.Activity, error) {
	activitydao, err := s.ActivityClient.GetActivityById(id)
	if err != nil {
		return domain.Activity{}, err
	}
	var activityDto domain.Activity
	activityDto.ID = activitydao.ID
	activityDto.Name = activitydao.Name
	activityDto.CoachName = activitydao.Coach.Name
	activityDto.Description = activitydao.Description
	activityDto.Duration = activitydao.Duration
	activityDto.ActivityType = activitydao.ActivityType.Name
	var activitieshours []domain.ActivityHours
	for _, hour := range activitydao.ActivityHours {
		activitieshours = append(activitieshours, domain.ActivityHours{
			Id:          hour.ID,
			Day:         hour.Day,
			Hour_start:  hour.Starting_Hour,
			Hour_finish: hour.Finish_hour,
		})
	}
	activityDto.ActivityHours = activitieshours
	return activityDto, nil
}

func (s *ActivityService) CreateActivity(ActivityInfo domain.NewActivity) error {
	var ActivityInfoDAO dao.Activity
	ActivityInfoDAO.Name = ActivityInfo.Name
	ActivityInfoDAO.Description = ActivityInfo.Description
	ActivityInfoDAO.ActivityTypeID = ActivityInfo.ActivityTypeID
	ActivityInfoDAO.Duration = uint(ActivityInfo.Duration)
	ActivityInfoDAO.Capacity = uint(ActivityInfo.Capacity)
	ActivityInfoDAO.CoachID = ActivityInfo.CoachId
	err := s.ActivityClient.CreateActivity(ActivityInfoDAO)
	if err != nil {
		return err
	}
	return nil
}

func (s *ActivityService) CreateActivityHour(ActivityInfoDto domain.NewHour) error {
	var ActivityHourDao dao.ActivityHour
	ActivityHourDao.ActivityID = ActivityInfoDto.ActivityID
	ActivityHourDao.Day = ActivityInfoDto.Day
	ActivityHourDao.Starting_Hour = ActivityInfoDto.Starting_Hour
	ActivityHourDao.Finish_hour = ActivityInfoDto.Finish_hour
	err := s.ActivityClient.CreateActivityHour(ActivityHourDao)
	if err != nil {
		return err
	}
	return nil
}
