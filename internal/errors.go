package internal

import (
	"encoding/json"
	"errors"
)

// кастомные ошибки
var (
	// ошибки config
	ErrEmptyConfigPath = errors.New("path to config file isn't defined")
	// ошибки ДБ
	ErrPingConn      = errors.New("connection to DB isn't opened")
	ErrSearchByLogin = errors.New("can't find information from DB")
	ErrNotExistUser  = errors.New("user doesn't exist")
	ErrAlreadyExist  = errors.New("already exist")
	// ошибки сервиса
	ErrService     = errors.New("user service")
	ErrEmptyEntity = errors.New("empty structure transfer to Service layer")
	ErrDelete      = errors.New("can't delete user")
	ErrAddUser     = errors.New("can't add user")
	ErrUpdateLogin = errors.New("can't change login")
	// ошибки хендлеров
	ErrHandlerAddUser = errors.New("handler add user error")
	ErrReadBodyReq    = errors.New("read request body")
	ErrUnMarshal      = errors.New("read request body")
	ErrEmptyBodyReq   = errors.New("empty request body")
)

// структура ошибки
type AppError struct {
	Err     error  // базовая ошибка
	Message string // сообщение ошибки
	Code    int    // код ошибки
}

// конструктор ошибки
func NewAppError(err error, msg string, code int) *AppError {
	return &AppError{
		Err:     err,
		Message: msg,
		Code:    code,
	}
}

func (e *AppError) Unwrap() error {
	return e.Err
}

func (e *AppError) Error() string {
	return e.Message
}

func (e *AppError) Masrshal() ([]byte, error) {
	errBytes, err := json.Marshal(e)
	if err != nil {
		return nil, err
	}
	return errBytes, nil
}
