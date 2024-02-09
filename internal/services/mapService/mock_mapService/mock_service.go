// Code generated by MockGen. DO NOT EDIT.
// Source: service.go

// Package mock_mapservice is a generated GoMock package.
package mock_mapservice

import (
	context "context"
	reflect "reflect"

	constants "github.com/RyuChk/ips-map-service/apps/constants"
	models "github.com/RyuChk/ips-map-service/apps/map/models"
	gomock "github.com/golang/mock/gomock"
)

// MockService is a mock of Service interface.
type MockService struct {
	ctrl     *gomock.Controller
	recorder *MockServiceMockRecorder
}

// MockServiceMockRecorder is the mock recorder for MockService.
type MockServiceMockRecorder struct {
	mock *MockService
}

// NewMockService creates a new mock instance.
func NewMockService(ctrl *gomock.Controller) *MockService {
	mock := &MockService{ctrl: ctrl}
	mock.recorder = &MockServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockService) EXPECT() *MockServiceMockRecorder {
	return m.recorder
}

// AddFloorToDB mocks base method.
func (m *MockService) AddFloorToDB(ctx context.Context, body models.Map) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddFloorToDB", ctx, body)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddFloorToDB indicates an expected call of AddFloorToDB.
func (mr *MockServiceMockRecorder) AddFloorToDB(ctx, body interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddFloorToDB", reflect.TypeOf((*MockService)(nil).AddFloorToDB), ctx, body)
}

// AddMapURL mocks base method.
func (m *MockService) AddMapURL(ctx context.Context, floor int, building, url string) (models.MapImageURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddMapURL", ctx, floor, building, url)
	ret0, _ := ret[0].(models.MapImageURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddMapURL indicates an expected call of AddMapURL.
func (mr *MockServiceMockRecorder) AddMapURL(ctx, floor, building, url interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddMapURL", reflect.TypeOf((*MockService)(nil).AddMapURL), ctx, floor, building, url)
}

// GetFloorListByBuilding mocks base method.
func (m *MockService) GetFloorListByBuilding(ctx context.Context, building string, role constants.UserRole) ([]models.Map, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFloorListByBuilding", ctx, building, role)
	ret0, _ := ret[0].([]models.Map)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFloorListByBuilding indicates an expected call of GetFloorListByBuilding.
func (mr *MockServiceMockRecorder) GetFloorListByBuilding(ctx, building, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloorListByBuilding", reflect.TypeOf((*MockService)(nil).GetFloorListByBuilding), ctx, building, role)
}

// GetMapURLFromKey mocks base method.
func (m *MockService) GetMapURLFromKey(ctx context.Context, floor int, building string, role constants.UserRole) (models.MapImageURL, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMapURLFromKey", ctx, floor, building, role)
	ret0, _ := ret[0].(models.MapImageURL)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMapURLFromKey indicates an expected call of GetMapURLFromKey.
func (mr *MockServiceMockRecorder) GetMapURLFromKey(ctx, floor, building, role interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMapURLFromKey", reflect.TypeOf((*MockService)(nil).GetMapURLFromKey), ctx, floor, building, role)
}
