module github.com/cloudwegox/thriftgo/test/golang/fieldmask

go 1.20

replace github.com/cloudwegox/thriftgo => ../../../.

replace github.com/apache/thrift => github.com/apache/thrift v0.13.0

require (
	github.com/apache/thrift v0.13.0
	github.com/cloudwegox/thriftgo v0.0.0-00010101000000-000000000000
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
