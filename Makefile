S3_CFN_BUCKET?=gourav1234567890
ENVIRONMENT?=dev
SERVICE_NAME?=go-lambda-bootstrap-project

update-vendor:
	go mod tidy
	go mod vendor

clean:
	rm -rf tmp

build: clean update-vendor
	mkdir -p tmp; \
	cp $(SERVICE_NAME).yaml tmp/; \
    export GOOS="linux"; \
    export GOARCH="amd64"; \
    go build -o tmp .

package: build
	cd ./tmp; \
	aws cloudformation package \
	--template-file $(SERVICE_NAME).yaml \
	--s3-bucket $(S3_CFN_BUCKET) \
	--output-template-file $(SERVICE_NAME)-output.yaml

deploy: package
	cd ./tmp; \
	aws cloudformation deploy \
	--template-file $(SERVICE_NAME)-output.yaml \
	--capabilities CAPABILITY_NAMED_IAM \
	--stack-name $(SERVICE_NAME)-$(ENVIRONMENT) \
	--parameter-overrides LogLevel=info ServiceName=$(SERVICE_NAME) Environment=$(ENVIRONMENT)
