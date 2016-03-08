package supermarket

type Goods struct {
    BarCode string
    Name string
    QuantityUnit string
    UnitPrice float32
    Category string
}

func NewGoods(barcode string, name string, quantity_unit string, unit_price float32, category string) *Goods {
    return &Goods {
        BarCode: barcode,
        Name: name,
        QuantityUnit: quantity_unit,
        UnitPrice: unit_price,
        Category: category,
    }
}
