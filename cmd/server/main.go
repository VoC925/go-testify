package main

import (
	"fmt"
	"os"

	"github.com/VoC925/go-testify/internal/api"
)

func main() {
	// Создание экземпляра структура Server
	s := api.New("localhost", ":8080")
	// Запуск сервера
	if err := s.Start(); err != nil {
		fmt.Printf("error: %s", err.Error())
		os.Exit(1)
	}
}
