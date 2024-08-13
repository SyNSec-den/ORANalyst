package utils

import (
	"encoding/gob"
	"net"
)

const (
	SERVER_HOST        = "localhost"
	MHO_SERVER_PORT    = 19960
	RC_SERVER_PORT     = 19961
	RC_PRE_SERVER_PORT = 19962
	KPM_SERVER_PORT    = 19963
	BUFFERSIZE         = 1 << 20
)

func SocketRecv(conn net.Conn) ([]byte, error) {
	buf := make([]byte, BUFFERSIZE)
	n, err := conn.Read(buf)
	if err != nil {
		return nil, err
	}
	return buf[:n], nil
}

// GobMarshal encodes obj into a gob and transmit over conn
func GobMarshal(conn net.Conn, obj interface{}) error {
	enc := gob.NewEncoder(conn)
	return enc.Encode(obj)
}

// GobUnmarshal recieves from conn and decodes gob into obj
func GobUnmarshal(conn net.Conn, obj interface{}) error {
	dec := gob.NewDecoder(conn)
	return dec.Decode(obj)
}
