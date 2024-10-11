.PHONY: build
build:
	@ go build -o ./bin/server ./main.go

.PHONY: gen-db
gen-db:
	@ go generate ./gen/ent

.PHONY: gen-mock
gen-mock:
	@ mockgen -source internal/db/intf.go -destination gen/mock/mock.go