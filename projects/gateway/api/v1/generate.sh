#!/usr/bin/env bash

OUT=${GOPATH}/src/github.com/solo-io/solo-kit/projects/gateway/pkg/api/v1
GOGO_OUT_FLAG="--gogo_out=Mgoogle/protobuf/struct.proto=github.com/gogo/protobuf/types,Mgoogle/protobuf/duration.proto=github.com/gogo/protobuf/types:${GOPATH}/src/"
SOLO_KIT_FLAG="--plugin=protoc-gen-solo-kit=${GOPATH}/bin/protoc-gen-solo-kit --solo-kit_out=${OUT}"

mkdir -p ${OUT}
protoc -I=./ \
    -I=${GOPATH}/src/github.com/gogo/protobuf/ \
    -I=${GOPATH}/src \
    ${GOGO_OUT_FLAG} \
    ${SOLO_KIT_FLAG} \
    *.proto