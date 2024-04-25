
all: webapp

webapp:
	go build -o bin/fotos-wa cmd/webapp/main.go


.PHONY: clean

clean:
	rm -rf bin/
