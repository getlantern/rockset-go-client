# \WorkspacesApi

All URIs are relative to *https://api.rs2.usw2.rockset.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateWorkspace**](WorkspacesApi.md#CreateWorkspace) | **Post** /v1/orgs/self/ws | Create Workspace
[**DeleteWorkspace**](WorkspacesApi.md#DeleteWorkspace) | **Delete** /v1/orgs/self/ws/{workspace} | Delete Workspace
[**GetWorkspace**](WorkspacesApi.md#GetWorkspace) | **Get** /v1/orgs/self/ws/{workspace} | Retrieve Workspace
[**ListWorkspaces**](WorkspacesApi.md#ListWorkspaces) | **Get** /v1/orgs/self/ws | List Workspaces



## CreateWorkspace

> CreateWorkspaceResponse CreateWorkspace(ctx).Body(body).Execute()

Create Workspace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := *openapiclient.NewCreateWorkspaceRequest("event_logs") // CreateWorkspaceRequest | workspace details

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspacesApi.CreateWorkspace(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.CreateWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateWorkspace`: CreateWorkspaceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.CreateWorkspace`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateWorkspaceRequest**](CreateWorkspaceRequest.md) | workspace details | 

### Return type

[**CreateWorkspaceResponse**](CreateWorkspaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteWorkspace

> DeleteWorkspaceResponse DeleteWorkspace(ctx, workspace).Execute()

Delete Workspace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspacesApi.DeleteWorkspace(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.DeleteWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteWorkspace`: DeleteWorkspaceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.DeleteWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**DeleteWorkspaceResponse**](DeleteWorkspaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWorkspace

> GetWorkspaceResponse GetWorkspace(ctx, workspace).Execute()

Retrieve Workspace



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    workspace := "workspace_example" // string | name of the workspace (default to "commons")

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspacesApi.GetWorkspace(context.Background(), workspace).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.GetWorkspace``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetWorkspace`: GetWorkspaceResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.GetWorkspace`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**workspace** | **string** | name of the workspace | [default to &quot;commons&quot;]

### Other Parameters

Other parameters are passed through a pointer to a apiGetWorkspaceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetWorkspaceResponse**](GetWorkspaceResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListWorkspaces

> ListWorkspacesResponse ListWorkspaces(ctx).FetchAcrossRegions(fetchAcrossRegions).Execute()

List Workspaces



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    fetchAcrossRegions := true // bool |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.WorkspacesApi.ListWorkspaces(context.Background()).FetchAcrossRegions(fetchAcrossRegions).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `WorkspacesApi.ListWorkspaces``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListWorkspaces`: ListWorkspacesResponse
    fmt.Fprintf(os.Stdout, "Response from `WorkspacesApi.ListWorkspaces`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListWorkspacesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **fetchAcrossRegions** | **bool** |  | 

### Return type

[**ListWorkspacesResponse**](ListWorkspacesResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

