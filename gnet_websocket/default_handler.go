package gnet_websocket

import "go-training/logger"

type DefaultHandler struct {
}

func (d DefaultHandler) OnText(s *Session, text string) {
	logger.Logger.Debugf("OnText conn[%v] message[%v]", s.Remote, text)

	s.SendTextAsync("hello "+text)

}

func (d DefaultHandler) OnBinary(s *Session, bytes []byte) {
	logger.Logger.Debugf("OnBinary conn[%v] bytes.len[%v]", s.Remote, _len(bytes))
}

func (d DefaultHandler) OnClose(s *Session) {
	logger.Logger.Debugf("OnClose conn[%v]", s.Remote)
}

func (d DefaultHandler) OnPing(s *Session) {
	logger.Logger.Debugf("OnPing conn[%v]", s.Remote)
	s.SendPong(ServerVersion)
}

func (d DefaultHandler) OnPong(s *Session) {
	logger.Logger.Debugf("OnPong conn[%v]", s.Remote)
	s.SendPing()
}

func _len(bytes []byte) int {
	if bytes == nil {
		return -1
	} else {
		return len(bytes)
	}
}
