package userServer

import (
	"encoding/json"
	"errors"
	"fmt"
	"grpcTest/wrpc/wrpc"
	"log"
)

type UserInfoServerInterface interface {
	GetUserInfoByID(req *GetUserInfoRequest) (*UserInfo, error)
	GetUserInfoByName(req *GetUserInfoByNameRequest) ([]*UserInfo, error)
}

type UserInfoServer struct {

}

func RegisterSayHiServiceServer(s *wrpc.RpcServer, srv UserInfoServerInterface) {
	s.Register(&_GetUserInfoByID_serviceDesc, srv)
}

func (u *UserInfoServer) GetUserInfoByID(req *GetUserInfoRequest) (*UserInfo, error) {
	if req == nil {
		log.Printf("Req is nil")
		return nil, errors.New("Get user info failed. ")
	}

	if req.ID <= 0 {
		return nil, errors.New("req.ID is < 0")
	}

	user := GetUserInfoByID(req.ID)
	if user == nil {
		return nil, errors.New(fmt.Sprintf("No user with id: %d", req.ID))
	}

	return &UserInfo{
		ID:   user.ID,
		Name: user.Name,
	}, nil
}

func (u *UserInfoServer) GetUserInfoByName(req *GetUserInfoByNameRequest) ([]*UserInfo, error) {
	if req == nil {
		log.Printf("Req is nil")
		return nil, errors.New("Request info is nil. ")
	}

	if len(req.Name) == 0 {
		return nil, errors.New("Request info name is invalid. ")
	}

	users := GetUserInfoByName(req.Name)
	if len(users) == 0 {
		return nil, errors.New(fmt.Sprintf("No user with name: %s", req.Name))
	}

	return users, nil
}

var _GetUserInfoByID_serviceDesc = wrpc.ServiceDesc{
	ServiceName: "userServer/UserInfoServer",
	Methods:     []wrpc.MethodDesc{
		{
			MethodName: "_GetUserInfoByID",
			Handler:  getUserInfoByIDHandler,
		},
		{
			MethodName: "_GetUserInfoByName",
			Handler:  getUserInfoByNameHandler,
		},
	},
}

func getUserInfoByIDHandler(req interface{}, sl interface{}) (interface{}, error) {
	bt, err := json.Marshal(req)
	if err != nil {
		return nil, errors.New("Request info invalid. ")
	}

	request  := GetUserInfoRequest{}
	err = json.Unmarshal(bt, &request)
	if err != nil {
		return nil, errors.New("Request info invalid. ")
	}

	server, ok := sl.(UserInfoServerInterface)
	if !ok {
		return nil, errors.New("server is not implement UserInfoServer")
	}

	return server.GetUserInfoByID(&request)
}

func getUserInfoByNameHandler(req interface{}, sl interface{}) (interface{}, error) {
	bt, err := json.Marshal(req)
	if err != nil {
		return nil, errors.New("Request info invalid. ")
	}

	request := GetUserInfoByNameRequest{}
	err = json.Unmarshal(bt, &request)
	if err != nil {
		return nil, errors.New("Request info invalid. ")
	}

	server, ok := sl.(UserInfoServerInterface)
	if !ok {
		return nil, errors.New("server is not implement of UserInfoServer")
	}

	return server.GetUserInfoByName(&request)
}

