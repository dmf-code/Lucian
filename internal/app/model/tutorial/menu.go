package tutorial

import (
	"app/library/helper"
	"github.com/gin-gonic/gin"
)

type TreeList struct {
	Id        	uint        `json:"id"`
	Name     	string      `json:"name"`
	Pid       	uint64      `json:"pid"`
	Type		uint8 		`json:"type"`
	Icon      	string      `json:"icon"`
	Label		string 		`json:"label"`	//冗余前端字段
	Value 		uint64		`json:"value"`	//冗余前端字段
	Children  	[]*TreeList `json:"children,omitempty"`
}

func getMenuTree(pid int) []*TreeList {
	db := helper.Db()
	var menus []Tutorial
	if err := db.Table("tutorial").Where("parent_id = ?", pid).Find(&menus).Error; err != nil {
		panic(err)
	}
	var treeList []*TreeList
	for _, v := range menus {
		// 筛除非选中节点
		child := getMenuTree(int(v.ID))
		node := &TreeList{
			Id: v.ID,
			Name: v.Title,
			Value: uint64(v.ID),
			Label: v.Title,
			Type: uint8(v.Type),
			Pid: uint64(v.ParentId),
			Icon: v.Icon,
		}

		node.Children = child
		treeList = append(treeList, node)
	}
	return treeList
}

func List(ctx *gin.Context) {
	t := ctx.Param("pid")
	treeList := getMenuTree(helper.Str2Int(t))
	helper.Success(ctx, treeList)
}