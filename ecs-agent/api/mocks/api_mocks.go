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
// Source: github.com/aws/amazon-ecs-agent/ecs-agent/api (interfaces: ECSTaskProtectionSDK)

// Package mock_api is a generated GoMock package.
package mock_api

import (
	context "context"
	reflect "reflect"

	ecs "github.com/aws/amazon-ecs-agent/ecs-agent/ecs_client/model/ecs"
	request "github.com/aws/aws-sdk-go/aws/request"
	gomock "github.com/golang/mock/gomock"
)

// MockECSTaskProtectionSDK is a mock of ECSTaskProtectionSDK interface.
type MockECSTaskProtectionSDK struct {
	ctrl     *gomock.Controller
	recorder *MockECSTaskProtectionSDKMockRecorder
}

// MockECSTaskProtectionSDKMockRecorder is the mock recorder for MockECSTaskProtectionSDK.
type MockECSTaskProtectionSDKMockRecorder struct {
	mock *MockECSTaskProtectionSDK
}

// NewMockECSTaskProtectionSDK creates a new mock instance.
func NewMockECSTaskProtectionSDK(ctrl *gomock.Controller) *MockECSTaskProtectionSDK {
	mock := &MockECSTaskProtectionSDK{ctrl: ctrl}
	mock.recorder = &MockECSTaskProtectionSDKMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockECSTaskProtectionSDK) EXPECT() *MockECSTaskProtectionSDKMockRecorder {
	return m.recorder
}

// GetTaskProtection mocks base method.
func (m *MockECSTaskProtectionSDK) GetTaskProtection(arg0 *ecs.GetTaskProtectionInput) (*ecs.GetTaskProtectionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTaskProtection", arg0)
	ret0, _ := ret[0].(*ecs.GetTaskProtectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskProtection indicates an expected call of GetTaskProtection.
func (mr *MockECSTaskProtectionSDKMockRecorder) GetTaskProtection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskProtection", reflect.TypeOf((*MockECSTaskProtectionSDK)(nil).GetTaskProtection), arg0)
}

// GetTaskProtectionWithContext mocks base method.
func (m *MockECSTaskProtectionSDK) GetTaskProtectionWithContext(arg0 context.Context, arg1 *ecs.GetTaskProtectionInput, arg2 ...request.Option) (*ecs.GetTaskProtectionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetTaskProtectionWithContext", varargs...)
	ret0, _ := ret[0].(*ecs.GetTaskProtectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetTaskProtectionWithContext indicates an expected call of GetTaskProtectionWithContext.
func (mr *MockECSTaskProtectionSDKMockRecorder) GetTaskProtectionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTaskProtectionWithContext", reflect.TypeOf((*MockECSTaskProtectionSDK)(nil).GetTaskProtectionWithContext), varargs...)
}

// UpdateTaskProtection mocks base method.
func (m *MockECSTaskProtectionSDK) UpdateTaskProtection(arg0 *ecs.UpdateTaskProtectionInput) (*ecs.UpdateTaskProtectionOutput, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateTaskProtection", arg0)
	ret0, _ := ret[0].(*ecs.UpdateTaskProtectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTaskProtection indicates an expected call of UpdateTaskProtection.
func (mr *MockECSTaskProtectionSDKMockRecorder) UpdateTaskProtection(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTaskProtection", reflect.TypeOf((*MockECSTaskProtectionSDK)(nil).UpdateTaskProtection), arg0)
}

// UpdateTaskProtectionWithContext mocks base method.
func (m *MockECSTaskProtectionSDK) UpdateTaskProtectionWithContext(arg0 context.Context, arg1 *ecs.UpdateTaskProtectionInput, arg2 ...request.Option) (*ecs.UpdateTaskProtectionOutput, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "UpdateTaskProtectionWithContext", varargs...)
	ret0, _ := ret[0].(*ecs.UpdateTaskProtectionOutput)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateTaskProtectionWithContext indicates an expected call of UpdateTaskProtectionWithContext.
func (mr *MockECSTaskProtectionSDKMockRecorder) UpdateTaskProtectionWithContext(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateTaskProtectionWithContext", reflect.TypeOf((*MockECSTaskProtectionSDK)(nil).UpdateTaskProtectionWithContext), varargs...)
}
