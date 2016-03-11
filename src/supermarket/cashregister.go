package supermarket

import (
	"fmt"
)

type CashRegister struct {
	belongs_supermarket *SuperMarket
}

func NewCashRegister(belongs_supermarket *SuperMarket) *CashRegister {
	return &CashRegister{
		belongs_supermarket: belongs_supermarket,
	}
}

func (this *CashRegister) SetPreferentials(preferentials *Preferentials) error {
    return nil
}

/*
***<没钱赚商店>购物清单***
名称：可口可乐，数量：3瓶，单价：3.00(元)，小计：9.00(元)
名称：羽毛球，数量：5个，单价：1.00(元)，小计：5.00(元)
名称：苹果，数量：2斤，单价：5.50(元)，小计：11.00(元)
----------------------
总计：25.00(元)
**********************
*/
func (this *CashRegister) SettleGoods(shoppingcart *ShoppingCart) error {
	goods_list := shoppingcart.GetGoodsList()

	fmt.Printf("***<%v>购物清单***\n", this.belongs_supermarket.GetName())

    var sum float32

	for goods_barcode, goods_num := range goods_list {
		goods := this.belongs_supermarket.QueryGoods(goods_barcode)
		if goods != nil {
            amount := float32(goods_num)*goods.UnitPrice
			fmt.Printf("名称: %v, 数量：%v%v, 单价：%v(元), 小计: %v(元)\n", goods.Name, goods_num, goods.QuantityUnit, goods.UnitPrice, amount)

            sum += amount
		}
	}

	fmt.Println("----------------------")
    fmt.Printf("总计：%v(元)\n", sum)
	fmt.Println("**********************")

	return nil
}
