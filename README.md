# hlf-control-plane

## TOC

- [hlf-control-plane](#hlf-control-plane)
    - [TOC](#toc)
    - [Description](#description)
    - [OpenAPI](#open-api)
    - [Configuration](#configuration)
    - [Monitoring](#monitoring)
    - [Development](#development)
        - [Utils](#utils)
        - [Protobuf management](#protobuf)
        - [Code generation](#code-generation)
    - [License](#license)
    - [Links](#links)

## Description

Утилита- сервис предоставляющая упрощенный доступ к администраторским функциям fabric, как например создание каналов,
инсталляция чейнкодов, апдейт каналов #admin#chaincode#deploy#hlf#controlplane#

## Open API

Сервис предоставляет функционал OpenAPI(Swagger). Список методов и форматов запросов доступен
по [ссылке](proto/plane.swagger.json)

## Configuration

```yaml
# Msp ID of your organization
mspId: atomyzeMSP
# Logging level
logLevel: debug
# Access token value
accessToken: "my_awesome_token"
# Path to cert and private key of your identity
identity:
  cert: certs/atomyze_admin.pem
  key: certs/atomyze_admin_key.pem
# or bccsp config of pkcs11
#  bccsp:
#    Default: PKCS11
#    PKCS11:
#      immutable: false
#      label: Org_Admin
#      library: /usr/lib/softhsm/libsofthsm2.so
#      pin: 123321
#      hash: SHA2
#      security: 256

# Tls credentials for mutual tls connection
tls:
  cert: cert.pem
  key: key.pem
  ca: ca.pem

# List of peers managed by organization
peers:
  - host: peer1.atomyze.stage.core.n-t.io:7051
  - host: peer2.atomyze.stage.core.n-t.io:7051

# ports configuration fo grpc and http
listen:
  http: ":8080"
  grpc: ":8081"
```

## Monitoring

Сервис содержит встроенный механизм **healthcheck**, доступный без авторизации по пути `/v1/healthz`, возвращающий информацию следующего вида:
```json
{
  "status": "SERVING"
}
```

## Development

### Utils

For generation tools needed to work, use

```bash
make build-tools
```

Tools and their dependencies are described in module `tools`. Installation path is `$(pwd)/bin` by default.

### Protobuf

1. protodep - protobuf dependency management
2. buf - generate and lint protobuf code

To download vendored protobuf dependencies:

```bash
bin/protodep up
```

### Code generation

Use `go generate` in root path to generate code from `.proto` and `swagger` client for tests:

```bash
go generate
```

## License

Apache-2.0

## Links

- [protodep](https://github.com/stormcat24/protodep)
- [buf](https://buf.build)