FROM golang:1.27-bookworm

ARG PROTOC_VERSION=25.3
RUN apt-get update && apt-get install -y --no-install-recommends unzip curl && \
    curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VERSION}/protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    unzip protoc-${PROTOC_VERSION}-linux-x86_64.zip -d /usr/local && \
    rm protoc-${PROTOC_VERSION}-linux-x86_64.zip && \
    rm -rf /var/lib/apt/lists/*

RUN go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest && \
    go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv3@latest && \
    go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

ARG GOOGLEAPIS_COMMIT=58bc461d695b3254a0ff88cc3541b6dca7e7a95f
RUN git clone --filter=blob:none --no-checkout --depth 1 \
      https://github.com/googleapis/googleapis.git /opt/googleapis && \
    cd /opt/googleapis && \
    git sparse-checkout init --no-cone && \
    git sparse-checkout set \
        "google/api/annotations.proto" \
        "google/api/http.proto" && \
    git checkout ${GOOGLEAPIS_COMMIT}

ENV PATH="/root/go/bin:${PATH}"
WORKDIR /workspace

COPY scripts/protobuilder.sh /usr/local/bin/protobuilder.sh
RUN chmod +x /usr/local/bin/protobuilder.sh

ENTRYPOINT ["/usr/local/bin/protobuilder.sh"]