package 中介者Mediator

import "testing"

func TestNewMediator(t *testing.T) {
	mediator := NewMediator(&Stock{
		Stock: 10,
	})
	for i := 0; i < 10; i++ {
		seller := NewSellColleague(mediator)
		seller.Sell()
		purchaser := NewPurchaseColleague(mediator)
		purchaser.Purchase()
		println("stock:", mediator.CheckStock())
	}
}
