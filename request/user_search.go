package request

type UserSearch struct {
	DomainId string `json:"domain_id"`
	Enabled  string `json:"enabled"`
	Name     string `json:"name"`
	UserId   string `json:"user_id"`
	PageSize int    `json:"page_size"`
	PageNum  int    `json:"page_num"`
}
