package deamon

import (
	"time"

	"github.com/soyoslab/soy_log_collector/pkg/container/ring"
)

// Listen does background jobs
// ring: contains jobs
// fn: function
// duration: sleep every x period(sec)
func Listen(ring *ring.Ring, fn func(...interface{}), duration time.Duration) {
	for true {
		for true {
			next, err := ring.Pop()
			if err == nil {
				fn(next)
			}
		}
		time.Sleep(time.Second * duration)
	}
}
