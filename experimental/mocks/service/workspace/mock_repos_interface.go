// Code generated by mockery v2.43.0. DO NOT EDIT.

package workspace

import (
	context "context"

	listing "github.com/databricks/databricks-sdk-go/listing"
	mock "github.com/stretchr/testify/mock"

	workspace "github.com/databricks/databricks-sdk-go/service/workspace"
)

// MockReposInterface is an autogenerated mock type for the ReposInterface type
type MockReposInterface struct {
	mock.Mock
}

type MockReposInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockReposInterface) EXPECT() *MockReposInterface_Expecter {
	return &MockReposInterface_Expecter{mock: &_m.Mock}
}

// Create provides a mock function with given fields: ctx, request
func (_m *MockReposInterface) Create(ctx context.Context, request workspace.CreateRepo) (*workspace.RepoInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Create")
	}

	var r0 *workspace.RepoInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.CreateRepo) (*workspace.RepoInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workspace.CreateRepo) *workspace.RepoInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.RepoInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workspace.CreateRepo) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReposInterface_Create_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Create'
type MockReposInterface_Create_Call struct {
	*mock.Call
}

// Create is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.CreateRepo
func (_e *MockReposInterface_Expecter) Create(ctx interface{}, request interface{}) *MockReposInterface_Create_Call {
	return &MockReposInterface_Create_Call{Call: _e.mock.On("Create", ctx, request)}
}

func (_c *MockReposInterface_Create_Call) Run(run func(ctx context.Context, request workspace.CreateRepo)) *MockReposInterface_Create_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.CreateRepo))
	})
	return _c
}

func (_c *MockReposInterface_Create_Call) Return(_a0 *workspace.RepoInfo, _a1 error) *MockReposInterface_Create_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReposInterface_Create_Call) RunAndReturn(run func(context.Context, workspace.CreateRepo) (*workspace.RepoInfo, error)) *MockReposInterface_Create_Call {
	_c.Call.Return(run)
	return _c
}

// Delete provides a mock function with given fields: ctx, request
func (_m *MockReposInterface) Delete(ctx context.Context, request workspace.DeleteRepoRequest) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Delete")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.DeleteRepoRequest) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockReposInterface_Delete_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Delete'
type MockReposInterface_Delete_Call struct {
	*mock.Call
}

// Delete is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.DeleteRepoRequest
func (_e *MockReposInterface_Expecter) Delete(ctx interface{}, request interface{}) *MockReposInterface_Delete_Call {
	return &MockReposInterface_Delete_Call{Call: _e.mock.On("Delete", ctx, request)}
}

func (_c *MockReposInterface_Delete_Call) Run(run func(ctx context.Context, request workspace.DeleteRepoRequest)) *MockReposInterface_Delete_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.DeleteRepoRequest))
	})
	return _c
}

func (_c *MockReposInterface_Delete_Call) Return(_a0 error) *MockReposInterface_Delete_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReposInterface_Delete_Call) RunAndReturn(run func(context.Context, workspace.DeleteRepoRequest) error) *MockReposInterface_Delete_Call {
	_c.Call.Return(run)
	return _c
}

// DeleteByRepoId provides a mock function with given fields: ctx, repoId
func (_m *MockReposInterface) DeleteByRepoId(ctx context.Context, repoId int64) error {
	ret := _m.Called(ctx, repoId)

	if len(ret) == 0 {
		panic("no return value specified for DeleteByRepoId")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) error); ok {
		r0 = rf(ctx, repoId)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockReposInterface_DeleteByRepoId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'DeleteByRepoId'
type MockReposInterface_DeleteByRepoId_Call struct {
	*mock.Call
}

// DeleteByRepoId is a helper method to define mock.On call
//   - ctx context.Context
//   - repoId int64
func (_e *MockReposInterface_Expecter) DeleteByRepoId(ctx interface{}, repoId interface{}) *MockReposInterface_DeleteByRepoId_Call {
	return &MockReposInterface_DeleteByRepoId_Call{Call: _e.mock.On("DeleteByRepoId", ctx, repoId)}
}

func (_c *MockReposInterface_DeleteByRepoId_Call) Run(run func(ctx context.Context, repoId int64)) *MockReposInterface_DeleteByRepoId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockReposInterface_DeleteByRepoId_Call) Return(_a0 error) *MockReposInterface_DeleteByRepoId_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReposInterface_DeleteByRepoId_Call) RunAndReturn(run func(context.Context, int64) error) *MockReposInterface_DeleteByRepoId_Call {
	_c.Call.Return(run)
	return _c
}

