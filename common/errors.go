package common

import (
	"errors"
)

var (
	// ErrInvalidArgs .
	ErrInvalidArgs = errors.New("invalid args")
	// ErrUnknown .
	ErrUnknown = errors.New("unknown")
	// ErrNoDevice .
	ErrNoDevice = errors.New("no device")
	// ErrNotEnoughData .
	ErrNotEnoughData = errors.New("not enough data")
	// ErrBadHeader .
	ErrBadHeader = errors.New("bad header")
	// ErrSslError .
	ErrSslError = errors.New("ssl error")
	// ErrReceiveTimeout .
	ErrReceiveTimeout = errors.New("receive timeout")
	// ErrMuxError .
	ErrMuxError = errors.New("mux error")
	// ErrNoRunningSession .
	ErrNoRunningSession = errors.New("no running session")
	// ErrInvalidResponse .
	ErrInvalidResponse = errors.New("invalid response")
	// ErrMissingKey .
	ErrMissingKey = errors.New("missing key")
	// ErrMissingValue .
	ErrMissingValue = errors.New("missing value")
	// ErrGetProhibited .
	ErrGetProhibited = errors.New("get prohibited")
	// ErrSetProhibited .
	ErrSetProhibited = errors.New("set prohibited")
	// ErrRemoveProhibited .
	ErrRemoveProhibited = errors.New("remove prohibited")
	// ErrImmutableValue .
	ErrImmutableValue = errors.New("immutable value")
	// ErrPasswordProtected .
	ErrPasswordProtected = errors.New("password protected")
	// ErrUserDeniedPairing .
	ErrUserDeniedPairing = errors.New("user denied pairing")
	// ErrPairingDialogResponsePending .
	ErrPairingDialogResponsePending = errors.New("pairing dialog response pending")
	// ErrMissingHostID .
	ErrMissingHostID = errors.New("missing host id")
	// ErrInvalidHostID .
	ErrInvalidHostID = errors.New("invalid host id")
	// ErrSessionActive .
	ErrSessionActive = errors.New("session active")
	// ErrSessionInactive .
	ErrSessionInactive = errors.New("session inactive")
	// ErrMissingSessionID .
	ErrMissingSessionID = errors.New("missing session id")
	// ErrInvalidSessionID .
	ErrInvalidSessionID = errors.New("invalid session id")
	// ErrMissingService .
	ErrMissingService = errors.New("missing service")
	// ErrInvalidService .
	ErrInvalidService = errors.New("invalid service")
	// ErrServiceLimit .
	ErrServiceLimit = errors.New("service limit")
	// ErrMissingPairRecord .
	ErrMissingPairRecord = errors.New("missing pair record")
	// ErrSavePairRecordFailed .
	ErrSavePairRecordFailed = errors.New("save pair record failed")
	// ErrInvalidPairRecord .
	ErrInvalidPairRecord = errors.New("invalid pair record")
	// ErrInvalidActivationPeriod .
	ErrInvalidActivationPeriod = errors.New("invalid activation period")
	// ErrMissingActivationPeriod .
	ErrMissingActivationPeriod = errors.New("missing activation period")
	// ErrServiceProhibited .
	ErrServiceProhibited = errors.New("service prohibited")
	// ErrEscrowLocked .
	ErrEscrowLocked = errors.New("escrow locked")
	// ErrPairingProhibitedOverThisConnection .
	ErrPairingProhibitedOverThisConnection = errors.New("pairing prohibited over this connection")
	// ErrFMIPProtected .
	ErrFMIPProtected = errors.New("FMIP protected")
	// ErrMCProtected .
	ErrMCProtected = errors.New("MC Protected")
	// ErrMCChallengeRequired .
	ErrMCChallengeRequired = errors.New("mc challenge required")
)
