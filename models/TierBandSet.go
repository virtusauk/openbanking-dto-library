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

// TierBandSet is an object representing the database table.
type TierBandSet struct {
	TierBandSetID     int         `boil:"tier_band_set_id" json:"tier_band_set_id" toml:"tier_band_set_id" yaml:"tier_band_set_id"`
	ProductID         int         `boil:"product_id" json:"product_id" toml:"product_id" yaml:"product_id"`
	TierBandMethod    null.String `boil:"tier_band_method" json:"tier_band_method,omitempty" toml:"tier_band_method" yaml:"tier_band_method,omitempty"`
	CalculationMethod null.String `boil:"calculation_method" json:"calculation_method,omitempty" toml:"calculation_method" yaml:"calculation_method,omitempty"`
	Destination       null.String `boil:"destination" json:"destination,omitempty" toml:"destination" yaml:"destination,omitempty"`
	Notes             null.String `boil:"notes" json:"notes,omitempty" toml:"notes" yaml:"notes,omitempty"`

	R *tierBandSetR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tierBandSetL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TierBandSetColumns = struct {
	TierBandSetID     string
	ProductID         string
	TierBandMethod    string
	CalculationMethod string
	Destination       string
	Notes             string
}{
	TierBandSetID:     "tier_band_set_id",
	ProductID:         "product_id",
	TierBandMethod:    "tier_band_method",
	CalculationMethod: "calculation_method",
	Destination:       "destination",
	Notes:             "notes",
}

// Generated where

var TierBandSetWhere = struct {
	TierBandSetID     whereHelperint
	ProductID         whereHelperint
	TierBandMethod    whereHelpernull_String
	CalculationMethod whereHelpernull_String
	Destination       whereHelpernull_String
	Notes             whereHelpernull_String
}{
	TierBandSetID:     whereHelperint{field: `tier_band_set_id`},
	ProductID:         whereHelperint{field: `product_id`},
	TierBandMethod:    whereHelpernull_String{field: `tier_band_method`},
	CalculationMethod: whereHelpernull_String{field: `calculation_method`},
	Destination:       whereHelpernull_String{field: `destination`},
	Notes:             whereHelpernull_String{field: `notes`},
}

// TierBandSetRels is where relationship names are stored.
var TierBandSetRels = struct {
	Product              string
	TierBandSetTierBands string
}{
	Product:              "Product",
	TierBandSetTierBands: "TierBandSetTierBands",
}

// tierBandSetR is where relationships are stored.
type tierBandSetR struct {
	Product              *Product
	TierBandSetTierBands TierBandSlice
}

// NewStruct creates a new relationship struct
func (*tierBandSetR) NewStruct() *tierBandSetR {
	return &tierBandSetR{}
}

// tierBandSetL is where Load methods for each relationship are stored.
type tierBandSetL struct{}

var (
	tierBandSetColumns               = []string{"tier_band_set_id", "product_id", "tier_band_method", "calculation_method", "destination", "notes"}
	tierBandSetColumnsWithoutDefault = []string{"tier_band_set_id", "product_id", "calculation_method", "destination", "notes"}
	tierBandSetColumnsWithDefault    = []string{"tier_band_method"}
	tierBandSetPrimaryKeyColumns     = []string{"tier_band_set_id"}
)

type (
	// TierBandSetSlice is an alias for a slice of pointers to TierBandSet.
	// This should generally be used opposed to []TierBandSet.
	TierBandSetSlice []*TierBandSet
	// TierBandSetHook is the signature for custom TierBandSet hook methods
	TierBandSetHook func(context.Context, boil.ContextExecutor, *TierBandSet) error

	tierBandSetQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tierBandSetType                 = reflect.TypeOf(&TierBandSet{})
	tierBandSetMapping              = queries.MakeStructMapping(tierBandSetType)
	tierBandSetPrimaryKeyMapping, _ = queries.BindMapping(tierBandSetType, tierBandSetMapping, tierBandSetPrimaryKeyColumns)
	tierBandSetInsertCacheMut       sync.RWMutex
	tierBandSetInsertCache          = make(map[string]insertCache)
	tierBandSetUpdateCacheMut       sync.RWMutex
	tierBandSetUpdateCache          = make(map[string]updateCache)
	tierBandSetUpsertCacheMut       sync.RWMutex
	tierBandSetUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tierBandSetBeforeInsertHooks []TierBandSetHook
var tierBandSetBeforeUpdateHooks []TierBandSetHook
var tierBandSetBeforeDeleteHooks []TierBandSetHook
var tierBandSetBeforeUpsertHooks []TierBandSetHook

var tierBandSetAfterInsertHooks []TierBandSetHook
var tierBandSetAfterSelectHooks []TierBandSetHook
var tierBandSetAfterUpdateHooks []TierBandSetHook
var tierBandSetAfterDeleteHooks []TierBandSetHook
var tierBandSetAfterUpsertHooks []TierBandSetHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TierBandSet) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tierBandSetBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TierBandSet) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tierBandSetBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TierBandSet) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tierBandSetBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TierBandSet) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tierBandSetBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TierBandSet) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tierBandSetAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TierBandSet) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tierBandSetAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TierBandSet) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tierBandSetAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TierBandSet) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tierBandSetAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TierBandSet) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tierBandSetAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTierBandSetHook registers your hook function for all future operations.
func AddTierBandSetHook(hookPoint boil.HookPoint, tierBandSetHook TierBandSetHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		tierBandSetBeforeInsertHooks = append(tierBandSetBeforeInsertHooks, tierBandSetHook)
	case boil.BeforeUpdateHook:
		tierBandSetBeforeUpdateHooks = append(tierBandSetBeforeUpdateHooks, tierBandSetHook)
	case boil.BeforeDeleteHook:
		tierBandSetBeforeDeleteHooks = append(tierBandSetBeforeDeleteHooks, tierBandSetHook)
	case boil.BeforeUpsertHook:
		tierBandSetBeforeUpsertHooks = append(tierBandSetBeforeUpsertHooks, tierBandSetHook)
	case boil.AfterInsertHook:
		tierBandSetAfterInsertHooks = append(tierBandSetAfterInsertHooks, tierBandSetHook)
	case boil.AfterSelectHook:
		tierBandSetAfterSelectHooks = append(tierBandSetAfterSelectHooks, tierBandSetHook)
	case boil.AfterUpdateHook:
		tierBandSetAfterUpdateHooks = append(tierBandSetAfterUpdateHooks, tierBandSetHook)
	case boil.AfterDeleteHook:
		tierBandSetAfterDeleteHooks = append(tierBandSetAfterDeleteHooks, tierBandSetHook)
	case boil.AfterUpsertHook:
		tierBandSetAfterUpsertHooks = append(tierBandSetAfterUpsertHooks, tierBandSetHook)
	}
}

