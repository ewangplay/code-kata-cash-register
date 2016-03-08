package supermarket

type SuperMarket struct {
}

func NewSuperMarket() *SuperMarket {
    return &SuperMarket {
    }
}

func (this *SuperMarket) RegisterGoods(goods *Goods) error {
    return nil
}
