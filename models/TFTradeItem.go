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

// TFTradeItem is an object representing the database table.
type TFTradeItem struct {
	TradeItemID       int               `boil:"trade_item_id" json:"trade_item_id" toml:"trade_item_id" yaml:"trade_item_id"`
	BranchID          null.Int          `boil:"branch_id" json:"branch_id,omitempty" toml:"branch_id" yaml:"branch_id,omitempty"`
	TradeID           int               `boil:"trade_id" json:"trade_id" toml:"trade_id" yaml:"trade_id"`
	NumberOfPackages  null.Int          `boil:"number_of_packages" json:"number_of_packages,omitempty" toml:"number_of_packages" yaml:"number_of_packages,omitempty"`
	CommodityID       int               `boil:"commodity_id" json:"commodity_id" toml:"commodity_id" yaml:"commodity_id"`
	HS2Code           string            `boil:"hs2_code" json:"hs2_code" toml:"hs2_code" yaml:"hs2_code"`
	GrossWeight       types.NullDecimal `boil:"gross_weight" json:"gross_weight,omitempty" toml:"gross_weight" yaml:"gross_weight,omitempty"`
	UnitOfMeasurement string            `boil:"unit_of_measurement" json:"unit_of_measurement" toml:"unit_of_measurement" yaml:"unit_of_measurement"`
	UnitPrice         types.NullDecimal `boil:"unit_price" json:"unit_price,omitempty" toml:"unit_price" yaml:"unit_price,omitempty"`
	BankID            int               `boil:"bank_id" json:"bank_id" toml:"bank_id" yaml:"bank_id"`
	MakerDate         time.Time         `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate       null.Time         `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID           string            `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID         null.String       `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy        null.String       `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate      null.Time         `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *tFTradeItemR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tFTradeItemL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TFTradeItemColumns = struct {
	TradeItemID       string
	BranchID          string
	TradeID           string
	NumberOfPackages  string
	CommodityID       string
	HS2Code           string
	GrossWeight       string
	UnitOfMeasurement string
	UnitPrice         string
	BankID            string
	MakerDate         string
	CheckerDate       string
	MakerID           string
	CheckerID         string
	ModifiedBy        string
	ModifiedDate      string
}{
	TradeItemID:       "trade_item_id",
	BranchID:          "branch_id",
	TradeID:           "trade_id",
	NumberOfPackages:  "number_of_packages",
	CommodityID:       "commodity_id",
	HS2Code:           "hs2_code",
	GrossWeight:       "gross_weight",
	UnitOfMeasurement: "unit_of_measurement",
	UnitPrice:         "unit_price",
	BankID:            "bank_id",
	MakerDate:         "maker_date",
	CheckerDate:       "checker_date",
	MakerID:           "maker_id",
	CheckerID:         "checker_id",
	ModifiedBy:        "modified_by",
	ModifiedDate:      "modified_date",
}

// Generated where

var TFTradeItemWhere = struct {
	TradeItemID       whereHelperint
	BranchID          whereHelpernull_Int
	TradeID           whereHelperint
	NumberOfPackages  whereHelpernull_Int
	CommodityID       whereHelperint
	HS2Code           whereHelperstring
	GrossWeight       whereHelpertypes_NullDecimal
	UnitOfMeasurement whereHelperstring
	UnitPrice         whereHelpertypes_NullDecimal
	BankID            whereHelperint
	MakerDate         whereHelpertime_Time
	CheckerDate       whereHelpernull_Time
	MakerID           whereHelperstring
	CheckerID         whereHelpernull_String
	ModifiedBy        whereHelpernull_String
	ModifiedDate      whereHelpernull_Time
}{
	TradeItemID:       whereHelperint{field: `trade_item_id`},
	BranchID:          whereHelpernull_Int{field: `branch_id`},
	TradeID:           whereHelperint{field: `trade_id`},
	NumberOfPackages:  whereHelpernull_Int{field: `number_of_packages`},
	CommodityID:       whereHelperint{field: `commodity_id`},
	HS2Code:           whereHelperstring{field: `hs2_code`},
	GrossWeight:       whereHelpertypes_NullDecimal{field: `gross_weight`},
	UnitOfMeasurement: whereHelperstring{field: `unit_of_measurement`},
	UnitPrice:         whereHelpertypes_NullDecimal{field: `unit_price`},
	BankID:            whereHelperint{field: `bank_id`},
	MakerDate:         whereHelpertime_Time{field: `maker_date`},
	CheckerDate:       whereHelpernull_Time{field: `checker_date`},
	MakerID:           whereHelperstring{field: `maker_id`},
	CheckerID:         whereHelpernull_String{field: `checker_id`},
	ModifiedBy:        whereHelpernull_String{field: `modified_by`},
	ModifiedDate:      whereHelpernull_Time{field: `modified_date`},
}

// TFTradeItemRels is where relationship names are stored.
var TFTradeItemRels = struct {
	Trade string
}{
	Trade: "Trade",
}

// tFTradeItemR is where relationships are stored.
type tFTradeItemR struct {
	Trade *TFTrade
}

// NewStruct creates a new relationship struct
func (*tFTradeItemR) NewStruct() *tFTradeItemR {
	return &tFTradeItemR{}
}

// tFTradeItemL is where Load methods for each relationship are stored.
type tFTradeItemL struct{}

var (
	tFTradeItemColumns               = []string{"trade_item_id", "branch_id", "trade_id", "number_of_packages", "commodity_id", "hs2_code", "gross_weight", "unit_of_measurement", "unit_price", "bank_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	tFTradeItemColumnsWithoutDefault = []string{"trade_item_id", "branch_id", "trade_id", "number_of_packages", "commodity_id", "hs2_code", "gross_weight", "unit_of_measurement", "unit_price", "bank_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	tFTradeItemColumnsWithDefault    = []string{}
	tFTradeItemPrimaryKeyColumns     = []string{"trade_item_id"}
)

type (
	// TFTradeItemSlice is an alias for a slice of pointers to TFTradeItem.
	// This should generally be used opposed to []TFTradeItem.
	TFTradeItemSlice []*TFTradeItem
	// TFTradeItemHook is the signature for custom TFTradeItem hook methods
	TFTradeItemHook func(context.Context, boil.ContextExecutor, *TFTradeItem) error

	tFTradeItemQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tFTradeItemType                 = reflect.TypeOf(&TFTradeItem{})
	tFTradeItemMapping              = queries.MakeStructMapping(tFTradeItemType)
	tFTradeItemPrimaryKeyMapping, _ = queries.BindMapping(tFTradeItemType, tFTradeItemMapping, tFTradeItemPrimaryKeyColumns)
	tFTradeItemInsertCacheMut       sync.RWMutex
	tFTradeItemInsertCache          = make(map[string]insertCache)
	tFTradeItemUpdateCacheMut       sync.RWMutex
	tFTradeItemUpdateCache          = make(map[string]updateCache)
	tFTradeItemUpsertCacheMut       sync.RWMutex
	tFTradeItemUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tFTradeItemBeforeInsertHooks []TFTradeItemHook
var tFTradeItemBeforeUpdateHooks []TFTradeItemHook
var tFTradeItemBeforeDeleteHooks []TFTradeItemHook
var tFTradeItemBeforeUpsertHooks []TFTradeItemHook

var tFTradeItemAfterInsertHooks []TFTradeItemHook
var tFTradeItemAfterSelectHooks []TFTradeItemHook
var tFTradeItemAfterUpdateHooks []TFTradeItemHook
var tFTradeItemAfterDeleteHooks []TFTradeItemHook
var tFTradeItemAfterUpsertHooks []TFTradeItemHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TFTradeItem) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFTradeItemBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TFTradeItem) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFTradeItemBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TFTradeItem) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFTradeItemBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TFTradeItem) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFTradeItemBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TFTradeItem) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFTradeItemAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TFTradeItem) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFTradeItemAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TFTradeItem) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFTradeItemAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TFTradeItem) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFTradeItemAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TFTradeItem) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFTradeItemAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTFTradeItemHook registers your hook function for all future operations.
func AddTFTradeItemHook(hookPoint boil.HookPoint, tFTradeItemHook TFTradeItemHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		tFTradeItemBeforeInsertHooks = append(tFTradeItemBeforeInsertHooks, tFTradeItemHook)
	case boil.BeforeUpdateHook:
		tFTradeItemBeforeUpdateHooks = append(tFTradeItemBeforeUpdateHooks, tFTradeItemHook)
	case boil.BeforeDeleteHook:
		tFTradeItemBeforeDeleteHooks = append(tFTradeItemBeforeDeleteHooks, tFTradeItemHook)
	case boil.BeforeUpsertHook:
		tFTradeItemBeforeUpsertHooks = append(tFTradeItemBeforeUpsertHooks, tFTradeItemHook)
	case boil.AfterInsertHook:
		tFTradeItemAfterInsertHooks = append(tFTradeItemAfterInsertHooks, tFTradeItemHook)
	case boil.AfterSelectHook:
		tFTradeItemAfterSelectHooks = append(tFTradeItemAfterSelectHooks, tFTradeItemHook)
	case boil.AfterUpdateHook:
		tFTradeItemAfterUpdateHooks = append(tFTradeItemAfterUpdateHooks, tFTradeItemHook)
	case boil.AfterDeleteHook:
		tFTradeItemAfterDeleteHooks = append(tFTradeItemAfterDeleteHooks, tFTradeItemHook)
	case boil.AfterUpsertHook:
		tFTradeItemAfterUpsertHooks = append(tFTradeItemAfterUpsertHooks, tFTradeItemHook)
	}
}

