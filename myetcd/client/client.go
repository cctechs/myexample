package main

import (
	"context"
	"encoding/json"
	"github.com/etcd-io/client"
	"log"
	"os"
	"strconv"
	"time"
)

// register service

type ServerInfo struct {
	Id int32 `json:"id"`
	IP string `json:"ip"`
	Port int32 `json:"port"`
}
type Service struct {
	processid int
	info ServerInfo
	KeysAPI client.KeysAPI
}

func (s *Service) HeartBeat(){
	api := s.KeysAPI
	for  {
		key := "lc_server/p_" + strconv.Itoa(s.processid)
		value, _ := json.Marshal(s.info)
		_, err := api.Set(context.Background(), key, value, &client.SetOptions{
			PrevValue:        "",
			PrevIndex:        0,
			PrevExist:        "",
			TTL:              time.Second*20,
			Refresh:          false,
			Dir:              false,
			NoValueOnSuccess: false,
		})
		if err != nil{
			log.Println("error update workInfo:", err)
		}
		time.Sleep(time.Second*10)
	}
}

func RegisterService(endpoints []string){
	cfg := client.Config{
		Endpoints:               endpoints,
		Transport:               client.DefaultTransport,
		CheckRedirect:           nil,
		Username:                "",
		Password:                "",
		HeaderTimeoutPerRequest: time.Second,
		SelectionMode:           0,
	}

	etcdClient, err := client.New(cfg)
	if err != nil{
		log.Fatalf("Error can not connect to etcd:%v", err)
	}

	s := &Service{
		processid: os.Getpid(),
		info:      ServerInfo{Id:1024, IP:"127.0.0.1", Port:100},
		KeysAPI:   client.NewKeysAPI(etcdClient),
	}
	go s.HeartBeat()
}

func main(){
	RegisterService([]string{"127.0.0.1"})
}
