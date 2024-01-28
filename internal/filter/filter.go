package filter

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/VoC925/go-testify/data"
)

// Filter - структура фильтра данных из GET параметров
type Filter struct {
	Count int    // количество кафе, которые нужно вернуть
	City  string // город, в котором нужно найти кафе
}

// New возвращает экзмепляр структуры Filter на основе GET запроса
func New(req *http.Request) (*Filter, error) {
	var filter Filter // экземпляр фильтра
	// Значение города из запроса
	city := req.URL.Query().Get("city")
	// Проверка на невведенный параметр
	if city == "" {
		return nil, fmt.Errorf("city value is missing")
	}
	// Слайс из названий кафе в городе city
	sliceCafe, ok := data.CafeList[city]
	// Проверка на существование города в мапе
	if !ok {
		return nil, fmt.Errorf("wrong city value")
	}
	filter.City = city
	// Значение количества кафе в городе city
	countStr := req.URL.Query().Get("count")
	// Проверка на невведенный параметр
	if countStr == "" {
		return nil, fmt.Errorf("count value is missing")
	}
	count, err := strconv.Atoi(countStr)
	if err != nil {
		return nil, fmt.Errorf("wrong count value")
	}
	// Случай, когда count < 0
	if count < 0 {
		return nil, fmt.Errorf("wrong count < 0")
	}
	// Случай, когда count больше, чем есть в списке
	if count > len(sliceCafe) {
		count = len(sliceCafe)
	}
	filter.Count = count

	return &filter, nil
}
