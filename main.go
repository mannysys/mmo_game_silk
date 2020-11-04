package main

import "game_server_silk/snet"

func main() {
	//创建silk server句柄
	s := snet.NewServer("MMO Game Silk Server")


	//链接创建和销毁的Hook钩子函数

	//注册一些路由业务


	//启动服务
	s.Serve()
}