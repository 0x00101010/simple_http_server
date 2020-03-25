package connpool

import (
	"errors"
	"net"
	"sync"
)

var (
	// ErrFullCapacity predefined error for showing connection pool is
	// at maximum capacity
	ErrFullCapacity = errors.New("Connection pool is full right now")
)

// ConnPool provides connection pool management capability
type ConnPool struct {
	liveConns int
	maxConns  int
	m         *sync.RWMutex
}

// InitConnPool initializes connection pool
func InitConnPool() ConnPool {
	return ConnPool{
		liveConns: 0,
		maxConns:  10,
		m:         &sync.RWMutex{},
	}
}

// TryAcceptNewConn tries to establish new connection when liveConns < maxConns
func (cp ConnPool) TryAcceptNewConn(listener net.Listener) (net.Conn, error) {
	cp.m.RLock()
	if cp.liveConns >= cp.maxConns {
		defer cp.m.Unlock()
		return nil, ErrFullCapacity
	}
	cp.m.RUnlock()

	// lock for acquireing new connections
	cp.m.Lock()
	conn, err := listener.Accept()
	if err != nil {
		defer cp.m.Unlock()
		return nil, err
	}
	cp.liveConns++
	cp.m.Unlock()

	return conn, nil
}
