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

// CodeDatum is an object representing the database table.
type CodeDatum struct {
	CDID         int         `boil:"cd_id" json:"cd_id" toml:"cd_id" yaml:"cd_id"`
	Code         string      `boil:"code" json:"code" toml:"code" yaml:"code"`
	Group        string      `boil:"group" json:"group" toml:"group" yaml:"group"`
	CodeVal      string      `boil:"code_val" json:"code_val" toml:"code_val" yaml:"code_val"`
	Description  string      `boil:"description" json:"description" toml:"description" yaml:"description"`
	MakerDate    time.Time   `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate  null.Time   `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID      string      `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID    null.String `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy   null.String `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate null.Time   `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *codeDatumR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L codeDatumL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CodeDatumColumns = struct {
	CDID         string
	Code         string
	Group        string
	CodeVal      string
	Description  string
	MakerDate    string
	CheckerDate  string
	MakerID      string
	CheckerID    string
	ModifiedBy   string
	ModifiedDate string
}{
	CDID:         "cd_id",
	Code:         "code",
	Group:        "group",
	CodeVal:      "code_val",
	Description:  "description",
	MakerDate:    "maker_date",
	CheckerDate:  "checker_date",
	MakerID:      "maker_id",
	CheckerID:    "checker_id",
	ModifiedBy:   "modified_by",
	ModifiedDate: "modified_date",
}

// Generated where

var CodeDatumWhere = struct {
	CDID         whereHelperint
	Code         whereHelperstring
	Group        whereHelperstring
	CodeVal      whereHelperstring
	Description  whereHelperstring
	MakerDate    whereHelpertime_Time
	CheckerDate  whereHelpernull_Time
	MakerID      whereHelperstring
	CheckerID    whereHelpernull_String
	ModifiedBy   whereHelpernull_String
	ModifiedDate whereHelpernull_Time
}{
	CDID:         whereHelperint{field: `cd_id`},
	Code:         whereHelperstring{field: `code`},
	Group:        whereHelperstring{field: `group`},
	CodeVal:      whereHelperstring{field: `code_val`},
	Description:  whereHelperstring{field: `description`},
	MakerDate:    whereHelpertime_Time{field: `maker_date`},
	CheckerDate:  whereHelpernull_Time{field: `checker_date`},
	MakerID:      whereHelperstring{field: `maker_id`},
	CheckerID:    whereHelpernull_String{field: `checker_id`},
	ModifiedBy:   whereHelpernull_String{field: `modified_by`},
	ModifiedDate: whereHelpernull_Time{field: `modified_date`},
}

// CodeDatumRels is where relationship names are stored.
var CodeDatumRels = struct {
}{}

// codeDatumR is where relationships are stored.
type codeDatumR struct {
}

// NewStruct creates a new relationship struct
func (*codeDatumR) NewStruct() *codeDatumR {
	return &codeDatumR{}
}

// codeDatumL is where Load methods for each relationship are stored.
type codeDatumL struct{}

var (
	codeDatumColumns               = []string{"cd_id", "code", "group", "code_val", "description", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	codeDatumColumnsWithoutDefault = []string{"cd_id", "code", "group", "code_val", "description", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	codeDatumColumnsWithDefault    = []string{}
	codeDatumPrimaryKeyColumns     = []string{"cd_id"}
)

type (
	// CodeDatumSlice is an alias for a slice of pointers to CodeDatum.
	// This should generally be used opposed to []CodeDatum.
	CodeDatumSlice []*CodeDatum
	// CodeDatumHook is the signature for custom CodeDatum hook methods
	CodeDatumHook func(context.Context, boil.ContextExecutor, *CodeDatum) error

	codeDatumQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	codeDatumType                 = reflect.TypeOf(&CodeDatum{})
	codeDatumMapping              = queries.MakeStructMapping(codeDatumType)
	codeDatumPrimaryKeyMapping, _ = queries.BindMapping(codeDatumType, codeDatumMapping, codeDatumPrimaryKeyColumns)
	codeDatumInsertCacheMut       sync.RWMutex
	codeDatumInsertCache          = make(map[string]insertCache)
	codeDatumUpdateCacheMut       sync.RWMutex
	codeDatumUpdateCache          = make(map[string]updateCache)
	codeDatumUpsertCacheMut       sync.RWMutex
	codeDatumUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var codeDatumBeforeInsertHooks []CodeDatumHook
var codeDatumBeforeUpdateHooks []CodeDatumHook
var codeDatumBeforeDeleteHooks []CodeDatumHook
var codeDatumBeforeUpsertHooks []CodeDatumHook

var codeDatumAfterInsertHooks []CodeDatumHook
var codeDatumAfterSelectHooks []CodeDatumHook
var codeDatumAfterUpdateHooks []CodeDatumHook
var codeDatumAfterDeleteHooks []CodeDatumHook
var codeDatumAfterUpsertHooks []CodeDatumHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CodeDatum) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range codeDatumBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CodeDatum) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range codeDatumBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CodeDatum) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range codeDatumBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CodeDatum) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range codeDatumBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CodeDatum) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range codeDatumAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CodeDatum) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range codeDatumAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CodeDatum) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range codeDatumAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CodeDatum) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range codeDatumAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CodeDatum) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range codeDatumAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCodeDatumHook registers your hook function for all future operations.
func AddCodeDatumHook(hookPoint boil.HookPoint, codeDatumHook CodeDatumHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		codeDatumBeforeInsertHooks = append(codeDatumBeforeInsertHooks, codeDatumHook)
	case boil.BeforeUpdateHook:
		codeDatumBeforeUpdateHooks = append(codeDatumBeforeUpdateHooks, codeDatumHook)
	case boil.BeforeDeleteHook:
		codeDatumBeforeDeleteHooks = append(codeDatumBeforeDeleteHooks, codeDatumHook)
	case boil.BeforeUpsertHook:
		codeDatumBeforeUpsertHooks = append(codeDatumBeforeUpsertHooks, codeDatumHook)
	case boil.AfterInsertHook:
		codeDatumAfterInsertHooks = append(codeDatumAfterInsertHooks, codeDatumHook)
	case boil.AfterSelectHook:
		codeDatumAfterSelectHooks = append(codeDatumAfterSelectHooks, codeDatumHook)
	case boil.AfterUpdateHook:
		codeDatumAfterUpdateHooks = append(codeDatumAfterUpdateHooks, codeDatumHook)
	case boil.AfterDeleteHook:
		codeDatumAfterDeleteHooks = append(codeDatumAfterDeleteHooks, codeDatumHook)
	case boil.AfterUpsertHook:
		codeDatumAfterUpsertHooks = append(codeDatumAfterUpsertHooks, codeDatumHook)
	}
}

