check: test vet gofumpt misspell unconvert staticcheck ineffassign unparam gomod

test:
	go test ./...

vet: | test
	go vet ./...

staticcheck:
	go get honnef.co/go/tools/cmd/staticcheck
	go mod vendor
	staticcheck -checks all ./...

misspell:
	go get github.com/client9/misspell/cmd/misspell
	go mod vendor
	misspell \
	-locale GB \
	-error \
	*.md *.go
unconvert:
	go get github.com/mdempsky/unconvert
	go mod vendor
	unconvert -v ./...

ineffassign:
	go get github.com/gordonklaus/ineffassign
	go mod vendor
	ineffassign ./...

unparam:
	go get mvdan.cc/unparam
	go mod vendor
	unparam ./...

gofumpt:
	gofumpt -l -w .

gomod:
	go mod tidy
	go mod vendor

install-hooks:
	cp .githooks/* .git/hooks
