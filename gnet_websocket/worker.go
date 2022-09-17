package gnet_websocket

import (
	"github.com/gobwas/ws"
	"sync"
)

//异步处理所有任务，包括WS的Text、Ping、Pong和下行的写入任务。
//0x3 在WS中保留，这里临时借用一下
type pool struct {
	workers   int
	maxTasks  int
	taskQueue chan runnable

	mu     sync.Mutex
	closed bool
	done   chan struct{}
}

type runnable struct {
	handler Handler
	session *Session
	op      ws.OpCode
	data    []byte
}

func newPool(w int, t int) *pool {
	return &pool{
		workers:   w,
		maxTasks:  t,
		taskQueue: make(chan runnable, t),
		done:      make(chan struct{}),
	}
}

func (p *pool) Close() {
	p.mu.Lock()
	p.closed = true
	close(p.done)
	close(p.taskQueue)
	p.mu.Unlock()
}

func (p *pool) addTask(run runnable) {
	p.mu.Lock()
	if p.closed {
		p.mu.Unlock()
		return
	}

	p.taskQueue <- run
	p.mu.Unlock()
}

func (p *pool) start() {
	for i := 0; i < p.workers; i++ {
		go p.startWorker()
	}
}

func (p *pool) startWorker() {
	for {
		select {
		case <-p.done:
			return
		case pp := <-p.taskQueue:
			switch pp.op {
			case ws.OpText:
				pp.handler.OnText(pp.session, string(pp.data))
				break
			case ws.OpBinary:
				pp.handler.OnBinary(pp.session, pp.data)
				break
			case ws.OpClose:
				pp.handler.OnClose(pp.session)
				break
			case OpWriteAsync:
				pp.session._doSendAsync(pp.data)
				break
			}

		}
	}
}
