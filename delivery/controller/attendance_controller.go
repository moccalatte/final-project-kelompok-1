package controller

import (
	"final-project-kelompok-1/model/dto"
	"final-project-kelompok-1/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AttendanceController struct {
	uc usecase.AttendanceUseCase
	rg *gin.RouterGroup
}

func (a *AttendanceController) CreateHandler(ctx *gin.Context) {
	var payload dto.AttendanceRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	createdAttendance, err := a.uc.AddAttendance(payload)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusCreated, "Attendance successfully created", createdAttendance)
}

func (a *AttendanceController) GetHandlerByID(ctx *gin.Context) {
	attendanceID := ctx.Param("id")

	attendance, err := a.uc.FindAttendanceByID(attendanceID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Get Attendance by ID", attendance)
}

func (a *AttendanceController) UpdateHandler(ctx *gin.Context) {
	var payload dto.AttendanceRequestDto
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		dto.SendSingleResponse(ctx, http.StatusBadRequest, err.Error(), nil)
		return
	}

	attendanceID := ctx.Param("id")
	updatedAttendance, err := a.uc.UpdateAttendance(payload, attendanceID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Attendance Updated", updatedAttendance)
}

func (a *AttendanceController) DeleteHandler(ctx *gin.Context) {
	attendanceID := ctx.Param("id")

	deletedAttendance, err := a.uc.Delete(attendanceID)
	if err != nil {
		dto.SendSingleResponse(ctx, http.StatusInternalServerError, err.Error(), nil)
		return
	}

	dto.SendSingleResponse(ctx, http.StatusOK, "Attendance Deleted", deletedAttendance)
}

func (a *AttendanceController) Route() {
	a.rg.POST("/attendance", a.CreateHandler)
	a.rg.GET("/attendance/:id", a.GetHandlerByID)
	a.rg.PUT("/attendance/:id", a.UpdateHandler)
	a.rg.DELETE("/attendance/:id", a.DeleteHandler)
}

func NewAttendanceController(uc usecase.AttendanceUseCase, rg *gin.RouterGroup) *AttendanceController {
	return &AttendanceController{uc: uc, rg: rg}
}