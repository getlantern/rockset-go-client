package option

import (
	"time"

	"github.com/getlantern/rockset-go-client/openapi"
)

type ListCollectionOptions struct {
	Workspace *string
}

type ListCollectionOption func(o *ListCollectionOptions)

func WithWorkspace(name string) func(o *ListCollectionOptions) {
	return func(o *ListCollectionOptions) {
		o.Workspace = &name
	}
}

type CollectionOption func(o *openapi.CreateCollectionRequest)

// WithCollectionRetention sets the retention in seconds for documents.
func WithCollectionRetention(d time.Duration) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		s := int64(d.Seconds())
		o.RetentionSecs = &s
	}
}

// WithDynamoDBMaxRCU sets the max RCU for a DynamoDB collection.
func WithDynamoDBMaxRCU(maxRCU int64) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		if o.Sources == nil || len(o.Sources) != 1 {
			return
		}
		if (o.Sources)[0].Dynamodb == nil {
			return
		}
		(o.Sources)[0].Dynamodb.Rcu = &maxRCU
	}
}

// WithFieldMappingQuery sets the field mapping query
func WithFieldMappingQuery(sql string) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		o.FieldMappingQuery = &openapi.FieldMappingQuery{
			Sql: &sql,
		}
	}
}

// WithInsertOnly enables insert only for the collection
func WithInsertOnly() CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		o.InsertOnly = openapi.PtrBool(true)
	}
}

type EventTimeInfoFormat string

const (
	MillisecondsSinceEpoch EventTimeInfoFormat = "milliseconds_since_epoch"
	SecondsSinceEpoch      EventTimeInfoFormat = "seconds_since_epoch"
)

// WithCollectionClusteringKey adds a clustering key. Can be specified multiple times.
func WithCollectionClusteringKey(fieldName, fieldType string, keys []string) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		if o.ClusteringKey == nil {
			o.ClusteringKey = []openapi.FieldPartition{}
		}

		o.ClusteringKey = append(o.ClusteringKey, openapi.FieldPartition{
			FieldName: &fieldName,
			Type:      &fieldType,
			Keys:      keys,
		})
	}
}

type FieldMissingAction string

func (f FieldMissingAction) String() string { return string(f) }

const (
	FieldMissingSkip FieldMissingAction = "SKIP"
	FieldMissingPass FieldMissingAction = "PASS"
)

type InputFieldFn func(field *openapi.InputField)

func InputField(fieldName string, ifMissing FieldMissingAction, drop bool, parameterName string) InputFieldFn {
	missing := ifMissing.String()
	return func(field *openapi.InputField) {
		field.FieldName = &fieldName
		field.IfMissing = &missing
		field.IsDrop = &drop
		field.Param = &parameterName
	}
}

type OnError string

func (e OnError) String() string { return string(e) }

const (
	OnErrorSkip OnError = "SKIP"
	OnErrorFail OnError = "FAIL"
)

type OutputFieldFn func(field *openapi.OutputField)

func OutputField(fieldName string, sql string, onError OnError) OutputFieldFn {
	return func(field *openapi.OutputField) {
		e := onError.String()
		field.FieldName = &fieldName
		field.Value = &openapi.SqlExpression{Sql: &sql}
		field.OnError = &e
	}
}

// WithCollectionFieldMapping adds a field mapping to the collection.
// If dropAll is true, the input and output fields are not set.
func WithCollectionFieldMapping(name string, dropAll bool, outputField OutputFieldFn,
	inputFields ...InputFieldFn) CollectionOption {
	return func(o *openapi.CreateCollectionRequest) {
		mapping := openapi.FieldMappingV2{
			Name: &name,
		}

		if dropAll {
			mapping.IsDropAllFields = &dropAll
		} else {
			out := openapi.OutputField{}
			outputField(&out)
			mapping.OutputField = &out

			inputs := make([]openapi.InputField, len(inputFields))
			for i, fn := range inputFields {
				var in openapi.InputField
				fn(&in)
				inputs[i] = in
			}
			mapping.InputFields = inputs
		}

		if o.FieldMappings == nil {
			o.FieldMappings = []openapi.FieldMappingV2{}
		}
		o.FieldMappings = append(o.FieldMappings, mapping)
	}
}

/*
	TODO: what about these two
	createParams.FieldSchemas
	createParams.InvertedIndexGroupEncodingOptions
*/
