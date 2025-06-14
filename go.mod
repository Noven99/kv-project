module kv-project

go 1.24

require (
	github.com/google/btree v1.1.3
	github.com/plar/go-adaptive-radix-tree v1.0.5 //这个不能用最新版本，只能用v1.0.5
	github.com/stretchr/testify v1.10.0
)

require go.etcd.io/bbolt v1.3.7 //这个不能用最新版本

require github.com/gofrs/flock v0.8.1 //这个不能用最新版本

require (
	github.com/tidwall/redcon v1.6.2
	golang.org/x/exp v0.0.0-20230321023759-10a507213a29
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/tidwall/btree v1.1.0 // indirect
	github.com/tidwall/match v1.1.1 // indirect
	golang.org/x/sys v0.29.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
