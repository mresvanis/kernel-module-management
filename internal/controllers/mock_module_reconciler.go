// Code generated by MockGen. DO NOT EDIT.
// Source: module_reconciler.go
//
// Generated by this command:
//
//	mockgen -source=module_reconciler.go -package=controllers -destination=mock_module_reconciler.go moduleReconcilerHelperAPI
//
// Package controllers is a generated GoMock package.
package controllers

import (
	context "context"
	reflect "reflect"

	v1beta1 "github.com/kubernetes-sigs/kernel-module-management/api/v1beta1"
	api "github.com/kubernetes-sigs/kernel-module-management/internal/api"
	gomock "go.uber.org/mock/gomock"
	v1 "k8s.io/api/apps/v1"
	v10 "k8s.io/api/core/v1"
	types "k8s.io/apimachinery/pkg/types"
)

// MockmoduleReconcilerHelperAPI is a mock of moduleReconcilerHelperAPI interface.
type MockmoduleReconcilerHelperAPI struct {
	ctrl     *gomock.Controller
	recorder *MockmoduleReconcilerHelperAPIMockRecorder
}

// MockmoduleReconcilerHelperAPIMockRecorder is the mock recorder for MockmoduleReconcilerHelperAPI.
type MockmoduleReconcilerHelperAPIMockRecorder struct {
	mock *MockmoduleReconcilerHelperAPI
}

// NewMockmoduleReconcilerHelperAPI creates a new mock instance.
func NewMockmoduleReconcilerHelperAPI(ctrl *gomock.Controller) *MockmoduleReconcilerHelperAPI {
	mock := &MockmoduleReconcilerHelperAPI{ctrl: ctrl}
	mock.recorder = &MockmoduleReconcilerHelperAPIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockmoduleReconcilerHelperAPI) EXPECT() *MockmoduleReconcilerHelperAPIMockRecorder {
	return m.recorder
}

// garbageCollect mocks base method.
func (m *MockmoduleReconcilerHelperAPI) garbageCollect(ctx context.Context, mod *v1beta1.Module, mldMappings map[string]*api.ModuleLoaderData, existingDS []v1.DaemonSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "garbageCollect", ctx, mod, mldMappings, existingDS)
	ret0, _ := ret[0].(error)
	return ret0
}

// garbageCollect indicates an expected call of garbageCollect.
func (mr *MockmoduleReconcilerHelperAPIMockRecorder) garbageCollect(ctx, mod, mldMappings, existingDS any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "garbageCollect", reflect.TypeOf((*MockmoduleReconcilerHelperAPI)(nil).garbageCollect), ctx, mod, mldMappings, existingDS)
}

// getNodesListBySelector mocks base method.
func (m *MockmoduleReconcilerHelperAPI) getNodesListBySelector(ctx context.Context, mod *v1beta1.Module) ([]v10.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getNodesListBySelector", ctx, mod)
	ret0, _ := ret[0].([]v10.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getNodesListBySelector indicates an expected call of getNodesListBySelector.
func (mr *MockmoduleReconcilerHelperAPIMockRecorder) getNodesListBySelector(ctx, mod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getNodesListBySelector", reflect.TypeOf((*MockmoduleReconcilerHelperAPI)(nil).getNodesListBySelector), ctx, mod)
}

// getRelevantKernelMappingsAndNodes mocks base method.
func (m *MockmoduleReconcilerHelperAPI) getRelevantKernelMappingsAndNodes(ctx context.Context, mod *v1beta1.Module, targetedNodes []v10.Node) (map[string]*api.ModuleLoaderData, []v10.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getRelevantKernelMappingsAndNodes", ctx, mod, targetedNodes)
	ret0, _ := ret[0].(map[string]*api.ModuleLoaderData)
	ret1, _ := ret[1].([]v10.Node)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// getRelevantKernelMappingsAndNodes indicates an expected call of getRelevantKernelMappingsAndNodes.
func (mr *MockmoduleReconcilerHelperAPIMockRecorder) getRelevantKernelMappingsAndNodes(ctx, mod, targetedNodes any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getRelevantKernelMappingsAndNodes", reflect.TypeOf((*MockmoduleReconcilerHelperAPI)(nil).getRelevantKernelMappingsAndNodes), ctx, mod, targetedNodes)
}

// getRequestedModule mocks base method.
func (m *MockmoduleReconcilerHelperAPI) getRequestedModule(ctx context.Context, namespacedName types.NamespacedName) (*v1beta1.Module, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "getRequestedModule", ctx, namespacedName)
	ret0, _ := ret[0].(*v1beta1.Module)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// getRequestedModule indicates an expected call of getRequestedModule.
func (mr *MockmoduleReconcilerHelperAPIMockRecorder) getRequestedModule(ctx, namespacedName any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getRequestedModule", reflect.TypeOf((*MockmoduleReconcilerHelperAPI)(nil).getRequestedModule), ctx, namespacedName)
}

// handleBuild mocks base method.
func (m *MockmoduleReconcilerHelperAPI) handleBuild(ctx context.Context, mld *api.ModuleLoaderData) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleBuild", ctx, mld)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// handleBuild indicates an expected call of handleBuild.
func (mr *MockmoduleReconcilerHelperAPIMockRecorder) handleBuild(ctx, mld any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleBuild", reflect.TypeOf((*MockmoduleReconcilerHelperAPI)(nil).handleBuild), ctx, mld)
}

// handleDevicePlugin mocks base method.
func (m *MockmoduleReconcilerHelperAPI) handleDevicePlugin(ctx context.Context, mod *v1beta1.Module, existingModuleDS []v1.DaemonSet) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleDevicePlugin", ctx, mod, existingModuleDS)
	ret0, _ := ret[0].(error)
	return ret0
}

// handleDevicePlugin indicates an expected call of handleDevicePlugin.
func (mr *MockmoduleReconcilerHelperAPIMockRecorder) handleDevicePlugin(ctx, mod, existingModuleDS any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleDevicePlugin", reflect.TypeOf((*MockmoduleReconcilerHelperAPI)(nil).handleDevicePlugin), ctx, mod, existingModuleDS)
}

// handleSigning mocks base method.
func (m *MockmoduleReconcilerHelperAPI) handleSigning(ctx context.Context, mld *api.ModuleLoaderData) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "handleSigning", ctx, mld)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// handleSigning indicates an expected call of handleSigning.
func (mr *MockmoduleReconcilerHelperAPIMockRecorder) handleSigning(ctx, mld any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "handleSigning", reflect.TypeOf((*MockmoduleReconcilerHelperAPI)(nil).handleSigning), ctx, mld)
}

// setKMMOMetrics mocks base method.
func (m *MockmoduleReconcilerHelperAPI) setKMMOMetrics(ctx context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "setKMMOMetrics", ctx)
}

// setKMMOMetrics indicates an expected call of setKMMOMetrics.
func (mr *MockmoduleReconcilerHelperAPIMockRecorder) setKMMOMetrics(ctx any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "setKMMOMetrics", reflect.TypeOf((*MockmoduleReconcilerHelperAPI)(nil).setKMMOMetrics), ctx)
}
