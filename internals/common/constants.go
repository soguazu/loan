package common

const (
	// CardTransactionBOTH transaction type for POS
	CardTransactionBOTH PricingIdentifier = "ngn-card-both"
	// CardTransactionATM transaction type for ATM
	CardTransactionATM PricingIdentifier = "ngn-card-atm"
	// CardCreation price identifier type for create card
	CardCreation PricingIdentifier = "ngn-card-create"
	// CardShipping price identifier type for card shipping
	CardShipping PricingIdentifier = "ngn-card-shipping"

	DebitTransaction  TransactionType = "debit"
	CreditTransaction TransactionType = "credit"
)

const (
	CustomerType   string = "company"
	CustomerStatus        = "active"
	Partner               = "sudo"
	Currency              = "NGN"
)

// PhysicalCardIdentifier list of charges for create physical card
var PhysicalCardIdentifier = []PricingIdentifier{CardShipping, CardCreation}

// VirtualCardIdentifier list of charges for create virtual card
var VirtualCardIdentifier = []PricingIdentifier{CardCreation}
