package request

type RoleSearch struct {
	Name     string `json:"name"`
	DomainId string `json:"domain_id"`
	PageSize int    `json:"page_size"`
	PageNum  int    `json:"page_num"`
}

type CreateRoleRequest struct {
	Role struct {
		Name     string `json:"name" binding:"required" validate:"required"`
		DomainId string `json:"domain_id"`
	}
}
type UpdateRoleRequest struct {
	Role struct {
		Name     string `json:"name"`
		DomainId string `json:"domain_id"`
	}
}
