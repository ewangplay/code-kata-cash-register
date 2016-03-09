package supermarket

type SuperMarket struct {
    goods_list map[string]*Goods
}

func NewSuperMarket() *SuperMarket {
    return &SuperMarket {
        goods_list: make(map[string]*Goods, 0),
    }
}

func (this *SuperMarket) RegisterGoods(goods *Goods) error {
    return nil
}
