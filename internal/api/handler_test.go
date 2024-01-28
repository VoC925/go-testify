package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Структура тестовых кейсов
type TestCases struct {
	Name             string // название теста
	URL              string
	ExpectedResponse // ожидаемый ответ
}

type ExpectedResponse struct {
	StatusCode int
	Body       string
}

// Слайс тестов на верные данные
var SliceCaseStatusCheckOK = []TestCases{
	{
		Name: "test_1",
		URL:  "http://localhost:8080/cafe?city=moscow&count=0",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusOK,
			Body:       `[]`,
		},
	},
	{
		Name: "test_2",
		URL:  "http://localhost:8080/cafe?city=moscow&count=1",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusOK,
			Body:       `["Мир кофе"]`,
		},
	},
	{
		Name: "test_3",
		URL:  "http://localhost:8080/cafe?city=moscow&count=2",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusOK,
			Body:       `["Мир кофе", "Сладкоежка"]`,
		},
	},
	{
		Name: "test_4",
		URL:  "http://localhost:8080/cafe?city=moscow&count=3",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusOK,
			Body:       `["Мир кофе", "Сладкоежка", "Кофе и завтраки"]`,
		},
	},
	{
		Name: "test_5",
		URL:  "http://localhost:8080/cafe?city=moscow&count=4",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusOK,
			Body:       `["Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"]`,
		},
	},
	{
		Name: "test_6",
		URL:  "http://localhost:8080/cafe?city=moscow&count=36435",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusOK,
			Body:       `["Мир кофе", "Сладкоежка", "Кофе и завтраки", "Сытый студент"]`,
		},
	},
	{
		Name: "test_7",
		URL:  "http://localhost:8080/cafe",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `error: empty GET param`,
		},
	},
}

// Тест на верные запросы и отсутствие GET параметров
func TestMainHandleStatusCheck(t *testing.T) {
	for _, testCase := range SliceCaseStatusCheckOK {
		req := httptest.NewRequest(
			http.MethodGet,
			testCase.URL,
			nil,
		)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(w, req)
		assert.Equal(t, testCase.StatusCode, w.Code, testCase.Name)
	}
}

// Слайс тестов на неверные данные параметра city
var SliceCaseStatusBadRequestCity = []TestCases{
	{
		Name: "test_1",
		URL:  "http://localhost:8080/cafe?count=3",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `error: city value is missing`,
		},
	},
	{
		Name: "test_2",
		URL:  "http://localhost:8080/cafe?city=&count=3",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `error: city value is missing`,
		},
	},
	{
		Name: "test_3",
		URL:  "http://localhost:8080/cafe?city=city&count=3",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `error: wrong city value`,
		},
	},
	{
		Name: "test_4",
		URL:  "http://localhost:8080/cafe?city=1&count=3",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `error: wrong city value`,
		},
	},
	{
		Name: "test_5",
		URL:  "http://localhost:8080/cafe?city=selknsdv&count=3",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `error: wrong city value`,
		},
	},
}

// Тест на случай не корректно введенного значения city
func TestMainHandleBadRequestCity(t *testing.T) {
	for _, testCase := range SliceCaseStatusBadRequestCity {
		req := httptest.NewRequest(
			http.MethodGet,
			testCase.URL,
			nil,
		)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(w, req)
		// Проверка статуса ответа
		require.Equal(t, testCase.StatusCode, w.Code)
		// Проверка тела ответа
		assert.Equal(t, testCase.Body, w.Body.String())
	}
}

// Слайс тестов неверных параметров count
var SliceCaseStatusBadRequestCount = []TestCases{
	{
		Name: "test_1",
		URL:  "http://localhost:8080/cafe?city=moscow",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `error: count value is missing`,
		},
	},
	{
		Name: "test_2",
		URL:  "http://localhost:8080/cafe?city=moscow&count=",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `error: count value is missing`,
		},
	},
	{
		Name: "test_3",
		URL:  "http://localhost:8080/cafe?city=moscow&count=jebf",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `error: wrong count value`,
		},
	},
	{
		Name: "test_4",
		URL:  "http://localhost:8080/cafe?city=moscow&count=-1",
		ExpectedResponse: ExpectedResponse{
			StatusCode: http.StatusBadRequest,
			Body:       `error: wrong count < 0`,
		},
	},
}

// Тест на случай не корректно введенного значения count
func TestMainHandleBadRequstCount(t *testing.T) {
	for _, testCase := range SliceCaseStatusBadRequestCount {
		req := httptest.NewRequest(
			http.MethodGet,
			testCase.URL,
			nil,
		)
		w := httptest.NewRecorder()
		handler := http.HandlerFunc(mainHandle)
		handler.ServeHTTP(w, req)
		// Проверка статуса ответа
		require.Equal(t, testCase.StatusCode, w.Code)
		// Проверка тела ответа
		assert.Equal(t, testCase.Body, w.Body.String())
	}
}
