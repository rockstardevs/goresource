package mocks

import (
	"io"
	"net/url"

	"goresource"

	gomock "github.com/golang/mock/gomock"
)

type MockEntity struct {
	Id string
}

func (m *MockEntity) HasId() bool {
	if m.Id != "" {
		return true
	}
	return false
}

func (m *MockEntity) GetId() string {
	return m.Id
}

// Mock of Store interface
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *_MockStoreRecorder
}

// Recorder for MockStore (not exported)
type _MockStoreRecorder struct {
	mock *MockStore
}

func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &_MockStoreRecorder{mock}
	return mock
}

func (_m *MockStore) EXPECT() *_MockStoreRecorder {
	return _m.recorder
}

func (_m *MockStore) Close() {
	_m.ctrl.Call(_m, "Close")
}

func (_mr *_MockStoreRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "Close")
}

func (_m *MockStore) CreateEntity(_param0 string, _param1 interface{}) (interface{}, error) {
	ret := _m.ctrl.Call(_m, "CreateEntity", _param0, _param1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoreRecorder) CreateEntity(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateEntity", arg0, arg1)
}

func (_m *MockStore) DeleteEntity(_param0 string, _param1 string) error {
	ret := _m.ctrl.Call(_m, "DeleteEntity", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockStoreRecorder) DeleteEntity(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteEntity", arg0, arg1)
}

func (_m *MockStore) GetEntity(_param0 string, _param1 string) (interface{}, error) {
	ret := _m.ctrl.Call(_m, "GetEntity", _param0, _param1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoreRecorder) GetEntity(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetEntity", arg0, arg1)
}

func (_m *MockStore) ListEntities(_param0 string, _param1 interface{}) ([]interface{}, error) {
	ret := _m.ctrl.Call(_m, "ListEntities", _param0, _param1)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoreRecorder) ListEntities(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListEntities", arg0, arg1)
}

func (_m *MockStore) UpdateEntity(_param0 string, _param1 string, _param2 interface{}) (interface{}, error) {
	ret := _m.ctrl.Call(_m, "UpdateEntity", _param0, _param1, _param2)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockStoreRecorder) UpdateEntity(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateEntity", arg0, arg1, arg2)
}

// Mock of ResourceManager interface
type MockResourceManager struct {
	ctrl     *gomock.Controller
	recorder *_MockResourceManagerRecorder
}

// Recorder for MockResourceManager (not exported)
type _MockResourceManagerRecorder struct {
	mock *MockResourceManager
}

func NewMockResourceManager(ctrl *gomock.Controller) *MockResourceManager {
	mock := &MockResourceManager{ctrl: ctrl}
	mock.recorder = &_MockResourceManagerRecorder{mock}
	return mock
}

func (_m *MockResourceManager) EXPECT() *_MockResourceManagerRecorder {
	return _m.recorder
}

func (_m *MockResourceManager) CreateEntity(_param0 goresource.Entity, _param1 url.Values) (interface{}, error) {
	ret := _m.ctrl.Call(_m, "CreateEntity", _param0, _param1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockResourceManagerRecorder) CreateEntity(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "CreateEntity", arg0, arg1)
}

func (_m *MockResourceManager) DeleteEntity(_param0 string, _param1 url.Values) error {
	ret := _m.ctrl.Call(_m, "DeleteEntity", _param0, _param1)
	ret0, _ := ret[0].(error)
	return ret0
}

func (_mr *_MockResourceManagerRecorder) DeleteEntity(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "DeleteEntity", arg0, arg1)
}

func (_m *MockResourceManager) GetEntity(_param0 string, _param1 url.Values) (interface{}, error) {
	ret := _m.ctrl.Call(_m, "GetEntity", _param0, _param1)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockResourceManagerRecorder) GetEntity(arg0, arg1 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetEntity", arg0, arg1)
}

func (_m *MockResourceManager) GetName() string {
	ret := _m.ctrl.Call(_m, "GetName")
	ret0, _ := ret[0].(string)
	return ret0
}

func (_mr *_MockResourceManagerRecorder) GetName() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "GetName")
}

func (_m *MockResourceManager) ListEntities(_param0 url.Values) ([]interface{}, error) {
	ret := _m.ctrl.Call(_m, "ListEntities", _param0)
	ret0, _ := ret[0].([]interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockResourceManagerRecorder) ListEntities(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ListEntities", arg0)
}

func (_m *MockResourceManager) New() goresource.Entity {
	ret := _m.ctrl.Call(_m, "New")
	ret0, _ := ret[0].(goresource.Entity)
	return ret0
}

func (_mr *_MockResourceManagerRecorder) New() *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "New")
}

func (_m *MockResourceManager) ParseJSON(_param0 io.ReadCloser) (goresource.Entity, error) {
	ret := _m.ctrl.Call(_m, "ParseJSON", _param0)
	ret0, _ := ret[0].(goresource.Entity)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockResourceManagerRecorder) ParseJSON(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "ParseJSON", arg0)
}

func (_m *MockResourceManager) UpdateEntity(_param0 string, _param1 goresource.Entity, _param2 url.Values) (interface{}, error) {
	ret := _m.ctrl.Call(_m, "UpdateEntity", _param0, _param1, _param2)
	ret0, _ := ret[0].(interface{})
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (_mr *_MockResourceManagerRecorder) UpdateEntity(arg0, arg1, arg2 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCall(_mr.mock, "UpdateEntity", arg0, arg1, arg2)
}
