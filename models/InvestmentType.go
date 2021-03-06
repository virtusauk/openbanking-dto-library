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

// InvestmentType is an object representing the database table.
type InvestmentType struct {
	InvestmenttypeID int         `boil:"investmenttype_id" json:"investmenttype_id" toml:"investmenttype_id" yaml:"investmenttype_id"`
	Investmenttype   string      `boil:"investmenttype" json:"investmenttype" toml:"investmenttype" yaml:"investmenttype"`
	Description      string      `boil:"description" json:"description" toml:"description" yaml:"description"`
	Active           null.String `boil:"active" json:"active,omitempty" toml:"active" yaml:"active,omitempty"`
	BankID           int         `boil:"bank_id" json:"bank_id" toml:"bank_id" yaml:"bank_id"`
	MakerDate        time.Time   `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate      null.Time   `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID          string      `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID        null.String `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy       null.String `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate     null.Time   `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *investmentTypeR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L investmentTypeL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InvestmentTypeColumns = struct {
	InvestmenttypeID string
	Investmenttype   string
	Description      string
	Active           string
	BankID           string
	MakerDate        string
	CheckerDate      string
	MakerID          string
	CheckerID        string
	ModifiedBy       string
	ModifiedDate     string
}{
	InvestmenttypeID: "investmenttype_id",
	Investmenttype:   "investmenttype",
	Description:      "description",
	Active:           "active",
	BankID:           "bank_id",
	MakerDate:        "maker_date",
	CheckerDate:      "checker_date",
	MakerID:          "maker_id",
	CheckerID:        "checker_id",
	ModifiedBy:       "modified_by",
	ModifiedDate:     "modified_date",
}

// Generated where

var InvestmentTypeWhere = struct {
	InvestmenttypeID whereHelperint
	Investmenttype   whereHelperstring
	Description      whereHelperstring
	Active           whereHelpernull_String
	BankID           whereHelperint
	MakerDate        whereHelpertime_Time
	CheckerDate      whereHelpernull_Time
	MakerID          whereHelperstring
	CheckerID        whereHelpernull_String
	ModifiedBy       whereHelpernull_String
	ModifiedDate     whereHelpernull_Time
}{
	InvestmenttypeID: whereHelperint{field: `investmenttype_id`},
	Investmenttype:   whereHelperstring{field: `investmenttype`},
	Description:      whereHelperstring{field: `description`},
	Active:           whereHelpernull_String{field: `active`},
	BankID:           whereHelperint{field: `bank_id`},
	MakerDate:        whereHelpertime_Time{field: `maker_date`},
	CheckerDate:      whereHelpernull_Time{field: `checker_date`},
	MakerID:          whereHelperstring{field: `maker_id`},
	CheckerID:        whereHelpernull_String{field: `checker_id`},
	ModifiedBy:       whereHelpernull_String{field: `modified_by`},
	ModifiedDate:     whereHelpernull_Time{field: `modified_date`},
}

// InvestmentTypeRels is where relationship names are stored.
var InvestmentTypeRels = struct {
	InvestmentTypePortfolios string
}{
	InvestmentTypePortfolios: "InvestmentTypePortfolios",
}

// investmentTypeR is where relationships are stored.
type investmentTypeR struct {
	InvestmentTypePortfolios PortfolioSlice
}

// NewStruct creates a new relationship struct
func (*investmentTypeR) NewStruct() *investmentTypeR {
	return &investmentTypeR{}
}

// investmentTypeL is where Load methods for each relationship are stored.
type investmentTypeL struct{}

var (
	investmentTypeColumns               = []string{"investmenttype_id", "investmenttype", "description", "active", "bank_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	investmentTypeColumnsWithoutDefault = []string{"investmenttype_id", "investmenttype", "description", "active", "bank_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	investmentTypeColumnsWithDefault    = []string{}
	investmentTypePrimaryKeyColumns     = []string{"investmenttype_id"}
)

type (
	// InvestmentTypeSlice is an alias for a slice of pointers to InvestmentType.
	// This should generally be used opposed to []InvestmentType.
	InvestmentTypeSlice []*InvestmentType
	// InvestmentTypeHook is the signature for custom InvestmentType hook methods
	InvestmentTypeHook func(context.Context, boil.ContextExecutor, *InvestmentType) error

	investmentTypeQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	investmentTypeType                 = reflect.TypeOf(&InvestmentType{})
	investmentTypeMapping              = queries.MakeStructMapping(investmentTypeType)
	investmentTypePrimaryKeyMapping, _ = queries.BindMapping(investmentTypeType, investmentTypeMapping, investmentTypePrimaryKeyColumns)
	investmentTypeInsertCacheMut       sync.RWMutex
	investmentTypeInsertCache          = make(map[string]insertCache)
	investmentTypeUpdateCacheMut       sync.RWMutex
	investmentTypeUpdateCache          = make(map[string]updateCache)
	investmentTypeUpsertCacheMut       sync.RWMutex
	investmentTypeUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var investmentTypeBeforeInsertHooks []InvestmentTypeHook
var investmentTypeBeforeUpdateHooks []InvestmentTypeHook
var investmentTypeBeforeDeleteHooks []InvestmentTypeHook
var investmentTypeBeforeUpsertHooks []InvestmentTypeHook

var investmentTypeAfterInsertHooks []InvestmentTypeHook
var investmentTypeAfterSelectHooks []InvestmentTypeHook
var investmentTypeAfterUpdateHooks []InvestmentTypeHook
var investmentTypeAfterDeleteHooks []InvestmentTypeHook
var investmentTypeAfterUpsertHooks []InvestmentTypeHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *InvestmentType) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range investmentTypeBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *InvestmentType) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range investmentTypeBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *InvestmentType) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range investmentTypeBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *InvestmentType) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range investmentTypeBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *InvestmentType) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range investmentTypeAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *InvestmentType) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range investmentTypeAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *InvestmentType) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range investmentTypeAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *InvestmentType) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range investmentTypeAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *InvestmentType) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range investmentTypeAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInvestmentTypeHook registers your hook function for all future operations.
func AddInvestmentTypeHook(hookPoint boil.HookPoint, investmentTypeHook InvestmentTypeHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		investmentTypeBeforeInsertHooks = append(investmentTypeBeforeInsertHooks, investmentTypeHook)
	case boil.BeforeUpdateHook:
		investmentTypeBeforeUpdateHooks = append(investmentTypeBeforeUpdateHooks, investmentTypeHook)
	case boil.BeforeDeleteHook:
		investmentTypeBeforeDeleteHooks = append(investmentTypeBeforeDeleteHooks, investmentTypeHook)
	case boil.BeforeUpsertHook:
		investmentTypeBeforeUpsertHooks = append(investmentTypeBeforeUpsertHooks, investmentTypeHook)
	case boil.AfterInsertHook:
		investmentTypeAfterInsertHooks = append(investmentTypeAfterInsertHooks, investmentTypeHook)
	case boil.AfterSelectHook:
		investmentTypeAfterSelectHooks = append(investmentTypeAfterSelectHooks, investmentTypeHook)
	case boil.AfterUpdateHook:
		investmentTypeAfterUpdateHooks = append(investmentTypeAfterUpdateHooks, investmentTypeHook)
	case boil.AfterDeleteHook:
		investmentTypeAfterDeleteHooks = append(investmentTypeAfterDeleteHooks, investmentTypeHook)
	case boil.AfterUpsertHook:
		investmentTypeAfterUpsertHooks = append(investmentTypeAfterUpsertHooks, investmentTypeHook)
	}
}

