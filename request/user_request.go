package request

type CreateUserRequest struct {
	User struct {
		Name        string `json:"name" binding:"required" validate:"required"` //用户名
		Description string `json:"description"`                                 //描述信息
		DomainId    string `json:"domain_id"`
		Enabled     *bool  `json:"enabled"`
		Email       string `json:"email"`
		Password    string `json:"password" binding:"required" validate:"required"`
	} `json:"user"`
}

type UpdateUserRequest struct {
	User struct {
		Name        string `json:"name"`        //用户名
		Description string `json:"description"` //描述信息
		DomainId    string `json:"domain_id"`
		Enabled     *bool  `json:"enabled"`
		Email       string `json:"email"`
	} `json:"user"`
}

type ChangePasswordRequest struct {
	User struct {
		Password         string `json:"password" binding:"required" validate:"required"`
		OriginalPassword string `json:"original_password" binding:"required" validate:"required"`
	} `json:"user"`
}
