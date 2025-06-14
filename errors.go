package kv_project

import "errors"

var (
	ErrKeyIsEmpty             = errors.New("key is empty")
	ErrIndexUpdateFailed      = errors.New("failed to update index")
	ErrKeyNotFound            = errors.New("key not found in database")
	ErrDataFileNotFound       = errors.New("data file is not found")
	ErrDataDirectoryCorrupted = errors.New("the database directory maybe corrupted")
	ErrExceedMaxBatchNum      = errors.New("exceed the max batch number")
	ErrMergeIsProgress        = errors.New("merge is in progress, please try again later")
	ErrDatabaseIsUsing        = errors.New("the database directory is used by another process")
	ErrMergeRationUnreached   = errors.New("the merge ration is unreachable")
	ErrNoEnoughSpaceForMerge  = errors.New("no enough space for merge")
)
