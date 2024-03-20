package request

type TokenRequest struct {
	UserId string `json:"user_id" binding:"required" validate:"required"`
}
