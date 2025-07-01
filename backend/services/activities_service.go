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
	EditActivity(domain.Activity) error
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

func (s *ActivityService) CreateCategory(CategoryDto domain.NewCategory) error {
	var NewCategory dao.ActivityType
	NewCategory.Name = CategoryDto.Name
	err := s.ActivityClient.CreateCategory(NewCategory)
	if err != nil {
		return err
	}
	return nil
}

func (s *ActivityService) EditActivity(ActivityInfoDto domain.UpdateActivity) error {
	var ActivityinfoDao dao.Activity
	ActivityinfoDao.ID = ActivityInfoDto.ID
	ActivityinfoDao.Name = ActivityInfoDto.Name
	ActivityinfoDao.Description = ActivityInfoDto.Description
	ActivityinfoDao.ActivityTypeID = ActivityInfoDto.ActivityTypeID
	ActivityinfoDao.CoachID = ActivityInfoDto.CoachId
	ActivityinfoDao.Duration = uint(ActivityInfoDto.Duration)
	ActivityinfoDao.Capacity = uint(ActivityInfoDto.Capacity)
	err := s.ActivityClient.EditActivity(ActivityinfoDao)
	if err != nil {
		return err
	}
	return nil
}

func (s *ActivityService) EditHour(ActivityInfo domain.UpdateActivityHours) error {

	var HourEditDao dao.ActivityHour
	var InscriptionEditDao dao.Inscription

	fmt.Println("DTO: ", ActivityInfo)

	HourEditDao.ID = ActivityInfo.Id
	HourEditDao.ActivityID = ActivityInfo.ActivityId
	HourEditDao.Day = ActivityInfo.Day
	HourEditDao.Starting_Hour = ActivityInfo.Hour_start
	HourEditDao.Finish_hour = ActivityInfo.Hour_finish

	InscriptionEditDao.ActivityID = ActivityInfo.ActivityId
	InscriptionEditDao.Day = ActivityInfo.Day
	InscriptionEditDao.Starting_Hour = ActivityInfo.Hour_start
	InscriptionEditDao.Finish_hour = ActivityInfo.Hour_finish

	fmt.Println("HourDAO: ", HourEditDao)
	fmt.Println("InscriptionDao: ", InscriptionEditDao)

	err := s.ActivityClient.EditActivityHour(HourEditDao)
	if err != nil {
		return err
	}

	err = s.ActivityClient.UpdateInscriptions(InscriptionEditDao)
	if err != nil {
		return err
	}
	return nil
}

func (s *ActivityService) DeleteActivity(id uint) error {
	err := s.ActivityClient.DeleteActivity(id)
	if err != nil {
		return err
	}
	return nil
}
