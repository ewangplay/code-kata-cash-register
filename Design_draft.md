## SuperMarket

* Supermarket
    + GetName()
    + RegisterGoods(&Goods)

* CashRegister
	+ SettleGoods(&Shoppingcart)
	+ SetPreferentials(&Preferentials)
	
* ShoppingCart
	+ AddGoods(GoodsBarCode, GoodsNum)
    + GetGoodsList()
	
* Goods
	- Name
	- QuantityUnit
	- UnitPrice
	- Category
	- BarCode

* Preferentials
	+ AddPreferential(id，priority，goods_barcode_list)
		- BUY_TWO_GET_ONE_FREE
		- 95_DISCOUNT
