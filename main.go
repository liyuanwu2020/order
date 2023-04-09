package main

import (
	"context"
	"fmt"
	"github.com/liyuanwu2020/goods/model"
	"github.com/liyuanwu2020/micro.service.pb/go/user"
	"github.com/liyuanwu2020/msgo"
	"github.com/liyuanwu2020/msgo/register"
	"github.com/liyuanwu2020/msgo/rpc"
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

		c, _ := rpc.NewGrpcClient(rpc.DefaultGrpcClientConfig(":9111"))
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {

			}
		}(c.Conn)
		c.Conn.Connect()
		rpcClient := user.NewUserServiceClient(c.Conn)
		rsp, err := rpcClient.Search(context.TODO(), &user.UserRequest{})
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, rsp)
	})

	group.Get("/findGrpc2", func(ctx *msgo.Context) {
		log.Println("grpc request2")
		nacosClient, nacosErr := register.CreateNacosClient()
		if nacosErr != nil {
			panic(nacosErr)
		}
		ip, port, nacosErr := register.GetInstance(nacosClient, "user")
		if nacosErr != nil {
			panic(nacosErr)
		}
		addr := fmt.Sprintf("%s:%d", ip, port)
		conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			panic(err)
		}
		defer func(conn *grpc.ClientConn) {
			err := conn.Close()
			if err != nil {

			}
		}(conn)
		rpcClient := user.NewUserServiceClient(conn)
		rsp, err := rpcClient.Search(context.TODO(), &user.UserRequest{})
		if err != nil {
			panic(err)
		}

		ctx.JSON(http.StatusOK, rsp)
	})

	engine.Run(":9003")
}
