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

// MccCode is an object representing the database table.
type MccCode struct {
	MCCCode             int         `boil:"mcc_code" json:"mcc_code" toml:"mcc_code" yaml:"mcc_code"`
	CombinedDescription null.String `boil:"combined_description" json:"combined_description,omitempty" toml:"combined_description" yaml:"combined_description,omitempty"`
	EditedDescription   null.String `boil:"edited_description" json:"edited_description,omitempty" toml:"edited_description" yaml:"edited_description,omitempty"`
	IrsDescription      null.String `boil:"irs_description" json:"irs_description,omitempty" toml:"irs_description" yaml:"irs_description,omitempty"`
	IrsReportable       null.String `boil:"irs_reportable" json:"irs_reportable,omitempty" toml:"irs_reportable" yaml:"irs_reportable,omitempty"`
	UsdaDescription     null.String `boil:"usda_description" json:"usda_description,omitempty" toml:"usda_description" yaml:"usda_description,omitempty"`
	MakerDate           time.Time   `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate         null.Time   `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID             string      `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID           null.String `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy          null.String `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate        null.Time   `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *mccCodeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L mccCodeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var MccCodeColumns = struct {
	MCCCode             string
	CombinedDescription string
	EditedDescription   string
	IrsDescription      string
	IrsReportable       string
	UsdaDescription     string
	MakerDate           string
	CheckerDate         string
	MakerID             string
	CheckerID           string
	ModifiedBy          string
	ModifiedDate        string
}{
	MCCCode:             "mcc_code",
	CombinedDescription: "combined_description",
	EditedDescription:   "edited_description",
	IrsDescription:      "irs_description",
	IrsReportable:       "irs_reportable",
	UsdaDescription:     "usda_description",
	MakerDate:           "maker_date",
	CheckerDate:         "checker_date",
	MakerID:             "maker_id",
	CheckerID:           "checker_id",
	ModifiedBy:          "modified_by",
	ModifiedDate:        "modified_date",
}

// Generated where

var MccCodeWhere = struct {
	MCCCode             whereHelperint
	CombinedDescription whereHelpernull_String
	EditedDescription   whereHelpernull_String
	IrsDescription      whereHelpernull_String
	IrsReportable       whereHelpernull_String
	UsdaDescription     whereHelpernull_String
	MakerDate           whereHelpertime_Time
	CheckerDate         whereHelpernull_Time
	MakerID             whereHelperstring
	CheckerID           whereHelpernull_String
	ModifiedBy          whereHelpernull_String
	ModifiedDate        whereHelpernull_Time
}{
	MCCCode:             whereHelperint{field: `mcc_code`},
	CombinedDescription: whereHelpernull_String{field: `combined_description`},
	EditedDescription:   whereHelpernull_String{field: `edited_description`},
	IrsDescription:      whereHelpernull_String{field: `irs_description`},
	IrsReportable:       whereHelpernull_String{field: `irs_reportable`},
	UsdaDescription:     whereHelpernull_String{field: `usda_description`},
	MakerDate:           whereHelpertime_Time{field: `maker_date`},
	CheckerDate:         whereHelpernull_Time{field: `checker_date`},
	MakerID:             whereHelperstring{field: `maker_id`},
	CheckerID:           whereHelpernull_String{field: `checker_id`},
	ModifiedBy:          whereHelpernull_String{field: `modified_by`},
	ModifiedDate:        whereHelpernull_Time{field: `modified_date`},
}

// MccCodeRels is where relationship names are stored.
var MccCodeRels = struct {
	CategoryBudgets string
}{
	CategoryBudgets: "CategoryBudgets",
}

// mccCodeR is where relationships are stored.
type mccCodeR struct {
	CategoryBudgets BudgetSlice
}

// NewStruct creates a new relationship struct
func (*mccCodeR) NewStruct() *mccCodeR {
	return &mccCodeR{}
}

// mccCodeL is where Load methods for each relationship are stored.
type mccCodeL struct{}

var (
	mccCodeColumns               = []string{"mcc_code", "combined_description", "edited_description", "irs_description", "irs_reportable", "usda_description", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	mccCodeColumnsWithoutDefault = []string{"mcc_code", "combined_description", "edited_description", "irs_description", "irs_reportable", "usda_description", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	mccCodeColumnsWithDefault    = []string{}
	mccCodePrimaryKeyColumns     = []string{"mcc_code"}
)

type (
	// MccCodeSlice is an alias for a slice of pointers to MccCode.
	// This should generally be used opposed to []MccCode.
	MccCodeSlice []*MccCode
	// MccCodeHook is the signature for custom MccCode hook methods
	MccCodeHook func(context.Context, boil.ContextExecutor, *MccCode) error

	mccCodeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	mccCodeType                 = reflect.TypeOf(&MccCode{})
	mccCodeMapping              = queries.MakeStructMapping(mccCodeType)
	mccCodePrimaryKeyMapping, _ = queries.BindMapping(mccCodeType, mccCodeMapping, mccCodePrimaryKeyColumns)
	mccCodeInsertCacheMut       sync.RWMutex
	mccCodeInsertCache          = make(map[string]insertCache)
	mccCodeUpdateCacheMut       sync.RWMutex
	mccCodeUpdateCache          = make(map[string]updateCache)
	mccCodeUpsertCacheMut       sync.RWMutex
	mccCodeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var mccCodeBeforeInsertHooks []MccCodeHook
var mccCodeBeforeUpdateHooks []MccCodeHook
var mccCodeBeforeDeleteHooks []MccCodeHook
var mccCodeBeforeUpsertHooks []MccCodeHook

var mccCodeAfterInsertHooks []MccCodeHook
var mccCodeAfterSelectHooks []MccCodeHook
var mccCodeAfterUpdateHooks []MccCodeHook
var mccCodeAfterDeleteHooks []MccCodeHook
var mccCodeAfterUpsertHooks []MccCodeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *MccCode) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mccCodeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *MccCode) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mccCodeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *MccCode) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mccCodeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *MccCode) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mccCodeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *MccCode) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mccCodeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *MccCode) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mccCodeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *MccCode) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mccCodeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *MccCode) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mccCodeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *MccCode) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range mccCodeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddMccCodeHook registers your hook function for all future operations.
func AddMccCodeHook(hookPoint boil.HookPoint, mccCodeHook MccCodeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		mccCodeBeforeInsertHooks = append(mccCodeBeforeInsertHooks, mccCodeHook)
	case boil.BeforeUpdateHook:
		mccCodeBeforeUpdateHooks = append(mccCodeBeforeUpdateHooks, mccCodeHook)
	case boil.BeforeDeleteHook:
		mccCodeBeforeDeleteHooks = append(mccCodeBeforeDeleteHooks, mccCodeHook)
	case boil.BeforeUpsertHook:
		mccCodeBeforeUpsertHooks = append(mccCodeBeforeUpsertHooks, mccCodeHook)
	case boil.AfterInsertHook:
		mccCodeAfterInsertHooks = append(mccCodeAfterInsertHooks, mccCodeHook)
	case boil.AfterSelectHook:
		mccCodeAfterSelectHooks = append(mccCodeAfterSelectHooks, mccCodeHook)
	case boil.AfterUpdateHook:
		mccCodeAfterUpdateHooks = append(mccCodeAfterUpdateHooks, mccCodeHook)
	case boil.AfterDeleteHook:
		mccCodeAfterDeleteHooks = append(mccCodeAfterDeleteHooks, mccCodeHook)
	case boil.AfterUpsertHook:
		mccCodeAfterUpsertHooks = append(mccCodeAfterUpsertHooks, mccCodeHook)
	}
}

