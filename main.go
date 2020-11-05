package main

import (
	"fmt"
	"game_server_silk/siface"
	"game_server_silk/snet"
	"mmo_game_silk/core"
)


//当前客户端建立连接之后的hook钩子函数
func OnConnectionAdd(conn siface.IConnection) {
	//创建一个player实例对象
	player := core.NewPlayer(conn)

	//给客户端发送MsgID：1 的消息，服务器同步当前player的ID 给客户端
	player.SyncPid()

	//给客户端发送MsgID：200 的消息，服务器同步当前player的初始位置坐标给客户端
	player.BroadCastStartPosition()

	fmt.Println("=====> Player pid = ", player.Pid, " is arrived <=====")
}

func main() {
	//创建silk server句柄
	s := snet.NewServer("MMO Game Silk Server")


	//链接创建和销毁的Hook钩子函数
	s.SetOnConnStart(OnConnectionAdd)

	//注册一些路由业务


	//启动服务
	s.Serve()
}