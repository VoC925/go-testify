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
	route.Post("/user", h.AddUser)
}

// mainHandle GET обработчик по эндпоинту "cafe/"
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
