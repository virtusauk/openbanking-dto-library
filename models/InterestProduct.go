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

// InterestProduct is an object representing the database table.
type InterestProduct struct {
	InterestProductID        int               `boil:"interest_product_id" json:"interest_product_id" toml:"interest_product_id" yaml:"interest_product_id"`
	ProductID                null.Int          `boil:"product_id" json:"product_id,omitempty" toml:"product_id" yaml:"product_id,omitempty"`
	Type                     null.String       `boil:"type" json:"type,omitempty" toml:"type" yaml:"type,omitempty"`
	DefaultInterestRate      types.NullDecimal `boil:"default_interest_rate" json:"default_interest_rate,omitempty" toml:"default_interest_rate" yaml:"default_interest_rate,omitempty"`
	MinInterestRate          types.NullDecimal `boil:"min_interest_rate" json:"min_interest_rate,omitempty" toml:"min_interest_rate" yaml:"min_interest_rate,omitempty"`
	MaxInterestRate          types.NullDecimal `boil:"max_interest_rate" json:"max_interest_rate,omitempty" toml:"max_interest_rate" yaml:"max_interest_rate,omitempty"`
	InterestRateFloorValue   types.NullDecimal `boil:"interest_rate_floor_value" json:"interest_rate_floor_value,omitempty" toml:"interest_rate_floor_value" yaml:"interest_rate_floor_value,omitempty"`
	InterestRateCeilingValue types.NullDecimal `boil:"interest_rate_ceiling_value" json:"interest_rate_ceiling_value,omitempty" toml:"interest_rate_ceiling_value" yaml:"interest_rate_ceiling_value,omitempty"`
	MakerDate                null.Time         `boil:"maker_date" json:"maker_date,omitempty" toml:"maker_date" yaml:"maker_date,omitempty"`
	CheckerDate              null.Time         `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID                  null.Int          `boil:"maker_id" json:"maker_id,omitempty" toml:"maker_id" yaml:"maker_id,omitempty"`
	CheckerID                null.Int          `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy               null.String       `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate             null.Time         `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *interestProductR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L interestProductL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var InterestProductColumns = struct {
	InterestProductID        string
	ProductID                string
	Type                     string
	DefaultInterestRate      string
	MinInterestRate          string
	MaxInterestRate          string
	InterestRateFloorValue   string
	InterestRateCeilingValue string
	MakerDate                string
	CheckerDate              string
	MakerID                  string
	CheckerID                string
	ModifiedBy               string
	ModifiedDate             string
}{
	InterestProductID:        "interest_product_id",
	ProductID:                "product_id",
	Type:                     "type",
	DefaultInterestRate:      "default_interest_rate",
	MinInterestRate:          "min_interest_rate",
	MaxInterestRate:          "max_interest_rate",
	InterestRateFloorValue:   "interest_rate_floor_value",
	InterestRateCeilingValue: "interest_rate_ceiling_value",
	MakerDate:                "maker_date",
	CheckerDate:              "checker_date",
	MakerID:                  "maker_id",
	CheckerID:                "checker_id",
	ModifiedBy:               "modified_by",
	ModifiedDate:             "modified_date",
}

// Generated where

var InterestProductWhere = struct {
	InterestProductID        whereHelperint
	ProductID                whereHelpernull_Int
	Type                     whereHelpernull_String
	DefaultInterestRate      whereHelpertypes_NullDecimal
	MinInterestRate          whereHelpertypes_NullDecimal
	MaxInterestRate          whereHelpertypes_NullDecimal
	InterestRateFloorValue   whereHelpertypes_NullDecimal
	InterestRateCeilingValue whereHelpertypes_NullDecimal
	MakerDate                whereHelpernull_Time
	CheckerDate              whereHelpernull_Time
	MakerID                  whereHelpernull_Int
	CheckerID                whereHelpernull_Int
	ModifiedBy               whereHelpernull_String
	ModifiedDate             whereHelpernull_Time
}{
	InterestProductID:        whereHelperint{field: `interest_product_id`},
	ProductID:                whereHelpernull_Int{field: `product_id`},
	Type:                     whereHelpernull_String{field: `type`},
	DefaultInterestRate:      whereHelpertypes_NullDecimal{field: `default_interest_rate`},
	MinInterestRate:          whereHelpertypes_NullDecimal{field: `min_interest_rate`},
	MaxInterestRate:          whereHelpertypes_NullDecimal{field: `max_interest_rate`},
	InterestRateFloorValue:   whereHelpertypes_NullDecimal{field: `interest_rate_floor_value`},
	InterestRateCeilingValue: whereHelpertypes_NullDecimal{field: `interest_rate_ceiling_value`},
	MakerDate:                whereHelpernull_Time{field: `maker_date`},
	CheckerDate:              whereHelpernull_Time{field: `checker_date`},
	MakerID:                  whereHelpernull_Int{field: `maker_id`},
	CheckerID:                whereHelpernull_Int{field: `checker_id`},
	ModifiedBy:               whereHelpernull_String{field: `modified_by`},
	ModifiedDate:             whereHelpernull_Time{field: `modified_date`},
}

// InterestProductRels is where relationship names are stored.
var InterestProductRels = struct {
	Product string
}{
	Product: "Product",
}

// interestProductR is where relationships are stored.
type interestProductR struct {
	Product *Product
}

// NewStruct creates a new relationship struct
func (*interestProductR) NewStruct() *interestProductR {
	return &interestProductR{}
}

// interestProductL is where Load methods for each relationship are stored.
type interestProductL struct{}

var (
	interestProductColumns               = []string{"interest_product_id", "product_id", "type", "default_interest_rate", "min_interest_rate", "max_interest_rate", "interest_rate_floor_value", "interest_rate_ceiling_value", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	interestProductColumnsWithoutDefault = []string{"interest_product_id", "product_id", "type", "default_interest_rate", "min_interest_rate", "max_interest_rate", "interest_rate_floor_value", "interest_rate_ceiling_value", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	interestProductColumnsWithDefault    = []string{}
	interestProductPrimaryKeyColumns     = []string{"interest_product_id"}
)

type (
	// InterestProductSlice is an alias for a slice of pointers to InterestProduct.
	// This should generally be used opposed to []InterestProduct.
	InterestProductSlice []*InterestProduct
	// InterestProductHook is the signature for custom InterestProduct hook methods
	InterestProductHook func(context.Context, boil.ContextExecutor, *InterestProduct) error

	interestProductQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	interestProductType                 = reflect.TypeOf(&InterestProduct{})
	interestProductMapping              = queries.MakeStructMapping(interestProductType)
	interestProductPrimaryKeyMapping, _ = queries.BindMapping(interestProductType, interestProductMapping, interestProductPrimaryKeyColumns)
	interestProductInsertCacheMut       sync.RWMutex
	interestProductInsertCache          = make(map[string]insertCache)
	interestProductUpdateCacheMut       sync.RWMutex
	interestProductUpdateCache          = make(map[string]updateCache)
	interestProductUpsertCacheMut       sync.RWMutex
	interestProductUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var interestProductBeforeInsertHooks []InterestProductHook
var interestProductBeforeUpdateHooks []InterestProductHook
var interestProductBeforeDeleteHooks []InterestProductHook
var interestProductBeforeUpsertHooks []InterestProductHook

var interestProductAfterInsertHooks []InterestProductHook
var interestProductAfterSelectHooks []InterestProductHook
var interestProductAfterUpdateHooks []InterestProductHook
var interestProductAfterDeleteHooks []InterestProductHook
var interestProductAfterUpsertHooks []InterestProductHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *InterestProduct) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range interestProductBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *InterestProduct) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range interestProductBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *InterestProduct) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range interestProductBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *InterestProduct) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range interestProductBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *InterestProduct) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range interestProductAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *InterestProduct) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range interestProductAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *InterestProduct) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range interestProductAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *InterestProduct) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range interestProductAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *InterestProduct) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range interestProductAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddInterestProductHook registers your hook function for all future operations.
func AddInterestProductHook(hookPoint boil.HookPoint, interestProductHook InterestProductHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		interestProductBeforeInsertHooks = append(interestProductBeforeInsertHooks, interestProductHook)
	case boil.BeforeUpdateHook:
		interestProductBeforeUpdateHooks = append(interestProductBeforeUpdateHooks, interestProductHook)
	case boil.BeforeDeleteHook:
		interestProductBeforeDeleteHooks = append(interestProductBeforeDeleteHooks, interestProductHook)
	case boil.BeforeUpsertHook:
		interestProductBeforeUpsertHooks = append(interestProductBeforeUpsertHooks, interestProductHook)
	case boil.AfterInsertHook:
		interestProductAfterInsertHooks = append(interestProductAfterInsertHooks, interestProductHook)
	case boil.AfterSelectHook:
		interestProductAfterSelectHooks = append(interestProductAfterSelectHooks, interestProductHook)
	case boil.AfterUpdateHook:
		interestProductAfterUpdateHooks = append(interestProductAfterUpdateHooks, interestProductHook)
	case boil.AfterDeleteHook:
		interestProductAfterDeleteHooks = append(interestProductAfterDeleteHooks, interestProductHook)
	case boil.AfterUpsertHook:
		interestProductAfterUpsertHooks = append(interestProductAfterUpsertHooks, interestProductHook)
	}
}

