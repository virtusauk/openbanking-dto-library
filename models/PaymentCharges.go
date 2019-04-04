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

// PaymentCharge is an object representing the database table.
type PaymentCharge struct {
	PaymentID    int               `boil:"payment_id" json:"payment_id" toml:"payment_id" yaml:"payment_id"`
	ChargeBearer null.String       `boil:"charge_bearer" json:"charge_bearer,omitempty" toml:"charge_bearer" yaml:"charge_bearer,omitempty"`
	Type         null.String       `boil:"type" json:"type,omitempty" toml:"type" yaml:"type,omitempty"`
	Amount       types.NullDecimal `boil:"amount" json:"amount,omitempty" toml:"amount" yaml:"amount,omitempty"`
	Currency     null.String       `boil:"currency" json:"currency,omitempty" toml:"currency" yaml:"currency,omitempty"`
	MakerDate    time.Time         `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate  null.Time         `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID      string            `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID    null.String       `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy   null.String       `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate null.Time         `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *paymentChargeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L paymentChargeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var PaymentChargeColumns = struct {
	PaymentID    string
	ChargeBearer string
	Type         string
	Amount       string
	Currency     string
	MakerDate    string
	CheckerDate  string
	MakerID      string
	CheckerID    string
	ModifiedBy   string
	ModifiedDate string
}{
	PaymentID:    "payment_id",
	ChargeBearer: "charge_bearer",
	Type:         "type",
	Amount:       "amount",
	Currency:     "currency",
	MakerDate:    "maker_date",
	CheckerDate:  "checker_date",
	MakerID:      "maker_id",
	CheckerID:    "checker_id",
	ModifiedBy:   "modified_by",
	ModifiedDate: "modified_date",
}

// Generated where

var PaymentChargeWhere = struct {
	PaymentID    whereHelperint
	ChargeBearer whereHelpernull_String
	Type         whereHelpernull_String
	Amount       whereHelpertypes_NullDecimal
	Currency     whereHelpernull_String
	MakerDate    whereHelpertime_Time
	CheckerDate  whereHelpernull_Time
	MakerID      whereHelperstring
	CheckerID    whereHelpernull_String
	ModifiedBy   whereHelpernull_String
	ModifiedDate whereHelpernull_Time
}{
	PaymentID:    whereHelperint{field: `payment_id`},
	ChargeBearer: whereHelpernull_String{field: `charge_bearer`},
	Type:         whereHelpernull_String{field: `type`},
	Amount:       whereHelpertypes_NullDecimal{field: `amount`},
	Currency:     whereHelpernull_String{field: `currency`},
	MakerDate:    whereHelpertime_Time{field: `maker_date`},
	CheckerDate:  whereHelpernull_Time{field: `checker_date`},
	MakerID:      whereHelperstring{field: `maker_id`},
	CheckerID:    whereHelpernull_String{field: `checker_id`},
	ModifiedBy:   whereHelpernull_String{field: `modified_by`},
	ModifiedDate: whereHelpernull_Time{field: `modified_date`},
}

// PaymentChargeRels is where relationship names are stored.
var PaymentChargeRels = struct {
	Payment string
}{
	Payment: "Payment",
}

// paymentChargeR is where relationships are stored.
type paymentChargeR struct {
	Payment *PaymentInitiation
}

// NewStruct creates a new relationship struct
func (*paymentChargeR) NewStruct() *paymentChargeR {
	return &paymentChargeR{}
}

// paymentChargeL is where Load methods for each relationship are stored.
type paymentChargeL struct{}

var (
	paymentChargeColumns               = []string{"payment_id", "charge_bearer", "type", "amount", "currency", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	paymentChargeColumnsWithoutDefault = []string{"charge_bearer", "type", "amount", "currency", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	paymentChargeColumnsWithDefault    = []string{"payment_id"}
	paymentChargePrimaryKeyColumns     = []string{"payment_id"}
)

type (
	// PaymentChargeSlice is an alias for a slice of pointers to PaymentCharge.
	// This should generally be used opposed to []PaymentCharge.
	PaymentChargeSlice []*PaymentCharge
	// PaymentChargeHook is the signature for custom PaymentCharge hook methods
	PaymentChargeHook func(context.Context, boil.ContextExecutor, *PaymentCharge) error

	paymentChargeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	paymentChargeType                 = reflect.TypeOf(&PaymentCharge{})
	paymentChargeMapping              = queries.MakeStructMapping(paymentChargeType)
	paymentChargePrimaryKeyMapping, _ = queries.BindMapping(paymentChargeType, paymentChargeMapping, paymentChargePrimaryKeyColumns)
	paymentChargeInsertCacheMut       sync.RWMutex
	paymentChargeInsertCache          = make(map[string]insertCache)
	paymentChargeUpdateCacheMut       sync.RWMutex
	paymentChargeUpdateCache          = make(map[string]updateCache)
	paymentChargeUpsertCacheMut       sync.RWMutex
	paymentChargeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var paymentChargeBeforeInsertHooks []PaymentChargeHook
var paymentChargeBeforeUpdateHooks []PaymentChargeHook
var paymentChargeBeforeDeleteHooks []PaymentChargeHook
var paymentChargeBeforeUpsertHooks []PaymentChargeHook

var paymentChargeAfterInsertHooks []PaymentChargeHook
var paymentChargeAfterSelectHooks []PaymentChargeHook
var paymentChargeAfterUpdateHooks []PaymentChargeHook
var paymentChargeAfterDeleteHooks []PaymentChargeHook
var paymentChargeAfterUpsertHooks []PaymentChargeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *PaymentCharge) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentChargeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *PaymentCharge) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentChargeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *PaymentCharge) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentChargeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *PaymentCharge) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentChargeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *PaymentCharge) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentChargeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *PaymentCharge) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentChargeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *PaymentCharge) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentChargeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *PaymentCharge) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentChargeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *PaymentCharge) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range paymentChargeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddPaymentChargeHook registers your hook function for all future operations.
func AddPaymentChargeHook(hookPoint boil.HookPoint, paymentChargeHook PaymentChargeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		paymentChargeBeforeInsertHooks = append(paymentChargeBeforeInsertHooks, paymentChargeHook)
	case boil.BeforeUpdateHook:
		paymentChargeBeforeUpdateHooks = append(paymentChargeBeforeUpdateHooks, paymentChargeHook)
	case boil.BeforeDeleteHook:
		paymentChargeBeforeDeleteHooks = append(paymentChargeBeforeDeleteHooks, paymentChargeHook)
	case boil.BeforeUpsertHook:
		paymentChargeBeforeUpsertHooks = append(paymentChargeBeforeUpsertHooks, paymentChargeHook)
	case boil.AfterInsertHook:
		paymentChargeAfterInsertHooks = append(paymentChargeAfterInsertHooks, paymentChargeHook)
	case boil.AfterSelectHook:
		paymentChargeAfterSelectHooks = append(paymentChargeAfterSelectHooks, paymentChargeHook)
	case boil.AfterUpdateHook:
		paymentChargeAfterUpdateHooks = append(paymentChargeAfterUpdateHooks, paymentChargeHook)
	case boil.AfterDeleteHook:
		paymentChargeAfterDeleteHooks = append(paymentChargeAfterDeleteHooks, paymentChargeHook)
	case boil.AfterUpsertHook:
		paymentChargeAfterUpsertHooks = append(paymentChargeAfterUpsertHooks, paymentChargeHook)
	}
}

