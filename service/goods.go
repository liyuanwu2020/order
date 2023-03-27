package service

import "github.com/liyuanwu2020/msgo/rpc"

type GoodsService struct {
	Find func(args map[string]any) ([]byte, error) `msrpc:"GET,/goods/find"`
}

func (s *GoodsService) Env() rpc.HttpConfig {

	return rpc.HttpConfig{
		Host: "localhost",
		Port: 9002,
	}
}
