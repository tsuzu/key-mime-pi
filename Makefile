bin:
	mkdir -p bin

.PHONY: build-rpi-zero
build-rpi-zero: bin
	GOOS=linux GOARCH=arm GOARM=6 go build -o bin/key-mime-pi ./app/cmd/key-mime-pi
