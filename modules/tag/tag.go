package tag

import (
	"fmt"
	"github.com/changwei4869/wedding/model"
	"net/http"
	"strconv"

	"github.com/changwei4869/wedding/modules/db"
	"github.com/gin-gonic/gin"
)

// GetTagById 根据id获取tag
// @summary 根据id获取tag
// @description 根据id获取tag
// @tags GetTagById
// @param id path string true "id"
// @produce application/json
// @success 200 {object} model.TagsResp "成功获取tag"
// @router /tag/:id [get]
func GetTagById(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "tag id empty")
		return
	}

	tagId, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "id is not a number")
		return
	}

	tag, err := NewTagsService(db.GetDb()).Detail(tagId)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("get tag from db error: %s", err))
		return
	}
	c.JSON(http.StatusOK, tag)
}

func AddTag(c *gin.Context) {

	tag := model.TagsAddReq{}
	// 从请求中解析 JSON 数据到 tag 结构体
	if err := c.BindJSON(&tag); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}

	if err := NewTagsService(db.GetDb()).Add(tag); err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error adding tag to db: %s", err))
		return
	}

	c.JSON(http.StatusCreated, tag) // 返回新创建的标签信息
}

func DeleteTag(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.String(http.StatusBadRequest, "tag id empty")
		return
	}

	tagID, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "id is not a number")
		return
	}
	err = NewTagsService(db.GetDb()).Del(tagID)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("delete tag from db error: %s", err))
		return
	}

	c.String(http.StatusOK, "tag deleted successfully")
}

func EditTag(c *gin.Context) {

	// 获取更新后的tag
	var updatedTag model.TagsEditReq
	if err := c.BindJSON(&updatedTag); err != nil {
		c.String(http.StatusBadRequest, "invalid JSON format")
		return
	}
	// 保存更新后的标签信息到数据库
	err := NewTagsService(db.GetDb()).Edit(updatedTag)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("error updating tag in db: %s", err))
		return
	}

	// 返回更新后的标签信息
	c.JSON(http.StatusOK, updatedTag)
}
