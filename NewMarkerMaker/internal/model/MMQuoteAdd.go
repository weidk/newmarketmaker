package model

type MMQuoteAdd struct {
	TYPE         string
	CLORDID      string
	QUOTEID      string
	QUOTESTATUS  string
	BONDCODE     string
	MAXFLOOR     int32
	BUYORDERQTY  int32
	BUYPRICE     float64
	BUYYTMYILED  float64
	BUYDEALQTY   float64
	SELLORDERQTY int32
	SELLPRICE    float64
	SELLYTMYILED float64
	SELLDEALQTY  float64
	SETTLTYPE    string
	ORDERTIME    string
	TRANSACTTIME string
	INFO         string
}
