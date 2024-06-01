// Code generated by MockGen. DO NOT EDIT.
// Source: webook/api/proto/gen/article/v1/article_grpc.pb.go
//
// Generated by this command:
//
//	mockgen -source=webook/api/proto/gen/article/v1/article_grpc.pb.go -package=artmocks -destination=webook/api/proto/gen/article/v1/mocks/article_grpc.mock.go
//
// Package artmocks is a generated GoMock package.
package artmocks

import (
	context "context"
	reflect "reflect"

	articlev1 "gitee.com/geekbang/basic-go/webook/api/proto/gen/article/v1"
	gomock "go.uber.org/mock/gomock"
	grpc "google.golang.org/grpc"
)

// MockArticleServiceClient is a mock of ArticleServiceClient interface.
type MockArticleServiceClient struct {
	ctrl     *gomock.Controller
	recorder *MockArticleServiceClientMockRecorder
}

// MockArticleServiceClientMockRecorder is the mock recorder for MockArticleServiceClient.
type MockArticleServiceClientMockRecorder struct {
	mock *MockArticleServiceClient
}

// NewMockArticleServiceClient creates a new mock instance.
func NewMockArticleServiceClient(ctrl *gomock.Controller) *MockArticleServiceClient {
	mock := &MockArticleServiceClient{ctrl: ctrl}
	mock.recorder = &MockArticleServiceClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleServiceClient) EXPECT() *MockArticleServiceClientMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockArticleServiceClient) GetById(ctx context.Context, in *articlev1.GetByIdRequest, opts ...grpc.CallOption) (*articlev1.GetByIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetById", varargs...)
	ret0, _ := ret[0].(*articlev1.GetByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockArticleServiceClientMockRecorder) GetById(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockArticleServiceClient)(nil).GetById), varargs...)
}

// GetPublishedById mocks base method.
func (m *MockArticleServiceClient) GetPublishedById(ctx context.Context, in *articlev1.GetPublishedByIdRequest, opts ...grpc.CallOption) (*articlev1.GetPublishedByIdResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "GetPublishedById", varargs...)
	ret0, _ := ret[0].(*articlev1.GetPublishedByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublishedById indicates an expected call of GetPublishedById.
func (mr *MockArticleServiceClientMockRecorder) GetPublishedById(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublishedById", reflect.TypeOf((*MockArticleServiceClient)(nil).GetPublishedById), varargs...)
}

// List mocks base method.
func (m *MockArticleServiceClient) List(ctx context.Context, in *articlev1.ListRequest, opts ...grpc.CallOption) (*articlev1.ListResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "List", varargs...)
	ret0, _ := ret[0].(*articlev1.ListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockArticleServiceClientMockRecorder) List(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockArticleServiceClient)(nil).List), varargs...)
}

// ListPub mocks base method.
func (m *MockArticleServiceClient) ListPub(ctx context.Context, in *articlev1.ListPubRequest, opts ...grpc.CallOption) (*articlev1.ListPubResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "ListPub", varargs...)
	ret0, _ := ret[0].(*articlev1.ListPubResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPub indicates an expected call of ListPub.
func (mr *MockArticleServiceClientMockRecorder) ListPub(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPub", reflect.TypeOf((*MockArticleServiceClient)(nil).ListPub), varargs...)
}

// Publish mocks base method.
func (m *MockArticleServiceClient) Publish(ctx context.Context, in *articlev1.PublishRequest, opts ...grpc.CallOption) (*articlev1.PublishResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Publish", varargs...)
	ret0, _ := ret[0].(*articlev1.PublishResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockArticleServiceClientMockRecorder) Publish(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockArticleServiceClient)(nil).Publish), varargs...)
}

// Save mocks base method.
func (m *MockArticleServiceClient) Save(ctx context.Context, in *articlev1.SaveRequest, opts ...grpc.CallOption) (*articlev1.SaveResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Save", varargs...)
	ret0, _ := ret[0].(*articlev1.SaveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockArticleServiceClientMockRecorder) Save(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockArticleServiceClient)(nil).Save), varargs...)
}

// Withdraw mocks base method.
func (m *MockArticleServiceClient) Withdraw(ctx context.Context, in *articlev1.WithdrawRequest, opts ...grpc.CallOption) (*articlev1.WithdrawResponse, error) {
	m.ctrl.T.Helper()
	varargs := []any{ctx, in}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Withdraw", varargs...)
	ret0, _ := ret[0].(*articlev1.WithdrawResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockArticleServiceClientMockRecorder) Withdraw(ctx, in any, opts ...any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]any{ctx, in}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockArticleServiceClient)(nil).Withdraw), varargs...)
}

// MockArticleServiceServer is a mock of ArticleServiceServer interface.
type MockArticleServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockArticleServiceServerMockRecorder
}

// MockArticleServiceServerMockRecorder is the mock recorder for MockArticleServiceServer.
type MockArticleServiceServerMockRecorder struct {
	mock *MockArticleServiceServer
}

// NewMockArticleServiceServer creates a new mock instance.
func NewMockArticleServiceServer(ctrl *gomock.Controller) *MockArticleServiceServer {
	mock := &MockArticleServiceServer{ctrl: ctrl}
	mock.recorder = &MockArticleServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockArticleServiceServer) EXPECT() *MockArticleServiceServerMockRecorder {
	return m.recorder
}

