package monitor

import "github.com/codeberserker/bus/pkg/event"

type Monitor struct {
	Config Config
}

func (_ *Monitor) Run() (<-chan *event.Event, error) {
	out := make(chan *event.Event)
	go func() {
		for {
			e := new(event.Event)
			out <- e
		}
	}()
	return out, nil
}
