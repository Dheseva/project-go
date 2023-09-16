package models

type Airsoftgun struct {
	Id         uint   `json:"id"`
	Name       string `json:"name"`
	Type       string `json:"type"`
	From       string `json:"from"`
	ASGData_id int    `json:"asgdata_id"`
	Available  int    `json:"available"`
	Created_by int    `json:"created_by"`
	Created_at int64  `json:"created_at"`
	Updated_at int64  `json:"updated_at"`
	Deleted_at int64  `json:"deleted_at"`
}

type AirsoftgunData struct {
	Id         uint   `json:"id"`
	Typeammo   string `json:"typeammo"`
	Capacity   string `json:"capacity"`
	Material   string `json:"material"`
	Color      string `json:"color"`
	Seri_no    int    `json:"seri_no"`
	Made_at    int64  `json:"made_at"`
	Created_at int64  `json:"created_at"`
	Updated_at int64  `json:"updated_at"`
	Deleted_at int64  `json:"deleted_at"`
}