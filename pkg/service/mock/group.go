// Code generated by MockGen. DO NOT EDIT.
// Source: group.go

// Package mock is a generated GoMock package.
package mock

import (
	types "iam/pkg/service/types"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	sqlx "github.com/jmoiron/sqlx"
)

// MockGroupService is a mock of GroupService interface.
type MockGroupService struct {
	ctrl     *gomock.Controller
	recorder *MockGroupServiceMockRecorder
}

// MockGroupServiceMockRecorder is the mock recorder for MockGroupService.
type MockGroupServiceMockRecorder struct {
	mock *MockGroupService
}

// NewMockGroupService creates a new mock instance.
func NewMockGroupService(ctrl *gomock.Controller) *MockGroupService {
	mock := &MockGroupService{ctrl: ctrl}
	mock.recorder = &MockGroupServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGroupService) EXPECT() *MockGroupServiceMockRecorder {
	return m.recorder
}

// AlterGroupAuthType mocks base method.
func (m *MockGroupService) AlterGroupAuthType(tx *sqlx.Tx, systemID string, groupPK, authType int64) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AlterGroupAuthType", tx, systemID, groupPK, authType)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AlterGroupAuthType indicates an expected call of AlterGroupAuthType.
func (mr *MockGroupServiceMockRecorder) AlterGroupAuthType(tx, systemID, groupPK, authType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AlterGroupAuthType", reflect.TypeOf((*MockGroupService)(nil).AlterGroupAuthType), tx, systemID, groupPK, authType)
}

