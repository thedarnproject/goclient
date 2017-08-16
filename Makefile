BIN_DIR := bin
CLIENT_DIR := client

default: binary

binary:
	mkdir -p bin
	go build -o ${BIN_DIR}/goclient ${CLIENT_DIR}/main.go

clean:
	rm -rf bin/

glide:
	glide update --strip-vendor

glide-vc:
	glide-vc --only-code --no-tests

glide-hard:
	rm -rf ~/.glide/
	glide update --strip-vendor

