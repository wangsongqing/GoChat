package group_people

import "GoChat/pkg/database"

func GetGroupMan(groupId int64) (gp []GroupPeople) {
	database.DB.Where("group_id = ? and status = ?", groupId, StatusOk).Find(&gp)
	return
}

func IsExistGroup(UserId int, GroupId int) bool {
	var count int64
	database.DB.Model(GroupPeople{}).Where("user_id = ? and group_id = ?", UserId, GroupId).Count(&count)
	return count > 0
}
