package appcommon

import "github.com/gorilla/websocket"

var (
	chanListenAddProcessStream (chan string)
	processStream              map[string](chan string)
	connectSocket              map[string]*websocket.Conn
)
