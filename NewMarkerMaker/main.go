package main

import (
	"NewMarkerMaker/internal/cmd"
	_ "github.com/gogf/gf/contrib/drivers/oracle/v2"
	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	cmd.Main.Run(gctx.New())
}
