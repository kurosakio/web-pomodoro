package handler

import (
	"fmt"
	"net/http"
	"strconv"
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
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	workMinutes, _ := strconv.Atoi(r.FormValue("workMinutes"))
	breakMinutes, _ := strconv.Atoi(r.FormValue("breakMinutes"))
	longBreakMinutes, _ := strconv.Atoi(r.FormValue("longBreakMinutes"))
	fmt.Println("workMinutes =", workMinutes)
	fmt.Println("breakMinutes =", breakMinutes)
	fmt.Println("longBreakMinutes =", longBreakMinutes)

	if workMinutes == 0 {
		workMinutes = 25
	}
	if breakMinutes == 0 {
		breakMinutes = 5
	}
	if longBreakMinutes == 0 {
		longBreakMinutes = 15
	}
	endTime := time.Now().Add(time.Duration(workMinutes) * time.Minute).Unix()

	data := map[string]interface{}{
		"Title":             "Таймер",
		"EndTime":           endTime,
		"WorkDuration":      workMinutes,
		"BreakDuration":     breakMinutes,
		"LongBreakDuration": longBreakMinutes,
	}

	h.renderer.Render(w, "timer.html", data)
}