// One returns a single codeDatum record from the query.
func (q codeDatumQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CodeDatum, error) {
	o := &CodeDatum{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for CodeData")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CodeDatum records from the query.
func (q codeDatumQuery) All(ctx context.Context, exec boil.ContextExecutor) (CodeDatumSlice, error) {
	var o []*CodeDatum

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CodeDatum slice")
	}

	if len(codeDatumAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CodeDatum records in the query.
func (q codeDatumQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count CodeData rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q codeDatumQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if CodeData exists")
	}

	return count > 0, nil
}

// CodeData retrieves all the records using an executor.
func CodeData(mods ...qm.QueryMod) codeDatumQuery {
	mods = append(mods, qm.From("`CodeData`"))
	return codeDatumQuery{NewQuery(mods...)}
}

// FindCodeDatum retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCodeDatum(ctx context.Context, exec boil.ContextExecutor, cDID int, selectCols ...string) (*CodeDatum, error) {
	codeDatumObj := &CodeDatum{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `CodeData` where `cd_id`=?", sel,
	)

	q := queries.Raw(query, cDID)

	err := q.Bind(ctx, exec, codeDatumObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from CodeData")
	}

	return codeDatumObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CodeDatum) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no CodeData provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(codeDatumColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	codeDatumInsertCacheMut.RLock()
	cache, cached := codeDatumInsertCache[key]
	codeDatumInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			codeDatumColumns,
			codeDatumColumnsWithDefault,
			codeDatumColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(codeDatumType, codeDatumMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(codeDatumType, codeDatumMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `CodeData` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `CodeData` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `CodeData` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, codeDatumPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into CodeData")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.CDID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for CodeData")
	}

CacheNoHooks:
	if !cached {
		codeDatumInsertCacheMut.Lock()
		codeDatumInsertCache[key] = cache
		codeDatumInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CodeDatum.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CodeDatum) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	codeDatumUpdateCacheMut.RLock()
	cache, cached := codeDatumUpdateCache[key]
	codeDatumUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			codeDatumColumns,
			codeDatumPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update CodeData, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `CodeData` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, codeDatumPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(codeDatumType, codeDatumMapping, append(wl, codeDatumPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update CodeData row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for CodeData")
	}

	if !cached {
		codeDatumUpdateCacheMut.Lock()
		codeDatumUpdateCache[key] = cache
		codeDatumUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q codeDatumQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for CodeData")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for CodeData")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CodeDatumSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), codeDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `CodeData` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, codeDatumPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in codeDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all codeDatum")
	}
	return rowsAff, nil
}

var mySQLCodeDatumUniqueColumns = []string{
	"cd_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CodeDatum) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no CodeData provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(codeDatumColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCodeDatumUniqueColumns, o)

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

	codeDatumUpsertCacheMut.RLock()
	cache, cached := codeDatumUpsertCache[key]
	codeDatumUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			codeDatumColumns,
			codeDatumColumnsWithDefault,
			codeDatumColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			codeDatumColumns,
			codeDatumPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert CodeData, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "CodeData", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `CodeData` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(codeDatumType, codeDatumMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(codeDatumType, codeDatumMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for CodeData")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(codeDatumType, codeDatumMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for CodeData")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for CodeData")
	}

CacheNoHooks:
	if !cached {
		codeDatumUpsertCacheMut.Lock()
		codeDatumUpsertCache[key] = cache
		codeDatumUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CodeDatum record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CodeDatum) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CodeDatum provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), codeDatumPrimaryKeyMapping)
	sql := "DELETE FROM `CodeData` WHERE `cd_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from CodeData")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for CodeData")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q codeDatumQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no codeDatumQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from CodeData")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for CodeData")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CodeDatumSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CodeDatum slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(codeDatumBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), codeDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `CodeData` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, codeDatumPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from codeDatum slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for CodeData")
	}

	if len(codeDatumAfterDeleteHooks) != 0 {
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
func (o *CodeDatum) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCodeDatum(ctx, exec, o.CDID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CodeDatumSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CodeDatumSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), codeDatumPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `CodeData`.* FROM `CodeData` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, codeDatumPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CodeDatumSlice")
	}

	*o = slice

	return nil
}

// CodeDatumExists checks if the CodeDatum row exists.
func CodeDatumExists(ctx context.Context, exec boil.ContextExecutor, cDID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `CodeData` where `cd_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, cDID)
	}

	row := exec.QueryRowContext(ctx, sql, cDID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if CodeData exists")
	}

	return exists, nil
}