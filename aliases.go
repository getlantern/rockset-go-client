package rockset

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/getlantern/rockset-go-client/openapi"
	"github.com/getlantern/rockset-go-client/option"
)

// https://docs.rockset.com/rest-api/#aliases

// CreateAlias creates a new alias for a list of fully qualified collections, with an optional alias description
// using option.WithAliasDescription()
//
// REST API documentation https://docs.rockset.com/rest-api/#createalias
func (rc *RockClient) CreateAlias(ctx context.Context, workspace, alias string, collections []string,
	options ...option.AliasOption) (openapi.Alias, error) {
	log := zerolog.Ctx(ctx)
	var err error
	var resp *openapi.CreateAliasResponse

	q := rc.AliasesApi.CreateAlias(ctx, workspace)
	req := openapi.NewCreateAliasRequest(alias, collections)

	opts := option.AliasOptions{}
	for _, o := range options {
		o(&opts)
	}

	if opts.Description != nil {
		req.Description = opts.Description
	}

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Body(*req).Execute()
		return err
	})

	if err != nil {
		return openapi.Alias{}, err
	}

	log.Debug().Str("alias", alias).Strs("target", collections).Msg("alias created")

	return resp.GetData(), nil
}

// DeleteAlias deletes an alias.
//
// REST API documentation https://docs.rockset.com/rest-api/#deletealias
func (rc *RockClient) DeleteAlias(ctx context.Context, workspace, alias string) error {
	log := zerolog.Ctx(ctx)
	q := rc.AliasesApi.DeleteAlias(ctx, workspace, alias)

	err := rc.Retry(ctx, func() error {
		_, _, err := q.Execute()
		return err
	})

	log.Debug().Str("alias", alias).Msg("alias deleted")

	return err
}

// GetAlias gets an alias
//
// REST API documentation https://docs.rockset.com/rest-api/#getalias
func (rc *RockClient) GetAlias(ctx context.Context, workspace, alias string) (openapi.Alias, error) {
	var err error
	var resp *openapi.GetAliasResponse
	log := zerolog.Ctx(ctx)
	q := rc.AliasesApi.GetAlias(ctx, workspace, alias)

	err = rc.Retry(ctx, func() error {
		resp, _, err = q.Execute()
		return err
	})

	if err != nil {
		return openapi.Alias{}, err
	}

	log.Debug().Str("alias", alias).Msg("get alias")

	return resp.GetData(), nil
}

// ListAliases lists all aliases for the organization, or for a workspace when using the
// option.WithAliasWorkspace() option.
//
// REST API documentation https://docs.rockset.com/rest-api/#listalias
func (rc *RockClient) ListAliases(ctx context.Context, options ...option.ListAliasesOption) ([]openapi.Alias, error) {
	var err error
	var resp *openapi.ListAliasesResponse
	log := zerolog.Ctx(ctx)

	opts := option.ListAliasesOptions{}
	for _, o := range options {
		o(&opts)
	}

	if opts.Workspace == "" {
		q := rc.AliasesApi.ListAliases(ctx)
		err = rc.Retry(ctx, func() error {
			resp, _, err = q.Execute()
			return err
		})
	} else {
		q := rc.AliasesApi.WorkspaceAliases(ctx, opts.Workspace)
		err = rc.Retry(ctx, func() error {
			resp, _, err = q.Execute()
			return err
		})
	}

	if err != nil {
		return nil, err
	}

	log.Debug().Int("aliases", len(resp.GetData())).Msg("list alias")

	return resp.GetData(), nil
}

// UpdateAlias updates an alias
//
// REST API documentation https://docs.rockset.com/rest-api/#updatealias
func (rc *RockClient) UpdateAlias(ctx context.Context, workspace, alias string, collections []string,
	options ...option.AliasOption) error {
	var err error
	log := zerolog.Ctx(ctx)

	q := rc.AliasesApi.UpdateAlias(ctx, workspace, alias)
	req := openapi.NewUpdateAliasRequest(collections)

	opts := option.AliasOptions{}
	for _, o := range options {
		o(&opts)
	}
	if opts.Description != nil {
		req.Description = opts.Description
	}

	err = rc.Retry(ctx, func() error {
		_, _, err = q.Body(*req).Execute()

		return err
	})

	if err != nil {
		return err
	}

	log.Debug().Str("alias", alias).Msg("alias updated successfully")

	return nil
}
