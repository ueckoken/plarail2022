package internal

import (
	"github.com/ueckoken/plarail2022/backend/external/pkg/synccontroller"
	"github.com/ueckoken/plarail2022/backend/external/spec"
	"go.uber.org/zap"
)

// startBlockSync starts sync controller for block state.
func startBlockSync(logger *zap.Logger, syncInput chan synccontroller.KV[spec.BlockId, spec.BlockStateEnum], syncOutput chan<- synccontroller.KV[spec.BlockId, spec.BlockStateEnum]) {
	s := synccontroller.NewSyncController(logger, syncInput, syncOutput)

	go s.Run()
}
