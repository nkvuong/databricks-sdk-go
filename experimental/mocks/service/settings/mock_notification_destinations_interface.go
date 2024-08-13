// Code generated by mockery v2.43.0. DO NOT EDIT.

package settings

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	settings "github.com/databricks/databricks-sdk-go/service/settings"
)

// MockNotificationDestinationsInterface is an autogenerated mock type for the NotificationDestinationsInterface type
type MockNotificationDestinationsInterface struct {
	mock.Mock
}

type MockNotificationDestinationsInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockNotificationDestinationsInterface) EXPECT() *MockNotificationDestinationsInterface_Expecter {
	return &MockNotificationDestinationsInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockNotificationDestinationsInterface) Create(ctx context.Context, request settings.CreateNotificationDestinationRequest) (*settings.NotificationDestination, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *settings.NotificationDestination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.CreateNotificationDestinationRequest) (*settings.NotificationDestination, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.CreateNotificationDestinationRequest) *settings.NotificationDestination); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.NotificationDestination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.CreateNotificationDestinationRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNotificationDestinationsInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockNotificationDestinationsInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.CreateNotificationDestinationRequest
func (_e *MockNotificationDestinationsInterface_Expecter) Create(ctx interface{}, request interface{}) *MockNotificationDestinationsInterface_Create_Call {
	return &MockNotificationDestinationsInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockNotificationDestinationsInterface_Create_Call) Run(run func(ctx context.Context, request settings.CreateNotificationDestinationRequest)) *MockNotificationDestinationsInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.CreateNotificationDestinationRequest))
	})
	return _c
}

func (_c *MockNotificationDestinationsInterface_Create_Call) Return(_a0 *settings.NotificationDestination, _a1 error) *MockNotificationDestinationsInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNotificationDestinationsInterface_Create_Call) RunAndReturn(run func(context.Context, settings.CreateNotificationDestinationRequest) (*settings.NotificationDestination, error)) *MockNotificationDestinationsInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockNotificationDestinationsInterface) Delete(ctx context.Context, request settings.DeleteNotificationDestinationRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.DeleteNotificationDestinationRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNotificationDestinationsInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockNotificationDestinationsInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.DeleteNotificationDestinationRequest
func (_e *MockNotificationDestinationsInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockNotificationDestinationsInterface_Delete_Call {
	return &MockNotificationDestinationsInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockNotificationDestinationsInterface_Delete_Call) Run(run func(ctx context.Context, request settings.DeleteNotificationDestinationRequest)) *MockNotificationDestinationsInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.DeleteNotificationDestinationRequest))
	})
	return _c
}

func (_c *MockNotificationDestinationsInterface_Delete_Call) Return(_a0 error) *MockNotificationDestinationsInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNotificationDestinationsInterface_Delete_Call) RunAndReturn(run func(context.Context, settings.DeleteNotificationDestinationRequest) error) *MockNotificationDestinationsInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteById provides a mock function with given fields: ctx, id
func (_m *MockNotificationDestinationsInterface) DeleteById(ctx context.Context, id string) error {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for DeleteById")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, string) error); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockNotificationDestinationsInterface_DeleteById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteById'
type MockNotificationDestinationsInterface_DeleteById_Call struct {
	*mock.Call
}

// DeleteById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockNotificationDestinationsInterface_Expecter) DeleteById(ctx interface{}, id interface{}) *MockNotificationDestinationsInterface_DeleteById_Call {
	return &MockNotificationDestinationsInterface_DeleteById_Call{Call: _e.mock.On("DeleteById", ctx, id)}
}

func (_c *MockNotificationDestinationsInterface_DeleteById_Call) Run(run func(ctx context.Context, id string)) *MockNotificationDestinationsInterface_DeleteById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockNotificationDestinationsInterface_DeleteById_Call) Return(_a0 error) *MockNotificationDestinationsInterface_DeleteById_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNotificationDestinationsInterface_DeleteById_Call) RunAndReturn(run func(context.Context, string) error) *MockNotificationDestinationsInterface_DeleteById_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockNotificationDestinationsInterface) Get(ctx context.Context, request settings.GetNotificationDestinationRequest) (*settings.NotificationDestination, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *settings.NotificationDestination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetNotificationDestinationRequest) (*settings.NotificationDestination, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.GetNotificationDestinationRequest) *settings.NotificationDestination); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.NotificationDestination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.GetNotificationDestinationRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNotificationDestinationsInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockNotificationDestinationsInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.GetNotificationDestinationRequest
func (_e *MockNotificationDestinationsInterface_Expecter) Get(ctx interface{}, request interface{}) *MockNotificationDestinationsInterface_Get_Call {
	return &MockNotificationDestinationsInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockNotificationDestinationsInterface_Get_Call) Run(run func(ctx context.Context, request settings.GetNotificationDestinationRequest)) *MockNotificationDestinationsInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.GetNotificationDestinationRequest))
	})
	return _c
}

func (_c *MockNotificationDestinationsInterface_Get_Call) Return(_a0 *settings.NotificationDestination, _a1 error) *MockNotificationDestinationsInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNotificationDestinationsInterface_Get_Call) RunAndReturn(run func(context.Context, settings.GetNotificationDestinationRequest) (*settings.NotificationDestination, error)) *MockNotificationDestinationsInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetById provides a mock function with given fields: ctx, id
func (_m *MockNotificationDestinationsInterface) GetById(ctx context.Context, id string) (*settings.NotificationDestination, error) {
	ret := _m.Called(ctx, id)

	if len(ret) == 0 {
		panic("no return value specified for GetById")
	}

	var r0 *settings.NotificationDestination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*settings.NotificationDestination, error)); ok {
		return rf(ctx, id)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *settings.NotificationDestination); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.NotificationDestination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNotificationDestinationsInterface_GetById_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetById'
