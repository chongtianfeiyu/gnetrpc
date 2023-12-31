package server

import (
	"github.com/cat3306/gnetrpc/protocol"
	"github.com/cat3306/gnetrpc/rpclog"
	"github.com/panjf2000/gnet/v2"
	"sync"
	"time"
)

type serverOption struct {
	printMethod bool
	gnetOptions gnet.Options
}
type Server struct {
	gnet.BuiltinEventEngine
	eng          gnet.Engine
	serviceMapMu sync.RWMutex
	serviceMap   map[string]*service
	option       serverOption
}

func (s *Server) OnBoot(engine gnet.Engine) (action gnet.Action) {
	return
}

func (s *Server) OnShutdown(engine gnet.Engine) {
}

func (s *Server) OnOpen(c gnet.Conn) (out []byte, action gnet.Action) {
	return
}

func (s *Server) OnClose(c gnet.Conn, err error) (action gnet.Action) {
	return
}

func (s *Server) OnTraffic(c gnet.Conn) (action gnet.Action) {

	ctx, err := protocol.Decode(c)
	if err != nil {
		rpclog.Warnf("OnTraffic err:%s", err.Error())
		return gnet.None
	}
	rpclog.Infof("%+v", ctx)
	return
}

func (s *Server) OnTick() (delay time.Duration, action gnet.Action) {
	return
}

func (s *Server) Run(netWork string, addr string) error {
	return gnet.Run(s, netWork+"://"+addr, gnet.WithOptions(s.option.gnetOptions))
}
func NewServer(options ...OptionFn) *Server {
	s := &Server{
		//Plugins:    &pluginContainer{},
		//options:    make(map[string]interface{}),
		//activeConn: make(map[net.Conn]struct{}),
		//doneChan:   make(chan struct{}),
		serviceMap: make(map[string]*service),
		//router:     make(map[string]Handler),
		//AsyncWrite: false, // 除非你想做进一步的优化测试，否则建议你设置为false
	}

	for _, op := range options {
		op(s)
	}

	//if s.options["TCPKeepAlivePeriod"] == nil {
	//	s.options["TCPKeepAlivePeriod"] = 3 * time.Minute
	//}
	return s
}
