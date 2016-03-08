package supermarket

import (
    "testing"
    "os"
)

var t_supermarket *SuperMarket

func TestMain(m *testing.M) {
    var goods *Goods

    //create one nwe supermarket
    t_supermarket = NewSuperMarket()

    //register goods 111
    goods = NewGoods("111", "苹果", "斤", 5.5, "水果")
    t_supermarket.RegisterGoods(goods)

    os.Exit(m.Run())
}

func TestSettleGoods(t *testing.T) {
    cr := NewCashRegister()

    sc := NewShoppingCart() 
    sc.AddGoods("111", 3)
    sc.AddGoods("222", 2)
    sc.AddGoods("333", 5)

    cr.ScanShoppingCar(sc)

    cr.SettleGoods()

    t.Log("hello")
}

