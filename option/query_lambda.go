package option

import "github.com/getlantern/rockset-go-client/openapi"

type ExecuteQueryLambdaRequest struct {
	openapi.ExecuteQueryLambdaRequest
	Tag     string
	Version string
}

type QueryLambdaOption func(request *ExecuteQueryLambdaRequest)

func WithTag(tag string) QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		o.Tag = tag
	}
}

func WithVersion(version string) QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		o.Version = version
	}
}

func WithRowLimit2(limit int32) QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		o.DefaultRowLimit = &limit
	}
}

func WithWarnings2() QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		o.GenerateWarnings = openapi.PtrBool(true)
	}
}

func WithParameter2(name, valueType, value string) QueryLambdaOption {
	return func(o *ExecuteQueryLambdaRequest) {
		o.Parameters = append(o.Parameters, openapi.QueryParameter{
			Name:  name,
			Type:  valueType,
			Value: value,
		})
	}
}

type ListQueryLambdaOptions struct {
	Workspace *string
}

type ListQueryLambdaOption func(request *ListQueryLambdaOptions)

func WithQueryLambdaWorkspace(name string) ListQueryLambdaOption {
	return func(o *ListQueryLambdaOptions) {
		o.Workspace = &name
	}
}

type CreateQueryLambdaOptions struct {
	Description     *string
	QueryParameters []openapi.QueryParameter
}

type CreateQueryLambdaOption func(request *CreateQueryLambdaOptions)

func WithQueryLambdaDescription(desc string) CreateQueryLambdaOption {
	return func(o *CreateQueryLambdaOptions) {
		o.Description = &desc
	}
}

func WithDefaultParameter(name, paramType, value string) CreateQueryLambdaOption {
	return func(o *CreateQueryLambdaOptions) {
		o.QueryParameters = append(o.QueryParameters, openapi.QueryParameter{
			Name:  name,
			Type:  paramType,
			Value: value,
		})
	}
}
