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
)

// PaymentMethod is an object representing the database table.
type PaymentMethod struct {
	PaymentMethodCode       int         `boil:"payment_method_code" json:"payment_method_code" toml:"payment_method_code" yaml:"payment_method_code"`
	Active                  null.String `boil:"active" json:"active,omitempty" toml:"active" yaml:"active,omitempty"`
	Description             null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	PaymentMethodName       string      `boil:"payment_method_name" json:"payment_method_name" toml:"payment_method_name" yaml:"payment_method_name"`
	Fee                     null.Int    `boil:"fee" json:"fee,omitempty" toml:"fee" yaml:"fee,omitempty"`
	TimeToResolve           null.Int    `boil:"time_to_resolve" json:"time_to_resolve,omitempty" toml:"time_to_resolve" yaml:"time_to_resolve,omitempty"`
	DefaultPaymentMethodFLG null.String `boil:"default_payment_method_flg" json:"default_payment_method_flg,omitempty" toml:"default_payment_method_flg" yaml:"default_payment_method_flg,omitempty"`
	PaymentActiveStartDate  null.Time   `boil:"payment_active_start_date" json:"payment_active_start_date,omitempty" toml:"payment_active_start_date" yaml:"payment_active_start_date,omitempty"`
	PaymentActiveEndDate    null.Time   `boil:"payment_active_end_date" json:"payment_active_end_date,omitempty" toml:"payment_active_end_date" yaml:"payment_active_end_date,omitempty"`
	MakerDate               time.Time   `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate             null.Time   `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID                 string      `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID               null.String `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy              null.String `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate            null.Time   `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *paymentMethodR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L paymentMethodL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PaymentMethodColumns = struct {
	PaymentMethodCode       string
	Active                  string
	Description             string
	PaymentMethodName       string
	Fee                     string
	TimeToResolve           string
	DefaultPaymentMethodFLG string
	PaymentActiveStartDate  string
	PaymentActiveEndDate    string
	MakerDate               string
	CheckerDate             string
	MakerID                 string
	CheckerID               string
	ModifiedBy              string
	ModifiedDate            string
}{
	PaymentMethodCode:       "payment_method_code",
	Active:                  "active",
	Description:             "description",
	PaymentMethodName:       "payment_method_name",
	Fee:                     "fee",
	TimeToResolve:           "time_to_resolve",
	DefaultPaymentMethodFLG: "default_payment_method_flg",
	PaymentActiveStartDate:  "payment_active_start_date",
	PaymentActiveEndDate:    "payment_active_end_date",
	MakerDate:               "maker_date",
	CheckerDate:             "checker_date",
	MakerID:                 "maker_id",
	CheckerID:               "checker_id",
	ModifiedBy:              "modified_by",
	ModifiedDate:            "modified_date",
}

// Generated where

var PaymentMethodWhere = struct {
	PaymentMethodCode       whereHelperint
	Active                  whereHelpernull_String
	Description             whereHelpernull_String
	PaymentMethodName       whereHelperstring
	Fee                     whereHelpernull_Int
	TimeToResolve           whereHelpernull_Int
	DefaultPaymentMethodFLG whereHelpernull_String
	PaymentActiveStartDate  whereHelpernull_Time
	PaymentActiveEndDate    whereHelpernull_Time
	MakerDate               whereHelpertime_Time
	CheckerDate             whereHelpernull_Time
	MakerID                 whereHelperstring
	CheckerID               whereHelpernull_String
	ModifiedBy              whereHelpernull_String
	ModifiedDate            whereHelpernull_Time
}{
	PaymentMethodCode:       whereHelperint{field: `payment_method_code`},
	Active:                  whereHelpernull_String{field: `active`},
	Description:             whereHelpernull_String{field: `description`},
	PaymentMethodName:       whereHelperstring{field: `payment_method_name`},
	Fee:                     whereHelpernull_Int{field: `fee`},
	TimeToResolve:           whereHelpernull_Int{field: `time_to_resolve`},
	DefaultPaymentMethodFLG: whereHelpernull_String{field: `default_payment_method_flg`},
	PaymentActiveStartDate:  whereHelpernull_Time{field: `payment_active_start_date`},
	PaymentActiveEndDate:    whereHelpernull_Time{field: `payment_active_end_date`},
	MakerDate:               whereHelpertime_Time{field: `maker_date`},
	CheckerDate:             whereHelpernull_Time{field: `checker_date`},
	MakerID:                 whereHelperstring{field: `maker_id`},
	CheckerID:               whereHelpernull_String{field: `checker_id`},
	ModifiedBy:              whereHelpernull_String{field: `modified_by`},
	ModifiedDate:            whereHelpernull_Time{field: `modified_date`},
}

// PaymentMethodRels is where relationship names are stored.
var PaymentMethodRels = struct {
	PaymentMethodCodePayments string
}{
	PaymentMethodCodePayments: "PaymentMethodCodePayments",
}

// paymentMethodR is where relationships are stored.
type paymentMethodR struct {
	PaymentMethodCodePayments PaymentSlice
}

// NewStruct creates a new relationship struct
func (*paymentMethodR) NewStruct() *paymentMethodR {
	return &paymentMethodR{}
}

// paymentMethodL is where Load methods for each relationship are stored.
type paymentMethodL struct{}

var (
	paymentMethodColumns               = []string{"payment_method_code", "active", "description", "payment_method_name", "fee", "time_to_resolve", "default_payment_method_flg", "payment_active_start_date", "payment_active_end_date", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	paymentMethodColumnsWithoutDefault = []string{"active", "description", "payment_method_name", "fee", "time_to_resolve", "default_payment_method_flg", "payment_active_start_date", "payment_active_end_date", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	paymentMethodColumnsWithDefault    = []string{"payment_method_code"}
	paymentMethodPrimaryKeyColumns     = []string{"payment_method_code"}
)

type (
	// PaymentMethodSlice is an alias for a slice of pointers to PaymentMethod.
	// This should generally be used opposed to []PaymentMethod.
	PaymentMethodSlice []*PaymentMethod
	// PaymentMethodHook is the signature for custom PaymentMethod hook methods
	PaymentMethodHook func(context.Context, boil.ContextExecutor, *PaymentMethod) error

	paymentMethodQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	paymentMethodType                 = reflect.TypeOf(&PaymentMethod{})
	paymentMethodMapping              = queries.MakeStructMapping(paymentMethodType)
	paymentMethodPrimaryKeyMapping, _ = queries.BindMapping(paymentMethodType, paymentMethodMapping, paymentMethodPrimaryKeyColumns)
	paymentMethodInsertCacheMut       sync.RWMutex
	paymentMethodInsertCache          = make(map[string]insertCache)
	paymentMethodUpdateCacheMut       sync.RWMutex
	paymentMethodUpdateCache          = make(map[string]updateCache)
	paymentMethodUpsertCacheMut       sync.RWMutex
	paymentMethodUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var paymentMethodBeforeInsertHooks []PaymentMethodHook
var paymentMethodBeforeUpdateHooks []PaymentMethodHook
var paymentMethodBeforeDeleteHooks []PaymentMethodHook
var paymentMethodBeforeUpsertHooks []PaymentMethodHook

var paymentMethodAfterInsertHooks []PaymentMethodHook
var paymentMethodAfterSelectHooks []PaymentMethodHook
var paymentMethodAfterUpdateHooks []PaymentMethodHook
var paymentMethodAfterDeleteHooks []PaymentMethodHook
var paymentMethodAfterUpsertHooks []PaymentMethodHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PaymentMethod) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentMethodBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PaymentMethod) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentMethodBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PaymentMethod) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentMethodBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PaymentMethod) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentMethodBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PaymentMethod) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentMethodAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PaymentMethod) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentMethodAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PaymentMethod) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentMethodAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PaymentMethod) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentMethodAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PaymentMethod) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentMethodAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPaymentMethodHook registers your hook function for all future operations.
func AddPaymentMethodHook(hookPoint boil.HookPoint, paymentMethodHook PaymentMethodHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		paymentMethodBeforeInsertHooks = append(paymentMethodBeforeInsertHooks, paymentMethodHook)
	case boil.BeforeUpdateHook:
		paymentMethodBeforeUpdateHooks = append(paymentMethodBeforeUpdateHooks, paymentMethodHook)
	case boil.BeforeDeleteHook:
		paymentMethodBeforeDeleteHooks = append(paymentMethodBeforeDeleteHooks, paymentMethodHook)
	case boil.BeforeUpsertHook:
		paymentMethodBeforeUpsertHooks = append(paymentMethodBeforeUpsertHooks, paymentMethodHook)
	case boil.AfterInsertHook:
		paymentMethodAfterInsertHooks = append(paymentMethodAfterInsertHooks, paymentMethodHook)
	case boil.AfterSelectHook:
		paymentMethodAfterSelectHooks = append(paymentMethodAfterSelectHooks, paymentMethodHook)
	case boil.AfterUpdateHook:
		paymentMethodAfterUpdateHooks = append(paymentMethodAfterUpdateHooks, paymentMethodHook)
	case boil.AfterDeleteHook:
		paymentMethodAfterDeleteHooks = append(paymentMethodAfterDeleteHooks, paymentMethodHook)
	case boil.AfterUpsertHook:
		paymentMethodAfterUpsertHooks = append(paymentMethodAfterUpsertHooks, paymentMethodHook)
	}
}

