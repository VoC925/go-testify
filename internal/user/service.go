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
	ChangeLogin(*UserDTO, string) error    // Обновить данные пользователя
	DeleteUser(*UserDTO) error             // Удалить пользователя
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
func (u *UserService) ChangeLogin(userDTO *UserDTO, newLogin string) error {
	// проверка на существование пользователя с таким же логином
	_, err := u.db.GetByLogin(userDTO.Login)
	if errors.Is(err, internal.ErrNotExistUser) {
		return fmt.Errorf("%s: %s", internal.ErrUpdateLogin, err)
	}
	err = u.db.UpdateLogin(userDTO.Login, newLogin)
	if err != nil {
		return fmt.Errorf("%s: %s", internal.ErrUpdateLogin, err)
	}
	u.logger.Infof("login changed successfully, before: %s ; now: %s", userDTO.Login, newLogin)
	return nil
}

// удаление пользователя
func (u *UserService) DeleteUser(userDTO *UserDTO) error {
	// проверка на существование пользователя с таким же логином
	_, err := u.db.GetByLogin(userDTO.Login)
	if errors.Is(err, internal.ErrNotExistUser) {
		return fmt.Errorf("%s: %s", internal.ErrDelete, err)
	}
	err = u.db.Delete(userDTO.Login)
	if err != nil {
		return fmt.Errorf("%s: %s", internal.ErrDelete, err)
	}
	u.logger.Infof("user with login: %s deleted successfully", userDTO.Login)
	return nil
}
