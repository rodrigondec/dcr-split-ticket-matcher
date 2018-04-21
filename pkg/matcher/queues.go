package matcher

import "github.com/decred/dcrd/dcrutil"

// splitTicketQueue is the queue of participants waiting for a split ticket
// session. May be named or not.
type splitTicketQueue struct {
	priceProvider       TicketPriceProvider
	waitingParticipants []*addParticipantRequest
}

func newSplitTicketQueue(priceProvider TicketPriceProvider) *splitTicketQueue {
	return &splitTicketQueue{
		priceProvider: priceProvider,
	}
}

func (q *splitTicketQueue) enoughForNewSession() bool {

	ticketPrice := dcrutil.Amount(q.priceProvider.CurrentTicketPrice())
	var availableSum uint64

	for _, r := range q.waitingParticipants {
		availableSum += r.maxAmount
	}

	ticketFee := SessionFeeEstimate(len(q.waitingParticipants))
	neededAmount := uint64(ticketPrice + ticketFee)
	return availableSum > neededAmount
}

func (q *splitTicketQueue) addWaitingParticipant(p *addParticipantRequest) {
	q.waitingParticipants = append(q.waitingParticipants, p)
}

func (q *splitTicketQueue) removeWaitingParticipant(p *addParticipantRequest) {
	waiting := q.waitingParticipants
	for i, ep := range waiting {
		if ep == p {
			q.waitingParticipants = append(waiting[:i], waiting[i+1:]...)
			break
		}
	}
}

func (q *splitTicketQueue) waitingAmounts() []dcrutil.Amount {
	res := make([]dcrutil.Amount, len(q.waitingParticipants))
	for i, p := range q.waitingParticipants {
		res[i] = dcrutil.Amount(p.maxAmount)
	}
	return res
}

func (q *splitTicketQueue) empty() bool {
	return len(q.waitingParticipants) == 0
}
