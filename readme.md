generate go pb stubs
protoc -I/usr/local/include -I.  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis gameserver.proto --go_out=plugins=grpc:gameserver

generate grpc gateway
protoc -I/usr/local/include -I. \
  -I$GOPATH/src \
  -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
  --grpc-gateway_out=logtostderr=true:gameserver \
  gameserver.proto

generate swagger docs
protoc -I/usr/local/include -I. \
-I$GOPATH/src \
-I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
--swagger_out=logtostderr=true:gameserver \
gameserver.proto

grpc-gateway conf (custom Marshal)
https://grpc-ecosystem.github.io/grpc-gateway/docs/customizingyourgateway.html
