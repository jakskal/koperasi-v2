package dto

type CreateSavingTypeRequest struct {
	Name string `json:"name" binding:"required"`
}

type UpdateSavingTypeRequest struct {
	ID   int
	Name string `json:"name"`
}
