package following_list

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/handler"
	"github.com/muxi-mini-project/2020-sharing-backend/model"

	//"encoding/json"
	. "fmt"
)

// @Summary 显示用户following
// @Description 显示用户fllowing列表与动态
// @Tags user
// @Accept json
// @Produce json
// @Param User.User_id body model.User.User_id true "查看人的ID"
// @Param token header string true "token"
// @Success 200 {object} model.Following2 "{"messsge":"successfully","total":数量，"following_lists",数组}"
// @Failure 401 {object} error.Error "{"error_code":"10001", "message":"Token Invalid."} 身份认证失败 重新登录 or {"error_code":"00002", "message":"wrong mysql."} or {"error_code":"10001", "message":"User does not Existed!"} 用户不存在```"
// @Failure 400 {object} error.Error "{"error_code":"00001", "message":"Fail."} or {"error_code":"00002", "message":"Lack Param Or Param Not Satisfiable."}"
// @Failure 500 {object} error.Error "{"error_code":"30001", "message":"Fail."} 失败"
// @Router /user/following_list/ [get]

func FollowingList(c *gin.Context) {

	Token := c.Request.Header.Get("Token")

	Println(Token)
	_, error := model.Token_info(Token)
	if !error {
		c.JSON(401, gin.H{
			"message": "wrong token",
		})
		return
	}

	var data model.User
	if err := c.BindJSON(&data); err != nil {
		handler.SendBadRequest(c)
		return
	}

	if !model.CheckUserByUser_id(data.User_id) {
		c.JSON(401, gin.H{
			"message": "User doesn't Existed!",
		})
		return
	}

	fid, err := model.GetFollowingid(data.User_id)
	rows, err := model.FollowingList(fid)
	num, err := model.FansNum(data.User_id)

	if err != nil {
		c.JSON(401, gin.H{
			"message": "wrong mysql",
			"err":     err,
		})
		return
	}

	c.JSON(200, gin.H{
		"message":         "successfully",
		"following_lists": rows,
		"total":           num,
	})
	return
}
