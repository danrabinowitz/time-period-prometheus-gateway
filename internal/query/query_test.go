package query

import "testing"

func Test_New_current_calendar_month(t *testing.T) {
	queryTemplate := "sum(increase(unifipoller_device_wan_receive_bytes_total[$DURATION]))"
	period := "current_calendar_month"
	_, err := New(queryTemplate, period)

	if err != nil {
		t.Errorf("New returned an unexpected error: %v", err)
	}
}

func Test_New_invalid_period(t *testing.T) {
	queryTemplate := "sum(increase(unifipoller_device_wan_receive_bytes_total[$DURATION]))"
	period := "invalid_period"
	_, err := New(queryTemplate, period)

	if err == nil {
		t.Error("New did not return an error")
	}
}
