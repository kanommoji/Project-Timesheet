package repository

import (
	"database/sql"
	"timesheet/internal/model"
)

type TimesheetRepositoryGateways interface {
	GetSummary(year, month int) ([]model.TransactionTimesheet, error)
	GetMemberByID(memberID string) ([]model.Member, error)
	GetIncomes(memberID string, year, month int) ([]model.Incomes, error)
	UpdateIncomeByID(year, month int, memberID string, income model.Incomes) error
}

type TimesheetRepository struct {
	DatabaseConnection *sql.DB
}

func (repository TimesheetRepository) GetSummary(year, month int) ([]model.TransactionTimesheet, error) {
	var transactionTimesheetList []model.TransactionTimesheet
	var transactionTimesheet model.TransactionTimesheet
	statement, err := repository.DatabaseConnection.Prepare(`SELECT * FROM timesheet.transactions WHERE year = ? AND month = ?`)
	if err != nil {
		return transactionTimesheetList, err
	}
	row, err := statement.Query(year, month)
	if err != nil {
		return transactionTimesheetList, err
	}
	for row.Next() {
		err = row.Scan(
			&transactionTimesheet.ID,
			&transactionTimesheet.MemberID,
			&transactionTimesheet.Month,
			&transactionTimesheet.Year,
			&transactionTimesheet.Company,
			&transactionTimesheet.MemberNameTH,
			&transactionTimesheet.Coaching,
			&transactionTimesheet.Training,
			&transactionTimesheet.Other,
			&transactionTimesheet.TotalIncomes,
			&transactionTimesheet.Salary,
			&transactionTimesheet.IncomeTax1,
			&transactionTimesheet.SocialSecurity,
			&transactionTimesheet.NetSalary,
			&transactionTimesheet.Wage,
			&transactionTimesheet.IncomeTax53Percentage,
			&transactionTimesheet.IncomeTax53,
			&transactionTimesheet.NetWage,
			&transactionTimesheet.NetTransfer,
			&transactionTimesheet.StatusCheckingTransfer,
			&transactionTimesheet.DateTransfer,
			&transactionTimesheet.Comment,
		)
		transactionTimesheetList = append(transactionTimesheetList, transactionTimesheet)
	}

	return transactionTimesheetList, nil
}

func (repository TimesheetRepository) UpdateIncomeByID(year, month int, memberID string, income model.Incomes) error {
	statement, err := repository.DatabaseConnection.Prepare(`INSERT INTO incomes (member_id, month, year, day, start_time_am_hours, start_time_am_minutes, start_time_am_seconds, end_time_am_hours, end_time_am_minutes, end_time_am_seconds, start_time_pm_hours, start_time_pm_minutes, start_time_pm_seconds, end_time_pm_hours, end_time_pm_minutes, end_time_pm_seconds, overtime, total_hours_hours, total_hours_minutes, total_hours_seconds, coaching_customer_charging, coaching_payment_rate, training_wage, other_wage, company, description) VALUES ( ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ? , ?, ? , ? , ?, ? , ? , ?, ? , ? , ?, ? , ? , ?, ? )`)
	if err != nil {
		return err
	}
	_, err = statement.Exec(memberID, month, year, income.Day, income.StartTimeAMHours, income.StartTimeAMMinutes, income.StartTimeAMSeconds, income.EndTimeAMHours, income.EndTimeAMMinutes, income.EndTimeAMSeconds, income.StartTimePMHours, income.StartTimePMMinutes, income.StartTimePMSeconds, income.EndTimePMHours, income.EndTimePMMinutes, income.EndTimePMSeconds, income.Overtime, income.TotalHoursHours, income.TotalHoursMinutes, income.TotalHoursSeconds, income.CoachingCustomerCharging, income.CoachingPaymentRate, income.TrainingWage, income.OtherWage, income.Company, income.Description)
	if err != nil {
		return err
	}
	return nil
}

