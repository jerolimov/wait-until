build:
	go build -o wu .

install: build
	cp -v wu ~/bin/
