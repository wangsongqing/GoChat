package group_people

import "GoChat/pkg/database"

func GetGroupMan(groupId int64) (gp []GroupPeople) {
	database.DB.Where("group_id = ? and status = ?", groupId, StatusOk).Find(&gp)
	return
}
