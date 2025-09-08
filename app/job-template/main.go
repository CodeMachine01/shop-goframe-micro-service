package main

import (
	_ "shop-goframe-micro-service/app/job-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"shop-goframe-micro-service/app/job-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
