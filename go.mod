module github.com/beiduoke/go-scaffold

go 1.23.0

toolchain go1.23.1

require (
	buf.build/gen/go/bufbuild/protovalidate/protocolbuffers/go v1.36.5-20250130201111-63bb56e20495.1
	entgo.io/ent v0.14.1
	github.com/BurntSushi/toml v1.4.0
	github.com/bwmarrin/snowflake v0.3.0
	github.com/casbin/casbin/v2 v2.103.0
	github.com/casbin/ent-adapter v0.4.0
	github.com/casbin/gorm-adapter/v3 v3.32.0
	github.com/casbin/redis-watcher/v2 v2.5.0
	github.com/envoyproxy/protoc-gen-validate v1.2.1
	github.com/go-chassis/sc-client v0.7.0
	github.com/go-kratos/aegis v0.2.0
	github.com/go-kratos/grpc-gateway/v2 v2.5.1-0.20210811062259-c92d36e434b1
	github.com/go-kratos/kratos/contrib/config/apollo/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/config/consul/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/config/etcd/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/config/nacos/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/log/aliyun/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/log/fluent/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/log/logrus/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/log/tencent/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/log/zap/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/registry/consul/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/registry/eureka/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/registry/kubernetes/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/registry/nacos/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/registry/servicecomb/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/contrib/registry/zookeeper/v2 v2.0.0-20250127095200-20970020a5ef
	github.com/go-kratos/kratos/v2 v2.8.3
	github.com/go-kratos/swagger-api v1.0.1
	github.com/go-zookeeper/zk v1.0.4
	github.com/golang-jwt/jwt/v4 v4.5.1
	github.com/golang-jwt/jwt/v5 v5.2.1
	github.com/google/gnostic v0.7.0
	github.com/google/subcommands v1.2.0
	github.com/google/wire v0.6.0
	github.com/gorilla/handlers v1.5.2
	github.com/hashicorp/consul/api v1.31.0
	github.com/iancoleman/strcase v0.3.0
	github.com/meilisearch/meilisearch-go v0.30.0
	github.com/minio/minio-go/v7 v7.0.85
	github.com/nacos-group/nacos-sdk-go v1.1.5
	github.com/nicksnyder/go-i18n/v2 v2.5.1
	github.com/olekukonko/tablewriter v0.0.5
	github.com/pkg/errors v0.9.1
	github.com/redis/go-redis/extra/redisotel/v9 v9.7.0
	github.com/redis/go-redis/v9 v9.7.0
	github.com/sirupsen/logrus v1.9.3
	github.com/spf13/cobra v1.8.1
	github.com/stretchr/testify v1.10.0
	github.com/tx7do/go-utils/entgo v1.1.24
	github.com/tx7do/kratos-swagger-ui v0.0.0-20241213153527-eca591f9f8b2
	go.etcd.io/etcd/client/v3 v3.5.18
	go.opentelemetry.io/otel v1.34.0
	go.opentelemetry.io/otel/exporters/jaeger v1.17.0
	go.opentelemetry.io/otel/exporters/zipkin v1.34.0
	go.opentelemetry.io/otel/sdk v1.34.0
	go.uber.org/zap v1.27.0
	golang.org/x/crypto v0.33.0
	golang.org/x/text v0.22.0
	golang.org/x/tools v0.30.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240318140521-94a12d6c2237
	google.golang.org/grpc v1.64.1
	google.golang.org/protobuf v1.36.5
	gopkg.in/natefinch/lumberjack.v2 v2.2.1
	gorm.io/driver/mysql v1.5.7
	gorm.io/gorm v1.25.12
	k8s.io/client-go v0.32.1
)

