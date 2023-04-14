package model

type CBXbondQuoteMarketData struct {
	/**
	 * 债券代码
	 * 必传
	 */
	SecurityId string

	/**
	 * 业务发生时间
	 */
	Transactime int64

	/**
	 * 可成交总量（元）
	 */
	TradeVolume int64

	/**
	 * 开盘净价
	 */
	OpenPx string

	/**
	 * 开盘到期（%）
	 */
	OpenYield string


	//****************************买一档
	/**
	 * 清算速度
	 */
	BuySettlType1 string

	/**
	 * 报买净价
	 */
	BuyPx1 string

	/**
	 * 买入未匹配量
	 */
	BuyUnMatchQty int64

	/**
	 * 最优行情时表示报买量（元）；
	 * 深度行情时表示报买可成交量（元）
	 */
	BuySize1 int64

	/**
	 * 报买到期（%）
	 */
	BuyMaturityYield1 string


	//****************************买二档
	/**
	 * 清算速度
	 */
	BuySettlType2 string

	/**
	 * 报买净价
	 */
	BuyPx2 string

	/**
	 * 最优行情时表示报买量（元）；
	 * 深度行情时表示报买可成交量（元）
	 */
	BuySize2 int64

	/**
	 * 报买到期（%）
	 */
	BuyMaturityYield2 string


	//****************************买三档
	/**
	 * 清算速度
	 */
	BuySettlType3 string

	/**
	 * 报买净价
	 */
	BuyPx3 string

	/**
	 * 最优行情时表示报买量（元）；
	 * 深度行情时表示报买可成交量（元）
	 */
	BuySize3 int64

	/**
	 * 报买到期（%）
	 */
	BuyMaturityYield3 string


	//****************************买四档
	/**
	 * 清算速度
	 */
	BuySettlType4 string

	/**
	 * 报买净价
	 */
	BuyPx4 string

	/**
	 * 最优行情时表示报买量（元）；
	 * 深度行情时表示报买可成交量（元）
	 */
	BuySize4 int64

	/**
	 * 报买到期（%）
	 */
	BuyMaturityYield4 string


	//****************************买五档
	/**
	 * 清算速度
	 */
	BuySettlType5 string

	/**
	 * 报买净价
	 */
	BuyPx5 string

	/**
	 * 最优行情时表示报买量（元）；
	 * 深度行情时表示报买可成交量（元）
	 */
	BuySize5 int64

	/**
	 * 报买到期（%）
	 */
	BuyMaturityYield5 string


	//****************************卖一档
	/**
	 * 清算速度
	 */
	SellSettlType1 string

	/**
	 * 报卖净价
	 */
	SellPx1 string

	/**
	 * 卖入未匹配量
	 */
	SellUnMatchQty int64

	/**
	 * 最优行情时表示报卖量（元）；
	 * 深度行情时表示报卖可成交量（元）
	 */
	SellSize1 int64

	/**
	 * 报卖到期（%）
	 */
	SellMaturityYield1 string

	//****************************卖二档
	/**
	 * 清算速度
	 */
	SellSettlType2 string

	/**
	 * 报卖净价
	 */
	SellPx2 string

	/**
	 * 最优行情时表示报卖量（元）；
	 * 深度行情时表示报卖可成交量（元）
	 */
	SellSize2 int64

	/**
	 * 报卖到期（%）
	 */
	sellMaturityYield2 string

	//****************************卖三档
	/**
	 * 清算速度
	 */
	SellSettlType3 string

	/**
	 * 报卖净价
	 */
	SellPx3 string

	/**
	 * 最优行情时表示报卖量（元）；
	 * 深度行情时表示报卖可成交量（元）
	 */
	SellSize3 int64

	/**
	 * 报卖到期（%）
	 */
	SellMaturityYield3 string

	//****************************卖四档
	/**
	 * 清算速度
	 */
	SellSettlType4 string

	/**
	 * 报卖净价
	 */
	SellPx4 string

	/**
	 * 最优行情时表示报卖量（元）；
	 * 深度行情时表示报卖可成交量（元）
	 */
	SellSize4 int64

	/**
	 * 报卖到期（%）
	 */
	SellMaturityYield4 string

	//****************************卖五档
	/**
	 * 清算速度
	 */
	SellSettlType5 string

	/**
	 * 报卖净价
	 */
	SellPx5 string

	/**
	 * 最优行情时表示报卖量（元）；
	 * 深度行情时表示报卖可成交量（元）
	 */
	SellSize5 int64

	/**
	 * 报卖到期（%）
	 */
	SellMaturityYield5 string
}