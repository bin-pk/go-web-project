package core

type CuteErrorCode int

const (
	InternalError CuteErrorCode = iota
	NullPointer
	NotFound
	Cancelled
	PermissionDenied
	DeadlineExceeded
)
