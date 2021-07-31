package config

import (
	"path/filepath"
	"runtime"
	"time"
)

const (
	TaxRate = 10
)

var (
	_, b, _, _ = runtime.Caller(0)
	PrjRoot    = filepath.Dir(b) + "/../.."
	Location   = time.FixedZone("Asia/Tokyo", 9*60*60)
)