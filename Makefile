start:
	go run *.go deploy.json

build: clean
	go build -ldflags "-s -w" -o ./bin/deployer
	upx -9 ./bin/deployer

build-darwin: clean
	GOOS=darwin go build -o ./bin/deployer

build-android: clean
	GOOS=linux GOARCH=arm GOARM=7 go build -o ./bin/deployer

clean:
	rm -f ./bin/*
