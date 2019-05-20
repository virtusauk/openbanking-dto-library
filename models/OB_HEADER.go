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

// OBHEADER is an object representing the database table.
type OBHEADER struct {
	ObHeaderID             int         `boil:"ob_header_id" json:"ob_header_id" toml:"ob_header_id" yaml:"ob_header_id"`
	PaymentID              null.Int    `boil:"payment_id" json:"payment_id,omitempty" toml:"payment_id" yaml:"payment_id,omitempty"`
	AccountIdentification  null.String `boil:"account_identification" json:"account_identification,omitempty" toml:"account_identification" yaml:"account_identification,omitempty"`
	IdempotencyKey         null.String `boil:"idempotency_key" json:"idempotency_key,omitempty" toml:"idempotency_key" yaml:"idempotency_key,omitempty"`
	JWSSignatureRequest    null.String `boil:"jws_signature_Request" json:"jws_signature_Request,omitempty" toml:"jws_signature_Request" yaml:"jws_signature_Request,omitempty"`
	InteractionID          null.String `boil:"interaction_id" json:"interaction_id,omitempty" toml:"interaction_id" yaml:"interaction_id,omitempty"`
	CustomerLastLoggedTime null.Time   `boil:"customer_last_logged_time" json:"customer_last_logged_time,omitempty" toml:"customer_last_logged_time" yaml:"customer_last_logged_time,omitempty"`
	CustomerIPAddress      null.String `boil:"customer_ip_address" json:"customer_ip_address,omitempty" toml:"customer_ip_address" yaml:"customer_ip_address,omitempty"`
	JWSSignatureResponse   null.String `boil:"jws_signature_Response" json:"jws_signature_Response,omitempty" toml:"jws_signature_Response" yaml:"jws_signature_Response,omitempty"`
	CustomerUserAgent      null.String `boil:"customer_user_agent" json:"customer_user_agent,omitempty" toml:"customer_user_agent" yaml:"customer_user_agent,omitempty"`

	R *oBHEADERR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L oBHEADERL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OBHEADERColumns = struct {
	ObHeaderID             string
	PaymentID              string
	AccountIdentification  string
	IdempotencyKey         string
	JWSSignatureRequest    string
	InteractionID          string
	CustomerLastLoggedTime string
	CustomerIPAddress      string
	JWSSignatureResponse   string
	CustomerUserAgent      string
}{
	ObHeaderID:             "ob_header_id",
	PaymentID:              "payment_id",
	AccountIdentification:  "account_identification",
	IdempotencyKey:         "idempotency_key",
	JWSSignatureRequest:    "jws_signature_Request",
	InteractionID:          "interaction_id",
	CustomerLastLoggedTime: "customer_last_logged_time",
	CustomerIPAddress:      "customer_ip_address",
	JWSSignatureResponse:   "jws_signature_Response",
	CustomerUserAgent:      "customer_user_agent",
}

// Generated where

var OBHEADERWhere = struct {
	ObHeaderID             whereHelperint
	PaymentID              whereHelpernull_Int
	AccountIdentification  whereHelpernull_String
	IdempotencyKey         whereHelpernull_String
	JWSSignatureRequest    whereHelpernull_String
	InteractionID          whereHelpernull_String
	CustomerLastLoggedTime whereHelpernull_Time
	CustomerIPAddress      whereHelpernull_String
	JWSSignatureResponse   whereHelpernull_String
	CustomerUserAgent      whereHelpernull_String
}{
	ObHeaderID:             whereHelperint{field: `ob_header_id`},
	PaymentID:              whereHelpernull_Int{field: `payment_id`},
	AccountIdentification:  whereHelpernull_String{field: `account_identification`},
	IdempotencyKey:         whereHelpernull_String{field: `idempotency_key`},
	JWSSignatureRequest:    whereHelpernull_String{field: `jws_signature_Request`},
	InteractionID:          whereHelpernull_String{field: `interaction_id`},
	CustomerLastLoggedTime: whereHelpernull_Time{field: `customer_last_logged_time`},
	CustomerIPAddress:      whereHelpernull_String{field: `customer_ip_address`},
	JWSSignatureResponse:   whereHelpernull_String{field: `jws_signature_Response`},
	CustomerUserAgent:      whereHelpernull_String{field: `customer_user_agent`},
}

// OBHEADERRels is where relationship names are stored.
var OBHEADERRels = struct {
}{}

// oBHEADERR is where relationships are stored.
type oBHEADERR struct {
}

// NewStruct creates a new relationship struct
func (*oBHEADERR) NewStruct() *oBHEADERR {
	return &oBHEADERR{}
}

// oBHEADERL is where Load methods for each relationship are stored.
type oBHEADERL struct{}

var (
	oBHEADERColumns               = []string{"ob_header_id", "payment_id", "account_identification", "idempotency_key", "jws_signature_Request", "interaction_id", "customer_last_logged_time", "customer_ip_address", "jws_signature_Response", "customer_user_agent"}
	oBHEADERColumnsWithoutDefault = []string{"ob_header_id", "payment_id", "account_identification", "idempotency_key", "jws_signature_Request", "interaction_id", "customer_last_logged_time", "customer_ip_address", "jws_signature_Response", "customer_user_agent"}
	oBHEADERColumnsWithDefault    = []string{}
	oBHEADERPrimaryKeyColumns     = []string{"ob_header_id"}
)

type (
	// OBHEADERSlice is an alias for a slice of pointers to OBHEADER.
	// This should generally be used opposed to []OBHEADER.
	OBHEADERSlice []*OBHEADER
	// OBHEADERHook is the signature for custom OBHEADER hook methods
	OBHEADERHook func(context.Context, boil.ContextExecutor, *OBHEADER) error

	oBHEADERQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	oBHEADERType                 = reflect.TypeOf(&OBHEADER{})
	oBHEADERMapping              = queries.MakeStructMapping(oBHEADERType)
	oBHEADERPrimaryKeyMapping, _ = queries.BindMapping(oBHEADERType, oBHEADERMapping, oBHEADERPrimaryKeyColumns)
	oBHEADERInsertCacheMut       sync.RWMutex
	oBHEADERInsertCache          = make(map[string]insertCache)
	oBHEADERUpdateCacheMut       sync.RWMutex
	oBHEADERUpdateCache          = make(map[string]updateCache)
	oBHEADERUpsertCacheMut       sync.RWMutex
	oBHEADERUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var oBHEADERBeforeInsertHooks []OBHEADERHook
var oBHEADERBeforeUpdateHooks []OBHEADERHook
var oBHEADERBeforeDeleteHooks []OBHEADERHook
var oBHEADERBeforeUpsertHooks []OBHEADERHook

var oBHEADERAfterInsertHooks []OBHEADERHook
var oBHEADERAfterSelectHooks []OBHEADERHook
var oBHEADERAfterUpdateHooks []OBHEADERHook
var oBHEADERAfterDeleteHooks []OBHEADERHook
var oBHEADERAfterUpsertHooks []OBHEADERHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OBHEADER) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBHEADERBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OBHEADER) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBHEADERBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OBHEADER) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBHEADERBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OBHEADER) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBHEADERBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OBHEADER) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBHEADERAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OBHEADER) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBHEADERAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OBHEADER) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBHEADERAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OBHEADER) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBHEADERAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OBHEADER) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBHEADERAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOBHEADERHook registers your hook function for all future operations.
func AddOBHEADERHook(hookPoint boil.HookPoint, oBHEADERHook OBHEADERHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		oBHEADERBeforeInsertHooks = append(oBHEADERBeforeInsertHooks, oBHEADERHook)
	case boil.BeforeUpdateHook:
		oBHEADERBeforeUpdateHooks = append(oBHEADERBeforeUpdateHooks, oBHEADERHook)
	case boil.BeforeDeleteHook:
		oBHEADERBeforeDeleteHooks = append(oBHEADERBeforeDeleteHooks, oBHEADERHook)
	case boil.BeforeUpsertHook:
		oBHEADERBeforeUpsertHooks = append(oBHEADERBeforeUpsertHooks, oBHEADERHook)
	case boil.AfterInsertHook:
		oBHEADERAfterInsertHooks = append(oBHEADERAfterInsertHooks, oBHEADERHook)
	case boil.AfterSelectHook:
		oBHEADERAfterSelectHooks = append(oBHEADERAfterSelectHooks, oBHEADERHook)
	case boil.AfterUpdateHook:
		oBHEADERAfterUpdateHooks = append(oBHEADERAfterUpdateHooks, oBHEADERHook)
	case boil.AfterDeleteHook:
		oBHEADERAfterDeleteHooks = append(oBHEADERAfterDeleteHooks, oBHEADERHook)
	case boil.AfterUpsertHook:
		oBHEADERAfterUpsertHooks = append(oBHEADERAfterUpsertHooks, oBHEADERHook)
	}
}