// One returns a single interestProduct record from the query.
func (q interestProductQuery) One(ctx context.Context, exec boil.ContextExecutor) (*InterestProduct, error) {
	o := &InterestProduct{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for InterestProduct")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all InterestProduct records from the query.
func (q interestProductQuery) All(ctx context.Context, exec boil.ContextExecutor) (InterestProductSlice, error) {
	var o []*InterestProduct

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to InterestProduct slice")
	}

	if len(interestProductAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all InterestProduct records in the query.
func (q interestProductQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count InterestProduct rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q interestProductQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if InterestProduct exists")
	}

	return count > 0, nil
}

// Product pointed to by the foreign key.
func (o *InterestProduct) Product(mods ...qm.QueryMod) productQuery {
	queryMods := []qm.QueryMod{
		qm.Where("product_id=?", o.ProductID),
	}

	queryMods = append(queryMods, mods...)

	query := Products(queryMods...)
	queries.SetFrom(query.Query, "`Product`")

	return query
}

// LoadProduct allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (interestProductL) LoadProduct(ctx context.Context, e boil.ContextExecutor, singular bool, maybeInterestProduct interface{}, mods queries.Applicator) error {
	var slice []*InterestProduct
	var object *InterestProduct

	if singular {
		object = maybeInterestProduct.(*InterestProduct)
	} else {
		slice = *maybeInterestProduct.(*[]*InterestProduct)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &interestProductR{}
		}
		if !queries.IsNil(object.ProductID) {
			args = append(args, object.ProductID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &interestProductR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ProductID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ProductID) {
				args = append(args, obj.ProductID)
			}

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

	if len(interestProductAfterSelectHooks) != 0 {
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
		foreign.R.ProductInterestProducts = append(foreign.R.ProductInterestProducts, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ProductID, foreign.ProductID) {
				local.R.Product = foreign
				if foreign.R == nil {
					foreign.R = &productR{}
				}
				foreign.R.ProductInterestProducts = append(foreign.R.ProductInterestProducts, local)
				break
			}
		}
	}

	return nil
}

// SetProduct of the interestProduct to the related item.
// Sets o.R.Product to related.
// Adds o to related.R.ProductInterestProducts.
func (o *InterestProduct) SetProduct(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Product) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `InterestProduct` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"product_id"}),
		strmangle.WhereClause("`", "`", 0, interestProductPrimaryKeyColumns),
	)
	values := []interface{}{related.ProductID, o.InterestProductID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ProductID, related.ProductID)
	if o.R == nil {
		o.R = &interestProductR{
			Product: related,
		}
	} else {
		o.R.Product = related
	}

	if related.R == nil {
		related.R = &productR{
			ProductInterestProducts: InterestProductSlice{o},
		}
	} else {
		related.R.ProductInterestProducts = append(related.R.ProductInterestProducts, o)
	}

	return nil
}

