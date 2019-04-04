// Code generated by SQLBoiler (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/volatiletech/null"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
	"github.com/volatiletech/sqlboiler/types"
)

// TFChargeCode is an object representing the database table.
type TFChargeCode struct {
	TFChargecodeID int           `boil:"tf_chargecode_id" json:"tf_chargecode_id" toml:"tf_chargecode_id" yaml:"tf_chargecode_id"`
	ChargeCode     string        `boil:"charge_code" json:"charge_code" toml:"charge_code" yaml:"charge_code"`
	ProdCode       null.String   `boil:"prod_code" json:"prod_code,omitempty" toml:"prod_code" yaml:"prod_code,omitempty"`
	Currency       null.String   `boil:"currency" json:"currency,omitempty" toml:"currency" yaml:"currency,omitempty"`
	Amount         types.Decimal `boil:"amount" json:"amount" toml:"amount" yaml:"amount"`
	BranchID       null.Int      `boil:"branch_id" json:"branch_id,omitempty" toml:"branch_id" yaml:"branch_id,omitempty"`
	BankID         int           `boil:"bank_id" json:"bank_id" toml:"bank_id" yaml:"bank_id"`
	PartyID        null.Int      `boil:"party_id" json:"party_id,omitempty" toml:"party_id" yaml:"party_id,omitempty"`
	MakerDate      time.Time     `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate    null.Time     `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID        string        `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID      null.String   `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy     null.String   `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate   null.Time     `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *tFChargeCodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tFChargeCodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TFChargeCodeColumns = struct {
	TFChargecodeID string
	ChargeCode     string
	ProdCode       string
	Currency       string
	Amount         string
	BranchID       string
	BankID         string
	PartyID        string
	MakerDate      string
	CheckerDate    string
	MakerID        string
	CheckerID      string
	ModifiedBy     string
	ModifiedDate   string
}{
	TFChargecodeID: "tf_chargecode_id",
	ChargeCode:     "charge_code",
	ProdCode:       "prod_code",
	Currency:       "currency",
	Amount:         "amount",
	BranchID:       "branch_id",
	BankID:         "bank_id",
	PartyID:        "party_id",
	MakerDate:      "maker_date",
	CheckerDate:    "checker_date",
	MakerID:        "maker_id",
	CheckerID:      "checker_id",
	ModifiedBy:     "modified_by",
	ModifiedDate:   "modified_date",
}

// Generated where

var TFChargeCodeWhere = struct {
	TFChargecodeID whereHelperint
	ChargeCode     whereHelperstring
	ProdCode       whereHelpernull_String
	Currency       whereHelpernull_String
	Amount         whereHelpertypes_Decimal
	BranchID       whereHelpernull_Int
	BankID         whereHelperint
	PartyID        whereHelpernull_Int
	MakerDate      whereHelpertime_Time
	CheckerDate    whereHelpernull_Time
	MakerID        whereHelperstring
	CheckerID      whereHelpernull_String
	ModifiedBy     whereHelpernull_String
	ModifiedDate   whereHelpernull_Time
}{
	TFChargecodeID: whereHelperint{field: `tf_chargecode_id`},
	ChargeCode:     whereHelperstring{field: `charge_code`},
	ProdCode:       whereHelpernull_String{field: `prod_code`},
	Currency:       whereHelpernull_String{field: `currency`},
	Amount:         whereHelpertypes_Decimal{field: `amount`},
	BranchID:       whereHelpernull_Int{field: `branch_id`},
	BankID:         whereHelperint{field: `bank_id`},
	PartyID:        whereHelpernull_Int{field: `party_id`},
	MakerDate:      whereHelpertime_Time{field: `maker_date`},
	CheckerDate:    whereHelpernull_Time{field: `checker_date`},
	MakerID:        whereHelperstring{field: `maker_id`},
	CheckerID:      whereHelpernull_String{field: `checker_id`},
	ModifiedBy:     whereHelpernull_String{field: `modified_by`},
	ModifiedDate:   whereHelpernull_Time{field: `modified_date`},
}

// TFChargeCodeRels is where relationship names are stored.
var TFChargeCodeRels = struct {
	ChargeCodeTFCharges     string
	ChargeCodeTFChargesTxns string
}{
	ChargeCodeTFCharges:     "ChargeCodeTFCharges",
	ChargeCodeTFChargesTxns: "ChargeCodeTFChargesTxns",
}

// tFChargeCodeR is where relationships are stored.
type tFChargeCodeR struct {
	ChargeCodeTFCharges     TFChargeSlice
	ChargeCodeTFChargesTxns TFChargesTxnSlice
}

// NewStruct creates a new relationship struct
func (*tFChargeCodeR) NewStruct() *tFChargeCodeR {
	return &tFChargeCodeR{}
}

// tFChargeCodeL is where Load methods for each relationship are stored.
type tFChargeCodeL struct{}

var (
	tFChargeCodeColumns               = []string{"tf_chargecode_id", "charge_code", "prod_code", "currency", "amount", "branch_id", "bank_id", "party_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	tFChargeCodeColumnsWithoutDefault = []string{"tf_chargecode_id", "charge_code", "prod_code", "currency", "amount", "branch_id", "bank_id", "party_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	tFChargeCodeColumnsWithDefault    = []string{}
	tFChargeCodePrimaryKeyColumns     = []string{"tf_chargecode_id"}
)

type (
	// TFChargeCodeSlice is an alias for a slice of pointers to TFChargeCode.
	// This should generally be used opposed to []TFChargeCode.
	TFChargeCodeSlice []*TFChargeCode
	// TFChargeCodeHook is the signature for custom TFChargeCode hook methods
	TFChargeCodeHook func(context.Context, boil.ContextExecutor, *TFChargeCode) error

	tFChargeCodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tFChargeCodeType                 = reflect.TypeOf(&TFChargeCode{})
	tFChargeCodeMapping              = queries.MakeStructMapping(tFChargeCodeType)
	tFChargeCodePrimaryKeyMapping, _ = queries.BindMapping(tFChargeCodeType, tFChargeCodeMapping, tFChargeCodePrimaryKeyColumns)
	tFChargeCodeInsertCacheMut       sync.RWMutex
	tFChargeCodeInsertCache          = make(map[string]insertCache)
	tFChargeCodeUpdateCacheMut       sync.RWMutex
	tFChargeCodeUpdateCache          = make(map[string]updateCache)
	tFChargeCodeUpsertCacheMut       sync.RWMutex
	tFChargeCodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tFChargeCodeBeforeInsertHooks []TFChargeCodeHook
var tFChargeCodeBeforeUpdateHooks []TFChargeCodeHook
var tFChargeCodeBeforeDeleteHooks []TFChargeCodeHook
var tFChargeCodeBeforeUpsertHooks []TFChargeCodeHook

var tFChargeCodeAfterInsertHooks []TFChargeCodeHook
var tFChargeCodeAfterSelectHooks []TFChargeCodeHook
var tFChargeCodeAfterUpdateHooks []TFChargeCodeHook
var tFChargeCodeAfterDeleteHooks []TFChargeCodeHook
var tFChargeCodeAfterUpsertHooks []TFChargeCodeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TFChargeCode) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFChargeCodeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TFChargeCode) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFChargeCodeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TFChargeCode) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFChargeCodeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TFChargeCode) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFChargeCodeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TFChargeCode) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFChargeCodeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TFChargeCode) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFChargeCodeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TFChargeCode) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFChargeCodeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TFChargeCode) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFChargeCodeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TFChargeCode) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFChargeCodeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTFChargeCodeHook registers your hook function for all future operations.
func AddTFChargeCodeHook(hookPoint boil.HookPoint, tFChargeCodeHook TFChargeCodeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		tFChargeCodeBeforeInsertHooks = append(tFChargeCodeBeforeInsertHooks, tFChargeCodeHook)
	case boil.BeforeUpdateHook:
		tFChargeCodeBeforeUpdateHooks = append(tFChargeCodeBeforeUpdateHooks, tFChargeCodeHook)
	case boil.BeforeDeleteHook:
		tFChargeCodeBeforeDeleteHooks = append(tFChargeCodeBeforeDeleteHooks, tFChargeCodeHook)
	case boil.BeforeUpsertHook:
		tFChargeCodeBeforeUpsertHooks = append(tFChargeCodeBeforeUpsertHooks, tFChargeCodeHook)
	case boil.AfterInsertHook:
		tFChargeCodeAfterInsertHooks = append(tFChargeCodeAfterInsertHooks, tFChargeCodeHook)
	case boil.AfterSelectHook:
		tFChargeCodeAfterSelectHooks = append(tFChargeCodeAfterSelectHooks, tFChargeCodeHook)
	case boil.AfterUpdateHook:
		tFChargeCodeAfterUpdateHooks = append(tFChargeCodeAfterUpdateHooks, tFChargeCodeHook)
	case boil.AfterDeleteHook:
		tFChargeCodeAfterDeleteHooks = append(tFChargeCodeAfterDeleteHooks, tFChargeCodeHook)
	case boil.AfterUpsertHook:
		tFChargeCodeAfterUpsertHooks = append(tFChargeCodeAfterUpsertHooks, tFChargeCodeHook)
	}
}

