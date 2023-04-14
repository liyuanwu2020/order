package service

import (
	"github.com/liyuanwu2020/msgo/engine"
	"time"
)

type Order struct {
	Id int    `json:"id"`
	No string `json:"no"`
}

func Route(ctx *engine.Context) {
	time.Sleep(time.Millisecond * 123)
	ctx.Logger.Info("执行顺序 main")
	panic("我错了")
	_ = ctx.Json(&Order{Id: 189485, No: "yjxnld"})
}
