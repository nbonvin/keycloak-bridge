// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/cloudtrust/keycloak-bridge/pkg/health (interfaces: InfluxModule,Influx)

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	health "github.com/cloudtrust/keycloak-bridge/pkg/health"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// InfluxModule is a mock of InfluxModule interface
type InfluxModule struct {
	ctrl     *gomock.Controller
	recorder *InfluxModuleMockRecorder
}

// InfluxModuleMockRecorder is the mock recorder for InfluxModule
type InfluxModuleMockRecorder struct {
	mock *InfluxModule
}

// NewInfluxModule creates a new mock instance
func NewInfluxModule(ctrl *gomock.Controller) *InfluxModule {
	mock := &InfluxModule{ctrl: ctrl}
	mock.recorder = &InfluxModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *InfluxModule) EXPECT() *InfluxModuleMockRecorder {
	return m.recorder
}

// HealthChecks mocks base method
func (m *InfluxModule) HealthChecks(arg0 context.Context) []health.InfluxReport {
	ret := m.ctrl.Call(m, "HealthChecks", arg0)
	ret0, _ := ret[0].([]health.InfluxReport)
	return ret0
}

// HealthChecks indicates an expected call of HealthChecks
func (mr *InfluxModuleMockRecorder) HealthChecks(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthChecks", reflect.TypeOf((*InfluxModule)(nil).HealthChecks), arg0)
}

// Influx is a mock of Influx interface
type Influx struct {
	ctrl     *gomock.Controller
	recorder *InfluxMockRecorder
}

// InfluxMockRecorder is the mock recorder for Influx
type InfluxMockRecorder struct {
	mock *Influx
}

// NewInflux creates a new mock instance
func NewInflux(ctrl *gomock.Controller) *Influx {
	mock := &Influx{ctrl: ctrl}
	mock.recorder = &InfluxMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *Influx) EXPECT() *InfluxMockRecorder {
	return m.recorder
}

// Ping mocks base method
func (m *Influx) Ping(arg0 time.Duration) (time.Duration, string, error) {
	ret := m.ctrl.Call(m, "Ping", arg0)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(string)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// Ping indicates an expected call of Ping
func (mr *InfluxMockRecorder) Ping(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ping", reflect.TypeOf((*Influx)(nil).Ping), arg0)
}
