// The MIT License (MIT)

// Copyright (c) 2017-2020 Uber Technologies Inc.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

// Code generated by MockGen. DO NOT EDIT.
// Source: clientInterface.go
//
// Generated by this command:
//
//	mockgen -package dynamicconfig -source clientInterface.go -destination clientInterface_mock.go -self_package github.com/uber/cadence/common/dynamicconfig
//

// Package dynamicconfig is a generated GoMock package.
package dynamicconfig

import (
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"

	types "github.com/uber/cadence/common/types"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
	isgomock struct{}
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// GetBoolValue mocks base method.
func (m *MockClient) GetBoolValue(name BoolKey, filters map[Filter]any) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBoolValue", name, filters)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBoolValue indicates an expected call of GetBoolValue.
func (mr *MockClientMockRecorder) GetBoolValue(name, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBoolValue", reflect.TypeOf((*MockClient)(nil).GetBoolValue), name, filters)
}

// GetDurationValue mocks base method.
func (m *MockClient) GetDurationValue(name DurationKey, filters map[Filter]any) (time.Duration, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetDurationValue", name, filters)
	ret0, _ := ret[0].(time.Duration)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetDurationValue indicates an expected call of GetDurationValue.
func (mr *MockClientMockRecorder) GetDurationValue(name, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetDurationValue", reflect.TypeOf((*MockClient)(nil).GetDurationValue), name, filters)
}

// GetFloatValue mocks base method.
func (m *MockClient) GetFloatValue(name FloatKey, filters map[Filter]any) (float64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetFloatValue", name, filters)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetFloatValue indicates an expected call of GetFloatValue.
func (mr *MockClientMockRecorder) GetFloatValue(name, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetFloatValue", reflect.TypeOf((*MockClient)(nil).GetFloatValue), name, filters)
}

// GetIntValue mocks base method.
func (m *MockClient) GetIntValue(name IntKey, filters map[Filter]any) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIntValue", name, filters)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIntValue indicates an expected call of GetIntValue.
func (mr *MockClientMockRecorder) GetIntValue(name, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIntValue", reflect.TypeOf((*MockClient)(nil).GetIntValue), name, filters)
}

// GetListValue mocks base method.
func (m *MockClient) GetListValue(name ListKey, filters map[Filter]any) ([]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetListValue", name, filters)
	ret0, _ := ret[0].([]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetListValue indicates an expected call of GetListValue.
func (mr *MockClientMockRecorder) GetListValue(name, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetListValue", reflect.TypeOf((*MockClient)(nil).GetListValue), name, filters)
}

// GetMapValue mocks base method.
func (m *MockClient) GetMapValue(name MapKey, filters map[Filter]any) (map[string]any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMapValue", name, filters)
	ret0, _ := ret[0].(map[string]any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMapValue indicates an expected call of GetMapValue.
func (mr *MockClientMockRecorder) GetMapValue(name, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMapValue", reflect.TypeOf((*MockClient)(nil).GetMapValue), name, filters)
}

// GetStringValue mocks base method.
func (m *MockClient) GetStringValue(name StringKey, filters map[Filter]any) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStringValue", name, filters)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStringValue indicates an expected call of GetStringValue.
func (mr *MockClientMockRecorder) GetStringValue(name, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStringValue", reflect.TypeOf((*MockClient)(nil).GetStringValue), name, filters)
}

// GetValue mocks base method.
func (m *MockClient) GetValue(name Key) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValue", name)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValue indicates an expected call of GetValue.
func (mr *MockClientMockRecorder) GetValue(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValue", reflect.TypeOf((*MockClient)(nil).GetValue), name)
}

// GetValueWithFilters mocks base method.
func (m *MockClient) GetValueWithFilters(name Key, filters map[Filter]any) (any, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetValueWithFilters", name, filters)
	ret0, _ := ret[0].(any)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetValueWithFilters indicates an expected call of GetValueWithFilters.
func (mr *MockClientMockRecorder) GetValueWithFilters(name, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetValueWithFilters", reflect.TypeOf((*MockClient)(nil).GetValueWithFilters), name, filters)
}

// ListValue mocks base method.
func (m *MockClient) ListValue(name Key) ([]*types.DynamicConfigEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListValue", name)
	ret0, _ := ret[0].([]*types.DynamicConfigEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListValue indicates an expected call of ListValue.
func (mr *MockClientMockRecorder) ListValue(name any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListValue", reflect.TypeOf((*MockClient)(nil).ListValue), name)
}

// RestoreValue mocks base method.
func (m *MockClient) RestoreValue(name Key, filters map[Filter]any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RestoreValue", name, filters)
	ret0, _ := ret[0].(error)
	return ret0
}

// RestoreValue indicates an expected call of RestoreValue.
func (mr *MockClientMockRecorder) RestoreValue(name, filters any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RestoreValue", reflect.TypeOf((*MockClient)(nil).RestoreValue), name, filters)
}

// UpdateValue mocks base method.
func (m *MockClient) UpdateValue(name Key, value any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateValue", name, value)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateValue indicates an expected call of UpdateValue.
func (mr *MockClientMockRecorder) UpdateValue(name, value any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateValue", reflect.TypeOf((*MockClient)(nil).UpdateValue), name, value)
}
