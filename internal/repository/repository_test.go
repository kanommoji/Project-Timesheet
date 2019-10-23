package repository_test

import (
	"testing"
	"time"
	"timesheet/internal/model"
	. "timesheet/internal/repository"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"

	"github.com/stretchr/testify/assert"
)

func Test_GetSummary_Input_Year_2017_Month_12_Should_Be_TransactionTimesheet(t *testing.T) {
	expected := []model.TransactionTimesheet{
		{
			ID:                     "001201712siam_chamnankit",
			MemberID:               "001",
			MemberNameTH:           "ประธาน ด่านสกุลเจริญกิจ",
			Month:                  12,
			Year:                   2017,
			Company:                "siam_chamnankit",
			Coaching:               85000.00,
			Training:               30000.00,
			Other:                  40000.00,
			TotalIncomes:           155000.00,
			Salary:                 80000.00,
			IncomeTax1:             5000.00,
			SocialSecurity:         0.00,
			NetSalary:              75000.00,
			Wage:                   75000.00,
			IncomeTax53Percentage:  10,
			IncomeTax53:            7500.00,
			NetWage:                67500.00,
			NetTransfer:            142500.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
			DateTransfer:           "",
			Comment:                "",
		},
	}
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}
	year := 2017
	month := 12

	actual, err := repository.GetSummary(year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_CreateIncome_Input_Year_2017_Month_12_MemberID_001_Income_Should_Be_No_Error(t *testing.T) {
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	year := 2017
	month := 12
	memberID := "001"
	incomes := model.Incomes{
		Day:                      28,
		StartTimeAM:              startTimeAM,
		EndTimeAM:                endTimeAM,
		StartTimePM:              startTimePM,
		EndTimePM:                endTimePM,
		Overtime:                 0,
		TotalHours:               totalHours,
		CoachingCustomerCharging: 15000.00,
		CoachingPaymentRate:      10000.00,
		TrainingWage:             0.00,
		OtherWage:                0.00,
		Company:                  "siam_chamnankit",
		Description:              "[KBTG] 2 Days Agile Project Management",
	}
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.CreateIncome(year, month, memberID, incomes)

	assert.Equal(t, nil, err)
}

func Test_GetMemberByID_Input_MemberID_001_Should_Be_Member(t *testing.T) {
	expected := []model.Member{
		{
			ID:                    1,
			MemberID:              "001",
			Company:               "siam_chamnankit",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			MemberNameENG:         "Prathan Dansakulcharoenkit",
			Email:                 "prathan@scrum123.com",
			OvertimeRate:          0.00,
			RatePerDay:            15000.00,
			RatePerHour:           1875.00,
			Salary:                80000.00,
			IncomeTax1:            5000.00,
			SocialSecurity:        0.00,
			IncomeTax53Percentage: 10,
			TravelExpense:         0.00,
			Status:                "wage",
		},
		{
			ID:                    2,
			MemberID:              "001",
			Company:               "shuhari",
			MemberNameTH:          "ประธาน ด่านสกุลเจริญกิจ",
			MemberNameENG:         "Prathan Dansakulcharoenkit",
			Email:                 "prathan@scrum123.com",
			OvertimeRate:          0.00,
			RatePerDay:            15000.00,
			RatePerHour:           1875.00,
			Salary:                0.00,
			IncomeTax1:            0.00,
			SocialSecurity:        0.00,
			IncomeTax53Percentage: 10,
			TravelExpense:         0.00,
			Status:                "wage",
		},
	}
	memberID := "001"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetMemberByID(memberID)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_GetIncomes_Input_MemberID_006_Year_2019_Month_12_Should_Be_Incomes_Day_11_And_12(t *testing.T) {
	startTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 09:00:00")
	endTimeAM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 12:00:00")
	startTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 13:00:00")
	endTimePM, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 18:00:00")
	totalHours, _ := time.Parse("2006-01-02 15:04:05", "2018-12-01 08:00:00")
	expected := []model.Incomes{
		{
			ID:                       58,
			MemberID:                 "006",
			Month:                    12,
			Year:                     2019,
			Day:                      11,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			Overtime:                 0,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "work at TN",
		}, {
			ID:                       59,
			MemberID:                 "006",
			Month:                    12,
			Year:                     2019,
			Day:                      12,
			StartTimeAM:              startTimeAM,
			EndTimeAM:                endTimeAM,
			StartTimePM:              startTimePM,
			EndTimePM:                endTimePM,
			Overtime:                 0,
			TotalHours:               totalHours,
			CoachingCustomerCharging: 0.00,
			CoachingPaymentRate:      0.00,
			TrainingWage:             0.00,
			OtherWage:                0.00,
			Company:                  "shuhari",
			Description:              "work at TN",
		},
	}
	memberID := "006"
	month := 12
	year := 2019
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet?parseTime=true")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetIncomes(memberID, year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_VerifyTransactionTimsheet_Input_Transaction_MemberID_001_Should_Be_Create_TransactionTimesheet_And_Update_TransactionTimesheet(t *testing.T) {
	transactionTimesheet := []model.TransactionTimesheet{
		{
			MemberID:               "001",
			Month:                  12,
			Year:                   2019,
			Company:                "shuhari",
			MemberNameTH:           "ประธาน ด่านสกุลเจริญกิจ",
			Coaching:               20000.00,
			Training:               0.00,
			Other:                  6500.00,
			TotalIncomes:           6500.00,
			Salary:                 25000.00,
			IncomeTax1:             0.00,
			SocialSecurity:         750.00,
			NetSalary:              24250.00,
			Wage:                   6500.00,
			IncomeTax53Percentage:  5,
			IncomeTax53:            325.00,
			NetWage:                6175.00,
			NetTransfer:            30425.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
		},
		{
			MemberID:               "001",
			Month:                  12,
			Year:                   2019,
			Company:                "siam_chamnankit",
			MemberNameTH:           "ประธาน ด่านสกุลเจริญกิจ",
			Coaching:               30000.00,
			Training:               0.00,
			Other:                  6500.00,
			TotalIncomes:           6500.00,
			Salary:                 25000.00,
			IncomeTax1:             0.00,
			SocialSecurity:         750.00,
			NetSalary:              24250.00,
			Wage:                   6500.00,
			IncomeTax53Percentage:  5,
			IncomeTax53:            325.00,
			NetWage:                6175.00,
			NetTransfer:            30425.00,
			StatusCheckingTransfer: "รอการตรวจสอบ",
		},
	}
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.VerifyTransactionTimsheet(transactionTimesheet)

	assert.Equal(t, nil, err)
}

func Test_CreateTransactionTimsheet_Input_Transaction_MemberID_006_Should_Be_No_Error(t *testing.T) {
	transactionID := "006201912shuhari"
	transactionTimesheet := model.TransactionTimesheet{
		MemberID:               "006",
		MemberNameTH:           "ภาณุมาศ แสนโท",
		Month:                  12,
		Year:                   2019,
		Company:                "shuhari",
		Coaching:               0.00,
		Training:               0.00,
		Other:                  6500.00,
		TotalIncomes:           6500.00,
		Salary:                 25000.00,
		IncomeTax1:             0.00,
		SocialSecurity:         750.00,
		NetSalary:              24250.00,
		Wage:                   6500.00,
		IncomeTax53Percentage:  5,
		IncomeTax53:            325.00,
		NetWage:                6175.00,
		NetTransfer:            30425.00,
		StatusCheckingTransfer: "รอการตรวจสอบ",
	}
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.CreateTransactionTimsheet(transactionTimesheet, transactionID)

	assert.Equal(t, nil, err)
}

func Test_UpdateTransactionTimsheet_Input_Transaction_MemberID_001_Should_Be_No_Error(t *testing.T) {
	transactionID := "001201911shuhari"
	transactionTimesheet := model.TransactionTimesheet{
		MemberID:              "001",
		Month:                 11,
		Year:                  2019,
		Company:               "shuhari",
		Coaching:              10000.00,
		Training:              10000.00,
		Other:                 6500.00,
		TotalIncomes:          6500.00,
		Salary:                25000.00,
		IncomeTax1:            1000.00,
		SocialSecurity:        750.00,
		NetSalary:             24250.00,
		Wage:                  6500.00,
		IncomeTax53Percentage: 5,
		IncomeTax53:           325.00,
		NetWage:               6175.00,
		NetTransfer:           30425.00,
	}
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdateTransactionTimsheet(transactionTimesheet, transactionID)

	assert.Equal(t, nil, err)
}

func Test_CreateTimsheet_Input_Timesheet_MemberID_006_Should_Be_No_Error(t *testing.T) {
	memberID := "006"
	month := 12
	year := 2019

	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.CreateTimesheet(memberID, year, month)

	assert.Equal(t, nil, err)
}

func Test_UpdateTimsheet_Input_Timesheet_MemberID_007_Year_2019_Month_12_Should_Be_No_Error(t *testing.T) {
	timesheet := model.Timesheet{
		TotalHours:                    "120:30:30",
		TotalCoachingCustomerCharging: 90000.00,
		TotalCoachingPaymentRate:      10000.00,
		TotalTrainigWage:              20000.00,
		TotalOtherWage:                30000.00,
		PaymentWage:                   60000.00,
	}
	memberID := "007"
	month := 12
	year := 2019
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdateTimesheet(timesheet, memberID, year, month)

	assert.Equal(t, nil, err)
}

func Test_GetTimsheet_Input_MemberID_003_Month_12_Year_2017_Should_Be_Timesheet(t *testing.T) {
	expected := model.Timesheet{
		ID:                            "003201712",
		MemberID:                      "003",
		Month:                         12,
		Year:                          2017,
		TotalHours:                    "88:00:00",
		TotalCoachingCustomerCharging: 0.00,
		TotalCoachingPaymentRate:      0.00,
		TotalTrainigWage:              120000.00,
		TotalOtherWage:                0.00,
		PaymentWage:                   120000.00,
	}
	memberID := "003"
	month := 12
	year := 2017

	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetTimesheet(memberID, year, month)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_UpdateStatusTransfer_Input_TransactionID_004201912siam_chamnankit_Status_TransferSuccess_Date_30_12_2019_Comment_FlightTicket_Should_Be_No_Error(t *testing.T) {
	transactionID := "004201912siam_chamnankit"
	status := "โอนเงินเรียบร้อยแล้ว"
	date := "30/12/2019"
	comment := "หักค่าตั๋วเครื่องบิน"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdateStatusTransfer(transactionID, status, date, comment)

	assert.Equal(t, nil, err)
}

func Test_DeleteIncome_Input_Year_2018_Month_12_MemberID_005_Day_15_Should_Be_No_Error(t *testing.T) {
	year := 2018
	month := 12
	day := 15
	memberID := "005"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.DeleteIncome(year, month, day, memberID)

	assert.Equal(t, nil, err)
}
