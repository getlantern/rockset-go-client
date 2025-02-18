# CreateIntegrationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** | descriptive label | 
**Description** | Pointer to **string** | longer explanation for the integration | [optional] 
**S3** | Pointer to [**S3Integration**](S3Integration.md) |  | [optional] 
**Kinesis** | Pointer to [**KinesisIntegration**](KinesisIntegration.md) |  | [optional] 
**Dynamodb** | Pointer to [**DynamodbIntegration**](DynamodbIntegration.md) |  | [optional] 
**Gcs** | Pointer to [**GcsIntegration**](GcsIntegration.md) |  | [optional] 
**AzureBlobStorage** | Pointer to [**AzureBlobStorageIntegration**](AzureBlobStorageIntegration.md) |  | [optional] 
**AzureEventHubs** | Pointer to [**AzureEventHubsIntegration**](AzureEventHubsIntegration.md) |  | [optional] 
**Segment** | Pointer to [**SegmentIntegration**](SegmentIntegration.md) |  | [optional] 
**Kafka** | Pointer to [**KafkaIntegration**](KafkaIntegration.md) |  | [optional] 
**Mongodb** | Pointer to [**MongoDbIntegration**](MongoDbIntegration.md) |  | [optional] 

## Methods

### NewCreateIntegrationRequest

`func NewCreateIntegrationRequest(name string, ) *CreateIntegrationRequest`

NewCreateIntegrationRequest instantiates a new CreateIntegrationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateIntegrationRequestWithDefaults

`func NewCreateIntegrationRequestWithDefaults() *CreateIntegrationRequest`

NewCreateIntegrationRequestWithDefaults instantiates a new CreateIntegrationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *CreateIntegrationRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CreateIntegrationRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CreateIntegrationRequest) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *CreateIntegrationRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CreateIntegrationRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CreateIntegrationRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CreateIntegrationRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetS3

`func (o *CreateIntegrationRequest) GetS3() S3Integration`

GetS3 returns the S3 field if non-nil, zero value otherwise.

### GetS3Ok

`func (o *CreateIntegrationRequest) GetS3Ok() (*S3Integration, bool)`

GetS3Ok returns a tuple with the S3 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetS3

`func (o *CreateIntegrationRequest) SetS3(v S3Integration)`

SetS3 sets S3 field to given value.

### HasS3

`func (o *CreateIntegrationRequest) HasS3() bool`

HasS3 returns a boolean if a field has been set.

### GetKinesis

`func (o *CreateIntegrationRequest) GetKinesis() KinesisIntegration`

GetKinesis returns the Kinesis field if non-nil, zero value otherwise.

### GetKinesisOk

`func (o *CreateIntegrationRequest) GetKinesisOk() (*KinesisIntegration, bool)`

GetKinesisOk returns a tuple with the Kinesis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKinesis

`func (o *CreateIntegrationRequest) SetKinesis(v KinesisIntegration)`

SetKinesis sets Kinesis field to given value.

### HasKinesis

`func (o *CreateIntegrationRequest) HasKinesis() bool`

HasKinesis returns a boolean if a field has been set.

### GetDynamodb

`func (o *CreateIntegrationRequest) GetDynamodb() DynamodbIntegration`

GetDynamodb returns the Dynamodb field if non-nil, zero value otherwise.

### GetDynamodbOk

`func (o *CreateIntegrationRequest) GetDynamodbOk() (*DynamodbIntegration, bool)`

GetDynamodbOk returns a tuple with the Dynamodb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDynamodb

`func (o *CreateIntegrationRequest) SetDynamodb(v DynamodbIntegration)`

SetDynamodb sets Dynamodb field to given value.

### HasDynamodb

`func (o *CreateIntegrationRequest) HasDynamodb() bool`

HasDynamodb returns a boolean if a field has been set.

### GetGcs

`func (o *CreateIntegrationRequest) GetGcs() GcsIntegration`

GetGcs returns the Gcs field if non-nil, zero value otherwise.

### GetGcsOk