// Get provides a mock function with given fields: ctx, request
func (_m *MockReposInterface) Get(ctx context.Context, request workspace.GetRepoRequest) (*workspace.RepoInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Get")
	}

	var r0 *workspace.RepoInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.GetRepoRequest) (*workspace.RepoInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workspace.GetRepoRequest) *workspace.RepoInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.RepoInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workspace.GetRepoRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReposInterface_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type MockReposInterface_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.GetRepoRequest
func (_e *MockReposInterface_Expecter) Get(ctx interface{}, request interface{}) *MockReposInterface_Get_Call {
	return &MockReposInterface_Get_Call{Call: _e.mock.On("Get", ctx, request)}
}

func (_c *MockReposInterface_Get_Call) Run(run func(ctx context.Context, request workspace.GetRepoRequest)) *MockReposInterface_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.GetRepoRequest))
	})
	return _c
}

func (_c *MockReposInterface_Get_Call) Return(_a0 *workspace.RepoInfo, _a1 error) *MockReposInterface_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReposInterface_Get_Call) RunAndReturn(run func(context.Context, workspace.GetRepoRequest) (*workspace.RepoInfo, error)) *MockReposInterface_Get_Call {
	_c.Call.Return(run)
	return _c
}

// GetByPath provides a mock function with given fields: ctx, name
func (_m *MockReposInterface) GetByPath(ctx context.Context, name string) (*workspace.RepoInfo, error) {
	ret := _m.Called(ctx, name)

	if len(ret) == 0 {
		panic("no return value specified for GetByPath")
	}

	var r0 *workspace.RepoInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*workspace.RepoInfo, error)); ok {
		return rf(ctx, name)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *workspace.RepoInfo); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.RepoInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReposInterface_GetByPath_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByPath'
type MockReposInterface_GetByPath_Call struct {
	*mock.Call
}

// GetByPath is a helper method to define mock.On call
//   - ctx context.Context
//   - name string
func (_e *MockReposInterface_Expecter) GetByPath(ctx interface{}, name interface{}) *MockReposInterface_GetByPath_Call {
	return &MockReposInterface_GetByPath_Call{Call: _e.mock.On("GetByPath", ctx, name)}
}

func (_c *MockReposInterface_GetByPath_Call) Run(run func(ctx context.Context, name string)) *MockReposInterface_GetByPath_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockReposInterface_GetByPath_Call) Return(_a0 *workspace.RepoInfo, _a1 error) *MockReposInterface_GetByPath_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReposInterface_GetByPath_Call) RunAndReturn(run func(context.Context, string) (*workspace.RepoInfo, error)) *MockReposInterface_GetByPath_Call {
	_c.Call.Return(run)
	return _c
}

// GetByRepoId provides a mock function with given fields: ctx, repoId
func (_m *MockReposInterface) GetByRepoId(ctx context.Context, repoId int64) (*workspace.RepoInfo, error) {
	ret := _m.Called(ctx, repoId)

	if len(ret) == 0 {
		panic("no return value specified for GetByRepoId")
	}

	var r0 *workspace.RepoInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, int64) (*workspace.RepoInfo, error)); ok {
		return rf(ctx, repoId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, int64) *workspace.RepoInfo); ok {
		r0 = rf(ctx, repoId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.RepoInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, int64) error); ok {
		r1 = rf(ctx, repoId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReposInterface_GetByRepoId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByRepoId'
type MockReposInterface_GetByRepoId_Call struct {
	*mock.Call
}

// GetByRepoId is a helper method to define mock.On call
//   - ctx context.Context
//   - repoId int64
func (_e *MockReposInterface_Expecter) GetByRepoId(ctx interface{}, repoId interface{}) *MockReposInterface_GetByRepoId_Call {
	return &MockReposInterface_GetByRepoId_Call{Call: _e.mock.On("GetByRepoId", ctx, repoId)}
}

func (_c *MockReposInterface_GetByRepoId_Call) Run(run func(ctx context.Context, repoId int64)) *MockReposInterface_GetByRepoId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(int64))
	})
	return _c
}

