package services

import (
	"github.com/maxabella/appgym/clients"
	"github.com/maxabella/appgym/domain"
)

/*
 */

type Activity interface {
	GetActivities() (domain.Activities, error)
	GetActivityByKeyword(keyword string) (domain.Activities, error)
}

type ActivityService struct {
	ActivityClient *clients.ActivityClient
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
