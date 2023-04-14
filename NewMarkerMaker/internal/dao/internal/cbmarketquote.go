// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CBMARKETQUOTEDao is the data access object for table CBMARKETQUOTE.
type CBMARKETQUOTEDao struct {
	table   string               // table is the underlying table name of the DAO.
	group   string               // group is the database configuration group name of current DAO.
	columns CBMARKETQUOTEColumns // columns contains all the column names of Table for convenient usage.
}

// CBMARKETQUOTEColumns defines and stores column names for table CBMARKETQUOTE.
type CBMARKETQUOTEColumns struct {
	ID                     string //
	BUSSINESSCODE          string //
	REMARK                 string //
	USERID                 string //
	BUYACCRUEDINTERESTAMT  string //
	BUYCLEARINGMETHOD      string //
	BUYDELIVERYTYPE        string //
	BUYORDERQTY            string //
	BUYPRICE               string //
	BUYSECURITYID          string //
	BUYSETTLCURRFXRATE     string //
	BUYSETTLCURRENCY       string //
	BUYSETTLTYPE           string //
	BUYSTRIKEYEILDYILED    string //
	BUYYTMYILED            string //
	CHANNEL                string //
	CHANNELTYPE            string //
	CLORDID                string //
	CONTINGENCYINDICATOR   string //
	DATASOURCESTRING       string //
	DATECONFIRMED          string //
	MARKETINDICATOR        string //
	MARKETSCOPE            string //
	MAXFLOOR               string //
	PARTYID101             string //
	PARTYID206             string //
	PARTYSUBID101803101    string //
	PARTYSUBID101803125    string //
	PARTYSUBID101803135    string //
	PARTYSUBID1018032      string //
	PARTYSUBID101803266    string //
	PARTYSUBID101803267    string //
	PARTYSUBID10180329     string //
	PARTYSUBID206803135    string //
	QUOTEID                string //
	QUOTESTATUS            string //
	QUOTETRANSTYPE         string //
	QUOTETYPE              string //
	ROUTINGTYPE            string //
	SECURITYDESC           string //
	SECURITYID             string //
	SECURITYTYPEID         string //
	SELLACCRUEDINTERESTAMT string //
	SELLCLEARINGMETHOD     string //
	SELLDELIVERYTYPE       string //
	SELLORDERQTY           string //
	SELLPRICE              string //
	SELLSECURITYID         string //
	SELLSETTLCURRFXRATE    string //
	SELLSETTLCURRENCY      string //
	SELLSETTLTYPE          string //
	SELLSTRIKEYEILDYILED   string //
	SELLYTMYILED           string //
	STATUS                 string //
	SYMBOL                 string //
	TERMTOMATURITYSTRING   string //
	TRANSACTTIME           string //
	USERREFERENCE1         string //
	USERREFERENCE2         string //
	USERREFERENCE3         string //
	USERREFERENCE4         string //
	USERREFERENCE5         string //
	USERREFERENCE6         string //
	VALIDUNTILTIME         string //
	QUOTESTATUSDESC        string //
	BUYLASTQTY             string //
	SELLLASTQTY            string //
}

