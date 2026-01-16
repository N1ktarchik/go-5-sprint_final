// Алгоритм реализации:
// 	Создайте структуру сервера с полями для логгера (log.Logger) и http-сервера (http.Server).
// 	Создайте функцию, в которой нужно создать http-роутер. Функция принимает log.Logger и возвращает экземпляр структуры вашего сервера.
// 	Зарегистрируйте ваши хендлеры в http-роутере.
// 	Создайте экземпляр структуры http.Server. Для настройки вашего сервера используйте следующие поля:
// 		Addr — используйте порт 8080.
// 		Handler — передайте ваш http-роутер.
// 		ErrorLog — передайте ваш логгер.
// 		ReadTimeout — таймаут для чтения. 5 секунд.
// 		WriteTimeout — таймаут для записи. 10 секунд.
// 		IdleTimeout — таймаут ожидания следующего запроса. 15 секунд.
// 	Верните ссылку на ваш сервер.

package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logs *log.Logger
	Serv *http.Server
}

func MakeRouter(Logs *log.Logger) *Server {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.MainHandler)
	mux.HandleFunc("/upload", handlers.DownoloadHandler)

	serv := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ErrorLog:     Logs,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	return &Server{
		Logs: Logs,
		Serv: &serv,
	}

}
