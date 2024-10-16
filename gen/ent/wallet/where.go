// Code generated by ent, DO NOT EDIT.

package wallet

import (
	"entgo.io/ent/dialect/sql"
	"github.com/tyagnii/wallets/gen/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldID, id))
}

// UUID applies equality check predicate on the "UUID" field. It's identical to UUIDEQ.
func UUID(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldUUID, v))
}

// Balance applies equality check predicate on the "balance" field. It's identical to BalanceEQ.
func Balance(v int) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldBalance, v))
}

// UUIDEQ applies the EQ predicate on the "UUID" field.
func UUIDEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "UUID" field.
func UUIDNEQ(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "UUID" field.
func UUIDIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "UUID" field.
func UUIDNotIn(vs ...string) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "UUID" field.
func UUIDGT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "UUID" field.
func UUIDGTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "UUID" field.
func UUIDLT(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "UUID" field.
func UUIDLTE(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldUUID, v))
}

// UUIDContains applies the Contains predicate on the "UUID" field.
func UUIDContains(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContains(FieldUUID, v))
}

// UUIDHasPrefix applies the HasPrefix predicate on the "UUID" field.
func UUIDHasPrefix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasPrefix(FieldUUID, v))
}

// UUIDHasSuffix applies the HasSuffix predicate on the "UUID" field.
func UUIDHasSuffix(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldHasSuffix(FieldUUID, v))
}

// UUIDEqualFold applies the EqualFold predicate on the "UUID" field.
func UUIDEqualFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldEqualFold(FieldUUID, v))
}

// UUIDContainsFold applies the ContainsFold predicate on the "UUID" field.
func UUIDContainsFold(v string) predicate.Wallet {
	return predicate.Wallet(sql.FieldContainsFold(FieldUUID, v))
}

// BalanceEQ applies the EQ predicate on the "balance" field.
func BalanceEQ(v int) predicate.Wallet {
	return predicate.Wallet(sql.FieldEQ(FieldBalance, v))
}

// BalanceNEQ applies the NEQ predicate on the "balance" field.
func BalanceNEQ(v int) predicate.Wallet {
	return predicate.Wallet(sql.FieldNEQ(FieldBalance, v))
}

// BalanceIn applies the In predicate on the "balance" field.
func BalanceIn(vs ...int) predicate.Wallet {
	return predicate.Wallet(sql.FieldIn(FieldBalance, vs...))
}

// BalanceNotIn applies the NotIn predicate on the "balance" field.
func BalanceNotIn(vs ...int) predicate.Wallet {
	return predicate.Wallet(sql.FieldNotIn(FieldBalance, vs...))
}

// BalanceGT applies the GT predicate on the "balance" field.
func BalanceGT(v int) predicate.Wallet {
	return predicate.Wallet(sql.FieldGT(FieldBalance, v))
}

// BalanceGTE applies the GTE predicate on the "balance" field.
func BalanceGTE(v int) predicate.Wallet {
	return predicate.Wallet(sql.FieldGTE(FieldBalance, v))
}

// BalanceLT applies the LT predicate on the "balance" field.
func BalanceLT(v int) predicate.Wallet {
	return predicate.Wallet(sql.FieldLT(FieldBalance, v))
}

// BalanceLTE applies the LTE predicate on the "balance" field.
func BalanceLTE(v int) predicate.Wallet {
	return predicate.Wallet(sql.FieldLTE(FieldBalance, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Wallet) predicate.Wallet {
	return predicate.Wallet(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Wallet) predicate.Wallet {
	return predicate.Wallet(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Wallet) predicate.Wallet {
	return predicate.Wallet(sql.NotPredicates(p))
}
