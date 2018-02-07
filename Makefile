.PHONY: test
test:
	go test -i -v .
	go test -v .

.PHONY: lint
lint:
	gometalinter --disable-all --enable=errcheck --enable=vet --enable=vetshadow --enable=golint --enable=gas --enable=ineffassign --enable=goconst --enable=goimports --enable=gofmt --exclude="should have comment" --enable=staticcheck --enable=gosimple --enable=misspell --deadline=60s .
