package models

type ForUse struct {
	Id            uint   `json:"id"`
	Use_for       string `json:"use_for"`
	User_id       int    `json:"user_id"`
	Airsoftgun_id int    `json:"airsoftgun_id"`
	Comment       string `json:"comment"`
	Expired_at    int64  `json:"expired_at"`
	Created_at    int64  `json:"created_at"`
	Updated_at    int64  `json:"updated_at"`
	Deleted_at    int64  `json:"deleted_at"`
}