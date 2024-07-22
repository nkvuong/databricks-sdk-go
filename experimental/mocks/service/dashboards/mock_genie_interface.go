// Code generated by mockery v2.43.0. DO NOT EDIT.

package dashboards

import (
	context "context"

	dashboards "github.com/databricks/databricks-sdk-go/service/dashboards"
	mock "github.com/stretchr/testify/mock"

	retries "github.com/databricks/databricks-sdk-go/retries"

	time "time"
)

// MockGenieInterface is an autogenerated mock type for the GenieInterface type
type MockGenieInterface struct {
	mock.Mock
}

type MockGenieInterface_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGenieInterface) EXPECT() *MockGenieInterface_Expecter {
	return &MockGenieInterface_Expecter{mock: &_m.Mock}
}

// CreateMessage provides a mock function with given fields: ctx, genieCreateConversationMessageRequest
func (_m *MockGenieInterface) CreateMessage(ctx context.Context, genieCreateConversationMessageRequest dashboards.GenieCreateConversationMessageRequest) (*dashboards.WaitGetMessageGenieCompleted[dashboards.GenieMessage], error) {
	ret := _m.Called(ctx, genieCreateConversationMessageRequest)

	if len(ret) == 0 {
		panic("no return value specified for CreateMessage")
	}

	var r0 *dashboards.WaitGetMessageGenieCompleted[dashboards.GenieMessage]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GenieCreateConversationMessageRequest) (*dashboards.WaitGetMessageGenieCompleted[dashboards.GenieMessage], error)); ok {
		return rf(ctx, genieCreateConversationMessageRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GenieCreateConversationMessageRequest) *dashboards.WaitGetMessageGenieCompleted[dashboards.GenieMessage]); ok {
		r0 = rf(ctx, genieCreateConversationMessageRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dashboards.WaitGetMessageGenieCompleted[dashboards.GenieMessage])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dashboards.GenieCreateConversationMessageRequest) error); ok {
		r1 = rf(ctx, genieCreateConversationMessageRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGenieInterface_CreateMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMessage'
type MockGenieInterface_CreateMessage_Call struct {
	*mock.Call
}

// CreateMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - genieCreateConversationMessageRequest dashboards.GenieCreateConversationMessageRequest
func (_e *MockGenieInterface_Expecter) CreateMessage(ctx interface{}, genieCreateConversationMessageRequest interface{}) *MockGenieInterface_CreateMessage_Call {
	return &MockGenieInterface_CreateMessage_Call{Call: _e.mock.On("CreateMessage", ctx, genieCreateConversationMessageRequest)}
}

func (_c *MockGenieInterface_CreateMessage_Call) Run(run func(ctx context.Context, genieCreateConversationMessageRequest dashboards.GenieCreateConversationMessageRequest)) *MockGenieInterface_CreateMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dashboards.GenieCreateConversationMessageRequest))
	})
	return _c
}

func (_c *MockGenieInterface_CreateMessage_Call) Return(_a0 *dashboards.WaitGetMessageGenieCompleted[dashboards.GenieMessage], _a1 error) *MockGenieInterface_CreateMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGenieInterface_CreateMessage_Call) RunAndReturn(run func(context.Context, dashboards.GenieCreateConversationMessageRequest) (*dashboards.WaitGetMessageGenieCompleted[dashboards.GenieMessage], error)) *MockGenieInterface_CreateMessage_Call {
	_c.Call.Return(run)
	return _c
}

