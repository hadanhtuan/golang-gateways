# Run services
run:
	@echo "Choose run options:"
	@echo "[1] > Run Service"
	@echo "[2] > Generate Proto"
	@echo "[3] > Cancel"
	@read -p "Press [Enter]: " choice; \
	case "$$choice" in \
		1) make -f makefile run-service;; \
		2) make -f makefile gen-all;; \
		3) exit 1;; \
		*) echo "Not Found Option"; exit 1;; \
	esac


#Run service
run-service:
	nodemon --exec "go run" ./main.go

# GRPC
gen-sdk:
	@protoc --go_out=. --go_opt=paths=source_relative \
				--go-grpc_out=. --go-grpc_opt=paths=source_relative \
				proto/sdk/*.proto

gen-user:
	@protoc --go_out=. --go_opt=paths=source_relative \
				--go-grpc_out=. --go-grpc_opt=paths=source_relative \
				proto/user/*.proto

gen-property:
	@protoc --go_out=. --go_opt=paths=source_relative \
				--go-grpc_out=. --go-grpc_opt=paths=source_relative \
				proto/property/*.proto

gen-search:
	@protoc --go_out=. --go_opt=paths=source_relative \
				--go-grpc_out=. --go-grpc_opt=paths=source_relative \
				proto/search/*.proto

gen-payment:
	@protoc --go_out=. --go_opt=paths=source_relative \
				--go-grpc_out=. --go-grpc_opt=paths=source_relative \
				proto/payment/*.proto


gen-all: gen-sdk gen-user gen-property gen-search gen-payment