// cBMARKETQUOTEColumns holds the columns for table CBMARKETQUOTE.
var cBMARKETQUOTEColumns = CBMARKETQUOTEColumns{
	ID:                     "ID",
	BUSSINESSCODE:          "BUSSINESSCODE",
	REMARK:                 "REMARK",
	USERID:                 "USERID",
	BUYACCRUEDINTERESTAMT:  "BUYACCRUEDINTERESTAMT",
	BUYCLEARINGMETHOD:      "BUYCLEARINGMETHOD",
	BUYDELIVERYTYPE:        "BUYDELIVERYTYPE",
	BUYORDERQTY:            "BUYORDERQTY",
	BUYPRICE:               "BUYPRICE",
	BUYSECURITYID:          "BUYSECURITYID",
	BUYSETTLCURRFXRATE:     "BUYSETTLCURRFXRATE",
	BUYSETTLCURRENCY:       "BUYSETTLCURRENCY",
	BUYSETTLTYPE:           "BUYSETTLTYPE",
	BUYSTRIKEYEILDYILED:    "BUYSTRIKEYEILDYILED",
	BUYYTMYILED:            "BUYYTMYILED",
	CHANNEL:                "CHANNEL",
	CHANNELTYPE:            "CHANNELTYPE",
	CLORDID:                "CLORDID",
	CONTINGENCYINDICATOR:   "CONTINGENCYINDICATOR",
	DATASOURCESTRING:       "DATASOURCESTRING",
	DATECONFIRMED:          "DATECONFIRMED",
	MARKETINDICATOR:        "MARKETINDICATOR",
	MARKETSCOPE:            "MARKETSCOPE",
	MAXFLOOR:               "MAXFLOOR",
	PARTYID101:             "PARTYID101",
	PARTYID206:             "PARTYID206",
	PARTYSUBID101803101:    "PARTYSUBID101_803_101",
	PARTYSUBID101803125:    "PARTYSUBID101_803_125",
	PARTYSUBID101803135:    "PARTYSUBID101_803_135",
	PARTYSUBID1018032:      "PARTYSUBID101_803_2",
	PARTYSUBID101803266:    "PARTYSUBID101_803_266",
	PARTYSUBID101803267:    "PARTYSUBID101_803_267",
	PARTYSUBID10180329:     "PARTYSUBID101_803_29",
	PARTYSUBID206803135:    "PARTYSUBID206_803_135",
	QUOTEID:                "QUOTEID",
	QUOTESTATUS:            "QUOTESTATUS",
	QUOTETRANSTYPE:         "QUOTETRANSTYPE",
	QUOTETYPE:              "QUOTETYPE",
	ROUTINGTYPE:            "ROUTINGTYPE",
	SECURITYDESC:           "SECURITYDESC",
	SECURITYID:             "SECURITYID",
	SECURITYTYPEID:         "SECURITYTYPEID",
	SELLACCRUEDINTERESTAMT: "SELLACCRUEDINTERESTAMT",
	SELLCLEARINGMETHOD:     "SELLCLEARINGMETHOD",
	SELLDELIVERYTYPE:       "SELLDELIVERYTYPE",
	SELLORDERQTY:           "SELLORDERQTY",
	SELLPRICE:              "SELLPRICE",
	SELLSECURITYID:         "SELLSECURITYID",
	SELLSETTLCURRFXRATE:    "SELLSETTLCURRFXRATE",
	SELLSETTLCURRENCY:      "SELLSETTLCURRENCY",
	SELLSETTLTYPE:          "SELLSETTLTYPE",
	SELLSTRIKEYEILDYILED:   "SELLSTRIKEYEILDYILED",
	SELLYTMYILED:           "SELLYTMYILED",
	STATUS:                 "STATUS",
	SYMBOL:                 "SYMBOL",
	TERMTOMATURITYSTRING:   "TERMTOMATURITYSTRING",
	TRANSACTTIME:           "TRANSACTTIME",
	USERREFERENCE1:         "USERREFERENCE1",
	USERREFERENCE2:         "USERREFERENCE2",
	USERREFERENCE3:         "USERREFERENCE3",
	USERREFERENCE4:         "USERREFERENCE4",
	USERREFERENCE5:         "USERREFERENCE5",
	USERREFERENCE6:         "USERREFERENCE6",
	VALIDUNTILTIME:         "VALIDUNTILTIME",
	QUOTESTATUSDESC:        "QUOTESTATUSDESC",
	BUYLASTQTY:             "BUYLASTQTY",
	SELLLASTQTY:            "SELLLASTQTY",
}

// NewCBMARKETQUOTEDao creates and returns a new DAO object for table data access.
func NewCBMARKETQUOTEDao() *CBMARKETQUOTEDao {
	return &CBMARKETQUOTEDao{
		group:   "default",
		table:   "CBMARKETQUOTE",
		columns: cBMARKETQUOTEColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CBMARKETQUOTEDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CBMARKETQUOTEDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *CBMARKETQUOTEDao) Columns() CBMARKETQUOTEColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *CBMARKETQUOTEDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CBMARKETQUOTEDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CBMARKETQUOTEDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
