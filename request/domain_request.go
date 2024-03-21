package request

type DomainSearch struct {
	Enabled string `json:"enabled"`
	Name    string `json:"name"`
}

type CreateDomainRequest struct {
	Domain struct {
		Description string `json:"description"`
		Name        string `json:"name" binding:"required" validate:"required"`
		Enabled     *bool  `json:"enabled"`
	} `json:"domain"`
}

type UpdateDomainRequest struct {
	Domain struct {
		Description string `json:"description"`
		Name        string `json:"name"`
		Enabled     *bool  `json:"enabled"`
	} `json:"domain"`
}
