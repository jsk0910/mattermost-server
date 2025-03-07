// Code generated by mockery v2.10.4. DO NOT EDIT.

// Regenerate this file using `make store-mocks`.

package mocks

import (
	model "github.com/mattermost/mattermost-server/v6/model"
	mock "github.com/stretchr/testify/mock"
)

// NotifyAdminStore is an autogenerated mock type for the NotifyAdminStore type
type NotifyAdminStore struct {
	mock.Mock
}

// DeleteBefore provides a mock function with given fields: trial, now
func (_m *NotifyAdminStore) DeleteBefore(trial bool, now int64) error {
	ret := _m.Called(trial, now)

	var r0 error
	if rf, ok := ret.Get(0).(func(bool, int64) error); ok {
		r0 = rf(trial, now)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: trial
func (_m *NotifyAdminStore) Get(trial bool) ([]*model.NotifyAdminData, error) {
	ret := _m.Called(trial)

	var r0 []*model.NotifyAdminData
	if rf, ok := ret.Get(0).(func(bool) []*model.NotifyAdminData); ok {
		r0 = rf(trial)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.NotifyAdminData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(bool) error); ok {
		r1 = rf(trial)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetDataByUserIdAndFeature provides a mock function with given fields: userId, feature
func (_m *NotifyAdminStore) GetDataByUserIdAndFeature(userId string, feature model.MattermostFeature) ([]*model.NotifyAdminData, error) {
	ret := _m.Called(userId, feature)

	var r0 []*model.NotifyAdminData
	if rf, ok := ret.Get(0).(func(string, model.MattermostFeature) []*model.NotifyAdminData); ok {
		r0 = rf(userId, feature)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*model.NotifyAdminData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, model.MattermostFeature) error); ok {
		r1 = rf(userId, feature)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Save provides a mock function with given fields: data
func (_m *NotifyAdminStore) Save(data *model.NotifyAdminData) (*model.NotifyAdminData, error) {
	ret := _m.Called(data)

	var r0 *model.NotifyAdminData
	if rf, ok := ret.Get(0).(func(*model.NotifyAdminData) *model.NotifyAdminData); ok {
		r0 = rf(data)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.NotifyAdminData)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*model.NotifyAdminData) error); ok {
		r1 = rf(data)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: userId, requiredPlan, requiredFeature, now
func (_m *NotifyAdminStore) Update(userId string, requiredPlan string, requiredFeature model.MattermostFeature, now int64) error {
	ret := _m.Called(userId, requiredPlan, requiredFeature, now)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, string, model.MattermostFeature, int64) error); ok {
		r0 = rf(userId, requiredPlan, requiredFeature, now)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
