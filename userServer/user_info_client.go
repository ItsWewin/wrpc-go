package userServer

import (
	"bufio"
	"encoding/json"
	"errors"
	"grpcTest/wrpc/wrpc"
	"log"
	"net"
)

type GetUserInfoClient interface {
	GetGoodsInfoByID(in *GetUserInfoRequest) (*UserInfo, error)
	GetGoodsInfoByName(name *GetUserInfoByNameRequest) (*[]UserInfo, error)
}

type GetUserInfoServerClient struct {
	conn       net.Conn
	ServerName string
}

func NewGetUserInfoClient(conn net.Conn) (*GetUserInfoServerClient, error) {
	if conn == nil {
		return nil, errors.New("net conn is nil")
	}
	return &GetUserInfoServerClient{
		conn:       conn,
		ServerName: "userServer/UserInfoServer",
	}, nil
}

func (client *GetUserInfoServerClient) GetGoodsInfoByID(in *GetUserInfoRequest) (*UserInfo, error) {

	result, err := client.sendAndReceiveMsg(in, "_GetUserInfoByID")

	bt, _ := json.Marshal(result)
	userinfo := UserInfo{}
	err = json.Unmarshal(bt, &userinfo)
	if err != nil {
		return nil, errors.New("Data from server unexpected. ")
	}

	return &userinfo, nil
}

func (client *GetUserInfoServerClient) GetGoodsInfoByName(in *GetUserInfoByNameRequest) ([]*UserInfo, error) {

	result, err := client.sendAndReceiveMsg(in, "_GetUserInfoByName")

	bt, _ := json.Marshal(result)
	var userinfo []*UserInfo
	err = json.Unmarshal(bt, &userinfo)
	if err != nil {
		return nil, errors.New("Data from server unexpected. ")
	}

	return userinfo, nil
}

func (client *GetUserInfoServerClient) sendAndReceiveMsg(in interface{}, methodName string) (interface{}, error) {
	if client == nil {
		return nil, errors.New("GetUserInfoServerClient is nil")
	}

	if client.conn == nil {
		return nil, errors.New("net conn is nil")
	}

	info := &wrpc.RequestInfoDesc{
		ServerName:  client.ServerName,
		MethodName:  methodName,
		RequestInfo: in,
	}

	bt, err := json.Marshal(info)
	if err != nil {
		return nil, errors.New("Request info marshal failed: " + err.Error())
	}
	_, err = client.conn.Write(append(bt, '\n'))
	if err != nil {
		log.Printf("client conn write failed: " + err.Error())
		return nil, err
	}

	respDesc := wrpc.ResponseInfoDesc{}
	for {
		data, err := bufio.NewReader(client.conn).ReadBytes('\n')
		if err != nil {
			log.Printf("Read data failed: %s", err)
			return nil, err
		}
		if len(data) == 0 {
			continue
		}

		err = json.Unmarshal(data, &respDesc)
		if err != nil {
			log.Printf("Read data from server failed: %s", err)
			return nil, err
		}

		if !respDesc.Succeed {
			return nil, errors.New("Read data from server failed" + respDesc.Msg)
		}

		return respDesc.Result, nil
	}
}