// CreateMessageAndWait provides a mock function with given fields: ctx, genieCreateConversationMessageRequest, options
func (_m *MockGenieInterface) CreateMessageAndWait(ctx context.Context, genieCreateConversationMessageRequest dashboards.GenieCreateConversationMessageRequest, options ...retries.Option[dashboards.GenieMessage]) (*dashboards.GenieMessage, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, genieCreateConversationMessageRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for CreateMessageAndWait")
	}

	var r0 *dashboards.GenieMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GenieCreateConversationMessageRequest, ...retries.Option[dashboards.GenieMessage]) (*dashboards.GenieMessage, error)); ok {
		return rf(ctx, genieCreateConversationMessageRequest, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GenieCreateConversationMessageRequest, ...retries.Option[dashboards.GenieMessage]) *dashboards.GenieMessage); ok {
		r0 = rf(ctx, genieCreateConversationMessageRequest, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dashboards.GenieMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dashboards.GenieCreateConversationMessageRequest, ...retries.Option[dashboards.GenieMessage]) error); ok {
		r1 = rf(ctx, genieCreateConversationMessageRequest, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGenieInterface_CreateMessageAndWait_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'CreateMessageAndWait'
type MockGenieInterface_CreateMessageAndWait_Call struct {
	*mock.Call
}

// CreateMessageAndWait is a helper method to define mock.On call
//   - ctx context.Context
//   - genieCreateConversationMessageRequest dashboards.GenieCreateConversationMessageRequest
//   - options ...retries.Option[dashboards.GenieMessage]
func (_e *MockGenieInterface_Expecter) CreateMessageAndWait(ctx interface{}, genieCreateConversationMessageRequest interface{}, options ...interface{}) *MockGenieInterface_CreateMessageAndWait_Call {
	return &MockGenieInterface_CreateMessageAndWait_Call{Call: _e.mock.On("CreateMessageAndWait",
		append([]interface{}{ctx, genieCreateConversationMessageRequest}, options...)...)}
}

func (_c *MockGenieInterface_CreateMessageAndWait_Call) Run(run func(ctx context.Context, genieCreateConversationMessageRequest dashboards.GenieCreateConversationMessageRequest, options ...retries.Option[dashboards.GenieMessage])) *MockGenieInterface_CreateMessageAndWait_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]retries.Option[dashboards.GenieMessage], len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(retries.Option[dashboards.GenieMessage])
			}
		}
		run(args[0].(context.Context), args[1].(dashboards.GenieCreateConversationMessageRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockGenieInterface_CreateMessageAndWait_Call) Return(_a0 *dashboards.GenieMessage, _a1 error) *MockGenieInterface_CreateMessageAndWait_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGenieInterface_CreateMessageAndWait_Call) RunAndReturn(run func(context.Context, dashboards.GenieCreateConversationMessageRequest, ...retries.Option[dashboards.GenieMessage]) (*dashboards.GenieMessage, error)) *MockGenieInterface_CreateMessageAndWait_Call {
	_c.Call.Return(run)
	return _c
}

// ExecuteMessageQuery provides a mock function with given fields: ctx, request
func (_m *MockGenieInterface) ExecuteMessageQuery(ctx context.Context, request dashboards.ExecuteMessageQueryRequest) (*dashboards.GenieGetMessageQueryResultResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for ExecuteMessageQuery")
	}

	var r0 *dashboards.GenieGetMessageQueryResultResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.ExecuteMessageQueryRequest) (*dashboards.GenieGetMessageQueryResultResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.ExecuteMessageQueryRequest) *dashboards.GenieGetMessageQueryResultResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dashboards.GenieGetMessageQueryResultResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dashboards.ExecuteMessageQueryRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGenieInterface_ExecuteMessageQuery_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ExecuteMessageQuery'
type MockGenieInterface_ExecuteMessageQuery_Call struct {
	*mock.Call
}

// ExecuteMessageQuery is a helper method to define mock.On call
//   - ctx context.Context
//   - request dashboards.ExecuteMessageQueryRequest
func (_e *MockGenieInterface_Expecter) ExecuteMessageQuery(ctx interface{}, request interface{}) *MockGenieInterface_ExecuteMessageQuery_Call {
	return &MockGenieInterface_ExecuteMessageQuery_Call{Call: _e.mock.On("ExecuteMessageQuery", ctx, request)}
}

func (_c *MockGenieInterface_ExecuteMessageQuery_Call) Run(run func(ctx context.Context, request dashboards.ExecuteMessageQueryRequest)) *MockGenieInterface_ExecuteMessageQuery_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dashboards.ExecuteMessageQueryRequest))
	})
	return _c
}

func (_c *MockGenieInterface_ExecuteMessageQuery_Call) Return(_a0 *dashboards.GenieGetMessageQueryResultResponse, _a1 error) *MockGenieInterface_ExecuteMessageQuery_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGenieInterface_ExecuteMessageQuery_Call) RunAndReturn(run func(context.Context, dashboards.ExecuteMessageQueryRequest) (*dashboards.GenieGetMessageQueryResultResponse, error)) *MockGenieInterface_ExecuteMessageQuery_Call {
	_c.Call.Return(run)
	return _c
}

