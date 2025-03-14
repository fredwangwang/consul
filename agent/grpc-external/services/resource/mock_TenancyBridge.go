// Code generated by mockery v2.20.0. DO NOT EDIT.

package resource

import mock "github.com/stretchr/testify/mock"

// MockTenancyBridge is an autogenerated mock type for the TenancyBridge type
type MockTenancyBridge struct {
	mock.Mock
}

// NamespaceExists provides a mock function with given fields: partition, namespace
func (_m *MockTenancyBridge) NamespaceExists(partition string, namespace string) (bool, error) {
	ret := _m.Called(partition, namespace)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string, string) (bool, error)); ok {
		return rf(partition, namespace)
	}
	if rf, ok := ret.Get(0).(func(string, string) bool); ok {
		r0 = rf(partition, namespace)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(partition, namespace)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PartitionExists provides a mock function with given fields: partition
func (_m *MockTenancyBridge) PartitionExists(partition string) (bool, error) {
	ret := _m.Called(partition)

	var r0 bool
	var r1 error
	if rf, ok := ret.Get(0).(func(string) (bool, error)); ok {
		return rf(partition)
	}
	if rf, ok := ret.Get(0).(func(string) bool); ok {
		r0 = rf(partition)
	} else {
		r0 = ret.Get(0).(bool)
	}

	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(partition)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewMockTenancyBridge interface {
	mock.TestingT
	Cleanup(func())
}

// NewMockTenancyBridge creates a new instance of MockTenancyBridge. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewMockTenancyBridge(t mockConstructorTestingTNewMockTenancyBridge) *MockTenancyBridge {
	mock := &MockTenancyBridge{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
