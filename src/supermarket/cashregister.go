package supermarket


type CashRegister struct {

}

func NewCashRegister() *CashRegister {
    return &CashRegister{}
}

func (this *CashRegister)SettleGoods() error {
    return nil
}

func (this *CashRegister) ScanShoppingCar(sc *ShoppingCart) error {
    return nil
}
