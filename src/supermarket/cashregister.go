package supermarket


type CashRegister struct {
    belongs_supermarket *SuperMarket

}

func NewCashRegister(belongs_supermarket *SuperMarket) *CashRegister {
    return &CashRegister{
        belongs_supermarket: belongs_supermarket,
    }
}

func (this *CashRegister)SettleGoods() error {
    return nil
}

func (this *CashRegister) ScanShoppingCar(sc *ShoppingCart) error {
    return nil
}