// GetMessage provides a mock function with given fields: ctx, request
func (_m *MockGenieInterface) GetMessage(ctx context.Context, request dashboards.GenieGetConversationMessageRequest) (*dashboards.GenieMessage, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetMessage")
	}

	var r0 *dashboards.GenieMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GenieGetConversationMessageRequest) (*dashboards.GenieMessage, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GenieGetConversationMessageRequest) *dashboards.GenieMessage); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dashboards.GenieMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dashboards.GenieGetConversationMessageRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGenieInterface_GetMessage_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMessage'
type MockGenieInterface_GetMessage_Call struct {
	*mock.Call
}

// GetMessage is a helper method to define mock.On call
//   - ctx context.Context
//   - request dashboards.GenieGetConversationMessageRequest
func (_e *MockGenieInterface_Expecter) GetMessage(ctx interface{}, request interface{}) *MockGenieInterface_GetMessage_Call {
	return &MockGenieInterface_GetMessage_Call{Call: _e.mock.On("GetMessage", ctx, request)}
}

func (_c *MockGenieInterface_GetMessage_Call) Run(run func(ctx context.Context, request dashboards.GenieGetConversationMessageRequest)) *MockGenieInterface_GetMessage_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dashboards.GenieGetConversationMessageRequest))
	})
	return _c
}

func (_c *MockGenieInterface_GetMessage_Call) Return(_a0 *dashboards.GenieMessage, _a1 error) *MockGenieInterface_GetMessage_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGenieInterface_GetMessage_Call) RunAndReturn(run func(context.Context, dashboards.GenieGetConversationMessageRequest) (*dashboards.GenieMessage, error)) *MockGenieInterface_GetMessage_Call {
	_c.Call.Return(run)
	return _c
}

// GetMessageBySpaceIdAndConversationIdAndMessageId provides a mock function with given fields: ctx, spaceId, conversationId, messageId
func (_m *MockGenieInterface) GetMessageBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*dashboards.GenieMessage, error) {
	ret := _m.Called(ctx, spaceId, conversationId, messageId)

	if len(ret) == 0 {
		panic("no return value specified for GetMessageBySpaceIdAndConversationIdAndMessageId")
	}

	var r0 *dashboards.GenieMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*dashboards.GenieMessage, error)); ok {
		return rf(ctx, spaceId, conversationId, messageId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *dashboards.GenieMessage); ok {
		r0 = rf(ctx, spaceId, conversationId, messageId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dashboards.GenieMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, spaceId, conversationId, messageId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGenieInterface_GetMessageBySpaceIdAndConversationIdAndMessageId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMessageBySpaceIdAndConversationIdAndMessageId'
type MockGenieInterface_GetMessageBySpaceIdAndConversationIdAndMessageId_Call struct {
	*mock.Call
}

// GetMessageBySpaceIdAndConversationIdAndMessageId is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - conversationId string
//   - messageId string
func (_e *MockGenieInterface_Expecter) GetMessageBySpaceIdAndConversationIdAndMessageId(ctx interface{}, spaceId interface{}, conversationId interface{}, messageId interface{}) *MockGenieInterface_GetMessageBySpaceIdAndConversationIdAndMessageId_Call {
	return &MockGenieInterface_GetMessageBySpaceIdAndConversationIdAndMessageId_Call{Call: _e.mock.On("GetMessageBySpaceIdAndConversationIdAndMessageId", ctx, spaceId, conversationId, messageId)}
}

func (_c *MockGenieInterface_GetMessageBySpaceIdAndConversationIdAndMessageId_Call) Run(run func(ctx context.Context, spaceId string, conversationId string, messageId string)) *MockGenieInterface_GetMessageBySpaceIdAndConversationIdAndMessageId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockGenieInterface_GetMessageBySpaceIdAndConversationIdAndMessageId_Call) Return(_a0 *dashboards.GenieMessage, _a1 error) *MockGenieInterface_GetMessageBySpaceIdAndConversationIdAndMessageId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGenieInterface_GetMessageBySpaceIdAndConversationIdAndMessageId_Call) RunAndReturn(run func(context.Context, string, string, string) (*dashboards.GenieMessage, error)) *MockGenieInterface_GetMessageBySpaceIdAndConversationIdAndMessageId_Call {
	_c.Call.Return(run)
	return _c
}

// GetMessageQueryResult provides a mock function with given fields: ctx, request
func (_m *MockGenieInterface) GetMessageQueryResult(ctx context.Context, request dashboards.GenieGetMessageQueryResultRequest) (*dashboards.GenieGetMessageQueryResultResponse, error) {
	ret := _m.Called(ctx, request)

	if len(ret) == 0 {
		panic("no return value specified for GetMessageQueryResult")
	}

	var r0 *dashboards.GenieGetMessageQueryResultResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GenieGetMessageQueryResultRequest) (*dashboards.GenieGetMessageQueryResultResponse, error)); ok {
		return rf(ctx, request)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GenieGetMessageQueryResultRequest) *dashboards.GenieGetMessageQueryResultResponse); ok {
		r0 = rf(ctx, request)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dashboards.GenieGetMessageQueryResultResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dashboards.GenieGetMessageQueryResultRequest) error); ok {
		r1 = rf(ctx, request)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGenieInterface_GetMessageQueryResult_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMessageQueryResult'
type MockGenieInterface_GetMessageQueryResult_Call struct {
	*mock.Call
}

// GetMessageQueryResult is a helper method to define mock.On call
//   - ctx context.Context
//   - request dashboards.GenieGetMessageQueryResultRequest
func (_e *MockGenieInterface_Expecter) GetMessageQueryResult(ctx interface{}, request interface{}) *MockGenieInterface_GetMessageQueryResult_Call {
	return &MockGenieInterface_GetMessageQueryResult_Call{Call: _e.mock.On("GetMessageQueryResult", ctx, request)}
}

func (_c *MockGenieInterface_GetMessageQueryResult_Call) Run(run func(ctx context.Context, request dashboards.GenieGetMessageQueryResultRequest)) *MockGenieInterface_GetMessageQueryResult_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dashboards.GenieGetMessageQueryResultRequest))
	})
	return _c
}

