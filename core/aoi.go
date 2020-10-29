package core

import (
	"fmt"
)

/*
	AOI区域管理模块
 */
type AOIManager struct {
	//区域的左边界坐标
	MinX int
	//区域的右边界坐标
	MaxX int
	//X轴方向格子的数量
	CntsX int

	//区域的上边界坐标
	MinY int
	//区域的下边界坐标
	MaxY int
	//Y轴方向格子的数量
	CntsY int

	//当前区域中有哪些格子map-key=格子的ID，value=格子对象
	grids map[int] *Grid
}

//初始化一个AOI区域管理模块
func NewAOIManager(minX, maxX, cntsX, minY, maxY, cntsY int) *AOIManager {
	aoiMgr := &AOIManager{
		MinX:  minX,
		MaxX:  maxX,
		CntsX: cntsX,
		MinY:  minY,
		MaxY:  maxY,
		CntsY: cntsY,
		grids: make(map[int] *Grid),
	}

	//给AOI初始化区域内的所有格子进行编号 和 初始化
	for  y := 0; y < cntsY; y++ {
		for x := 0; x < cntsX; x++ {
			//计算格子ID  根据 x，y编号
			//计算格子编号公式 id = idy * cntsX + idx
			gid := y * cntsX + x

			//初始化gid格子
			aoiMgr.grids[gid] = NewGrid(gid,
				aoiMgr.MinX + x * aoiMgr.gridWidth(),
				aoiMgr.MinX + (x+1) * aoiMgr.gridWidth(),
				aoiMgr.MinY + y * aoiMgr.gridLength(),
				aoiMgr.MinY + (y+1) * aoiMgr.gridLength())
		}
	}

	return aoiMgr
}
//得到每个格子在 X轴方向的宽度
func (m *AOIManager) gridWidth() int {
	return (m.MaxX - m.MinX) / m.CntsX
}
//得到每个格子在 Y轴方向的长度
func (m *AOIManager) gridLength() int {
	return (m.MaxY - m.MinY) / m.CntsY
}

//打印格子信息
func (m *AOIManager) String() string {
	//打印AOIManager信息
	s := fmt.Sprintf("AOIManager:\n minX:%d, maxX:%d, cntsX:%d, minY:%d, maxY:%d, cntsY:%d\n Grids in AOIManager:\n",
		m.MinX, m.MaxX, m.CntsX, m.MinY, m.MaxY, m.CntsY)

	//打印全部格子信息
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}
	return s
}


/*
	根据格子GID 得到周边九宫格格子的ID集合
 */
func (m *AOIManager) GetSurroundGridsByGid(gID int) (grids []*Grid) {
	//判断gID是否在 AOIManager中
	if _, ok := m.grids[gID]; !ok {
		return
	}

	//将当前gid本身 加入九宫格切片中
	grids = append(grids, m.grids[gID]) //8

	//需要判断gID的左边是否有格子？右边是否有格子
	//需要通过gID得到当前格子x轴的编号 -- idx = id % nx
	idx := gID % m.CntsX //3

	//判断idx编号是否左边还有格子
	if idx > 0 {
		grids = append(grids, m.grids[gID - 1]) //7
	}
	//判断idx编号是否右边还有格子
	if idx < m.CntsX-1 {
		grids = append(grids, m.grids[gID + 1]) //9
	}

	//将x轴当前的格子都取出，进行遍历，再分别得到每个格子上下是否还有格子
	//得到当前x轴格子的ID集合
	gidsX := make([]int, 0, len(grids))
	for _, v := range grids {
		gidsX = append(gidsX, v.GID)
	}

	//遍历gidsX 集合中每个格子的gid
	for _, v := range gidsX {
		//得到当前格子id的y轴的编号 idy = id / ny
		idy := v / m.CntsY
		//gid 上边是否还有格子
		if idy > 0 {
			grids = append(grids, m.grids[v - m.CntsX])
		}
		//gid 下边是否还有格子
		if idy < m.CntsY - 1 {
			grids = append(grids, m.grids[v + m.CntsX])
		}
	}

	return
}


//通过x、y横纵轴坐标得到当前的GID格子编号
func (m *AOIManager) GetGidByPos(x, y float32) int {
	idx := (int(x) - m.MinX) / m.gridWidth()
	idy := (int(y) - m.MinY) / m.gridLength()

	return idy * m.CntsX + idx
}

/*
	通过x、y横纵坐标得到周边九宫格内全部的PlayerIDs
 */
func (m *AOIManager) GetPidsByPos(x, y float32) (playerIDs []int) {
	//得到当前玩家的GID格子id
	//gid := y * m.CntsX + x
	gID := m.GetGidByPos(x, y)

	//通过GID得到周边九宫格信息
	grids := m.GetSurroundGridsByGid(gID)

	//将九宫格的信息里全部的Player的id 累加到 playerIDs
	for _, grid := range grids {
		playerIDs = append(playerIDs, grid.GetPlayerIDs()...)
		fmt.Println("===> grid ID: %d, pids: %v ===", grid.GID, grid.GetPlayerIDs())
	}

	return
}

//-------------------网格的增删改查方法-------------------------------

//添加一个PlayerID 到一个格子中
func (m *AOIManager) AddPidToGrid(pID, gID int) {
	m.grids[gID].Add(pID)
}
//移除一个格子中的PlayerID
func (m *AOIManager) RemovePidFromGrid(pID, gID int) {
	m.grids[gID].Remove(pID)
}
//通过GID获取全部的PlayerID
func (m *AOIManager) GetPidsByGid(gID int) (playerIDs []int) {
	playerIDs = m.grids[gID].GetPlayerIDs()
	return
}
//通过坐标将Player添加到一个格子中
func (m *AOIManager) AddToGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Add(pID)
}
//通过坐标把一个Player从一个格子中删除
func (m *AOIManager) RemoveFromGridByPos(pID int, x, y float32) {
	gID := m.GetGidByPos(x, y)
	grid := m.grids[gID]
	grid.Remove(pID)
}



