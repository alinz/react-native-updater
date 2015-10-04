clean:
	@rm -rf ./bin

dist-tools:
	go get github.com/pilu/fresh
	go get github.com/robfig/glock

deps:
	@glock sync github.com/alinz/react-native-updater/server

build:
	@mkdir -p ./bin
	cd main && GOGC=off go build -i -o ../bin/updater-server

dev:
	@(export CONFIG=$$PWD/etc/server.conf && cd ./main && fresh -c ../etc/fresh-runner.conf -w=../)

reset-db:
	bash scripts/db.sh
