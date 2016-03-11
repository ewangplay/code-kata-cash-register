package supermarket

import (
    "testing"
    "os"
)

var t_cashregister *CashRegister

func TestMain(m *testing.M) {
    //创建超市并登记商品
    supermarket := NewSuperMarket("没钱赚商店")
    supermarket.RegisterGoods(NewGoods("111", "苹果", "斤", 5.5, "水果"))
    supermarket.RegisterGoods(NewGoods("222", "可乐", "瓶", 3.0, "饮料"))
    supermarket.RegisterGoods(NewGoods("333", "康师傅方便面", "袋", 4.5, "食品"))

    //创建优惠信息
    preferentials := NewPreferentials()

    //创建收银机并关联超市，并设置优惠信息
    t_cashregister = NewCashRegister(supermarket)
    t_cashregister.SetPreferentials(preferentials)

    os.Exit(m.Run())
}

func TestSettleGoods(t *testing.T) {
    //添加购物车
    shoppingcart := NewShoppingCart() 
    shoppingcart.AddGoods("111", 3)
    shoppingcart.AddGoods("222", 2)
    shoppingcart.AddGoods("333", 5)

    //收银机结算
    t_cashregister.SettleGoods(shoppingcart)
}

