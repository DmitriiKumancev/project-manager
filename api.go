// создаем fpi для запуска сервера и всего что с ним связано
package main

import (
	"net/http"
	"project-manager/adapters"

	"github.com/gorilla/mux"
)

type APIServer struct {
	addr  string
	store Store //хранилище, которое имеет все подключения к базе данных
}

var log = adapters.GetLogger()

// создадим конструктор для сервиса
func NewApiServer(addr string, store Store) *APIServer {
	return &APIServer{addr: addr, store: store}
}

// единственный метод, который будет использоваться на этом сервере (инициализация маршрутизатора, а затем зарегестрировать все службы и их зависимости)
func (s *APIServer) Serve() {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	// регистрируем наши услуги
	log.Info("Starting the API server at", s.addr)

	// Запускаем HTTP-сервер и проверяем на наличие ошибки
	err := http.ListenAndServe(s.addr, subrouter)
	if err != nil {
		// Логируем ошибку, если она произошла
		log.Error("Failed to start HTTP server", err)
	}

}
