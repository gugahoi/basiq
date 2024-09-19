//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -package=basiq -generate=types,client -o=../internal/api/webhooks.gen.go ../.api/webhooks.json

package tools
