package handler

import (
	"net/http"
	"timesheet/internal/model"
	"timesheet/internal/repository"
	"timesheet/internal/timesheet"

	"github.com/gin-gonic/gin"
)

type Date struct {
	Year  int `json:"year"`
	Month int `json:"month"`
}

type RequestIncome struct {
	Year     int             `json:"year"`
	Month    int             `json:"month"`
	MemberID string          `json:"member_id"`
	Incomes  []model.Incomes `json:"incomes"`
}

type TimesheetAPI struct {
	Timesheet           timesheet.TimesheetGateways
	TimesheetRepository repository.TimesheetRepositoryGateways
}

func (api TimesheetAPI) GetSummaryHandler(context *gin.Context) {
	var date Date
	context.ShouldBindJSON(&date)

	transactionTimesheet, _ := api.TimesheetRepository.GetSummary(date.Year, date.Month)

	context.JSON(http.StatusOK, transactionTimesheet)
}

func (api TimesheetAPI) UpdateIncomeHandler(context *gin.Context) {
	var requestIncome RequestIncome
	context.ShouldBindJSON(&requestIncome)

	context.JSON(http.StatusOK, "")
}
