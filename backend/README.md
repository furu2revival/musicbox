# backend

## Getting Started

### Prerequisites

- Golang
- Docker

### Installation

```shell
# Prepare tools
make install-tools
docker-compose up -d

# Prepare application config
export MUSICBOX_CONFIG_FILEPATH=$(pwd)/config/default.json
# Optional: Use custom config file
cp config/default.json config/{custom_config_name}.json  
export MUSICBOX_CONFIG_FILEPATH=$(pwd)/config/{custom_config_name}.json
```

## ポケットリファレンス

### コード生成する

[schema](./schema) の内容を元に、コードを自動生成します。

```shell
make codegen
```

### code format

```shell
make format
```

### lint

```shell
make lint
```

### DB マイグレーションする

1. [schema/db/ddl.sql`](./schema/db/ddl.sql) を編集する。
2. `make db-migrate` を実行する。

### DB をクリーンアップする

```shell
make db-clean
```

### curl で API を叩く

```shell
curl -X POST http://localhost:8000/api.debug.EchoService/EchoV1 -H "Content-Type: application/json" -H "Connect-Protocol-Version: 1" -d '{"message": "Hello, Connect-RPC!"}'
```
