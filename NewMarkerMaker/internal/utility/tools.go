package utility

import (
	"NewMarkerMaker/internal/consts"
	"github.com/gogf/gf/v2/frame/g"
)

func Recov(){
	if p := recover(); p != nil {
		g.Log().Error(consts.Ctx, p)
	}
}