func (_c *MockReposInterface_GetByRepoId_Call) Return(_a0 *workspace.RepoInfo, _a1 error) *MockReposInterface_GetByRepoId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReposInterface_GetByRepoId_Call) RunAndReturn(run func(context.Context, int64) (*workspace.RepoInfo, error)) *MockReposInterface_GetByRepoId_Call {
	_c.Call.Return(run)
	return _c
}

// GetPermissionLevels provides a mock function with given fields: ctx, request
func (_m *MockReposInterface) GetPermissionLevels(ctx context.Context, request workspace.GetRepoPermissionLevelsRequest) (*workspace.GetRepoPermissionLevelsResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetPermissionLevels")
	}

	var r0 *workspace.GetRepoPermissionLevelsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.GetRepoPermissionLevelsRequest) (*workspace.GetRepoPermissionLevelsResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workspace.GetRepoPermissionLevelsRequest) *workspace.GetRepoPermissionLevelsResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.GetRepoPermissionLevelsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workspace.GetRepoPermissionLevelsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReposInterface_GetPermissionLevels_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPermissionLevels'
type MockReposInterface_GetPermissionLevels_Call struct {
	*mock.Call
}

// GetPermissionLevels is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.GetRepoPermissionLevelsRequest
func (_e *MockReposInterface_Expecter) GetPermissionLevels(ctx interface{}, request interface{}) *MockReposInterface_GetPermissionLevels_Call {
	return &MockReposInterface_GetPermissionLevels_Call{Call: _e.mock.On("GetPermissionLevels", ctx, request)}
}

func (_c *MockReposInterface_GetPermissionLevels_Call) Run(run func(ctx context.Context, request workspace.GetRepoPermissionLevelsRequest)) *MockReposInterface_GetPermissionLevels_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.GetRepoPermissionLevelsRequest))
	})
	return _c
}

func (_c *MockReposInterface_GetPermissionLevels_Call) Return(_a0 *workspace.GetRepoPermissionLevelsResponse, _a1 error) *MockReposInterface_GetPermissionLevels_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReposInterface_GetPermissionLevels_Call) RunAndReturn(run func(context.Context, workspace.GetRepoPermissionLevelsRequest) (*workspace.GetRepoPermissionLevelsResponse, error)) *MockReposInterface_GetPermissionLevels_Call {
	_c.Call.Return(run)
	return _c
}

// GetPermissionLevelsByRepoId provides a mock function with given fields: ctx, repoId
func (_m *MockReposInterface) GetPermissionLevelsByRepoId(ctx context.Context, repoId string) (*workspace.GetRepoPermissionLevelsResponse, error) {
	ret := _m.Called(ctx, repoId)

	if len(ret) == 0 {
		panic("no return value specified for GetPermissionLevelsByRepoId")
	}

	var r0 *workspace.GetRepoPermissionLevelsResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*workspace.GetRepoPermissionLevelsResponse, error)); ok {
		return rf(ctx, repoId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *workspace.GetRepoPermissionLevelsResponse); ok {
		r0 = rf(ctx, repoId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.GetRepoPermissionLevelsResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, repoId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReposInterface_GetPermissionLevelsByRepoId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPermissionLevelsByRepoId'
type MockReposInterface_GetPermissionLevelsByRepoId_Call struct {
	*mock.Call
}

// GetPermissionLevelsByRepoId is a helper method to define mock.On call
//   - ctx context.Context
//   - repoId string
func (_e *MockReposInterface_Expecter) GetPermissionLevelsByRepoId(ctx interface{}, repoId interface{}) *MockReposInterface_GetPermissionLevelsByRepoId_Call {
	return &MockReposInterface_GetPermissionLevelsByRepoId_Call{Call: _e.mock.On("GetPermissionLevelsByRepoId", ctx, repoId)}
}

func (_c *MockReposInterface_GetPermissionLevelsByRepoId_Call) Run(run func(ctx context.Context, repoId string)) *MockReposInterface_GetPermissionLevelsByRepoId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockReposInterface_GetPermissionLevelsByRepoId_Call) Return(_a0 *workspace.GetRepoPermissionLevelsResponse, _a1 error) *MockReposInterface_GetPermissionLevelsByRepoId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReposInterface_GetPermissionLevelsByRepoId_Call) RunAndReturn(run func(context.Context, string) (*workspace.GetRepoPermissionLevelsResponse, error)) *MockReposInterface_GetPermissionLevelsByRepoId_Call {
	_c.Call.Return(run)
	return _c
}

// GetPermissions provides a mock function with given fields: ctx, request
func (_m *MockReposInterface) GetPermissions(ctx context.Context, request workspace.GetRepoPermissionsRequest) (*workspace.RepoPermissions, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetPermissions")
	}

	var r0 *workspace.RepoPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.GetRepoPermissionsRequest) (*workspace.RepoPermissions, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workspace.GetRepoPermissionsRequest) *workspace.RepoPermissions); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.RepoPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workspace.GetRepoPermissionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReposInterface_GetPermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPermissions'
