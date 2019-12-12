.PHONY: bindings
bindings:
	abigen --abi abi/cbat.json --pkg bindings --out bindings/cbat/cbat.go
	abigen --abi abi/cdai.json --pkg bindings --out bindings/cdai/cdai.go
	abigen --abi abi/csai.json --pkg bindings --out bindings/csai/csai.go
	abigen --abi abi/ceth.json --pkg bindings --out bindings/ceth/ceth.go
	abigen --abi abi/crep.json --pkg bindings --out bindings/crep/crep.go
	abigen --abi abi/cusdc.json --pkg bindings --out bindings/cusdc/cusdc.go
	abigen --abi abi/cwbtc.json --pkg bindings --out bindings/cwbtc/cwbtc.go
	abigen --abi abi/czrx.json --pkg bindings --out bindings/czrx/czrx.go
	abigen --abi abi/comptroller.json --pkg bindings --out bindings/comptroller/comptroller.go
	abigen --abi abi/unitroller.json  --pkg bindings --out bindings/unitroller/unitroller.go

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