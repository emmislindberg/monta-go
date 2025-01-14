package monta

import "time"

// Charge is a charging transaction.
type Charge struct {

	// ID of the charge
	ID int64 `json:"id"`

	// ID of the charge point related to this charge
	ChargePointID int64 `json:"chargePointId"`

	// CreatedAt is the creation date of the charge.
	CreatedAt time.Time `json:"createdAt"`

	// Date when cable was plugged in
	CablePluggedInAt *time.Time `json:"cablePluggedInAt"`

	// Date when charge started
	StartedAt *time.Time `json:"startedAt"`

	// Date when charge stopped
	StoppedAt *time.Time `json:"stoppedAt"`

	// Date when EV was fully charged
	FullyChargedAt *time.Time `json:"fullyChargedAt"`

	// Date when charge failed
	FailedAt *time.Time `json:"failedAt"`

	// Date when charge timed out
	TimeoutAt *time.Time `json:"timeoutAt"`

	// State of the charge
	State ChargeState `json:"state"`

	// Consumed Kwh
	ConsumedKWh *float64 `json:"consumedKwh"`

	// Kwh of the meter before charging started
	StartMeterKWh *float64 `json:"startMeterKwh"`

	// Kwh of the meter after charging stopped
	EndMeterKWh *float64 `json:"endMeterKwh"`

	// Price for this charge
	Price *float64 `json:"price"`

	// Average price per Kwh
	AveragePricePerKWh *float64 `json:"averagePricePerKwh"`

	// Average CO2 consumption per Kwh
	AverageCo2PerKWh *float64 `json:"averageCo2PerKwh"`

	// Average percentage of renewable energy per Kwh
	AverageRenewablePerKWh *float64 `json:"averageRenewablePerKwh"`

	// Failure reason for this charge
	FailureReason string `json:"failureReason"`

	// Payment method for this charge
	PaymentMethod PaymentMethod `json:"paymentMethod"`

	// A note taken for this charge
	Note string `json:"note"`

	// Configured Kwh limit for this charge
	KWhLimit *float64 `json:"kwhLimit"`

	// Currency for paying the charge.
	Currency *Currency `json:"currency"`

	// PayingTeam is the team paying for the charge.
	PayingTeam *Team `json:"payingTeam"`
}
