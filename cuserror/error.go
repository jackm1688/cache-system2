package cuserror

import (
	"errors"
)

var(
	ErrNotFoundPath = errors.New("not found log_path")
	ErrNotFoundName  = errors.New("not found log_name")
	ErrNotFoundLevel = errors.New("not found log_level")
	ErrNotFondSize   = errors.New("not found log_chan_size")
)
