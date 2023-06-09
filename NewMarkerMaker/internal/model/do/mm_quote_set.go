// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// MMQUOTESET is the golang structure of table MM_QUOTE_SET for DAO operations like Where/Data.
type MMQUOTESET struct {
	g.Meta         `orm:"table:MM_QUOTE_SET, do:true"`
	BONDCODE       interface{} //
	BASEYIELD      interface{} //
	MIDYIELDADD    interface{} //
	MIDYIELDSUB    interface{} //
	BESTYIELDADD   interface{} //
	BESTYIELDSUB   interface{} //
	BESTYIELDTYPE  interface{} //
	SETYIELDBUY    interface{} //
	SETYIELDSELL   interface{} //
	BESTBID        interface{} //
	SECONDBID      interface{} //
	BESTOFR        interface{} //
	SECONDOFR      interface{} //
	SELFBID        interface{} //
	SELFOFR        interface{} //
	BESTMIDYIELD   interface{} //
	SECONDMIDYIELD interface{} //
	SETTLETYPE     interface{} //
	BUYQTY         interface{} //
	SELLQTY        interface{} //
	MAXFLOOR       interface{} //
	CANSEND        interface{} //
}
