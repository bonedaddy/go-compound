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
	# make common proto
	protoc \
		-I=pb \
		-I=${GOPATH}/src \
		-I=${GOPATH}/src/github.com/gogo/protobuf/protobuf \
		--gogo_out=plugins=grpc:pb \
		pb/common.proto 


# run standard go tooling for better readability
.PHONY: tidy
tidy: imports fmt
	go vet ./...
	golint ./...

# automatically add missing imports
.PHONY: imports
imports:
	find . -type f -name '*.go' -exec goimports -w {} \;

# format code and simplify if possible
.PHONY: fmt
fmt:
	find . -type f -name '*.go' -exec gofmt -s -w {} \;

.PHONY: omit
omit:
	sed -i 's/\"`/,omitempty\"\`/g' models/account.go
	sed -i 's/\"`/,omitempty\"\`/g' models/ctoken.go 