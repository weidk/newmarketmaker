package model

import "github.com/gogf/gf/v2/os/gtime"

type CBMARKETQUOTESW struct {
	ID               int         `json:"iD"               ` //
	MDBOOKTYPE       string      `json:"mDBOOKTYPE"       ` //
	MDENTRYDATE      string      `json:"mDENTRYDATE"      ` //
	MDENTRYPX        string      `json:"mDENTRYPX"        ` //
	MDENTRYSIZE      string      `json:"mDENTRYSIZE"      ` //
	MDENTRYTIME      string      `json:"mDENTRYTIME"      ` //
	MDENTRYTYPE      string      `json:"mDENTRYTYPE"      ` //
	MDPRICELEVEL     string      `json:"mDPRICELEVEL"     ` //
	MDQUOTETYPE      string      `json:"mDQUOTETYPE"      ` //
	MDSUBBOOKTYPE    string      `json:"mDSUBBOOKTYPE"    ` //
	CLEARINGMETHOD   string      `json:"cLEARINGMETHOD"   ` //
	DELIVERYTYPE     string      `json:"dELIVERYTYPE"     ` //
	LASTPX           string      `json:"lASTPX"           ` //
	MARKETDEPTH      string      `json:"mARKETDEPTH"      ` //
	MARKETINDICATOR  string      `json:"mARKETINDICATOR"  ` //
	MATURITYYIELD    float64     `json:"mATURITYYIELD"    ` //
	MDID             string      `json:"mDID"             ` //
	MDREQID          string      `json:"mDREQID"          ` //
	PARTYID101       string      `json:"pARTYID101"       ` //
	PARTYID1018032   string      `json:"pARTYID1018032"   ` //
	PARTYID101803266 string      `json:"pARTYID101803266" ` //
	QUOTEENTRYID     string      `json:"qUOTEENTRYID"     ` //
	SECURITYID       string      `json:"sECURITYID"       ` //
	SECURITYTYPE     string      `json:"sECURITYTYPE"     ` //
	SETTLCURRFXRATE  string      `json:"sETTLCURRFXRATE"  ` //
	SETTLCURRENCY    string      `json:"sETTLCURRENCY"    ` //
	SETTLDATE        string      `json:"sETTLDATE"        ` //
	SETTLTYPE        string      `json:"sETTLTYPE"        ` //
	SYMBOL           string      `json:"sYMBOL"           ` //
	INSERTTIME       *gtime.Time `json:"iNSERTTIME"       ` //
}
