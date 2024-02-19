package user

type User struct {
	Name          string `json:"name"`
	Login         string `json:"login"`
	PasswordHash  string `json:"hash"`
	CreatedAt     string `json:"date_created_at"`
	LastChangedAt string `json:"date_updated_at"`
}

type UserDTO struct {
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

/*
{
	"name": "Alex",
	"login": "Alex23",
	"password": "pass23"
}
*/
