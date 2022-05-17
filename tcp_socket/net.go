package tcp_socket

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"errors"
	"net"
	"time"
	"unsafe"
)

// HEADER_SIZE 消息头长度
const HEADER_SIZE = int(unsafe.Sizeof(MessageHeader{}))

// TcpClient tcp连接客户端
type TcpClient struct {
	conn net.Conn
	r    *bufio.Reader
}

// NewTcpClient 创建简单的连接客户端
func NewTcpClient(conn net.Conn) *TcpClient {
	return &TcpClient{conn: conn, r: bufio.NewReader(conn)}
}

// LocalAddr 本地通信地址
func (c *TcpClient) LocalAddr() net.Addr {
	return c.conn.LocalAddr()
}

// RemoteAddr 对方通信地址
func (c *TcpClient) RemoteAddr() net.Addr {
	return c.conn.RemoteAddr()
}

// Close 关闭连接
func (c *TcpClient) Close() error {
	return c.conn.Close()
}

// Write 向连接中写入消息
func (c *TcpClient) Write(msg Message) (int, error) {
	// 读取消息的长度
	msg.Header.Length = int32(len(msg.Body))
	var pkg = new(bytes.Buffer)

	//写入消息头
	err := binary.Write(pkg, binary.BigEndian, msg.Header)
	if err != nil {
		return 0, err
	}
	//写入消息体
	err = binary.Write(pkg, binary.BigEndian, msg.Body)
	if err != nil {
		return 0, err
	}
	if err = c.conn.SetWriteDeadline(time.Now().Add(time.Second * 5)); err != nil {
		return 0, err
	}
	nn, err := c.conn.Write(pkg.Bytes())
	if err != nil {
		return 0, err
	}
	return nn, nil
}

// Read 从连接中读取消息
func (c *TcpClient) Read() (Message, error) {
	// Peek 返回缓存的一个切片，该切片引用缓存中前 n 个字节的数据，
	// 该操作不会将数据读出，只是引用，引用的数据在下一次读取操作之
	// 前是有效的。如果切片长度小于 n，则返回一个错误信息说明原因。
	// 如果 n 大于缓存的总大小，则返回 ErrBufferFull。
	if err := c.conn.SetReadDeadline(time.Now().Add(time.Second * 5)); err != nil {
		return Message{}, err
	}
	lengthByte, err := c.r.Peek(HEADER_SIZE)
	if err != nil {
		return Message{}, err
	}
	//创建 Buffer缓冲器
	lengthBuff := bytes.NewBuffer(lengthByte)
	var header MessageHeader
	// 通过Read接口可以将buf中得内容填充到data参数表示的数据结构中
	err = binary.Read(lengthBuff, binary.BigEndian, &header)
	if err != nil {
		return Message{}, err
	}
	// Buffered 返回缓存中未读取的数据的长度
	if c.r.Buffered() < int(header.Length)+HEADER_SIZE {
		return Message{}, errors.New("not read all")
	}
	// 读取消息真正的内容
	pack := make([]byte, HEADER_SIZE+int(header.Length))
	// Read 从 b 中读出数据到 p 中，返回读出的字节数和遇到的错误。
	// 如果缓存不为空，则只能读出缓存中的数据，不会从底层 io.Reader
	// 中提取数据，如果缓存为空，则：
	// 1、len(p) >= 缓存大小，则跳过缓存，直接从底层 io.Reader 中读
	// 出到 p 中。
	// 2、len(p) < 缓存大小，则先将数据从底层 io.Reader 中读取到缓存
	// 中，再从缓存读取到 p 中。
	_, err = c.r.Read(pack)
	if err != nil {
		return Message{}, err
	}
	return Message{Header: header, Body: pack[HEADER_SIZE:]}, nil
}

// MessageHeader 消息头部 字段可自定义
type MessageHeader struct {
	MsgType int32  // 消息类型
	Length  int32  // 消息体长度
	Index   uint64 // 自定义参数
}

// Message 简单的消息体
type Message struct {
	Header MessageHeader
	Body   []byte
}
