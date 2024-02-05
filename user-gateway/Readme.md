Project Require: 
1. install protoc https://www.geeksforgeeks.org/how-to-install-protocol-buffers-on-windows/
2. install grpc for golang https://grpc.io/docs/languages/go/quickstart/
3. install nodemon global: npm i nodemon -g

Generate Swagger:
1. [Docs](https://github.com/swaggo/swag)
2. Init after change
  swag init -g ./cmd/main.go -o ./docs
3. Format
  swag fmt
