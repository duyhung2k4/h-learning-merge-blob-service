package appcommon

import "github.com/gorilla/websocket"

func makeMapSocket() {
	connectSocket = map[string]*websocket.Conn{}
}

func GetSocket(uuid string) *websocket.Conn {
	return connectSocket[uuid]
}

func CreateSocket(uuid string, connect *websocket.Conn) {
	connectSocket[uuid] = connect
}
