package config

import (
	"io/ioutil"
	"time"

	"github.com/MixinNetwork/mixin/crypto"
	"github.com/pelletier/go-toml"
)

const (
	Debug        = true
	BuildVersion = "v0.8.3-BUILD_VERSION"

	MainnetId                  = "6430225c42bb015b4da03102fa962e4f4ef3969e03e04345db229f8377ef7997"
	SnapshotRoundGap           = uint64(3 * time.Second)
	SnapshotReferenceThreshold = 10
	SnapshotSyncRoundThreshold = 100
	SnapshotRoundSize          = 200
	TransactionMaximumSize     = 1024 * 1024
	WithdrawalClaimFee         = "0.0001"
	GossipSize                 = 3

	KernelMintTimeBegin = 7
	KernelMintTimeEnd   = 9

	KernelNodeAcceptTimeBegin     = 13
	KernelNodeAcceptTimeEnd       = 19
	KernelNodePledgePeriodMinimum = 12 * time.Hour
	KernelNodeAcceptPeriodMinimum = 12 * time.Hour
	KernelNodeAcceptPeriodMaximum = 7 * 24 * time.Hour
)

type Custom struct {
	Node struct {
		Signer               crypto.Key `toml:"-"`
		SignerStr            string     `toml:"signer-key"`
		ConsensusOnly        bool       `toml:"consensus-only"`
		KernelOprationPeriod int        `toml:"kernel-operation-period"`
		MemoryCacheSize      int        `toml:"memory-cache-size"`
		CacheTTL             int        `toml:"cache-ttl"`
		RingCacheSize        uint64     `toml:"ring-cache-size"`
		RingFinalSize        uint64     `toml:"ring-final-size"`
	} `toml:"node"`
	Storage struct {
		ValueLogGC bool `toml:"value-log-gc"`
	} `toml:"storage"`
	Network struct {
		Listener        string `toml:"listener"`
		GossipNeighbors bool   `toml:"gossip-neighbors"`
	} `toml:"network"`
	RPC struct {
		Runtime bool `toml:"runtime"`
	} `toml:"rpc"`
}

func Initialize(file string) (*Custom, error) {
	f, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	var config Custom
	err = toml.Unmarshal(f, &config)
	if err != nil {
		return nil, err
	}
	key, err := crypto.KeyFromString(config.Node.SignerStr)
	if err != nil {
		return nil, err
	}
	config.Node.Signer = key
	if config.Node.KernelOprationPeriod == 0 {
		config.Node.KernelOprationPeriod = 700
	}
	if config.Node.MemoryCacheSize == 0 {
		config.Node.MemoryCacheSize = 1024 * 16
	}
	if config.Node.CacheTTL == 0 {
		config.Node.CacheTTL = 3600 * 2
	}
	if config.Node.RingCacheSize == 0 {
		config.Node.RingCacheSize = 1024 * 1024
	}
	if config.Node.RingFinalSize == 0 {
		config.Node.RingFinalSize = 1024 * 1024 * 16
	}
	return &config, nil
}