// One returns a single tFChargeCode record from the query.
func (q tFChargeCodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TFChargeCode, error) {
	o := &TFChargeCode{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for TFChargeCodes")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TFChargeCode records from the query.
func (q tFChargeCodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (TFChargeCodeSlice, error) {
	var o []*TFChargeCode

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TFChargeCode slice")
	}

	if len(tFChargeCodeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TFChargeCode records in the query.
func (q tFChargeCodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count TFChargeCodes rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tFChargeCodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if TFChargeCodes exists")
	}

	return count > 0, nil
}

// ChargeCodeTFCharges retrieves all the TFCharge's TFCharges with an executor via charge_code column.
func (o *TFChargeCode) ChargeCodeTFCharges(mods ...qm.QueryMod) tFChargeQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`TFCharges`.`charge_code`=?", o.ChargeCode),
	)

	query := TFCharges(queryMods...)
	queries.SetFrom(query.Query, "`TFCharges`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`TFCharges`.*"})
	}

	return query
}

// ChargeCodeTFChargesTxns retrieves all the TFChargesTxn's TFChargesTxns with an executor via charge_code column.
func (o *TFChargeCode) ChargeCodeTFChargesTxns(mods ...qm.QueryMod) tFChargesTxnQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`TFChargesTxn`.`charge_code`=?", o.ChargeCode),
	)

	query := TFChargesTxns(queryMods...)
	queries.SetFrom(query.Query, "`TFChargesTxn`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`TFChargesTxn`.*"})
	}

	return query
}