require (
	ariga.io/atlas v0.29.0 // indirect
	dario.cat/mergo v1.0.0 // indirect
	entgo.io/contrib v0.6.0 // indirect
	github.com/agext/levenshtein v1.2.3 // indirect
	github.com/aliyun/alibaba-cloud-sdk-go v1.61.18 // indirect
	github.com/aliyun/aliyun-log-go-sdk v0.1.75 // indirect
	github.com/andybalholm/brotli v1.1.1 // indirect
	github.com/apolloconfig/agollo/v4 v4.3.1 // indirect
	github.com/apparentlymart/go-textseg/v15 v15.0.0 // indirect
	github.com/armon/go-metrics v0.4.1 // indirect
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/bmatcuk/doublestar/v4 v4.6.1 // indirect
	github.com/bufbuild/protocompile v0.14.1 // indirect
	github.com/buger/jsonparser v1.1.1 // indirect
	github.com/casbin/govaluate v1.3.0 // indirect
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cenkalti/backoff/v4 v4.1.1 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.2-0.20180830191138-d8f796af33cc // indirect
	github.com/deckarep/golang-set v1.7.1 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/dustin/go-humanize v1.0.1 // indirect
	github.com/emicklei/go-restful/v3 v3.11.0 // indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/felixge/httpsnoop v1.0.3 // indirect
	github.com/fluent/fluent-logger-golang v1.9.0 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/fxamacker/cbor/v2 v2.7.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/glebarez/go-sqlite v1.20.3 // indirect
	github.com/glebarez/sqlite v1.7.0 // indirect
	github.com/go-chassis/cari v0.6.0 // indirect
	github.com/go-chassis/foundation v0.4.0 // indirect
	github.com/go-chassis/openlog v1.1.3 // indirect
	github.com/go-errors/errors v1.0.1 // indirect
	github.com/go-ini/ini v1.67.0 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/inflect v0.21.0 // indirect
	github.com/go-openapi/jsonpointer v0.21.0 // indirect
	github.com/go-openapi/jsonreference v0.20.2 // indirect
	github.com/go-openapi/swag v0.23.0 // indirect
	github.com/go-playground/form/v4 v4.2.0 // indirect
	github.com/go-sql-driver/mysql v1.7.0 // indirect
	github.com/goccy/go-json v0.10.4 // indirect
	github.com/gofrs/uuid v4.2.0+incompatible // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-sql/civil v0.0.0-20220223132316-b832511892a9 // indirect
	github.com/golang-sql/sqlexp v0.1.0 // indirect
	github.com/golang/glog v1.2.3 // indirect
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/gnostic-models v0.6.9-0.20230804172637-c7be7c783f49 // indirect
	github.com/google/go-cmp v0.6.0 // indirect
	github.com/google/gofuzz v1.2.0 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/mux v1.8.1 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.11.3 // indirect
	github.com/hashicorp/errwrap v1.1.0 // indirect
	github.com/hashicorp/go-cleanhttp v0.5.2 // indirect
	github.com/hashicorp/go-hclog v1.5.0 // indirect
	github.com/hashicorp/go-immutable-radix v1.3.1 // indirect
	github.com/hashicorp/go-multierror v1.1.1 // indirect
	github.com/hashicorp/go-rootcerts v1.0.2 // indirect
	github.com/hashicorp/golang-lru v0.5.4 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/hashicorp/hcl/v2 v2.23.0 // indirect
	github.com/hashicorp/serf v0.10.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/jackc/pgpassfile v1.0.0 // indirect
	github.com/jackc/pgservicefile v0.0.0-20221227161230-091c0ba34f0a // indirect
	github.com/jackc/pgx/v5 v5.6.0 // indirect
	github.com/jackc/puddle/v2 v2.2.1 // indirect
	github.com/jhump/protoreflect v1.17.0 // indirect
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.5 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/karlseguin/ccache/v2 v2.0.8 // indirect
	github.com/klauspost/compress v1.17.11 // indirect
	github.com/klauspost/cpuid/v2 v2.2.9 // indirect
	github.com/lib/pq v1.10.2 // indirect
	github.com/lufia/plan9stats v0.0.0-20230326075908-cb1d2100619a // indirect
	github.com/magiconair/properties v1.8.6 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	github.com/mattn/go-runewidth v0.0.9 // indirect
	github.com/microsoft/go-mssqldb v1.6.0 // indirect
	github.com/minio/md5-simd v1.1.2 // indirect
	github.com/mitchellh/go-homedir v1.1.0 // indirect
	github.com/mitchellh/go-wordwrap v1.0.1 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/munnerz/goautoneg v0.0.0-20191010083416-a7dc8b61c822 // indirect
	github.com/openzipkin/zipkin-go v0.4.3 // indirect
	github.com/patrickmn/go-cache v2.1.0+incompatible // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pelletier/go-toml/v2 v2.0.0-beta.8 // indirect
	github.com/philhofer/fwd v1.1.1 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pmezard/go-difflib v1.0.1-0.20181226105442-5d4384ee4fb2 // indirect
	github.com/power-devops/perfstat v0.0.0-20221212215047-62379fc7944b // indirect
	github.com/rakyll/statik v0.1.7 // indirect
	github.com/redis/go-redis/extra/rediscmd/v9 v9.7.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20230126093431-47fa9a501578 // indirect
	github.com/rs/xid v1.6.0 // indirect
	github.com/shirou/gopsutil/v3 v3.23.6 // indirect
	github.com/shoenig/go-m1cpu v0.1.6 // indirect
	github.com/sony/sonyflake v1.2.0 // indirect
	github.com/spf13/afero v1.10.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.11.0 // indirect
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/swaggest/swgui v1.8.2 // indirect
	github.com/tencentcloud/tencentcloud-cls-sdk-go v1.0.2 // indirect
	github.com/tinylib/msgp v1.1.6 // indirect
	github.com/tklauser/go-sysconf v0.3.11 // indirect
	github.com/tklauser/numcpus v0.6.1 // indirect
	github.com/tx7do/go-utils v1.1.13 // indirect
	github.com/vearutop/statigz v1.4.3 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	github.com/yusufpapurcu/wmi v1.2.3 // indirect
	github.com/zclconf/go-cty v1.15.1 // indirect
	go.etcd.io/etcd/api/v3 v3.5.18 // indirect
	go.etcd.io/etcd/client/pkg/v3 v3.5.18 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel/metric v1.34.0 // indirect
	go.opentelemetry.io/otel/trace v1.34.0 // indirect
	go.uber.org/atomic v1.10.0 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/exp v0.0.0-20241108190413-2d47ceb2692f // indirect
	golang.org/x/mod v0.23.0 // indirect
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/oauth2 v0.24.0 // indirect
	golang.org/x/sync v0.11.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/term v0.29.0 // indirect
	golang.org/x/time v0.7.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250124145028-65684f501c47 // indirect
	gopkg.in/evanphx/json-patch.v4 v4.12.0 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	gopkg.in/ini.v1 v1.66.4 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
	gorm.io/driver/postgres v1.5.9 // indirect
	gorm.io/driver/sqlserver v1.5.3 // indirect
	gorm.io/plugin/dbresolver v1.5.3 // indirect
	k8s.io/api v0.32.1 // indirect
	k8s.io/apimachinery v0.32.1 // indirect
	k8s.io/klog/v2 v2.130.1 // indirect
	k8s.io/kube-openapi v0.0.0-20241105132330-32ad38e42d3f // indirect
	k8s.io/utils v0.0.0-20241104100929-3ea5e8cea738 // indirect
	modernc.org/libc v1.22.2 // indirect
	modernc.org/mathutil v1.5.0 // indirect
	modernc.org/memory v1.5.0 // indirect
	modernc.org/sqlite v1.20.3 // indirect
	sigs.k8s.io/json v0.0.0-20241010143419-9aa6b5e7a4b3 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.4.2 // indirect
	sigs.k8s.io/yaml v1.4.0 // indirect
)
