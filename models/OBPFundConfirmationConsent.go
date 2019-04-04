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

// OBPFundConfirmationConsent is an object representing the database table.
type OBPFundConfirmationConsent struct {
	ObfundconfirmationconsentID int         `boil:"obfundconfirmationconsent_id" json:"obfundconfirmationconsent_id" toml:"obfundconfirmationconsent_id" yaml:"obfundconfirmationconsent_id"`
	ConsentID                   string      `boil:"consent_id" json:"consent_id" toml:"consent_id" yaml:"consent_id"`
	AccountID                   int         `boil:"account_id" json:"account_id" toml:"account_id" yaml:"account_id"`
	CreationDateTime            time.Time   `boil:"creation_date_time" json:"creation_date_time" toml:"creation_date_time" yaml:"creation_date_time"`
	Status                      string      `boil:"status" json:"status" toml:"status" yaml:"status"`
	StatusUpdateDatetime        time.Time   `boil:"status_update_datetime" json:"status_update_datetime" toml:"status_update_datetime" yaml:"status_update_datetime"`
	ExpirationDateTime          time.Time   `boil:"expiration_date_time" json:"expiration_date_time" toml:"expiration_date_time" yaml:"expiration_date_time"`
	MakerDate                   time.Time   `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate                 null.Time   `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID                     string      `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID                   null.String `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy                  null.String `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate                null.Time   `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *oBPFundConfirmationConsentR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L oBPFundConfirmationConsentL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OBPFundConfirmationConsentColumns = struct {
	ObfundconfirmationconsentID string
	ConsentID                   string
	AccountID                   string
	CreationDateTime            string
	Status                      string
	StatusUpdateDatetime        string
	ExpirationDateTime          string
	MakerDate                   string
	CheckerDate                 string
	MakerID                     string
	CheckerID                   string
	ModifiedBy                  string
	ModifiedDate                string
}{
	ObfundconfirmationconsentID: "obfundconfirmationconsent_id",
	ConsentID:                   "consent_id",
	AccountID:                   "account_id",
	CreationDateTime:            "creation_date_time",
	Status:                      "status",
	StatusUpdateDatetime:        "status_update_datetime",
	ExpirationDateTime:          "expiration_date_time",
	MakerDate:                   "maker_date",
	CheckerDate:                 "checker_date",
	MakerID:                     "maker_id",
	CheckerID:                   "checker_id",
	ModifiedBy:                  "modified_by",
	ModifiedDate:                "modified_date",
}

// Generated where

var OBPFundConfirmationConsentWhere = struct {
	ObfundconfirmationconsentID whereHelperint
	ConsentID                   whereHelperstring
	AccountID                   whereHelperint
	CreationDateTime            whereHelpertime_Time
	Status                      whereHelperstring
	StatusUpdateDatetime        whereHelpertime_Time
	ExpirationDateTime          whereHelpertime_Time
	MakerDate                   whereHelpertime_Time
	CheckerDate                 whereHelpernull_Time
	MakerID                     whereHelperstring
	CheckerID                   whereHelpernull_String
	ModifiedBy                  whereHelpernull_String
	ModifiedDate                whereHelpernull_Time
}{
	ObfundconfirmationconsentID: whereHelperint{field: `obfundconfirmationconsent_id`},
	ConsentID:                   whereHelperstring{field: `consent_id`},
	AccountID:                   whereHelperint{field: `account_id`},
	CreationDateTime:            whereHelpertime_Time{field: `creation_date_time`},
	Status:                      whereHelperstring{field: `status`},
	StatusUpdateDatetime:        whereHelpertime_Time{field: `status_update_datetime`},
	ExpirationDateTime:          whereHelpertime_Time{field: `expiration_date_time`},
	MakerDate:                   whereHelpertime_Time{field: `maker_date`},
	CheckerDate:                 whereHelpernull_Time{field: `checker_date`},
	MakerID:                     whereHelperstring{field: `maker_id`},
	CheckerID:                   whereHelpernull_String{field: `checker_id`},
	ModifiedBy:                  whereHelpernull_String{field: `modified_by`},
	ModifiedDate:                whereHelpernull_Time{field: `modified_date`},
}

// OBPFundConfirmationConsentRels is where relationship names are stored.
var OBPFundConfirmationConsentRels = struct {
	Account string
}{
	Account: "Account",
}

// oBPFundConfirmationConsentR is where relationships are stored.
type oBPFundConfirmationConsentR struct {
	Account *Account
}

// NewStruct creates a new relationship struct
func (*oBPFundConfirmationConsentR) NewStruct() *oBPFundConfirmationConsentR {
	return &oBPFundConfirmationConsentR{}
}

// oBPFundConfirmationConsentL is where Load methods for each relationship are stored.
type oBPFundConfirmationConsentL struct{}

var (
	oBPFundConfirmationConsentColumns               = []string{"obfundconfirmationconsent_id", "consent_id", "account_id", "creation_date_time", "status", "status_update_datetime", "expiration_date_time", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	oBPFundConfirmationConsentColumnsWithoutDefault = []string{"consent_id", "account_id", "creation_date_time", "status", "status_update_datetime", "expiration_date_time", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	oBPFundConfirmationConsentColumnsWithDefault    = []string{"obfundconfirmationconsent_id"}
	oBPFundConfirmationConsentPrimaryKeyColumns     = []string{"obfundconfirmationconsent_id"}
)

type (
	// OBPFundConfirmationConsentSlice is an alias for a slice of pointers to OBPFundConfirmationConsent.
	// This should generally be used opposed to []OBPFundConfirmationConsent.
	OBPFundConfirmationConsentSlice []*OBPFundConfirmationConsent
	// OBPFundConfirmationConsentHook is the signature for custom OBPFundConfirmationConsent hook methods
	OBPFundConfirmationConsentHook func(context.Context, boil.ContextExecutor, *OBPFundConfirmationConsent) error

	oBPFundConfirmationConsentQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	oBPFundConfirmationConsentType                 = reflect.TypeOf(&OBPFundConfirmationConsent{})
	oBPFundConfirmationConsentMapping              = queries.MakeStructMapping(oBPFundConfirmationConsentType)
	oBPFundConfirmationConsentPrimaryKeyMapping, _ = queries.BindMapping(oBPFundConfirmationConsentType, oBPFundConfirmationConsentMapping, oBPFundConfirmationConsentPrimaryKeyColumns)
	oBPFundConfirmationConsentInsertCacheMut       sync.RWMutex
	oBPFundConfirmationConsentInsertCache          = make(map[string]insertCache)
	oBPFundConfirmationConsentUpdateCacheMut       sync.RWMutex
	oBPFundConfirmationConsentUpdateCache          = make(map[string]updateCache)
	oBPFundConfirmationConsentUpsertCacheMut       sync.RWMutex
	oBPFundConfirmationConsentUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var oBPFundConfirmationConsentBeforeInsertHooks []OBPFundConfirmationConsentHook
var oBPFundConfirmationConsentBeforeUpdateHooks []OBPFundConfirmationConsentHook
var oBPFundConfirmationConsentBeforeDeleteHooks []OBPFundConfirmationConsentHook
var oBPFundConfirmationConsentBeforeUpsertHooks []OBPFundConfirmationConsentHook

var oBPFundConfirmationConsentAfterInsertHooks []OBPFundConfirmationConsentHook
var oBPFundConfirmationConsentAfterSelectHooks []OBPFundConfirmationConsentHook
var oBPFundConfirmationConsentAfterUpdateHooks []OBPFundConfirmationConsentHook
var oBPFundConfirmationConsentAfterDeleteHooks []OBPFundConfirmationConsentHook
var oBPFundConfirmationConsentAfterUpsertHooks []OBPFundConfirmationConsentHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *OBPFundConfirmationConsent) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBPFundConfirmationConsentBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *OBPFundConfirmationConsent) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBPFundConfirmationConsentBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *OBPFundConfirmationConsent) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBPFundConfirmationConsentBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *OBPFundConfirmationConsent) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBPFundConfirmationConsentBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *OBPFundConfirmationConsent) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBPFundConfirmationConsentAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *OBPFundConfirmationConsent) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBPFundConfirmationConsentAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *OBPFundConfirmationConsent) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBPFundConfirmationConsentAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *OBPFundConfirmationConsent) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBPFundConfirmationConsentAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *OBPFundConfirmationConsent) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range oBPFundConfirmationConsentAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddOBPFundConfirmationConsentHook registers your hook function for all future operations.
func AddOBPFundConfirmationConsentHook(hookPoint boil.HookPoint, oBPFundConfirmationConsentHook OBPFundConfirmationConsentHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		oBPFundConfirmationConsentBeforeInsertHooks = append(oBPFundConfirmationConsentBeforeInsertHooks, oBPFundConfirmationConsentHook)
	case boil.BeforeUpdateHook:
		oBPFundConfirmationConsentBeforeUpdateHooks = append(oBPFundConfirmationConsentBeforeUpdateHooks, oBPFundConfirmationConsentHook)
	case boil.BeforeDeleteHook:
		oBPFundConfirmationConsentBeforeDeleteHooks = append(oBPFundConfirmationConsentBeforeDeleteHooks, oBPFundConfirmationConsentHook)
	case boil.BeforeUpsertHook:
		oBPFundConfirmationConsentBeforeUpsertHooks = append(oBPFundConfirmationConsentBeforeUpsertHooks, oBPFundConfirmationConsentHook)
	case boil.AfterInsertHook:
		oBPFundConfirmationConsentAfterInsertHooks = append(oBPFundConfirmationConsentAfterInsertHooks, oBPFundConfirmationConsentHook)
	case boil.AfterSelectHook:
		oBPFundConfirmationConsentAfterSelectHooks = append(oBPFundConfirmationConsentAfterSelectHooks, oBPFundConfirmationConsentHook)
	case boil.AfterUpdateHook:
		oBPFundConfirmationConsentAfterUpdateHooks = append(oBPFundConfirmationConsentAfterUpdateHooks, oBPFundConfirmationConsentHook)
	case boil.AfterDeleteHook:
		oBPFundConfirmationConsentAfterDeleteHooks = append(oBPFundConfirmationConsentAfterDeleteHooks, oBPFundConfirmationConsentHook)
	case boil.AfterUpsertHook:
		oBPFundConfirmationConsentAfterUpsertHooks = append(oBPFundConfirmationConsentAfterUpsertHooks, oBPFundConfirmationConsentHook)
	}
}

