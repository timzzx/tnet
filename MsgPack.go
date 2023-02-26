package tnet

import (
	"bytes"
	"encoding/binary"
	"io"
	"net"
)

// 消息解包
func Unpack(conn net.Conn) (int, string, error) {
	// 消息组成 routerId 4字节 | datalen 4字节| data
	// 获取路由id
	routerID := make([]byte, 4)
	_, err := io.ReadFull(conn, routerID)
	if err != nil {
		return 0, "", err
	}
	rid := int(binary.LittleEndian.Uint32(routerID))

	// 获取消息长度
	dataLen := make([]byte, 4)
	_, err = io.ReadFull(conn, dataLen)
	if err != nil {
		return 0, "", err
	}
	msgLen := int(binary.LittleEndian.Uint32(dataLen))

	// 获取消息
	data := make([]byte, msgLen)
	_, err = io.ReadFull(conn, data)
	if err != nil {
		return 0, "", err
	}
	msg := bytes.NewBuffer([]byte{})
	binary.Write(msg, binary.LittleEndian, data)

	return rid, string(data), nil
}

// 消息封包并发送
func PackSend(rid int, data string, conn net.Conn) ([]byte, error) {
	databuff := bytes.NewBuffer([]byte{})

	// 写msgID
	if err := binary.Write(databuff, binary.LittleEndian, uint32(rid)); err != nil {
		return nil, err
	}

	//写dataLen
	if err := binary.Write(databuff, binary.LittleEndian, uint32(len(data))); err != nil {
		return nil, err
	}

	//写data数据
	if err := binary.Write(databuff, binary.LittleEndian, []byte(data)); err != nil {
		return nil, err
	}

	return databuff.Bytes(), nil
}