func (_c *MockGenieInterface_GetMessageQueryResult_Call) Return(_a0 *dashboards.GenieGetMessageQueryResultResponse, _a1 error) *MockGenieInterface_GetMessageQueryResult_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGenieInterface_GetMessageQueryResult_Call) RunAndReturn(run func(context.Context, dashboards.GenieGetMessageQueryResultRequest) (*dashboards.GenieGetMessageQueryResultResponse, error)) *MockGenieInterface_GetMessageQueryResult_Call {
	_c.Call.Return(run)
	return _c
}

// GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId provides a mock function with given fields: ctx, spaceId, conversationId, messageId
func (_m *MockGenieInterface) GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx context.Context, spaceId string, conversationId string, messageId string) (*dashboards.GenieGetMessageQueryResultResponse, error) {
	ret := _m.Called(ctx, spaceId, conversationId, messageId)

	if len(ret) == 0 {
		panic("no return value specified for GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId")
	}

	var r0 *dashboards.GenieGetMessageQueryResultResponse
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) (*dashboards.GenieGetMessageQueryResultResponse, error)); ok {
		return rf(ctx, spaceId, conversationId, messageId)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string) *dashboards.GenieGetMessageQueryResultResponse); ok {
		r0 = rf(ctx, spaceId, conversationId, messageId)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dashboards.GenieGetMessageQueryResultResponse)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string) error); ok {
		r1 = rf(ctx, spaceId, conversationId, messageId)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGenieInterface_GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId'
type MockGenieInterface_GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId_Call struct {
	*mock.Call
}

// GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId is a helper method to define mock.On call
//   - ctx context.Context
//   - spaceId string
//   - conversationId string
//   - messageId string
func (_e *MockGenieInterface_Expecter) GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId(ctx interface{}, spaceId interface{}, conversationId interface{}, messageId interface{}) *MockGenieInterface_GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId_Call {
	return &MockGenieInterface_GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId_Call{Call: _e.mock.On("GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId", ctx, spaceId, conversationId, messageId)}
}

func (_c *MockGenieInterface_GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId_Call) Run(run func(ctx context.Context, spaceId string, conversationId string, messageId string)) *MockGenieInterface_GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string))
	})
	return _c
}