// BulkCreateGroupMembersWithTx mocks base method.
func (m *MockGroupService) BulkCreateGroupMembersWithTx(tx *sqlx.Tx, groupPK int64, relations []types.SubjectRelationForCreate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkCreateGroupMembersWithTx", tx, groupPK, relations)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkCreateGroupMembersWithTx indicates an expected call of BulkCreateGroupMembersWithTx.
func (mr *MockGroupServiceMockRecorder) BulkCreateGroupMembersWithTx(tx, groupPK, relations interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkCreateGroupMembersWithTx", reflect.TypeOf((*MockGroupService)(nil).BulkCreateGroupMembersWithTx), tx, groupPK, relations)
}

// BulkDeleteByGroupPKsWithTx mocks base method.
func (m *MockGroupService) BulkDeleteByGroupPKsWithTx(tx *sqlx.Tx, subjectPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteByGroupPKsWithTx", tx, subjectPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteByGroupPKsWithTx indicates an expected call of BulkDeleteByGroupPKsWithTx.
func (mr *MockGroupServiceMockRecorder) BulkDeleteByGroupPKsWithTx(tx, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteByGroupPKsWithTx", reflect.TypeOf((*MockGroupService)(nil).BulkDeleteByGroupPKsWithTx), tx, subjectPKs)
}

// BulkDeleteBySubjectPKsWithTx mocks base method.
func (m *MockGroupService) BulkDeleteBySubjectPKsWithTx(tx *sqlx.Tx, subjectPKs []int64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteBySubjectPKsWithTx", tx, subjectPKs)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkDeleteBySubjectPKsWithTx indicates an expected call of BulkDeleteBySubjectPKsWithTx.
func (mr *MockGroupServiceMockRecorder) BulkDeleteBySubjectPKsWithTx(tx, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteBySubjectPKsWithTx", reflect.TypeOf((*MockGroupService)(nil).BulkDeleteBySubjectPKsWithTx), tx, subjectPKs)
}

// BulkDeleteGroupMembers mocks base method.
func (m *MockGroupService) BulkDeleteGroupMembers(groupPK int64, userPKs, departmentPKs []int64) (map[string]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkDeleteGroupMembers", groupPK, userPKs, departmentPKs)
	ret0, _ := ret[0].(map[string]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// BulkDeleteGroupMembers indicates an expected call of BulkDeleteGroupMembers.
func (mr *MockGroupServiceMockRecorder) BulkDeleteGroupMembers(groupPK, userPKs, departmentPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkDeleteGroupMembers", reflect.TypeOf((*MockGroupService)(nil).BulkDeleteGroupMembers), groupPK, userPKs, departmentPKs)
}

// FilterGroupPKsHasMemberBeforeExpiredAt mocks base method.
func (m *MockGroupService) FilterGroupPKsHasMemberBeforeExpiredAt(groupPKs []int64, expiredAt int64) ([]int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FilterGroupPKsHasMemberBeforeExpiredAt", groupPKs, expiredAt)
	ret0, _ := ret[0].([]int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FilterGroupPKsHasMemberBeforeExpiredAt indicates an expected call of FilterGroupPKsHasMemberBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) FilterGroupPKsHasMemberBeforeExpiredAt(groupPKs, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FilterGroupPKsHasMemberBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).FilterGroupPKsHasMemberBeforeExpiredAt), groupPKs, expiredAt)
}

// GetExpiredAtBySubjectGroup mocks base method.
func (m *MockGroupService) GetExpiredAtBySubjectGroup(subjectPK, groupPK int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetExpiredAtBySubjectGroup", subjectPK, groupPK)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetExpiredAtBySubjectGroup indicates an expected call of GetExpiredAtBySubjectGroup.
func (mr *MockGroupServiceMockRecorder) GetExpiredAtBySubjectGroup(subjectPK, groupPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetExpiredAtBySubjectGroup", reflect.TypeOf((*MockGroupService)(nil).GetExpiredAtBySubjectGroup), subjectPK, groupPK)
}

// GetGroupMemberCount mocks base method.
func (m *MockGroupService) GetGroupMemberCount(groupPK int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMemberCount", groupPK)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupMemberCount indicates an expected call of GetGroupMemberCount.
func (mr *MockGroupServiceMockRecorder) GetGroupMemberCount(groupPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMemberCount", reflect.TypeOf((*MockGroupService)(nil).GetGroupMemberCount), groupPK)
}

// GetGroupMemberCountBeforeExpiredAt mocks base method.
func (m *MockGroupService) GetGroupMemberCountBeforeExpiredAt(groupPK, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupMemberCountBeforeExpiredAt", groupPK, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupMemberCountBeforeExpiredAt indicates an expected call of GetGroupMemberCountBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) GetGroupMemberCountBeforeExpiredAt(groupPK, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupMemberCountBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).GetGroupMemberCountBeforeExpiredAt), groupPK, expiredAt)
}

// GetGroupSubjectCountBeforeExpiredAt mocks base method.
func (m *MockGroupService) GetGroupSubjectCountBeforeExpiredAt(expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGroupSubjectCountBeforeExpiredAt", expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGroupSubjectCountBeforeExpiredAt indicates an expected call of GetGroupSubjectCountBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) GetGroupSubjectCountBeforeExpiredAt(expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGroupSubjectCountBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).GetGroupSubjectCountBeforeExpiredAt), expiredAt)
}

// GetSubjectGroupCountBeforeExpiredAt mocks base method.
func (m *MockGroupService) GetSubjectGroupCountBeforeExpiredAt(subjectPK, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectGroupCountBeforeExpiredAt", subjectPK, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectGroupCountBeforeExpiredAt indicates an expected call of GetSubjectGroupCountBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) GetSubjectGroupCountBeforeExpiredAt(subjectPK, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectGroupCountBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).GetSubjectGroupCountBeforeExpiredAt), subjectPK, expiredAt)
}

// GetSubjectSystemGroupCountBeforeExpiredAt mocks base method.
func (m *MockGroupService) GetSubjectSystemGroupCountBeforeExpiredAt(subjectPK int64, systemID string, expiredAt int64) (int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSubjectSystemGroupCountBeforeExpiredAt", subjectPK, systemID, expiredAt)
	ret0, _ := ret[0].(int64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetSubjectSystemGroupCountBeforeExpiredAt indicates an expected call of GetSubjectSystemGroupCountBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) GetSubjectSystemGroupCountBeforeExpiredAt(subjectPK, systemID, expiredAt interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSubjectSystemGroupCountBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).GetSubjectSystemGroupCountBeforeExpiredAt), subjectPK, systemID, expiredAt)
}

// ListEffectSubjectGroupsBySubjectPKGroupPKs mocks base method.
func (m *MockGroupService) ListEffectSubjectGroupsBySubjectPKGroupPKs(subjectPK int64, groupPKs []int64) ([]types.SubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEffectSubjectGroupsBySubjectPKGroupPKs", subjectPK, groupPKs)
	ret0, _ := ret[0].([]types.SubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEffectSubjectGroupsBySubjectPKGroupPKs indicates an expected call of ListEffectSubjectGroupsBySubjectPKGroupPKs.
func (mr *MockGroupServiceMockRecorder) ListEffectSubjectGroupsBySubjectPKGroupPKs(subjectPK, groupPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEffectSubjectGroupsBySubjectPKGroupPKs", reflect.TypeOf((*MockGroupService)(nil).ListEffectSubjectGroupsBySubjectPKGroupPKs), subjectPK, groupPKs)
}

// ListEffectThinSubjectGroups mocks base method.
func (m *MockGroupService) ListEffectThinSubjectGroups(systemID string, subjectPKs []int64) (map[int64][]types.ThinSubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEffectThinSubjectGroups", systemID, subjectPKs)
	ret0, _ := ret[0].(map[int64][]types.ThinSubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEffectThinSubjectGroups indicates an expected call of ListEffectThinSubjectGroups.
func (mr *MockGroupServiceMockRecorder) ListEffectThinSubjectGroups(systemID, subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEffectThinSubjectGroups", reflect.TypeOf((*MockGroupService)(nil).ListEffectThinSubjectGroups), systemID, subjectPKs)
}

// ListEffectThinSubjectGroupsBySubjectPKs mocks base method.
func (m *MockGroupService) ListEffectThinSubjectGroupsBySubjectPKs(subjectPKs []int64) ([]types.ThinSubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEffectThinSubjectGroupsBySubjectPKs", subjectPKs)
	ret0, _ := ret[0].([]types.ThinSubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEffectThinSubjectGroupsBySubjectPKs indicates an expected call of ListEffectThinSubjectGroupsBySubjectPKs.
func (mr *MockGroupServiceMockRecorder) ListEffectThinSubjectGroupsBySubjectPKs(subjectPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEffectThinSubjectGroupsBySubjectPKs", reflect.TypeOf((*MockGroupService)(nil).ListEffectThinSubjectGroupsBySubjectPKs), subjectPKs)
}

// ListGroupAuthBySystemGroupPKs mocks base method.
func (m *MockGroupService) ListGroupAuthBySystemGroupPKs(systemID string, groupPKs []int64) ([]types.GroupAuthType, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupAuthBySystemGroupPKs", systemID, groupPKs)
	ret0, _ := ret[0].([]types.GroupAuthType)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupAuthBySystemGroupPKs indicates an expected call of ListGroupAuthBySystemGroupPKs.
func (mr *MockGroupServiceMockRecorder) ListGroupAuthBySystemGroupPKs(systemID, groupPKs interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupAuthBySystemGroupPKs", reflect.TypeOf((*MockGroupService)(nil).ListGroupAuthBySystemGroupPKs), systemID, groupPKs)
}

// ListGroupAuthSystemIDs mocks base method.
func (m *MockGroupService) ListGroupAuthSystemIDs(groupPK int64) ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupAuthSystemIDs", groupPK)
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupAuthSystemIDs indicates an expected call of ListGroupAuthSystemIDs.
func (mr *MockGroupServiceMockRecorder) ListGroupAuthSystemIDs(groupPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupAuthSystemIDs", reflect.TypeOf((*MockGroupService)(nil).ListGroupAuthSystemIDs), groupPK)
}

// ListGroupMember mocks base method.
func (m *MockGroupService) ListGroupMember(groupPK int64) ([]types.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListGroupMember", groupPK)
	ret0, _ := ret[0].([]types.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListGroupMember indicates an expected call of ListGroupMember.
func (mr *MockGroupServiceMockRecorder) ListGroupMember(groupPK interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListGroupMember", reflect.TypeOf((*MockGroupService)(nil).ListGroupMember), groupPK)
}

// ListPagingGroupMember mocks base method.
func (m *MockGroupService) ListPagingGroupMember(groupPK, limit, offset int64) ([]types.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingGroupMember", groupPK, limit, offset)
	ret0, _ := ret[0].([]types.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingGroupMember indicates an expected call of ListPagingGroupMember.
func (mr *MockGroupServiceMockRecorder) ListPagingGroupMember(groupPK, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingGroupMember", reflect.TypeOf((*MockGroupService)(nil).ListPagingGroupMember), groupPK, limit, offset)
}

// ListPagingGroupMemberBeforeExpiredAt mocks base method.
func (m *MockGroupService) ListPagingGroupMemberBeforeExpiredAt(groupPK, expiredAt, limit, offset int64) ([]types.GroupMember, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingGroupMemberBeforeExpiredAt", groupPK, expiredAt, limit, offset)
	ret0, _ := ret[0].([]types.GroupMember)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingGroupMemberBeforeExpiredAt indicates an expected call of ListPagingGroupMemberBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) ListPagingGroupMemberBeforeExpiredAt(groupPK, expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingGroupMemberBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).ListPagingGroupMemberBeforeExpiredAt), groupPK, expiredAt, limit, offset)
}

// ListPagingGroupSubjectBeforeExpiredAt mocks base method.
func (m *MockGroupService) ListPagingGroupSubjectBeforeExpiredAt(expiredAt, limit, offset int64) ([]types.GroupSubject, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingGroupSubjectBeforeExpiredAt", expiredAt, limit, offset)
	ret0, _ := ret[0].([]types.GroupSubject)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingGroupSubjectBeforeExpiredAt indicates an expected call of ListPagingGroupSubjectBeforeExpiredAt.
func (mr *MockGroupServiceMockRecorder) ListPagingGroupSubjectBeforeExpiredAt(expiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingGroupSubjectBeforeExpiredAt", reflect.TypeOf((*MockGroupService)(nil).ListPagingGroupSubjectBeforeExpiredAt), expiredAt, limit, offset)
}

// ListPagingSubjectGroups mocks base method.
func (m *MockGroupService) ListPagingSubjectGroups(subjectPK, beforeExpiredAt, limit, offset int64) ([]types.SubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingSubjectGroups", subjectPK, beforeExpiredAt, limit, offset)
	ret0, _ := ret[0].([]types.SubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingSubjectGroups indicates an expected call of ListPagingSubjectGroups.
func (mr *MockGroupServiceMockRecorder) ListPagingSubjectGroups(subjectPK, beforeExpiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingSubjectGroups", reflect.TypeOf((*MockGroupService)(nil).ListPagingSubjectGroups), subjectPK, beforeExpiredAt, limit, offset)
}

// ListPagingSubjectSystemGroups mocks base method.
func (m *MockGroupService) ListPagingSubjectSystemGroups(subjectPK int64, systemID string, beforeExpiredAt, limit, offset int64) ([]types.SubjectGroup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPagingSubjectSystemGroups", subjectPK, systemID, beforeExpiredAt, limit, offset)
	ret0, _ := ret[0].([]types.SubjectGroup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPagingSubjectSystemGroups indicates an expected call of ListPagingSubjectSystemGroups.
func (mr *MockGroupServiceMockRecorder) ListPagingSubjectSystemGroups(subjectPK, systemID, beforeExpiredAt, limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPagingSubjectSystemGroups", reflect.TypeOf((*MockGroupService)(nil).ListPagingSubjectSystemGroups), subjectPK, systemID, beforeExpiredAt, limit, offset)
}

// UpdateGroupMembersExpiredAtWithTx mocks base method.
func (m *MockGroupService) UpdateGroupMembersExpiredAtWithTx(tx *sqlx.Tx, groupPK int64, members []types.SubjectRelationForUpdate) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateGroupMembersExpiredAtWithTx", tx, groupPK, members)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateGroupMembersExpiredAtWithTx indicates an expected call of UpdateGroupMembersExpiredAtWithTx.
func (mr *MockGroupServiceMockRecorder) UpdateGroupMembersExpiredAtWithTx(tx, groupPK, members interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateGroupMembersExpiredAtWithTx", reflect.TypeOf((*MockGroupService)(nil).UpdateGroupMembersExpiredAtWithTx), tx, groupPK, members)
}
