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

// TFShipment is an object representing the database table.
type TFShipment struct {
	ShipmentID             int         `boil:"shipment_id" json:"shipment_id" toml:"shipment_id" yaml:"shipment_id"`
	BranchID               int         `boil:"branch_id" json:"branch_id" toml:"branch_id" yaml:"branch_id"`
	TradeID                int         `boil:"trade_id" json:"trade_id" toml:"trade_id" yaml:"trade_id"`
	ShippingOrganizationID int         `boil:"shipping_organization_id" json:"shipping_organization_id" toml:"shipping_organization_id" yaml:"shipping_organization_id"`
	ShipmentDate           null.Time   `boil:"shipment_date" json:"shipment_date,omitempty" toml:"shipment_date" yaml:"shipment_date,omitempty"`
	ExportPort             string      `boil:"export_port" json:"export_port" toml:"export_port" yaml:"export_port"`
	ExportPortAddress      null.String `boil:"export_port_address" json:"export_port_address,omitempty" toml:"export_port_address" yaml:"export_port_address,omitempty"`
	DestinationPort        string      `boil:"destination_port" json:"destination_port" toml:"destination_port" yaml:"destination_port"`
	DestinationPortAddress string      `boil:"destination_port_address" json:"destination_port_address" toml:"destination_port_address" yaml:"destination_port_address"`
	ContainerNo            string      `boil:"container_no" json:"container_no" toml:"container_no" yaml:"container_no"`
	Modifiedby             int         `boil:"modifiedby" json:"modifiedby" toml:"modifiedby" yaml:"modifiedby"`
	ModifiedDatetime       null.Time   `boil:"modified_datetime" json:"modified_datetime,omitempty" toml:"modified_datetime" yaml:"modified_datetime,omitempty"`
	ExportCountryCode      null.String `boil:"export_country_code" json:"export_country_code,omitempty" toml:"export_country_code" yaml:"export_country_code,omitempty"`
	ImportCountryCode      null.String `boil:"import_country_code" json:"import_country_code,omitempty" toml:"import_country_code" yaml:"import_country_code,omitempty"`
	BankID                 int         `boil:"bank_id" json:"bank_id" toml:"bank_id" yaml:"bank_id"`
	MakerDate              time.Time   `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate            null.Time   `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID                string      `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID              null.String `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy             null.String `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate           null.Time   `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *tFShipmentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L tFShipmentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var TFShipmentColumns = struct {
	ShipmentID             string
	BranchID               string
	TradeID                string
	ShippingOrganizationID string
	ShipmentDate           string
	ExportPort             string
	ExportPortAddress      string
	DestinationPort        string
	DestinationPortAddress string
	ContainerNo            string
	Modifiedby             string
	ModifiedDatetime       string
	ExportCountryCode      string
	ImportCountryCode      string
	BankID                 string
	MakerDate              string
	CheckerDate            string
	MakerID                string
	CheckerID              string
	ModifiedBy             string
	ModifiedDate           string
}{
	ShipmentID:             "shipment_id",
	BranchID:               "branch_id",
	TradeID:                "trade_id",
	ShippingOrganizationID: "shipping_organization_id",
	ShipmentDate:           "shipment_date",
	ExportPort:             "export_port",
	ExportPortAddress:      "export_port_address",
	DestinationPort:        "destination_port",
	DestinationPortAddress: "destination_port_address",
	ContainerNo:            "container_no",
	Modifiedby:             "modifiedby",
	ModifiedDatetime:       "modified_datetime",
	ExportCountryCode:      "export_country_code",
	ImportCountryCode:      "import_country_code",
	BankID:                 "bank_id",
	MakerDate:              "maker_date",
	CheckerDate:            "checker_date",
	MakerID:                "maker_id",
	CheckerID:              "checker_id",
	ModifiedBy:             "modified_by",
	ModifiedDate:           "modified_date",
}

// Generated where

var TFShipmentWhere = struct {
	ShipmentID             whereHelperint
	BranchID               whereHelperint
	TradeID                whereHelperint
	ShippingOrganizationID whereHelperint
	ShipmentDate           whereHelpernull_Time
	ExportPort             whereHelperstring
	ExportPortAddress      whereHelpernull_String
	DestinationPort        whereHelperstring
	DestinationPortAddress whereHelperstring
	ContainerNo            whereHelperstring
	Modifiedby             whereHelperint
	ModifiedDatetime       whereHelpernull_Time
	ExportCountryCode      whereHelpernull_String
	ImportCountryCode      whereHelpernull_String
	BankID                 whereHelperint
	MakerDate              whereHelpertime_Time
	CheckerDate            whereHelpernull_Time
	MakerID                whereHelperstring
	CheckerID              whereHelpernull_String
	ModifiedBy             whereHelpernull_String
	ModifiedDate           whereHelpernull_Time
}{
	ShipmentID:             whereHelperint{field: `shipment_id`},
	BranchID:               whereHelperint{field: `branch_id`},
	TradeID:                whereHelperint{field: `trade_id`},
	ShippingOrganizationID: whereHelperint{field: `shipping_organization_id`},
	ShipmentDate:           whereHelpernull_Time{field: `shipment_date`},
	ExportPort:             whereHelperstring{field: `export_port`},
	ExportPortAddress:      whereHelpernull_String{field: `export_port_address`},
	DestinationPort:        whereHelperstring{field: `destination_port`},
	DestinationPortAddress: whereHelperstring{field: `destination_port_address`},
	ContainerNo:            whereHelperstring{field: `container_no`},
	Modifiedby:             whereHelperint{field: `modifiedby`},
	ModifiedDatetime:       whereHelpernull_Time{field: `modified_datetime`},
	ExportCountryCode:      whereHelpernull_String{field: `export_country_code`},
	ImportCountryCode:      whereHelpernull_String{field: `import_country_code`},
	BankID:                 whereHelperint{field: `bank_id`},
	MakerDate:              whereHelpertime_Time{field: `maker_date`},
	CheckerDate:            whereHelpernull_Time{field: `checker_date`},
	MakerID:                whereHelperstring{field: `maker_id`},
	CheckerID:              whereHelpernull_String{field: `checker_id`},
	ModifiedBy:             whereHelpernull_String{field: `modified_by`},
	ModifiedDate:           whereHelpernull_Time{field: `modified_date`},
}

// TFShipmentRels is where relationship names are stored.
var TFShipmentRels = struct {
	Trade string
}{
	Trade: "Trade",
}

// tFShipmentR is where relationships are stored.
type tFShipmentR struct {
	Trade *TFTrade
}

// NewStruct creates a new relationship struct
func (*tFShipmentR) NewStruct() *tFShipmentR {
	return &tFShipmentR{}
}

// tFShipmentL is where Load methods for each relationship are stored.
type tFShipmentL struct{}

var (
	tFShipmentColumns               = []string{"shipment_id", "branch_id", "trade_id", "shipping_organization_id", "shipment_date", "export_port", "export_port_address", "destination_port", "destination_port_address", "container_no", "modifiedby", "modified_datetime", "export_country_code", "import_country_code", "bank_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	tFShipmentColumnsWithoutDefault = []string{"shipment_id", "branch_id", "trade_id", "shipping_organization_id", "shipment_date", "export_port", "export_port_address", "destination_port", "destination_port_address", "container_no", "modifiedby", "modified_datetime", "export_country_code", "import_country_code", "bank_id", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	tFShipmentColumnsWithDefault    = []string{}
	tFShipmentPrimaryKeyColumns     = []string{"shipment_id"}
)

type (
	// TFShipmentSlice is an alias for a slice of pointers to TFShipment.
	// This should generally be used opposed to []TFShipment.
	TFShipmentSlice []*TFShipment
	// TFShipmentHook is the signature for custom TFShipment hook methods
	TFShipmentHook func(context.Context, boil.ContextExecutor, *TFShipment) error

	tFShipmentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	tFShipmentType                 = reflect.TypeOf(&TFShipment{})
	tFShipmentMapping              = queries.MakeStructMapping(tFShipmentType)
	tFShipmentPrimaryKeyMapping, _ = queries.BindMapping(tFShipmentType, tFShipmentMapping, tFShipmentPrimaryKeyColumns)
	tFShipmentInsertCacheMut       sync.RWMutex
	tFShipmentInsertCache          = make(map[string]insertCache)
	tFShipmentUpdateCacheMut       sync.RWMutex
	tFShipmentUpdateCache          = make(map[string]updateCache)
	tFShipmentUpsertCacheMut       sync.RWMutex
	tFShipmentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var tFShipmentBeforeInsertHooks []TFShipmentHook
var tFShipmentBeforeUpdateHooks []TFShipmentHook
var tFShipmentBeforeDeleteHooks []TFShipmentHook
var tFShipmentBeforeUpsertHooks []TFShipmentHook

var tFShipmentAfterInsertHooks []TFShipmentHook
var tFShipmentAfterSelectHooks []TFShipmentHook
var tFShipmentAfterUpdateHooks []TFShipmentHook
var tFShipmentAfterDeleteHooks []TFShipmentHook
var tFShipmentAfterUpsertHooks []TFShipmentHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *TFShipment) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFShipmentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *TFShipment) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFShipmentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *TFShipment) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFShipmentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *TFShipment) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFShipmentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *TFShipment) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFShipmentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *TFShipment) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFShipmentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *TFShipment) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFShipmentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *TFShipment) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFShipmentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *TFShipment) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range tFShipmentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddTFShipmentHook registers your hook function for all future operations.
func AddTFShipmentHook(hookPoint boil.HookPoint, tFShipmentHook TFShipmentHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		tFShipmentBeforeInsertHooks = append(tFShipmentBeforeInsertHooks, tFShipmentHook)
	case boil.BeforeUpdateHook:
		tFShipmentBeforeUpdateHooks = append(tFShipmentBeforeUpdateHooks, tFShipmentHook)
	case boil.BeforeDeleteHook:
		tFShipmentBeforeDeleteHooks = append(tFShipmentBeforeDeleteHooks, tFShipmentHook)
	case boil.BeforeUpsertHook:
		tFShipmentBeforeUpsertHooks = append(tFShipmentBeforeUpsertHooks, tFShipmentHook)
	case boil.AfterInsertHook:
		tFShipmentAfterInsertHooks = append(tFShipmentAfterInsertHooks, tFShipmentHook)
	case boil.AfterSelectHook:
		tFShipmentAfterSelectHooks = append(tFShipmentAfterSelectHooks, tFShipmentHook)
	case boil.AfterUpdateHook:
		tFShipmentAfterUpdateHooks = append(tFShipmentAfterUpdateHooks, tFShipmentHook)
	case boil.AfterDeleteHook:
		tFShipmentAfterDeleteHooks = append(tFShipmentAfterDeleteHooks, tFShipmentHook)
	case boil.AfterUpsertHook:
		tFShipmentAfterUpsertHooks = append(tFShipmentAfterUpsertHooks, tFShipmentHook)
	}
}