// One returns a single paymentCharge record from the query.
func (q paymentChargeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*PaymentCharge, error) {
	o := &PaymentCharge{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for PaymentCharges")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all PaymentCharge records from the query.
func (q paymentChargeQuery) All(ctx context.Context, exec boil.ContextExecutor) (PaymentChargeSlice, error) {
	var o []*PaymentCharge

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to PaymentCharge slice")
	}

	if len(paymentChargeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all PaymentCharge records in the query.
func (q paymentChargeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count PaymentCharges rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q paymentChargeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if PaymentCharges exists")
	}

	return count > 0, nil
}

// Payment pointed to by the foreign key.
func (o *PaymentCharge) Payment(mods ...qm.QueryMod) paymentInitiationQuery {
	queryMods := []qm.QueryMod{
		qm.Where("payment_id=?", o.PaymentID),
	}

	queryMods = append(queryMods, mods...)

	query := PaymentInitiations(queryMods...)
	queries.SetFrom(query.Query, "`PaymentInitiation`")

	return query
}

// LoadPayment allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (paymentChargeL) LoadPayment(ctx context.Context, e boil.ContextExecutor, singular bool, maybePaymentCharge interface{}, mods queries.Applicator) error {
	var slice []*PaymentCharge
	var object *PaymentCharge

	if singular {
		object = maybePaymentCharge.(*PaymentCharge)
	} else {
		slice = *maybePaymentCharge.(*[]*PaymentCharge)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &paymentChargeR{}
		}
		args = append(args, object.PaymentID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &paymentChargeR{}
			}

			for _, a := range args {
				if a == obj.PaymentID {
					continue Outer
				}
			}

			args = append(args, obj.PaymentID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`PaymentInitiation`), qm.WhereIn(`payment_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load PaymentInitiation")
	}

	var resultSlice []*PaymentInitiation
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice PaymentInitiation")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for PaymentInitiation")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for PaymentInitiation")
	}

	if len(paymentChargeAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Payment = foreign
		if foreign.R == nil {
			foreign.R = &paymentInitiationR{}
		}
		foreign.R.PaymentPaymentCharge = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PaymentID == foreign.PaymentID {
				local.R.Payment = foreign
				if foreign.R == nil {
					foreign.R = &paymentInitiationR{}
				}
				foreign.R.PaymentPaymentCharge = local
				break
			}
		}
	}

	return nil
}

// SetPayment of the paymentCharge to the related item.
// Sets o.R.Payment to related.
// Adds o to related.R.PaymentPaymentCharge.
func (o *PaymentCharge) SetPayment(ctx context.Context, exec boil.ContextExecutor, insert bool, related *PaymentInitiation) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `PaymentCharges` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"payment_id"}),
		strmangle.WhereClause("`", "`", 0, paymentChargePrimaryKeyColumns),
	)
	values := []interface{}{related.PaymentID, o.PaymentID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PaymentID = related.PaymentID
	if o.R == nil {
		o.R = &paymentChargeR{
			Payment: related,
		}
	} else {
		o.R.Payment = related
	}

	if related.R == nil {
		related.R = &paymentInitiationR{
			PaymentPaymentCharge: o,
		}
	} else {
		related.R.PaymentPaymentCharge = o
	}

	return nil
}

// PaymentCharges retrieves all the records using an executor.
func PaymentCharges(mods ...qm.QueryMod) paymentChargeQuery {
	mods = append(mods, qm.From("`PaymentCharges`"))
	return paymentChargeQuery{NewQuery(mods...)}
}

// FindPaymentCharge retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindPaymentCharge(ctx context.Context, exec boil.ContextExecutor, paymentID int, selectCols ...string) (*PaymentCharge, error) {
	paymentChargeObj := &PaymentCharge{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `PaymentCharges` where `payment_id`=?", sel,
	)

	q := queries.Raw(query, paymentID)

	err := q.Bind(ctx, exec, paymentChargeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from PaymentCharges")
	}

	return paymentChargeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *PaymentCharge) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no PaymentCharges provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(paymentChargeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	paymentChargeInsertCacheMut.RLock()
	cache, cached := paymentChargeInsertCache[key]
	paymentChargeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			paymentChargeColumns,
			paymentChargeColumnsWithDefault,
			paymentChargeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(paymentChargeType, paymentChargeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(paymentChargeType, paymentChargeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `PaymentCharges` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `PaymentCharges` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `PaymentCharges` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, paymentChargePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into PaymentCharges")
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

	o.PaymentID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == paymentChargeMapping["PaymentID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PaymentID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for PaymentCharges")
	}

CacheNoHooks:
	if !cached {
		paymentChargeInsertCacheMut.Lock()
		paymentChargeInsertCache[key] = cache
		paymentChargeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the PaymentCharge.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *PaymentCharge) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	paymentChargeUpdateCacheMut.RLock()
	cache, cached := paymentChargeUpdateCache[key]
	paymentChargeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			paymentChargeColumns,
			paymentChargePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update PaymentCharges, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `PaymentCharges` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, paymentChargePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(paymentChargeType, paymentChargeMapping, append(wl, paymentChargePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update PaymentCharges row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for PaymentCharges")
	}

	if !cached {
		paymentChargeUpdateCacheMut.Lock()
		paymentChargeUpdateCache[key] = cache
		paymentChargeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q paymentChargeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for PaymentCharges")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for PaymentCharges")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o PaymentChargeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), paymentChargePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `PaymentCharges` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, paymentChargePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in paymentCharge slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all paymentCharge")
	}
	return rowsAff, nil
}

var mySQLPaymentChargeUniqueColumns = []string{
	"payment_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *PaymentCharge) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no PaymentCharges provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(paymentChargeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLPaymentChargeUniqueColumns, o)

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

	paymentChargeUpsertCacheMut.RLock()
	cache, cached := paymentChargeUpsertCache[key]
	paymentChargeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			paymentChargeColumns,
			paymentChargeColumnsWithDefault,
			paymentChargeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			paymentChargeColumns,
			paymentChargePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert PaymentCharges, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "PaymentCharges", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `PaymentCharges` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(paymentChargeType, paymentChargeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(paymentChargeType, paymentChargeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for PaymentCharges")
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

	o.PaymentID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == paymentChargeMapping["payment_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(paymentChargeType, paymentChargeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for PaymentCharges")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for PaymentCharges")
	}

CacheNoHooks:
	if !cached {
		paymentChargeUpsertCacheMut.Lock()
		paymentChargeUpsertCache[key] = cache
		paymentChargeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single PaymentCharge record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *PaymentCharge) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PaymentCharge provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), paymentChargePrimaryKeyMapping)
	sql := "DELETE FROM `PaymentCharges` WHERE `payment_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from PaymentCharges")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for PaymentCharges")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q paymentChargeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no paymentChargeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from PaymentCharges")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for PaymentCharges")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o PaymentChargeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no PaymentCharge slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(paymentChargeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), paymentChargePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `PaymentCharges` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, paymentChargePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from paymentCharge slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for PaymentCharges")
	}

	if len(paymentChargeAfterDeleteHooks) != 0 {
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
func (o *PaymentCharge) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindPaymentCharge(ctx, exec, o.PaymentID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *PaymentChargeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := PaymentChargeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), paymentChargePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `PaymentCharges`.* FROM `PaymentCharges` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, paymentChargePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in PaymentChargeSlice")
	}

	*o = slice

	return nil
}

// PaymentChargeExists checks if the PaymentCharge row exists.
func PaymentChargeExists(ctx context.Context, exec boil.ContextExecutor, paymentID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `PaymentCharges` where `payment_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, paymentID)
	}

	row := exec.QueryRowContext(ctx, sql, paymentID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if PaymentCharges exists")
	}

	return exists, nil
}