func (_c *MockGenieInterface_GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId_Call) Return(_a0 *dashboards.GenieGetMessageQueryResultResponse, _a1 error) *MockGenieInterface_GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGenieInterface_GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId_Call) RunAndReturn(run func(context.Context, string, string, string) (*dashboards.GenieGetMessageQueryResultResponse, error)) *MockGenieInterface_GetMessageQueryResultBySpaceIdAndConversationIdAndMessageId_Call {
	_c.Call.Return(run)
	return _c
}

// Impl provides a mock function with given fields:
func (_m *MockGenieInterface) Impl() dashboards.GenieService {
	ret := _m.Called()

	if len(ret) == 0 {
		panic("no return value specified for Impl")
	}

	var r0 dashboards.GenieService
	if rf, ok := ret.Get(0).(func() dashboards.GenieService); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dashboards.GenieService)
		}
	}

	return r0
}

// MockGenieInterface_Impl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Impl'
type MockGenieInterface_Impl_Call struct {
	*mock.Call
}

// Impl is a helper method to define mock.On call
func (_e *MockGenieInterface_Expecter) Impl() *MockGenieInterface_Impl_Call {
	return &MockGenieInterface_Impl_Call{Call: _e.mock.On("Impl")}
}

func (_c *MockGenieInterface_Impl_Call) Run(run func()) *MockGenieInterface_Impl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run()
	})
	return _c
}

func (_c *MockGenieInterface_Impl_Call) Return(_a0 dashboards.GenieService) *MockGenieInterface_Impl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGenieInterface_Impl_Call) RunAndReturn(run func() dashboards.GenieService) *MockGenieInterface_Impl_Call {
	_c.Call.Return(run)
	return _c
}

// StartConversation provides a mock function with given fields: ctx, genieStartConversationMessageRequest
func (_m *MockGenieInterface) StartConversation(ctx context.Context, genieStartConversationMessageRequest dashboards.GenieStartConversationMessageRequest) (*dashboards.WaitGetMessageGenieCompleted[dashboards.GenieStartConversationResponse], error) {
	ret := _m.Called(ctx, genieStartConversationMessageRequest)

	if len(ret) == 0 {
		panic("no return value specified for StartConversation")
	}

	var r0 *dashboards.WaitGetMessageGenieCompleted[dashboards.GenieStartConversationResponse]
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GenieStartConversationMessageRequest) (*dashboards.WaitGetMessageGenieCompleted[dashboards.GenieStartConversationResponse], error)); ok {
		return rf(ctx, genieStartConversationMessageRequest)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GenieStartConversationMessageRequest) *dashboards.WaitGetMessageGenieCompleted[dashboards.GenieStartConversationResponse]); ok {
		r0 = rf(ctx, genieStartConversationMessageRequest)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dashboards.WaitGetMessageGenieCompleted[dashboards.GenieStartConversationResponse])
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dashboards.GenieStartConversationMessageRequest) error); ok {
		r1 = rf(ctx, genieStartConversationMessageRequest)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGenieInterface_StartConversation_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartConversation'
type MockGenieInterface_StartConversation_Call struct {
	*mock.Call
}

// StartConversation is a helper method to define mock.On call
//   - ctx context.Context
//   - genieStartConversationMessageRequest dashboards.GenieStartConversationMessageRequest
func (_e *MockGenieInterface_Expecter) StartConversation(ctx interface{}, genieStartConversationMessageRequest interface{}) *MockGenieInterface_StartConversation_Call {
	return &MockGenieInterface_StartConversation_Call{Call: _e.mock.On("StartConversation", ctx, genieStartConversationMessageRequest)}
}

func (_c *MockGenieInterface_StartConversation_Call) Run(run func(ctx context.Context, genieStartConversationMessageRequest dashboards.GenieStartConversationMessageRequest)) *MockGenieInterface_StartConversation_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(dashboards.GenieStartConversationMessageRequest))
	})
	return _c
}

func (_c *MockGenieInterface_StartConversation_Call) Return(_a0 *dashboards.WaitGetMessageGenieCompleted[dashboards.GenieStartConversationResponse], _a1 error) *MockGenieInterface_StartConversation_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGenieInterface_StartConversation_Call) RunAndReturn(run func(context.Context, dashboards.GenieStartConversationMessageRequest) (*dashboards.WaitGetMessageGenieCompleted[dashboards.GenieStartConversationResponse], error)) *MockGenieInterface_StartConversation_Call {
	_c.Call.Return(run)
	return _c
}

