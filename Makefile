mod:
	GO111MODULE=on go mod tidy

fmt:
	find . -name \*.go  | xargs goimports -w

lint:
	golangci-lint run -c .golangci.yml
	revive -exclude pkg/apis/... -exclude pkg/client/... -config .revive.toml -formatter friendly ./pkg/...


test:
	go test  -v  ./hash/...
