package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/muxi-mini-project/2020-sharing-backend/model"
	"log"
)

type data struct {
	CollectlistId   int    `json:"collectlist_id"`
	CollectlistName string `json:"collectlist_name"`
}

type tmpstring struct {
	CollectlistName string `json:"collectlist_name"`
}

type tmpint struct {
	CollectlistId int `json:"collectlist_id"`
}

// @Summary 新建收藏夹
// @Description 新建收藏夹
// @Tags collectlist
// @Accept json
// @Produce json
// @Param token header string true "user的认证令牌"
// @Param data body tmpstring true "收藏夹名称"
// @Success 200 {object} model.Res "{"message":"收藏夹建立成功"}"
// @Failure 401 {object} model.Error "{"message":"身份认证错误，请先登录或注册！"}"
// @Failure 400 {object} model.Error "{"message":"Bad Request!"}"
// @Failure 404 {object} model.Error "{"message":"收藏夹建立失败“} "
// @Router /user/collect_list/create/ [post]
func CreateNewCollectlist(c *gin.Context) {
	var data tmpstring
	token := c.Request.Header.Get("token")
	if len(token) == 0 {
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
	key, _ := model.Token_info(token)
	if err := c.BindJSON(&data); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if err := model.CreateNewcollectlist(data.CollectlistName, key); !err {
		log.Print("创建收藏夹失败")
		c.JSON(404, gin.H{
			"message": "收藏夹建立失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "收藏夹建立成功",
	})
}

// @Summary 收藏夹改名
// @Description 收藏夹重命名
// @Tags collectlist
// @Accept json
// @Produce json
// @Param token header string true "user的认证令牌"
// @Param data body data true "收藏夹id与新收藏夹名字"
// @Success 200 {object} model.Res "{"message":"收藏夹改名成功"}"
// @Failure 401 {object} model.Error "{"message":"身份认证错误，请先登录或注册！"}"
// @Failure 400 {object} model.Error "{"message":"Bad Request!"}"
// @Failure 404 {object} model.Error "{"message":"未找到收藏夹“} or {"message":"数据更新失败“}"
// @Router /user/collect_list/ [put]
func ChangeCollectionlistName(c *gin.Context) {
	var data data
	var tmpcollectlist model.Collect_list
	//利用token解码出的userid来检验进行该操作的是否为本人
	token := c.Request.Header.Get("token")
	if len(token) == 0 {
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
	//key, _ := model.Token_info(token)
	//if err := model.DB.Self.Model(&model.Collect_list{}).Where(&model.Collect_list{UserID:key,CollectlistId:data.CollectlistId})
	if err := c.BindJSON(&data); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if err := model.DB.Self.Model(&model.Collect_list{}).Where(&model.Collect_list{CollectlistId: data.CollectlistId}).First(&tmpcollectlist).Error; err != nil {
		log.Println(err)
		log.Print("获取收藏夹信息失败")
		c.JSON(404, gin.H{
			"message": "未找到收藏夹",
		})
		return
	}
	tmpcollectlist.CollectlistName = data.CollectlistName
	if err := model.DB.Self.Model(&model.Collect_list{}).Where(&model.Collect_list{CollectlistId: data.CollectlistId}).Update("collectlist_name", tmpcollectlist.CollectlistName).Error; err != nil {
		log.Println(err)
		log.Print("更新数据失败")
		c.JSON(404, gin.H{
			"message": "更新数据失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "收藏夹改名成功",
	})
}

// @Summary 删除收藏夹
// @Description 删除收藏夹
// @Tags collectlist
// @Accept json
// @Produce json
// @Param token header string true "user的认证令牌"
// @Param data body tmpint true "收藏夹id"
// @Success 200 {object} model.Res "{"message":"收藏夹删除成功"}"
// @Failure 401 {object} model.Error "{"message":"身份认证错误，请先登录或注册！"}"
// @Failure 400 {object} model.Error "{"message":"Bad Request!"}"
// @Failure 404 {object} model.Error "{"message":"收藏夹删除失败“}"
// @Router /user/collect_list/delete/ [delete]
func DeleteCollectlist(c *gin.Context) {
	var data tmpint
	token := c.Request.Header.Get("token")
	if len(token) == 0 {
		c.JSON(401, gin.H{
			"message": "身份认证错误，请先登录或注册！",
		})
		return
	}
	//key, _ := model.Token_info(token)
	if err := c.BindJSON(&data); err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request!",
		})
		return
	}
	if err := model.DB.Self.Where(&model.Collect_list{CollectlistId: data.CollectlistId}).Delete(&model.Collect_list{}).Error; err != nil {
		log.Println(err)
		log.Print("收藏夹删除失败")
		c.JSON(404, gin.H{
			"message": "收藏夹删除失败",
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "收藏夹删除成功",
	})
}
