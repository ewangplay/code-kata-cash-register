package supermarket

type PREFERENTIAL_ID int

const (
	NO_PREFERENTIAL PREFERENTIAL_ID = iota
	BUY_TWO_GET_ONE_FREE
	DISCOUNT_95
)

type Preferentials struct {
}

func NewPreferentials() *Preferentials {
	return &Preferentials{}
}

func (this *Preferentials) AddPreferential(id PREFERENTIAL_ID, priority int, goods_barcode_list []string) error {
	return nil
}

func (this *Preferentials) QueryGoodsPreferential(goods_barcode string) PREFERENTIAL_ID {
    return NO_PREFERENTIAL
}
