package constant

type Role string

const (
	RoleBuyer  Role = "Buyer"
	RoleSeller Role = "Seller"
	RoleBroker Role = "Broker"
)

type TxStatus string

const (
	StatusPendingSellerConfirm TxStatus = "PendingSellerConfirm"
	StatusAwaitingPayment      TxStatus = "AwaitingPayment"
	StatusPaid                 TxStatus = "Paid"
	StatusShipped              TxStatus = "Shipped"
	StatusDelivered            TxStatus = "Delivered"
	StatusInspection           TxStatus = "Inspection"
	StatusClosed               TxStatus = "Closed"
)
