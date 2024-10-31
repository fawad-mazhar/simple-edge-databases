package models

type BucketPage struct {
	Buckets []string
}

type KeyValuePage struct {
	BucketName string
	Pairs      []KeyValue
}

type KeyValue struct {
	Key       string
	Value     string
	IsJSON    bool
	JSONValue interface{}
}
