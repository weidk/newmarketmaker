// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"NewMarkerMaker/internal/dao/internal"
)

// internalMMQUOTESETDao is internal type for wrapping internal DAO implements.
type internalMMQUOTESETDao = *internal.MMQUOTESETDao

// mMQUOTESETDao is the data access object for table MM_QUOTE_SET.
// You can define custom methods on it to extend its functionality as you wish.
type mMQUOTESETDao struct {
	internalMMQUOTESETDao
}

var (
	// MMQUOTESET is globally public accessible object for table MM_QUOTE_SET operations.
	MMQUOTESET = mMQUOTESETDao{
		internal.NewMMQUOTESETDao(),
	}
)

// Fill with you ideas below.
