lint:
	golangci-lint run

lint-fix:
	goimports -w -l $(shell find . -type f -name '*.go' -not \( -path "./vendor/*" -o -path "*fakes/*" -o -path "*wrapped/*" -o -path "*changed_generated.go" \))
	make lint

test:
	ginkgo -r -tags=integration --trace --race

test-cover:
	ginkgo -r -tags=integration --trace --race --cover --coverprofile=.coverage.txt --outputdir=.
	go tool cover -func=.coverage.txt

install-tools-globally:
	@cat tools/tools.go | grep _ | awk -F'"' '{		\
package = $$2;							\
tags = $$3;							\
gsub("//"," ",tags);						\
print("go install", tags, " ", package)				\
}' | while read line ; do echo $$line; eval $$line ; done

generate:
	go generate ./...
