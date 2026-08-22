.PHONY: proto
proto:
	docker build -t onyx-protobuilder -f build/protobuilder/Dockerfile .
	docker run --rm -v $(PWD):/workspace onyx-protobuilder