// RemoveProduct relationship.
// Sets o.R.Product to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *InterestProduct) RemoveProduct(ctx context.Context, exec boil.ContextExecutor, related *Product) error {
	var err error

	queries.SetScanner(&o.ProductID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("product_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Product = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ProductInterestProducts {
		if queries.Equal(o.ProductID, ri.ProductID) {
			continue
		}

		ln := len(related.R.ProductInterestProducts)
		if ln > 1 && i < ln-1 {
			related.R.ProductInterestProducts[i] = related.R.ProductInterestProducts[ln-1]
		}
		related.R.ProductInterestProducts = related.R.ProductInterestProducts[:ln-1]
		break
	}
	return nil
}

// InterestProducts retrieves all the records using an executor.
func InterestProducts(mods ...qm.QueryMod) interestProductQuery {
	mods = append(mods, qm.From("`InterestProduct`"))
	return interestProductQuery{NewQuery(mods...)}
}

// FindInterestProduct retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindInterestProduct(ctx context.Context, exec boil.ContextExecutor, interestProductID int, selectCols ...string) (*InterestProduct, error) {
	interestProductObj := &InterestProduct{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `InterestProduct` where `interest_product_id`=?", sel,
	)

	q := queries.Raw(query, interestProductID)

	err := q.Bind(ctx, exec, interestProductObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from InterestProduct")
	}

	return interestProductObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *InterestProduct) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no InterestProduct provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(interestProductColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	interestProductInsertCacheMut.RLock()
	cache, cached := interestProductInsertCache[key]
	interestProductInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			interestProductColumns,
			interestProductColumnsWithDefault,
			interestProductColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(interestProductType, interestProductMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(interestProductType, interestProductMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `InterestProduct` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `InterestProduct` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `InterestProduct` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, interestProductPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into InterestProduct")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.InterestProductID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for InterestProduct")
	}

CacheNoHooks:
	if !cached {
		interestProductInsertCacheMut.Lock()
		interestProductInsertCache[key] = cache
		interestProductInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the InterestProduct.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *InterestProduct) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	interestProductUpdateCacheMut.RLock()
	cache, cached := interestProductUpdateCache[key]
	interestProductUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			interestProductColumns,
			interestProductPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update InterestProduct, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `InterestProduct` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, interestProductPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(interestProductType, interestProductMapping, append(wl, interestProductPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update InterestProduct row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for InterestProduct")
	}

	if !cached {
		interestProductUpdateCacheMut.Lock()
		interestProductUpdateCache[key] = cache
		interestProductUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q interestProductQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for InterestProduct")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for InterestProduct")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o InterestProductSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), interestProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `InterestProduct` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, interestProductPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in interestProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all interestProduct")
	}
	return rowsAff, nil
}

