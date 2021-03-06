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

// CustomerGoal is an object representing the database table.
type CustomerGoal struct {
	CustomerGoalID int           `boil:"customer_goal_id" json:"customer_goal_id" toml:"customer_goal_id" yaml:"customer_goal_id"`
	PartyID        int           `boil:"party_id" json:"party_id" toml:"party_id" yaml:"party_id"`
	GoalID         int           `boil:"goal_id" json:"goal_id" toml:"goal_id" yaml:"goal_id"`
	Percentage     types.Decimal `boil:"percentage" json:"percentage" toml:"percentage" yaml:"percentage"`
	TargetAmount   types.Decimal `boil:"target_amount" json:"target_amount" toml:"target_amount" yaml:"target_amount"`
	TargetDate     time.Time     `boil:"target_date" json:"target_date" toml:"target_date" yaml:"target_date"`
	Status         null.String   `boil:"status" json:"status,omitempty" toml:"status" yaml:"status,omitempty"`
	BankID         int           `boil:"bank_id" json:"bank_id" toml:"bank_id" yaml:"bank_id"`
	MakerDate      time.Time     `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate    null.Time     `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID        string        `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID      null.String   `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy     null.String   `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate   null.Time     `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *customerGoalR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L customerGoalL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var CustomerGoalColumns = struct {
	CustomerGoalID string
	PartyID        string
	GoalID         string
	Percentage     string
	TargetAmount   string
	TargetDate     string
	Status         string
	BankID         string
	MakerDate      string
	CheckerDate    string
	MakerID        string
	CheckerID      string
	ModifiedBy     string
	ModifiedDate   string
}{
	CustomerGoalID: "customer_goal_id",
	PartyID:        "party_id",
	GoalID:         "goal_id",
	Percentage:     "percentage",
	TargetAmount:   "target_amount",
	TargetDate:     "target_date",
	Status:         "status",
	BankID:         "bank_id",
	MakerDate:      "maker_date",
	CheckerDate:    "checker_date",
	MakerID:        "maker_id",
	CheckerID:      "checker_id",
	ModifiedBy:     "modified_by",
	ModifiedDate:   "modified_date",
}

// Generated where

var CustomerGoalWhere = struct {
	CustomerGoalID whereHelperint
	PartyID        whereHelperint
	GoalID         whereHelperint
	Percentage     whereHelpertypes_Decimal
	TargetAmount   whereHelpertypes_Decimal
	TargetDate     whereHelpertime_Time
	Status         whereHelpernull_String
	BankID         whereHelperint
	MakerDate      whereHelpertime_Time
	CheckerDate    whereHelpernull_Time
	MakerID        whereHelperstring
	CheckerID      whereHelpernull_String
	ModifiedBy     whereHelpernull_String
	ModifiedDate   whereHelpernull_Time
}{
	CustomerGoalID: whereHelperint{field: `customer_goal_id`},
	PartyID:        whereHelperint{field: `party_id`},
	GoalID:         whereHelperint{field: `goal_id`},
	Percentage:     whereHelpertypes_Decimal{field: `percentage`},
	TargetAmount:   whereHelpertypes_Decimal{field: `target_amount`},
	TargetDate:     whereHelpertime_Time{field: `target_date`},
	Status:         whereHelpernull_String{field: `status`},
	BankID:         whereHelperint{field: `bank_id`},
	MakerDate:      whereHelpertime_Time{field: `maker_date`},
	CheckerDate:    whereHelpernull_Time{field: `checker_date`},
	MakerID:        whereHelperstring{field: `maker_id`},
	CheckerID:      whereHelpernull_String{field: `checker_id`},
	ModifiedBy:     whereHelpernull_String{field: `modified_by`},
	ModifiedDate:   whereHelpernull_Time{field: `modified_date`},
}

// CustomerGoalRels is where relationship names are stored.
var CustomerGoalRels = struct {
}{}

// customerGoalR is where relationships are stored.
type customerGoalR struct {
}

// NewStruct creates a new relationship struct
func (*customerGoalR) NewStruct() *customerGoalR {
	return &customerGoalR{}
}

// customerGoalL is where Load methods for each relationship are stored.
type customerGoalL struct{}

var (
	customerGoalColumns               = []string{"customer_goal_id", "party_id", "goal_id", "percentage", "target_amount", "target_date", "status", "bank_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	customerGoalColumnsWithoutDefault = []string{"customer_goal_id", "party_id", "goal_id", "percentage", "target_amount", "target_date", "status", "bank_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	customerGoalColumnsWithDefault    = []string{}
	customerGoalPrimaryKeyColumns     = []string{"customer_goal_id"}
)

type (
	// CustomerGoalSlice is an alias for a slice of pointers to CustomerGoal.
	// This should generally be used opposed to []CustomerGoal.
	CustomerGoalSlice []*CustomerGoal
	// CustomerGoalHook is the signature for custom CustomerGoal hook methods
	CustomerGoalHook func(context.Context, boil.ContextExecutor, *CustomerGoal) error

	customerGoalQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	customerGoalType                 = reflect.TypeOf(&CustomerGoal{})
	customerGoalMapping              = queries.MakeStructMapping(customerGoalType)
	customerGoalPrimaryKeyMapping, _ = queries.BindMapping(customerGoalType, customerGoalMapping, customerGoalPrimaryKeyColumns)
	customerGoalInsertCacheMut       sync.RWMutex
	customerGoalInsertCache          = make(map[string]insertCache)
	customerGoalUpdateCacheMut       sync.RWMutex
	customerGoalUpdateCache          = make(map[string]updateCache)
	customerGoalUpsertCacheMut       sync.RWMutex
	customerGoalUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var customerGoalBeforeInsertHooks []CustomerGoalHook
var customerGoalBeforeUpdateHooks []CustomerGoalHook
var customerGoalBeforeDeleteHooks []CustomerGoalHook
var customerGoalBeforeUpsertHooks []CustomerGoalHook

var customerGoalAfterInsertHooks []CustomerGoalHook
var customerGoalAfterSelectHooks []CustomerGoalHook
var customerGoalAfterUpdateHooks []CustomerGoalHook
var customerGoalAfterDeleteHooks []CustomerGoalHook
var customerGoalAfterUpsertHooks []CustomerGoalHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CustomerGoal) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerGoalBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CustomerGoal) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerGoalBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CustomerGoal) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerGoalBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CustomerGoal) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerGoalBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CustomerGoal) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerGoalAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CustomerGoal) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerGoalAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CustomerGoal) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerGoalAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CustomerGoal) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerGoalAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CustomerGoal) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range customerGoalAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCustomerGoalHook registers your hook function for all future operations.
func AddCustomerGoalHook(hookPoint boil.HookPoint, customerGoalHook CustomerGoalHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		customerGoalBeforeInsertHooks = append(customerGoalBeforeInsertHooks, customerGoalHook)
	case boil.BeforeUpdateHook:
		customerGoalBeforeUpdateHooks = append(customerGoalBeforeUpdateHooks, customerGoalHook)
	case boil.BeforeDeleteHook:
		customerGoalBeforeDeleteHooks = append(customerGoalBeforeDeleteHooks, customerGoalHook)
	case boil.BeforeUpsertHook:
		customerGoalBeforeUpsertHooks = append(customerGoalBeforeUpsertHooks, customerGoalHook)
	case boil.AfterInsertHook:
		customerGoalAfterInsertHooks = append(customerGoalAfterInsertHooks, customerGoalHook)
	case boil.AfterSelectHook:
		customerGoalAfterSelectHooks = append(customerGoalAfterSelectHooks, customerGoalHook)
	case boil.AfterUpdateHook:
		customerGoalAfterUpdateHooks = append(customerGoalAfterUpdateHooks, customerGoalHook)
	case boil.AfterDeleteHook:
		customerGoalAfterDeleteHooks = append(customerGoalAfterDeleteHooks, customerGoalHook)
	case boil.AfterUpsertHook:
		customerGoalAfterUpsertHooks = append(customerGoalAfterUpsertHooks, customerGoalHook)
	}
}