// One returns a single tierBandSet record from the query.
func (q tierBandSetQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TierBandSet, error) {
	o := &TierBandSet{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for TierBandSet")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TierBandSet records from the query.
func (q tierBandSetQuery) All(ctx context.Context, exec boil.ContextExecutor) (TierBandSetSlice, error) {
	var o []*TierBandSet

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TierBandSet slice")
	}

	if len(tierBandSetAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TierBandSet records in the query.
func (q tierBandSetQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count TierBandSet rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tierBandSetQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if TierBandSet exists")
	}

	return count > 0, nil
}

// Product pointed to by the foreign key.
func (o *TierBandSet) Product(mods ...qm.QueryMod) productQuery {
	queryMods := []qm.QueryMod{
		qm.Where("product_id=?", o.ProductID),
	}

	queryMods = append(queryMods, mods...)

	query := Products(queryMods...)
	queries.SetFrom(query.Query, "`Product`")

	return query
}

// TierBandSetTierBands retrieves all the TierBand's TierBands with an executor via tier_band_set_id column.
func (o *TierBandSet) TierBandSetTierBands(mods ...qm.QueryMod) tierBandQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`TierBand`.`tier_band_set_id`=?", o.TierBandSetID),
	)

	query := TierBands(queryMods...)
	queries.SetFrom(query.Query, "`TierBand`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`TierBand`.*"})
	}

	return query
}

// LoadProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (tierBandSetL) LoadProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTierBandSet interface{}, mods queries.Applicator) error {
	var slice []*TierBandSet
	var object *TierBandSet

	if singular {
		object = maybeTierBandSet.(*TierBandSet)
	} else {
		slice = *maybeTierBandSet.(*[]*TierBandSet)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tierBandSetR{}
		}
		args = append(args, object.ProductID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tierBandSetR{}
			}

			for _, a := range args {
				if a == obj.ProductID {
					continue Outer
				}
			}

			args = append(args, obj.ProductID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`Product`), qm.WhereIn(`product_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Product")
	}

	var resultSlice []*Product
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Product")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Product")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Product")
	}

	if len(tierBandSetAfterSelectHooks) != 0 {
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
		object.R.Product = foreign
		if foreign.R == nil {
			foreign.R = &productR{}
		}
		foreign.R.ProductTierBandSets = append(foreign.R.ProductTierBandSets, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.ProductID == foreign.ProductID {
				local.R.Product = foreign
				if foreign.R == nil {
					foreign.R = &productR{}
				}
				foreign.R.ProductTierBandSets = append(foreign.R.ProductTierBandSets, local)
				break
			}
		}
	}

	return nil
}

// LoadTierBandSetTierBands allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (tierBandSetL) LoadTierBandSetTierBands(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTierBandSet interface{}, mods queries.Applicator) error {
	var slice []*TierBandSet
	var object *TierBandSet

	if singular {
		object = maybeTierBandSet.(*TierBandSet)
	} else {
		slice = *maybeTierBandSet.(*[]*TierBandSet)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tierBandSetR{}
		}
		args = append(args, object.TierBandSetID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tierBandSetR{}
			}

			for _, a := range args {
				if a == obj.TierBandSetID {
					continue Outer
				}
			}

			args = append(args, obj.TierBandSetID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`TierBand`), qm.WhereIn(`tier_band_set_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load TierBand")
	}

	var resultSlice []*TierBand
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice TierBand")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on TierBand")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for TierBand")
	}

	if len(tierBandAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.TierBandSetTierBands = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &tierBandR{}
			}
			foreign.R.TierBandSet = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.TierBandSetID == foreign.TierBandSetID {
				local.R.TierBandSetTierBands = append(local.R.TierBandSetTierBands, foreign)
				if foreign.R == nil {
					foreign.R = &tierBandR{}
				}
				foreign.R.TierBandSet = local
				break
			}
		}
	}

	return nil
}

// SetProduct of the tierBandSet to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.ProductTierBandSets.
func (o *TierBandSet) SetProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Product) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `TierBandSet` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"product_id"}),
		strmangle.WhereClause("`", "`", 0, tierBandSetPrimaryKeyColumns),
	)
	values := []interface{}{related.ProductID, o.TierBandSetID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.ProductID = related.ProductID
	if o.R == nil {
		o.R = &tierBandSetR{
			Product: related,
		}
	} else {
		o.R.Product = related
	}

	if related.R == nil {
		related.R = &productR{
			ProductTierBandSets: TierBandSetSlice{o},
		}
	} else {
		related.R.ProductTierBandSets = append(related.R.ProductTierBandSets, o)
	}

	return nil
}

// AddTierBandSetTierBands adds the given related objects to the existing relationships
// of the TierBandSet, optionally inserting them as new records.
// Appends related to o.R.TierBandSetTierBands.
// Sets related.R.TierBandSet appropriately.
func (o *TierBandSet) AddTierBandSetTierBands(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*TierBand) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.TierBandSetID = o.TierBandSetID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `TierBand` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"tier_band_set_id"}),
				strmangle.WhereClause("`", "`", 0, tierBandPrimaryKeyColumns),
			)
			values := []interface{}{o.TierBandSetID, rel.TierbandID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.TierBandSetID = o.TierBandSetID
		}
	}

	if o.R == nil {
		o.R = &tierBandSetR{
			TierBandSetTierBands: related,
		}
	} else {
		o.R.TierBandSetTierBands = append(o.R.TierBandSetTierBands, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &tierBandR{
				TierBandSet: o,
			}
		} else {
			rel.R.TierBandSet = o
		}
	}
	return nil
}

// TierBandSets retrieves all the records using an executor.
func TierBandSets(mods ...qm.QueryMod) tierBandSetQuery {
	mods = append(mods, qm.From("`TierBandSet`"))
	return tierBandSetQuery{NewQuery(mods...)}
}

