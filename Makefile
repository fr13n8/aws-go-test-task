.PHONY: build clean deploy

build:
	go get ./...
	go mod vendor
	go env -w GOARCH=amd64 && go env -w GOOS=linux && go env -w CGO_ENABLED=0 && go build -ldflags="-d -s -w" -a -tags netgo -installsuffix netgo -o bin/area api/area/main.go
	go env -w GOARCH=amd64 && go env -w GOOS=linux && go env -w CGO_ENABLED=0 && go build -ldflags="-d -s -w" -a -tags netgo -installsuffix netgo -o bin/distance api/distance/main.go

clean:
	rm -r ./bin ./vendor Gopkg.lock

deploy:
	serverless deploy --verbose
