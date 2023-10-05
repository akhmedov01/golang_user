package models

type User struct {
	Id         string `json:"id"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	Name       string `json:"name"`
	Age        int    `json:"age"`
	Created_at string `json:"created_at"`
	Updated_at string `json:"updated_at"`
}

type CreateUser struct {
	Login    string `json:"login"`
	Password string `json:"password"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
}

type UpdateUser struct {
	Login string `json:"login"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
}

type IdRequest struct {
	Id string
}

type GetAllUserRequest struct {
	Page  int
	Limit int
	Name  string
}

type GetAllUser struct {
	Users []User
	Count int
}

type GetByLoginReq struct {
	Login string
}

type GetByLoginRes struct {
	Id       string
	Password string
}

type LoginReq struct {
	Login    string
	Password string
}

type LoginRes struct {
	Token string
}
