package userServer

var userInfoDataBase = []*UserInfo{
	{
		ID:   1,
		Name: "user1",
	},
	{
		ID:   2,
		Name: "user2",
	},
	{
		ID:   3,
		Name: "user3",
	},
	{
		ID:   4,
		Name: "user4",
	},
	{
		ID:   5,
		Name: "user5",
	},
	{
		ID:   6,
		Name: "user6",
	},
	{
		ID:   7,
		Name: "user7",
	},
	{
		ID:   8,
		Name: "user8",
	},
}

func GetUserInfoByID(id int64) *UserInfo {
	for _, user := range userInfoDataBase {
		if user.ID == id {
			return user
		}
	}

	return nil
}


func GetUserInfoByName(name string) []*UserInfo {
	var users []*UserInfo
	for _, user := range userInfoDataBase {
		if user.Name == name {
			users = append(users, user)
		}
	}

	return users
}

