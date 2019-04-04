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

// TFCommodity is an object representing the database table.
type TFCommodity struct {
	CommodityID    int               `boil:"commodity_id" json:"commodity_id" toml:"commodity_id" yaml:"commodity_id"`
	CommodityName  string            `boil:"commodity_name" json:"commodity_name" toml:"commodity_name" yaml:"commodity_name"`
	CommodityGroup string            `boil:"commodity_group" json:"commodity_group" toml:"commodity_group" yaml:"commodity_group"`
	CommodityRisk  types.NullDecimal `boil:"commodity_risk" json:"commodity_risk,omitempty" toml:"commodity_risk" yaml:"commodity_risk,omitempty"`
	BankID         int               `boil:"bank_id" json:"bank_id" toml:"bank_id" yaml:"bank_id"`
	MakerDate      time.Time         `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate    null.Time         `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID        string            `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID      null.String       `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy     null.String       `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate   null.Time         `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *tFCommodityR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tFCommodityL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TFCommodityColumns = struct {
	CommodityID    string
	CommodityName  string
	CommodityGroup string
	CommodityRisk  string
	BankID         string
	MakerDate      string
	CheckerDate    string
	MakerID        string
	CheckerID      string
	ModifiedBy     string
	ModifiedDate   string
}{
	CommodityID:    "commodity_id",
	CommodityName:  "commodity_name",
	CommodityGroup: "commodity_group",
	CommodityRisk:  "commodity_risk",
	BankID:         "bank_id",
	MakerDate:      "maker_date",
	CheckerDate:    "checker_date",
	MakerID:        "maker_id",
	CheckerID:      "checker_id",
	ModifiedBy:     "modified_by",
	ModifiedDate:   "modified_date",
}

// Generated where

var TFCommodityWhere = struct {
	CommodityID    whereHelperint
	CommodityName  whereHelperstring
	CommodityGroup whereHelperstring
	CommodityRisk  whereHelpertypes_NullDecimal
	BankID         whereHelperint
	MakerDate      whereHelpertime_Time
	CheckerDate    whereHelpernull_Time
	MakerID        whereHelperstring
	CheckerID      whereHelpernull_String
	ModifiedBy     whereHelpernull_String
	ModifiedDate   whereHelpernull_Time
}{
	CommodityID:    whereHelperint{field: `commodity_id`},
	CommodityName:  whereHelperstring{field: `commodity_name`},
	CommodityGroup: whereHelperstring{field: `commodity_group`},
	CommodityRisk:  whereHelpertypes_NullDecimal{field: `commodity_risk`},
	BankID:         whereHelperint{field: `bank_id`},
	MakerDate:      whereHelpertime_Time{field: `maker_date`},
	CheckerDate:    whereHelpernull_Time{field: `checker_date`},
	MakerID:        whereHelperstring{field: `maker_id`},
	CheckerID:      whereHelpernull_String{field: `checker_id`},
	ModifiedBy:     whereHelpernull_String{field: `modified_by`},
	ModifiedDate:   whereHelpernull_Time{field: `modified_date`},
}

// TFCommodityRels is where relationship names are stored.
var TFCommodityRels = struct {
}{}

// tFCommodityR is where relationships are stored.
type tFCommodityR struct {
}

// NewStruct creates a new relationship struct
func (*tFCommodityR) NewStruct() *tFCommodityR {
	return &tFCommodityR{}
}

// tFCommodityL is where Load methods for each relationship are stored.
type tFCommodityL struct{}

var (
	tFCommodityColumns               = []string{"commodity_id", "commodity_name", "commodity_group", "commodity_risk", "bank_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	tFCommodityColumnsWithoutDefault = []string{"commodity_id", "commodity_name", "commodity_group", "commodity_risk", "bank_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	tFCommodityColumnsWithDefault    = []string{}
	tFCommodityPrimaryKeyColumns     = []string{"commodity_id"}
)

type (
	// TFCommoditySlice is an alias for a slice of pointers to TFCommodity.
	// This should generally be used opposed to []TFCommodity.
	TFCommoditySlice []*TFCommodity
	// TFCommodityHook is the signature for custom TFCommodity hook methods
	TFCommodityHook func(context.Context, boil.ContextExecutor, *TFCommodity) error

	tFCommodityQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tFCommodityType                 = reflect.TypeOf(&TFCommodity{})
	tFCommodityMapping              = queries.MakeStructMapping(tFCommodityType)
	tFCommodityPrimaryKeyMapping, _ = queries.BindMapping(tFCommodityType, tFCommodityMapping, tFCommodityPrimaryKeyColumns)
	tFCommodityInsertCacheMut       sync.RWMutex
	tFCommodityInsertCache          = make(map[string]insertCache)
	tFCommodityUpdateCacheMut       sync.RWMutex
	tFCommodityUpdateCache          = make(map[string]updateCache)
	tFCommodityUpsertCacheMut       sync.RWMutex
	tFCommodityUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tFCommodityBeforeInsertHooks []TFCommodityHook
var tFCommodityBeforeUpdateHooks []TFCommodityHook
var tFCommodityBeforeDeleteHooks []TFCommodityHook
var tFCommodityBeforeUpsertHooks []TFCommodityHook

var tFCommodityAfterInsertHooks []TFCommodityHook
var tFCommodityAfterSelectHooks []TFCommodityHook
var tFCommodityAfterUpdateHooks []TFCommodityHook
var tFCommodityAfterDeleteHooks []TFCommodityHook
var tFCommodityAfterUpsertHooks []TFCommodityHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TFCommodity) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFCommodityBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TFCommodity) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFCommodityBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TFCommodity) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFCommodityBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TFCommodity) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFCommodityBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TFCommodity) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFCommodityAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TFCommodity) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFCommodityAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TFCommodity) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFCommodityAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TFCommodity) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFCommodityAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TFCommodity) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFCommodityAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTFCommodityHook registers your hook function for all future operations.
func AddTFCommodityHook(hookPoint boil.HookPoint, tFCommodityHook TFCommodityHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		tFCommodityBeforeInsertHooks = append(tFCommodityBeforeInsertHooks, tFCommodityHook)
	case boil.BeforeUpdateHook:
		tFCommodityBeforeUpdateHooks = append(tFCommodityBeforeUpdateHooks, tFCommodityHook)
	case boil.BeforeDeleteHook:
		tFCommodityBeforeDeleteHooks = append(tFCommodityBeforeDeleteHooks, tFCommodityHook)
	case boil.BeforeUpsertHook:
		tFCommodityBeforeUpsertHooks = append(tFCommodityBeforeUpsertHooks, tFCommodityHook)
	case boil.AfterInsertHook:
		tFCommodityAfterInsertHooks = append(tFCommodityAfterInsertHooks, tFCommodityHook)
	case boil.AfterSelectHook:
		tFCommodityAfterSelectHooks = append(tFCommodityAfterSelectHooks, tFCommodityHook)
	case boil.AfterUpdateHook:
		tFCommodityAfterUpdateHooks = append(tFCommodityAfterUpdateHooks, tFCommodityHook)
	case boil.AfterDeleteHook:
		tFCommodityAfterDeleteHooks = append(tFCommodityAfterDeleteHooks, tFCommodityHook)
	case boil.AfterUpsertHook:
		tFCommodityAfterUpsertHooks = append(tFCommodityAfterUpsertHooks, tFCommodityHook)
	}
}

