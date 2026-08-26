package migrations

import "embed"

// Files contains SQL migrations applied in filename order.
//
//go:embed *.sql
var Files embed.FS
