package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/VoC925/go-testify/data"
	"github.com/VoC925/go-testify/internal/filter"
)

// mainHandle GET обработчик по эндпоинту "cafe/"
func mainHandle(w http.ResponseWriter, req *http.Request) {
	// Экземпляр структуры Filter
	filter, err := filter.New(req)
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(fmt.Sprintf("error: %s", err.Error())))
		return
	}
	// JSON данные кафе
	dataJson, err := json.Marshal(dataByFilter(filter))
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error: marshal JSON data"))
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if _, err = w.Write([]byte(dataJson)); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error: write JSON data"))
		return
	}
}

// DataByFilter возвращает список кафе в городе по фильтру
func dataByFilter(filter *filter.Filter) []string {
	// Список кафе по городу в мапе
	sliceCafe := data.CafeList[filter.City]
	return sliceCafe[:filter.Count]
}
