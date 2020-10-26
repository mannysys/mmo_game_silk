package core

import (
	"fmt"
	"sync"
)

/*
	一个AOI地图中的格子类型
 */
type Grid struct {
	//格子ID
	GID int
	//格子的左边 边界坐标
	MinX int
	//格子的右边 边界坐标
	MaxX int
	//格子的上边 边界坐标
	MinY int
	//格子的下边 边界坐标
	MaxY int
	//当前格子内玩家或者物体成员的ID集合
	playerIDs map[int]bool
	//保护当前集合的锁
	pIDLock sync.RWMutex

}

//初始化当前格子的方法
func NewGrid(gID, minX, maxX, minY, maxY int) *Grid {
	return &Grid{
		GID:       gID,
		MinX:      minX,
		MaxX:      maxX,
		MinY:      minY,
		MaxY:      maxY,
		playerIDs: make(map[int]bool),
	}
}

//给格子添加一个玩家
func (g *Grid) Add(playerID int) {
	//加上写锁
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()
	//将玩家ID加入当前格子的集合中
	g.playerIDs[playerID] = true
}
//从格子中删除一个玩家
func (g *Grid) Remove(playerID int)  {
	//加上写锁
	g.pIDLock.Lock()
	defer g.pIDLock.Unlock()

	delete(g.playerIDs, playerID)
}
//得到当前格子中所有玩家ID
func (g *Grid) GetPlayerIDs() (playerIDs []int) {
	//加上读锁
	g.pIDLock.RLock()
	defer g.pIDLock.RUnlock()

	//遍历map集合格子内玩家ID，加入到playerIDs切片中
	for k, _ := range g.playerIDs {
		playerIDs = append(playerIDs, k)
	}
	return
}
//调试使用-打印出格子的基本信息
func (g *Grid) String() string {
	return fmt.Sprintf("Grid id: %d, minX:%d, maxX:%d, minY:%d, maxY:%d, playerIDs:%v",
		g.GID, g.MinX, g.MaxX, g.MinY, g.MaxY, g.playerIDs)
}


