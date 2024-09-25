module grpcgatewayeg

go 1.22.5

require (
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240924160255-9d4c2d233b61
	google.golang.org/grpc v1.67.0
	google.golang.org/protobuf v1.34.2
)

require (
	golang.org/x/net v0.28.0 // indirect
	golang.org/x/sys v0.24.0 // indirect
	golang.org/x/text v0.18.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240924160255-9d4c2d233b61 // indirect

)

replace grpcgatewayeg => ./

replace github.com/grpc-ecosystem/grpc-gateway/v2 v2.22.0 => /Users/mahikajaguste/UL-Code/Soul-searching/grpc-gateway/
