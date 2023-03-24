package constants

var OrderStatus = struct {
	PLACED     string
	DISPATCHED string
	COMPLETED  string
	CANCELLED  string
}{
	PLACED:     "placed",
	DISPATCHED: "dispatched",
	COMPLETED:  "completed",
	CANCELLED:  "cancelled",
}
