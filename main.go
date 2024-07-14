package main

import (
	_ "cert-auto-cdn/internal/packed"

	_ "cert-auto-cdn/internal/logic"

	"github.com/gogf/gf/v2/os/gctx"

	"cert-auto-cdn/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
