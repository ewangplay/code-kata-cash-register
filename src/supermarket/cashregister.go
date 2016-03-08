package supermarket


type CashRegister struct {

}

func NewCashRegister() *CashRegister {
    return &CashRegister{}
}

func (this *CashRegister)SettleGoods() error {
    return nil
}
