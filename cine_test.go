package main

import "testing"

func TestCalculateNetIncome(t *testing.T) {
	t.Run("Un ticket general", func(t *testing.T) {
		var tickets []Ticket
		tickets = append(tickets, TicketGeneral{basePrice: Pesos(400)})
		got := CalculateNetIncome(tickets)
		want := Pesos(400)

		assertPesos(t, got, want)
	})

	t.Run("Ocho tickets generales", func(t *testing.T) {
		var tickets []Ticket
		aTicket := TicketGeneral{Pesos(400)}
		for i := 0; i < 8; i++ {
			tickets = append(tickets, aTicket)
		}
		got := CalculateNetIncome(tickets)
		want := Pesos(400 * 8)

		assertPesos(t, got, want)
	})

	t.Run("Ocho tickets invitados", func(t *testing.T) {
		var tickets []Ticket
		aTicket := TicketInvitado{}
		for i := 0; i < 8; i++ {
			tickets = append(tickets, aTicket)
		}
		got := CalculateNetIncome(tickets)
		want := Pesos(0)

		assertPesos(t, got, want)
	})

	t.Run("Un ticket jubilado", func(t *testing.T) {
		var tickets []Ticket
		tickets = append(tickets, TicketJubilado{basePrice: Pesos(400), discount: 0.5})
		got := CalculateNetIncome(tickets)
		want := Pesos(200)

		assertPesos(t, got, want)
	})

	t.Run("Dos tickets de cada tipo", func(t *testing.T) {
		var tickets []Ticket
		ticketBasePrice := Pesos(500)
		tickets = append(tickets, TicketJubilado{basePrice: ticketBasePrice, discount: 0.5})
		tickets = append(tickets, TicketJubilado{basePrice: ticketBasePrice, discount: 0.5})
		tickets = append(tickets, TicketInvitado{})
		tickets = append(tickets, TicketInvitado{})
		tickets = append(tickets, TicketGeneral{basePrice: ticketBasePrice})
		tickets = append(tickets, TicketGeneral{basePrice: ticketBasePrice})

		got := CalculateNetIncome(tickets)
		want := Pesos(1500)

		assertPesos(t, got, want)
	})
}

func assertPesos(t *testing.T, got, want Pesos) {
	t.Helper()

	if got != want {
		t.Errorf("got %v want %v", got, want)
	}
}
