package model

import (
	"time"
)

type XBondQuoteAdd struct {
	TYPE         string
	CLORDID      string
	QUOTEID      string
	QUOTESTATUS  string
	BONDCODE     string
	SIDE         string
	PRICE        float32
	YTMYILED     float32
	DEALQTY      float32
	DEALPRICE    float32
	SETTLTYPE    int32
	ORDERTIME    time.Time
	TRANSACTTIME time.Time
	INFO         string
}
