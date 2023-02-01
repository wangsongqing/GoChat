package group

import "GoChat/pkg/database"

const StatusOk = 1
const StatusOut = 2

// UpdateStatus 解散群
func UpdateStatus(id int) int64 {
	res := database.DB.Model(GroupChat{}).Where("id = ?", id).Update("status", StatusOut)

	return res.RowsAffected
}

func GetGroupInfo(id int) (gc GroupChat) {
	database.DB.Model(GroupChat{}).Where("id = ?", id).Find(&gc)
	return
}
