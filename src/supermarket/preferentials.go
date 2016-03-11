package supermarket

type PREFERENTIAL_ID int

const (
	NO_PREFERENTIAL PREFERENTIAL_ID = iota
	BUY_TWO_GET_ONE_FREE
	DISCOUNT_95
)

type Preferentials struct {
	data map[string]map[PREFERENTIAL_ID]int
}

func NewPreferentials() *Preferentials {
	return &Preferentials{
		data: make(map[string]map[PREFERENTIAL_ID]int, 0),
	}
}

func (this *Preferentials) AddPreferential(id PREFERENTIAL_ID, priority int, goods_barcode_list []string) error {
	for _, goods_barcode := range goods_barcode_list {
		_, found := this.data[goods_barcode]
		if !found {
			this.data[goods_barcode] = make(map[PREFERENTIAL_ID]int, 1)
		}
		this.data[goods_barcode][id] = priority
	}

	return nil
}

func (this *Preferentials) QueryGoodsPreferential(goods_barcode string) PREFERENTIAL_ID {
    id := NO_PREFERENTIAL
    priority := 0

	m, found := this.data[goods_barcode]
    if found {
        for i, p := range m {
            if p > priority {
                id = i
            }
        }
    }

	return id
}
