package handler

import (
	"net/http"
	"time"

	"github.com/kurosakio/web-pomodoro/internal/template"
)

type TimerHandler struct {
	renderer *template.Renderer
}

func NewTimerHandler(renderer *template.Renderer) *TimerHandler {
	return &TimerHandler{renderer: renderer}
}

func (h *TimerHandler) TimerPage(w http.ResponseWriter, r *http.Request) {
	endTime := time.Now().Add(5 * time.Minute).Unix()

	data := map[string]interface{}{
		"Title":    "Таймер",
		"EndTime":  endTime,
		"Duration": 300,
	}

	h.renderer.Render(w, "timer.html", data)
}
