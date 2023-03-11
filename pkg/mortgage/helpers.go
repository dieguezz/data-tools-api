package amortizations

import (
	pb "github.com/dieguezz/nif-tools/proto"
)

func MapToAmortizations(capital float64, terms int, interestType float64, amortizationAmount float64, amortizationYear int, amortizationMonth int) *pb.MortgageAmortizationResponse {
	interestSavingsForPrice,
		monthlyPrice,
		pendingPayments,
		timeSavingsYear,
		timeSavingsMonth,
		totalTimeInterest,

		fees := CalcMortgageAmortization(
		float64(capital),
		int(terms),
		float64(interestType),
		float64(amortizationAmount),
		int(amortizationYear),
		int(amortizationMonth))
	var mappedFees []*pb.Fee
	for _, fee := range fees {
		mapped := pb.Fee{}
		mapped.Year = int32(fee.Year)
		mapped.Amortization = fee.Amortization
		mapped.AmortizationForTime = fee.AmortizationFortime
		mapped.Interest = fee.Interest
		mapped.InterestForTime = fee.InterestForTime
		mapped.Month = int32(fee.Month)
		mapped.PendingCapital = fee.PendingCapital
		mapped.PendingCapitalForTime = fee.PendingCapitalForTime
		mapped.Price = fee.Price
		mappedFees = append(mappedFees, &mapped)
	}

	return &pb.MortgageAmortizationResponse{
		Fees:                    mappedFees,
		MonthlyPrice:            int32(monthlyPrice),
		PendingPayments:         int32(pendingPayments),
		TimeSavingsYear:         int32(timeSavingsYear),
		TimeSavingsMonth:        int32(timeSavingsMonth),
		TotalTimeInterest:       int32(totalTimeInterest),
		InterestSavingsForPrice: int32(interestSavingsForPrice),
	}
}
