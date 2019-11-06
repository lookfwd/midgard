module gitlab.com/thorchain/bepswap/chain-service

go 1.13

require (
	github.com/99designs/gqlgen v0.10.1
	github.com/DATA-DOG/godog v0.7.13 // indirect
	github.com/binance-chain/go-sdk v1.1.3
	github.com/bradfitz/gomemcache v0.0.0-20190913173617-a41fca850d0b // indirect
	github.com/btcsuite/btcd v0.0.0-20190926002857-ba530c4abb35 // indirect
	github.com/btcsuite/btcutil v0.0.0-20190425235716-9e5f4b9a998d
	github.com/cenkalti/backoff v2.2.1+incompatible
	github.com/cosmos/cosmos-sdk v0.37.3
	github.com/cpuguy83/go-md2man/v2 v2.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1
	github.com/deepmap/oapi-codegen v1.3.0
	github.com/getkin/kin-openapi v0.2.0
	github.com/gin-contrib/cache v1.1.0 // indirect
	github.com/gin-contrib/logger v0.0.1 // indirect
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.4.0
	github.com/google/go-cmp v0.3.1 // indirect
	github.com/gorilla/mux v1.7.0
	github.com/hashicorp/golang-lru v0.5.3 // indirect
	github.com/influxdata/influxdb1-client v0.0.0-20190809212627-fc22c7df067e
	github.com/kazukousen/gouml v0.0.0-20190718105346-f20d094b56c5 // indirect
	github.com/labstack/echo/v4 v4.1.11
	github.com/mattn/go-colorable v0.1.4 // indirect
	github.com/onsi/ginkgo v1.10.2 // indirect
	github.com/onsi/gomega v1.7.0 // indirect
	github.com/openlyinc/pointy v1.1.2
	github.com/pelletier/go-toml v1.5.0 // indirect
	github.com/pkg/errors v0.8.1
	github.com/prometheus/client_golang v1.2.1 // indirect
	github.com/rs/zerolog v1.15.0
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.4.0
	github.com/superoo7/go-gecko v0.0.0-20190607060444-a448b0c99969 // indirect
	github.com/tendermint/go-amino v0.15.1 // indirect
	github.com/ugorji/go v1.1.7 // indirect
	github.com/urfave/cli v1.22.1 // indirect
	github.com/valyala/fasttemplate v1.1.0 // indirect
	github.com/vektah/gqlparser v1.1.2
	github.com/yogendra/plantuml-go v0.0.0-20170731163123-1e0758c537a3 // indirect
	github.com/ziflex/lecho/v2 v2.0.0
	golang.org/x/crypto v0.0.0-20191002192127-34f69633bfdc // indirect
	golang.org/x/net v0.0.0-20191007182048-72f939374954 // indirect
	golang.org/x/xerrors v0.0.0-20191011141410-1b5146add898 // indirect
	google.golang.org/genproto v0.0.0-20191007204434-a023cd5227bd
	google.golang.org/grpc v1.24.0 // indirect
	gopkg.in/check.v1 v1.0.0-20190902080502-41f04d3bba15
	gopkg.in/h2non/gock.v1 v1.0.15 // indirect
)

replace github.com/tendermint/go-amino => github.com/binance-chain/bnc-go-amino v0.14.1-binance.1

replace github.com/ugorji/go v1.1.4 => github.com/ugorji/go/codec v0.0.0-20190204201341-e444a5086c43
