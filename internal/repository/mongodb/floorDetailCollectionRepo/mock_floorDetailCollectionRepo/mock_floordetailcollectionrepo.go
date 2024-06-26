// Code generated by MockGen. DO NOT EDIT.
// Source: floor_detail_collection_repo.go

// Package mock_floordetailcollectionrepo is a generated GoMock package.
package mock_floordetailcollectionrepo

import (
	context "context"
	reflect "reflect"

	models "github.com/RyuChk/ips-map-service/internal/models"
	mongodb "github.com/RyuChk/ips-map-service/internal/repository/mongodb"
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

// FindOne mocks base method.
func (m *MockRepository) FindOne(ctx context.Context, filter mongodb.Filter) (models.FloorDetail, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindOne", ctx, filter)
	ret0, _ := ret[0].(models.FloorDetail)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindOne indicates an expected call of FindOne.
func (mr *MockRepositoryMockRecorder) FindOne(ctx, filter interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindOne", reflect.TypeOf((*MockRepository)(nil).FindOne), ctx, filter)
}

// InsertOne mocks base method.
func (m *MockRepository) InsertOne(ctx context.Context, document models.FloorDetail) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InsertOne", ctx, document)
	ret0, _ := ret[0].(error)
	return ret0
}

// InsertOne indicates an expected call of InsertOne.
func (mr *MockRepositoryMockRecorder) InsertOne(ctx, document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InsertOne", reflect.TypeOf((*MockRepository)(nil).InsertOne), ctx, document)
}

// Upsert mocks base method.
func (m *MockRepository) Upsert(ctx context.Context, document models.FloorDetail) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", ctx, document)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockRepositoryMockRecorder) Upsert(ctx, document interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockRepository)(nil).Upsert), ctx, document)
}
