module github.com/dexidp/dex

go 1.17

require (
	entgo.io/ent v0.10.0
	github.com/AppsFlyer/go-sundheit v0.5.0
	github.com/Masterminds/semver v1.5.0
	github.com/Masterminds/sprig/v3 v3.2.2
	github.com/beevik/etree v1.1.0
	github.com/coreos/go-oidc/v3 v3.1.0
	github.com/dexidp/dex/api/v2 v2.0.0
	github.com/felixge/httpsnoop v1.0.2
	github.com/ghodss/yaml v1.0.0
	github.com/go-ldap/ldap/v3 v3.4.1
	github.com/go-sql-driver/mysql v1.6.0
	github.com/gorilla/handlers v1.5.1
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/kylelemons/godebug v1.1.0
	github.com/lib/pq v1.10.4
	github.com/mattermost/xml-roundtrip-validator v0.1.0
	github.com/mattn/go-sqlite3 v1.14.10
	github.com/oklog/run v1.1.0
	github.com/pkg/errors v0.9.1
	github.com/prometheus/client_golang v1.11.0
	github.com/russellhaering/goxmldsig v1.1.1
	github.com/sirupsen/logrus v1.8.1
	github.com/spf13/cobra v1.3.0
	github.com/stretchr/testify v1.7.1-0.20210427113832-6241f9ab9942
	go.etcd.io/etcd/client/pkg/v3 v3.5.1
	go.etcd.io/etcd/client/v3 v3.5.1
	golang.org/x/crypto v0.0.0-20210817164053-32db794688a5
	golang.org/x/net v0.0.0-20210813160813-60bc85c4be6d
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8
	google.golang.org/api v0.63.0
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/square/go-jose.v2 v2.6.0
)

require (
	ariga.io/atlas v0.3.2-0.20220120225051-c3fac7d636dd // indirect
	cloud.google.com/go v0.99.0 // indirect
	github.com/Azure/go-ntlmssp v0.0.0-20200615164410-66371956d46c // indirect
	github.com/Masterminds/goutils v1.1.1 // indirect
	github.com/Masterminds/semver/v3 v3.1.1 // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/go-systemd/v22 v22.3.2 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/go-asn1-ber/asn1-ber v1.5.1 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/googleapis/gax-go/v2 v2.1.1 // indirect
	github.com/hashicorp/hcl/v2 v2.10.0 // indirect
	github.com/huandu/xstrings v1.3.1 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/matttproud/golang_protobuf_extensions v1.0.1 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/mitchellh/reflectwalk v1.0.0 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0 // indirect
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/shopspring/decimal v1.2.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	go.etcd.io/etcd/api/v3 v3.5.1 // indirect
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.17.0 // indirect
	golang.org/x/mod v0.5.1 // indirect
	golang.org/x/sys v0.0.0-20211210111614-af8b64212486 // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/genproto v0.0.0-20211208223120-3a66f561d7aa // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
)

replace github.com/dexidp/dex/api/v2 => ./api/v2