// StartConversationAndWait provides a mock function with given fields: ctx, genieStartConversationMessageRequest, options
func (_m *MockGenieInterface) StartConversationAndWait(ctx context.Context, genieStartConversationMessageRequest dashboards.GenieStartConversationMessageRequest, options ...retries.Option[dashboards.GenieMessage]) (*dashboards.GenieMessage, error) {
	_va := make([]interface{}, len(options))
	for _i := range options {
		_va[_i] = options[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx, genieStartConversationMessageRequest)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	if len(ret) == 0 {
		panic("no return value specified for StartConversationAndWait")
	}

	var r0 *dashboards.GenieMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GenieStartConversationMessageRequest, ...retries.Option[dashboards.GenieMessage]) (*dashboards.GenieMessage, error)); ok {
		return rf(ctx, genieStartConversationMessageRequest, options...)
	}
	if rf, ok := ret.Get(0).(func(context.Context, dashboards.GenieStartConversationMessageRequest, ...retries.Option[dashboards.GenieMessage]) *dashboards.GenieMessage); ok {
		r0 = rf(ctx, genieStartConversationMessageRequest, options...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dashboards.GenieMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, dashboards.GenieStartConversationMessageRequest, ...retries.Option[dashboards.GenieMessage]) error); ok {
		r1 = rf(ctx, genieStartConversationMessageRequest, options...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGenieInterface_StartConversationAndWait_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'StartConversationAndWait'
type MockGenieInterface_StartConversationAndWait_Call struct {
	*mock.Call
}

// StartConversationAndWait is a helper method to define mock.On call
//   - ctx context.Context
//   - genieStartConversationMessageRequest dashboards.GenieStartConversationMessageRequest
//   - options ...retries.Option[dashboards.GenieMessage]
func (_e *MockGenieInterface_Expecter) StartConversationAndWait(ctx interface{}, genieStartConversationMessageRequest interface{}, options ...interface{}) *MockGenieInterface_StartConversationAndWait_Call {
	return &MockGenieInterface_StartConversationAndWait_Call{Call: _e.mock.On("StartConversationAndWait",
		append([]interface{}{ctx, genieStartConversationMessageRequest}, options...)...)}
}

func (_c *MockGenieInterface_StartConversationAndWait_Call) Run(run func(ctx context.Context, genieStartConversationMessageRequest dashboards.GenieStartConversationMessageRequest, options ...retries.Option[dashboards.GenieMessage])) *MockGenieInterface_StartConversationAndWait_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]retries.Option[dashboards.GenieMessage], len(args)-2)
		for i, a := range args[2:] {
			if a != nil {
				variadicArgs[i] = a.(retries.Option[dashboards.GenieMessage])
			}
		}
		run(args[0].(context.Context), args[1].(dashboards.GenieStartConversationMessageRequest), variadicArgs...)
	})
	return _c
}

func (_c *MockGenieInterface_StartConversationAndWait_Call) Return(_a0 *dashboards.GenieMessage, _a1 error) *MockGenieInterface_StartConversationAndWait_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGenieInterface_StartConversationAndWait_Call) RunAndReturn(run func(context.Context, dashboards.GenieStartConversationMessageRequest, ...retries.Option[dashboards.GenieMessage]) (*dashboards.GenieMessage, error)) *MockGenieInterface_StartConversationAndWait_Call {
	_c.Call.Return(run)
	return _c
}

