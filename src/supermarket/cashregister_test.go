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

    cr.SettleGoods()

    t.Log("hello")
}

