package request

type ctxUserIDKey string
type SettlementID string

const (
	UserIDKey       ctxUserIDKey = "userId"
	SettlementIDKey SettlementID = "settlementId"
)
