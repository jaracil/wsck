// +build js

package websocket

import (
	"net"

	"github.com/jaracil/websocket"
)

func Dial(url, _ string) (net.Conn, error) {
	return websocket.Dial(url)
}
