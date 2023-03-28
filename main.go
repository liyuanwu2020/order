package main

import (
	"context"
	"github.com/liyuanwu2020/goods/model"
	"github.com/liyuanwu2020/msgo"
	"github.com/liyuanwu2020/msgo/rpc"
	"github.com/liyuanwu2020/order/api"
	"github.com/liyuanwu2020/order/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
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
		//改造目标 rpc 方式
		body, err := client.Do("goods", "Find").(*service.GoodsService).Find(params)
		if err != nil {
			log.Println(err)
		}
		ctx.JSON(http.StatusOK, &model.Result{Code: 1000, Msg: "success from rpc", Data: string(body)})

	})

	group.Get("/findGrpc", func(ctx *msgo.Context) {
		conn, err := grpc.Dial("localhost:9111", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)
		rpcClient := api.NewGoodsApiClient(conn)
		rsp, err := rpcClient.Find(context.TODO(), &api.GoodsRequest{})
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, rsp)
	})

	engine.Run(":9003")
}
