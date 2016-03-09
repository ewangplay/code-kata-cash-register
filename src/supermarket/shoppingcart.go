package supermarket

type ShoppingCart struct {
    goods_list map[string]int
}

func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{
        goods_list: make(map[string]int, 0),
    }
}

func (this *ShoppingCart) AddGoods(goods_barcode string, goods_num int) error {
    _, ok := this.goods_list[goods_barcode]
    if ok {
        this.goods_list[goods_barcode] += goods_num
    } else {
        this.goods_list[goods_barcode] = goods_num
    }

    return nil
}

func (this *ShoppingCart) GetGoodsList() map[string]int {
    return this.goods_list
}
