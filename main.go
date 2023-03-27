package main

import (
	"github.com/liyuanwu2020/msgo"
	"github.com/liyuanwu2020/msgo/rpc"
	"github.com/liyuanwu2020/order/service"
	"log"
)

func main() {
	engine := msgo.Default()
	group := engine.Group("order")
	client := rpc.NewHttpClient()
	client.RegisterHttpService("goods", &service.GoodsService{})
	group.Get("/find", func(ctx *msgo.Context) {
		params := make(map[string]any)
		params["id"] = 100
		params["name"] = "xiaotian"
		//通过商品中心 查询商品信息
		//http的调用的方式
		//body, err := client.Get("http://127.0.0.1:9002/goods/find")
		//改造目标
		body, err := client.Do("goods", "Find").(*service.GoodsService).Find(params)
		log.Println(err)
		log.Println(string(body))

	})
	engine.Run(":9003")
}
