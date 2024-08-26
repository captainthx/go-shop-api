package response

type LoginResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type CategoryResponse struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
}

type UpLodaFileResponse struct {
	FileName string
	FileUrl  string
	Size     float32
}
