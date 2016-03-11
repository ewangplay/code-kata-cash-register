package supermarket

const (
    BUY_TWO_GET_ONE_FREE = 1
    DISCOUNT_95 = 2
)

type Preferentials struct {
}

func NewPreferentials() *Preferentials {
    return &Preferentials{
    }
}

func (this *Preferentials) AddPreferential(id int, priority int, goods_barcode_list []string) error {
    return nil
}