// One returns a single customerGoal record from the query.
func (q customerGoalQuery) One(ctx context.Context, exec boil.ContextExecutor) (*CustomerGoal, error) {
	o := &CustomerGoal{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for CustomerGoal")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all CustomerGoal records from the query.
func (q customerGoalQuery) All(ctx context.Context, exec boil.ContextExecutor) (CustomerGoalSlice, error) {
	var o []*CustomerGoal

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CustomerGoal slice")
	}

	if len(customerGoalAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all CustomerGoal records in the query.
func (q customerGoalQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count CustomerGoal rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q customerGoalQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if CustomerGoal exists")
	}

	return count > 0, nil
}

// CustomerGoals retrieves all the records using an executor.
func CustomerGoals(mods ...qm.QueryMod) customerGoalQuery {
	mods = append(mods, qm.From("`CustomerGoal`"))
	return customerGoalQuery{NewQuery(mods...)}
}

// FindCustomerGoal retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCustomerGoal(ctx context.Context, exec boil.ContextExecutor, customerGoalID int, selectCols ...string) (*CustomerGoal, error) {
	customerGoalObj := &CustomerGoal{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `CustomerGoal` where `customer_goal_id`=?", sel,
	)

	q := queries.Raw(query, customerGoalID)

	err := q.Bind(ctx, exec, customerGoalObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from CustomerGoal")
	}

	return customerGoalObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *CustomerGoal) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no CustomerGoal provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(customerGoalColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	customerGoalInsertCacheMut.RLock()
	cache, cached := customerGoalInsertCache[key]
	customerGoalInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			customerGoalColumns,
			customerGoalColumnsWithDefault,
			customerGoalColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(customerGoalType, customerGoalMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(customerGoalType, customerGoalMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `CustomerGoal` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `CustomerGoal` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `CustomerGoal` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, customerGoalPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into CustomerGoal")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.CustomerGoalID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for CustomerGoal")
	}

CacheNoHooks:
	if !cached {
		customerGoalInsertCacheMut.Lock()
		customerGoalInsertCache[key] = cache
		customerGoalInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the CustomerGoal.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *CustomerGoal) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	customerGoalUpdateCacheMut.RLock()
	cache, cached := customerGoalUpdateCache[key]
	customerGoalUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			customerGoalColumns,
			customerGoalPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update CustomerGoal, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `CustomerGoal` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, customerGoalPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(customerGoalType, customerGoalMapping, append(wl, customerGoalPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update CustomerGoal row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for CustomerGoal")
	}

	if !cached {
		customerGoalUpdateCacheMut.Lock()
		customerGoalUpdateCache[key] = cache
		customerGoalUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q customerGoalQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for CustomerGoal")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for CustomerGoal")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CustomerGoalSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customerGoalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `CustomerGoal` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, customerGoalPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in customerGoal slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all customerGoal")
	}
	return rowsAff, nil
}

var mySQLCustomerGoalUniqueColumns = []string{
	"customer_goal_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *CustomerGoal) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no CustomerGoal provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(customerGoalColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLCustomerGoalUniqueColumns, o)

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

	customerGoalUpsertCacheMut.RLock()
	cache, cached := customerGoalUpsertCache[key]
	customerGoalUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			customerGoalColumns,
			customerGoalColumnsWithDefault,
			customerGoalColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			customerGoalColumns,
			customerGoalPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert CustomerGoal, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "CustomerGoal", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `CustomerGoal` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(customerGoalType, customerGoalMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(customerGoalType, customerGoalMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for CustomerGoal")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(customerGoalType, customerGoalMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for CustomerGoal")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for CustomerGoal")
	}

CacheNoHooks:
	if !cached {
		customerGoalUpsertCacheMut.Lock()
		customerGoalUpsertCache[key] = cache
		customerGoalUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single CustomerGoal record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CustomerGoal) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CustomerGoal provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), customerGoalPrimaryKeyMapping)
	sql := "DELETE FROM `CustomerGoal` WHERE `customer_goal_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from CustomerGoal")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for CustomerGoal")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q customerGoalQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no customerGoalQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from CustomerGoal")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for CustomerGoal")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CustomerGoalSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no CustomerGoal slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(customerGoalBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customerGoalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `CustomerGoal` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, customerGoalPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from customerGoal slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for CustomerGoal")
	}

	if len(customerGoalAfterDeleteHooks) != 0 {
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
func (o *CustomerGoal) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindCustomerGoal(ctx, exec, o.CustomerGoalID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CustomerGoalSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := CustomerGoalSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), customerGoalPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `CustomerGoal`.* FROM `CustomerGoal` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, customerGoalPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CustomerGoalSlice")
	}

	*o = slice

	return nil
}

// CustomerGoalExists checks if the CustomerGoal row exists.
func CustomerGoalExists(ctx context.Context, exec boil.ContextExecutor, customerGoalID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `CustomerGoal` where `customer_goal_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, customerGoalID)
	}

	row := exec.QueryRowContext(ctx, sql, customerGoalID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if CustomerGoal exists")
	}

	return exists, nil
}