// GetById mocks base method.
func (m *MockArticleServiceServer) GetById(arg0 context.Context, arg1 *articlev1.GetByIdRequest) (*articlev1.GetByIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetById", arg0, arg1)
	ret0, _ := ret[0].(*articlev1.GetByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetById indicates an expected call of GetById.
func (mr *MockArticleServiceServerMockRecorder) GetById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetById", reflect.TypeOf((*MockArticleServiceServer)(nil).GetById), arg0, arg1)
}

// GetPublishedById mocks base method.
func (m *MockArticleServiceServer) GetPublishedById(arg0 context.Context, arg1 *articlev1.GetPublishedByIdRequest) (*articlev1.GetPublishedByIdResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublishedById", arg0, arg1)
	ret0, _ := ret[0].(*articlev1.GetPublishedByIdResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublishedById indicates an expected call of GetPublishedById.
func (mr *MockArticleServiceServerMockRecorder) GetPublishedById(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublishedById", reflect.TypeOf((*MockArticleServiceServer)(nil).GetPublishedById), arg0, arg1)
}

// List mocks base method.
func (m *MockArticleServiceServer) List(arg0 context.Context, arg1 *articlev1.ListRequest) (*articlev1.ListResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "List", arg0, arg1)
	ret0, _ := ret[0].(*articlev1.ListResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// List indicates an expected call of List.
func (mr *MockArticleServiceServerMockRecorder) List(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "List", reflect.TypeOf((*MockArticleServiceServer)(nil).List), arg0, arg1)
}

// ListPub mocks base method.
func (m *MockArticleServiceServer) ListPub(arg0 context.Context, arg1 *articlev1.ListPubRequest) (*articlev1.ListPubResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPub", arg0, arg1)
	ret0, _ := ret[0].(*articlev1.ListPubResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPub indicates an expected call of ListPub.
func (mr *MockArticleServiceServerMockRecorder) ListPub(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPub", reflect.TypeOf((*MockArticleServiceServer)(nil).ListPub), arg0, arg1)
}

// Publish mocks base method.
func (m *MockArticleServiceServer) Publish(arg0 context.Context, arg1 *articlev1.PublishRequest) (*articlev1.PublishResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Publish", arg0, arg1)
	ret0, _ := ret[0].(*articlev1.PublishResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Publish indicates an expected call of Publish.
func (mr *MockArticleServiceServerMockRecorder) Publish(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Publish", reflect.TypeOf((*MockArticleServiceServer)(nil).Publish), arg0, arg1)
}

// Save mocks base method.
func (m *MockArticleServiceServer) Save(arg0 context.Context, arg1 *articlev1.SaveRequest) (*articlev1.SaveResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", arg0, arg1)
	ret0, _ := ret[0].(*articlev1.SaveResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Save indicates an expected call of Save.
func (mr *MockArticleServiceServerMockRecorder) Save(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockArticleServiceServer)(nil).Save), arg0, arg1)
}

// Withdraw mocks base method.
func (m *MockArticleServiceServer) Withdraw(arg0 context.Context, arg1 *articlev1.WithdrawRequest) (*articlev1.WithdrawResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", arg0, arg1)
	ret0, _ := ret[0].(*articlev1.WithdrawResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Withdraw indicates an expected call of Withdraw.
func (mr *MockArticleServiceServerMockRecorder) Withdraw(arg0, arg1 any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockArticleServiceServer)(nil).Withdraw), arg0, arg1)
}

// mustEmbedUnimplementedArticleServiceServer mocks base method.
func (m *MockArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedArticleServiceServer")
}

// mustEmbedUnimplementedArticleServiceServer indicates an expected call of mustEmbedUnimplementedArticleServiceServer.
func (mr *MockArticleServiceServerMockRecorder) mustEmbedUnimplementedArticleServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedArticleServiceServer", reflect.TypeOf((*MockArticleServiceServer)(nil).mustEmbedUnimplementedArticleServiceServer))
}

// MockUnsafeArticleServiceServer is a mock of UnsafeArticleServiceServer interface.
type MockUnsafeArticleServiceServer struct {
	ctrl     *gomock.Controller
	recorder *MockUnsafeArticleServiceServerMockRecorder
}

// MockUnsafeArticleServiceServerMockRecorder is the mock recorder for MockUnsafeArticleServiceServer.
type MockUnsafeArticleServiceServerMockRecorder struct {
	mock *MockUnsafeArticleServiceServer
}

// NewMockUnsafeArticleServiceServer creates a new mock instance.
func NewMockUnsafeArticleServiceServer(ctrl *gomock.Controller) *MockUnsafeArticleServiceServer {
	mock := &MockUnsafeArticleServiceServer{ctrl: ctrl}
	mock.recorder = &MockUnsafeArticleServiceServerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUnsafeArticleServiceServer) EXPECT() *MockUnsafeArticleServiceServerMockRecorder {
	return m.recorder
}

// mustEmbedUnimplementedArticleServiceServer mocks base method.
func (m *MockUnsafeArticleServiceServer) mustEmbedUnimplementedArticleServiceServer() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "mustEmbedUnimplementedArticleServiceServer")
}

// mustEmbedUnimplementedArticleServiceServer indicates an expected call of mustEmbedUnimplementedArticleServiceServer.
func (mr *MockUnsafeArticleServiceServerMockRecorder) mustEmbedUnimplementedArticleServiceServer() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "mustEmbedUnimplementedArticleServiceServer", reflect.TypeOf((*MockUnsafeArticleServiceServer)(nil).mustEmbedUnimplementedArticleServiceServer))
}
