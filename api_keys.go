package rockset

import (
	"context"

	"github.com/getlantern/rockset-go-client/openapi"
	"github.com/getlantern/rockset-go-client/option"
)

// CreateAPIKey creates a new API key for the current user with the specified, with an optional role.
//
// REST API documentation https://docs.rockset.com/rest-api/#createapikey
func (rc *RockClient) CreateAPIKey(ctx context.Context, keyName string,
	options ...option.APIKeyRoleOption) (openapi.ApiKey, error) {
	var err error
	var resp *openapi.CreateApiKeyResponse

	createReq := rc.APIKeysApi.CreateApiKey(ctx)
	b := openapi.NewCreateApiKeyRequest(keyName)

	var opts option.APIKeyRoleOptions
	for _, o := range options {
		o(&opts)
	}

	if opts.Role != nil {
		b.Role = opts.Role
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = createReq.Body(*b).Execute()
		return err
	})

	if err != nil {
		return openapi.ApiKey{}, err
	}

	return resp.GetData(), nil
}

const self = "self"

// GetAPIKey gets the named API key.
// An admin can get an api key for another user with option.ForUser().
//
// REST API documentation https://docs.rockset.com/rest-api/#getapikey
func (rc *RockClient) GetAPIKey(ctx context.Context, name string,
	options ...option.APIKeyOption) (openapi.ApiKey, error) {
	var err error
	var resp *openapi.GetApiKeyResponse

	opts := option.APIKeyOptions{}
	for _, o := range options {
		o(&opts)
	}

	user := self
	if opts.User != nil {
		user = *opts.User
	}

	getReq := rc.APIKeysApi.GetApiKey(ctx, user, name)

	err = rc.Retry(ctx, func() error {
		resp, _, err = getReq.Execute()
		return err
	})

	if err != nil {
		return openapi.ApiKey{}, err
	}

	return resp.GetData(), nil
}

// DeleteAPIKey deletes an API key.
// An admin can delete an api key for another user with option.ForUser().
//
// REST API documentation https://docs.rockset.com/rest-api/#deleteapikey
func (rc *RockClient) DeleteAPIKey(ctx context.Context, keyName string, options ...option.APIKeyOption) error {
	var err error

	opts := option.APIKeyOptions{}
	for _, o := range options {
		o(&opts)
	}

	user := self
	if opts.User != nil {
		user = *opts.User
	}

	getReq := rc.APIKeysApi.DeleteApiKey(ctx, keyName, user)

	err = rc.Retry(ctx, func() error {
		_, _, err = getReq.Execute()
		return err
	})

	if err != nil {
		return err
	}

	return nil
}

// ListAPIKeys list API keys.
// An admin can list api keys for another user with option.ForUser().
//
// REST API documentation https://docs.rockset.com/rest-api/#listapikey
func (rc *RockClient) ListAPIKeys(ctx context.Context, options ...option.APIKeyOption) ([]openapi.ApiKey, error) {
	var err error
	var resp *openapi.ListApiKeysResponse

	opts := option.APIKeyOptions{}
	for _, o := range options {
		o(&opts)
	}

	user := self
	if opts.User != nil {
		user = *opts.User
	}

	getReq := rc.APIKeysApi.ListApiKeys(ctx, user)

	err = rc.Retry(ctx, func() error {
		resp, _, err = getReq.Execute()
		return err
	})

	if err != nil {
		return nil, err
	}

	return resp.GetData(), nil
}
