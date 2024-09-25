//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config ./cfg.client.yaml ./api/2.0.spec.yaml
//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config ./cfg.server.yaml ./api/2.0.spec.yaml

package oscp_v20
