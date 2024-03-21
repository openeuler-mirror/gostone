package request

type ProjectSearch struct {
	DomainId string `json:"domain_id"`
	Enabled  string `json:"enabled"`
	IsDomain string `json:"is_domain"`
	Name     string `json:"name"`
	ParentId string `json:"parent_id"`
	PageSize int    `json:"page_size"`
	PageNum  int    `json:"page_num"`
}

type CreateProjectRequest struct {
	Project struct {
		Description string `json:"description"`
		Name        string `json:"name" binding:"required" validate:"required"`
		Enabled     *bool  `json:"enabled"`
		DomainId    string `json:"domain_id"`
		ParentId    string `json:"parent_id"`
	} `json:"project"`
}

type UpdateProjectRequest struct {
	Project struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		Enabled     *bool  `json:"enabled"`
		DomainId    string `json:"domain_id"`
		ParentId    string `json:"parent_id"`
	} `json:"project"`
}