// One returns a single oBHEADER record from the query.
func (q oBHEADERQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OBHEADER, error) {
	o := &OBHEADER{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for OB_HEADER")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OBHEADER records from the query.
func (q oBHEADERQuery) All(ctx context.Context, exec boil.ContextExecutor) (OBHEADERSlice, error) {
	var o []*OBHEADER

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OBHEADER slice")
	}

	if len(oBHEADERAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OBHEADER records in the query.
func (q oBHEADERQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count OB_HEADER rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q oBHEADERQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if OB_HEADER exists")
	}

	return count > 0, nil
}

// OBHEADERS retrieves all the records using an executor.
func OBHEADERS(mods ...qm.QueryMod) oBHEADERQuery {
	mods = append(mods, qm.From("`OB_HEADER`"))
	return oBHEADERQuery{NewQuery(mods...)}
}

// FindOBHEADER retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOBHEADER(ctx context.Context, exec boil.ContextExecutor, obHeaderID int, selectCols ...string) (*OBHEADER, error) {
	oBHEADERObj := &OBHEADER{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `OB_HEADER` where `ob_header_id`=?", sel,
	)

	q := queries.Raw(query, obHeaderID)

	err := q.Bind(ctx, exec, oBHEADERObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from OB_HEADER")
	}

	return oBHEADERObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OBHEADER) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no OB_HEADER provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(oBHEADERColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	oBHEADERInsertCacheMut.RLock()
	cache, cached := oBHEADERInsertCache[key]
	oBHEADERInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			oBHEADERColumns,
			oBHEADERColumnsWithDefault,
			oBHEADERColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(oBHEADERType, oBHEADERMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(oBHEADERType, oBHEADERMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `OB_HEADER` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `OB_HEADER` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `OB_HEADER` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, oBHEADERPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into OB_HEADER")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ObHeaderID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for OB_HEADER")
	}

CacheNoHooks:
	if !cached {
		oBHEADERInsertCacheMut.Lock()
		oBHEADERInsertCache[key] = cache
		oBHEADERInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OBHEADER.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OBHEADER) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	oBHEADERUpdateCacheMut.RLock()
	cache, cached := oBHEADERUpdateCache[key]
	oBHEADERUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			oBHEADERColumns,
			oBHEADERPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update OB_HEADER, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `OB_HEADER` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, oBHEADERPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(oBHEADERType, oBHEADERMapping, append(wl, oBHEADERPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update OB_HEADER row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for OB_HEADER")
	}

	if !cached {
		oBHEADERUpdateCacheMut.Lock()
		oBHEADERUpdateCache[key] = cache
		oBHEADERUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q oBHEADERQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for OB_HEADER")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for OB_HEADER")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OBHEADERSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oBHEADERPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `OB_HEADER` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oBHEADERPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in oBHEADER slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all oBHEADER")
	}
	return rowsAff, nil
}

var mySQLOBHEADERUniqueColumns = []string{
	"ob_header_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OBHEADER) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no OB_HEADER provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(oBHEADERColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOBHEADERUniqueColumns, o)

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

	oBHEADERUpsertCacheMut.RLock()
	cache, cached := oBHEADERUpsertCache[key]
	oBHEADERUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			oBHEADERColumns,
			oBHEADERColumnsWithDefault,
			oBHEADERColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			oBHEADERColumns,
			oBHEADERPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert OB_HEADER, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "OB_HEADER", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `OB_HEADER` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(oBHEADERType, oBHEADERMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(oBHEADERType, oBHEADERMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for OB_HEADER")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(oBHEADERType, oBHEADERMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for OB_HEADER")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for OB_HEADER")
	}

CacheNoHooks:
	if !cached {
		oBHEADERUpsertCacheMut.Lock()
		oBHEADERUpsertCache[key] = cache
		oBHEADERUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OBHEADER record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OBHEADER) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OBHEADER provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), oBHEADERPrimaryKeyMapping)
	sql := "DELETE FROM `OB_HEADER` WHERE `ob_header_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from OB_HEADER")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for OB_HEADER")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q oBHEADERQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no oBHEADERQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from OB_HEADER")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for OB_HEADER")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OBHEADERSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OBHEADER slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(oBHEADERBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oBHEADERPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `OB_HEADER` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oBHEADERPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from oBHEADER slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for OB_HEADER")
	}

	if len(oBHEADERAfterDeleteHooks) != 0 {
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
func (o *OBHEADER) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOBHEADER(ctx, exec, o.ObHeaderID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OBHEADERSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OBHEADERSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oBHEADERPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `OB_HEADER`.* FROM `OB_HEADER` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oBHEADERPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OBHEADERSlice")
	}

	*o = slice

	return nil
}

// OBHEADERExists checks if the OBHEADER row exists.
func OBHEADERExists(ctx context.Context, exec boil.ContextExecutor, obHeaderID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `OB_HEADER` where `ob_header_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, obHeaderID)
	}

	row := exec.QueryRowContext(ctx, sql, obHeaderID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if OB_HEADER exists")
	}

	return exists, nil
}