// LoadChargeCodeTFCharges allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (tFChargeCodeL) LoadChargeCodeTFCharges(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTFChargeCode interface{}, mods queries.Applicator) error {
	var slice []*TFChargeCode
	var object *TFChargeCode

	if singular {
		object = maybeTFChargeCode.(*TFChargeCode)
	} else {
		slice = *maybeTFChargeCode.(*[]*TFChargeCode)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tFChargeCodeR{}
		}
		args = append(args, object.ChargeCode)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tFChargeCodeR{}
			}

			for _, a := range args {
				if a == obj.ChargeCode {
					continue Outer
				}
			}

			args = append(args, obj.ChargeCode)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`TFCharges`), qm.WhereIn(`charge_code in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load TFCharges")
	}

	var resultSlice []*TFCharge
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice TFCharges")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on TFCharges")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for TFCharges")
	}

	if len(tFChargeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ChargeCodeTFCharges = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &tFChargeR{}
			}
			foreign.R.ChargeCode = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ChargeCode == foreign.ChargeCode {
				local.R.ChargeCodeTFCharges = append(local.R.ChargeCodeTFCharges, foreign)
				if foreign.R == nil {
					foreign.R = &tFChargeR{}
				}
				foreign.R.ChargeCode = local
				break
			}
		}
	}

	return nil
}

// LoadChargeCodeTFChargesTxns allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (tFChargeCodeL) LoadChargeCodeTFChargesTxns(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTFChargeCode interface{}, mods queries.Applicator) error {
	var slice []*TFChargeCode
	var object *TFChargeCode

	if singular {
		object = maybeTFChargeCode.(*TFChargeCode)
	} else {
		slice = *maybeTFChargeCode.(*[]*TFChargeCode)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tFChargeCodeR{}
		}
		args = append(args, object.ChargeCode)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tFChargeCodeR{}
			}

			for _, a := range args {
				if a == obj.ChargeCode {
					continue Outer
				}
			}

			args = append(args, obj.ChargeCode)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`TFChargesTxn`), qm.WhereIn(`charge_code in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load TFChargesTxn")
	}

	var resultSlice []*TFChargesTxn
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice TFChargesTxn")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on TFChargesTxn")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for TFChargesTxn")
	}

	if len(tFChargesTxnAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ChargeCodeTFChargesTxns = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &tFChargesTxnR{}
			}
			foreign.R.ChargeCode = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ChargeCode == foreign.ChargeCode {
				local.R.ChargeCodeTFChargesTxns = append(local.R.ChargeCodeTFChargesTxns, foreign)
				if foreign.R == nil {
					foreign.R = &tFChargesTxnR{}
				}
				foreign.R.ChargeCode = local
				break
			}
		}
	}

	return nil
}

