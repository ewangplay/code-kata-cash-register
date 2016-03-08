package supermarket

import (
    "testing"
    "os"
)

func TestMain(m *testing.M) {
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

