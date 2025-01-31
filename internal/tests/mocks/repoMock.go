// Code generated by MockGen. DO NOT EDIT.
// Source: ./internal/repositories/repository.go

// Package mock_repositories is a generated GoMock package.
package mock_repositories

import (
	context "context"
	reflect "reflect"

	entity "github.com/MaksKazantsev/DriverGO/internal/entity"
	models "github.com/MaksKazantsev/DriverGO/internal/service/models"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// AboutMe mocks base method.
func (m *MockRepository) AboutMe(ctx context.Context, userID string) (entity.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AboutMe", ctx, userID)
	ret0, _ := ret[0].(entity.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AboutMe indicates an expected call of AboutMe.
func (mr *MockRepositoryMockRecorder) AboutMe(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AboutMe", reflect.TypeOf((*MockRepository)(nil).AboutMe), ctx, userID)
}

// AddCar mocks base method.
func (m *MockRepository) AddCar(ctx context.Context, car entity.Car) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCar", ctx, car)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCar indicates an expected call of AddCar.
func (mr *MockRepositoryMockRecorder) AddCar(ctx, car interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCar", reflect.TypeOf((*MockRepository)(nil).AddCar), ctx, car)
}

// EditCar mocks base method.
func (m *MockRepository) EditCar(ctx context.Context, data models.CarReq, carID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditCar", ctx, data, carID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditCar indicates an expected call of EditCar.
func (mr *MockRepositoryMockRecorder) EditCar(ctx, data, carID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditCar", reflect.TypeOf((*MockRepository)(nil).EditCar), ctx, data, carID)
}

// FinishRent mocks base method.
func (m *MockRepository) FinishRent(ctx context.Context, userID, rentID string) (entity.Bill, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishRent", ctx, userID, rentID)
	ret0, _ := ret[0].(entity.Bill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FinishRent indicates an expected call of FinishRent.
func (mr *MockRepositoryMockRecorder) FinishRent(ctx, userID, rentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishRent", reflect.TypeOf((*MockRepository)(nil).FinishRent), ctx, userID, rentID)
}

// GetAvailableCars mocks base method.
func (m *MockRepository) GetAvailableCars(ctx context.Context) ([]entity.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableCars", ctx)
	ret0, _ := ret[0].([]entity.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableCars indicates an expected call of GetAvailableCars.
func (mr *MockRepositoryMockRecorder) GetAvailableCars(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableCars", reflect.TypeOf((*MockRepository)(nil).GetAvailableCars), ctx)
}

// GetFBToken mocks base method.
func (m *MockRepository) GetFBToken(ctx context.Context, userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFBToken", ctx, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFBToken indicates an expected call of GetFBToken.
func (mr *MockRepositoryMockRecorder) GetFBToken(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFBToken", reflect.TypeOf((*MockRepository)(nil).GetFBToken), ctx, userID)
}

// GetNotifications mocks base method.
func (m *MockRepository) GetNotifications(ctx context.Context, userID string) ([]entity.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotifications", ctx, userID)
	ret0, _ := ret[0].([]entity.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotifications indicates an expected call of GetNotifications.
func (mr *MockRepositoryMockRecorder) GetNotifications(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotifications", reflect.TypeOf((*MockRepository)(nil).GetNotifications), ctx, userID)
}

// GetPasswordAndID mocks base method.
func (m *MockRepository) GetPasswordAndID(ctx context.Context, email string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPasswordAndID", ctx, email)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPasswordAndID indicates an expected call of GetPasswordAndID.
func (mr *MockRepositoryMockRecorder) GetPasswordAndID(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPasswordAndID", reflect.TypeOf((*MockRepository)(nil).GetPasswordAndID), ctx, email)
}

// GetProfile mocks base method.
func (m *MockRepository) GetProfile(ctx context.Context, userID string) (entity.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", ctx, userID)
	ret0, _ := ret[0].(entity.UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockRepositoryMockRecorder) GetProfile(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockRepository)(nil).GetProfile), ctx, userID)
}

// GetRentHistory mocks base method.
func (m *MockRepository) GetRentHistory(ctx context.Context, userID string) ([]entity.RentHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRentHistory", ctx, userID)
	ret0, _ := ret[0].([]entity.RentHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRentHistory indicates an expected call of GetRentHistory.
func (mr *MockRepositoryMockRecorder) GetRentHistory(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRentHistory", reflect.TypeOf((*MockRepository)(nil).GetRentHistory), ctx, userID)
}

// Login mocks base method.
func (m *MockRepository) Login(ctx context.Context, data models.LoginReq) (models.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, data)
	ret0, _ := ret[0].(models.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockRepositoryMockRecorder) Login(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockRepository)(nil).Login), ctx, data)
}

// Refresh mocks base method.
func (m *MockRepository) Refresh(ctx context.Context, uuid, token string) (models.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", ctx, uuid, token)
	ret0, _ := ret[0].(models.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh.
func (mr *MockRepositoryMockRecorder) Refresh(ctx, uuid, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockRepository)(nil).Refresh), ctx, uuid, token)
}

// Register mocks base method.
func (m *MockRepository) Register(ctx context.Context, data entity.User) (models.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, data)
	ret0, _ := ret[0].(models.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockRepositoryMockRecorder) Register(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockRepository)(nil).Register), ctx, data)
}

// RemoveCar mocks base method.
func (m *MockRepository) RemoveCar(ctx context.Context, carID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCar", ctx, carID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCar indicates an expected call of RemoveCar.
func (mr *MockRepositoryMockRecorder) RemoveCar(ctx, carID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCar", reflect.TypeOf((*MockRepository)(nil).RemoveCar), ctx, carID)
}

// SaveNotification mocks base method.
func (m *MockRepository) SaveNotification(ctx context.Context, notification entity.Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNotification", ctx, notification)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNotification indicates an expected call of SaveNotification.
func (mr *MockRepositoryMockRecorder) SaveNotification(ctx, notification interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNotification", reflect.TypeOf((*MockRepository)(nil).SaveNotification), ctx, notification)
}

// StartRent mocks base method.
func (m *MockRepository) StartRent(ctx context.Context, userID, carID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartRent", ctx, userID, carID)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartRent indicates an expected call of StartRent.
func (mr *MockRepositoryMockRecorder) StartRent(ctx, userID, carID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartRent", reflect.TypeOf((*MockRepository)(nil).StartRent), ctx, userID, carID)
}

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// GetPasswordAndID mocks base method.
func (m *MockAuth) GetPasswordAndID(ctx context.Context, email string) (string, string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPasswordAndID", ctx, email)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPasswordAndID indicates an expected call of GetPasswordAndID.
func (mr *MockAuthMockRecorder) GetPasswordAndID(ctx, email interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPasswordAndID", reflect.TypeOf((*MockAuth)(nil).GetPasswordAndID), ctx, email)
}

// Login mocks base method.
func (m *MockAuth) Login(ctx context.Context, data models.LoginReq) (models.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Login", ctx, data)
	ret0, _ := ret[0].(models.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Login indicates an expected call of Login.
func (mr *MockAuthMockRecorder) Login(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Login", reflect.TypeOf((*MockAuth)(nil).Login), ctx, data)
}

// Refresh mocks base method.
func (m *MockAuth) Refresh(ctx context.Context, uuid, token string) (models.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Refresh", ctx, uuid, token)
	ret0, _ := ret[0].(models.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Refresh indicates an expected call of Refresh.
func (mr *MockAuthMockRecorder) Refresh(ctx, uuid, token interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Refresh", reflect.TypeOf((*MockAuth)(nil).Refresh), ctx, uuid, token)
}

// Register mocks base method.
func (m *MockAuth) Register(ctx context.Context, data entity.User) (models.AuthResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", ctx, data)
	ret0, _ := ret[0].(models.AuthResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Register indicates an expected call of Register.
func (mr *MockAuthMockRecorder) Register(ctx, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockAuth)(nil).Register), ctx, data)
}

// MockRent is a mock of Rent interface.
type MockRent struct {
	ctrl     *gomock.Controller
	recorder *MockRentMockRecorder
}

// MockRentMockRecorder is the mock recorder for MockRent.
type MockRentMockRecorder struct {
	mock *MockRent
}

// NewMockRent creates a new mock instance.
func NewMockRent(ctrl *gomock.Controller) *MockRent {
	mock := &MockRent{ctrl: ctrl}
	mock.recorder = &MockRentMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRent) EXPECT() *MockRentMockRecorder {
	return m.recorder
}

// FinishRent mocks base method.
func (m *MockRent) FinishRent(ctx context.Context, userID, rentID string) (entity.Bill, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FinishRent", ctx, userID, rentID)
	ret0, _ := ret[0].(entity.Bill)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FinishRent indicates an expected call of FinishRent.
func (mr *MockRentMockRecorder) FinishRent(ctx, userID, rentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FinishRent", reflect.TypeOf((*MockRent)(nil).FinishRent), ctx, userID, rentID)
}

// GetAvailableCars mocks base method.
func (m *MockRent) GetAvailableCars(ctx context.Context) ([]entity.Car, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAvailableCars", ctx)
	ret0, _ := ret[0].([]entity.Car)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAvailableCars indicates an expected call of GetAvailableCars.
func (mr *MockRentMockRecorder) GetAvailableCars(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAvailableCars", reflect.TypeOf((*MockRent)(nil).GetAvailableCars), ctx)
}

// GetRentHistory mocks base method.
func (m *MockRent) GetRentHistory(ctx context.Context, userID string) ([]entity.RentHistory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRentHistory", ctx, userID)
	ret0, _ := ret[0].([]entity.RentHistory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRentHistory indicates an expected call of GetRentHistory.
func (mr *MockRentMockRecorder) GetRentHistory(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRentHistory", reflect.TypeOf((*MockRent)(nil).GetRentHistory), ctx, userID)
}

// StartRent mocks base method.
func (m *MockRent) StartRent(ctx context.Context, userID, carID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartRent", ctx, userID, carID)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartRent indicates an expected call of StartRent.
func (mr *MockRentMockRecorder) StartRent(ctx, userID, carID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartRent", reflect.TypeOf((*MockRent)(nil).StartRent), ctx, userID, carID)
}

// MockCarManagement is a mock of CarManagement interface.
type MockCarManagement struct {
	ctrl     *gomock.Controller
	recorder *MockCarManagementMockRecorder
}

// MockCarManagementMockRecorder is the mock recorder for MockCarManagement.
type MockCarManagementMockRecorder struct {
	mock *MockCarManagement
}

// NewMockCarManagement creates a new mock instance.
func NewMockCarManagement(ctrl *gomock.Controller) *MockCarManagement {
	mock := &MockCarManagement{ctrl: ctrl}
	mock.recorder = &MockCarManagementMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCarManagement) EXPECT() *MockCarManagementMockRecorder {
	return m.recorder
}

// AddCar mocks base method.
func (m *MockCarManagement) AddCar(ctx context.Context, car entity.Car) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddCar", ctx, car)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddCar indicates an expected call of AddCar.
func (mr *MockCarManagementMockRecorder) AddCar(ctx, car interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddCar", reflect.TypeOf((*MockCarManagement)(nil).AddCar), ctx, car)
}

// EditCar mocks base method.
func (m *MockCarManagement) EditCar(ctx context.Context, data models.CarReq, carID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EditCar", ctx, data, carID)
	ret0, _ := ret[0].(error)
	return ret0
}

// EditCar indicates an expected call of EditCar.
func (mr *MockCarManagementMockRecorder) EditCar(ctx, data, carID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EditCar", reflect.TypeOf((*MockCarManagement)(nil).EditCar), ctx, data, carID)
}

// RemoveCar mocks base method.
func (m *MockCarManagement) RemoveCar(ctx context.Context, carID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveCar", ctx, carID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveCar indicates an expected call of RemoveCar.
func (mr *MockCarManagementMockRecorder) RemoveCar(ctx, carID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveCar", reflect.TypeOf((*MockCarManagement)(nil).RemoveCar), ctx, carID)
}

// MockUser is a mock of User interface.
type MockUser struct {
	ctrl     *gomock.Controller
	recorder *MockUserMockRecorder
}

// MockUserMockRecorder is the mock recorder for MockUser.
type MockUserMockRecorder struct {
	mock *MockUser
}

// NewMockUser creates a new mock instance.
func NewMockUser(ctrl *gomock.Controller) *MockUser {
	mock := &MockUser{ctrl: ctrl}
	mock.recorder = &MockUserMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUser) EXPECT() *MockUserMockRecorder {
	return m.recorder
}

// AboutMe mocks base method.
func (m *MockUser) AboutMe(ctx context.Context, userID string) (entity.UserInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AboutMe", ctx, userID)
	ret0, _ := ret[0].(entity.UserInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AboutMe indicates an expected call of AboutMe.
func (mr *MockUserMockRecorder) AboutMe(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AboutMe", reflect.TypeOf((*MockUser)(nil).AboutMe), ctx, userID)
}

// GetNotifications mocks base method.
func (m *MockUser) GetNotifications(ctx context.Context, userID string) ([]entity.Notification, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotifications", ctx, userID)
	ret0, _ := ret[0].([]entity.Notification)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotifications indicates an expected call of GetNotifications.
func (mr *MockUserMockRecorder) GetNotifications(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotifications", reflect.TypeOf((*MockUser)(nil).GetNotifications), ctx, userID)
}

// GetProfile mocks base method.
func (m *MockUser) GetProfile(ctx context.Context, userID string) (entity.UserProfile, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProfile", ctx, userID)
	ret0, _ := ret[0].(entity.UserProfile)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProfile indicates an expected call of GetProfile.
func (mr *MockUserMockRecorder) GetProfile(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProfile", reflect.TypeOf((*MockUser)(nil).GetProfile), ctx, userID)
}

// MockNotifierRepo is a mock of NotifierRepo interface.
type MockNotifierRepo struct {
	ctrl     *gomock.Controller
	recorder *MockNotifierRepoMockRecorder
}

// MockNotifierRepoMockRecorder is the mock recorder for MockNotifierRepo.
type MockNotifierRepoMockRecorder struct {
	mock *MockNotifierRepo
}

// NewMockNotifierRepo creates a new mock instance.
func NewMockNotifierRepo(ctrl *gomock.Controller) *MockNotifierRepo {
	mock := &MockNotifierRepo{ctrl: ctrl}
	mock.recorder = &MockNotifierRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotifierRepo) EXPECT() *MockNotifierRepoMockRecorder {
	return m.recorder
}

// GetFBToken mocks base method.
func (m *MockNotifierRepo) GetFBToken(ctx context.Context, userID string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFBToken", ctx, userID)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFBToken indicates an expected call of GetFBToken.
func (mr *MockNotifierRepoMockRecorder) GetFBToken(ctx, userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFBToken", reflect.TypeOf((*MockNotifierRepo)(nil).GetFBToken), ctx, userID)
}

// SaveNotification mocks base method.
func (m *MockNotifierRepo) SaveNotification(ctx context.Context, notification entity.Notification) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SaveNotification", ctx, notification)
	ret0, _ := ret[0].(error)
	return ret0
}

// SaveNotification indicates an expected call of SaveNotification.
func (mr *MockNotifierRepoMockRecorder) SaveNotification(ctx, notification interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SaveNotification", reflect.TypeOf((*MockNotifierRepo)(nil).SaveNotification), ctx, notification)
}
