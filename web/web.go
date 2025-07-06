package web

import (
	"embed"
)

//go:embed template/*
var EMFS embed.FS
