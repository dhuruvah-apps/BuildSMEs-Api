// Code generated by MockGen. DO NOT EDIT.
// Source: usecase.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	models "github.com/dhuruvah-apps/BuildSMEs-Api/internal/models"
	utils "github.com/dhuruvah-apps/BuildSMEs-Api/pkg/utils"
	gomock "github.com/golang/mock/gomock"
	uuid "github.com/google/uuid"
	reflect "reflect"
)

// MockUseCase is a mock of UseCase interface
type MockUseCase struct {
	ctrl     *gomock.Controller
	recorder *MockUseCaseMockRecorder
}

// MockUseCaseMockRecorder is the mock recorder for MockUseCase
type MockUseCaseMockRecorder struct {
	mock *MockUseCase
}

// NewMockUseCase creates a new mock instance
func NewMockUseCase(ctrl *gomock.Controller) *MockUseCase {
	mock := &MockUseCase{ctrl: ctrl}
	mock.recorder = &MockUseCaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUseCase) EXPECT() *MockUseCaseMockRecorder {
	return m.recorder
}

// Create mocks base method
func (m *MockUseCase) Create(ctx context.Context, news *models.News) (*models.News, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, news)
	ret0, _ := ret[0].(*models.News)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create
func (mr *MockUseCaseMockRecorder) Create(ctx, news interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockUseCase)(nil).Create), ctx, news)
}

// Update mocks base method
func (m *MockUseCase) Update(ctx context.Context, news *models.News) (*models.News, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", ctx, news)
	ret0, _ := ret[0].(*models.News)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update
func (mr *MockUseCaseMockRecorder) Update(ctx, news interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockUseCase)(nil).Update), ctx, news)
}

// GetNewsByID mocks base method
func (m *MockUseCase) GetNewsByID(ctx context.Context, newsID uuid.UUID) (*models.NewsBase, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewsByID", ctx, newsID)
	ret0, _ := ret[0].(*models.NewsBase)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNewsByID indicates an expected call of GetNewsByID
func (mr *MockUseCaseMockRecorder) GetNewsByID(ctx, newsID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewsByID", reflect.TypeOf((*MockUseCase)(nil).GetNewsByID), ctx, newsID)
}

// Delete mocks base method
func (m *MockUseCase) Delete(ctx context.Context, newsID uuid.UUID) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", ctx, newsID)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete
func (mr *MockUseCaseMockRecorder) Delete(ctx, newsID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockUseCase)(nil).Delete), ctx, newsID)
}

// GetNews mocks base method
func (m *MockUseCase) GetNews(ctx context.Context, pq *utils.PaginationQuery) (*models.NewsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNews", ctx, pq)
	ret0, _ := ret[0].(*models.NewsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNews indicates an expected call of GetNews
func (mr *MockUseCaseMockRecorder) GetNews(ctx, pq interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNews", reflect.TypeOf((*MockUseCase)(nil).GetNews), ctx, pq)
}

// SearchByTitle mocks base method
func (m *MockUseCase) SearchByTitle(ctx context.Context, title string, query *utils.PaginationQuery) (*models.NewsList, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchByTitle", ctx, title, query)
	ret0, _ := ret[0].(*models.NewsList)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchByTitle indicates an expected call of SearchByTitle
func (mr *MockUseCaseMockRecorder) SearchByTitle(ctx, title, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchByTitle", reflect.TypeOf((*MockUseCase)(nil).SearchByTitle), ctx, title, query)
}
