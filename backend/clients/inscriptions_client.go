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
	DeleteUserInscription(id uint) error
	CountInscriptionByActivity(activityid int) error
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

func (s *InscriptionClient) CountInscriptionByActivity(activityid int) error {
	var count int64
	err := s.DbClient.db.Model(dao.Inscription{}).Where("activity_id = ?", activityid).Count(&count)
	if err.Error != nil {
		return err.Error
	}
	fmt.Println(count)
	var TheActivity dao.Activity
	err = s.DbClient.db.Where("id = ?", activityid).First(&TheActivity)
	if err != nil {
		return err.Error
	}
	if TheActivity.Capacity == uint(count) {
		return fmt.Errorf("No hay mas cupo en esta actividad con este horario")
	}
	return nil
}

func (s *InscriptionClient) MakeInscription(inscriptionData domain.MakeInscription) error {
	// Primero nos aseguramos de que una inscripcion como la que se esta solicitando no exista ya en la base de datos.
	var inscriptiondao dao.Inscription
	var i dao.Inscription
	inscriptiondao.UserID = inscriptionData.UserId
	inscriptiondao.ActivityID = inscriptionData.ActivityId
	inscriptiondao.Day = inscriptionData.Day
	inscriptiondao.Starting_Hour = inscriptionData.Hour_start
	inscriptiondao.Finish_hour = inscriptionData.Hour_finish
	if err := s.DbClient.db.Where("user_id = ? AND activity_id = ? AND day = ? AND starting_hour = ? AND finish_hour = ?", inscriptionData.UserId, inscriptionData.ActivityId, inscriptionData.Day, inscriptionData.Hour_start, inscriptionData.Hour_finish).First(&i).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		err = s.CountInscriptionByActivity(int(inscriptionData.ActivityId))
		if err != nil {
			return err
		}
		fmt.Println("user can be added to activity") // DEBUG
		err := s.DbClient.db.Create(&inscriptiondao)
		if err != nil {
			return err.Error
		}
	} else {
		return fmt.Errorf("El usuario ya esta inscripto a esta actividad") // Si el usuario ya esta inscripto entonces devolvemos el codigo para el front
	}
	return nil
}

func (s *InscriptionClient) DeleteUserInscription(id uint) error {
	err := s.DbClient.db.Where("id = ?", id).Delete(dao.Inscription{})
	if err.Error != nil {
		return err.Error
	}
	return nil
}
