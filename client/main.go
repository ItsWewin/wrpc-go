package main

import (
	"grpcTest/wrpc/conf"
	"grpcTest/wrpc/userServer"
	"grpcTest/wrpc/wrpc/tls"
	"log"
	"sync"
)

func main() {
	conn, err := tls.Dial(conf.Network(), conf.Addr())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	getUserInfoClient, err := userServer.NewGetUserInfoClient(conn)
	if err != nil {
		log.Printf("NewGetUserInfoClient failed")
		return
	}

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go getUserInfo(wg, getUserInfoClient)

	wg.Wait()
}

func getUserInfo(wg *sync.WaitGroup, getUserInfoClient *userServer.GetUserInfoServerClient) {
	defer wg.Done()

	resp, err := getUserInfoClient.GetGoodsInfoByID(&userServer.GetUserInfoRequest{ID: 1})
	if err != nil {
		log.Printf("get user info by id failed: %s", err)
		return
	}
	log.Printf("user info: %v", resp)

	users, err := getUserInfoClient.GetGoodsInfoByName(&userServer.GetUserInfoByNameRequest{Name: "user1"})
	if err != nil {
		log.Printf("get user info by id failed: %s", err)
		return
	}

	log.Printf("user info: %v", users[0].Name)
}
