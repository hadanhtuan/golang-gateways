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
	nodemon --exec "go run" ./cmd/main.go --signal SIGTERM

# GRPC
gen-common:
	@protoc --go_out=. --go_opt=paths=source_relative \
				--go-grpc_out=. --go-grpc_opt=paths=source_relative \
				proto/common/*.proto

gen-user:
	@protoc --go_out=. --go_opt=paths=source_relative \
				--go-grpc_out=. --go-grpc_opt=paths=source_relative \
				proto/user/*.proto

gen-all: gen-common gen-user
