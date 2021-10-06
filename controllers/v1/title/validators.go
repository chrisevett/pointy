package controllers

type CreateTitleInput struct {
	Name     string `json:"name" binding:"required"`
	Designer string `json:"designer" binding:"required"`
}
