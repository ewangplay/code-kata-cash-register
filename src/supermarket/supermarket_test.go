package supermarket

import (
	"os"
	"testing"
)

var t_cashregister *CashRegister

func TestMain(m *testing.M) {
	//创建超市并登记商品
	supermarket := NewSuperMarket("没钱赚商店")
	supermarket.RegisterGoods(NewGoods("ITEM000001", "苹果", "斤", 5.5, "水果"))
	supermarket.RegisterGoods(NewGoods("ITEM000002", "可口可乐", "瓶", 3.0, "饮料"))
	supermarket.RegisterGoods(NewGoods("ITEM000003", "康师傅方便面", "袋", 4.5, "食品"))
	supermarket.RegisterGoods(NewGoods("ITEM000004", "羽毛球", "个", 1.5, "体育用品"))
	supermarket.RegisterGoods(NewGoods("ITEM000005", "西瓜", "斤", 6, "水果"))

	//创建优惠信息
	preferentials := NewPreferentials()
	preferentials.AddPreferential(BUY_TWO_GET_ONE_FREE, 2, []string{"ITEM000004"})
	preferentials.AddPreferential(DISCOUNT_95, 1, []string{"ITEM000001"})

	//创建收银机并关联超市，并设置优惠信息
	t_cashregister = NewCashRegister(supermarket)
	t_cashregister.SetPreferentials(preferentials)

	os.Exit(m.Run())
}

func TestNotBuyGoods(t *testing.T) {
	//添加购物车
	shoppingcart := NewShoppingCart()

	//收银机结算
	t_cashregister.SettleGoods(shoppingcart)
}

func TestNoPreferentialGoods(t *testing.T) {
	//添加购物车
	shoppingcart := NewShoppingCart()
	shoppingcart.AddGoods("ITEM000002", 3)
	shoppingcart.AddGoods("ITEM000003", 2)
	shoppingcart.AddGoods("ITEM000005", 1)

	//收银机结算
	t_cashregister.SettleGoods(shoppingcart)
}

func TestBuyTwoGetOneFreePreferentialGoods(t *testing.T) {
	//添加购物车
	shoppingcart := NewShoppingCart()
	shoppingcart.AddGoods("ITEM000002", 3)
	shoppingcart.AddGoods("ITEM000003", 2)
	shoppingcart.AddGoods("ITEM000004", 5)

	//收银机结算
	t_cashregister.SettleGoods(shoppingcart)
}

func TestDiscount95PreferentialGoods(t *testing.T) {
	//添加购物车
	shoppingcart := NewShoppingCart()
	shoppingcart.AddGoods("ITEM000001", 3)
	shoppingcart.AddGoods("ITEM000003", 3)
	shoppingcart.AddGoods("ITEM000005", 1)

	//收银机结算
	t_cashregister.SettleGoods(shoppingcart)
}

func TestBuyTwoGetOneFreeAndDiscount95PreferentialGoods(t *testing.T) {
	//添加购物车
	shoppingcart := NewShoppingCart()
	shoppingcart.AddGoods("ITEM000001", 3)
	shoppingcart.AddGoods("ITEM000002", 3)
	shoppingcart.AddGoods("ITEM000004", 5)

	//收银机结算
	t_cashregister.SettleGoods(shoppingcart)
}
