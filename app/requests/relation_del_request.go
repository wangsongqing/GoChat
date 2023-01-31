package requests

import (
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type RelationDelRequest struct {
	Type       int    `json:"type,omitempty" valid:"type"`
	Desc       string `json:"desc,omitempty" valid:"desc"`
	TargetName string `json:"target_name,omitempty" valid:"target_name"`
}

func CheckRelation(data interface{}, ctx *gin.Context) map[string][]string {
	rules := govalidator.MapData{
		"type":        []string{"required"},
		"desc":        []string{"required"},
		"target_name": []string{"required"},
	}
	messages := govalidator.MapData{
		"type": []string{
			"required:好友名称必填，参数名称 type",
		},
		"desc": []string{
			"required:描述字段不能为空，参数名称 desc",
		},
		"target_name": []string{
			"required:用户名为必填，参数名称 target_name",
		},
	}

	errs := validate(data, rules, messages)

	return errs
}
