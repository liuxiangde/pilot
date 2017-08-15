// Automatically generated by MockGen. DO NOT EDIT!
// Source: model/config.go

package model

import (
	gomock "github.com/golang/mock/gomock"
	proto "github.com/golang/protobuf/proto"
	config "istio.io/api/proxy/v1/config"
)

// Mock of ConfigStore interface
type MockConfigStore struct {
	ctrl     *gomock.Controller
	recorder *_MockConfigStoreRecorder
}

// Recorder for MockConfigStore (not exported)
type _MockConfigStoreRecorder struct {
	mock *MockConfigStore
}

func NewMockConfigStore(ctrl *gomock.Controller) *MockConfigStore {
	mock := &MockConfigStore{ctrl: ctrl}
	mock.recorder = &_MockConfigStoreRecorder{mock}
	return mock
}

func (_m *MockConfigStore) EXPECT() *_MockConfigStoreRecorder {
	return _m.recorder
}

func (_m *MockConfigStore) ConfigDescriptor() ConfigDescriptor {
	ret := _m.ctrl.Call(_m, "ConfigDescriptor")
	ret0, _ := ret[0].(ConfigDescriptor)
	return ret0
}

func (_mr *_MockConfigStoreRecorder) ConfigDescriptor() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConfigDescriptor")
}

func (_m *MockConfigStore) Get(typ string, key string) (proto.Message, bool, string) {
	ret := _m.ctrl.Call(_m, "Get", typ, key)
	ret0, _ := ret[0].(proto.Message)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(string)
	return ret0, ret1, ret2
}

func (_mr *_MockConfigStoreRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}

func (_m *MockConfigStore) List(typ string) ([]Config, error) {
	ret := _m.ctrl.Call(_m, "List", typ)
	ret0, _ := ret[0].([]Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConfigStoreRecorder) List(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List", arg0)
}

func (_m *MockConfigStore) Post(config proto.Message) (string, error) {
	ret := _m.ctrl.Call(_m, "Post", config)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConfigStoreRecorder) Post(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Post", arg0)
}

func (_m *MockConfigStore) Put(config proto.Message, oldRevision string) (string, error) {
	ret := _m.ctrl.Call(_m, "Put", config, oldRevision)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConfigStoreRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0, arg1)
}

func (_m *MockConfigStore) Delete(typ string, key string) error {
	ret := _m.ctrl.Call(_m, "Delete", typ, key)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConfigStoreRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1)
}

// Mock of ConfigStoreCache interface
type MockConfigStoreCache struct {
	ctrl     *gomock.Controller
	recorder *_MockConfigStoreCacheRecorder
}

// Recorder for MockConfigStoreCache (not exported)
type _MockConfigStoreCacheRecorder struct {
	mock *MockConfigStoreCache
}

func NewMockConfigStoreCache(ctrl *gomock.Controller) *MockConfigStoreCache {
	mock := &MockConfigStoreCache{ctrl: ctrl}
	mock.recorder = &_MockConfigStoreCacheRecorder{mock}
	return mock
}

func (_m *MockConfigStoreCache) EXPECT() *_MockConfigStoreCacheRecorder {
	return _m.recorder
}

func (_m *MockConfigStoreCache) ConfigDescriptor() ConfigDescriptor {
	ret := _m.ctrl.Call(_m, "ConfigDescriptor")
	ret0, _ := ret[0].(ConfigDescriptor)
	return ret0
}

func (_mr *_MockConfigStoreCacheRecorder) ConfigDescriptor() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ConfigDescriptor")
}

func (_m *MockConfigStoreCache) Get(typ string, key string) (proto.Message, bool, string) {
	ret := _m.ctrl.Call(_m, "Get", typ, key)
	ret0, _ := ret[0].(proto.Message)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(string)
	return ret0, ret1, ret2
}

func (_mr *_MockConfigStoreCacheRecorder) Get(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Get", arg0, arg1)
}

func (_m *MockConfigStoreCache) List(typ string) ([]Config, error) {
	ret := _m.ctrl.Call(_m, "List", typ)
	ret0, _ := ret[0].([]Config)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConfigStoreCacheRecorder) List(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "List", arg0)
}