// AddChargeCodeTFCharges adds the given related objects to the existing relationships
// of the TFChargeCode, optionally inserting them as new records.
// Appends related to o.R.ChargeCodeTFCharges.
// Sets related.R.ChargeCode appropriately.
func (o *TFChargeCode) AddChargeCodeTFCharges(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*TFCharge) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChargeCode = o.ChargeCode
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `TFCharges` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"charge_code"}),
				strmangle.WhereClause("`", "`", 0, tFChargePrimaryKeyColumns),
			)
			values := []interface{}{o.ChargeCode, rel.TFChargesID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ChargeCode = o.ChargeCode
		}
	}

	if o.R == nil {
		o.R = &tFChargeCodeR{
			ChargeCodeTFCharges: related,
		}
	} else {
		o.R.ChargeCodeTFCharges = append(o.R.ChargeCodeTFCharges, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &tFChargeR{
				ChargeCode: o,
			}
		} else {
			rel.R.ChargeCode = o
		}
	}
	return nil
}

// AddChargeCodeTFChargesTxns adds the given related objects to the existing relationships
// of the TFChargeCode, optionally inserting them as new records.
// Appends related to o.R.ChargeCodeTFChargesTxns.
// Sets related.R.ChargeCode appropriately.
func (o *TFChargeCode) AddChargeCodeTFChargesTxns(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*TFChargesTxn) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ChargeCode = o.ChargeCode
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `TFChargesTxn` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"charge_code"}),
				strmangle.WhereClause("`", "`", 0, tFChargesTxnPrimaryKeyColumns),
			)
			values := []interface{}{o.ChargeCode, rel.TFChargesID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ChargeCode = o.ChargeCode
		}
	}

	if o.R == nil {
		o.R = &tFChargeCodeR{
			ChargeCodeTFChargesTxns: related,
		}
	} else {
		o.R.ChargeCodeTFChargesTxns = append(o.R.ChargeCodeTFChargesTxns, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &tFChargesTxnR{
				ChargeCode: o,
			}
		} else {
			rel.R.ChargeCode = o
		}
	}
	return nil
}

// TFChargeCodes retrieves all the records using an executor.
func TFChargeCodes(mods ...qm.QueryMod) tFChargeCodeQuery {
	mods = append(mods, qm.From("`TFChargeCodes`"))
	return tFChargeCodeQuery{NewQuery(mods...)}
}

// FindTFChargeCode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTFChargeCode(ctx context.Context, exec boil.ContextExecutor, tFChargecodeID int, selectCols ...string) (*TFChargeCode, error) {
	tFChargeCodeObj := &TFChargeCode{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `TFChargeCodes` where `tf_chargecode_id`=?", sel,
	)

	q := queries.Raw(query, tFChargecodeID)

	err := q.Bind(ctx, exec, tFChargeCodeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from TFChargeCodes")
	}

	return tFChargeCodeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TFChargeCode) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TFChargeCodes provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tFChargeCodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tFChargeCodeInsertCacheMut.RLock()
	cache, cached := tFChargeCodeInsertCache[key]
	tFChargeCodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tFChargeCodeColumns,
			tFChargeCodeColumnsWithDefault,
			tFChargeCodeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tFChargeCodeType, tFChargeCodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tFChargeCodeType, tFChargeCodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `TFChargeCodes` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `TFChargeCodes` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `TFChargeCodes` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tFChargeCodePrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into TFChargeCodes")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.TFChargecodeID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TFChargeCodes")
	}

