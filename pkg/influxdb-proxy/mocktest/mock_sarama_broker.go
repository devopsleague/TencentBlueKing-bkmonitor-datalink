// Tencent is pleased to support the open source community by making
// 蓝鲸智云 - 监控平台 (BlueKing - Monitor) available.
// Copyright (C) 2022 THL A29 Limited, a Tencent company. All rights reserved.
// Licensed under the MIT License (the "License"); you may not use this file except in compliance with the License.
// You may obtain a copy of the License at http://opensource.org/licenses/MIT
// Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on
// an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the
// specific language governing permissions and limitations under the License.

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/TencentBlueKing/bkmonitor-datalink/pkg/influxdb-proxy/backend/influxdb (interfaces: Broker)

// Package mocktest is a generated GoMock package.
package mocktest

import (
	reflect "reflect"

	sarama "github.com/Shopify/sarama"
	gomock "github.com/golang/mock/gomock"
)

// MockBroker is a mock of Broker interface.
type MockBroker struct {
	ctrl     *gomock.Controller
	recorder *MockBrokerMockRecorder
}

// MockBrokerMockRecorder is the mock recorder for MockBroker.
type MockBrokerMockRecorder struct {
	mock *MockBroker
}

// NewMockBroker creates a new mock instance.
func NewMockBroker(ctrl *gomock.Controller) *MockBroker {
	mock := &MockBroker{ctrl: ctrl}
	mock.recorder = &MockBrokerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBroker) EXPECT() *MockBrokerMockRecorder {
	return m.recorder
}

// Close mocks base method.
func (m *MockBroker) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close.
func (mr *MockBrokerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockBroker)(nil).Close))
}

// Connected mocks base method.
func (m *MockBroker) Connected() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Connected")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Connected indicates an expected call of Connected.
func (mr *MockBrokerMockRecorder) Connected() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Connected", reflect.TypeOf((*MockBroker)(nil).Connected))
}

// CreateTopics mocks base method.
func (m *MockBroker) CreateTopics(arg0 *sarama.CreateTopicsRequest) (*sarama.CreateTopicsResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTopics", arg0)
	ret0, _ := ret[0].(*sarama.CreateTopicsResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTopics indicates an expected call of CreateTopics.
func (mr *MockBrokerMockRecorder) CreateTopics(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTopics", reflect.TypeOf((*MockBroker)(nil).CreateTopics), arg0)
}

// Open mocks base method.
func (m *MockBroker) Open(arg0 *sarama.Config) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Open", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Open indicates an expected call of Open.
func (mr *MockBrokerMockRecorder) Open(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Open", reflect.TypeOf((*MockBroker)(nil).Open), arg0)
}