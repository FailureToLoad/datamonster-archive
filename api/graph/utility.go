package graph

import (
	"entgo.io/ent/dialect/sql"
	"github.com/failuretoload/datamonster/ent"
	"github.com/failuretoload/datamonster/ent/survivor"
)

func survivorOrderFunc(order *ent.SurvivorOrder) func(opts ...sql.OrderTermOption) survivor.OrderOption {
	switch order.Field.String() {
	case "BORN":
		return survivor.ByBorn
	case "GENDER":
		return survivor.ByGender
	case "HUNTXP":
		return survivor.ByHuntxp
	case "SURVIVAL":
		return survivor.BySurvival
	case "MOVEMENT":
		return survivor.ByMovement
	case "ACCURACY":
		return survivor.ByAccuracy
	case "STRENGTH":
		return survivor.ByStrength
	case "EVASION":
		return survivor.ByEvasion
	case "LUCK":
		return survivor.ByLuck
	case "SPEED":
		return survivor.BySpeed
	case "SYSTEMICPRESSURE":
		return survivor.BySystemicpressure
	case "TORMENT":
		return survivor.ByTorment
	case "INSANITY":
		return survivor.ByInsanity
	case "LUMI":
		return survivor.ByLumi
	case "COURAGE":
		return survivor.ByCourage
	case "UNDERSTANDING":
		return survivor.ByUnderstanding
	case "SETTLEMENTID":
		return survivor.BySettlementID
	}
	return survivor.ByName
}
