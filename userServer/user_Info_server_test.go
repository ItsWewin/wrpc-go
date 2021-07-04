package userServer

import (
	"grpcTest/wrpc/wrpc"
	"testing"
)

func TestNewAndRegister(t *testing.T) {
 	wRpcServer := wrpcServer.New()

 	userInfo := new(UserInfoServer)
	RegisterSayHiServiceServer(wRpcServer, userInfo)

	for sn, si := range wRpcServer.ServerInfoMap {
		t.Logf("serverName: %s\n", sn)

		for mn, md := range si.Methods {
			t.Logf("method name: %s", mn)

			req := GetUserInfoRequest{ID: 1}
			resp, err := md.Handler(req, si.ServiceImpl)
			if err != nil {
				t.Fatalf("md.handler failed: %s", err)
			}

			t.Logf("resp: %s", resp)
		}
	}

 	server, ok := wRpcServer.ServerInfoMap["userServer/UserInfoServer"]
 	if !ok {
 		t.Fatalf("no server")
	}

	method, ok := server.Methods["GetUserInfoByIDHandler"]
	if !ok {
		t.Fatalf("no server")
	}

	req := GetUserInfoRequest{ID: 1}

	response, err := method.Handler(req, server.ServiceImpl)
	if err != nil {
		t.Fatalf("some error")
	}

	t.Logf("response: %v",response)
}
