generate-client:
	docker run --rm -v $(pwd):/local openapitools/openapi-generator-cli:v7.12.0 generate --git-user-id dhimasan0206 --git-repo-id adapter/sdk/client --package-name client -v -i /local/internal/oas/openapi.json -g go -o /local/sdk/client && cd sdk/client && go mod tidy && exit

generate-server:
	docker run --rm -v $(pwd):/local openapitools/openapi-generator-cli:v7.12.0 generate --git-user-id dhimasan0206 --git-repo-id adapter/sdk/server -s -v --package-name server --minimal-update -v -i /local/internal/oas/openapi.json -g go-gin-server -o /local/sdk/server --additional-properties=interfaceOnly=true && cd sdk/server && go mod tidy && exit