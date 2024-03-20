package request

type RegionSearch struct {
	ParentRegionId string `json:"parent_region_id"`
}

type CreateRegionRequest struct {
	Region struct {
		Id             string `json:"id"`
		Description    string `json:"description" binding:"required" validate:"required"`
		ParentRegionId string `json:"parent_region_id"`
	} `json:"region"`
}

type UpdateRegionRequest struct {
	Region struct {
		Id             string `json:"id"`
		Description    string `json:"description"`
		ParentRegionId string `json:"parent_region_id"`
	} `json:"region"`
}
