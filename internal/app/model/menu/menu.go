package menu

import (
	"app/bootstrap/Table"
	"app/utils/helper"
	"fmt"
	"github.com/gin-gonic/gin"
)

type TreeList struct {
	Id        	uint        `json:"id"`
	Name     	string      `json:"name"`
	Pid       	uint64      `json:"pid"`
	Label		string 		`json:"label"`	//冗余前端字段
	Value 		uint64		`json:"value"`	//冗余前端字段
	Status 		int			`json:"status"`
	Type		uint8 		`json:"type"`
	Memo 		string		`json:"memo"`
	Sequence    int         `json:"sequence"`
	Url      	string      `json:"url"`
	FullUrl		string		`json:"full_url"`
	Component 	string      `json:"component"`
	Icon      	string      `json:"icon"`
	Children  	[]*TreeList `json:"children,omitempty"`
}

func getMenu(pid int, path string) []*TreeList {
	var menus []Table.Menu
	if pid == 0 {
		helper.Db().
			Table("menu").
			Where("parent_id = ?", pid).
			Where("url = ?", "/admin").
			Order("sequence").
			Find(&menus)
	} else {
		helper.Db().Table("menu").Where("parent_id = ?", pid).Order("sequence").Find(&menus)
	}
	var treeList []*TreeList
	for _, v := range menus {
		if v.Type >= 3 {
			continue
		}
		child := getMenu(int(v.ID), v.Url)
		node := &TreeList{
			Id: v.ID,
			Name: v.Name,
			Label: v.Name,
			Value: uint64(v.ID),
			Status: int(v.Status),
			Type: v.Type,
			Memo: v.Memo,
			Component: v.Component,
			Sequence: v.Sequence,
			Url: v.Url,
			FullUrl: path + "/" + v.Url,
			Pid: v.ParentID,
			Icon: v.Icon,
		}

		node.Children = child
		treeList = append(treeList, node)

	}

	return treeList
}

func getApi(pid int, path string) []*TreeList {
	var menus []Table.Menu
	if pid == 0 {
		helper.Db().
			Table("menu").
			Where("parent_id = ?", pid).
			Order("sequence").
			Find(&menus)
	} else {
		helper.Db().Table("menu").Where("parent_id = ?", pid).Order("sequence").Find(&menus)
	}
	var treeList []*TreeList
	for _, v := range menus {
		child := getMenu(int(v.ID), v.Url)
		node := &TreeList{
			Id: v.ID,
			Name: v.Name,
			Label: v.Name,
			Value: uint64(v.ID),
			Component: v.Component,
			Sequence: v.Sequence,
			Url: v.Url,
			FullUrl: path + "/" + v.Url,
			Pid: v.ParentID,
			Icon: v.Icon,
		}

		node.Children = child
		treeList = append(treeList, node)

	}

	return treeList
}


func Index(ctx *gin.Context) {
	db := helper.Db()
	var fields []Table.Menu
	if err := db.Table("menu").Find(&fields).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, fields)
}

func Show(ctx *gin.Context) {
	var field Table.Menu
	if err := helper.Db().Table("menu").Where("id = ?", ctx.Param("id")).First(&field).Error; err != nil {
		helper.Fail(ctx, "查询失败")
		return
	}

	helper.Success(ctx, field)
}

func Store(ctx *gin.Context) {
	var field Table.Menu
	err := ctx.Bind(&field)
	fmt.Println(field)
	if err != nil {
		fmt.Println(err)
		helper.Fail(ctx, "绑定数据失败")
		return
	}
	if err = helper.Db().Table("menu").Create(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "success")
}

func Update(ctx *gin.Context) {
	var filed Table.Menu
	requestJson := helper.GetRequestJson(ctx)
	filed.ID = helper.Str2Uint(ctx.Param("id"))
	fmt.Println(requestJson)
	if err := helper.Db().Table("menu").Model(&filed).Updates(requestJson).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "更新成功")
}

func Destroy(ctx *gin.Context) {
	var field Table.Menu
	field.ID = helper.Str2Uint(ctx.Param("id"))
	if err := helper.Db().Table("menu").Delete(&field).Error; err != nil {
		helper.Fail(ctx, err.Error())
		return
	}

	helper.Success(ctx, "删除成功")
}

func List(ctx *gin.Context) {
	treeList := getMenu(0, "")
	helper.Success(ctx, treeList)
}

func ApiList(ctx *gin.Context)  {
	treeList := getApi(0, "")
	helper.Success(ctx, treeList)
}
