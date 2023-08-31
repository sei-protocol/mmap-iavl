package store

import (
	"github.com/sei-protocol/mmap-iavl/memiavl"
	"github.com/sei-protocol/mmap-iavl/store/rootmulti"
	"path/filepath"

	"github.com/spf13/cast"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/baseapp"
	servertypes "github.com/cosmos/cosmos-sdk/server/types"
)

const (
	FlagMemIAVL            = "memiavl.enable"
	FlagAsyncCommitBuffer  = "memiavl.async-commit-buffer"
	FlagZeroCopy           = "memiavl.zero-copy"
	FlagSnapshotKeepRecent = "memiavl.snapshot-keep-recent"
	FlagSnapshotInterval   = "memiavl.snapshot-interval"
	FlagCacheSize          = "memiavl.cache-size"
)

// SetupMemIAVL insert the memiavl setter in front of baseapp options, so that
// the default rootmulti store is replaced by memiavl store,
func SetupMemIAVL(logger log.Logger, homePath string, appOpts servertypes.AppOptions, sdk46Compact bool, baseAppOptions []func(*baseapp.BaseApp)) []func(*baseapp.BaseApp) {
	if cast.ToBool(appOpts.Get(FlagMemIAVL)) {
		opts := memiavl.Options{
			AsyncCommitBuffer:  cast.ToInt(appOpts.Get(FlagAsyncCommitBuffer)),
			ZeroCopy:           cast.ToBool(appOpts.Get(FlagZeroCopy)),
			SnapshotKeepRecent: cast.ToUint32(appOpts.Get(FlagSnapshotKeepRecent)),
			SnapshotInterval:   cast.ToUint32(appOpts.Get(FlagSnapshotInterval)),
			CacheSize:          cast.ToInt(appOpts.Get(FlagCacheSize)),
		}

		// cms must be overridden before the other options, because they may use the cms,
		// make sure the cms aren't be overridden by the other options later on.
		baseAppOptions = append([]func(*baseapp.BaseApp){setMemIAVL(homePath, logger, opts, sdk46Compact)}, baseAppOptions...)
	}

	return baseAppOptions
}

func setMemIAVL(homePath string, logger log.Logger, opts memiavl.Options, sdk46Compact bool) func(*baseapp.BaseApp) {
	return func(bapp *baseapp.BaseApp) {
		// trigger state-sync snapshot creation by memiavl
		opts.TriggerStateSyncExport = func(height int64) {
			logger.Info("Triggering memIAVL snapshot creation")
			if !shouldTakeSnapshot(height, int64(opts.SnapshotInterval)) {
				logger.Debug("snapshot is skipped", "height", height)
			} else {
				go bapp.Snapshot(height)
			}

		}
		cms := rootmulti.NewStore(filepath.Join(homePath, "data", "memiavl.db"), logger, sdk46Compact)
		cms.SetMemIAVLOptions(opts)
		bapp.SetCMS(cms)
	}
}

// shouldTakeSnapshot returns true is snapshot should be taken at height.
func shouldTakeSnapshot(height int64, snapshotInterval int64) bool {
	return snapshotInterval > 0 && height%snapshotInterval == 0
}
