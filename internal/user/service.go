package user

import (
	"errors"
	"fmt"
	"time"

	"github.com/VoC925/go-testify/internal"
	"github.com/VoC925/go-testify/pkg/logging"
	"github.com/VoC925/go-testify/pkg/utils"
)

// интерфейс сервиса
type Service interface {
	RegisterNewUser(*UserDTO) (int, error) // Зарегистрировать нового пользователя
	ChangeLogin(string, string) error      // Обновить данные пользователя
	DeleteUser(string) error               // Удалить пользователя
	GetUserByLogin(string) (*User, error)  // Получить пользователя по логину
	GetAllUsers() ([]*User, error)         // Получить всех пользователей
}

// Конструктор
func NewService(db Storage, logger *logging.Logger) Service {
	return &UserService{
		db:     db,
		logger: logger,
	}
}

type UserService struct {
	db     Storage
	logger *logging.Logger
}

// Регистрация нового пользователя
func (u *UserService) RegisterNewUser(userDTO *UserDTO) (int, error) {
	if userDTO == nil {
		return 0, fmt.Errorf("%s: %s", internal.ErrAddUser, internal.ErrEmptyEntity)
	}
	// проверка на существование пользователя с таким же логином
	_, err := u.db.GetByLogin(userDTO.Login)
	if err == nil {
		return 0, fmt.Errorf("%s: %s", internal.ErrAddUser, internal.ErrAlreadyExist)
	}

	// структура пользователя User
	newUser := &User{
		Name:          userDTO.Name,
		Login:         userDTO.Login,
		PasswordHash:  utils.PasswordToHash(userDTO.Password),
		CreatedAt:     time.Now().Format(time.DateTime),
		LastChangedAt: time.Now().Format(time.DateTime),
	}

	// Добавление нового пользователя
	idUser, err := u.db.Add(newUser)
	if err != nil {
		return 0, fmt.Errorf("%s: %s", internal.ErrAddUser, err)
	}
	u.logger.Infof("user with id:%d add successfully", idUser)
	return idUser, nil
}

// изменение данных пользователя
func (u *UserService) ChangeLogin(login string, newLogin string) error {
	// проверка на существование пользователя с таким же логином
	_, err := u.db.GetByLogin(login)
	if errors.Is(err, internal.ErrNotExistUser) {
		return fmt.Errorf("%s: %s", internal.ErrUpdateLogin, err)
	}
	// изменение логина
	err = u.db.UpdateLogin(login, newLogin)
	if err != nil {
		return fmt.Errorf("%s: %s", internal.ErrUpdateLogin, err)
	}
	// изменение времени последних изменений
	err = u.db.UpdateTime(newLogin)
	if err != nil {
		return fmt.Errorf("%s: %s", internal.ErrUpdateLogin, err)
	}
	u.logger.Infof("login changed successfully, before: %s ; now: %s", login, newLogin)
	return nil
}

// удаление пользователя
func (u *UserService) DeleteUser(login string) error {
	// проверка на существование пользователя с таким же логином
	_, err := u.db.GetByLogin(login)
	if errors.Is(err, internal.ErrNotExistUser) {
		return fmt.Errorf("%s: %s", internal.ErrDelete, err)
	}
	err = u.db.Delete(login)
	if err != nil {
		return fmt.Errorf("%s: %s", internal.ErrDelete, err)
	}
	u.logger.Infof("user with login: %s deleted successfully", login)
	return nil
}

// извлчение пользователя
func (u *UserService) GetUserByLogin(login string) (*User, error) {
	// проверка на существование пользователя с таким же логином
	userObj, err := u.db.GetByLogin(login)
	if errors.Is(err, internal.ErrNotExistUser) {
		return nil, fmt.Errorf("%s: %s", internal.ErrGetUser, err)
	}
	return userObj, nil
}

func (u *UserService) GetAllUsers() ([]*User, error) {
	users, err := u.db.GetAllUsers()
	if err != nil {
		return nil, fmt.Errorf("%s: %s", internal.ErrGetUser, err)
	}
	if users == nil {
		return nil, fmt.Errorf("%s: %s", internal.ErrGetUser, internal.ErrEmptyList)
	}
	return users, nil
}