// WaitGetMessageGenieCompleted provides a mock function with given fields: ctx, conversationId, messageId, spaceId, timeout, callback
func (_m *MockGenieInterface) WaitGetMessageGenieCompleted(ctx context.Context, conversationId string, messageId string, spaceId string, timeout time.Duration, callback func(*dashboards.GenieMessage)) (*dashboards.GenieMessage, error) {
	ret := _m.Called(ctx, conversationId, messageId, spaceId, timeout, callback)

	if len(ret) == 0 {
		panic("no return value specified for WaitGetMessageGenieCompleted")
	}

	var r0 *dashboards.GenieMessage
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, time.Duration, func(*dashboards.GenieMessage)) (*dashboards.GenieMessage, error)); ok {
		return rf(ctx, conversationId, messageId, spaceId, timeout, callback)
	}
	if rf, ok := ret.Get(0).(func(context.Context, string, string, string, time.Duration, func(*dashboards.GenieMessage)) *dashboards.GenieMessage); ok {
		r0 = rf(ctx, conversationId, messageId, spaceId, timeout, callback)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*dashboards.GenieMessage)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, string, string, string, time.Duration, func(*dashboards.GenieMessage)) error); ok {
		r1 = rf(ctx, conversationId, messageId, spaceId, timeout, callback)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGenieInterface_WaitGetMessageGenieCompleted_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WaitGetMessageGenieCompleted'
type MockGenieInterface_WaitGetMessageGenieCompleted_Call struct {
	*mock.Call
}

// WaitGetMessageGenieCompleted is a helper method to define mock.On call
//   - ctx context.Context
//   - conversationId string
//   - messageId string
//   - spaceId string
//   - timeout time.Duration
//   - callback func(*dashboards.GenieMessage)
func (_e *MockGenieInterface_Expecter) WaitGetMessageGenieCompleted(ctx interface{}, conversationId interface{}, messageId interface{}, spaceId interface{}, timeout interface{}, callback interface{}) *MockGenieInterface_WaitGetMessageGenieCompleted_Call {
	return &MockGenieInterface_WaitGetMessageGenieCompleted_Call{Call: _e.mock.On("WaitGetMessageGenieCompleted", ctx, conversationId, messageId, spaceId, timeout, callback)}
}

func (_c *MockGenieInterface_WaitGetMessageGenieCompleted_Call) Run(run func(ctx context.Context, conversationId string, messageId string, spaceId string, timeout time.Duration, callback func(*dashboards.GenieMessage))) *MockGenieInterface_WaitGetMessageGenieCompleted_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(string), args[2].(string), args[3].(string), args[4].(time.Duration), args[5].(func(*dashboards.GenieMessage)))
	})
	return _c
}

func (_c *MockGenieInterface_WaitGetMessageGenieCompleted_Call) Return(_a0 *dashboards.GenieMessage, _a1 error) *MockGenieInterface_WaitGetMessageGenieCompleted_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGenieInterface_WaitGetMessageGenieCompleted_Call) RunAndReturn(run func(context.Context, string, string, string, time.Duration, func(*dashboards.GenieMessage)) (*dashboards.GenieMessage, error)) *MockGenieInterface_WaitGetMessageGenieCompleted_Call {
	_c.Call.Return(run)
	return _c
}

// WithImpl provides a mock function with given fields: impl
func (_m *MockGenieInterface) WithImpl(impl dashboards.GenieService) dashboards.GenieInterface {
	ret := _m.Called(impl)

	if len(ret) == 0 {
		panic("no return value specified for WithImpl")
	}

	var r0 dashboards.GenieInterface
	if rf, ok := ret.Get(0).(func(dashboards.GenieService) dashboards.GenieInterface); ok {
		r0 = rf(impl)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(dashboards.GenieInterface)
		}
	}

	return r0
}

// MockGenieInterface_WithImpl_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'WithImpl'
type MockGenieInterface_WithImpl_Call struct {
	*mock.Call
}

// WithImpl is a helper method to define mock.On call
//   - impl dashboards.GenieService
func (_e *MockGenieInterface_Expecter) WithImpl(impl interface{}) *MockGenieInterface_WithImpl_Call {
	return &MockGenieInterface_WithImpl_Call{Call: _e.mock.On("WithImpl", impl)}
}

func (_c *MockGenieInterface_WithImpl_Call) Run(run func(impl dashboards.GenieService)) *MockGenieInterface_WithImpl_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(dashboards.GenieService))
	})
	return _c
}

func (_c *MockGenieInterface_WithImpl_Call) Return(_a0 dashboards.GenieInterface) *MockGenieInterface_WithImpl_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockGenieInterface_WithImpl_Call) RunAndReturn(run func(dashboards.GenieService) dashboards.GenieInterface) *MockGenieInterface_WithImpl_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGenieInterface creates a new instance of MockGenieInterface. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGenieInterface(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGenieInterface {
	mock := &MockGenieInterface{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
