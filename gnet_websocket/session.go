package gnet_websocket

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/panjf2000/gnet/v2"
	"go-training/logger"
)

type Session struct {
	upgraded bool
	Remote   string
	c        gnet.Conn
	pool     *pool
}

// SendTextNow 在当前routine中直接写入
func (s *Session) SendTextNow(text string) {
	if err := wsutil.WriteServerMessage(s.c, ws.OpText, []byte(text)); err != nil {
		logger.Logger.Error("Session SentTextNow conn[%v] [err=%v]", s.Remote, err.Error())
	}
}

// SendTextAsync 在另一个routine中异步写入
func (s *Session) SendTextAsync(text string) {
	s.pool.addTask(runnable{handler: nil, session: s, op: OpWriteAsync, data: []byte(text)})
}

// SendPong 写一个Pong，当客户端Ping时，写入SERVER_VERSION
func (s *Session) SendPong(data string) {
	wsutil.WriteMessage(s.c, ws.StateServerSide, ws.OpPong, []byte(data))
}

// SendPing 写一个Ping，当客户端Pong时，写入一个Ping。
//我个人认为客户端ping，服务端pong的方式比较好。但是客户端如果上行pong，这里返回一个ping，以此提示客户端应该ping，而不应该pong
func (s *Session) SendPing() {
	wsutil.WriteMessage(s.c, ws.StateServerSide, ws.OpPing, nil)
}

// 内部方法，业务方不应该使用
func (s *Session) _doSendAsync(text []byte) {
	if err := wsutil.WriteServerMessage(s.c, ws.OpText, text); err != nil {
		logger.Logger.Error("Session SendTextAsync conn[%v] [err=%v]", s.Remote, err.Error())
	}
}
