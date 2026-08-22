FROM golang:1.26.6-bookworm

ARG PROTOC_VERSION=25.3
RUN apt-get update && apt-get install -y --no-install-recommends unzip curl && \
    curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local && \
    rm protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    rm -rf /var/lib/apt/lists/*

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.34.2 && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.5.1

ENV PATH="/root/go/bin:${PATH}"
WORKDIR /workspace

COPY scripts/protobuilder.sh /usr/local/bin/protobuilder.sh
RUN chmod +x /usr/local/bin/protobuilder.sh

ENTRYPOINT ["/usr/local/bin/protobuilder.sh"]