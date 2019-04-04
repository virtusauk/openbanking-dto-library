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

// FundHistoricalReturn is an object representing the database table.
type FundHistoricalReturn struct {
	FundID        int               `boil:"fund_id" json:"fund_id" toml:"fund_id" yaml:"fund_id"`
	Date          time.Time         `boil:"date" json:"date" toml:"date" yaml:"date"`
	ReturnPercent types.NullDecimal `boil:"return_percent" json:"return_percent,omitempty" toml:"return_percent" yaml:"return_percent,omitempty"`
	BankID        int               `boil:"bank_id" json:"bank_id" toml:"bank_id" yaml:"bank_id"`
	MakerDate     time.Time         `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate   null.Time         `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID       string            `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID     null.String       `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy    null.String       `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate  null.Time         `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *fundHistoricalReturnR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L fundHistoricalReturnL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var FundHistoricalReturnColumns = struct {
	FundID        string
	Date          string
	ReturnPercent string
	BankID        string
	MakerDate     string
	CheckerDate   string
	MakerID       string
	CheckerID     string
	ModifiedBy    string
	ModifiedDate  string
}{
	FundID:        "fund_id",
	Date:          "date",
	ReturnPercent: "return_percent",
	BankID:        "bank_id",
	MakerDate:     "maker_date",
	CheckerDate:   "checker_date",
	MakerID:       "maker_id",
	CheckerID:     "checker_id",
	ModifiedBy:    "modified_by",
	ModifiedDate:  "modified_date",
}

// Generated where

var FundHistoricalReturnWhere = struct {
	FundID        whereHelperint
	Date          whereHelpertime_Time
	ReturnPercent whereHelpertypes_NullDecimal
	BankID        whereHelperint
	MakerDate     whereHelpertime_Time
	CheckerDate   whereHelpernull_Time
	MakerID       whereHelperstring
	CheckerID     whereHelpernull_String
	ModifiedBy    whereHelpernull_String
	ModifiedDate  whereHelpernull_Time
}{
	FundID:        whereHelperint{field: `fund_id`},
	Date:          whereHelpertime_Time{field: `date`},
	ReturnPercent: whereHelpertypes_NullDecimal{field: `return_percent`},
	BankID:        whereHelperint{field: `bank_id`},
	MakerDate:     whereHelpertime_Time{field: `maker_date`},
	CheckerDate:   whereHelpernull_Time{field: `checker_date`},
	MakerID:       whereHelperstring{field: `maker_id`},
	CheckerID:     whereHelpernull_String{field: `checker_id`},
	ModifiedBy:    whereHelpernull_String{field: `modified_by`},
	ModifiedDate:  whereHelpernull_Time{field: `modified_date`},
}

// FundHistoricalReturnRels is where relationship names are stored.
var FundHistoricalReturnRels = struct {
	Fund string
}{
	Fund: "Fund",
}

// fundHistoricalReturnR is where relationships are stored.
type fundHistoricalReturnR struct {
	Fund *FundDescription
}

// NewStruct creates a new relationship struct
func (*fundHistoricalReturnR) NewStruct() *fundHistoricalReturnR {
	return &fundHistoricalReturnR{}
}

// fundHistoricalReturnL is where Load methods for each relationship are stored.
type fundHistoricalReturnL struct{}

var (
	fundHistoricalReturnColumns               = []string{"fund_id", "date", "return_percent", "bank_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	fundHistoricalReturnColumnsWithoutDefault = []string{"fund_id", "date", "return_percent", "bank_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	fundHistoricalReturnColumnsWithDefault    = []string{}
	fundHistoricalReturnPrimaryKeyColumns     = []string{"fund_id", "date"}
)

type (
	// FundHistoricalReturnSlice is an alias for a slice of pointers to FundHistoricalReturn.
	// This should generally be used opposed to []FundHistoricalReturn.
	FundHistoricalReturnSlice []*FundHistoricalReturn
	// FundHistoricalReturnHook is the signature for custom FundHistoricalReturn hook methods
	FundHistoricalReturnHook func(context.Context, boil.ContextExecutor, *FundHistoricalReturn) error

	fundHistoricalReturnQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	fundHistoricalReturnType                 = reflect.TypeOf(&FundHistoricalReturn{})
	fundHistoricalReturnMapping              = queries.MakeStructMapping(fundHistoricalReturnType)
	fundHistoricalReturnPrimaryKeyMapping, _ = queries.BindMapping(fundHistoricalReturnType, fundHistoricalReturnMapping, fundHistoricalReturnPrimaryKeyColumns)
	fundHistoricalReturnInsertCacheMut       sync.RWMutex
	fundHistoricalReturnInsertCache          = make(map[string]insertCache)
	fundHistoricalReturnUpdateCacheMut       sync.RWMutex
	fundHistoricalReturnUpdateCache          = make(map[string]updateCache)
	fundHistoricalReturnUpsertCacheMut       sync.RWMutex
	fundHistoricalReturnUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var fundHistoricalReturnBeforeInsertHooks []FundHistoricalReturnHook
var fundHistoricalReturnBeforeUpdateHooks []FundHistoricalReturnHook
var fundHistoricalReturnBeforeDeleteHooks []FundHistoricalReturnHook
var fundHistoricalReturnBeforeUpsertHooks []FundHistoricalReturnHook

var fundHistoricalReturnAfterInsertHooks []FundHistoricalReturnHook
var fundHistoricalReturnAfterSelectHooks []FundHistoricalReturnHook
var fundHistoricalReturnAfterUpdateHooks []FundHistoricalReturnHook
var fundHistoricalReturnAfterDeleteHooks []FundHistoricalReturnHook
var fundHistoricalReturnAfterUpsertHooks []FundHistoricalReturnHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *FundHistoricalReturn) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fundHistoricalReturnBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *FundHistoricalReturn) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fundHistoricalReturnBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *FundHistoricalReturn) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fundHistoricalReturnBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *FundHistoricalReturn) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fundHistoricalReturnBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *FundHistoricalReturn) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fundHistoricalReturnAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *FundHistoricalReturn) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fundHistoricalReturnAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *FundHistoricalReturn) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fundHistoricalReturnAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *FundHistoricalReturn) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fundHistoricalReturnAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *FundHistoricalReturn) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range fundHistoricalReturnAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddFundHistoricalReturnHook registers your hook function for all future operations.
func AddFundHistoricalReturnHook(hookPoint boil.HookPoint, fundHistoricalReturnHook FundHistoricalReturnHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		fundHistoricalReturnBeforeInsertHooks = append(fundHistoricalReturnBeforeInsertHooks, fundHistoricalReturnHook)
	case boil.BeforeUpdateHook:
		fundHistoricalReturnBeforeUpdateHooks = append(fundHistoricalReturnBeforeUpdateHooks, fundHistoricalReturnHook)
	case boil.BeforeDeleteHook:
		fundHistoricalReturnBeforeDeleteHooks = append(fundHistoricalReturnBeforeDeleteHooks, fundHistoricalReturnHook)
	case boil.BeforeUpsertHook:
		fundHistoricalReturnBeforeUpsertHooks = append(fundHistoricalReturnBeforeUpsertHooks, fundHistoricalReturnHook)
	case boil.AfterInsertHook:
		fundHistoricalReturnAfterInsertHooks = append(fundHistoricalReturnAfterInsertHooks, fundHistoricalReturnHook)
	case boil.AfterSelectHook:
		fundHistoricalReturnAfterSelectHooks = append(fundHistoricalReturnAfterSelectHooks, fundHistoricalReturnHook)
	case boil.AfterUpdateHook:
		fundHistoricalReturnAfterUpdateHooks = append(fundHistoricalReturnAfterUpdateHooks, fundHistoricalReturnHook)
	case boil.AfterDeleteHook:
		fundHistoricalReturnAfterDeleteHooks = append(fundHistoricalReturnAfterDeleteHooks, fundHistoricalReturnHook)
	case boil.AfterUpsertHook:
		fundHistoricalReturnAfterUpsertHooks = append(fundHistoricalReturnAfterUpsertHooks, fundHistoricalReturnHook)
	}
}