func (repository TimesheetRepository) GetMemberByID(memberID string) ([]model.Member, error) {
	var memberList []model.Member
	var member model.Member
	statement, err := repository.DatabaseConnection.Prepare(`SELECT * FROM timesheet.members WHERE member_id = ?`)
	if err != nil {
		return memberList, err
	}
	row, err := statement.Query(memberID)
	if err != nil {
		return memberList, err
	}
	for row.Next() {
		err = row.Scan(
			&member.ID,
			&member.MemberID,
			&member.Company,
			&member.MemberNameTH,
			&member.MemberNameENG,
			&member.Email,
			&member.OvertimeRate,
			&member.RatePerDay,
			&member.RatePerHour,
			&member.Salary,
			&member.IncomeTax1,
			&member.IncomeTax53Percentage,
			&member.SocialSecurity,
			&member.Status,
			&member.TravelExpense,
		)
		memberList = append(memberList, member)
	}
	return memberList, nil
}

func (repository TimesheetRepository) GetIncomes(memberID string, year, month int) ([]model.Incomes, error) {
	var incomeList []model.Incomes
	var income model.Incomes
	statement, err := repository.DatabaseConnection.Prepare(`SELECT * FROM timesheet.incomes WHERE member_id = ? AND year = ? AND month = ?`)
	if err != nil {
		return incomeList, err
	}
	row, err := statement.Query(memberID, year, month)
	if err != nil {
		return incomeList, err
	}
	for row.Next() {
		err = row.Scan(
			&income.ID,
			&income.MemberID,
			&income.Month,
			&income.Year,
			&income.Day,
			&income.StartTimeAMHours,
			&income.StartTimeAMMinutes,
			&income.StartTimeAMSeconds,
			&income.EndTimeAMHours,
			&income.EndTimeAMMinutes,
			&income.EndTimeAMSeconds,
			&income.StartTimePMHours,
			&income.StartTimePMMinutes,
			&income.StartTimePMSeconds,
			&income.EndTimePMHours,
			&income.EndTimePMMinutes,
			&income.EndTimePMSeconds,
			&income.Overtime,
			&income.TotalHoursHours,
			&income.TotalHoursMinutes,
			&income.TotalHoursSeconds,
			&income.CoachingCustomerCharging,
			&income.CoachingPaymentRate,
			&income.TrainingWage,
			&income.OtherWage,
			&income.Company,
			&income.Description,
		)
		incomeList = append(incomeList, income)
	}
	return incomeList, nil
}

func (repository TimesheetRepository) UpdateTransactionTimsheet(transactionTimesheet []model.TransactionTimesheet) error {
	for index := range transactionTimesheet {
		statement, err := repository.DatabaseConnection.Prepare(`INSERT INTO transactions (member_id, month, year, company, member_name_th, coaching, training, other, total_incomes, salary, income_tax_1, social_security, net_salary, wage, income_tax_53_percentage, income_tax_53, net_wage, net_transfer, status_checking_transfer, date_transfer, comment) VALUES ( ? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ? ,? , ? , ? )`)
		if err != nil {
			return err
		}
		_, err = statement.Exec(
			transactionTimesheet[index].MemberID,
			transactionTimesheet[index].Month,
			transactionTimesheet[index].Year,
			transactionTimesheet[index].Company,
			transactionTimesheet[index].MemberNameTH,
			transactionTimesheet[index].Coaching,
			transactionTimesheet[index].Training,
			transactionTimesheet[index].Other,
			transactionTimesheet[index].TotalIncomes,
			transactionTimesheet[index].Salary,
			transactionTimesheet[index].IncomeTax1,
			transactionTimesheet[index].SocialSecurity,
			transactionTimesheet[index].NetSalary,
			transactionTimesheet[index].Wage,
			transactionTimesheet[index].IncomeTax53Percentage,
			transactionTimesheet[index].IncomeTax53,
			transactionTimesheet[index].NetWage,
			transactionTimesheet[index].NetTransfer,
			transactionTimesheet[index].StatusCheckingTransfer,
			transactionTimesheet[index].DateTransfer,
			transactionTimesheet[index].Comment)
		if err != nil {
			return err
		}
	}
	return nil
}