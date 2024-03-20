package request

type EndpointSearch struct {
	Interface string `json:"interface"`

	ServiceId string `json:"service_id"`
}

type CreateEndpointRequest struct {
	Endpoint struct {
		Enabled   *bool  `json:"enabled"`
		Interface string `json:"interface" binding:"required" validate:"required"`
		Region    string `json:"region"`
		RegionId  string `json:"region_id"`
		ServiceId string `json:"service_id" binding:"required" validate:"required"`
		Url       string `json:"url" binding:"required" validate:"required"`
	} `json:"endpoint"`
}

type UpdateEndpointRequest struct {
	Endpoint struct {
		Enabled   *bool  `json:"enabled"`
		Interface string `json:"interface"`
		Region    string `json:"region"`
		RegionId  string `json:"region_id"`
		ServiceId string `json:"service_id" `
		Url       string `json:"url"`
	} `json:"endpoint"`
}