// One returns a single tFShipment record from the query.
func (q tFShipmentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*TFShipment, error) {
	o := &TFShipment{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for TFShipment")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all TFShipment records from the query.
func (q tFShipmentQuery) All(ctx context.Context, exec boil.ContextExecutor) (TFShipmentSlice, error) {
	var o []*TFShipment

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to TFShipment slice")
	}

	if len(tFShipmentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all TFShipment records in the query.
func (q tFShipmentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count TFShipment rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q tFShipmentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if TFShipment exists")
	}

	return count > 0, nil
}

// Trade pointed to by the foreign key.
func (o *TFShipment) Trade(mods ...qm.QueryMod) tFTradeQuery {
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
func (tFShipmentL) LoadTrade(ctx context.Context, e boil.ContextExecutor, singular bool, maybeTFShipment interface{}, mods queries.Applicator) error {
	var slice []*TFShipment
	var object *TFShipment

	if singular {
		object = maybeTFShipment.(*TFShipment)
	} else {
		slice = *maybeTFShipment.(*[]*TFShipment)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &tFShipmentR{}
		}
		args = append(args, object.TradeID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &tFShipmentR{}
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

	if len(tFShipmentAfterSelectHooks) != 0 {
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
		foreign.R.TradeTFShipments = append(foreign.R.TradeTFShipments, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.TradeID == foreign.TradeID {
				local.R.Trade = foreign
				if foreign.R == nil {
					foreign.R = &tFTradeR{}
				}
				foreign.R.TradeTFShipments = append(foreign.R.TradeTFShipments, local)
				break
			}
		}
	}

	return nil
}

// SetTrade of the tFShipment to the related item.
// Sets o.R.Trade to related.
// Adds o to related.R.TradeTFShipments.
func (o *TFShipment) SetTrade(ctx context.Context, exec boil.ContextExecutor, insert bool, related *TFTrade) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `TFShipment` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"trade_id"}),
		strmangle.WhereClause("`", "`", 0, tFShipmentPrimaryKeyColumns),
	)
	values := []interface{}{related.TradeID, o.ShipmentID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.TradeID = related.TradeID
	if o.R == nil {
		o.R = &tFShipmentR{
			Trade: related,
		}
	} else {
		o.R.Trade = related
	}

	if related.R == nil {
		related.R = &tFTradeR{
			TradeTFShipments: TFShipmentSlice{o},
		}
	} else {
		related.R.TradeTFShipments = append(related.R.TradeTFShipments, o)
	}

	return nil
}

// TFShipments retrieves all the records using an executor.
func TFShipments(mods ...qm.QueryMod) tFShipmentQuery {
	mods = append(mods, qm.From("`TFShipment`"))
	return tFShipmentQuery{NewQuery(mods...)}
}

// FindTFShipment retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindTFShipment(ctx context.Context, exec boil.ContextExecutor, shipmentID int, selectCols ...string) (*TFShipment, error) {
	tFShipmentObj := &TFShipment{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `TFShipment` where `shipment_id`=?", sel,
	)

	q := queries.Raw(query, shipmentID)

	err := q.Bind(ctx, exec, tFShipmentObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from TFShipment")
	}

	return tFShipmentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *TFShipment) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TFShipment provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tFShipmentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	tFShipmentInsertCacheMut.RLock()
	cache, cached := tFShipmentInsertCache[key]
	tFShipmentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			tFShipmentColumns,
			tFShipmentColumnsWithDefault,
			tFShipmentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(tFShipmentType, tFShipmentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(tFShipmentType, tFShipmentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `TFShipment` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `TFShipment` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `TFShipment` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, tFShipmentPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into TFShipment")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ShipmentID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TFShipment")
	}

