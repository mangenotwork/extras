package model

import "sync"

var HashSize = 10

// web socket client hash table
var _wsClientTable = InitWsConnTable()
var _wsClientTableOnce sync.Once

// 提供调用
func WsClientTable() *WsHashTable{
	_wsClientTableOnce.Do(func() {
		_wsClientTable = InitWsConnTable()
	})
	return _wsClientTable
}

// tcp client hash table
var _tcpClientTable = InitTcpConnTable()
var _tcpClientTableOnce sync.Once

// 提供调用
func TcpClientTable() *TcpHashTable{
	_tcpClientTableOnce.Do(func() {
		_tcpClientTable = InitTcpConnTable()
	})
	return _tcpClientTable
}
