clean:
	@rm -rf ./bin

dist-tools:
	@mkdir -p ./bin
	go get github.com/robfig/glock

deps:
	@glock sync github.com/alinz/react-native-updater/server

build:
	@mkdir -p ./bin
	cd cmd && GOGC=off go build -i -o ../bin/updater-server
