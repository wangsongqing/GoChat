package user

import (
	"GoChat/pkg/app"
	"GoChat/pkg/database"
	"GoChat/pkg/paginator"
	"github.com/gin-gonic/gin"
)

// IsEmailExist 判断 Email 已被注册
func IsEmailExist(email string) bool {
	var count int64
	database.DB.Model(User{}).Where("email = ?", email).Count(&count)
	return count > 0
}

// IsPhoneExist 判断手机号已被注册
func IsPhoneExist(phone string) bool {
	var count int64
	database.DB.Model(User{}).Where("phone = ?", phone).Count(&count)
	return count > 0
}

func GetUser(id int) (userModel User) {
	database.DB.Where("id = ?", id).Find(&userModel)
	return
}

func IsUserExist(id int) bool {
	var count int64
	database.DB.Model(User{}).Where("id = ?", id).Count(&count)
	return count > 0
}

func GetByPhone(phone string) (userModel User) {
	database.DB.Where("phone = ?", phone).Find(&userModel)
	return
}

func GetByName(name string) (userModel User) {
	database.DB.Where("name = ?", name).Find(&userModel)
	return
}

func GetByEmail(email string) (userModel User) {
	database.DB.Where("email = ?", email).Find(&userModel)
	return
}

func GetAccount(account string) (userModel User) {
	database.DB.Where("(phone = ? or email = ?)", account, account).Find(&userModel)
	return
}

func Get(id string) (userModel User) {
	database.DB.Where("id = ?", id).Find(&userModel)
	return
}

// All 获取所有用户数据
func All() (users []User) {
	database.DB.Find(&users)
	return
}

// Paginate 分页内容
func Paginate(c *gin.Context, perPage int) (users []User, paging paginator.Paging) {
	paging = paginator.Paginate(
		c,
		database.DB.Model(User{}),
		&users,
		app.V1URL(database.TableName(&User{})),
		perPage,
	)
	return
}
