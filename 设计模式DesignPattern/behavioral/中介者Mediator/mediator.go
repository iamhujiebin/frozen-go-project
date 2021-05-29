package 中介者Mediator

import "fmt"

// 中介者模式
// 理解为中间人,例如房产中介,销售|采购|库存的中介
// 有两个抽奖类即可,Mediator:中介;Colleague:同事
// 下面用销售|采购|库存的中介做例子
type Mediator interface {
	CheckStock() int
	AfterPurchase()
	AfterSale()
}

type Colleague interface {
	SetMediator(mediator Mediator)
}

func NewSellColleague(mediator Mediator) *SellColleague {
	return &SellColleague{Mediator: mediator}
}

func NewPurchaseColleague(mediator Mediator) *PurchaseColleague {
	return &PurchaseColleague{Mediator: mediator}
}

type SellColleague struct {
	Mediator Mediator
}

func (s *SellColleague) SetMediator(mediator Mediator) {
	s.Mediator = mediator
}

func (s *SellColleague) Sell() {
	if s.Mediator.CheckStock() > 0 {
		s.Mediator.AfterSale()
		fmt.Println("do Sell")
	} else {
		fmt.Println("no stock")
	}
}

type PurchaseColleague struct {
	Mediator Mediator
}

func (p *PurchaseColleague) SetMediator(mediator Mediator) {
	p.Mediator = mediator
}

func (p *PurchaseColleague) Purchase() {
	if p.Mediator.CheckStock() < 5 {
		fmt.Println("do Purchase")
		p.Mediator.AfterPurchase()
	} else {
		fmt.Println("stock enough")
	}
}

type Stock struct {
	Stock int
}

func (s *Stock) CheckStock() int {
	return s.Stock
}

func (s *Stock) AddStock() {
	s.Stock++
}

func (s *Stock) DecrStock() bool {
	if s.Stock <= 0 {
		return false
	}
	s.Stock--
	return true
}

type ConcreteMediator struct {
	StockColleague *Stock
}

func (c *ConcreteMediator) AfterPurchase() {
	c.StockColleague.AddStock()
}

func (c *ConcreteMediator) AfterSale() {
	c.StockColleague.DecrStock()
}

func (c *ConcreteMediator) CheckStock() int {
	return c.StockColleague.CheckStock()
}

func NewMediator(colleague *Stock) *ConcreteMediator {
	return &ConcreteMediator{
		StockColleague: colleague,
	}
}
