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

// TransferMethod is an object representing the database table.
type TransferMethod struct {
	TransferMethodCode int         `boil:"transfer_method_code" json:"transfer_method_code" toml:"transfer_method_code" yaml:"transfer_method_code"`
	Active             null.String `boil:"active" json:"active,omitempty" toml:"active" yaml:"active,omitempty"`
	Description        null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Fee                null.Int    `boil:"fee" json:"fee,omitempty" toml:"fee" yaml:"fee,omitempty"`
	TimeToResolve      null.Int    `boil:"time_to_resolve" json:"time_to_resolve,omitempty" toml:"time_to_resolve" yaml:"time_to_resolve,omitempty"`
	TransferMethodName string      `boil:"transfer_method_name" json:"transfer_method_name" toml:"transfer_method_name" yaml:"transfer_method_name"`
	MakerDate          time.Time   `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate        null.Time   `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID            string      `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID          null.String `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy         null.String `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate       null.Time   `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *transferMethodR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L transferMethodL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TransferMethodColumns = struct {
	TransferMethodCode string
	Active             string
	Description        string
	Fee                string
	TimeToResolve      string
	TransferMethodName string
	MakerDate          string
	CheckerDate        string
	MakerID            string
	CheckerID          string
	ModifiedBy         string
	ModifiedDate       string
}{
	TransferMethodCode: "transfer_method_code",
	Active:             "active",
	Description:        "description",
	Fee:                "fee",
	TimeToResolve:      "time_to_resolve",
	TransferMethodName: "transfer_method_name",
	MakerDate:          "maker_date",
	CheckerDate:        "checker_date",
	MakerID:            "maker_id",
	CheckerID:          "checker_id",
	ModifiedBy:         "modified_by",
	ModifiedDate:       "modified_date",
}

// Generated where

var TransferMethodWhere = struct {
	TransferMethodCode whereHelperint
	Active             whereHelpernull_String
	Description        whereHelpernull_String
	Fee                whereHelpernull_Int
	TimeToResolve      whereHelpernull_Int
	TransferMethodName whereHelperstring
	MakerDate          whereHelpertime_Time
	CheckerDate        whereHelpernull_Time
	MakerID            whereHelperstring
	CheckerID          whereHelpernull_String
	ModifiedBy         whereHelpernull_String
	ModifiedDate       whereHelpernull_Time
}{
	TransferMethodCode: whereHelperint{field: `transfer_method_code`},
	Active:             whereHelpernull_String{field: `active`},
	Description:        whereHelpernull_String{field: `description`},
	Fee:                whereHelpernull_Int{field: `fee`},
	TimeToResolve:      whereHelpernull_Int{field: `time_to_resolve`},
	TransferMethodName: whereHelperstring{field: `transfer_method_name`},
	MakerDate:          whereHelpertime_Time{field: `maker_date`},
	CheckerDate:        whereHelpernull_Time{field: `checker_date`},
	MakerID:            whereHelperstring{field: `maker_id`},
	CheckerID:          whereHelpernull_String{field: `checker_id`},
	ModifiedBy:         whereHelpernull_String{field: `modified_by`},
	ModifiedDate:       whereHelpernull_Time{field: `modified_date`},
}

// TransferMethodRels is where relationship names are stored.
var TransferMethodRels = struct {
	TransferMethodCodeTransfers string
}{
	TransferMethodCodeTransfers: "TransferMethodCodeTransfers",
}

// transferMethodR is where relationships are stored.
type transferMethodR struct {
	TransferMethodCodeTransfers TransferSlice
}

// NewStruct creates a new relationship struct
func (*transferMethodR) NewStruct() *transferMethodR {
	return &transferMethodR{}
}

// transferMethodL is where Load methods for each relationship are stored.
type transferMethodL struct{}

var (
	transferMethodColumns               = []string{"transfer_method_code", "active", "description", "fee", "time_to_resolve", "transfer_method_name", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	transferMethodColumnsWithoutDefault = []string{"active", "description", "fee", "time_to_resolve", "transfer_method_name", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	transferMethodColumnsWithDefault    = []string{"transfer_method_code"}
	transferMethodPrimaryKeyColumns     = []string{"transfer_method_code"}
)

type (
	// TransferMethodSlice is an alias for a slice of pointers to TransferMethod.
	// This should generally be used opposed to []TransferMethod.
	TransferMethodSlice []*TransferMethod
	// TransferMethodHook is the signature for custom TransferMethod hook methods
	TransferMethodHook func(context.Context, boil.ContextExecutor, *TransferMethod) error

	transferMethodQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	transferMethodType                 = reflect.TypeOf(&TransferMethod{})
	transferMethodMapping              = queries.MakeStructMapping(transferMethodType)
	transferMethodPrimaryKeyMapping, _ = queries.BindMapping(transferMethodType, transferMethodMapping, transferMethodPrimaryKeyColumns)
	transferMethodInsertCacheMut       sync.RWMutex
	transferMethodInsertCache          = make(map[string]insertCache)
	transferMethodUpdateCacheMut       sync.RWMutex
	transferMethodUpdateCache          = make(map[string]updateCache)
	transferMethodUpsertCacheMut       sync.RWMutex
	transferMethodUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var transferMethodBeforeInsertHooks []TransferMethodHook
var transferMethodBeforeUpdateHooks []TransferMethodHook
var transferMethodBeforeDeleteHooks []TransferMethodHook
var transferMethodBeforeUpsertHooks []TransferMethodHook

var transferMethodAfterInsertHooks []TransferMethodHook
var transferMethodAfterSelectHooks []TransferMethodHook
var transferMethodAfterUpdateHooks []TransferMethodHook
var transferMethodAfterDeleteHooks []TransferMethodHook
var transferMethodAfterUpsertHooks []TransferMethodHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TransferMethod) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transferMethodBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TransferMethod) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transferMethodBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TransferMethod) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transferMethodBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TransferMethod) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transferMethodBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TransferMethod) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transferMethodAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TransferMethod) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transferMethodAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TransferMethod) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transferMethodAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TransferMethod) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transferMethodAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TransferMethod) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range transferMethodAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTransferMethodHook registers your hook function for all future operations.
func AddTransferMethodHook(hookPoint boil.HookPoint, transferMethodHook TransferMethodHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		transferMethodBeforeInsertHooks = append(transferMethodBeforeInsertHooks, transferMethodHook)
	case boil.BeforeUpdateHook:
		transferMethodBeforeUpdateHooks = append(transferMethodBeforeUpdateHooks, transferMethodHook)
	case boil.BeforeDeleteHook:
		transferMethodBeforeDeleteHooks = append(transferMethodBeforeDeleteHooks, transferMethodHook)
	case boil.BeforeUpsertHook:
		transferMethodBeforeUpsertHooks = append(transferMethodBeforeUpsertHooks, transferMethodHook)
	case boil.AfterInsertHook:
		transferMethodAfterInsertHooks = append(transferMethodAfterInsertHooks, transferMethodHook)
	case boil.AfterSelectHook:
		transferMethodAfterSelectHooks = append(transferMethodAfterSelectHooks, transferMethodHook)
	case boil.AfterUpdateHook:
		transferMethodAfterUpdateHooks = append(transferMethodAfterUpdateHooks, transferMethodHook)
	case boil.AfterDeleteHook:
		transferMethodAfterDeleteHooks = append(transferMethodAfterDeleteHooks, transferMethodHook)
	case boil.AfterUpsertHook:
		transferMethodAfterUpsertHooks = append(transferMethodAfterUpsertHooks, transferMethodHook)
	}
}