// One returns a single tFCommodity record from the query.
func (q tFCommodityQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TFCommodity, error) {
	o := &TFCommodity{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for TFCommodity")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TFCommodity records from the query.
func (q tFCommodityQuery) All(ctx context.Context, exec boil.ContextExecutor) (TFCommoditySlice, error) {
	var o []*TFCommodity

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TFCommodity slice")
	}

	if len(tFCommodityAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TFCommodity records in the query.
func (q tFCommodityQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count TFCommodity rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tFCommodityQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if TFCommodity exists")
	}

	return count > 0, nil
}

// TFCommodities retrieves all the records using an executor.
func TFCommodities(mods ...qm.QueryMod) tFCommodityQuery {
	mods = append(mods, qm.From("`TFCommodity`"))
	return tFCommodityQuery{NewQuery(mods...)}
}

// FindTFCommodity retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTFCommodity(ctx context.Context, exec boil.ContextExecutor, commodityID int, selectCols ...string) (*TFCommodity, error) {
	tFCommodityObj := &TFCommodity{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `TFCommodity` where `commodity_id`=?", sel,
	)

	q := queries.Raw(query, commodityID)

	err := q.Bind(ctx, exec, tFCommodityObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from TFCommodity")
	}

	return tFCommodityObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TFCommodity) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TFCommodity provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tFCommodityColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tFCommodityInsertCacheMut.RLock()
	cache, cached := tFCommodityInsertCache[key]
	tFCommodityInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tFCommodityColumns,
			tFCommodityColumnsWithDefault,
			tFCommodityColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tFCommodityType, tFCommodityMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tFCommodityType, tFCommodityMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `TFCommodity` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `TFCommodity` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `TFCommodity` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tFCommodityPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into TFCommodity")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.CommodityID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TFCommodity")
	}

