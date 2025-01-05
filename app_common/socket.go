package appcommon

import "github.com/gorilla/websocket"

func makeMapSocket() {
	connectSocket = map[string][]*websocket.Conn{}
}

func GetSocket(uuid string) []*websocket.Conn {
	return connectSocket[uuid]
}

func CreateSocket(uuid string, connect *websocket.Conn) {
	if connectSocket[uuid] == nil {
		connectSocket[uuid] = make([]*websocket.Conn, 0)
	}
	connectSocket[uuid] = append(connectSocket[uuid], connect)
}