type MockNotificationDestinationsInterface_GetById_Call struct {
	*mock.Call
}

// GetById is a helper method to define mock.On call
//   - ctx context.Context
//   - id string
func (_e *MockNotificationDestinationsInterface_Expecter) GetById(ctx interface{}, id interface{}) *MockNotificationDestinationsInterface_GetById_Call {
	return &MockNotificationDestinationsInterface_GetById_Call{Call: _e.mock.On("GetById", ctx, id)}
}

func (_c *MockNotificationDestinationsInterface_GetById_Call) Run(run func(ctx context.Context, id string)) *MockNotificationDestinationsInterface_GetById_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockNotificationDestinationsInterface_GetById_Call) Return(_a0 *settings.NotificationDestination, _a1 error) *MockNotificationDestinationsInterface_GetById_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNotificationDestinationsInterface_GetById_Call) RunAndReturn(run func(context.Context, string) (*settings.NotificationDestination, error)) *MockNotificationDestinationsInterface_GetById_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockNotificationDestinationsInterface) List(ctx context.Context, request settings.ListNotificationDestinationsRequest) listing.Iterator[settings.ListNotificationDestinationsResult] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[settings.ListNotificationDestinationsResult]
	if rf, ok := ret.Get(0).(func(context.Context, settings.ListNotificationDestinationsRequest) listing.Iterator[settings.ListNotificationDestinationsResult]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[settings.ListNotificationDestinationsResult])
		}
	}

	return r0
}

// MockNotificationDestinationsInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockNotificationDestinationsInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.ListNotificationDestinationsRequest
func (_e *MockNotificationDestinationsInterface_Expecter) List(ctx interface{}, request interface{}) *MockNotificationDestinationsInterface_List_Call {
	return &MockNotificationDestinationsInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockNotificationDestinationsInterface_List_Call) Run(run func(ctx context.Context, request settings.ListNotificationDestinationsRequest)) *MockNotificationDestinationsInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.ListNotificationDestinationsRequest))
	})
	return _c
}

func (_c *MockNotificationDestinationsInterface_List_Call) Return(_a0 listing.Iterator[settings.ListNotificationDestinationsResult]) *MockNotificationDestinationsInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockNotificationDestinationsInterface_List_Call) RunAndReturn(run func(context.Context, settings.ListNotificationDestinationsRequest) listing.Iterator[settings.ListNotificationDestinationsResult]) *MockNotificationDestinationsInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockNotificationDestinationsInterface) ListAll(ctx context.Context, request settings.ListNotificationDestinationsRequest) ([]settings.ListNotificationDestinationsResult, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []settings.ListNotificationDestinationsResult
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.ListNotificationDestinationsRequest) ([]settings.ListNotificationDestinationsResult, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.ListNotificationDestinationsRequest) []settings.ListNotificationDestinationsResult); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]settings.ListNotificationDestinationsResult)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.ListNotificationDestinationsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNotificationDestinationsInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockNotificationDestinationsInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.ListNotificationDestinationsRequest
func (_e *MockNotificationDestinationsInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockNotificationDestinationsInterface_ListAll_Call {
	return &MockNotificationDestinationsInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockNotificationDestinationsInterface_ListAll_Call) Run(run func(ctx context.Context, request settings.ListNotificationDestinationsRequest)) *MockNotificationDestinationsInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.ListNotificationDestinationsRequest))
	})
	return _c
}

func (_c *MockNotificationDestinationsInterface_ListAll_Call) Return(_a0 []settings.ListNotificationDestinationsResult, _a1 error) *MockNotificationDestinationsInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNotificationDestinationsInterface_ListAll_Call) RunAndReturn(run func(context.Context, settings.ListNotificationDestinationsRequest) ([]settings.ListNotificationDestinationsResult, error)) *MockNotificationDestinationsInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockNotificationDestinationsInterface) Update(ctx context.Context, request settings.UpdateNotificationDestinationRequest) (*settings.NotificationDestination, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 *settings.NotificationDestination
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateNotificationDestinationRequest) (*settings.NotificationDestination, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, settings.UpdateNotificationDestinationRequest) *settings.NotificationDestination); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*settings.NotificationDestination)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, settings.UpdateNotificationDestinationRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockNotificationDestinationsInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockNotificationDestinationsInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request settings.UpdateNotificationDestinationRequest
func (_e *MockNotificationDestinationsInterface_Expecter) Update(ctx interface{}, request interface{}) *MockNotificationDestinationsInterface_Update_Call {
	return &MockNotificationDestinationsInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockNotificationDestinationsInterface_Update_Call) Run(run func(ctx context.Context, request settings.UpdateNotificationDestinationRequest)) *MockNotificationDestinationsInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(settings.UpdateNotificationDestinationRequest))
	})
	return _c
}

func (_c *MockNotificationDestinationsInterface_Update_Call) Return(_a0 *settings.NotificationDestination, _a1 error) *MockNotificationDestinationsInterface_Update_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockNotificationDestinationsInterface_Update_Call) RunAndReturn(run func(context.Context, settings.UpdateNotificationDestinationRequest) (*settings.NotificationDestination, error)) *MockNotificationDestinationsInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockNotificationDestinationsInterface creates a new instance of MockNotificationDestinationsInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockNotificationDestinationsInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockNotificationDestinationsInterface {
	mock := &MockNotificationDestinationsInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