CacheNoHooks:
	if !cached {
		tFCommodityInsertCacheMut.Lock()
		tFCommodityInsertCache[key] = cache
		tFCommodityInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TFCommodity.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TFCommodity) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tFCommodityUpdateCacheMut.RLock()
	cache, cached := tFCommodityUpdateCache[key]
	tFCommodityUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tFCommodityColumns,
			tFCommodityPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update TFCommodity, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `TFCommodity` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tFCommodityPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tFCommodityType, tFCommodityMapping, append(wl, tFCommodityPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update TFCommodity row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for TFCommodity")
	}

	if !cached {
		tFCommodityUpdateCacheMut.Lock()
		tFCommodityUpdateCache[key] = cache
		tFCommodityUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tFCommodityQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for TFCommodity")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for TFCommodity")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TFCommoditySlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tFCommodityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `TFCommodity` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tFCommodityPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tFCommodity slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tFCommodity")
	}
	return rowsAff, nil
}

var mySQLTFCommodityUniqueColumns = []string{
	"commodity_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TFCommodity) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TFCommodity provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tFCommodityColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTFCommodityUniqueColumns, o)

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

	tFCommodityUpsertCacheMut.RLock()
	cache, cached := tFCommodityUpsertCache[key]
	tFCommodityUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tFCommodityColumns,
			tFCommodityColumnsWithDefault,
			tFCommodityColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			tFCommodityColumns,
			tFCommodityPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert TFCommodity, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "TFCommodity", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `TFCommodity` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tFCommodityType, tFCommodityMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tFCommodityType, tFCommodityMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for TFCommodity")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tFCommodityType, tFCommodityMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for TFCommodity")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TFCommodity")
	}

CacheNoHooks:
	if !cached {
		tFCommodityUpsertCacheMut.Lock()
		tFCommodityUpsertCache[key] = cache
		tFCommodityUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TFCommodity record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TFCommodity) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TFCommodity provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tFCommodityPrimaryKeyMapping)
	sql := "DELETE FROM `TFCommodity` WHERE `commodity_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from TFCommodity")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for TFCommodity")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tFCommodityQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tFCommodityQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from TFCommodity")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TFCommodity")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TFCommoditySlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TFCommodity slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(tFCommodityBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tFCommodityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `TFCommodity` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tFCommodityPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tFCommodity slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TFCommodity")
	}

	if len(tFCommodityAfterDeleteHooks) != 0 {
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
func (o *TFCommodity) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTFCommodity(ctx, exec, o.CommodityID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TFCommoditySlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TFCommoditySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tFCommodityPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `TFCommodity`.* FROM `TFCommodity` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tFCommodityPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TFCommoditySlice")
	}

	*o = slice

	return nil
}

// TFCommodityExists checks if the TFCommodity row exists.
func TFCommodityExists(ctx context.Context, exec boil.ContextExecutor, commodityID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `TFCommodity` where `commodity_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, commodityID)
	}

	row := exec.QueryRowContext(ctx, sql, commodityID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if TFCommodity exists")
	}

	return exists, nil
}
