package supermarket

type SuperMarket struct {
	name       string
	goods_list map[string]*Goods
}

func NewSuperMarket(name string) *SuperMarket {
	return &SuperMarket{
		name:       name,
		goods_list: make(map[string]*Goods, 0),
	}
}

func (this *SuperMarket) RegisterGoods(goods *Goods) error {
	this.goods_list[goods.BarCode] = goods

	return nil
}

func (this *SuperMarket) GetName() string {
	return this.name
}

func (this *SuperMarket) QueryGoods(barcode string) *Goods {
	goods, _ := this.goods_list[barcode]
	return goods
}
