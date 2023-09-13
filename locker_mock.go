package websocket_conn_pool

import (
	"sync"
	"testing"
)

const (
	name =
)

// MockMutex mock mutex
// lock state: locked, unlock, starved(in queue)
type MockMutex struct {
	state int32 // 锁状态, 4个字节的32位整型
}

type MockLocker interface {
	mockLock()
	mockUnlock()
	tryLocker()
}

func (m MockMutex) mockLock() {
	// 判断locker是否锁定状态
	if m.state ==
	// 如果空闲状态，尝试加锁
}

func (MockMutex) tryLocker() {

}

func (MockMutex) mockUnlock() {

}

func TestMutex(t *testing.T) {
	var local sync.Mutex

}
