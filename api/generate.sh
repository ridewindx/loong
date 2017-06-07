swagger generate server --spec api/authorization-v2.yaml --name LoongAuthorizationApi --target pkg/api/authorization
swagger generate server --spec api/content-v2.yaml --name LoongContentApi --target pkg/api/content
swagger generate server --spec api/token-v2.yaml --name LoongTokenApi --target pkg/api/token
swagger generate server --spec api/upload-v2.yaml --name LoongUploadApi --target pkg/api/upload

