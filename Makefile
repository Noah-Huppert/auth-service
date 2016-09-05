.PHONY: build run all
all: build run
build:
	docker build -t noah-huppert/auth-service:dev-latest .
run:
	docker run -it --rm --name auth-service -p 3000:80 noah-huppert/auth-service:dev-latest