// One returns a single fundHistoricalReturn record from the query.
func (q fundHistoricalReturnQuery) One(ctx context.Context, exec boil.ContextExecutor) (*FundHistoricalReturn, error) {
	o := &FundHistoricalReturn{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for FundHistoricalReturn")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all FundHistoricalReturn records from the query.
func (q fundHistoricalReturnQuery) All(ctx context.Context, exec boil.ContextExecutor) (FundHistoricalReturnSlice, error) {
	var o []*FundHistoricalReturn

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to FundHistoricalReturn slice")
	}

	if len(fundHistoricalReturnAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all FundHistoricalReturn records in the query.
func (q fundHistoricalReturnQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count FundHistoricalReturn rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q fundHistoricalReturnQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if FundHistoricalReturn exists")
	}

	return count > 0, nil
}

// Fund pointed to by the foreign key.
func (o *FundHistoricalReturn) Fund(mods ...qm.QueryMod) fundDescriptionQuery {
	queryMods := []qm.QueryMod{
		qm.Where("fund_id=?", o.FundID),
	}

	queryMods = append(queryMods, mods...)

	query := FundDescriptions(queryMods...)
	queries.SetFrom(query.Query, "`FundDescription`")

	return query
}

// LoadFund allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (fundHistoricalReturnL) LoadFund(ctx context.Context, e boil.ContextExecutor, singular bool, maybeFundHistoricalReturn interface{}, mods queries.Applicator) error {
	var slice []*FundHistoricalReturn
	var object *FundHistoricalReturn

	if singular {
		object = maybeFundHistoricalReturn.(*FundHistoricalReturn)
	} else {
		slice = *maybeFundHistoricalReturn.(*[]*FundHistoricalReturn)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &fundHistoricalReturnR{}
		}
		args = append(args, object.FundID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &fundHistoricalReturnR{}
			}

			for _, a := range args {
				if a == obj.FundID {
					continue Outer
				}
			}

			args = append(args, obj.FundID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`FundDescription`), qm.WhereIn(`fund_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load FundDescription")
	}

	var resultSlice []*FundDescription
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice FundDescription")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for FundDescription")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for FundDescription")
	}

	if len(fundHistoricalReturnAfterSelectHooks) != 0 {
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
		object.R.Fund = foreign
		if foreign.R == nil {
			foreign.R = &fundDescriptionR{}
		}
		foreign.R.FundFundHistoricalReturns = append(foreign.R.FundFundHistoricalReturns, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.FundID == foreign.FundID {
				local.R.Fund = foreign
				if foreign.R == nil {
					foreign.R = &fundDescriptionR{}
				}
				foreign.R.FundFundHistoricalReturns = append(foreign.R.FundFundHistoricalReturns, local)
				break
			}
		}
	}

	return nil
}

// SetFund of the fundHistoricalReturn to the related item.
// Sets o.R.Fund to related.
// Adds o to related.R.FundFundHistoricalReturns.
func (o *FundHistoricalReturn) SetFund(ctx context.Context, exec boil.ContextExecutor, insert bool, related *FundDescription) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `FundHistoricalReturn` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"fund_id"}),
		strmangle.WhereClause("`", "`", 0, fundHistoricalReturnPrimaryKeyColumns),
	)
	values := []interface{}{related.FundID, o.FundID, o.Date}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.FundID = related.FundID
	if o.R == nil {
		o.R = &fundHistoricalReturnR{
			Fund: related,
		}
	} else {
		o.R.Fund = related
	}

	if related.R == nil {
		related.R = &fundDescriptionR{
			FundFundHistoricalReturns: FundHistoricalReturnSlice{o},
		}
	} else {
		related.R.FundFundHistoricalReturns = append(related.R.FundFundHistoricalReturns, o)
	}

	return nil
}

// FundHistoricalReturns retrieves all the records using an executor.
func FundHistoricalReturns(mods ...qm.QueryMod) fundHistoricalReturnQuery {
	mods = append(mods, qm.From("`FundHistoricalReturn`"))
	return fundHistoricalReturnQuery{NewQuery(mods...)}
}

// FindFundHistoricalReturn retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindFundHistoricalReturn(ctx context.Context, exec boil.ContextExecutor, fundID int, date time.Time, selectCols ...string) (*FundHistoricalReturn, error) {
	fundHistoricalReturnObj := &FundHistoricalReturn{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `FundHistoricalReturn` where `fund_id`=? AND `date`=?", sel,
	)

	q := queries.Raw(query, fundID, date)

	err := q.Bind(ctx, exec, fundHistoricalReturnObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from FundHistoricalReturn")
	}

	return fundHistoricalReturnObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *FundHistoricalReturn) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no FundHistoricalReturn provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fundHistoricalReturnColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	fundHistoricalReturnInsertCacheMut.RLock()
	cache, cached := fundHistoricalReturnInsertCache[key]
	fundHistoricalReturnInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			fundHistoricalReturnColumns,
			fundHistoricalReturnColumnsWithDefault,
			fundHistoricalReturnColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(fundHistoricalReturnType, fundHistoricalReturnMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(fundHistoricalReturnType, fundHistoricalReturnMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `FundHistoricalReturn` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `FundHistoricalReturn` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `FundHistoricalReturn` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, fundHistoricalReturnPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into FundHistoricalReturn")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.FundID,
		o.Date,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for FundHistoricalReturn")
	}

CacheNoHooks:
	if !cached {
		fundHistoricalReturnInsertCacheMut.Lock()
		fundHistoricalReturnInsertCache[key] = cache
		fundHistoricalReturnInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the FundHistoricalReturn.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *FundHistoricalReturn) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	fundHistoricalReturnUpdateCacheMut.RLock()
	cache, cached := fundHistoricalReturnUpdateCache[key]
	fundHistoricalReturnUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			fundHistoricalReturnColumns,
			fundHistoricalReturnPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update FundHistoricalReturn, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `FundHistoricalReturn` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, fundHistoricalReturnPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(fundHistoricalReturnType, fundHistoricalReturnMapping, append(wl, fundHistoricalReturnPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update FundHistoricalReturn row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for FundHistoricalReturn")
	}

	if !cached {
		fundHistoricalReturnUpdateCacheMut.Lock()
		fundHistoricalReturnUpdateCache[key] = cache
		fundHistoricalReturnUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q fundHistoricalReturnQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for FundHistoricalReturn")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for FundHistoricalReturn")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o FundHistoricalReturnSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fundHistoricalReturnPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `FundHistoricalReturn` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, fundHistoricalReturnPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in fundHistoricalReturn slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all fundHistoricalReturn")
	}
	return rowsAff, nil
}

var mySQLFundHistoricalReturnUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *FundHistoricalReturn) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no FundHistoricalReturn provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(fundHistoricalReturnColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLFundHistoricalReturnUniqueColumns, o)

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

	fundHistoricalReturnUpsertCacheMut.RLock()
	cache, cached := fundHistoricalReturnUpsertCache[key]
	fundHistoricalReturnUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			fundHistoricalReturnColumns,
			fundHistoricalReturnColumnsWithDefault,
			fundHistoricalReturnColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			fundHistoricalReturnColumns,
			fundHistoricalReturnPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert FundHistoricalReturn, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "FundHistoricalReturn", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `FundHistoricalReturn` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(fundHistoricalReturnType, fundHistoricalReturnMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(fundHistoricalReturnType, fundHistoricalReturnMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for FundHistoricalReturn")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(fundHistoricalReturnType, fundHistoricalReturnMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for FundHistoricalReturn")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for FundHistoricalReturn")
	}

CacheNoHooks:
	if !cached {
		fundHistoricalReturnUpsertCacheMut.Lock()
		fundHistoricalReturnUpsertCache[key] = cache
		fundHistoricalReturnUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single FundHistoricalReturn record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *FundHistoricalReturn) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no FundHistoricalReturn provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), fundHistoricalReturnPrimaryKeyMapping)
	sql := "DELETE FROM `FundHistoricalReturn` WHERE `fund_id`=? AND `date`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from FundHistoricalReturn")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for FundHistoricalReturn")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q fundHistoricalReturnQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no fundHistoricalReturnQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from FundHistoricalReturn")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for FundHistoricalReturn")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o FundHistoricalReturnSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no FundHistoricalReturn slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(fundHistoricalReturnBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fundHistoricalReturnPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `FundHistoricalReturn` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, fundHistoricalReturnPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from fundHistoricalReturn slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for FundHistoricalReturn")
	}

	if len(fundHistoricalReturnAfterDeleteHooks) != 0 {
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
func (o *FundHistoricalReturn) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindFundHistoricalReturn(ctx, exec, o.FundID, o.Date)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *FundHistoricalReturnSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := FundHistoricalReturnSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), fundHistoricalReturnPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `FundHistoricalReturn`.* FROM `FundHistoricalReturn` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, fundHistoricalReturnPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in FundHistoricalReturnSlice")
	}

	*o = slice

	return nil
}

// FundHistoricalReturnExists checks if the FundHistoricalReturn row exists.
func FundHistoricalReturnExists(ctx context.Context, exec boil.ContextExecutor, fundID int, date time.Time) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `FundHistoricalReturn` where `fund_id`=? AND `date`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, fundID, date)
	}

	row := exec.QueryRowContext(ctx, sql, fundID, date)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if FundHistoricalReturn exists")
	}

	return exists, nil
}
