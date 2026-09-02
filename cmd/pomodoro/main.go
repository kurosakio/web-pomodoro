package main

import (
	"log"
	"net/http"

	"github.com/kurosakio/web-pomodoro/internal/handler"
	"github.com/kurosakio/web-pomodoro/internal/template"
)

func main() {
	renderer, err := template.NewRenderer()
	if err != nil {
		log.Fatal(err)
	}
	timerHandler := handler.NewTimerHandler(renderer)

	http.HandleFunc("/", timerHandler.TimerPage)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
