package gnet_websocket

import (
	"context"
	"flag"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	"github.com/panjf2000/gnet/v2"
	"go-training/logger"
	"os"
	"runtime"
)

const (
	serverVersion = "SERVER_VERSION"
	OpWriteAsync  = 0x3
)

var ServerVersion string = "UNKNOWN"

type server struct {
	gnet.BuiltinEventEngine
	rootCtx context.Context
	handler Handler
	pool    *pool
}

func (s *server) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	logger.Logger.Debugf("open conn[%v]", c.RemoteAddr().String())
	session := &Session{
		upgraded: false,
		c:        c,
		Remote:   c.RemoteAddr().String(),
		pool:     s.pool,
	}
	c.SetContext(session)
	return
}

func (s *server) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	logger.Logger.Debugf("close conn[%v]", c.RemoteAddr().String())
	session := c.Context().(*Session)
	s.pool.addTask(runnable{handler: s.handler, session: session, op: ws.OpClose, data: nil})
	return
}

func (s *server) OnTraffic(c gnet.Conn) (action gnet.Action) {
	logger.Logger.Debugf("OnTraffic conn[%v]", c.RemoteAddr().String())

	session := c.Context().(*Session)

	if !session.upgraded {
		_, err := ws.Upgrade(c)
		if err != nil {
			logger.Logger.Errorf("OnTraffic upgrade conn[%v] [err=%v]", c.RemoteAddr().String(), err.Error())
			return gnet.Close
		} else {
			logger.Logger.Debugf("OnTraffic upgraded conn[%v]", c.RemoteAddr().String())
		}
		session.upgraded = true
		c.SetContext(session)
	} else {
		msg, op, err := wsutil.ReadClientData(c)
		logger.Logger.Debugf("OnTraffic read op[%v] conn[%v]", op, c.RemoteAddr().String())

		if err != nil {
			logger.Logger.Errorf("OnTraffic read conn[%v] [err=%v]", c.RemoteAddr().String(), err.Error())
			return gnet.Close
		}

		if op == ws.OpPing {
			s.handler.OnPing(session)
		} else if op == ws.OpPong {
			s.handler.OnPong(session)
		} else {
			s.pool.addTask(runnable{handler: s.handler, session: session, op: op, data: msg})
		}
	}

	return
}

func Start() {

	flag.StringVar(&ServerVersion, serverVersion, os.Getenv(serverVersion), serverVersion)
	flag.Parse()

	server := &server{
		rootCtx: context.Background(),
		handler: DefaultHandler{},
		pool:    newPool(runtime.NumCPU(), 65535),
	}

	server.pool.start()
	op1 := gnet.WithMulticore(true)
	op2 := gnet.WithNumEventLoop(50)

	if err := gnet.Run(server, "tcp://0.0.0.0:8080", op1, op2); err != nil {
		logger.Logger.Errorf("gnet_websocket start [err=%v]", err.Error())
	}
}
