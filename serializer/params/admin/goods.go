package admin

type CreateCakeAttributeParam struct {
	AttributeName string `json:"attributeName" binding:"required"`
}

type UpdateCakeAttributeParam struct {
	AttributeName string `json:"attributeName" binding:"required"`
	AttributeId   string `json:"attributeId" binding:"required"`
}

type DeleteCakeAttributeParam struct {
	AttributeId string `json:"attributeId" binding:"required"`
}

type CreateCakeBrandParam struct {
	BrandName string `json:"brandName" binding:"required"`
}

type UpdateCakeBrandParam struct {
	BrandName string `json:"brandName" binding:"required"`
	BrandId   string `json:"brandId" binding:"required"`
}

type DeleteCakeBrandParam struct {
	BrandId string `json:"brandId" binding:"required"`
}

type CreateCakeCategoryParam struct {
	CategoryName string `json:"categoryName" binding:"required"`
}

type UpdateCakeCategoryParam struct {
	CategoryName string `json:"categoryName" binding:"required"`
	CategoryId   string `json:"categoryId" binding:"required"`
}

type DeleteCakeCategoryParam struct {
	CategoryId string `json:"categoryId" binding:"required"`
}
