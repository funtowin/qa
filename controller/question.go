package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"qa/model"
	util "qa/utils"
	"strconv"
)

// 创建问题
func AddQuestion(c *gin.Context) {
	var q model.Question

	if err := c.ShouldBind(&q); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    util.QuestionInvalidParams,
			"message": util.QuestionInvalidParams.Msg(),
		})
		return
	}

	code := q.Create()
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": code.Msg(),
	})
}

//  查询分类下的所有文章
func GetAllQuestion(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))

	switch {
	case pageSize >= 100:
		pageSize = 100
	case pageSize <= 0:
		pageSize = 10
	}

	if pageNum == 0 {
		pageNum = 1
	}

	var data []model.Question
	var code util.MyCode
	var total int64
	data, total, code = model.GetAllQuestion(pageSize, pageNum)

	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": code.Msg(),
		"data": gin.H{
			"questionList": data,
			"total":        total,
		},
	})
}
