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
// Source: types.go
//
// Generated by this command:
//
//	mockgen -package store -source types.go -destination mocks.go -self_package github.com/uber/cadence/common/reconciliation/store
//

// Package store is a generated GoMock package.
package store

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockScanOutputIterator is a mock of ScanOutputIterator interface.
type MockScanOutputIterator struct {
	ctrl     *gomock.Controller
	recorder *MockScanOutputIteratorMockRecorder
	isgomock struct{}
}

// MockScanOutputIteratorMockRecorder is the mock recorder for MockScanOutputIterator.
type MockScanOutputIteratorMockRecorder struct {
	mock *MockScanOutputIterator
}

// NewMockScanOutputIterator creates a new mock instance.
func NewMockScanOutputIterator(ctrl *gomock.Controller) *MockScanOutputIterator {
	mock := &MockScanOutputIterator{ctrl: ctrl}
	mock.recorder = &MockScanOutputIteratorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockScanOutputIterator) EXPECT() *MockScanOutputIteratorMockRecorder {
	return m.recorder
}

// HasNext mocks base method.
func (m *MockScanOutputIterator) HasNext() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HasNext")
	ret0, _ := ret[0].(bool)
	return ret0
}

// HasNext indicates an expected call of HasNext.
func (mr *MockScanOutputIteratorMockRecorder) HasNext() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HasNext", reflect.TypeOf((*MockScanOutputIterator)(nil).HasNext))
}

// Next mocks base method.
func (m *MockScanOutputIterator) Next() (*ScanOutputEntity, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Next")
	ret0, _ := ret[0].(*ScanOutputEntity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Next indicates an expected call of Next.
func (mr *MockScanOutputIteratorMockRecorder) Next() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Next", reflect.TypeOf((*MockScanOutputIterator)(nil).Next))
}

// MockExecutionWriter is a mock of ExecutionWriter interface.
type MockExecutionWriter struct {
	ctrl     *gomock.Controller
	recorder *MockExecutionWriterMockRecorder
	isgomock struct{}
}

// MockExecutionWriterMockRecorder is the mock recorder for MockExecutionWriter.
type MockExecutionWriterMockRecorder struct {
	mock *MockExecutionWriter
}

// NewMockExecutionWriter creates a new mock instance.
func NewMockExecutionWriter(ctrl *gomock.Controller) *MockExecutionWriter {
	mock := &MockExecutionWriter{ctrl: ctrl}
	mock.recorder = &MockExecutionWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExecutionWriter) EXPECT() *MockExecutionWriterMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockExecutionWriter) Add(arg0 any) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Add", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Add indicates an expected call of Add.
func (mr *MockExecutionWriterMockRecorder) Add(arg0 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockExecutionWriter)(nil).Add), arg0)
}

// Flush mocks base method.
func (m *MockExecutionWriter) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush.
func (mr *MockExecutionWriterMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockExecutionWriter)(nil).Flush))
}

// FlushedKeys mocks base method.
func (m *MockExecutionWriter) FlushedKeys() *Keys {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlushedKeys")
	ret0, _ := ret[0].(*Keys)
	return ret0
}

// FlushedKeys indicates an expected call of FlushedKeys.
func (mr *MockExecutionWriterMockRecorder) FlushedKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlushedKeys", reflect.TypeOf((*MockExecutionWriter)(nil).FlushedKeys))
}
