package controller

import (
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CourseController struct {
	uc usecase.CourseUseCase
	rg *gin.RouterGroup
}

func (c *CourseController) CreateHandler(ctx *gin.Context) {
	var payload dto.CourseRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	createdCourse, err := c.uc.AddCourse(payload)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusCreated, "Course successfuly Created", createdCourse)
}

func (c *CourseController) GetHandlerByID(ctx *gin.Context) {
	courseId := ctx.Param("id")

	course, err := c.uc.FindCourseByID(courseId)

	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	dto.SendSingleResponse(ctx, http.StatusOK, "Get Student by ID", course)
}

func (c *CourseController) UpdateHandler(ctx *gin.Context) {
	var payload dto.CourseRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}
	coursetId := ctx.Param("id")

	updatedCourse, err := c.uc.UpdateCourse(payload, coursetId)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Customer Updated", updatedCourse)

}

func (c *CourseController) DeleteHandler(ctx *gin.Context) {
	coursetId := ctx.Param("id")

	deletedCourse, err := c.uc.DeleteCourse(coursetId)

	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Customer Deleted", deletedCourse)
}

func (c *CourseController) Route() {
	c.rg.POST("/course", c.CreateHandler)
	c.rg.GET("/course/:id", c.GetHandlerByID)
	c.rg.PUT("/course/:id", c.UpdateHandler)
	c.rg.DELETE("/course/:id", c.DeleteHandler)
}

func NewCourseController(uc usecase.CourseUseCase, rg *gin.RouterGroup) *CourseController {
	return &CourseController{uc: uc, rg: rg}
}