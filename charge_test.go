package monta

import (
	"encoding/json"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestCharge_MarshalJSON(t *testing.T) {
	expected := strings.TrimSpace(`
{
  "id": 1,
  "chargePointId": 21,
  "createdAt": "2022-05-12T15:56:45.999189Z",
  "cablePluggedInAt": "2022-05-12T15:56:45.999189Z",
  "startedAt": "2022-05-12T15:56:45.999189Z",
  "stoppedAt": "2022-05-12T15:56:45.999189Z",
  "fullyChargedAt": "2022-05-12T15:56:45.999189Z",
  "failedAt": "2022-05-12T15:56:45.999189Z",
  "timeoutAt": "2022-05-12T15:56:45.999189Z",
  "state": "charging",
  "consumedKwh": 20.4,
  "startMeterKwh": 123.45,
  "endMeterKwh": 163.85,
  "price": 122.4,
  "averagePricePerKwh": 6,
  "averageCo2PerKwh": 100,
  "averageRenewablePerKwh": 72.5,
  "failureReason": "Some reason why we couldn't charge.",
  "paymentMethod": "free",
  "note": "Lorem Ipsum",
  "kwhLimit": 21,
  "currency": {
    "identifier": "DKK",
    "name": "Danish krone",
    "decimals": 2
  },
  "payingTeam": {
    "id": 14,
    "publicName": "Monta HQ"
  }
}
	`)
	var charge Charge
	assert.NilError(t, json.Unmarshal([]byte(expected), &charge))
	actual, err := json.MarshalIndent(&charge, "", "  ")
	assert.NilError(t, err)
	assert.Equal(t, expected, string(actual))
}
