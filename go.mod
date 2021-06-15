module github.com/ozoncp/ocp-issue-api

go 1.16

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0 // indirect
	github.com/Masterminds/squirrel v1.5.0
	github.com/cockroachdb/apd v1.1.0 // indirect
	github.com/envoyproxy/protoc-gen-validate v0.1.0
	github.com/gofrs/uuid v4.0.0+incompatible // indirect
	github.com/golang/mock v1.5.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/jackc/fake v0.0.0-20150926172116-812a484cc733 // indirect
	github.com/jackc/pgx v3.6.2+incompatible
	github.com/lib/pq v1.10.2 // indirect
	github.com/onsi/ginkgo v1.16.2
	github.com/onsi/gomega v1.13.0
	github.com/pressly/goose v2.7.0+incompatible // indirect
	github.com/rs/zerolog v1.22.0
	github.com/shopspring/decimal v1.2.0 // indirect
	golang.org/x/crypto v0.0.0-20210513164829-c07d793c2f9a // indirect
	golang.org/x/net v0.0.0-20210614182718-04defd469f4e // indirect
	golang.org/x/sys v0.0.0-20210611083646-a4fc73990273 // indirect
	google.golang.org/genproto v0.0.0-20210614182748-5b3b54cad159
	google.golang.org/grpc v1.38.0
	google.golang.org/protobuf v1.26.0
)

replace github.com/ozoncp/ocp-issue-api/pkg/ocp-issue-api => ./pkg/ocp-issue-api