// One returns a single tFTradeItem record from the query.
func (q tFTradeItemQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TFTradeItem, error) {
	o := &TFTradeItem{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for TFTradeItem")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TFTradeItem records from the query.
func (q tFTradeItemQuery) All(ctx context.Context, exec boil.ContextExecutor) (TFTradeItemSlice, error) {
	var o []*TFTradeItem

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TFTradeItem slice")
	}

	if len(tFTradeItemAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TFTradeItem records in the query.
func (q tFTradeItemQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count TFTradeItem rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tFTradeItemQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if TFTradeItem exists")
	}

	return count > 0, nil
}

// Trade pointed to by the foreign key.
func (o *TFTradeItem) Trade(mods ...qm.QueryMod) tFTradeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("trade_id=?", o.TradeID),
	}

	queryMods = append(queryMods, mods...)

	query := TFTrades(queryMods...)
	queries.SetFrom(query.Query, "`TFTrade`")

	return query
}

// LoadTrade allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (tFTradeItemL) LoadTrade(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTFTradeItem interface{}, mods queries.Applicator) error {
	var slice []*TFTradeItem
	var object *TFTradeItem

	if singular {
		object = maybeTFTradeItem.(*TFTradeItem)
	} else {
		slice = *maybeTFTradeItem.(*[]*TFTradeItem)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tFTradeItemR{}
		}
		args = append(args, object.TradeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tFTradeItemR{}
			}

			for _, a := range args {
				if a == obj.TradeID {
					continue Outer
				}
			}

			args = append(args, obj.TradeID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`TFTrade`), qm.WhereIn(`trade_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load TFTrade")
	}

	var resultSlice []*TFTrade
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice TFTrade")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for TFTrade")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for TFTrade")
	}

	if len(tFTradeItemAfterSelectHooks) != 0 {
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
		object.R.Trade = foreign
		if foreign.R == nil {
			foreign.R = &tFTradeR{}
		}
		foreign.R.TradeTFTradeItems = append(foreign.R.TradeTFTradeItems, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TradeID == foreign.TradeID {
				local.R.Trade = foreign
				if foreign.R == nil {
					foreign.R = &tFTradeR{}
				}
				foreign.R.TradeTFTradeItems = append(foreign.R.TradeTFTradeItems, local)
				break
			}
		}
	}

	return nil
}

// SetTrade of the tFTradeItem to the related item.
// Sets o.R.Trade to related.
// Adds o to related.R.TradeTFTradeItems.
func (o *TFTradeItem) SetTrade(ctx context.Context, exec boil.ContextExecutor, insert bool, related *TFTrade) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `TFTradeItem` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"trade_id"}),
		strmangle.WhereClause("`", "`", 0, tFTradeItemPrimaryKeyColumns),
	)
	values := []interface{}{related.TradeID, o.TradeItemID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TradeID = related.TradeID
	if o.R == nil {
		o.R = &tFTradeItemR{
			Trade: related,
		}
	} else {
		o.R.Trade = related
	}

	if related.R == nil {
		related.R = &tFTradeR{
			TradeTFTradeItems: TFTradeItemSlice{o},
		}
	} else {
		related.R.TradeTFTradeItems = append(related.R.TradeTFTradeItems, o)
	}

	return nil
}

// TFTradeItems retrieves all the records using an executor.
func TFTradeItems(mods ...qm.QueryMod) tFTradeItemQuery {
	mods = append(mods, qm.From("`TFTradeItem`"))
	return tFTradeItemQuery{NewQuery(mods...)}
}

// FindTFTradeItem retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTFTradeItem(ctx context.Context, exec boil.ContextExecutor, tradeItemID int, selectCols ...string) (*TFTradeItem, error) {
	tFTradeItemObj := &TFTradeItem{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `TFTradeItem` where `trade_item_id`=?", sel,
	)

	q := queries.Raw(query, tradeItemID)

	err := q.Bind(ctx, exec, tFTradeItemObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from TFTradeItem")
	}

	return tFTradeItemObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TFTradeItem) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TFTradeItem provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tFTradeItemColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tFTradeItemInsertCacheMut.RLock()
	cache, cached := tFTradeItemInsertCache[key]
	tFTradeItemInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tFTradeItemColumns,
			tFTradeItemColumnsWithDefault,
			tFTradeItemColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tFTradeItemType, tFTradeItemMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tFTradeItemType, tFTradeItemMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `TFTradeItem` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `TFTradeItem` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `TFTradeItem` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tFTradeItemPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into TFTradeItem")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.TradeItemID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TFTradeItem")
	}

CacheNoHooks:
	if !cached {
		tFTradeItemInsertCacheMut.Lock()
		tFTradeItemInsertCache[key] = cache
		tFTradeItemInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TFTradeItem.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TFTradeItem) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tFTradeItemUpdateCacheMut.RLock()
	cache, cached := tFTradeItemUpdateCache[key]
	tFTradeItemUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tFTradeItemColumns,
			tFTradeItemPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update TFTradeItem, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `TFTradeItem` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tFTradeItemPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tFTradeItemType, tFTradeItemMapping, append(wl, tFTradeItemPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update TFTradeItem row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for TFTradeItem")
	}

	if !cached {
		tFTradeItemUpdateCacheMut.Lock()
		tFTradeItemUpdateCache[key] = cache
		tFTradeItemUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tFTradeItemQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for TFTradeItem")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for TFTradeItem")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TFTradeItemSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tFTradeItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `TFTradeItem` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tFTradeItemPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tFTradeItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tFTradeItem")
	}
	return rowsAff, nil
}

var mySQLTFTradeItemUniqueColumns = []string{
	"trade_item_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TFTradeItem) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TFTradeItem provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tFTradeItemColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTFTradeItemUniqueColumns, o)

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

	tFTradeItemUpsertCacheMut.RLock()
	cache, cached := tFTradeItemUpsertCache[key]
	tFTradeItemUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tFTradeItemColumns,
			tFTradeItemColumnsWithDefault,
			tFTradeItemColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			tFTradeItemColumns,
			tFTradeItemPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert TFTradeItem, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "TFTradeItem", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `TFTradeItem` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tFTradeItemType, tFTradeItemMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tFTradeItemType, tFTradeItemMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for TFTradeItem")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tFTradeItemType, tFTradeItemMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for TFTradeItem")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TFTradeItem")
	}

CacheNoHooks:
	if !cached {
		tFTradeItemUpsertCacheMut.Lock()
		tFTradeItemUpsertCache[key] = cache
		tFTradeItemUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TFTradeItem record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TFTradeItem) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TFTradeItem provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tFTradeItemPrimaryKeyMapping)
	sql := "DELETE FROM `TFTradeItem` WHERE `trade_item_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from TFTradeItem")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for TFTradeItem")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tFTradeItemQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tFTradeItemQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from TFTradeItem")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TFTradeItem")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TFTradeItemSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TFTradeItem slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(tFTradeItemBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tFTradeItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `TFTradeItem` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tFTradeItemPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tFTradeItem slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TFTradeItem")
	}

	if len(tFTradeItemAfterDeleteHooks) != 0 {
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
func (o *TFTradeItem) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTFTradeItem(ctx, exec, o.TradeItemID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TFTradeItemSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TFTradeItemSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tFTradeItemPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `TFTradeItem`.* FROM `TFTradeItem` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tFTradeItemPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TFTradeItemSlice")
	}

	*o = slice

	return nil
}

// TFTradeItemExists checks if the TFTradeItem row exists.
func TFTradeItemExists(ctx context.Context, exec boil.ContextExecutor, tradeItemID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `TFTradeItem` where `trade_item_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, tradeItemID)
	}

	row := exec.QueryRowContext(ctx, sql, tradeItemID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if TFTradeItem exists")
	}

	return exists, nil
}
