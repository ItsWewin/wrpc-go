package wrpc

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net"
)

type RpcServer struct {
	ServerInfoMap map[string]*ServerInfo
}

type ServerInfo struct {
	Name        string
	ServiceImpl interface{}
	Methods     map[string]*MethodDesc
}

type MethodDesc struct {
	MethodName string
	Handler    FuncHandler
}

type FuncHandler func(req interface{}, sl interface{}) (interface{}, error)

type ResponseDesc struct {
	MethodName string
}

type ServiceDesc struct {
	ServiceName string
	Methods     []MethodDesc
}

type RequestInfoDesc struct {
	ServerName  string      `json:"server_name"`
	MethodName  string      `json:"method_name"`
	RequestInfo interface{} `json:"request_info"`
}

type ResponseInfoDesc struct {
	ServerName string      `json:"server_name"`
	MethodName string      `json:"method_name"`
	Succeed    bool        `json:"succeed"`
	Msg        string      `json:"msg"`
	Result     interface{} `json:"result"`
}

func New() *RpcServer {
	return &RpcServer{ServerInfoMap: make(map[string]*ServerInfo)}
}

func (s *RpcServer) Register(sd *ServiceDesc, ss interface{}) {
	_, ok := s.ServerInfoMap[sd.ServiceName]
	if ok {
		panic(fmt.Sprintf("servers: %s is already register", sd.ServiceName))
	}

	info := &ServerInfo{
		ServiceImpl: ss,
		Methods:     make(map[string]*MethodDesc),
	}

	for i := range sd.Methods {
		d := &sd.Methods[i]
		info.Methods[d.MethodName] = d
	}

	s.ServerInfoMap[sd.ServiceName] = info
}

func (s *RpcServer) HandlerServer(data []byte) (interface{}, error) {
	log.Printf("HandlerServer info: %s", string(data))

	req := RequestInfoDesc{}
	err := json.Unmarshal(data, &req)
	if err != nil {
		return nil, errors.New("Handler server failed. json.Unmarshal failed. ")
	}

	return s.handlerServer(req.ServerName, req.MethodName, req.RequestInfo)
}

func (s *RpcServer) handlerServer(serverName string, methodName string, request interface{}) (interface{}, error) {
	if request == nil {
		return nil, errors.New("Request info is nil. ")
	}

	server, ok := s.ServerInfoMap[serverName]
	if !ok {
		return nil, errors.New(fmt.Sprintf("server: %s not existed", serverName))
	}

	return server.Methods[methodName].Handler(request, server.ServiceImpl)
}

func (s *RpcServer) Server(lis net.Listener) error {
	for {
		conn, err := lis.Accept()
		if err != nil {
			return errors.New("Server accept failed. err: " + err.Error())
		}

		go s.handlerConn(conn)
	}
}

func (s *RpcServer) handlerConn(conn net.Conn) {
	defer conn.Close()

	for {
		netData, err := bufio.NewReader(conn).ReadBytes('\n')
		if err == io.EOF {
			log.Printf("Read data finished")
			return
		}

		if err != nil {
			log.Printf("Read data failed: %s", err)
			return
		}

		err = s.handlerData(conn, netData)
		if err != nil {
			log.Printf("Read data failed")
		}
	}
}

func (s *RpcServer) handlerData(conn net.Conn, data []byte) error {
	var resp = ResponseInfoDesc{Succeed: true}

	result, err := s.HandlerServer(data)
	if err != nil {
		log.Printf("Handler server failed: %s\n", err)

		resp.Succeed = false
		resp.Msg = fmt.Sprintf("Handler exce failed: %s", err)
	}
	resp.Result = result

	data, _ = json.Marshal(resp)
	_, err = conn.Write(append(data, '\n'))
	if err != nil {
		log.Printf("conn write failed: %s", err)
	}

	return nil
}
