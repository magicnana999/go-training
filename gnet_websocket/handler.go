package gnet_websocket

type Handler interface {
	OnText(s *Session,text string)
	OnBinary(s *Session,bytes []byte)
	OnClose(s *Session)
	OnPing(s *Session)
	OnPong(s *Session)
}
