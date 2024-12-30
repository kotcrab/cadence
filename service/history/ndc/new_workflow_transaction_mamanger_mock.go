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
// Source: new_workflow_transaction_manager.go
//
// Generated by this command:
//
//	mockgen -package ndc -source new_workflow_transaction_manager.go -destination new_workflow_transaction_mamanger_mock.go
//

// Package ndc is a generated GoMock package.
package ndc

import (
	context "context"
	reflect "reflect"
	time "time"

	gomock "go.uber.org/mock/gomock"

	execution "github.com/uber/cadence/service/history/execution"
)

// MocktransactionManagerForNewWorkflow is a mock of transactionManagerForNewWorkflow interface.
type MocktransactionManagerForNewWorkflow struct {
	ctrl     *gomock.Controller
	recorder *MocktransactionManagerForNewWorkflowMockRecorder
	isgomock struct{}
}

// MocktransactionManagerForNewWorkflowMockRecorder is the mock recorder for MocktransactionManagerForNewWorkflow.
type MocktransactionManagerForNewWorkflowMockRecorder struct {
	mock *MocktransactionManagerForNewWorkflow
}

// NewMocktransactionManagerForNewWorkflow creates a new mock instance.
func NewMocktransactionManagerForNewWorkflow(ctrl *gomock.Controller) *MocktransactionManagerForNewWorkflow {
	mock := &MocktransactionManagerForNewWorkflow{ctrl: ctrl}
	mock.recorder = &MocktransactionManagerForNewWorkflowMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MocktransactionManagerForNewWorkflow) EXPECT() *MocktransactionManagerForNewWorkflowMockRecorder {
	return m.recorder
}

// dispatchForNewWorkflow mocks base method.
func (m *MocktransactionManagerForNewWorkflow) dispatchForNewWorkflow(ctx context.Context, now time.Time, targetWorkflow execution.Workflow) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "dispatchForNewWorkflow", ctx, now, targetWorkflow)
	ret0, _ := ret[0].(error)
	return ret0
}

// dispatchForNewWorkflow indicates an expected call of dispatchForNewWorkflow.
func (mr *MocktransactionManagerForNewWorkflowMockRecorder) dispatchForNewWorkflow(ctx, now, targetWorkflow any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "dispatchForNewWorkflow", reflect.TypeOf((*MocktransactionManagerForNewWorkflow)(nil).dispatchForNewWorkflow), ctx, now, targetWorkflow)
}
