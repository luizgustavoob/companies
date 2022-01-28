// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	context "context"

	entities "github.com/api-scanapi/internal/core/entities"
	mock "github.com/stretchr/testify/mock"
)

// CompanyClient is an autogenerated mock type for the CompanyClient type
type CompanyClient struct {
	mock.Mock
}

// IsValidCompany provides a mock function with given fields: ctx, company
func (_m *CompanyClient) IsValidCompany(ctx context.Context, company *entities.Company) (bool, error) {
	ret := _m.Called(ctx, company)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, *entities.Company) bool); ok {
		r0 = rf(ctx, company)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, *entities.Company) error); ok {
		r1 = rf(ctx, company)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}