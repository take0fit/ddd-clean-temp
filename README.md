# ddd-clean-temp

## 概要
モジュラーモノリス ❌ ドメイン駆動設計（DDD） ❌ クリーンアーキテクチャ
ソフトウェアの複雑さを管理し、ビジネスロジックと技術的な詳細を明確に分離することにより、メンテナンス性とスケーラビリティを向上を目的とする。

## 特徴
モジュラーモノリス: 単一のコードベースでありながら、モジュール間の疎結合を保持。
ドメイン駆動設計: ビジネスドメインを中心にソフトウェアを設計します。
クリーンアーキテクチャ: ビジネスロジックと技術的詳細の分離を促進します。

## 技術スタック
言語: Go
フレームワーク: Gin (HTTPサーバーとして)
データベース: PostgreSQL
その他のツール: Docker, Wire (依存性注入), gin, gorm

## プロジェクト構造
```
.
├── README.md
├── api
│   └── openapi.yaml
├── cmd
│   └── rest_api
│       └── main.go
├── docker
│   ├── dev
│   │   └── Dockerfile
│   └── local
│       └── Dockerfile
├── docker-compose.yml
├── go.mod
├── go.sum
├── internal
│   ├── auth
│   │   ├── application
│   │   │   ├── dto
│   │   │   └── service
│   │   ├── domain
│   │   │   ├── entity
│   │   │   └── repository
│   │   ├── infrastructure
│   │   │   └── repository
│   │   └── interface
│   │       └── controller
│   ├── common
│   │   ├── db
│   │   │   └── mysql.go
│   │   └── http
│   │       └── handlers.go
│   └── user
│       ├── application
│       │   ├── dto
│       │   │   └── user_dto.go
│       │   └── service
│       │       ├── transaction.go
│       │       ├── user_service.go
│       │       └── user_service_impl.go
│       ├── domain
│       │   ├── entity
│       │   │   └── user.go
│       │   └── repository
│       │       └── user_repository.go
│       ├── infrastructure
│       │   └── user_repository_impl.go
│       ├── injection
│       │   ├── wire.go
│       │   └── wire_gen.go
│       └── interface
│           ├── controller
│           │   └── user_controller.go
│           └── presenter
│               └── response
│                   └── user_response.go
├── openapitools.json
├── pkg
│   ├── config
│   │   └── const
│   ├── middleware
│   └── util
└── scripts

```


コントリビューション
プロジェクトへの貢献に興味がある方は、プルリクエストを歓迎します。また、問題が発生した場合は、GitHubのIssuesに報告してください。
