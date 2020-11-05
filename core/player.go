package core

import (
	"fmt"
	"game_server_silk/siface"
	"github.com/golang/protobuf/proto"
	"math/rand"
	"mmo_game_silk/pb"
	"sync"
)

//玩家实例对象
type Player struct {
	Pid int32  //玩家ID
	Conn siface.IConnection    //当前玩家的链接模块（用于和客户端的链接）
	X float32  //平面的x坐标
	Y float32  //高度
	Z float32  //平面y坐标
	V float32  //旋转的0-360角度
}

/*
	Player ID 生成器
 */
var PidGen int32 = 1 //用来生成玩家ID的计数器
var IdLock sync.Mutex //保护PidGen的Mutex 互斥锁


//创建一个玩家的方法
func NewPlayer(conn siface.IConnection) *Player {
	//生成一个玩家ID
	IdLock.Lock()
	id := PidGen
	PidGen ++
	IdLock.Unlock()

	//创建一个玩家对象
	p := &Player{
		Pid:  id,
		Conn: conn,
		X:    float32(160 + rand.Intn(10)), //随机在160坐标点，基于x轴若干偏移
		Y:    0,
		Z:    float32(140 + rand.Intn(20)), //随机在140坐标点，基于y轴若干偏移
		V:    0,
	}

	return p
}
/*
	提供一个发送给客户端消息的方法
	主要是将pb的protobuf数据序列化之后，再调用silk框架的SendMsg方法
 */
func (p *Player) SendMsg(msgId uint32, data proto.Message) {
	//将proto message结构体序列化 转换成二进制
	msg, err := proto.Marshal(data)
	if err != nil {
		fmt.Println("[转换proto消息失败]marshal msg err: ", err)
		return
	}
	//检查客户端链接是否在线链接的状态
	if(p.Conn == nil) {
		fmt.Println("[玩家客户端链接已关闭]connection in player is nil")
		return
	}

	//将二进制数据 通过silk框架的sendmsg将数据发送给客户端
	if err := p.Conn.SendMsg(msgId, msg); err != nil {
		fmt.Println("[玩家消息发送失败]Player SendMsg error!")
		return
	}

	return
}

/*
     以下为消息类型-----------------------------------------
 */

//告知客户端玩家Pid，同步已经生成的玩家ID给客户端
func (p *Player) SyncPid() {
	//组建MsgID：1 的proto数据
	proto_msg := &pb.SyncPid{
		Pid: p.Pid,
	}

	//将消息发送给客户端
	p.SendMsg(1, proto_msg)
}

//广播玩家自己的出生地点位置
func (p *Player) BroadCastStartPosition() {
	//组建MsgID：200 的proto数据
	proto_msg := &pb.BroadCast{
		Pid:  p.Pid,
		Tp:   2,   //Tp2 代表广播的位置坐标
		Data: &pb.BroadCast_P{
			P: &pb.Position{
				X: p.X,
				Y: p.Y,
				Z: p.Z,
				V: p.V,
			},
		},
	}

	//将消息发送给客户端
	p.SendMsg(200, proto_msg)
}




