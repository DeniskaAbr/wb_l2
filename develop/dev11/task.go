package main

import (
	"dev11/api/event"
	v1 "dev11/api/event/delivery/http/v1"
	"dev11/conf"
	"dev11/internal/server"
	"log"
	"os"
)

/*
=== HTTP server ===

Реализовать HTTP сервер для работы с календарем. В рамках задания необходимо работать строго со стандартной HTTP библиотекой.
В рамках задания необходимо:
	1. Реализовать вспомогательные функции для сериализации объектов доменной област
и в JSON.
	2. Реализовать вспомогательные функции для парсинга и валидации параметров методов /create_event и /update_event.
	3. Реализовать HTTP обработчики для каждого из методов API, используя вспомогательные функции и объекты доменной области.
	4. Реализовать middleware для логирования запросов
Методы API: POST /create_event POST /update_event POST /delete_event GET /events_for_day GET /events_for_week GET /events_for_month
Параметры передаются в виде www-url-form-encoded (т.е. обычные user_id=3&date=2019-09-09).
В GET методах параметры передаются через queryString, в POST через тело запроса.
В результате каждого запроса должен возвращаться JSON документ содержащий либо {"result": "..."} в случае успешного выполнения метода,
либо {"error": "..."} в случае ошибки бизнес-логики.

В рамках задачи необходимо:
	1. Реализовать все методы.
	2. Бизнес логика НЕ должна зависеть от кода HTTP сервера.
	3. В случае ошибки бизнес-логики сервер должен возвращать HTTP 503. В случае ошибки входных данных (невалидный int например) сервер должен возвращать HTTP 400. В случае остальных ошибок сервер должен возвращать HTTP 500. Web-сервер должен запускаться на порту указанном в конфиге и выводить в лог каждый обработанный запрос.
	4. Код должен проходить проверки go vet и golint.
*/

func main() {

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	storage := event.NewInmemStore(logger)
	eventsRepository := event.NewMemoryEventRepository(storage, logger)
	eventUsercase := event.NewEventUsecase(eventsRepository, logger)
	handler := v1.NewEventHandler(eventUsercase, logger)

	// var conf server.Configuration

	var config conf.Configuration

	config, _ = config.ReadFromFile("config.json")

	server := server.NewServer(config, logger)

	server.Init(handler.Mux)
	server.Run()
	// server.Stop()
}

//https://stackoverflow.com/questions/23695479/how-to-format-timestamp-in-outgoing-json

// d:\WB\repos\neuralIme\WB_Task_2\develop\dev11\

// https://habr.com/ru/company/ruvds/blog/566198/
// https://github.com/minhthong582000/go-movies-api/blob/master/domain/movie.go
// https://github.com/cvenkman/
// https://habr.com/ru/company/domclick/blog/592087/
// https://github.com/percybolmer/ddd-go
// https://github.com/jojoarianto/go-ddd-api
// https://github.com/lsw45/freedom_base
