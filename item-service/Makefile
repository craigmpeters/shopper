build:
	protoc -I. --go_out=plugins=micro:$(GOPATH)/src/github.com/craigmpeters/shopper/item-service \
	  proto/item/item.proto
	protoc-go-inject-tag -input=$(GOPATH)/src/github.com/craigmpeters/shopper/item-service \
      proto/item/item.pb.go

run:
	docker-run -p 50051:50051 -e MICRO_SERVER_ADDRESS=:50051 -e MICRO_REGISTRY=mdns item-service