type MockReposInterface_GetPermissions_Call struct {
	*mock.Call
}

// GetPermissions is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.GetRepoPermissionsRequest
func (_e *MockReposInterface_Expecter) GetPermissions(ctx interface{}, request interface{}) *MockReposInterface_GetPermissions_Call {
	return &MockReposInterface_GetPermissions_Call{Call: _e.mock.On("GetPermissions", ctx, request)}
}

func (_c *MockReposInterface_GetPermissions_Call) Run(run func(ctx context.Context, request workspace.GetRepoPermissionsRequest)) *MockReposInterface_GetPermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.GetRepoPermissionsRequest))
	})
	return _c
}

func (_c *MockReposInterface_GetPermissions_Call) Return(_a0 *workspace.RepoPermissions, _a1 error) *MockReposInterface_GetPermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReposInterface_GetPermissions_Call) RunAndReturn(run func(context.Context, workspace.GetRepoPermissionsRequest) (*workspace.RepoPermissions, error)) *MockReposInterface_GetPermissions_Call {
	_c.Call.Return(run)
	return _c
}

// GetPermissionsByRepoId provides a mock function with given fields: ctx, repoId
func (_m *MockReposInterface) GetPermissionsByRepoId(ctx context.Context, repoId string) (*workspace.RepoPermissions, error) {
	ret := _m.Called(ctx, repoId)

	if len(ret) == 0 {
		panic("no return value specified for GetPermissionsByRepoId")
	}

	var r0 *workspace.RepoPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string) (*workspace.RepoPermissions, error)); ok {
		return rf(ctx, repoId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string) *workspace.RepoPermissions); ok {
		r0 = rf(ctx, repoId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.RepoPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, repoId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReposInterface_GetPermissionsByRepoId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetPermissionsByRepoId'
type MockReposInterface_GetPermissionsByRepoId_Call struct {
	*mock.Call
}

// GetPermissionsByRepoId is a helper method to define mock.On call
//   - ctx context.Context
//   - repoId string
func (_e *MockReposInterface_Expecter) GetPermissionsByRepoId(ctx interface{}, repoId interface{}) *MockReposInterface_GetPermissionsByRepoId_Call {
	return &MockReposInterface_GetPermissionsByRepoId_Call{Call: _e.mock.On("GetPermissionsByRepoId", ctx, repoId)}
}

func (_c *MockReposInterface_GetPermissionsByRepoId_Call) Run(run func(ctx context.Context, repoId string)) *MockReposInterface_GetPermissionsByRepoId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string))
	})
	return _c
}

func (_c *MockReposInterface_GetPermissionsByRepoId_Call) Return(_a0 *workspace.RepoPermissions, _a1 error) *MockReposInterface_GetPermissionsByRepoId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReposInterface_GetPermissionsByRepoId_Call) RunAndReturn(run func(context.Context, string) (*workspace.RepoPermissions, error)) *MockReposInterface_GetPermissionsByRepoId_Call {
	_c.Call.Return(run)
	return _c
}

// List provides a mock function with given fields: ctx, request
func (_m *MockReposInterface) List(ctx context.Context, request workspace.ListReposRequest) listing.Iterator[workspace.RepoInfo] {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for List")
	}

	var r0 listing.Iterator[workspace.RepoInfo]
	if rf, ok := ret.Get(0).(func(context.Context, workspace.ListReposRequest) listing.Iterator[workspace.RepoInfo]); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(listing.Iterator[workspace.RepoInfo])
		}
	}

	return r0
}

// MockReposInterface_List_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'List'
type MockReposInterface_List_Call struct {
	*mock.Call
}

