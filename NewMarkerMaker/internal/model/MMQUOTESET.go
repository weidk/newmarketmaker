package model

type MMQUOTESET struct {
	BONDCODE      string  `json:"BONDCODE"      ` //
	MIDYIELDADD   float64 `json:"MIDYIELDADD"   ` //
	MIDYIELDSUB   float64 `json:"MIDYIELDSUB"   ` //
	BESTYIELDADD  float64 `json:"BESTYIELDADD"  ` //
	BESTYIELDSUB  float64 `json:"BESTYIELDSUB"  ` //
	BESTYIELDTYPE int     `json:"BESTYIELDTYPE" ` //
	SETYIELDBUY   float64 `json:"SETYIELDBUY"   ` //
	SETYIELDSELL  float64 `json:"SETYIELDSELL"  ` //

	BASEYIELD float64 `json:"BASEYIELD"     ` //

	BESTBID   float64 `json:"BESTBID"       ` //
	SECONDBID float64 `json:"SECONDBID"     ` //

	BESTOFR   float64 `json:"BESTOFR"       ` //
	SECONDOFR float64 `json:"SECONDOFR"     ` //

	BESTMIDYIELD   float64 `json:"BESTMIDYIELD"   ` //
	SECONDMIDYIELD float64 `json:"SECONDMIDYIELD" ` //

	SELFBID float64 `json:"SELFBID"       ` //
	SELFOFR float64 `json:"SELFOFR"       ` //

	SETTLETYPE string `json:"SETTLETYPE"     ` //

	BUYQTY   int32 `json:"BUYQTY"         ` //
	SELLQTY  int32 `json:"SELLQTY"        ` //
	MAXFLOOR int32 `json:"MAXFLOOR"       ` //

	LASTBID float64
	LASTOFR float64
	IsNew   bool
	CANSEND bool
}
