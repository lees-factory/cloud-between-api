package request

type SocialLoginRequest struct {
	SocialID string `json:"socialId" binding:"required"`
	Provider string `json:"provider" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
}
