package main

type Ticket interface {
	Price() Pesos
}

type TicketGeneral struct {
	basePrice Pesos
}

func (tg TicketGeneral) Price() Pesos {
	return tg.basePrice
}

type TicketJubilado struct {
	basePrice Pesos
	discount  float64
}

func (tj TicketJubilado) Price() Pesos {
	return tj.basePrice * Pesos(tj.discount)
}

type TicketInvitado struct {
}

func (ti TicketInvitado) Price() Pesos {
	return Pesos(0)
}

func CalculateNetIncome(tickets []Ticket) (income Pesos) {
	for _, ticket := range tickets {
		income += ticket.Price()
	}
	return
}
