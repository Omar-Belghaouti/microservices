check_install:
	which swagger || go get -u github.com/go-swagger/go-swagger/cmd/swagger

swagger: check_install
	swagger generate spec -o ./swagger.yml --scan-models

client: check_install
	swagger generate client -f swagger.yml -A product-api -t client/     

.PHONY: check_install swagger client