// One returns a single investmentType record from the query.
func (q investmentTypeQuery) One(ctx context.Context, exec boil.ContextExecutor) (*InvestmentType, error) {
	o := &InvestmentType{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for InvestmentType")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all InvestmentType records from the query.
func (q investmentTypeQuery) All(ctx context.Context, exec boil.ContextExecutor) (InvestmentTypeSlice, error) {
	var o []*InvestmentType

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to InvestmentType slice")
	}

	if len(investmentTypeAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all InvestmentType records in the query.
func (q investmentTypeQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count InvestmentType rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q investmentTypeQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if InvestmentType exists")
	}

	return count > 0, nil
}

// InvestmentTypePortfolios retrieves all the Portfolio's Portfolios with an executor via investment_type_id column.
func (o *InvestmentType) InvestmentTypePortfolios(mods ...qm.QueryMod) portfolioQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`Portfolio`.`investment_type_id`=?", o.InvestmenttypeID),
	)

	query := Portfolios(queryMods...)
	queries.SetFrom(query.Query, "`Portfolio`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`Portfolio`.*"})
	}

	return query
}

// LoadInvestmentTypePortfolios allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (investmentTypeL) LoadInvestmentTypePortfolios(ctx context.Context, e boil.ContextExecutor, singular bool, maybeInvestmentType interface{}, mods queries.Applicator) error {
	var slice []*InvestmentType
	var object *InvestmentType

	if singular {
		object = maybeInvestmentType.(*InvestmentType)
	} else {
		slice = *maybeInvestmentType.(*[]*InvestmentType)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &investmentTypeR{}
		}
		args = append(args, object.InvestmenttypeID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &investmentTypeR{}
			}

			for _, a := range args {
				if a == obj.InvestmenttypeID {
					continue Outer
				}
			}

			args = append(args, obj.InvestmenttypeID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`Portfolio`), qm.WhereIn(`investment_type_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Portfolio")
	}

	var resultSlice []*Portfolio
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Portfolio")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on Portfolio")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Portfolio")
	}

	if len(portfolioAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.InvestmentTypePortfolios = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &portfolioR{}
			}
			foreign.R.InvestmentType = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.InvestmenttypeID == foreign.InvestmentTypeID {
				local.R.InvestmentTypePortfolios = append(local.R.InvestmentTypePortfolios, foreign)
				if foreign.R == nil {
					foreign.R = &portfolioR{}
				}
				foreign.R.InvestmentType = local
				break
			}
		}
	}

	return nil
}

// AddInvestmentTypePortfolios adds the given related objects to the existing relationships
// of the InvestmentType, optionally inserting them as new records.
// Appends related to o.R.InvestmentTypePortfolios.
// Sets related.R.InvestmentType appropriately.
func (o *InvestmentType) AddInvestmentTypePortfolios(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*Portfolio) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.InvestmentTypeID = o.InvestmenttypeID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `Portfolio` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"investment_type_id"}),
				strmangle.WhereClause("`", "`", 0, portfolioPrimaryKeyColumns),
			)
			values := []interface{}{o.InvestmenttypeID, rel.PortfolioID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.InvestmentTypeID = o.InvestmenttypeID
		}
	}

	if o.R == nil {
		o.R = &investmentTypeR{
			InvestmentTypePortfolios: related,
		}
	} else {
		o.R.InvestmentTypePortfolios = append(o.R.InvestmentTypePortfolios, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &portfolioR{
				InvestmentType: o,
			}
		} else {
			rel.R.InvestmentType = o
		}
	}
	return nil
}

// InvestmentTypes retrieves all the records using an executor.
func InvestmentTypes(mods ...qm.QueryMod) investmentTypeQuery {
	mods = append(mods, qm.From("`InvestmentType`"))
	return investmentTypeQuery{NewQuery(mods...)}
}

// FindInvestmentType retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInvestmentType(ctx context.Context, exec boil.ContextExecutor, investmenttypeID int, selectCols ...string) (*InvestmentType, error) {
	investmentTypeObj := &InvestmentType{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `InvestmentType` where `investmenttype_id`=?", sel,
	)

	q := queries.Raw(query, investmenttypeID)

	err := q.Bind(ctx, exec, investmentTypeObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from InvestmentType")
	}

	return investmentTypeObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *InvestmentType) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no InvestmentType provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(investmentTypeColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	investmentTypeInsertCacheMut.RLock()
	cache, cached := investmentTypeInsertCache[key]
	investmentTypeInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			investmentTypeColumns,
			investmentTypeColumnsWithDefault,
			investmentTypeColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(investmentTypeType, investmentTypeMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(investmentTypeType, investmentTypeMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `InvestmentType` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `InvestmentType` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `InvestmentType` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, investmentTypePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into InvestmentType")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.InvestmenttypeID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for InvestmentType")
	}

CacheNoHooks:
	if !cached {
		investmentTypeInsertCacheMut.Lock()
		investmentTypeInsertCache[key] = cache
		investmentTypeInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the InvestmentType.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *InvestmentType) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	investmentTypeUpdateCacheMut.RLock()
	cache, cached := investmentTypeUpdateCache[key]
	investmentTypeUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			investmentTypeColumns,
			investmentTypePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update InvestmentType, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `InvestmentType` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, investmentTypePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(investmentTypeType, investmentTypeMapping, append(wl, investmentTypePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update InvestmentType row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for InvestmentType")
	}

	if !cached {
		investmentTypeUpdateCacheMut.Lock()
		investmentTypeUpdateCache[key] = cache
		investmentTypeUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q investmentTypeQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for InvestmentType")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for InvestmentType")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InvestmentTypeSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), investmentTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `InvestmentType` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, investmentTypePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in investmentType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all investmentType")
	}
	return rowsAff, nil
}

var mySQLInvestmentTypeUniqueColumns = []string{
	"investmenttype_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *InvestmentType) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no InvestmentType provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(investmentTypeColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLInvestmentTypeUniqueColumns, o)

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

	investmentTypeUpsertCacheMut.RLock()
	cache, cached := investmentTypeUpsertCache[key]
	investmentTypeUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			investmentTypeColumns,
			investmentTypeColumnsWithDefault,
			investmentTypeColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			investmentTypeColumns,
			investmentTypePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert InvestmentType, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "InvestmentType", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `InvestmentType` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(investmentTypeType, investmentTypeMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(investmentTypeType, investmentTypeMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for InvestmentType")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(investmentTypeType, investmentTypeMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for InvestmentType")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for InvestmentType")
	}

CacheNoHooks:
	if !cached {
		investmentTypeUpsertCacheMut.Lock()
		investmentTypeUpsertCache[key] = cache
		investmentTypeUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single InvestmentType record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *InvestmentType) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no InvestmentType provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), investmentTypePrimaryKeyMapping)
	sql := "DELETE FROM `InvestmentType` WHERE `investmenttype_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from InvestmentType")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for InvestmentType")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q investmentTypeQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no investmentTypeQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from InvestmentType")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for InvestmentType")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InvestmentTypeSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no InvestmentType slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(investmentTypeBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), investmentTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `InvestmentType` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, investmentTypePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from investmentType slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for InvestmentType")
	}

	if len(investmentTypeAfterDeleteHooks) != 0 {
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
func (o *InvestmentType) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInvestmentType(ctx, exec, o.InvestmenttypeID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InvestmentTypeSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InvestmentTypeSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), investmentTypePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `InvestmentType`.* FROM `InvestmentType` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, investmentTypePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InvestmentTypeSlice")
	}

	*o = slice

	return nil
}

// InvestmentTypeExists checks if the InvestmentType row exists.
func InvestmentTypeExists(ctx context.Context, exec boil.ContextExecutor, investmenttypeID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `InvestmentType` where `investmenttype_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, investmenttypeID)
	}

	row := exec.QueryRowContext(ctx, sql, investmenttypeID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if InvestmentType exists")
	}

	return exists, nil
}
