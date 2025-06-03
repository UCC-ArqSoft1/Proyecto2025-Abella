package services

import (
	"github.com/maxabella/appgym/clients"
	"github.com/maxabella/appgym/domain"
)

type Inscription interface {
	GetUserActivities(userid uint) (domain.Inscriptions, error)
	Makeinscription(inscriptionData domain.MakeInscription) error
}

type InscriptionService struct {
	InscriptionClient *clients.InscriptionClient
}

func (s *InscriptionService) GetUserActivities(userid uint) (domain.Inscriptions, error) {
	UserActivitiesDAO, err := s.InscriptionClient.GetUserActivities(userid)
	if err != nil {
		return domain.Inscriptions{}, err
	}
	var InscriptionsDTO domain.Inscriptions
	for _, Inscription := range UserActivitiesDAO {
		InscriptionsDTO = append(InscriptionsDTO, domain.Inscription{
			ID:          Inscription.ID,
			Name:        Inscription.Activity.Name,
			CoachName:   Inscription.Activity.Coach.Name,
			Duration:    Inscription.Activity.Duration,
			Day:         Inscription.Day,
			Hour_start:  Inscription.Starting_Hour,
			Hour_finish: Inscription.Finish_hour,
		})
	}
	return InscriptionsDTO, nil
}

func (s *InscriptionService) Makeinscription(inscriptionData domain.MakeInscription) error {
	err := s.InscriptionClient.MakeInscription(inscriptionData)
	if err != nil {
		return err
	}
	return nil
}