func (_m *MockConfigStoreCache) Post(config proto.Message) (string, error) {
	ret := _m.ctrl.Call(_m, "Post", config)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConfigStoreCacheRecorder) Post(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Post", arg0)
}

func (_m *MockConfigStoreCache) Put(config proto.Message, oldRevision string) (string, error) {
	ret := _m.ctrl.Call(_m, "Put", config, oldRevision)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockConfigStoreCacheRecorder) Put(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Put", arg0, arg1)
}

func (_m *MockConfigStoreCache) Delete(typ string, key string) error {
	ret := _m.ctrl.Call(_m, "Delete", typ, key)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockConfigStoreCacheRecorder) Delete(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Delete", arg0, arg1)
}

func (_m *MockConfigStoreCache) RegisterEventHandler(typ string, handler func(Config, Event)) {
	_m.ctrl.Call(_m, "RegisterEventHandler", typ, handler)
}

func (_mr *_MockConfigStoreCacheRecorder) RegisterEventHandler(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RegisterEventHandler", arg0, arg1)
}

func (_m *MockConfigStoreCache) Run(stop <-chan struct{}) {
	_m.ctrl.Call(_m, "Run", stop)
}

func (_mr *_MockConfigStoreCacheRecorder) Run(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Run", arg0)
}

func (_m *MockConfigStoreCache) HasSynced() bool {
	ret := _m.ctrl.Call(_m, "HasSynced")
	ret0, _ := ret[0].(bool)
	return ret0
}

func (_mr *_MockConfigStoreCacheRecorder) HasSynced() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "HasSynced")
}

// Mock of IstioConfigStore interface
type MockIstioConfigStore struct {
	ctrl     *gomock.Controller
	recorder *_MockIstioConfigStoreRecorder
}

// Recorder for MockIstioConfigStore (not exported)
type _MockIstioConfigStoreRecorder struct {
	mock *MockIstioConfigStore
}

func NewMockIstioConfigStore(ctrl *gomock.Controller) *MockIstioConfigStore {
	mock := &MockIstioConfigStore{ctrl: ctrl}
	mock.recorder = &_MockIstioConfigStoreRecorder{mock}
	return mock
}

func (_m *MockIstioConfigStore) EXPECT() *_MockIstioConfigStoreRecorder {
	return _m.recorder
}

func (_m *MockIstioConfigStore) RouteRules() map[string]*config.RouteRule {
	ret := _m.ctrl.Call(_m, "RouteRules")
	ret0, _ := ret[0].(map[string]*config.RouteRule)
	return ret0
}

func (_mr *_MockIstioConfigStoreRecorder) RouteRules() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RouteRules")
}

func (_m *MockIstioConfigStore) IngressRules() map[string]*config.IngressRule {
	ret := _m.ctrl.Call(_m, "IngressRules")
	ret0, _ := ret[0].(map[string]*config.IngressRule)
	return ret0
}

func (_mr *_MockIstioConfigStoreRecorder) IngressRules() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "IngressRules")
}

func (_m *MockIstioConfigStore) DestinationPolicies() []*config.DestinationPolicy {
	ret := _m.ctrl.Call(_m, "DestinationPolicies")
	ret0, _ := ret[0].([]*config.DestinationPolicy)
	return ret0
}

func (_mr *_MockIstioConfigStoreRecorder) DestinationPolicies() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DestinationPolicies")
}

func (_m *MockIstioConfigStore) RouteRulesBySource(instances []*ServiceInstance) []*config.RouteRule {
	ret := _m.ctrl.Call(_m, "RouteRulesBySource", instances)
	ret0, _ := ret[0].([]*config.RouteRule)
	return ret0
}

func (_mr *_MockIstioConfigStoreRecorder) RouteRulesBySource(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "RouteRulesBySource", arg0)
}

func (_m *MockIstioConfigStore) DestinationPolicy(destination string, tags Tags) *config.DestinationVersionPolicy {
	ret := _m.ctrl.Call(_m, "DestinationPolicy", destination, tags)
	ret0, _ := ret[0].(*config.DestinationVersionPolicy)
	return ret0
}

func (_mr *_MockIstioConfigStoreRecorder) DestinationPolicy(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DestinationPolicy", arg0, arg1)
}