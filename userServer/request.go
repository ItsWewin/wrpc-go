package userServer

type GetUserInfoRequest struct {
	ID int64 `json:"id"`
}

type GetUserInfoByNameRequest struct {
	Name string `json:"name"`
}