PROJECT=csv2json

.PHONY=all get-deps build push

all: clean .build build

.build:
	mkdir -p .build/bin

clean:
	rm -rf .build

build:
	cp -R src .build/
	docker run --rm -v "$(CURDIR)/.build":/opt/workspace -w /opt/workspace/src golang:1.3.1 go build -v -o csv2json
	mv $(CURDIR)/.build/src/csv2json $(CURDIR)/.build/bin
	docker build -t pseiffert/$(PROJECT) .

push:
	docker push pseiffert/$(PROJECT)
