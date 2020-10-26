package core

import "fmt"

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
	s := fmt.Sprintf("AOIManager:\n minX:%d, maxX:%d, cntsX:%d, minY:%d, maxY:%d, cntsY:%d\n Grids in AOIManager:\n")

	//打印全部格子信息
	for _, grid := range m.grids {
		s += fmt.Sprintln(grid)
	}
	return s
}