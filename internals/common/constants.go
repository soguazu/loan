package common

const (
	// CardTransactionPOS transaction type for POS
	CardTransactionPOS PricingIdentifier = "ngn-card-pos"
	// CardTransactionWEB transaction type for WEB
	CardTransactionWEB PricingIdentifier = "ngn-card-web"
	// CardTransactionATM transaction type for ATM
	CardTransactionATM PricingIdentifier = "ngn-card-atm"
	// CardCreation price identifier type for create card
	CardCreation PricingIdentifier = "ngn-card-create"
	// CardShipping price identifier type for card shipping
	CardShipping PricingIdentifier = "ngn-card-shipping"

	CustomerType   = "company"
	CustomerStatus = "active"
	Partner        = "sudo"

	DebitTransaction  TransactionType = "debit"
	CreditTransaction TransactionType = "credit"
)

// PhysicalCardIdentifier list of charges for create physical card
var PhysicalCardIdentifier = []PricingIdentifier{CardShipping, CardCreation}

// VirtualCardIdentifier list of charges for create virtual card
var VirtualCardIdentifier = []PricingIdentifier{CardCreation}
