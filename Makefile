.PHONY: build
build:
	@ go build -o ./bin/server ./main.go

.PHONY: gen-db
gen-db:
	@ go generate ./gen/ent