package services

import (
	"fmt"

	"github.com/maxabella/appgym/clients"
	"github.com/maxabella/appgym/dao"
	"github.com/maxabella/appgym/domain"
	"github.com/maxabella/appgym/utils"
)

type User interface {
	CreateUser(user domain.UserRegister) (domain.UserLoginResponse, error)
	Login(user domain.UserLoginRequest) (domain.UserLoginResponse, error)
	GetUserByEmail(string)
	CreateUserType()
	GetCoaches() (domain.Coaches, error)
}

type UserService struct {
	Userclient *clients.UserClient
}

func (s *UserService) Login(user domain.UserLoginRequest) (domain.UserLoginResponse, error) {
	UserDAO, err := s.Userclient.GetUserByEmail(user.Email)
	if err != nil {
		return domain.UserLoginResponse{}, fmt.Errorf("User not found! %v", err)
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

func (s *UserService) GetCoaches() (domain.Coaches, error) {
	var Coaches domain.Coaches
	CoachesInfoDto, err := s.Userclient.GetCoaches()
	if err != nil {
		return domain.Coaches{}, err
	}
	for _, Coach := range CoachesInfoDto {
		Coaches = append(Coaches, domain.Coach{
			ID:   Coach.ID,
			Name: Coach.Name + " " + Coach.LastName,
		})
	}
	return Coaches, nil
}
