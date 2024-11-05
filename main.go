package main

import (
	"log/slog"

	"github.com/marcelbednarczyk/Golang-Dual-EEG/cortex"
)

func main() {
	slog.Info("Hello, World!")
	_ = cortex.GetDefaultInfoRequest()
}
