module github.com/sei-protocol/mmap-iavl

go 1.20

require (
	cosmossdk.io/errors v1.0.0-beta.7
	github.com/confio/ics23/go v0.9.0
	github.com/cosmos/cosmos-sdk v0.45.10
	github.com/cosmos/gogoproto v1.4.7
	github.com/cosmos/iavl v0.19.6
	github.com/gogo/protobuf v1.3.3
	github.com/hashicorp/golang-lru v0.5.5-0.20210104140557-80c98217689d
	github.com/ledgerwatch/erigon-lib v0.0.0-20230210071639-db0e7ed11263
	github.com/spf13/cast v1.5.0
	github.com/stretchr/testify v1.8.4
	github.com/tendermint/tendermint v0.37.0-dev
	github.com/tendermint/tm-db v0.6.8-0.20220519162814-e24b96538a12
	github.com/tidwall/btree v1.5.0
	github.com/tidwall/gjson v1.10.2
	github.com/tidwall/wal v1.1.7
	github.com/zbiljic/go-filelock v0.0.0-20170914061330-1dbf7103ab7d
	golang.org/x/exp v0.0.0-20230310171629-522b1b587ee0
	golang.org/x/sync v0.1.0

)

require (

)

replace (
	github.com/cosmos/cosmos-sdk => ../sei-cosmos
	github.com/cosmos/iavl => ../sei-iavl
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/tendermint/tendermint => github.com/sei-protocol/sei-tendermint v0.2.24
	github.com/tendermint/tm-db => github.com/sei-protocol/tm-db v0.0.4
	google.golang.org/grpc => google.golang.org/grpc v1.33.2
)
