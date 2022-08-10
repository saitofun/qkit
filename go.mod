module github.com/saitofun/qkit

go 1.18

require (
	github.com/DATA-DOG/go-sqlmock v1.5.0
	github.com/eclipse/paho.mqtt.golang v1.4.1
	github.com/fatih/color v1.13.0
	github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/gomodule/redigo v1.8.9
	github.com/google/uuid v1.3.0
	github.com/julienschmidt/httprouter v1.3.0
	github.com/lib/pq v1.10.6
	github.com/onsi/gomega v1.20.0
	github.com/pkg/errors v0.9.1
	github.com/saitofun/qlib v0.0.0-20220804014931-3a213f937710
	github.com/sirupsen/logrus v1.9.0
	github.com/spf13/cobra v1.5.0
	go.opentelemetry.io/contrib/propagators/b3 v0.0.0-00010101000000-000000000000
	go.opentelemetry.io/otel v1.9.0
	go.opentelemetry.io/otel/exporters/zipkin v1.7.0
	go.opentelemetry.io/otel/sdk v1.9.0
	go.opentelemetry.io/otel/trace v1.9.0
	golang.org/x/net v0.0.0-20220520000938-2e3eb7b945c2
	golang.org/x/tools v0.1.11
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/google/go-cmp v0.5.8 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/kr/pretty v0.3.0 // indirect
	github.com/mattn/go-colorable v0.1.12 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/openzipkin/zipkin-go v0.4.0 // indirect
	github.com/rogpeppe/go-internal v1.6.2 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220419223038-86c51ed26bb4 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
	golang.org/x/sys v0.0.0-20220727055044-e65921a090b8 // indirect
	golang.org/x/text v0.3.7 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	go.opentelemetry.io/contrib/propagators/b3 => go.opentelemetry.io/contrib/propagators/b3 v1.8.0
	go.opentelemetry.io/otel => go.opentelemetry.io/otel v1.9.0
	go.opentelemetry.io/otel/exporters/zipkin => go.opentelemetry.io/otel/exporters/zipkin v1.9.0
	go.opentelemetry.io/otel/sdk => go.opentelemetry.io/otel/sdk v1.9.0
	go.opentelemetry.io/otel/trace => go.opentelemetry.io/otel/trace v1.9.0
)
