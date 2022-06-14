swag fmt
swag init -g ./internal/api/ignite.go -o ./internal/api/docs
GIN_MODE=release
go run github.com/io-boxies/api