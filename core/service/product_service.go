package service

import (
	"go-shop-api/adapters/errs"
	"go-shop-api/common"
	"go-shop-api/core/domain"
	"go-shop-api/core/model/response"
	"go-shop-api/core/ports"
	"time"
)

type productService struct {
	repo ports.ProductRepository
}

func NewProductService(repo ports.ProductRepository) ports.ProductService {
	return &productService{repo: repo}
}

// GetProductList implements ports.ProductService.
func (p *productService) GetProductList(page, limit int, sort string) (*common.Pagination, error) {
	pagiantion := &common.Pagination{}

	if page != 0 {
		pagiantion.Page = page
	}
	if limit != 0 {
		pagiantion.Limit = limit
	}
	if sort != "" {
		pagiantion.Sort = sort
	}

	result, err := p.repo.FindAll(pagiantion)
	if err != nil {
		return nil, errs.NewUnexpectedError(err.Error())
	}

	products, ok := result.Items.([]domain.Product)
	if !ok {
		return nil, errs.NewUnexpectedError("Invalid product data")
	}

	var response []response.ProductResponse

	for _, product := range products {
		response = append(response, mapProductToResponse(product))
	}

	result.Items = response
	return result, nil
}

func mapProductToResponse(product domain.Product) response.ProductResponse {
	return response.ProductResponse{
		ID:            product.ID,
		Name:          product.Name,
		Price:         product.Price,
		Quantity:      product.Quantity,
		CreatedAt:     product.CreatedAt.Format(time.RFC3339),
		ProductImages: mapProductImagesToResponse(product.ProductImage),
	}
}

func mapProductImagesToResponse(images []domain.ProductImage) []response.ProductImageResponse {
	var responseImages []response.ProductImageResponse
	for _, img := range images {
		responseImages = append(responseImages, response.ProductImageResponse{
			ID:        img.ID,
			ProductID: img.ProductID,
			Url:       img.URL,
		})
	}
	return responseImages
}
