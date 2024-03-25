# hlf-control-plane


### !!Disclaimer!!

Please note the repo is in syncing mode now Due to some integration and testing issues the source is in closed system. We are planning to come into Github only in May 2024


## TOC

- [hlf-control-plane](#hlf-control-plane)
    - [!!Disclaimer!!](#disclaimer)
  - [TOC](#toc)
  - [Description](#description)
  - [Open API](#open-api)
  - [Configuration](#configuration)
  - [Monitoring](#monitoring)
  - [Development](#development)
    - [Utils](#utils)
    - [Protobuf](#protobuf)
    - [Code generation](#code-generation)
  - [License](#license)
  - [Links](#links)

## Description

Tool service for administraive function of fabric net: channel creation, chainconde install, etc
#admin#chaincode#deploy#hlf#controlplane#

## Open API

The service has an API.  Specs see in Swagger(proto/plane.swagger.json),     

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
  - host: peer1.org1.org:7051
  - host: peer2.org1.org:7051

# ports configuration fo grpc and http
listen:
  http: ":8080"
  grpc: ":8081"
```

## Monitoring

Service has a  **healthcheck** endpoint without authorisation: `/v1/healthz` with answer:
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