// One returns a single paymentMethod record from the query.
func (q paymentMethodQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PaymentMethod, error) {
	o := &PaymentMethod{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for PaymentMethod")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PaymentMethod records from the query.
func (q paymentMethodQuery) All(ctx context.Context, exec boil.ContextExecutor) (PaymentMethodSlice, error) {
	var o []*PaymentMethod

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PaymentMethod slice")
	}

	if len(paymentMethodAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PaymentMethod records in the query.
func (q paymentMethodQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count PaymentMethod rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q paymentMethodQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if PaymentMethod exists")
	}

	return count > 0, nil
}

// PaymentMethodCodePayments retrieves all the Payment's Payments with an executor via payment_method_code column.
func (o *PaymentMethod) PaymentMethodCodePayments(mods ...qm.QueryMod) paymentQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`Payment`.`payment_method_code`=?", o.PaymentMethodCode),
	)

	query := Payments(queryMods...)
	queries.SetFrom(query.Query, "`Payment`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`Payment`.*"})
	}

	return query
}

// LoadPaymentMethodCodePayments allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (paymentMethodL) LoadPaymentMethodCodePayments(ctx context.Context, e boil.ContextExecutor, singular bool, maybePaymentMethod interface{}, mods queries.Applicator) error {
	var slice []*PaymentMethod
	var object *PaymentMethod

	if singular {
		object = maybePaymentMethod.(*PaymentMethod)
	} else {
		slice = *maybePaymentMethod.(*[]*PaymentMethod)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &paymentMethodR{}
		}
		args = append(args, object.PaymentMethodCode)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &paymentMethodR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.PaymentMethodCode) {
					continue Outer
				}
			}

			args = append(args, obj.PaymentMethodCode)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`Payment`), qm.WhereIn(`payment_method_code in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Payment")
	}

	var resultSlice []*Payment
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Payment")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on Payment")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Payment")
	}

	if len(paymentAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.PaymentMethodCodePayments = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &paymentR{}
			}
			foreign.R.PaymentMethodCode = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if queries.Equal(local.PaymentMethodCode, foreign.PaymentMethodCode) {
				local.R.PaymentMethodCodePayments = append(local.R.PaymentMethodCodePayments, foreign)
				if foreign.R == nil {
					foreign.R = &paymentR{}
				}
				foreign.R.PaymentMethodCode = local
				break
			}
		}
	}

	return nil
}

// AddPaymentMethodCodePayments adds the given related objects to the existing relationships
// of the PaymentMethod, optionally inserting them as new records.
// Appends related to o.R.PaymentMethodCodePayments.
// Sets related.R.PaymentMethodCode appropriately.
func (o *PaymentMethod) AddPaymentMethodCodePayments(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Payment) error {
	var err error
	for _, rel := range related {
		if insert {
			queries.Assign(&rel.PaymentMethodCode, o.PaymentMethodCode)
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `Payment` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"payment_method_code"}),
				strmangle.WhereClause("`", "`", 0, paymentPrimaryKeyColumns),
			)
			values := []interface{}{o.PaymentMethodCode, rel.PaymentID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			queries.Assign(&rel.PaymentMethodCode, o.PaymentMethodCode)
		}
	}

	if o.R == nil {
		o.R = &paymentMethodR{
			PaymentMethodCodePayments: related,
		}
	} else {
		o.R.PaymentMethodCodePayments = append(o.R.PaymentMethodCodePayments, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &paymentR{
				PaymentMethodCode: o,
			}
		} else {
			rel.R.PaymentMethodCode = o
		}
	}
	return nil
}

// SetPaymentMethodCodePayments removes all previously related items of the
// PaymentMethod replacing them completely with the passed
// in related items, optionally inserting them as new records.
// Sets o.R.PaymentMethodCode's PaymentMethodCodePayments accordingly.
// Replaces o.R.PaymentMethodCodePayments with related.
// Sets related.R.PaymentMethodCode's PaymentMethodCodePayments accordingly.
func (o *PaymentMethod) SetPaymentMethodCodePayments(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Payment) error {
	query := "update `Payment` set `payment_method_code` = null where `payment_method_code` = ?"
	values := []interface{}{o.PaymentMethodCode}
	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err := exec.ExecContext(ctx, query, values...)
	if err != nil {
		return errors.Wrap(err, "failed to remove relationships before set")
	}

	if o.R != nil {
		for _, rel := range o.R.PaymentMethodCodePayments {
			queries.SetScanner(&rel.PaymentMethodCode, nil)
			if rel.R == nil {
				continue
			}

			rel.R.PaymentMethodCode = nil
		}

		o.R.PaymentMethodCodePayments = nil
	}
	return o.AddPaymentMethodCodePayments(ctx, exec, insert, related...)
}

// RemovePaymentMethodCodePayments relationships from objects passed in.
// Removes related items from R.PaymentMethodCodePayments (uses pointer comparison, removal does not keep order)
// Sets related.R.PaymentMethodCode.
func (o *PaymentMethod) RemovePaymentMethodCodePayments(ctx context.Context, exec boil.ContextExecutor, related ...*Payment) error {
	var err error
	for _, rel := range related {
		queries.SetScanner(&rel.PaymentMethodCode, nil)
		if rel.R != nil {
			rel.R.PaymentMethodCode = nil
		}
		if _, err = rel.Update(ctx, exec, boil.Whitelist("payment_method_code")); err != nil {
			return err
		}
	}
	if o.R == nil {
		return nil
	}

	for _, rel := range related {
		for i, ri := range o.R.PaymentMethodCodePayments {
			if rel != ri {
				continue
			}

			ln := len(o.R.PaymentMethodCodePayments)
			if ln > 1 && i < ln-1 {
				o.R.PaymentMethodCodePayments[i] = o.R.PaymentMethodCodePayments[ln-1]
			}
			o.R.PaymentMethodCodePayments = o.R.PaymentMethodCodePayments[:ln-1]
			break
		}
	}

	return nil
}

// PaymentMethods retrieves all the records using an executor.
func PaymentMethods(mods ...qm.QueryMod) paymentMethodQuery {
	mods = append(mods, qm.From("`PaymentMethod`"))
	return paymentMethodQuery{NewQuery(mods...)}
}

// FindPaymentMethod retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPaymentMethod(ctx context.Context, exec boil.ContextExecutor, paymentMethodCode int, selectCols ...string) (*PaymentMethod, error) {
	paymentMethodObj := &PaymentMethod{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `PaymentMethod` where `payment_method_code`=?", sel,
	)

	q := queries.Raw(query, paymentMethodCode)

	err := q.Bind(ctx, exec, paymentMethodObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from PaymentMethod")
	}

	return paymentMethodObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PaymentMethod) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no PaymentMethod provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(paymentMethodColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	paymentMethodInsertCacheMut.RLock()
	cache, cached := paymentMethodInsertCache[key]
	paymentMethodInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			paymentMethodColumns,
			paymentMethodColumnsWithDefault,
			paymentMethodColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(paymentMethodType, paymentMethodMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(paymentMethodType, paymentMethodMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `PaymentMethod` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `PaymentMethod` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `PaymentMethod` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, paymentMethodPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into PaymentMethod")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.PaymentMethodCode = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == paymentMethodMapping["PaymentMethodCode"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PaymentMethodCode,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for PaymentMethod")
	}

CacheNoHooks:
	if !cached {
		paymentMethodInsertCacheMut.Lock()
		paymentMethodInsertCache[key] = cache
		paymentMethodInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PaymentMethod.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PaymentMethod) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	paymentMethodUpdateCacheMut.RLock()
	cache, cached := paymentMethodUpdateCache[key]
	paymentMethodUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			paymentMethodColumns,
			paymentMethodPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update PaymentMethod, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `PaymentMethod` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, paymentMethodPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(paymentMethodType, paymentMethodMapping, append(wl, paymentMethodPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update PaymentMethod row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for PaymentMethod")
	}

	if !cached {
		paymentMethodUpdateCacheMut.Lock()
		paymentMethodUpdateCache[key] = cache
		paymentMethodUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q paymentMethodQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for PaymentMethod")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for PaymentMethod")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PaymentMethodSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), paymentMethodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `PaymentMethod` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, paymentMethodPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in paymentMethod slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all paymentMethod")
	}
	return rowsAff, nil
}

var mySQLPaymentMethodUniqueColumns = []string{
	"payment_method_code",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PaymentMethod) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no PaymentMethod provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(paymentMethodColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPaymentMethodUniqueColumns, o)

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

	paymentMethodUpsertCacheMut.RLock()
	cache, cached := paymentMethodUpsertCache[key]
	paymentMethodUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			paymentMethodColumns,
			paymentMethodColumnsWithDefault,
			paymentMethodColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			paymentMethodColumns,
			paymentMethodPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert PaymentMethod, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "PaymentMethod", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `PaymentMethod` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(paymentMethodType, paymentMethodMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(paymentMethodType, paymentMethodMapping, ret)
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

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for PaymentMethod")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.PaymentMethodCode = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == paymentMethodMapping["payment_method_code"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(paymentMethodType, paymentMethodMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for PaymentMethod")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for PaymentMethod")
	}

CacheNoHooks:
	if !cached {
		paymentMethodUpsertCacheMut.Lock()
		paymentMethodUpsertCache[key] = cache
		paymentMethodUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PaymentMethod record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PaymentMethod) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PaymentMethod provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), paymentMethodPrimaryKeyMapping)
	sql := "DELETE FROM `PaymentMethod` WHERE `payment_method_code`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from PaymentMethod")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for PaymentMethod")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q paymentMethodQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no paymentMethodQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from PaymentMethod")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for PaymentMethod")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PaymentMethodSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PaymentMethod slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(paymentMethodBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), paymentMethodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `PaymentMethod` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, paymentMethodPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from paymentMethod slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for PaymentMethod")
	}

	if len(paymentMethodAfterDeleteHooks) != 0 {
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
func (o *PaymentMethod) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPaymentMethod(ctx, exec, o.PaymentMethodCode)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PaymentMethodSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PaymentMethodSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), paymentMethodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `PaymentMethod`.* FROM `PaymentMethod` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, paymentMethodPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PaymentMethodSlice")
	}

	*o = slice

	return nil
}

// PaymentMethodExists checks if the PaymentMethod row exists.
func PaymentMethodExists(ctx context.Context, exec boil.ContextExecutor, paymentMethodCode int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `PaymentMethod` where `payment_method_code`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, paymentMethodCode)
	}

	row := exec.QueryRowContext(ctx, sql, paymentMethodCode)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if PaymentMethod exists")
	}

	return exists, nil
}
