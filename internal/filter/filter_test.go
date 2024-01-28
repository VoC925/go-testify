package filter

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Структура данных теста
type TestCases struct {
	Name           string
	URL            string
	ExpectedFilter // ожидаемый фильтр
}

type ExpectedFilter struct {
	Filter
}

var SliceCaseFilter = []TestCases{
	{
		Name: "test_1",
		URL:  "http://localhost:8080/cafe?city=moscow&count=1",
		ExpectedFilter: ExpectedFilter{
			Filter: Filter{
				Count: 1,
				City:  `moscow`,
			},
		},
	},
	{
		Name: "test_2",
		URL:  "http://localhost:8080/cafe?city=moscow&count=2",
		ExpectedFilter: ExpectedFilter{
			Filter: Filter{
				Count: 2,
				City:  `moscow`,
			},
		},
	},
	{
		Name: "test_3",
		URL:  "http://localhost:8080/cafe?city=moscow&count=3",
		ExpectedFilter: ExpectedFilter{
			Filter: Filter{
				Count: 3,
				City:  `moscow`,
			},
		},
	},
	{
		Name: "test_4",
		URL:  "http://localhost:8080/cafe?city=moscow&count=4",
		ExpectedFilter: ExpectedFilter{
			Filter: Filter{
				Count: 4,
				City:  `moscow`,
			},
		},
	},
	{
		Name: "test_5",
		URL:  "http://localhost:8080/cafe?city=moscow&count=5",
		ExpectedFilter: ExpectedFilter{
			Filter: Filter{
				Count: 4,
				City:  `moscow`,
			},
		},
	},
	{
		Name: "test_6",
		URL:  "http://localhost:8080/cafe?city=moscow&count=651",
		ExpectedFilter: ExpectedFilter{
			Filter: Filter{
				Count: 4,
				City:  `moscow`,
			},
		},
	},
}

// Тест функции New, возвращающей фильтр на основе запроса
func TestFilterNew(t *testing.T) {
	for _, testCase := range SliceCaseFilter {
		req := httptest.NewRequest(
			http.MethodGet,
			testCase.URL,
			nil,
		)
		filterByFunc, _ := New(req)
		assert.Equal(t, testCase.Count, filterByFunc.Count, testCase.Name)
		assert.Equal(t, testCase.City, filterByFunc.City, testCase.Name)
	}
}
