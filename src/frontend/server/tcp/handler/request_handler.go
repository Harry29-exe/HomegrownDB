package handler

import "net"

type requestHandler struct {
	conn           net.Conn
	requestLenBuff []byte
}

//func NewRequestHandler() *requestHandler {
//
//}
//
//func (rh *requestHandler) Handle(conn net.Conn) {
//
//	n, err := rh.conn.Read(rh.requestLenBuff)
//	if n != 4 {
//
//	}
//}

func (rh requestHandler) Read() {

}
