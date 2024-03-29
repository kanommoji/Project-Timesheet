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
	income := model.Incomes{
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

	err := repository.CreateIncome(year, month, memberID, income)

	assert.Equal(t, nil, err)
}

func Test_GetMemberListByMemberID_Input_MemberID_001_Should_Be_MemberList(t *testing.T) {
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
			Picture:               "",
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
			Picture:               "",
		},
	}
	memberID := "001"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetMemberListByMemberID(memberID)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_GetIncomes_Input_MemberID_006_Year_2019_Month_12_Should_Be_IncomeList(t *testing.T) {
	startTimeAMDay11, _ := time.Parse("2006-01-02 15:04:05", "2018-12-11 09:00:00")
	endTimeAMDay11, _ := time.Parse("2006-01-02 15:04:05", "2018-12-11 12:00:00")
	startTimePMDay11, _ := time.Parse("2006-01-02 15:04:05", "2018-12-11 13:00:00")
	endTimePMDay11, _ := time.Parse("2006-01-02 15:04:05", "2018-12-11 18:00:00")
	totalHoursDay11, _ := time.Parse("2006-01-02 15:04:05", "2018-12-11 08:00:00")

	startTimeAMDay12, _ := time.Parse("2006-01-02 15:04:05", "2018-12-12 09:00:00")
	endTimeAMDay12, _ := time.Parse("2006-01-02 15:04:05", "2018-12-12 12:00:00")
	startTimePMDay12, _ := time.Parse("2006-01-02 15:04:05", "2018-12-12 13:00:00")
	endTimePMDay12, _ := time.Parse("2006-01-02 15:04:05", "2018-12-12 18:00:00")
	totalHoursDay12, _ := time.Parse("2006-01-02 15:04:05", "2018-12-12 08:00:00")
	expected := []model.Incomes{
		{
			ID:                       58,
			MemberID:                 "006",
			Month:                    12,
			Year:                     2019,
			Day:                      11,
			StartTimeAM:              startTimeAMDay11,
			EndTimeAM:                endTimeAMDay11,
			StartTimePM:              startTimePMDay11,
			EndTimePM:                endTimePMDay11,
			Overtime:                 0,
			TotalHours:               totalHoursDay11,
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
			StartTimeAM:              startTimeAMDay12,
			EndTimeAM:                endTimeAMDay12,
			StartTimePM:              startTimePMDay12,
			EndTimePM:                endTimePMDay12,
			Overtime:                 0,
			TotalHours:               totalHoursDay12,
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

func Test_VerifyTransactionTimesheet_Input_Transaction_MemberID_001_Should_Be_Create_TransactionTimesheet_And_Update_TransactionTimesheet(t *testing.T) {
	transactionTimesheetList := []model.TransactionTimesheet{
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

	err := repository.VerifyTransactionTimesheet(transactionTimesheetList)

	assert.Equal(t, nil, err)
}

func Test_CreateTransactionTimesheet_Input_TransactionID_006201912shuhari_TransactionTimesheet_MemberID_006_Should_Be_No_Error(t *testing.T) {
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

	err := repository.CreateTransactionTimesheet(transactionTimesheet, transactionID)

	assert.Equal(t, nil, err)
}

func Test_UpdateTransactionTimesheet_Input_TransactionID_001201911shuhari_TransactionTimesheet_MemberID_001_Should_Be_No_Error(t *testing.T) {
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

	err := repository.UpdateTransactionTimesheet(transactionTimesheet, transactionID)

	assert.Equal(t, nil, err)
}

func Test_CreateTimesheet_Input_MemberID_006_Month_12_Year_2019_Should_Be_No_Error(t *testing.T) {
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

func Test_UpdateTimesheet_Input_Timesheet_MemberID_007_Year_2019_Month_12_Should_Be_No_Error(t *testing.T) {
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

func Test_GetTimesheet_Input_MemberID_003_Month_12_Year_2017_Should_Be_Timesheet(t *testing.T) {
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

func Test_DeleteIncome_Input_IncomeID_47_Should_Be_No_Error(t *testing.T) {
	incomeID := 47
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.DeleteIncome(incomeID)

	assert.Equal(t, nil, err)
}

func Test_UpdateMemberDetails_Input_Member_Should_Be_No_Error(t *testing.T) {
	memberDetails := model.Member{
		ID:                    10,
		MemberNameTH:          "ภาณุมาศ แสนโท",
		MemberNameENG:         "Panumars Seanto",
		Email:                 "panumars@scrum123.com",
		OvertimeRate:          0.00,
		RatePerDay:            15000.00,
		RatePerHour:           1875.00,
		Salary:                25000.00,
		IncomeTax1:            0.00,
		SocialSecurity:        750.00,
		IncomeTax53Percentage: 5,
		Status:                "wage",
		TravelExpense:         0.00,
	}
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdateMemberDetails(memberDetails)

	assert.Equal(t, nil, err)
}

func Test_GetMemberIDByEmail_Input_Email_prathan_scrum123_com_Should_Be_001(t *testing.T) {
	expected := "001"
	email := "prathan@scrum123.com"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetMemberIDByEmail(email)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_GetMemberIDByEmail_Input_Email_somkiat_scrum123_com_Should_Be_003(t *testing.T) {
	expected := "003"
	email := "somkiat@scrum123.com"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetMemberIDByEmail(email)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_GetMemberIDByEmail_Input_Email_logintest535_gmail_com_Should_Be_007(t *testing.T) {
	expected := "007"
	email := "logintest535@gmail.com"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetMemberIDByEmail(email)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_UpdatePictureToMembers_Input_Email_prathan_scrum123_com_And_Picture_Should_Be_No_Error(t *testing.T) {
	email := "prathan@scrum123.com"
	picture := "https://lh4.googleusercontent.com/-nA86bkk5Icc/AAAAAAAAAAI/AAAAAAAAAAA/Wixwdu9UCfU/photo.jpg"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	err := repository.UpdatePictureToMembers(picture, email)

	assert.Equal(t, nil, err)
}

func Test_GetTransactionTimesheets_Input_MemberID_001_Year_2017_Should_Be_TransactionTimesheetList(t *testing.T) {
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
	memberID := "001"
	year := 2017
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet?parseTime=true")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetTransactionTimesheets(memberID, year)

	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}

func Test_GetProfileByEmail_Input_Email_nareenart_scrum123_com_Should_Be_MemberID_002_And_Picture(t *testing.T) {
	expected := model.Profile{
		MemberID: "002",
		Email:    "nareenart@scrum123.com",
		Picture:  "https://lh4.googleusercontent.com/-nA86bkk5Icc/AAAAAAAAAAI/AAAAAAAAAAA/Wixwdu9UCfU/photo.jpg",
	}
	email := "nareenart@scrum123.com"
	databaseConnection, _ := sqlx.Connect("mysql", "root:root@tcp(localhost:3306)/timesheet")
	defer databaseConnection.Close()
	repository := TimesheetRepository{
		DatabaseConnection: databaseConnection,
	}

	actual, err := repository.GetProfileByEmail(email)
	assert.Equal(t, nil, err)
	assert.Equal(t, expected, actual)
}
