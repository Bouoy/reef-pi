package doser

import (
	"github.com/reef-pi/reef-pi/controller/utils"
	"log"
	"time"
)

type Runner struct {
	pin      int
	duration time.Duration
	speed    int
	vv       utils.PWM
}

func (r *Runner) Run() {
	log.Println("doser sub system: setting pwm pin:", r.pin, "at speed", r.speed)
	if err := r.vv.On(r.pin); err != nil {
		log.Println("ERROR: doser sub system: failed to enable pwm pin:", r.pin, "Error:", err)
		return
	}
	r.vv.Set(r.pin, r.speed)
	select {
	case <-time.After(r.duration * time.Second):
		r.vv.Set(r.pin, 0)
		log.Println("doser sub system: setting pwm pin:", r.pin, "at speed", 0)
	}
	if err := r.vv.Off(r.pin); err != nil {
		log.Println("ERROR: doser sub system: failed to disable pwm pin:", r.pin, "Error:", err)
		return
	}
}
