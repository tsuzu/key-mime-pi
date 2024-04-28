package templates

import "embed"

//go:embed index.html css/* js/* third-party/*
var FS embed.FS
