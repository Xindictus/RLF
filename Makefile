.DEFAULT_GOAL := swagger
command = "export PATH=$$PATH:$$(go env GOPATH)/bin"

install_swagger:
	go get -u github.com/swaggo/swag/cmd/swag
	go install github.com/swaggo/swag/cmd/swag@latest

swagger:
	@echo Ensure you have the swagger CLI or this command will fail.
	@echo You can install the swagger CLI with: go get -u github.com/swaggo/swag/cmd/swag
	@echo 'If swagger cmd does not work, verify that you have $(command) in your .bashrc file'
	@echo ....

	swag init