package supermarket

type ShoppingCart struct {
}

func NewShoppingCart() *ShoppingCart {
    return &ShoppingCart{}
}

func (this *ShoppingCart) AddGoods(goods_barcode string, goods_num int) error {
    return nil
}