// FindTierBandSet retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTierBandSet(ctx context.Context, exec boil.ContextExecutor, tierBandSetID int, selectCols ...string) (*TierBandSet, error) {
	tierBandSetObj := &TierBandSet{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `TierBandSet` where `tier_band_set_id`=?", sel,
	)

	q := queries.Raw(query, tierBandSetID)

	err := q.Bind(ctx, exec, tierBandSetObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from TierBandSet")
	}

	return tierBandSetObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TierBandSet) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TierBandSet provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tierBandSetColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tierBandSetInsertCacheMut.RLock()
	cache, cached := tierBandSetInsertCache[key]
	tierBandSetInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tierBandSetColumns,
			tierBandSetColumnsWithDefault,
			tierBandSetColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tierBandSetType, tierBandSetMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tierBandSetType, tierBandSetMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `TierBandSet` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `TierBandSet` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `TierBandSet` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tierBandSetPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into TierBandSet")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.TierBandSetID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TierBandSet")
	}

CacheNoHooks:
	if !cached {
		tierBandSetInsertCacheMut.Lock()
		tierBandSetInsertCache[key] = cache
		tierBandSetInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TierBandSet.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TierBandSet) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tierBandSetUpdateCacheMut.RLock()
	cache, cached := tierBandSetUpdateCache[key]
	tierBandSetUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tierBandSetColumns,
			tierBandSetPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update TierBandSet, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `TierBandSet` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tierBandSetPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tierBandSetType, tierBandSetMapping, append(wl, tierBandSetPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update TierBandSet row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for TierBandSet")
	}

	if !cached {
		tierBandSetUpdateCacheMut.Lock()
		tierBandSetUpdateCache[key] = cache
		tierBandSetUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tierBandSetQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for TierBandSet")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for TierBandSet")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TierBandSetSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tierBandSetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `TierBandSet` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tierBandSetPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tierBandSet slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tierBandSet")
	}
	return rowsAff, nil
}

var mySQLTierBandSetUniqueColumns = []string{
	"tier_band_set_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TierBandSet) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TierBandSet provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tierBandSetColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTierBandSetUniqueColumns, o)

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

	tierBandSetUpsertCacheMut.RLock()
	cache, cached := tierBandSetUpsertCache[key]
	tierBandSetUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tierBandSetColumns,
			tierBandSetColumnsWithDefault,
			tierBandSetColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			tierBandSetColumns,
			tierBandSetPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert TierBandSet, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "TierBandSet", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `TierBandSet` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tierBandSetType, tierBandSetMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tierBandSetType, tierBandSetMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for TierBandSet")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tierBandSetType, tierBandSetMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for TierBandSet")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TierBandSet")
	}

CacheNoHooks:
	if !cached {
		tierBandSetUpsertCacheMut.Lock()
		tierBandSetUpsertCache[key] = cache
		tierBandSetUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TierBandSet record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TierBandSet) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TierBandSet provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tierBandSetPrimaryKeyMapping)
	sql := "DELETE FROM `TierBandSet` WHERE `tier_band_set_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from TierBandSet")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for TierBandSet")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tierBandSetQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tierBandSetQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from TierBandSet")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TierBandSet")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TierBandSetSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TierBandSet slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(tierBandSetBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tierBandSetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `TierBandSet` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tierBandSetPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tierBandSet slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TierBandSet")
	}

	if len(tierBandSetAfterDeleteHooks) != 0 {
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
func (o *TierBandSet) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTierBandSet(ctx, exec, o.TierBandSetID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TierBandSetSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TierBandSetSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tierBandSetPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `TierBandSet`.* FROM `TierBandSet` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tierBandSetPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TierBandSetSlice")
	}

	*o = slice

	return nil
}

// TierBandSetExists checks if the TierBandSet row exists.
func TierBandSetExists(ctx context.Context, exec boil.ContextExecutor, tierBandSetID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `TierBandSet` where `tier_band_set_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, tierBandSetID)
	}

	row := exec.QueryRowContext(ctx, sql, tierBandSetID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if TierBandSet exists")
	}

	return exists, nil
}
