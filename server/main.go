package main

import (
	"grpcTest/wrpc/conf"
	"grpcTest/wrpc/userServer"
	"grpcTest/wrpc/wrpc"
	"grpcTest/wrpc/wrpc/tls"
)

func main() {
	wRpcServer := wrpc.New()
	userInfo := new(userServer.UserInfoServer)
	userServer.RegisterSayHiServiceServer(wRpcServer, userInfo)

	lis, err := tls.Listen(conf.Network(), conf.Addr())
	if err != nil {
		panic(err)
	}

	err = wRpcServer.Server(lis)
	if err != nil {
		panic("Server on lis failed: " + err.Error())
		return
	}
}