`func (o *CreateIntegrationRequest) GetGcsOk() (*GcsIntegration, bool)`

GetGcsOk returns a tuple with the Gcs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGcs

`func (o *CreateIntegrationRequest) SetGcs(v GcsIntegration)`

SetGcs sets Gcs field to given value.

### HasGcs

`func (o *CreateIntegrationRequest) HasGcs() bool`

HasGcs returns a boolean if a field has been set.

### GetAzureBlobStorage

`func (o *CreateIntegrationRequest) GetAzureBlobStorage() AzureBlobStorageIntegration`

GetAzureBlobStorage returns the AzureBlobStorage field if non-nil, zero value otherwise.

### GetAzureBlobStorageOk

`func (o *CreateIntegrationRequest) GetAzureBlobStorageOk() (*AzureBlobStorageIntegration, bool)`

GetAzureBlobStorageOk returns a tuple with the AzureBlobStorage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureBlobStorage

`func (o *CreateIntegrationRequest) SetAzureBlobStorage(v AzureBlobStorageIntegration)`

SetAzureBlobStorage sets AzureBlobStorage field to given value.

### HasAzureBlobStorage

`func (o *CreateIntegrationRequest) HasAzureBlobStorage() bool`

HasAzureBlobStorage returns a boolean if a field has been set.

### GetAzureEventHubs

`func (o *CreateIntegrationRequest) GetAzureEventHubs() AzureEventHubsIntegration`

GetAzureEventHubs returns the AzureEventHubs field if non-nil, zero value otherwise.

### GetAzureEventHubsOk

`func (o *CreateIntegrationRequest) GetAzureEventHubsOk() (*AzureEventHubsIntegration, bool)`

GetAzureEventHubsOk returns a tuple with the AzureEventHubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureEventHubs

`func (o *CreateIntegrationRequest) SetAzureEventHubs(v AzureEventHubsIntegration)`

SetAzureEventHubs sets AzureEventHubs field to given value.

### HasAzureEventHubs

`func (o *CreateIntegrationRequest) HasAzureEventHubs() bool`

HasAzureEventHubs returns a boolean if a field has been set.

### GetSegment

`func (o *CreateIntegrationRequest) GetSegment() SegmentIntegration`

GetSegment returns the Segment field if non-nil, zero value otherwise.

### GetSegmentOk

`func (o *CreateIntegrationRequest) GetSegmentOk() (*SegmentIntegration, bool)`

GetSegmentOk returns a tuple with the Segment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSegment

`func (o *CreateIntegrationRequest) SetSegment(v SegmentIntegration)`

SetSegment sets Segment field to given value.

### HasSegment

`func (o *CreateIntegrationRequest) HasSegment() bool`

HasSegment returns a boolean if a field has been set.

### GetKafka

`func (o *CreateIntegrationRequest) GetKafka() KafkaIntegration`

GetKafka returns the Kafka field if non-nil, zero value otherwise.

### GetKafkaOk

`func (o *CreateIntegrationRequest) GetKafkaOk() (*KafkaIntegration, bool)`

GetKafkaOk returns a tuple with the Kafka field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKafka

`func (o *CreateIntegrationRequest) SetKafka(v KafkaIntegration)`

SetKafka sets Kafka field to given value.

### HasKafka

`func (o *CreateIntegrationRequest) HasKafka() bool`

HasKafka returns a boolean if a field has been set.

### GetMongodb

`func (o *CreateIntegrationRequest) GetMongodb() MongoDbIntegration`

GetMongodb returns the Mongodb field if non-nil, zero value otherwise.

### GetMongodbOk

`func (o *CreateIntegrationRequest) GetMongodbOk() (*MongoDbIntegration, bool)`

GetMongodbOk returns a tuple with the Mongodb field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMongodb

`func (o *CreateIntegrationRequest) SetMongodb(v MongoDbIntegration)`

SetMongodb sets Mongodb field to given value.

### HasMongodb

`func (o *CreateIntegrationRequest) HasMongodb() bool`

HasMongodb returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


