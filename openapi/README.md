# Go API client for openapi

Rockset's REST API allows for creating and managing all resources in Rockset. Each supported endpoint is documented below.

All requests must be authorized with a Rockset API key, which can be created in the [Rockset console](https://console.rockset.com). The API key must be provided as `ApiKey <api_key>` in the `Authorization` request header. For example:
```
Authorization: ApiKey aB35kDjg93J5nsf4GjwMeErAVd832F7ad4vhsW1S02kfZiab42sTsfW5Sxt25asT
```

All endpoints are only accessible via https.

Build something awesome!

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: v1
- Package version: 0.14.4
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/getlantern/rockset-go-client"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *https://api.rs2.usw2.rockset.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*APIKeysApi* | [**CreateApiKey**](docs/APIKeysApi.md#createapikey) | **Post** /v1/orgs/self/users/self/apikeys | Create API Key
*APIKeysApi* | [**DeleteApiKey**](docs/APIKeysApi.md#deleteapikey) | **Delete** /v1/orgs/self/users/{user}/apikeys/{name} | Delete API Key
*APIKeysApi* | [**GetApiKey**](docs/APIKeysApi.md#getapikey) | **Get** /v1/orgs/self/users/{user}/apikeys/{name} | Retrieve API Key
*APIKeysApi* | [**ListApiKeys**](docs/APIKeysApi.md#listapikeys) | **Get** /v1/orgs/self/users/{user}/apikeys | List API Keys
*APIKeysApi* | [**UpdateApiKey**](docs/APIKeysApi.md#updateapikey) | **Post** /v1/orgs/self/users/{user}/apikeys/{name} | Update API Key State
*AliasesApi* | [**CreateAlias**](docs/AliasesApi.md#createalias) | **Post** /v1/orgs/self/ws/{workspace}/aliases | Create Alias
*AliasesApi* | [**DeleteAlias**](docs/AliasesApi.md#deletealias) | **Delete** /v1/orgs/self/ws/{workspace}/aliases/{alias} | Delete Alias
*AliasesApi* | [**GetAlias**](docs/AliasesApi.md#getalias) | **Get** /v1/orgs/self/ws/{workspace}/aliases/{alias} | Retrieve Alias
*AliasesApi* | [**ListAliases**](docs/AliasesApi.md#listaliases) | **Get** /v1/orgs/self/aliases | List Aliases
*AliasesApi* | [**UpdateAlias**](docs/AliasesApi.md#updatealias) | **Post** /v1/orgs/self/ws/{workspace}/aliases/{alias} | Update Alias
*AliasesApi* | [**WorkspaceAliases**](docs/AliasesApi.md#workspacealiases) | **Get** /v1/orgs/self/ws/{workspace}/aliases | List Aliases in Workspace
*CollectionsApi* | [**CreateCollection**](docs/CollectionsApi.md#createcollection) | **Post** /v1/orgs/self/ws/{workspace}/collections | Create Collection
*CollectionsApi* | [**DeleteCollection**](docs/CollectionsApi.md#deletecollection) | **Delete** /v1/orgs/self/ws/{workspace}/collections/{collection} | Delete Collection
*CollectionsApi* | [**GetCollection**](docs/CollectionsApi.md#getcollection) | **Get** /v1/orgs/self/ws/{workspace}/collections/{collection} | Retrieve Collection
*CollectionsApi* | [**ListCollections**](docs/CollectionsApi.md#listcollections) | **Get** /v1/orgs/self/collections | List Collections
*CollectionsApi* | [**WorkspaceCollections**](docs/CollectionsApi.md#workspacecollections) | **Get** /v1/orgs/self/ws/{workspace}/collections | List Collections in Workspace
*CustomRolesApi* | [**CreateRole**](docs/CustomRolesApi.md#createrole) | **Post** /v1/orgs/self/roles | Create a Role
*CustomRolesApi* | [**DeleteRole**](docs/CustomRolesApi.md#deleterole) | **Delete** /v1/orgs/self/roles/{roleName} | Delete a Role
*CustomRolesApi* | [**GetRole**](docs/CustomRolesApi.md#getrole) | **Get** /v1/orgs/self/roles/{roleName} | Retrieve role
*CustomRolesApi* | [**ListRoles**](docs/CustomRolesApi.md#listroles) | **Get** /v1/orgs/self/roles | List Roles
*CustomRolesApi* | [**UpdateRole**](docs/CustomRolesApi.md#updaterole) | **Post** /v1/orgs/self/roles/{roleName} | Update a Role
*DocumentsApi* | [**AddDocuments**](docs/DocumentsApi.md#adddocuments) | **Post** /v1/orgs/self/ws/{workspace}/collections/{collection}/docs | Add Documents
*DocumentsApi* | [**DeleteDocuments**](docs/DocumentsApi.md#deletedocuments) | **Delete** /v1/orgs/self/ws/{workspace}/collections/{collection}/docs | Delete Documents
*DocumentsApi* | [**PatchDocuments**](docs/DocumentsApi.md#patchdocuments) | **Patch** /v1/orgs/self/ws/{workspace}/collections/{collection}/docs | Patch Documents
*IntegrationsApi* | [**CreateIntegration**](docs/IntegrationsApi.md#createintegration) | **Post** /v1/orgs/self/integrations | Create Integration
*IntegrationsApi* | [**DeleteIntegration**](docs/IntegrationsApi.md#deleteintegration) | **Delete** /v1/orgs/self/integrations/{integration} | Delete Integration
*IntegrationsApi* | [**GetIntegration**](docs/IntegrationsApi.md#getintegration) | **Get** /v1/orgs/self/integrations/{integration} | Retrieve Integration
*IntegrationsApi* | [**ListIntegrations**](docs/IntegrationsApi.md#listintegrations) | **Get** /v1/orgs/self/integrations | List Integrations
*OrganizationsApi* | [**GetOrganization**](docs/OrganizationsApi.md#getorganization) | **Get** /v1/orgs/self | Get Organization
*QueriesApi* | [**Query**](docs/QueriesApi.md#query) | **Post** /v1/orgs/self/queries | Query
*QueriesApi* | [**Validate**](docs/QueriesApi.md#validate) | **Post** /v1/orgs/self/queries/validations | Validate Query
*QueryLambdasApi* | [**CreateQueryLambda**](docs/QueryLambdasApi.md#createquerylambda) | **Post** /v1/orgs/self/ws/{workspace}/lambdas | Create Query Lambda
*QueryLambdasApi* | [**CreateQueryLambdaTag**](docs/QueryLambdasApi.md#createquerylambdatag) | **Post** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/tags | Create Query Lambda Tag
*QueryLambdasApi* | [**DeleteQueryLambda**](docs/QueryLambdasApi.md#deletequerylambda) | **Delete** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda} | Delete Query Lambda
*QueryLambdasApi* | [**DeleteQueryLambdaTag**](docs/QueryLambdasApi.md#deletequerylambdatag) | **Delete** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/tags/{tag} | Delete Query Lambda Tag Version
*QueryLambdasApi* | [**DeleteQueryLambdaVersion**](docs/QueryLambdasApi.md#deletequerylambdaversion) | **Delete** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/version/{version} | Delete Query Lambda Version
*QueryLambdasApi* | [**ExecuteQueryLambda**](docs/QueryLambdasApi.md#executequerylambda) | **Post** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/versions/{version} | Execute Query Lambda By Version
*QueryLambdasApi* | [**ExecuteQueryLambdaByTag**](docs/QueryLambdasApi.md#executequerylambdabytag) | **Post** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/tags/{tag} | Execute Query Lambda By Tag
*QueryLambdasApi* | [**GetQueryLambdaTagVersion**](docs/QueryLambdasApi.md#getquerylambdatagversion) | **Get** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/tags/{tag} | Retrieve Query Lambda Tag
*QueryLambdasApi* | [**GetQueryLambdaVersion**](docs/QueryLambdasApi.md#getquerylambdaversion) | **Get** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/versions/{version} | Retrieve Query Lambda Version
*QueryLambdasApi* | [**ListAllQueryLambdas**](docs/QueryLambdasApi.md#listallquerylambdas) | **Get** /v1/orgs/self/lambdas | List Query Lambdas
*QueryLambdasApi* | [**ListQueryLambdaTags**](docs/QueryLambdasApi.md#listquerylambdatags) | **Get** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/tags | List Query Lambda Tags
*QueryLambdasApi* | [**ListQueryLambdaVersions**](docs/QueryLambdasApi.md#listquerylambdaversions) | **Get** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/versions | List Query Lambda Versions
*QueryLambdasApi* | [**ListQueryLambdasInWorkspace**](docs/QueryLambdasApi.md#listquerylambdasinworkspace) | **Get** /v1/orgs/self/ws/{workspace}/lambdas | List Query Lambdas in Workspace
*QueryLambdasApi* | [**UpdateQueryLambda**](docs/QueryLambdasApi.md#updatequerylambda) | **Post** /v1/orgs/self/ws/{workspace}/lambdas/{queryLambda}/versions | Update Query Lambda
*UsersApi* | [**CreateUser**](docs/UsersApi.md#createuser) | **Post** /v1/orgs/self/users | Create User
*UsersApi* | [**DeleteUser**](docs/UsersApi.md#deleteuser) | **Delete** /v1/orgs/self/users/{user} | Delete User
*UsersApi* | [**GetCurrentUser**](docs/UsersApi.md#getcurrentuser) | **Get** /v1/orgs/self/users/self | Retrieve Current User
*UsersApi* | [**GetUser**](docs/UsersApi.md#getuser) | **Get** /v1/orgs/self/users/{user} | Retrieve User
*UsersApi* | [**ListUnsubscribePreferences**](docs/UsersApi.md#listunsubscribepreferences) | **Get** /v1/orgs/self/users/self/preferences | Retrieve Notification Preferences
*UsersApi* | [**ListUsers**](docs/UsersApi.md#listusers) | **Get** /v1/orgs/self/users | List Users
*UsersApi* | [**UpdateUnsubscribePreferences**](docs/UsersApi.md#updateunsubscribepreferences) | **Post** /v1/orgs/self/users/self/preferences | Update Notification Preferences
*ViewsApi* | [**CreateView**](docs/ViewsApi.md#createview) | **Post** /v1/orgs/self/ws/{workspace}/views | Create View
*ViewsApi* | [**DeleteView**](docs/ViewsApi.md#deleteview) | **Delete** /v1/orgs/self/ws/{workspace}/views/{view} | Delete View
*ViewsApi* | [**GetView**](docs/ViewsApi.md#getview) | **Get** /v1/orgs/self/ws/{workspace}/views/{view} | Retrieve View
*ViewsApi* | [**ListViews**](docs/ViewsApi.md#listviews) | **Get** /v1/orgs/self/views | List Views
*ViewsApi* | [**UpdateView**](docs/ViewsApi.md#updateview) | **Post** /v1/orgs/self/ws/{workspace}/views/{view} | Update View
*ViewsApi* | [**WorkspaceViews**](docs/ViewsApi.md#workspaceviews) | **Get** /v1/orgs/self/ws/{workspace}/views | List Views in Workspace
*VirtualInstancesApi* | [**GetVirtualInstance**](docs/VirtualInstancesApi.md#getvirtualinstance) | **Get** /v1/orgs/self/virtualinstances/{virtualInstanceId} | Retrieve Virtual Instance
*VirtualInstancesApi* | [**ListVirtualInstances**](docs/VirtualInstancesApi.md#listvirtualinstances) | **Get** /v1/orgs/self/virtualinstances | List Virtual Instances
*VirtualInstancesApi* | [**SetVirtualInstance**](docs/VirtualInstancesApi.md#setvirtualinstance) | **Post** /v1/orgs/self/virtualinstances/{virtualInstanceId} | Update Virtual Instance
*WorkspacesApi* | [**CreateWorkspace**](docs/WorkspacesApi.md#createworkspace) | **Post** /v1/orgs/self/ws | Create Workspace
*WorkspacesApi* | [**DeleteWorkspace**](docs/WorkspacesApi.md#deleteworkspace) | **Delete** /v1/orgs/self/ws/{workspace} | Delete Workspace
*WorkspacesApi* | [**GetWorkspace**](docs/WorkspacesApi.md#getworkspace) | **Get** /v1/orgs/self/ws/{workspace} | Retrieve Workspace
*WorkspacesApi* | [**ListWorkspaces**](docs/WorkspacesApi.md#listworkspaces) | **Get** /v1/orgs/self/ws | List Workspaces


## Documentation For Models

 - [AddDocumentsRequest](docs/AddDocumentsRequest.md)
 - [AddDocumentsResponse](docs/AddDocumentsResponse.md)
 - [Alias](docs/Alias.md)
 - [ApiKey](docs/ApiKey.md)
 - [AwsAccessKey](docs/AwsAccessKey.md)
 - [AwsRole](docs/AwsRole.md)
 - [AzureBlobStorageIntegration](docs/AzureBlobStorageIntegration.md)
 - [AzureEventHubsIntegration](docs/AzureEventHubsIntegration.md)
 - [AzureServiceBusIntegration](docs/AzureServiceBusIntegration.md)
 - [Cluster](docs/Cluster.md)
 - [Collection](docs/Collection.md)
 - [CollectionStats](docs/CollectionStats.md)
 - [CreateAliasRequest](docs/CreateAliasRequest.md)
 - [CreateAliasResponse](docs/CreateAliasResponse.md)
 - [CreateApiKeyRequest](docs/CreateApiKeyRequest.md)
 - [CreateApiKeyResponse](docs/CreateApiKeyResponse.md)
 - [CreateCollectionRequest](docs/CreateCollectionRequest.md)
 - [CreateCollectionResponse](docs/CreateCollectionResponse.md)
 - [CreateIntegrationRequest](docs/CreateIntegrationRequest.md)
 - [CreateIntegrationResponse](docs/CreateIntegrationResponse.md)
 - [CreateQueryLambdaRequest](docs/CreateQueryLambdaRequest.md)
 - [CreateQueryLambdaTagRequest](docs/CreateQueryLambdaTagRequest.md)
 - [CreateRoleRequest](docs/CreateRoleRequest.md)
 - [CreateUserRequest](docs/CreateUserRequest.md)
 - [CreateUserResponse](docs/CreateUserResponse.md)
 - [CreateViewRequest](docs/CreateViewRequest.md)
 - [CreateViewResponse](docs/CreateViewResponse.md)
 - [CreateWorkspaceRequest](docs/CreateWorkspaceRequest.md)
 - [CreateWorkspaceResponse](docs/CreateWorkspaceResponse.md)
 - [CsvParams](docs/CsvParams.md)
 - [DeleteAliasResponse](docs/DeleteAliasResponse.md)
 - [DeleteApiKeyResponse](docs/DeleteApiKeyResponse.md)
 - [DeleteCollectionResponse](docs/DeleteCollectionResponse.md)
 - [DeleteDocumentsRequest](docs/DeleteDocumentsRequest.md)
 - [DeleteDocumentsRequestData](docs/DeleteDocumentsRequestData.md)
 - [DeleteDocumentsResponse](docs/DeleteDocumentsResponse.md)
 - [DeleteIntegrationResponse](docs/DeleteIntegrationResponse.md)
 - [DeleteQueryLambdaResponse](docs/DeleteQueryLambdaResponse.md)
 - [DeleteUserResponse](docs/DeleteUserResponse.md)
 - [DeleteViewResponse](docs/DeleteViewResponse.md)
 - [DeleteWorkspaceResponse](docs/DeleteWorkspaceResponse.md)
 - [DocumentStatus](docs/DocumentStatus.md)
 - [DynamodbIntegration](docs/DynamodbIntegration.md)
 - [ErrorModel](docs/ErrorModel.md)
 - [EventTimeInfo](docs/EventTimeInfo.md)
 - [ExecuteQueryLambdaRequest](docs/ExecuteQueryLambdaRequest.md)
 - [FieldMappingQuery](docs/FieldMappingQuery.md)
 - [FieldMappingV2](docs/FieldMappingV2.md)
 - [FieldPartition](docs/FieldPartition.md)
 - [FormatParams](docs/FormatParams.md)
 - [GcpServiceAccount](docs/GcpServiceAccount.md)
 - [GcsIntegration](docs/GcsIntegration.md)
 - [GetAliasResponse](docs/GetAliasResponse.md)
 - [GetApiKeyResponse](docs/GetApiKeyResponse.md)
 - [GetCollectionResponse](docs/GetCollectionResponse.md)
 - [GetIntegrationResponse](docs/GetIntegrationResponse.md)
 - [GetViewResponse](docs/GetViewResponse.md)
 - [GetVirtualInstanceResponse](docs/GetVirtualInstanceResponse.md)
 - [GetWorkspaceResponse](docs/GetWorkspaceResponse.md)
 - [InputField](docs/InputField.md)
 - [Integration](docs/Integration.md)
 - [KafkaIntegration](docs/KafkaIntegration.md)
 - [KafkaV3SecurityConfig](docs/KafkaV3SecurityConfig.md)
 - [KinesisIntegration](docs/KinesisIntegration.md)
 - [ListAliasesResponse](docs/ListAliasesResponse.md)
 - [ListApiKeysResponse](docs/ListApiKeysResponse.md)
 - [ListCollectionsResponse](docs/ListCollectionsResponse.md)
 - [ListIntegrationsResponse](docs/ListIntegrationsResponse.md)
 - [ListQueryLambdaTagsResponse](docs/ListQueryLambdaTagsResponse.md)
 - [ListQueryLambdaVersionsResponse](docs/ListQueryLambdaVersionsResponse.md)
 - [ListQueryLambdasResponse](docs/ListQueryLambdasResponse.md)
 - [ListRolesResponse](docs/ListRolesResponse.md)
 - [ListUnsubscribePreferencesResponse](docs/ListUnsubscribePreferencesResponse.md)
 - [ListUsersResponse](docs/ListUsersResponse.md)
 - [ListViewsResponse](docs/ListViewsResponse.md)
 - [ListVirtualInstancesResponse](docs/ListVirtualInstancesResponse.md)
 - [ListWorkspacesResponse](docs/ListWorkspacesResponse.md)
 - [MongoDbIntegration](docs/MongoDbIntegration.md)
 - [Organization](docs/Organization.md)
 - [OrganizationResponse](docs/OrganizationResponse.md)
 - [OutputField](docs/OutputField.md)
 - [PaginationInfo](docs/PaginationInfo.md)
 - [PatchDocument](docs/PatchDocument.md)
 - [PatchDocumentsRequest](docs/PatchDocumentsRequest.md)
 - [PatchDocumentsResponse](docs/PatchDocumentsResponse.md)
 - [PatchOperation](docs/PatchOperation.md)
 - [Privilege](docs/Privilege.md)
 - [QueryError](docs/QueryError.md)
 - [QueryFieldType](docs/QueryFieldType.md)
 - [QueryLambda](docs/QueryLambda.md)
 - [QueryLambdaSql](docs/QueryLambdaSql.md)
 - [QueryLambdaStats](docs/QueryLambdaStats.md)
 - [QueryLambdaTag](docs/QueryLambdaTag.md)
 - [QueryLambdaTagResponse](docs/QueryLambdaTagResponse.md)
 - [QueryLambdaVersion](docs/QueryLambdaVersion.md)
 - [QueryLambdaVersionResponse](docs/QueryLambdaVersionResponse.md)
 - [QueryParameter](docs/QueryParameter.md)
 - [QueryRequest](docs/QueryRequest.md)
 - [QueryRequestSql](docs/QueryRequestSql.md)
 - [QueryResponse](docs/QueryResponse.md)
 - [QueryResponseStats](docs/QueryResponseStats.md)
 - [Role](docs/Role.md)
 - [RoleResponse](docs/RoleResponse.md)
 - [S3Integration](docs/S3Integration.md)
 - [SegmentIntegration](docs/SegmentIntegration.md)
 - [Source](docs/Source.md)
 - [SourceAzureBlobStorage](docs/SourceAzureBlobStorage.md)
 - [SourceAzureEventHubs](docs/SourceAzureEventHubs.md)
 - [SourceAzureServiceBus](docs/SourceAzureServiceBus.md)
 - [SourceDynamoDb](docs/SourceDynamoDb.md)
 - [SourceFileUpload](docs/SourceFileUpload.md)
 - [SourceGcs](docs/SourceGcs.md)
 - [SourceKafka](docs/SourceKafka.md)
 - [SourceKinesis](docs/SourceKinesis.md)
 - [SourceMongoDb](docs/SourceMongoDb.md)
 - [SourceS3](docs/SourceS3.md)
 - [SqlExpression](docs/SqlExpression.md)
 - [Status](docs/Status.md)
 - [StatusAzureEventHubs](docs/StatusAzureEventHubs.md)
 - [StatusAzureEventHubsPartition](docs/StatusAzureEventHubsPartition.md)
 - [StatusAzureServiceBus](docs/StatusAzureServiceBus.md)
 - [StatusAzureServiceBusSession](docs/StatusAzureServiceBusSession.md)
 - [StatusDynamoDb](docs/StatusDynamoDb.md)
 - [StatusDynamoDbV2](docs/StatusDynamoDbV2.md)
 - [StatusKafka](docs/StatusKafka.md)
 - [StatusKafkaPartition](docs/StatusKafkaPartition.md)
 - [StatusMongoDb](docs/StatusMongoDb.md)
 - [UnsubscribePreference](docs/UnsubscribePreference.md)
 - [UpdateAliasRequest](docs/UpdateAliasRequest.md)
 - [UpdateApiKeyRequest](docs/UpdateApiKeyRequest.md)
 - [UpdateApiKeyResponse](docs/UpdateApiKeyResponse.md)
 - [UpdateQueryLambdaRequest](docs/UpdateQueryLambdaRequest.md)
 - [UpdateRoleRequest](docs/UpdateRoleRequest.md)
 - [UpdateUnsubscribePreferencesRequest](docs/UpdateUnsubscribePreferencesRequest.md)
 - [UpdateUnsubscribePreferencesResponse](docs/UpdateUnsubscribePreferencesResponse.md)
 - [UpdateViewRequest](docs/UpdateViewRequest.md)
 - [UpdateViewResponse](docs/UpdateViewResponse.md)
 - [UpdateVirtualInstanceRequest](docs/UpdateVirtualInstanceRequest.md)
 - [UpdateVirtualInstanceResponse](docs/UpdateVirtualInstanceResponse.md)
 - [User](docs/User.md)
 - [ValidateQueryResponse](docs/ValidateQueryResponse.md)
 - [View](docs/View.md)
 - [VirtualInstance](docs/VirtualInstance.md)
 - [Workspace](docs/Workspace.md)
 - [XmlParams](docs/XmlParams.md)


## Documentation For Authorization

 Endpoints do not require authorization.


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



