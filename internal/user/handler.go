package user

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VoC925/go-testify/internal"
	"github.com/VoC925/go-testify/internal/handlers"
	"github.com/VoC925/go-testify/pkg/logging"
	"github.com/go-chi/chi/v5"
)

var _ handlers.Handler = &HandlerUser{}

type HandlerUser struct {
	logger  *logging.Logger
	service Service
}

func NewHandlerUser(service Service, logger *logging.Logger) handlers.Handler {
	return &HandlerUser{
		logger:  logger,
		service: service,
	}
}

func (h *HandlerUser) Register(route *chi.Mux) {
	route.Get("/user/{login}", h.GetUser)
	route.Get("/users", h.GetAllUsers)
	route.Post("/user", h.AddUser)
	route.Post("/user/update/{login}/{newLogin}", h.UpdateLoginUser)
	route.Delete("/user/delete/{login}", h.DeleteUser)
}

// GetAllUsers обработчик по эндпоинту "users/"
func (h *HandlerUser) GetAllUsers(w http.ResponseWriter, req *http.Request) {
	var errStr string
	// сервис
	users, err := h.service.GetAllUsers()
	if err != nil {
		errStr = fmt.Sprintf("%s: %s",
			internal.ErrHandlerGetAllUsers,
			err)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errStr))
		return
	}
	jsonData, err := json.Marshal(users)
	if err != nil {
		errStr = fmt.Sprintf("%s: %s: %s",
			internal.ErrHandlerGetUser,
			internal.ErrMarshal,
			err)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errStr))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// GetUser обработчик по эндпоинту "users/"
func (h *HandlerUser) GetUser(w http.ResponseWriter, req *http.Request) {
	// параметр из request
	var errStr string
	login := chi.URLParam(req, "login")
	if login == "" {
		errStr = fmt.Sprintf("%s: %s",
			internal.ErrHandlerGetUser,
			internal.ErrEmptyGetParam)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errStr))
		return
	}
	// сервис
	userObj, err := h.service.GetUserByLogin(login)
	if err != nil {
		errStr = fmt.Sprintf("%s: %s",
			internal.ErrHandlerGetUser,
			err)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errStr))
		return
	}
	jsonData, err := json.Marshal(userObj)
	if err != nil {
		errStr = fmt.Sprintf("%s: %s: %s",
			internal.ErrHandlerGetUser,
			internal.ErrMarshal,
			err)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errStr))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

// DeleteUser обработчик по эндпоинту "users/"
func (h *HandlerUser) DeleteUser(w http.ResponseWriter, req *http.Request) {
	// параметр из request
	var errStr string
	deleteLogin := chi.URLParam(req, "login")
	if deleteLogin == "" {
		errStr = fmt.Sprintf("%s: %s",
			internal.ErrHandlerDeleteUser,
			internal.ErrEmptyGetParam)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errStr))
		return
	}
	// сервис
	err := h.service.DeleteUser(deleteLogin)
	if err != nil {
		errStr = fmt.Sprintf("%s: %s",
			internal.ErrHandlerDeleteUser,
			err)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errStr))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Пользователь удален"))
}

// AddUser POST обработчик по эндпоинту "users/"
func (h *HandlerUser) AddUser(w http.ResponseWriter, req *http.Request) {
	var (
		buf    bytes.Buffer
		errStr string
	)
	_, err := buf.ReadFrom(req.Body) // Чтение JSON из тела запроса
	if err != nil {
		errStr = fmt.Sprintf("%s: %s: %s",
			internal.ErrHandlerAddUser,
			internal.ErrReadBodyReq,
			err)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errStr))
		return
	}
	if buf.String() == "" {
		errStr = fmt.Sprintf("%s: %s",
			internal.ErrHandlerAddUser,
			internal.ErrEmptyBodyReq)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errStr))
		return
	}
	defer req.Body.Close()
	// указатель на структуру нового пользователя userDTO
	newUser := new(UserDTO)
	// UnMarshal buf
	err = json.Unmarshal(buf.Bytes(), newUser)
	if err != nil {
		errStr = fmt.Sprintf("%s: %s: %s",
			internal.ErrHandlerAddUser,
			internal.ErrUnMarshal,
			err)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errStr))
		return
	}
	// запрос к сервису
	_, err = h.service.RegisterNewUser(newUser)
	if err != nil {
		errStr = fmt.Sprintf("%s: %s",
			internal.ErrHandlerAddUser,
			err)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errStr))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Пользователь добавлен"))
}

// DeleteUser обработчик по эндпоинту "users/"
func (h *HandlerUser) UpdateLoginUser(w http.ResponseWriter, req *http.Request) {

	var errStr string
	// параметр из request
	login := chi.URLParam(req, "login")
	newLogin := chi.URLParam(req, "newLogin")

	if login == "" || newLogin == "" {
		errStr = fmt.Sprintf("%s: %s",
			internal.ErrHandlerUpdateLoginUser,
			internal.ErrEmptyGetParam)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(errStr))
		return
	}
	// сервис
	err := h.service.ChangeLogin(login, newLogin)
	if err != nil {
		errStr = fmt.Sprintf("%s: %s",
			internal.ErrHandlerUpdateLoginUser,
			err)
		h.logger.Error(errStr)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(errStr))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Логин изменен"))
}
