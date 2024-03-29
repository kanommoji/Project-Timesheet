package mockapi

import (
	"timesheet/internal/model"

	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

func (mock *MockRepository) GetSummary(year, month int) ([]model.TransactionTimesheet, error) {
	argument := mock.Called(year, month)
	return argument.Get(0).([]model.TransactionTimesheet), argument.Error(1)
}

func (mock *MockRepository) CreateIncome(year, month int, memberID string, income model.Incomes) error {
	argument := mock.Called(year, month, memberID, income)
	return argument.Error(0)
}

func (mock *MockRepository) VerifyTransactionTimesheet(transactionTimesheetList []model.TransactionTimesheet) error {
	argument := mock.Called(transactionTimesheetList)
	return argument.Error(0)
}

func (mock *MockRepository) UpdateTimesheet(timesheet model.Timesheet, memberID string, year, month int) error {
	argument := mock.Called(timesheet, memberID, year, month)
	return argument.Error(0)
}

func (mock *MockRepository) UpdateStatusTransfer(transactionID, status, date, comment string) error {
	argument := mock.Called(transactionID, status, date, comment)
	return argument.Error(0)
}

func (mock *MockRepository) DeleteIncome(incomeID int) error {
	argument := mock.Called(incomeID)
	return argument.Error(0)
}

func (mock *MockRepository) UpdateMemberDetails(memberDetails model.Member) error {
	argument := mock.Called(memberDetails)
	return argument.Error(0)
}

func (mock *MockRepository) CreateAuthentication(userInfo model.UserInfo, token model.Token) error {
	argument := mock.Called(userInfo, token)
	return argument.Error(0)
}

func (mock *MockRepository) GetProfileByAccessToken(accessToken string) (model.Profile, error) {
	argument := mock.Called(accessToken)
	return argument.Get(0).(model.Profile), argument.Error(1)
}

func (mock *MockRepository) DeleteAuthentication(accessToken string) error {
	argument := mock.Called(accessToken)
	return argument.Error(0)
}

func (mock *MockRepository) UpdatePictureToMembers(picture, email string) error {
	argument := mock.Called(picture, email)
	return argument.Error(0)
}

func (mock *MockRepository) GetProfileByEmail(email string) (model.Profile, error) {
	argument := mock.Called(email)
	return argument.Get(0).(model.Profile), argument.Error(1)
}

type MockRepositoryToTimesheet struct {
	mock.Mock
}

func (mock *MockRepositoryToTimesheet) GetIncomes(memberID string, year, month int) ([]model.Incomes, error) {
	argument := mock.Called(memberID, year, month)
	return argument.Get(0).([]model.Incomes), argument.Error(1)
}

func (mock *MockRepositoryToTimesheet) GetMemberListByMemberID(memberID string) ([]model.Member, error) {
	argument := mock.Called(memberID)
	return argument.Get(0).([]model.Member), argument.Error(1)
}

func (mock *MockRepositoryToTimesheet) CreateTimesheet(memberID string, year int, month int) error {
	argument := mock.Called(memberID, year, month)
	return argument.Error(0)
}

func (mock *MockRepositoryToTimesheet) GetTimesheet(memberID string, year, month int) (model.Timesheet, error) {
	argument := mock.Called(memberID, year, month)
	return argument.Get(0).(model.Timesheet), argument.Error(1)
}

func (mock *MockRepositoryToTimesheet) GetTransactionTimesheets(memberID string, year int) ([]model.TransactionTimesheet, error) {
	argument := mock.Called(memberID, year)
	return argument.Get(0).([]model.TransactionTimesheet), argument.Error(1)
}

func (mock *MockRepositoryToTimesheet) GetMemberIDByEmail(email string) (string, error) {
	argument := mock.Called(email)
	return argument.String(0), argument.Error(1)
}

type MockTimesheet struct {
	mock.Mock
}

func (mock *MockTimesheet) CalculatePayment(incomes []model.Incomes) model.Timesheet {
	argument := mock.Called(incomes)
	return argument.Get(0).(model.Timesheet)
}

func (mock *MockTimesheet) CalculatePaymentSummary(member []model.Member, incomes []model.Incomes, year, month int) []model.TransactionTimesheet {
	argument := mock.Called(member, incomes, year, month)
	return argument.Get(0).([]model.TransactionTimesheet)
}

func (mock *MockTimesheet) GetSummaryByID(memberID string, year, month int) (model.SummaryTimesheet, error) {
	argument := mock.Called(memberID, year, month)
	return argument.Get(0).(model.SummaryTimesheet), argument.Error(1)
}

func (mock *MockTimesheet) VerifyAuthentication(email string, expiry float64, memberID string) string {
	argument := mock.Called(email, expiry, memberID)
	return argument.String(0)
}

func (mock *MockTimesheet) GetSummaryInYearByID(memberID string, year int) (model.SummaryTransactionTimesheet, error) {
	argument := mock.Called(memberID, year)
	return argument.Get(0).(model.SummaryTransactionTimesheet), argument.Error(1)
}