CacheNoHooks:
	if !cached {
		tFChargeCodeInsertCacheMut.Lock()
		tFChargeCodeInsertCache[key] = cache
		tFChargeCodeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TFChargeCode.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TFChargeCode) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tFChargeCodeUpdateCacheMut.RLock()
	cache, cached := tFChargeCodeUpdateCache[key]
	tFChargeCodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tFChargeCodeColumns,
			tFChargeCodePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update TFChargeCodes, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `TFChargeCodes` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tFChargeCodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tFChargeCodeType, tFChargeCodeMapping, append(wl, tFChargeCodePrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update TFChargeCodes row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for TFChargeCodes")
	}

	if !cached {
		tFChargeCodeUpdateCacheMut.Lock()
		tFChargeCodeUpdateCache[key] = cache
		tFChargeCodeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tFChargeCodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for TFChargeCodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for TFChargeCodes")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TFChargeCodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tFChargeCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `TFChargeCodes` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tFChargeCodePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tFChargeCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tFChargeCode")
	}
	return rowsAff, nil
}

var mySQLTFChargeCodeUniqueColumns = []string{
	"tf_chargecode_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TFChargeCode) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TFChargeCodes provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tFChargeCodeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTFChargeCodeUniqueColumns, o)

	if len(nzUniques) == 0 {
		return errors.New("cannot upsert with a table that cannot conflict on a unique column")
	}

	// Build cache key in-line uglily - mysql vs psql problems
	buf := strmangle.GetBuffer()
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzUniques {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	tFChargeCodeUpsertCacheMut.RLock()
	cache, cached := tFChargeCodeUpsertCache[key]
	tFChargeCodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tFChargeCodeColumns,
			tFChargeCodeColumnsWithDefault,
			tFChargeCodeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			tFChargeCodeColumns,
			tFChargeCodePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert TFChargeCodes, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "TFChargeCodes", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `TFChargeCodes` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tFChargeCodeType, tFChargeCodeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tFChargeCodeType, tFChargeCodeMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	_, err = exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for TFChargeCodes")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tFChargeCodeType, tFChargeCodeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for TFChargeCodes")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TFChargeCodes")
	}

CacheNoHooks:
	if !cached {
		tFChargeCodeUpsertCacheMut.Lock()
		tFChargeCodeUpsertCache[key] = cache
		tFChargeCodeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TFChargeCode record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TFChargeCode) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TFChargeCode provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tFChargeCodePrimaryKeyMapping)
	sql := "DELETE FROM `TFChargeCodes` WHERE `tf_chargecode_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from TFChargeCodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for TFChargeCodes")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tFChargeCodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tFChargeCodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from TFChargeCodes")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TFChargeCodes")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TFChargeCodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TFChargeCode slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(tFChargeCodeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tFChargeCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `TFChargeCodes` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tFChargeCodePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tFChargeCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TFChargeCodes")
	}

	if len(tFChargeCodeAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *TFChargeCode) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTFChargeCode(ctx, exec, o.TFChargecodeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TFChargeCodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TFChargeCodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tFChargeCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `TFChargeCodes`.* FROM `TFChargeCodes` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tFChargeCodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TFChargeCodeSlice")
	}

	*o = slice

	return nil
}

// TFChargeCodeExists checks if the TFChargeCode row exists.
func TFChargeCodeExists(ctx context.Context, exec boil.ContextExecutor, tFChargecodeID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `TFChargeCodes` where `tf_chargecode_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, tFChargecodeID)
	}

	row := exec.QueryRowContext(ctx, sql, tFChargecodeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if TFChargeCodes exists")
	}

	return exists, nil
}