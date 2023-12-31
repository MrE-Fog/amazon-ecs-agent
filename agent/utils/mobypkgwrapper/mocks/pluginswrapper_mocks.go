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
// Source: github.com/aws/amazon-ecs-agent/agent/utils/mobypkgwrapper (interfaces: Plugins)

// Package mock_mobypkgwrapper is a generated GoMock package.
package mock_mobypkgwrapper

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockPlugins is a mock of Plugins interface.
type MockPlugins struct {
	ctrl     *gomock.Controller
	recorder *MockPluginsMockRecorder
}

// MockPluginsMockRecorder is the mock recorder for MockPlugins.
type MockPluginsMockRecorder struct {
	mock *MockPlugins
}

// NewMockPlugins creates a new mock instance.
func NewMockPlugins(ctrl *gomock.Controller) *MockPlugins {
	mock := &MockPlugins{ctrl: ctrl}
	mock.recorder = &MockPluginsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPlugins) EXPECT() *MockPluginsMockRecorder {
	return m.recorder
}

// Scan mocks base method.
func (m *MockPlugins) Scan() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Scan")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Scan indicates an expected call of Scan.
func (mr *MockPluginsMockRecorder) Scan() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Scan", reflect.TypeOf((*MockPlugins)(nil).Scan))
}
