# Gateway
run:
	@echo "Choose run options:"
	@echo "[1] > User Gateway"
	@echo "[2] > Admin Gateway"
	@echo "[3] > Cancel"
	@read -p "Press [Enter]: " choice; \
    case "$$choice" in \
        1) make user-gateway;; \
        2) make admin-gateway;; \
        3) exit 1;; \
        *) echo "Not Found Option"; exit 1;; \
    esac

user-gateway:
	nodemon --exec "go run" ./cmd/user-gateway/main.go --signal SIGTERM

admin-gateway:
	nodemon --exec "go run" ./cmd/admin-gateway/main.go --signal SIGTERM

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

