DIR=/Users/sky0621/work/temp/testdata

.PHONY: build
build:
	docker build -t sky0621:kreu-crud-for-sqlc .

.PHONY: run
run:
	docker container run --rm --mount type=bind,src=${DIR},dst=/dir sky0621:kreu-crud-for-sqlc
