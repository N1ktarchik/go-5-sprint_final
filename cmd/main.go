package main

import (
	"log"
	"net/http"
	"os"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {

	file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	mylog := log.New(file, `serv `, log.LstdFlags|log.Lshortfile)

	mylog.Println(`Start server`)
	defer mylog.Println(`Finish server`)

	server := server.MakeRouter(mylog)

	err = http.ListenAndServe(server.Serv.Addr, server.Serv.Handler)
	if err != nil {
		mylog.Fatal(err)
	}
}
