module restprotobuff

go 1.22.5

require (
	github.com/golang/protobuf v1.5.4
	google.golang.org/protobuf v1.34.2
	restwithprotobuff.com v0.0.0-00010101000000-000000000000
)

require github.com/gorilla/mux v1.8.1 // indirect

replace restwithprotobuff.com => ./