// List is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.ListReposRequest
func (_e *MockReposInterface_Expecter) List(ctx interface{}, request interface{}) *MockReposInterface_List_Call {
	return &MockReposInterface_List_Call{Call: _e.mock.On("List", ctx, request)}
}

func (_c *MockReposInterface_List_Call) Run(run func(ctx context.Context, request workspace.ListReposRequest)) *MockReposInterface_List_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.ListReposRequest))
	})
	return _c
}

func (_c *MockReposInterface_List_Call) Return(_a0 listing.Iterator[workspace.RepoInfo]) *MockReposInterface_List_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReposInterface_List_Call) RunAndReturn(run func(context.Context, workspace.ListReposRequest) listing.Iterator[workspace.RepoInfo]) *MockReposInterface_List_Call {
	_c.Call.Return(run)
	return _c
}

// ListAll provides a mock function with given fields: ctx, request
func (_m *MockReposInterface) ListAll(ctx context.Context, request workspace.ListReposRequest) ([]workspace.RepoInfo, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ListAll")
	}

	var r0 []workspace.RepoInfo
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.ListReposRequest) ([]workspace.RepoInfo, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workspace.ListReposRequest) []workspace.RepoInfo); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]workspace.RepoInfo)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workspace.ListReposRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReposInterface_ListAll_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ListAll'
type MockReposInterface_ListAll_Call struct {
	*mock.Call
}

// ListAll is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.ListReposRequest
func (_e *MockReposInterface_Expecter) ListAll(ctx interface{}, request interface{}) *MockReposInterface_ListAll_Call {
	return &MockReposInterface_ListAll_Call{Call: _e.mock.On("ListAll", ctx, request)}
}

func (_c *MockReposInterface_ListAll_Call) Run(run func(ctx context.Context, request workspace.ListReposRequest)) *MockReposInterface_ListAll_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.ListReposRequest))
	})
	return _c
}

func (_c *MockReposInterface_ListAll_Call) Return(_a0 []workspace.RepoInfo, _a1 error) *MockReposInterface_ListAll_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReposInterface_ListAll_Call) RunAndReturn(run func(context.Context, workspace.ListReposRequest) ([]workspace.RepoInfo, error)) *MockReposInterface_ListAll_Call {
	_c.Call.Return(run)
	return _c
}

// RepoInfoPathToIdMap provides a mock function with given fields: ctx, request
func (_m *MockReposInterface) RepoInfoPathToIdMap(ctx context.Context, request workspace.ListReposRequest) (map[string]int64, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for RepoInfoPathToIdMap")
	}

	var r0 map[string]int64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.ListReposRequest) (map[string]int64, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workspace.ListReposRequest) map[string]int64); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]int64)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workspace.ListReposRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReposInterface_RepoInfoPathToIdMap_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RepoInfoPathToIdMap'
type MockReposInterface_RepoInfoPathToIdMap_Call struct {
	*mock.Call
}

// RepoInfoPathToIdMap is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.ListReposRequest
func (_e *MockReposInterface_Expecter) RepoInfoPathToIdMap(ctx interface{}, request interface{}) *MockReposInterface_RepoInfoPathToIdMap_Call {
	return &MockReposInterface_RepoInfoPathToIdMap_Call{Call: _e.mock.On("RepoInfoPathToIdMap", ctx, request)}
}

func (_c *MockReposInterface_RepoInfoPathToIdMap_Call) Run(run func(ctx context.Context, request workspace.ListReposRequest)) *MockReposInterface_RepoInfoPathToIdMap_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.ListReposRequest))
	})
	return _c
}

func (_c *MockReposInterface_RepoInfoPathToIdMap_Call) Return(_a0 map[string]int64, _a1 error) *MockReposInterface_RepoInfoPathToIdMap_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReposInterface_RepoInfoPathToIdMap_Call) RunAndReturn(run func(context.Context, workspace.ListReposRequest) (map[string]int64, error)) *MockReposInterface_RepoInfoPathToIdMap_Call {
	_c.Call.Return(run)
	return _c
}

// SetPermissions provides a mock function with given fields: ctx, request
func (_m *MockReposInterface) SetPermissions(ctx context.Context, request workspace.RepoPermissionsRequest) (*workspace.RepoPermissions, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for SetPermissions")
	}

	var r0 *workspace.RepoPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.RepoPermissionsRequest) (*workspace.RepoPermissions, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workspace.RepoPermissionsRequest) *workspace.RepoPermissions); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.RepoPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workspace.RepoPermissionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReposInterface_SetPermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SetPermissions'
