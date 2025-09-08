package main

import (
	_ "shop-goframe-micro-service/app/api-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"shop-goframe-micro-service/app/api-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
