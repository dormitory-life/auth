package requestmodels

type (
	RefreshTokensRequest struct {
		UserId       string
		DormitoryId  string
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}

	RefreshTokensResponse struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
	}
)
