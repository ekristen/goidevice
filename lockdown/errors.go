package lockdown

// #cgo pkg-config: libimobiledevice-1.0
// #include <libimobiledevice/lockdown.h>
import "C"

import (
	"github.com/pauldotknopf/goidevice/common"
)

func resultToError(result C.idevice_error_t) error {
	switch result {
	case 0:
		return nil
	case -1:
		return common.ErrInvalidArgs
	case -2:
		return common.ErrUnknown
	case -3:
		return common.ErrNoDevice
	case -4:
		return common.ErrNotEnoughData
	case -5:
		return common.ErrBadHeader
	case -6:
		return common.ErrSslError
	case -7:
		return common.ErrReceiveTimeout
	case -8:
		return common.ErrMuxError
	case -9:
		return common.ErrNoRunningSession
	case -10:
		return common.ErrInvalidResponse
	case -11:
		return common.ErrMissingKey
	case -12:
		return common.ErrMissingValue
	case -13:
		return common.ErrGetProhibited
	case -14:
		return common.ErrSetProhibited
	case -15:
		return common.ErrRemoveProhibited
	case -16:
		return common.ErrImmutableValue
	case -17:
		return common.ErrPasswordProtected
	case -18:
		return common.ErrUserDeniedPairing
	case -19:
		return common.ErrPairingDialogResponsePending
	case -20:
		return common.ErrMissingHostID
	case -21:
		return common.ErrInvalidHostID
	case -22:
		return common.ErrSessionActive
	case -23:
		return common.ErrSessionInactive
	case -24:
		return common.ErrMissingSessionID
	case -25:
		return common.ErrInvalidSessionID
	case -26:
		return common.ErrMissingService
	case -27:
		return common.ErrInvalidService
	case -28:
		return common.ErrServiceLimit
	case -29:
		return common.ErrMissingPairRecord
	case -30:
		return common.ErrSavePairRecordFailed
	case -31:
		return common.ErrInvalidPairRecord
	case -32:
		return common.ErrInvalidActivationPeriod
	case -33:
		return common.ErrMissingActivationPeriod
	case -34:
		return common.ErrServiceProhibited
	case -35:
		return common.ErrEscrowLocked
	case -36:
		return common.ErrPairingProhibitedOverThisConnection
	case -37:
		return common.ErrFMIPProtected
	case -38:
		return common.ErrMCProtected
	case -39:
		return common.ErrMCChallengeRequired
	default:
		return common.ErrUnknown
	}
}
