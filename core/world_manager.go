package core

import "sync"

/*
	当前游戏的世界总管理模块
 */
type WorldManager struct {
	//AOIManager 当前世界地图的AOI的管理模块
	AoiMgr *AOIManager
	//当前全部在线的Players集合
	Players map[int32] *Player
	//保护Players集合的读写锁
	pLock sync.RWMutex
}

//提供一个对外的世界管理模块的句柄（全局）
var WorldMgrObj *WorldManager

//初始化方法（init方法是在其它模块中导入该包的时候自动调用一次）
func init()  {
	WorldMgrObj = &WorldManager{
		//创建世界AOI地图规划
		AoiMgr:  NewAOIManager(AOI_MIN_X, AOI_MAX_X, AOI_CNTS_X, AOI_MIN_Y, AOI_MAX_Y, AOI_CNTS_Y),
		//初始化Player map集合
		Players: make(map[int32] *Player),
	}
}

//添加一个玩家
func (wm *WorldManager) AddPlayer(player *Player) {
	//保护共享资源map，加写锁
	wm.pLock.Lock()
	//将玩家加入集合中
	wm.Players[player.Pid] = player
	wm.pLock.Unlock()

	//将player添加到 AOIManager区域管理网格中的 一个格子内
	wm.AoiMgr.AddToGridByPos(int(player.Pid), player.X, player.Z)
}

//删除一个玩家
func (wm *WorldManager) RemovePlayerByPid(pid int32) {
	//得到当前的玩家
	player := wm.Players[pid]
	//将玩家从AOIManager区域管理中删除
	wm.AoiMgr.RemoveFromGridByPos(int(pid), player.X, player.Z)

	//加上写锁
	wm.pLock.Lock()
	//将玩家从世界管理玩家集合中删除
	delete(wm.Players, pid)
	wm.pLock.Unlock()
}

//通过玩家ID查询Player对象
func (wm *WorldManager) GetPlayerByPid(pid int32) *Player {
	//加读锁
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	return wm.Players[pid]
}

//获取全部的在线玩家
func (wm *WorldManager) GetAllPlayers() []*Player {
	//加读锁
	wm.pLock.RLock()
	defer wm.pLock.RUnlock()

	//初始化一个切片
	players := make([]*Player, 0)

	//添加到切片中
	for _, p := range wm.Players {
		players = append(players, p)
	}

	return players
}