// One returns a single mccCode record from the query.
func (q mccCodeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*MccCode, error) {
	o := &MccCode{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for MccCode")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all MccCode records from the query.
func (q mccCodeQuery) All(ctx context.Context, exec boil.ContextExecutor) (MccCodeSlice, error) {
	var o []*MccCode

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to MccCode slice")
	}

	if len(mccCodeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all MccCode records in the query.
func (q mccCodeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count MccCode rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q mccCodeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if MccCode exists")
	}

	return count > 0, nil
}

// CategoryBudgets retrieves all the Budget's Budgets with an executor via category_id column.
func (o *MccCode) CategoryBudgets(mods ...qm.QueryMod) budgetQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`Budget`.`category_id`=?", o.MCCCode),
	)

	query := Budgets(queryMods...)
	queries.SetFrom(query.Query, "`Budget`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`Budget`.*"})
	}

	return query
}

// LoadCategoryBudgets allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (mccCodeL) LoadCategoryBudgets(ctx context.Context, e boil.ContextExecutor, singular bool, maybeMccCode interface{}, mods queries.Applicator) error {
	var slice []*MccCode
	var object *MccCode

	if singular {
		object = maybeMccCode.(*MccCode)
	} else {
		slice = *maybeMccCode.(*[]*MccCode)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &mccCodeR{}
		}
		args = append(args, object.MCCCode)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &mccCodeR{}
			}

			for _, a := range args {
				if a == obj.MCCCode {
					continue Outer
				}
			}

			args = append(args, obj.MCCCode)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`Budget`), qm.WhereIn(`category_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Budget")
	}

	var resultSlice []*Budget
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Budget")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on Budget")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Budget")
	}

	if len(budgetAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.CategoryBudgets = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &budgetR{}
			}
			foreign.R.Category = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.MCCCode == foreign.CategoryID {
				local.R.CategoryBudgets = append(local.R.CategoryBudgets, foreign)
				if foreign.R == nil {
					foreign.R = &budgetR{}
				}
				foreign.R.Category = local
				break
			}
		}
	}

	return nil
}

// AddCategoryBudgets adds the given related objects to the existing relationships
// of the MccCode, optionally inserting them as new records.
// Appends related to o.R.CategoryBudgets.
// Sets related.R.Category appropriately.
func (o *MccCode) AddCategoryBudgets(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Budget) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.CategoryID = o.MCCCode
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `Budget` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"category_id"}),
				strmangle.WhereClause("`", "`", 0, budgetPrimaryKeyColumns),
			)
			values := []interface{}{o.MCCCode, rel.BudgetID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.CategoryID = o.MCCCode
		}
	}

	if o.R == nil {
		o.R = &mccCodeR{
			CategoryBudgets: related,
		}
	} else {
		o.R.CategoryBudgets = append(o.R.CategoryBudgets, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &budgetR{
				Category: o,
			}
		} else {
			rel.R.Category = o
		}
	}
	return nil
}

// MccCodes retrieves all the records using an executor.
func MccCodes(mods ...qm.QueryMod) mccCodeQuery {
	mods = append(mods, qm.From("`MccCode`"))
	return mccCodeQuery{NewQuery(mods...)}
}

// FindMccCode retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindMccCode(ctx context.Context, exec boil.ContextExecutor, mCCCode int, selectCols ...string) (*MccCode, error) {
	mccCodeObj := &MccCode{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `MccCode` where `mcc_code`=?", sel,
	)

	q := queries.Raw(query, mCCCode)

	err := q.Bind(ctx, exec, mccCodeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from MccCode")
	}

	return mccCodeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *MccCode) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no MccCode provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mccCodeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	mccCodeInsertCacheMut.RLock()
	cache, cached := mccCodeInsertCache[key]
	mccCodeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			mccCodeColumns,
			mccCodeColumnsWithDefault,
			mccCodeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(mccCodeType, mccCodeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(mccCodeType, mccCodeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `MccCode` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `MccCode` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `MccCode` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, mccCodePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into MccCode")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.MCCCode,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for MccCode")
	}

CacheNoHooks:
	if !cached {
		mccCodeInsertCacheMut.Lock()
		mccCodeInsertCache[key] = cache
		mccCodeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the MccCode.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *MccCode) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	mccCodeUpdateCacheMut.RLock()
	cache, cached := mccCodeUpdateCache[key]
	mccCodeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			mccCodeColumns,
			mccCodePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update MccCode, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `MccCode` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, mccCodePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(mccCodeType, mccCodeMapping, append(wl, mccCodePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update MccCode row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for MccCode")
	}

	if !cached {
		mccCodeUpdateCacheMut.Lock()
		mccCodeUpdateCache[key] = cache
		mccCodeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q mccCodeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for MccCode")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for MccCode")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o MccCodeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mccCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `MccCode` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mccCodePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in mccCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all mccCode")
	}
	return rowsAff, nil
}

var mySQLMccCodeUniqueColumns = []string{
	"mcc_code",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *MccCode) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no MccCode provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(mccCodeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLMccCodeUniqueColumns, o)

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

	mccCodeUpsertCacheMut.RLock()
	cache, cached := mccCodeUpsertCache[key]
	mccCodeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			mccCodeColumns,
			mccCodeColumnsWithDefault,
			mccCodeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			mccCodeColumns,
			mccCodePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert MccCode, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "MccCode", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `MccCode` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(mccCodeType, mccCodeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(mccCodeType, mccCodeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for MccCode")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(mccCodeType, mccCodeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for MccCode")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for MccCode")
	}

CacheNoHooks:
	if !cached {
		mccCodeUpsertCacheMut.Lock()
		mccCodeUpsertCache[key] = cache
		mccCodeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single MccCode record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *MccCode) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MccCode provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), mccCodePrimaryKeyMapping)
	sql := "DELETE FROM `MccCode` WHERE `mcc_code`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from MccCode")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for MccCode")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q mccCodeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no mccCodeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from MccCode")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for MccCode")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o MccCodeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no MccCode slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(mccCodeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mccCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `MccCode` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mccCodePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from mccCode slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for MccCode")
	}

	if len(mccCodeAfterDeleteHooks) != 0 {
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
func (o *MccCode) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindMccCode(ctx, exec, o.MCCCode)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *MccCodeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := MccCodeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), mccCodePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `MccCode`.* FROM `MccCode` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, mccCodePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in MccCodeSlice")
	}

	*o = slice

	return nil
}

// MccCodeExists checks if the MccCode row exists.
func MccCodeExists(ctx context.Context, exec boil.ContextExecutor, mCCCode int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `MccCode` where `mcc_code`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, mCCCode)
	}

	row := exec.QueryRowContext(ctx, sql, mCCCode)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if MccCode exists")
	}

	return exists, nil
}