// One returns a single oBPFundConfirmationConsent record from the query.
func (q oBPFundConfirmationConsentQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OBPFundConfirmationConsent, error) {
	o := &OBPFundConfirmationConsent{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for OBPFundConfirmationConsent")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all OBPFundConfirmationConsent records from the query.
func (q oBPFundConfirmationConsentQuery) All(ctx context.Context, exec boil.ContextExecutor) (OBPFundConfirmationConsentSlice, error) {
	var o []*OBPFundConfirmationConsent

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OBPFundConfirmationConsent slice")
	}

	if len(oBPFundConfirmationConsentAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all OBPFundConfirmationConsent records in the query.
func (q oBPFundConfirmationConsentQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count OBPFundConfirmationConsent rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q oBPFundConfirmationConsentQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if OBPFundConfirmationConsent exists")
	}

	return count > 0, nil
}

// Account pointed to by the foreign key.
func (o *OBPFundConfirmationConsent) Account(mods ...qm.QueryMod) accountQuery {
	queryMods := []qm.QueryMod{
		qm.Where("account_id=?", o.AccountID),
	}

	queryMods = append(queryMods, mods...)

	query := Accounts(queryMods...)
	queries.SetFrom(query.Query, "`Account`")

	return query
}

// LoadAccount allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (oBPFundConfirmationConsentL) LoadAccount(ctx context.Context, e boil.ContextExecutor, singular bool, maybeOBPFundConfirmationConsent interface{}, mods queries.Applicator) error {
	var slice []*OBPFundConfirmationConsent
	var object *OBPFundConfirmationConsent

	if singular {
		object = maybeOBPFundConfirmationConsent.(*OBPFundConfirmationConsent)
	} else {
		slice = *maybeOBPFundConfirmationConsent.(*[]*OBPFundConfirmationConsent)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &oBPFundConfirmationConsentR{}
		}
		args = append(args, object.AccountID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &oBPFundConfirmationConsentR{}
			}

			for _, a := range args {
				if a == obj.AccountID {
					continue Outer
				}
			}

			args = append(args, obj.AccountID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`Account`), qm.WhereIn(`account_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Account")
	}

	var resultSlice []*Account
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Account")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Account")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Account")
	}

	if len(oBPFundConfirmationConsentAfterSelectHooks) != 0 {
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
		object.R.Account = foreign
		if foreign.R == nil {
			foreign.R = &accountR{}
		}
		foreign.R.AccountOBPFundConfirmationConsents = append(foreign.R.AccountOBPFundConfirmationConsents, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.AccountID == foreign.AccountID {
				local.R.Account = foreign
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.AccountOBPFundConfirmationConsents = append(foreign.R.AccountOBPFundConfirmationConsents, local)
				break
			}
		}
	}

	return nil
}

// SetAccount of the oBPFundConfirmationConsent to the related item.
// Sets o.R.Account to related.
// Adds o to related.R.AccountOBPFundConfirmationConsents.
func (o *OBPFundConfirmationConsent) SetAccount(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Account) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `OBPFundConfirmationConsent` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"account_id"}),
		strmangle.WhereClause("`", "`", 0, oBPFundConfirmationConsentPrimaryKeyColumns),
	)
	values := []interface{}{related.AccountID, o.ObfundconfirmationconsentID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.AccountID = related.AccountID
	if o.R == nil {
		o.R = &oBPFundConfirmationConsentR{
			Account: related,
		}
	} else {
		o.R.Account = related
	}

	if related.R == nil {
		related.R = &accountR{
			AccountOBPFundConfirmationConsents: OBPFundConfirmationConsentSlice{o},
		}
	} else {
		related.R.AccountOBPFundConfirmationConsents = append(related.R.AccountOBPFundConfirmationConsents, o)
	}

	return nil
}

// OBPFundConfirmationConsents retrieves all the records using an executor.
func OBPFundConfirmationConsents(mods ...qm.QueryMod) oBPFundConfirmationConsentQuery {
	mods = append(mods, qm.From("`OBPFundConfirmationConsent`"))
	return oBPFundConfirmationConsentQuery{NewQuery(mods...)}
}

// FindOBPFundConfirmationConsent retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOBPFundConfirmationConsent(ctx context.Context, exec boil.ContextExecutor, obfundconfirmationconsentID int, selectCols ...string) (*OBPFundConfirmationConsent, error) {
	oBPFundConfirmationConsentObj := &OBPFundConfirmationConsent{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `OBPFundConfirmationConsent` where `obfundconfirmationconsent_id`=?", sel,
	)

	q := queries.Raw(query, obfundconfirmationconsentID)

	err := q.Bind(ctx, exec, oBPFundConfirmationConsentObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from OBPFundConfirmationConsent")
	}

	return oBPFundConfirmationConsentObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OBPFundConfirmationConsent) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no OBPFundConfirmationConsent provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(oBPFundConfirmationConsentColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	oBPFundConfirmationConsentInsertCacheMut.RLock()
	cache, cached := oBPFundConfirmationConsentInsertCache[key]
	oBPFundConfirmationConsentInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			oBPFundConfirmationConsentColumns,
			oBPFundConfirmationConsentColumnsWithDefault,
			oBPFundConfirmationConsentColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(oBPFundConfirmationConsentType, oBPFundConfirmationConsentMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(oBPFundConfirmationConsentType, oBPFundConfirmationConsentMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `OBPFundConfirmationConsent` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `OBPFundConfirmationConsent` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `OBPFundConfirmationConsent` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, oBPFundConfirmationConsentPrimaryKeyColumns))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into OBPFundConfirmationConsent")
	}

	var lastID int64
	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ObfundconfirmationconsentID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == oBPFundConfirmationConsentMapping["ObfundconfirmationconsentID"] {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ObfundconfirmationconsentID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for OBPFundConfirmationConsent")
	}

CacheNoHooks:
	if !cached {
		oBPFundConfirmationConsentInsertCacheMut.Lock()
		oBPFundConfirmationConsentInsertCache[key] = cache
		oBPFundConfirmationConsentInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the OBPFundConfirmationConsent.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OBPFundConfirmationConsent) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	oBPFundConfirmationConsentUpdateCacheMut.RLock()
	cache, cached := oBPFundConfirmationConsentUpdateCache[key]
	oBPFundConfirmationConsentUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			oBPFundConfirmationConsentColumns,
			oBPFundConfirmationConsentPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update OBPFundConfirmationConsent, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `OBPFundConfirmationConsent` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, oBPFundConfirmationConsentPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(oBPFundConfirmationConsentType, oBPFundConfirmationConsentMapping, append(wl, oBPFundConfirmationConsentPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update OBPFundConfirmationConsent row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for OBPFundConfirmationConsent")
	}

	if !cached {
		oBPFundConfirmationConsentUpdateCacheMut.Lock()
		oBPFundConfirmationConsentUpdateCache[key] = cache
		oBPFundConfirmationConsentUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q oBPFundConfirmationConsentQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for OBPFundConfirmationConsent")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for OBPFundConfirmationConsent")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OBPFundConfirmationConsentSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oBPFundConfirmationConsentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `OBPFundConfirmationConsent` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oBPFundConfirmationConsentPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in oBPFundConfirmationConsent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all oBPFundConfirmationConsent")
	}
	return rowsAff, nil
}

var mySQLOBPFundConfirmationConsentUniqueColumns = []string{
	"obfundconfirmationconsent_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OBPFundConfirmationConsent) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no OBPFundConfirmationConsent provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(oBPFundConfirmationConsentColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLOBPFundConfirmationConsentUniqueColumns, o)

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

	oBPFundConfirmationConsentUpsertCacheMut.RLock()
	cache, cached := oBPFundConfirmationConsentUpsertCache[key]
	oBPFundConfirmationConsentUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			oBPFundConfirmationConsentColumns,
			oBPFundConfirmationConsentColumnsWithDefault,
			oBPFundConfirmationConsentColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			oBPFundConfirmationConsentColumns,
			oBPFundConfirmationConsentPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert OBPFundConfirmationConsent, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "OBPFundConfirmationConsent", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `OBPFundConfirmationConsent` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(oBPFundConfirmationConsentType, oBPFundConfirmationConsentMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(oBPFundConfirmationConsentType, oBPFundConfirmationConsentMapping, ret)
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

	result, err := exec.ExecContext(ctx, cache.query, vals...)

	if err != nil {
		return errors.Wrap(err, "models: unable to upsert for OBPFundConfirmationConsent")
	}

	var lastID int64
	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	lastID, err = result.LastInsertId()
	if err != nil {
		return ErrSyncFail
	}

	o.ObfundconfirmationconsentID = int(lastID)
	if lastID != 0 && len(cache.retMapping) == 1 && cache.retMapping[0] == oBPFundConfirmationConsentMapping["obfundconfirmationconsent_id"] {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(oBPFundConfirmationConsentType, oBPFundConfirmationConsentMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for OBPFundConfirmationConsent")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for OBPFundConfirmationConsent")
	}

CacheNoHooks:
	if !cached {
		oBPFundConfirmationConsentUpsertCacheMut.Lock()
		oBPFundConfirmationConsentUpsertCache[key] = cache
		oBPFundConfirmationConsentUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single OBPFundConfirmationConsent record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OBPFundConfirmationConsent) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OBPFundConfirmationConsent provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), oBPFundConfirmationConsentPrimaryKeyMapping)
	sql := "DELETE FROM `OBPFundConfirmationConsent` WHERE `obfundconfirmationconsent_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from OBPFundConfirmationConsent")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for OBPFundConfirmationConsent")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q oBPFundConfirmationConsentQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no oBPFundConfirmationConsentQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from OBPFundConfirmationConsent")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for OBPFundConfirmationConsent")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OBPFundConfirmationConsentSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OBPFundConfirmationConsent slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(oBPFundConfirmationConsentBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oBPFundConfirmationConsentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `OBPFundConfirmationConsent` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oBPFundConfirmationConsentPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from oBPFundConfirmationConsent slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for OBPFundConfirmationConsent")
	}

	if len(oBPFundConfirmationConsentAfterDeleteHooks) != 0 {
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
func (o *OBPFundConfirmationConsent) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOBPFundConfirmationConsent(ctx, exec, o.ObfundconfirmationconsentID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OBPFundConfirmationConsentSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OBPFundConfirmationConsentSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), oBPFundConfirmationConsentPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `OBPFundConfirmationConsent`.* FROM `OBPFundConfirmationConsent` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, oBPFundConfirmationConsentPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OBPFundConfirmationConsentSlice")
	}

	*o = slice

	return nil
}

// OBPFundConfirmationConsentExists checks if the OBPFundConfirmationConsent row exists.
func OBPFundConfirmationConsentExists(ctx context.Context, exec boil.ContextExecutor, obfundconfirmationconsentID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `OBPFundConfirmationConsent` where `obfundconfirmationconsent_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, obfundconfirmationconsentID)
	}

	row := exec.QueryRowContext(ctx, sql, obfundconfirmationconsentID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if OBPFundConfirmationConsent exists")
	}

	return exists, nil
}
