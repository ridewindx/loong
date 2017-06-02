protoc \
  -I=$GOPATH/src:$GOPATH/src/github.com/gogo/protobuf:$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis:$(dirname $1) \
  --gogoslick_out=plugins=grpc:$2 \
  --grpc-gateway_out=logtostderr=true:$2 \
  --swagger_out=logtostderr=true:$2 \
  $1
