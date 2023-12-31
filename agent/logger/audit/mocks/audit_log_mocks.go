// Copyright Amazon.com Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
//     http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/aws/amazon-ecs-agent/agent/logger/audit (interfaces: InfoLogger)

// Package mock_audit is a generated GoMock package.
package mock_audit

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockInfoLogger is a mock of InfoLogger interface.
type MockInfoLogger struct {
	ctrl     *gomock.Controller
	recorder *MockInfoLoggerMockRecorder
}

// MockInfoLoggerMockRecorder is the mock recorder for MockInfoLogger.
type MockInfoLoggerMockRecorder struct {
	mock *MockInfoLogger
}

// NewMockInfoLogger creates a new mock instance.
func NewMockInfoLogger(ctrl *gomock.Controller) *MockInfoLogger {
	mock := &MockInfoLogger{ctrl: ctrl}
	mock.recorder = &MockInfoLoggerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockInfoLogger) EXPECT() *MockInfoLoggerMockRecorder {
	return m.recorder
}

// Info mocks base method.
func (m *MockInfoLogger) Info(arg0 ...interface{}) {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	m.ctrl.Call(m, "Info", varargs...)
}

// Info indicates an expected call of Info.
func (mr *MockInfoLoggerMockRecorder) Info(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Info", reflect.TypeOf((*MockInfoLogger)(nil).Info), arg0...)
}
