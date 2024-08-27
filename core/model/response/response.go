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

type ProductResponse struct {
	ID            uint                   `json:"id"`
	Name          string                 `json:"name"`
	Price         float64                `json:"price"`
	Quantity      int                    `json:"quantity"`
	CreatedAt     string                 `json:"created_at"`
	ProductImages []ProductImageResponse `json:"product_images"`
}

type ProductImageResponse struct {
	ID        uint   `json:"id"`
	ProductID uint   `json:"product_id"`
	Url       string `json:"url"`
}