// One returns a single transferMethod record from the query.
func (q transferMethodQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TransferMethod, error) {
	o := &TransferMethod{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for TransferMethod")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TransferMethod records from the query.
func (q transferMethodQuery) All(ctx context.Context, exec boil.ContextExecutor) (TransferMethodSlice, error) {
	var o []*TransferMethod

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TransferMethod slice")
	}

	if len(transferMethodAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TransferMethod records in the query.
func (q transferMethodQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count TransferMethod rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q transferMethodQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if TransferMethod exists")
	}

	return count > 0, nil
}

// TransferMethodCodeTransfers retrieves all the Transfer's Transfers with an executor via transfer_method_code column.
func (o *TransferMethod) TransferMethodCodeTransfers(mods ...qm.QueryMod) transferQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`Transfer`.`transfer_method_code`=?", o.TransferMethodCode),
	)

	query := Transfers(queryMods...)
	queries.SetFrom(query.Query, "`Transfer`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`Transfer`.*"})
	}

	return query
}

// LoadTransferMethodCodeTransfers allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (transferMethodL) LoadTransferMethodCodeTransfers(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTransferMethod interface{}, mods queries.Applicator) error {
	var slice []*TransferMethod
	var object *TransferMethod

	if singular {
		object = maybeTransferMethod.(*TransferMethod)
	} else {
		slice = *maybeTransferMethod.(*[]*TransferMethod)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &transferMethodR{}
		}
		args = append(args, object.TransferMethodCode)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &transferMethodR{}
			}

			for _, a := range args {
				if a == obj.TransferMethodCode {
					continue Outer
				}
			}

			args = append(args, obj.TransferMethodCode)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`Transfer`), qm.WhereIn(`transfer_method_code in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Transfer")
	}

	var resultSlice []*Transfer
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Transfer")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on Transfer")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Transfer")
	}

	if len(transferAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.TransferMethodCodeTransfers = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &transferR{}
			}
			foreign.R.TransferMethodCode = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.TransferMethodCode == foreign.TransferMethodCode {
				local.R.TransferMethodCodeTransfers = append(local.R.TransferMethodCodeTransfers, foreign)
				if foreign.R == nil {
					foreign.R = &transferR{}
				}
				foreign.R.TransferMethodCode = local
				break
			}
		}
	}

	return nil
}

// AddTransferMethodCodeTransfers adds the given related objects to the existing relationships
// of the TransferMethod, optionally inserting them as new records.
// Appends related to o.R.TransferMethodCodeTransfers.
// Sets related.R.TransferMethodCode appropriately.
func (o *TransferMethod) AddTransferMethodCodeTransfers(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Transfer) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TransferMethodCode = o.TransferMethodCode
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `Transfer` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"transfer_method_code"}),
				strmangle.WhereClause("`", "`", 0, transferPrimaryKeyColumns),
			)
			values := []interface{}{o.TransferMethodCode, rel.TransferID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.TransferMethodCode = o.TransferMethodCode
		}
	}

	if o.R == nil {
		o.R = &transferMethodR{
			TransferMethodCodeTransfers: related,
		}
	} else {
		o.R.TransferMethodCodeTransfers = append(o.R.TransferMethodCodeTransfers, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &transferR{
				TransferMethodCode: o,
			}
		} else {
			rel.R.TransferMethodCode = o
		}
	}
	return nil
}

// TransferMethods retrieves all the records using an executor.
func TransferMethods(mods ...qm.QueryMod) transferMethodQuery {
	mods = append(mods, qm.From("`TransferMethod`"))
	return transferMethodQuery{NewQuery(mods...)}
}

// FindTransferMethod retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTransferMethod(ctx context.Context, exec boil.ContextExecutor, transferMethodCode int, selectCols ...string) (*TransferMethod, error) {
	transferMethodObj := &TransferMethod{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `TransferMethod` where `transfer_method_code`=?", sel,
	)

	q := queries.Raw(query, transferMethodCode)

	err := q.Bind(ctx, exec, transferMethodObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from TransferMethod")
	}

	return transferMethodObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TransferMethod) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TransferMethod provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(transferMethodColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	transferMethodInsertCacheMut.RLock()
	cache, cached := transferMethodInsertCache[key]
	transferMethodInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			transferMethodColumns,
			transferMethodColumnsWithDefault,
			transferMethodColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(transferMethodType, transferMethodMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(transferMethodType, transferMethodMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `TransferMethod` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `TransferMethod` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `TransferMethod` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, transferMethodPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into TransferMethod")
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

	o.TransferMethodCode = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == transferMethodMapping["TransferMethodCode"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.TransferMethodCode,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TransferMethod")
	}

CacheNoHooks:
	if !cached {
		transferMethodInsertCacheMut.Lock()
		transferMethodInsertCache[key] = cache
		transferMethodInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TransferMethod.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TransferMethod) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	transferMethodUpdateCacheMut.RLock()
	cache, cached := transferMethodUpdateCache[key]
	transferMethodUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			transferMethodColumns,
			transferMethodPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update TransferMethod, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `TransferMethod` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, transferMethodPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(transferMethodType, transferMethodMapping, append(wl, transferMethodPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update TransferMethod row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for TransferMethod")
	}

	if !cached {
		transferMethodUpdateCacheMut.Lock()
		transferMethodUpdateCache[key] = cache
		transferMethodUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q transferMethodQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for TransferMethod")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for TransferMethod")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TransferMethodSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transferMethodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `TransferMethod` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, transferMethodPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in transferMethod slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all transferMethod")
	}
	return rowsAff, nil
}

var mySQLTransferMethodUniqueColumns = []string{
	"transfer_method_code",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TransferMethod) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TransferMethod provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(transferMethodColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTransferMethodUniqueColumns, o)

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

	transferMethodUpsertCacheMut.RLock()
	cache, cached := transferMethodUpsertCache[key]
	transferMethodUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			transferMethodColumns,
			transferMethodColumnsWithDefault,
			transferMethodColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			transferMethodColumns,
			transferMethodPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert TransferMethod, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "TransferMethod", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `TransferMethod` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(transferMethodType, transferMethodMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(transferMethodType, transferMethodMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for TransferMethod")
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

	o.TransferMethodCode = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == transferMethodMapping["transfer_method_code"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(transferMethodType, transferMethodMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for TransferMethod")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TransferMethod")
	}

CacheNoHooks:
	if !cached {
		transferMethodUpsertCacheMut.Lock()
		transferMethodUpsertCache[key] = cache
		transferMethodUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TransferMethod record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TransferMethod) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TransferMethod provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), transferMethodPrimaryKeyMapping)
	sql := "DELETE FROM `TransferMethod` WHERE `transfer_method_code`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from TransferMethod")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for TransferMethod")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q transferMethodQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no transferMethodQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from TransferMethod")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TransferMethod")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TransferMethodSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TransferMethod slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(transferMethodBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transferMethodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `TransferMethod` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, transferMethodPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from transferMethod slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TransferMethod")
	}

	if len(transferMethodAfterDeleteHooks) != 0 {
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
func (o *TransferMethod) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTransferMethod(ctx, exec, o.TransferMethodCode)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TransferMethodSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TransferMethodSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), transferMethodPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `TransferMethod`.* FROM `TransferMethod` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, transferMethodPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TransferMethodSlice")
	}

	*o = slice

	return nil
}

// TransferMethodExists checks if the TransferMethod row exists.
func TransferMethodExists(ctx context.Context, exec boil.ContextExecutor, transferMethodCode int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `TransferMethod` where `transfer_method_code`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, transferMethodCode)
	}

	row := exec.QueryRowContext(ctx, sql, transferMethodCode)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if TransferMethod exists")
	}

	return exists, nil
}
