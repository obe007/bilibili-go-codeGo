// Code generated by protoc-gen-bm v0.1, DO NOT EDIT.
// source: api.proto

/*
Package v0 is a generated blademaster stub package.
This code was generated with go-common/app/tool/bmgen/protoc-gen-bm v0.1.

It is generated from these files:
	api.proto
*/
package v0

import (
	"context"

	bm "go-common/library/net/http/blademaster"
	"go-common/library/net/http/blademaster/binding"
)

// to suppressed 'imported but not used warning'
var _ *bm.Context
var _ context.Context
var _ binding.StructValidator

var PathCreateDataCreateCacheList = "/live.daoanchor.v0.CreateData/CreateCacheList"
var PathCreateDataCreateLiveCacheList = "/live.daoanchor.v0.CreateData/CreateLiveCacheList"
var PathCreateDataGetContentMap = "/live.daoanchor.v0.CreateData/GetContentMap"
var PathCreateDataCreateDBData = "/live.daoanchor.v0.CreateData/CreateDBData"

var PathPopularityGetAnchorGradeList = "/live.daoanchor.v0.Popularity/GetAnchorGradeList"
var PathPopularityEditAnchorGrade = "/live.daoanchor.v0.Popularity/EditAnchorGrade"
var PathPopularityGetContentList = "/live.daoanchor.v0.Popularity/GetContentList"
var PathPopularityAddContent = "/live.daoanchor.v0.Popularity/AddContent"
var PathPopularityEditContent = "/live.daoanchor.v0.Popularity/EditContent"
var PathPopularityDeleteContent = "/live.daoanchor.v0.Popularity/DeleteContent"

// ====================
// CreateData Interface
// ====================

type CreateDataBMServer interface {
	// CreateCacheList 生成历史数据缓存列表
	CreateCacheList(ctx context.Context, req *CreateCacheListReq) (resp *CreateCacheListResp, err error)

	// CreateLiveCacheList 生成开播历史数据缓存列表
	CreateLiveCacheList(ctx context.Context, req *CreateLiveCacheListReq) (resp *CreateLiveCacheListResp, err error)

	// GetContentMap 获取需要生成历史数据的对象列表
	GetContentMap(ctx context.Context, req *GetContentMapReq) (resp *GetContentMapResp, err error)

	CreateDBData(ctx context.Context, req *CreateDBDataReq) (resp *CreateDBDataResp, err error)
}

var v0CreateDataSvc CreateDataBMServer

func createDataCreateCacheList(c *bm.Context) {
	p := new(CreateCacheListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v0CreateDataSvc.CreateCacheList(c, p)
	c.JSON(resp, err)
}

func createDataCreateLiveCacheList(c *bm.Context) {
	p := new(CreateLiveCacheListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v0CreateDataSvc.CreateLiveCacheList(c, p)
	c.JSON(resp, err)
}

func createDataGetContentMap(c *bm.Context) {
	p := new(GetContentMapReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v0CreateDataSvc.GetContentMap(c, p)
	c.JSON(resp, err)
}

func createDataCreateDBData(c *bm.Context) {
	p := new(CreateDBDataReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v0CreateDataSvc.CreateDBData(c, p)
	c.JSON(resp, err)
}

// RegisterCreateDataBMServer Register the blademaster route
func RegisterCreateDataBMServer(e *bm.Engine, server CreateDataBMServer) {
	v0CreateDataSvc = server
	e.GET("/live.daoanchor.v0.CreateData/CreateCacheList", createDataCreateCacheList)
	e.GET("/live.daoanchor.v0.CreateData/CreateLiveCacheList", createDataCreateLiveCacheList)
	e.GET("/live.daoanchor.v0.CreateData/GetContentMap", createDataGetContentMap)
	e.GET("/live.daoanchor.v0.CreateData/CreateDBData", createDataCreateDBData)
}

// ====================
// Popularity Interface
// ====================

// 人气相关接口
type PopularityBMServer interface {
	// GetAnchorGradeList 获取人气值主播评级列表
	GetAnchorGradeList(ctx context.Context, req *GetAnchorGradeListReq) (resp *GetAnchorGradeListResp, err error)

	// EditAnchorGrade  编辑主播评级对应的人气值数据
	EditAnchorGrade(ctx context.Context, req *EditAnchorGradeReq) (resp *EditAnchorGradeResp, err error)

	// GetContentList  人气内容系数列表
	GetContentList(ctx context.Context, req *GetContentListReq) (resp *GetContentListResp, err error)

	// AddContent 添加内容系数
	AddContent(ctx context.Context, req *AddContentReq) (resp *AddContentResp, err error)

	// EditContent 编辑内容系数
	EditContent(ctx context.Context, req *EditContentReq) (resp *EditContentResp, err error)

	// DeleteContent 删除内容系数
	DeleteContent(ctx context.Context, req *DeleteContentReq) (resp *DeleteContentResp, err error)
}

var v0PopularitySvc PopularityBMServer

func popularityGetAnchorGradeList(c *bm.Context) {
	p := new(GetAnchorGradeListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v0PopularitySvc.GetAnchorGradeList(c, p)
	c.JSON(resp, err)
}

func popularityEditAnchorGrade(c *bm.Context) {
	p := new(EditAnchorGradeReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v0PopularitySvc.EditAnchorGrade(c, p)
	c.JSON(resp, err)
}

func popularityGetContentList(c *bm.Context) {
	p := new(GetContentListReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v0PopularitySvc.GetContentList(c, p)
	c.JSON(resp, err)
}

func popularityAddContent(c *bm.Context) {
	p := new(AddContentReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v0PopularitySvc.AddContent(c, p)
	c.JSON(resp, err)
}

func popularityEditContent(c *bm.Context) {
	p := new(EditContentReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v0PopularitySvc.EditContent(c, p)
	c.JSON(resp, err)
}

func popularityDeleteContent(c *bm.Context) {
	p := new(DeleteContentReq)
	if err := c.BindWith(p, binding.Default(c.Request.Method, c.Request.Header.Get("Content-Type"))); err != nil {
		return
	}
	resp, err := v0PopularitySvc.DeleteContent(c, p)
	c.JSON(resp, err)
}

// RegisterPopularityBMServer Register the blademaster route
func RegisterPopularityBMServer(e *bm.Engine, server PopularityBMServer) {
	v0PopularitySvc = server
	e.GET("/live.daoanchor.v0.Popularity/GetAnchorGradeList", popularityGetAnchorGradeList)
	e.GET("/live.daoanchor.v0.Popularity/EditAnchorGrade", popularityEditAnchorGrade)
	e.GET("/live.daoanchor.v0.Popularity/GetContentList", popularityGetContentList)
	e.GET("/live.daoanchor.v0.Popularity/AddContent", popularityAddContent)
	e.GET("/live.daoanchor.v0.Popularity/EditContent", popularityEditContent)
	e.GET("/live.daoanchor.v0.Popularity/DeleteContent", popularityDeleteContent)
}
