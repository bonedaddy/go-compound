.PHONY: gen
gen:
	# make account proto
	protoc \
		-I=pb \
		-I=${GOPATH}/src \
		-I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		--gogo_out=plugins=grpc:pb \
		pb/account.proto 
	# make ctoken proto
	protoc \
		-I=pb \
		-I=${GOPATH}/src \
		-I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		--gogo_out=plugins=grpc:pb \
		pb/ctoken.proto 