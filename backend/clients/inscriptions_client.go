package clients

import (
	"errors"
	"fmt"

	"github.com/maxabella/appgym/dao"
	"github.com/maxabella/appgym/domain"
	"gorm.io/gorm"
)

type Inscriptions interface {
	GetUserActivities(userid uint) (dao.Inscriptions, error)
	MakeInscription(userid uint, activityid uint) error
}

type InscriptionClient struct {
	DbClient *Mysql_Client
}

func (s *InscriptionClient) GetUserActivities(userid uint) (dao.Inscriptions, error) {
	var UserActivitiesDAO dao.Inscriptions
	if err := s.DbClient.db.Where("user_id = ?", userid).Preload("Activity").Find(&UserActivitiesDAO); err.Error != nil {
		return dao.Inscriptions{}, err.Error
	}
	return UserActivitiesDAO, nil
}

func (s *InscriptionClient) MakeInscription(inscriptionData domain.MakeInscription) error {
	// First we check that the inscription doesn't exist already, we don't want the same user to have the same inscription twice. The frontend has no idea that they are the same.
	var inscriptiondao dao.Inscription
	var i dao.Inscription
	inscriptiondao.UserID = inscriptionData.UserId
	inscriptiondao.ActivityID = inscriptionData.ActivityId
	inscriptiondao.Day = inscriptionData.Day
	inscriptiondao.Starting_Hour = inscriptionData.Hour_start
	inscriptiondao.Finish_hour = inscriptionData.Hour_finish
	if err := s.DbClient.db.Where("user_id = ? AND activity_id = ? AND day = ? AND starting_hour = ? AND finish_hour = ?", inscriptionData.UserId, inscriptionData.ActivityId, inscriptionData.Day, inscriptionData.Hour_start, inscriptionData.Hour_finish).First(&i).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		fmt.Println("user can be added to activity")
		err := s.DbClient.db.Create(&inscriptiondao)
		if err != nil {
			return err.Error
		}
	} else {
		return fmt.Errorf("El usuario ya esta inscripto a esta actividad")
	}
	return nil
}
