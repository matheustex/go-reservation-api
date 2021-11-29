.PHONY: build clean deploy

build: 
	./gobuild.sh

clean:
	rm -rf ./bin ./vendor Gopkg.lock

deploy: clean build
	sls deploy --verbose