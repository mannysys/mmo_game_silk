package core

import (
	"fmt"
	"testing"
)

//给出x和y 的坐标轴，x、y轴方向格子数量，得出整个网格
func TestNewAOIManager(t *testing.T) {
	//初始化AOIManager
	aoiMgr := NewAOIManager(0,250,5,0,250,5)

	//打印AOIManager
	fmt.Println(aoiMgr)
}

//根据格子的编号gid 得出格子周围九宫格的格子编号gid
func TestAOIManagerSuroundGridsByGid(t *testing.T) {
	//初始化AOIManager
	aoiMgr := NewAOIManager(0,250,5,0,250,5)

	for gid, _ := range aoiMgr.grids {
		//得到当前gid的周边九宫格信息
		grids := aoiMgr.GetSurroundGridsByGid(gid)
		fmt.Println("gid : ", gid, "grids len = ", len(grids))
		gIDs := make([]int, 0, len(grids))

		for _, grid := range grids {
			gIDs = append(gIDs, grid.GID)
		}

		fmt.Println("surrounding grid IDs are ", gIDs)
	}

}