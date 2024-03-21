package request

type ServiceSearch struct {
	Type string `json:"type"`
}

type CreateServiceRequest struct {
	Service struct {
		Name        string `json:"name" binding:"required" validate:"required"`
		Type        string `json:"type" binding:"required" validate:"required"`
		Description string `json:"description"`
		Enabled     *bool  `json:"enabled"`
	} `json:"service"`
}

type UpdateServiceRequest struct {
	Service struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		Description string `json:"description"`
		Enabled     *bool  `json:"enabled"`
	} `json:"service"`
}
