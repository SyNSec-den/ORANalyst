package control

import (
	"gerrit.o-ran-sc.org/r/ric-plt/xapp-frame/pkg/xapp"
	gofuzzdep "github.com/dvyukov/go-fuzz/go-fuzz-dep"
)

func (c *Control) controlLoop() {
	//Handle receiving message based on message type
	for {
		gofuzzdep.LoopPos()
		msg := <-c.RMR
		xapp.Logger.Debug("Received message type: %d", msg.Mtype)
		switch msg.Mtype {
		case xapp.RIC_INDICATION:
			go c.handleIndication(msg)
		default:
			xapp.Logger.Error("Unknown Message Type '%d', discarding", msg.Mtype)
		}
	}
}
