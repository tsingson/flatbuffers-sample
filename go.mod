module github.com/tsingson/flatbuffers-sample

go 1.11

require (
	github.com/google/flatbuffers v1.12.0
	github.com/stretchr/testify v1.5.1
	github.com/tsingson/goflatbuffers v0.0.0-20200420205847-de8b4cd9cf67
	golang.org/x/lint v0.0.0-20190313153728-d0100b6bd8b3 // indirect
	golang.org/x/tools v0.0.0-20190524140312-2c0ae7006135 // indirect
	honnef.co/go/tools v0.0.0-20190523083050-ea95bdfd59fc // indirect
)

replace github.com/tsingson/goflatbuffers v0.0.0-20200420205847-de8b4cd9cf67 => /Users/qinshen/go/src/goflatbuffers

replace github.com/google/flatbuffers v1.12.0 => /Users/qinshen/go/src/flatbuffers
