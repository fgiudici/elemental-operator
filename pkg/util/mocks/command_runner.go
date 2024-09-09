// /*
// Copyright © 2022 - 2024 SUSE LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
// */
//
//

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/rancher/elemental-operator/pkg/util (interfaces: CommandRunner)
//
// Generated by this command:
//
//	mockgen -copyright_file=scripts/boilerplate.go.txt -destination=pkg/util/mocks/command_runner.go -package=mocks github.com/rancher/elemental-operator/pkg/util CommandRunner
//

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockCommandRunner is a mock of CommandRunner interface.
type MockCommandRunner struct {
	ctrl     *gomock.Controller
	recorder *MockCommandRunnerMockRecorder
}

// MockCommandRunnerMockRecorder is the mock recorder for MockCommandRunner.
type MockCommandRunnerMockRecorder struct {
	mock *MockCommandRunner
}

// NewMockCommandRunner creates a new mock instance.
func NewMockCommandRunner(ctrl *gomock.Controller) *MockCommandRunner {
	mock := &MockCommandRunner{ctrl: ctrl}
	mock.recorder = &MockCommandRunnerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCommandRunner) EXPECT() *MockCommandRunnerMockRecorder {
	return m.recorder
}

// Run mocks base method.
func (m *MockCommandRunner) Run(arg0 string, arg1 ...string) error {
	m.ctrl.T.Helper()
	varargs := []any{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Run", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// Run indicates an expected call of Run.
func (mr *MockCommandRunnerMockRecorder) Run(arg0 any, arg1 ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Run", reflect.TypeOf((*MockCommandRunner)(nil).Run), varargs...)
}