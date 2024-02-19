package user

// интерфейс хранилища
type Storage interface {
	Add(*User) (int, error)           // Добавление пользователя
	GetByLogin(string) (*User, error) // Получить пользователя по логину
	GetAllUsers() ([]*User, error)    // Получить всех пользователей
	Delete(string) error              // Удалить пользователя
	UpdateLogin(string, string) error // Изменить логин
	UpdateTime(string) error          // Обновить время изменений
}
