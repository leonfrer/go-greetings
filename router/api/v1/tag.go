package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/go-greeting/model"
	"github.com/go-greeting/pkg/e"
	"github.com/go-greeting/pkg/setting"
	"github.com/go-greeting/pkg/util"
	"github.com/unknwon/com"
	"net/http"
)

//获取多个文章标签
func GetTags(c *gin.Context) {

	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != "" {
		maps["name"] = name
	}

	state := -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["list"] = model.GetTags(util.GetPage(c), setting.PageSize, maps)
	data["total"] = model.GetTagTotal(maps)

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

//新增文章标签 todo exchange to 'application/json' type body & use a better validator
func AddTag(c *gin.Context) {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name, "name").Message("name cannot be empty")
	valid.MaxSize(name, 100, "name").Message("the max length of name is 100")
	valid.Required(createdBy, "created_by").Message("must have create person")
	valid.MaxSize(createdBy, 100, "created_by").Message("the max length of create person name is 100")
	valid.Range(state, 0, 1, "state").Message("state only can be 0 or 1")

	code := e.INVALID_PARAMS
	if !valid.HasErrors() {
		if !model.ExistTagByName(name) {
			code = e.SUCCESS
			model.AddTag(name, state, createdBy)
		} else {
			code = e.ERROR_EXIST_TAG
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

//修改文章标签
func EditTag(c *gin.Context) {
}

//删除文章标签
func DeleteTag(c *gin.Context) {
}
