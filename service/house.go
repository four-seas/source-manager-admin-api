package service

import (
	"source_manager_admin_api/global"
	"source_manager_admin_api/model"
	"source_manager_admin_api/model/request"
)

// @title    GetHousesInfoList
// @description   get house list by pagination, 分页获取数据
// @auth                      （2020/08/24  20:22）
// @param     info             request.PageInfo
// @return    err              error
// @return    list             interface{}
// @return    total            int

func GetHousesInfoList(info request.PageInfo) (err error, list interface{}, total int) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&model.Houses{})
	var hourseList []model.Houses
	err = db.Count(&total).Error
	err = db.Limit(limit).Offset(offset).Find(&hourseList).Error
	return err, hourseList, total
}
