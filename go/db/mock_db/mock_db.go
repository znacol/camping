// Code generated by MockGen. DO NOT EDIT.
// Source: go/db/db.go

// Package mock_db is a generated GoMock package.
package mock_db

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	proto "github.com/znacol/camping/go/proto"
	reflect "reflect"
)

// MockAPI is a mock of API interface
type MockAPI struct {
	ctrl     *gomock.Controller
	recorder *MockAPIMockRecorder
}

// MockAPIMockRecorder is the mock recorder for MockAPI
type MockAPIMockRecorder struct {
	mock *MockAPI
}

// NewMockAPI creates a new mock instance
func NewMockAPI(ctrl *gomock.Controller) *MockAPI {
	mock := &MockAPI{ctrl: ctrl}
	mock.recorder = &MockAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAPI) EXPECT() *MockAPIMockRecorder {
	return m.recorder
}

// SitesGet mocks base method
func (m *MockAPI) SitesGet(ctx context.Context, id uint64) ([]*proto.Site, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SitesGet", ctx, id)
	ret0, _ := ret[0].([]*proto.Site)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SitesGet indicates an expected call of SitesGet
func (mr *MockAPIMockRecorder) SitesGet(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SitesGet", reflect.TypeOf((*MockAPI)(nil).SitesGet), ctx, id)
}

// SiteUpsert mocks base method
func (m *MockAPI) SiteUpsert(ctx context.Context, latitude, longitude float32, nationalForestID, districtID, altitude uint64, notes string) (*proto.Site, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SiteUpsert", ctx, latitude, longitude, nationalForestID, districtID, altitude, notes)
	ret0, _ := ret[0].(*proto.Site)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SiteUpsert indicates an expected call of SiteUpsert
func (mr *MockAPIMockRecorder) SiteUpsert(ctx, latitude, longitude, nationalForestID, districtID, altitude, notes interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SiteUpsert", reflect.TypeOf((*MockAPI)(nil).SiteUpsert), ctx, latitude, longitude, nationalForestID, districtID, altitude, notes)
}

// NationalForestsGet mocks base method
func (m *MockAPI) NationalForestsGet(ctx context.Context, id uint64) ([]*proto.NationalForest, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NationalForestsGet", ctx, id)
	ret0, _ := ret[0].([]*proto.NationalForest)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NationalForestsGet indicates an expected call of NationalForestsGet
func (mr *MockAPIMockRecorder) NationalForestsGet(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NationalForestsGet", reflect.TypeOf((*MockAPI)(nil).NationalForestsGet), ctx, id)
}

// DistrictsGet mocks base method
func (m *MockAPI) DistrictsGet(ctx context.Context, id uint64) ([]*proto.District, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DistrictsGet", ctx, id)
	ret0, _ := ret[0].([]*proto.District)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DistrictsGet indicates an expected call of DistrictsGet
func (mr *MockAPIMockRecorder) DistrictsGet(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DistrictsGet", reflect.TypeOf((*MockAPI)(nil).DistrictsGet), ctx, id)
}