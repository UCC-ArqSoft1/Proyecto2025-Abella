package services

import (
	"fmt"

	"github.com/maxabella/appgym/clients"
	"github.com/maxabella/appgym/dao"
	"github.com/maxabella/appgym/domain"
	"github.com/maxabella/appgym/utils"
)

type User interface {
	CreateUser(domain.User)
	Login(domain.UserLoginRequest)
	GetUserByEmail(string)
	CreateUserType()
}

type UserService struct {
	Userclient *clients.UserClient
}

func (s *UserService) Login(user domain.UserLoginRequest) (domain.UserLoginResponse, error) {
	UserDAO, err := s.Userclient.GetUserByEmail(user.Email)
	if err != nil {
		return domain.UserLoginResponse{}, fmt.Errorf("User not found!", err)
	}
	hashed_password := utils.HashSHA256(user.Password)
	if UserDAO.HashedPassword != hashed_password {
		return domain.UserLoginResponse{}, fmt.Errorf("Invalid Password")
	}
	var Response domain.UserLoginResponse
	Response.UserID = UserDAO.ID
	token, err := utils.GenerateJWT(int(UserDAO.ID), UserDAO.UserTypeID)
	if err != nil {
		return domain.UserLoginResponse{}, fmt.Errorf("Error Generating token", err)
	}
	Response.Token = token
	return Response, nil // Final return if everything is ok
}

func (s *UserService) CreateUser(user domain.UserRegister) (domain.UserLoginResponse, error) {
	var userRes domain.UserLoginResponse
	_, err := s.Userclient.GetUserByEmail(user.Email)
	if err != nil {
		hashed_password := utils.HashSHA256(user.Passwordstring)
		userDao := dao.User{
			UserTypeID:     1,
			Email:          user.Email,
			HashedPassword: hashed_password,
			Name:           user.Name,
			LastName:       user.LastName,
			Documentation:  user.Documentation,
			IsCoach:        false,
		}
		id, usertype, err := s.Userclient.CreateUser(userDao)
		if err != nil {
			return domain.UserLoginResponse{}, err
		}

		userRes.UserID = id
		token, err := utils.GenerateJWT(int(id), usertype)
		if err != nil {
			panic(err.Error())
		}
		userRes.Token = token
	} else {
		fmt.Println("User with that email already exists")
	}
	return userRes, nil
}

func (s *UserService) GetUserActivities(userid uint) (domain.Inscriptions, error) {
	UserActivitiesDAO, err := s.Userclient.GetUserActivities(userid)
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

func (s *UserService) Makeinscription(inscriptionData domain.MakeInscription) {
	var inscriptiondao dao.Inscription
	inscriptiondao.UserID = inscriptionData.UserId
	inscriptiondao.ActivityID = inscriptionData.ActivityId
	inscriptiondao.Day = inscriptionData.Day
	inscriptiondao.Starting_Hour = inscriptionData.Hour_start
	inscriptiondao.Finish_hour = inscriptionData.Hour_finish

}