type MockReposInterface_SetPermissions_Call struct {
	*mock.Call
}

// SetPermissions is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.RepoPermissionsRequest
func (_e *MockReposInterface_Expecter) SetPermissions(ctx interface{}, request interface{}) *MockReposInterface_SetPermissions_Call {
	return &MockReposInterface_SetPermissions_Call{Call: _e.mock.On("SetPermissions", ctx, request)}
}

func (_c *MockReposInterface_SetPermissions_Call) Run(run func(ctx context.Context, request workspace.RepoPermissionsRequest)) *MockReposInterface_SetPermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.RepoPermissionsRequest))
	})
	return _c
}

func (_c *MockReposInterface_SetPermissions_Call) Return(_a0 *workspace.RepoPermissions, _a1 error) *MockReposInterface_SetPermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReposInterface_SetPermissions_Call) RunAndReturn(run func(context.Context, workspace.RepoPermissionsRequest) (*workspace.RepoPermissions, error)) *MockReposInterface_SetPermissions_Call {
	_c.Call.Return(run)
	return _c
}

// Update provides a mock function with given fields: ctx, request
func (_m *MockReposInterface) Update(ctx context.Context, request workspace.UpdateRepo) error {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for Update")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.UpdateRepo) error); ok {
		r0 = rf(ctx, request)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockReposInterface_Update_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Update'
type MockReposInterface_Update_Call struct {
	*mock.Call
}

// Update is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.UpdateRepo
func (_e *MockReposInterface_Expecter) Update(ctx interface{}, request interface{}) *MockReposInterface_Update_Call {
	return &MockReposInterface_Update_Call{Call: _e.mock.On("Update", ctx, request)}
}

func (_c *MockReposInterface_Update_Call) Run(run func(ctx context.Context, request workspace.UpdateRepo)) *MockReposInterface_Update_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.UpdateRepo))
	})
	return _c
}

func (_c *MockReposInterface_Update_Call) Return(_a0 error) *MockReposInterface_Update_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockReposInterface_Update_Call) RunAndReturn(run func(context.Context, workspace.UpdateRepo) error) *MockReposInterface_Update_Call {
	_c.Call.Return(run)
	return _c
}

// UpdatePermissions provides a mock function with given fields: ctx, request
func (_m *MockReposInterface) UpdatePermissions(ctx context.Context, request workspace.RepoPermissionsRequest) (*workspace.RepoPermissions, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for UpdatePermissions")
	}

	var r0 *workspace.RepoPermissions
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, workspace.RepoPermissionsRequest) (*workspace.RepoPermissions, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, workspace.RepoPermissionsRequest) *workspace.RepoPermissions); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*workspace.RepoPermissions)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, workspace.RepoPermissionsRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockReposInterface_UpdatePermissions_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'UpdatePermissions'
type MockReposInterface_UpdatePermissions_Call struct {
	*mock.Call
}

// UpdatePermissions is a helper method to define mock.On call
//   - ctx context.Context
//   - request workspace.RepoPermissionsRequest
func (_e *MockReposInterface_Expecter) UpdatePermissions(ctx interface{}, request interface{}) *MockReposInterface_UpdatePermissions_Call {
	return &MockReposInterface_UpdatePermissions_Call{Call: _e.mock.On("UpdatePermissions", ctx, request)}
}

func (_c *MockReposInterface_UpdatePermissions_Call) Run(run func(ctx context.Context, request workspace.RepoPermissionsRequest)) *MockReposInterface_UpdatePermissions_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(workspace.RepoPermissionsRequest))
	})
	return _c
}

func (_c *MockReposInterface_UpdatePermissions_Call) Return(_a0 *workspace.RepoPermissions, _a1 error) *MockReposInterface_UpdatePermissions_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockReposInterface_UpdatePermissions_Call) RunAndReturn(run func(context.Context, workspace.RepoPermissionsRequest) (*workspace.RepoPermissions, error)) *MockReposInterface_UpdatePermissions_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockReposInterface creates a new instance of MockReposInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockReposInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockReposInterface {
	mock := &MockReposInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