CacheNoHooks:
	if !cached {
		tFShipmentInsertCacheMut.Lock()
		tFShipmentInsertCache[key] = cache
		tFShipmentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the TFShipment.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *TFShipment) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	tFShipmentUpdateCacheMut.RLock()
	cache, cached := tFShipmentUpdateCache[key]
	tFShipmentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			tFShipmentColumns,
			tFShipmentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update TFShipment, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `TFShipment` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, tFShipmentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(tFShipmentType, tFShipmentMapping, append(wl, tFShipmentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update TFShipment row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for TFShipment")
	}

	if !cached {
		tFShipmentUpdateCacheMut.Lock()
		tFShipmentUpdateCache[key] = cache
		tFShipmentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q tFShipmentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for TFShipment")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for TFShipment")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o TFShipmentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tFShipmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `TFShipment` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tFShipmentPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in tFShipment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all tFShipment")
	}
	return rowsAff, nil
}

var mySQLTFShipmentUniqueColumns = []string{
	"shipment_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *TFShipment) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no TFShipment provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(tFShipmentColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLTFShipmentUniqueColumns, o)

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

	tFShipmentUpsertCacheMut.RLock()
	cache, cached := tFShipmentUpsertCache[key]
	tFShipmentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			tFShipmentColumns,
			tFShipmentColumnsWithDefault,
			tFShipmentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			tFShipmentColumns,
			tFShipmentPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert TFShipment, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "TFShipment", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `TFShipment` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(tFShipmentType, tFShipmentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(tFShipmentType, tFShipmentMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for TFShipment")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(tFShipmentType, tFShipmentMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for TFShipment")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for TFShipment")
	}

CacheNoHooks:
	if !cached {
		tFShipmentUpsertCacheMut.Lock()
		tFShipmentUpsertCache[key] = cache
		tFShipmentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single TFShipment record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *TFShipment) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TFShipment provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), tFShipmentPrimaryKeyMapping)
	sql := "DELETE FROM `TFShipment` WHERE `shipment_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from TFShipment")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for TFShipment")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q tFShipmentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no tFShipmentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from TFShipment")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TFShipment")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o TFShipmentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no TFShipment slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(tFShipmentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tFShipmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `TFShipment` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tFShipmentPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from tFShipment slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for TFShipment")
	}

	if len(tFShipmentAfterDeleteHooks) != 0 {
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
func (o *TFShipment) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindTFShipment(ctx, exec, o.ShipmentID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *TFShipmentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := TFShipmentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), tFShipmentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `TFShipment`.* FROM `TFShipment` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, tFShipmentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in TFShipmentSlice")
	}

	*o = slice

	return nil
}

// TFShipmentExists checks if the TFShipment row exists.
func TFShipmentExists(ctx context.Context, exec boil.ContextExecutor, shipmentID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `TFShipment` where `shipment_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, shipmentID)
	}

	row := exec.QueryRowContext(ctx, sql, shipmentID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if TFShipment exists")
	}

	return exists, nil
}
