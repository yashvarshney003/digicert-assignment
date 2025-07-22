package controllers

type BookInput struct {
	Title       string `json:"title" binding:"required,min=3,max=100"`
	Description string `json:"description" binding:"required,min=3,max=500"`
	Author      string `json:"author" binding:"required,min=3,max=100"`
}
type BookOutput struct {
	ID          uint   `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Author      string `json:"author"`
}
type BookUpdateInput struct {
	Title       string `json:"title" binding:"required"`
	Description string `json:"description" binding:"required"`
	ID          uint   `json:"id" binding:"required"`
	Author      string `json:"author" binding:"required"`
}