var mySQLInterestProductUniqueColumns = []string{
	"interest_product_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *InterestProduct) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no InterestProduct provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(interestProductColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLInterestProductUniqueColumns, o)

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

	interestProductUpsertCacheMut.RLock()
	cache, cached := interestProductUpsertCache[key]
	interestProductUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			interestProductColumns,
			interestProductColumnsWithDefault,
			interestProductColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			interestProductColumns,
			interestProductPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert InterestProduct, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "InterestProduct", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `InterestProduct` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(interestProductType, interestProductMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(interestProductType, interestProductMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for InterestProduct")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(interestProductType, interestProductMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for InterestProduct")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for InterestProduct")
	}

CacheNoHooks:
	if !cached {
		interestProductUpsertCacheMut.Lock()
		interestProductUpsertCache[key] = cache
		interestProductUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single InterestProduct record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *InterestProduct) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no InterestProduct provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), interestProductPrimaryKeyMapping)
	sql := "DELETE FROM `InterestProduct` WHERE `interest_product_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from InterestProduct")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for InterestProduct")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q interestProductQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no interestProductQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from InterestProduct")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for InterestProduct")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o InterestProductSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no InterestProduct slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(interestProductBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), interestProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `InterestProduct` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, interestProductPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from interestProduct slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for InterestProduct")
	}

	if len(interestProductAfterDeleteHooks) != 0 {
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
func (o *InterestProduct) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindInterestProduct(ctx, exec, o.InterestProductID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *InterestProductSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := InterestProductSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), interestProductPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `InterestProduct`.* FROM `InterestProduct` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, interestProductPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in InterestProductSlice")
	}

	*o = slice

	return nil
}

// InterestProductExists checks if the InterestProduct row exists.
func InterestProductExists(ctx context.Context, exec boil.ContextExecutor, interestProductID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `InterestProduct` where `interest_product_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, interestProductID)
	}

	row := exec.QueryRowContext(ctx, sql, interestProductID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if InterestProduct exists")
	}

	return exists, nil
}