package main

import "github.com/maxabella/appgym/app"

func main() {
	app.Start()
}

/*
	Myuser := domain.UserRegister{
		UserTypeId:     1,
		Email:          "maximoabella13@gmail.com",
		Passwordstring: "12345678",
		Name:           "Maximo",
		LastName:       "Abella",
		Documentation:  45843938,
	}

	Usuario, err := services.Userservice.CreateUser(Myuser)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(Usuario)
	fmt.Println("Hello world")
*/

/*
	ID             uint   `json:"id"`
	UserTypeId     uint   `json:"usertypeid"`
	Email          string `json:"email"`
	Passwordstring string `json:"password"`
	Name           string `json:"name"`
	LastName       string `json:"lastname"`
	Documentation  int    `json:"documentation"`
*/
