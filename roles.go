package rockset

import (
	"context"

	"github.com/getlantern/rockset-go-client/openapi"
	"github.com/getlantern/rockset-go-client/option"
)

// CreateRole creates a new role
//
// REST API documentation https://docs.rockset.com/rest-api/#createrole
func (rc *RockClient) CreateRole(ctx context.Context, roleName string,
	options ...option.RoleOption) (openapi.Role, error) {
	var err error
	var resp *openapi.RoleResponse

	createReq := rc.CustomRolesApi.CreateRole(ctx)
	b := openapi.NewCreateRoleRequest()
	b.RoleName = &roleName

	var opts option.RoleOptions
	for _, o := range options {
		o(&opts)
	}

	if opts.Description != nil {
		b.Description = opts.Description
	}
	if len(opts.Privileges) > 0 {
		b.Privileges = opts.Privileges
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = createReq.Body(*b).Execute()
		return err
	})

	if err != nil {
		return openapi.Role{}, err
	}

	return resp.GetData(), nil
}

// UpdateRole updates a role.
//
// REST API documentation https://docs.rockset.com/rest-api/#updaterole
func (rc *RockClient) UpdateRole(ctx context.Context, roleName string,
	options ...option.RoleOption) (openapi.Role, error) {
	var err error
	var resp *openapi.RoleResponse

	createReq := rc.CustomRolesApi.UpdateRole(ctx, roleName)
	b := openapi.NewUpdateRoleRequest()

	var opts option.RoleOptions
	for _, o := range options {
		o(&opts)
	}

	if opts.Description != nil {
		b.Description = opts.Description
	}
	if len(opts.Privileges) > 0 {
		b.Privileges = opts.Privileges
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = createReq.Body(*b).Execute()
		return err
	})

	if err != nil {
		return openapi.Role{}, err
	}

	return resp.GetData(), nil
}

// DeleteRole deletes a role.
//
// REST API documentation https://docs.rockset.com/rest-api/#deleterole
func (rc *RockClient) DeleteRole(ctx context.Context, roleName string) error {
	var err error

	getReq := rc.CustomRolesApi.DeleteRole(ctx, roleName)

	err = rc.Retry(ctx, func() error {
		_, _, err = getReq.Execute()
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

// GetRole retrieve a role.
//
// REST API documentation https://docs.rockset.com/rest-api/#getrole
func (rc *RockClient) GetRole(ctx context.Context, roleName string) (openapi.Role, error) {
	var err error
	var resp *openapi.RoleResponse

	getReq := rc.CustomRolesApi.GetRole(ctx, roleName)

	err = rc.Retry(ctx, func() error {
		resp, _, err = getReq.Execute()
		return err
	})

	if err != nil {
		return openapi.Role{}, err
	}

	return resp.GetData(), nil
}

// ListRoles list all roles.
//
// REST API documentation https://docs.rockset.com/rest-api/#listroles
func (rc *RockClient) ListRoles(ctx context.Context) ([]openapi.Role, error) {
	var err error
	var resp *openapi.ListRolesResponse

	getReq := rc.CustomRolesApi.ListRoles(ctx)

	err = rc.Retry(ctx, func() error {
		resp, _, err = getReq.Execute()
		return err
	})

	if err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}
