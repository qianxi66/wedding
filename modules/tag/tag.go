package tag

import (
	"fmt"
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
