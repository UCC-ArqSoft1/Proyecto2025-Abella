package domain

type User struct {
	ID              uint   `json:"id"`
	UserTypeId      uint   `json:"usertypeid"`
	Email           string `json:"email"`
	Hashed_Password string `json:"password"`
	Name            string `json:"name"`
	LastName        string `json:"lastname"`
	Documentation   int    `json:"documentation"`
}

type UserRegister struct {
	Email          string `json:"email"`
	Passwordstring string `json:"password"`
	Name           string `json:"name"`
	LastName       string `json:"lastname"`
	Documentation  int    `json:"documentation"`
}

type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponse struct {
	UserID uint   `json:"userid"`
	Token  string `json:"token"`
}

type Coach struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type Coaches []Coach
