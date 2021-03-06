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

// AccountCreditLine is an object representing the database table.
type AccountCreditLine struct {
	AccountCreditLineID  int               `boil:"account_credit_line_id" json:"account_credit_line_id" toml:"account_credit_line_id" yaml:"account_credit_line_id"`
	AccountID            null.Int          `boil:"account_id" json:"account_id,omitempty" toml:"account_id" yaml:"account_id,omitempty"`
	BankID               null.Int          `boil:"bank_id" json:"bank_id,omitempty" toml:"bank_id" yaml:"bank_id,omitempty"`
	CreditLineType       null.String       `boil:"credit_line_type" json:"credit_line_type,omitempty" toml:"credit_line_type" yaml:"credit_line_type,omitempty"`
	AccountCreditLinecol null.String       `boil:"AccountCreditLinecol" json:"AccountCreditLinecol,omitempty" toml:"AccountCreditLinecol" yaml:"AccountCreditLinecol,omitempty"`
	CreditLineIncluded   null.String       `boil:"credit_line_included" json:"credit_line_included,omitempty" toml:"credit_line_included" yaml:"credit_line_included,omitempty"`
	CreditLineAmount     types.NullDecimal `boil:"credit_line_amount" json:"credit_line_amount,omitempty" toml:"credit_line_amount" yaml:"credit_line_amount,omitempty"`
	CreditLineCurrency   null.String       `boil:"credit_line_currency" json:"credit_line_currency,omitempty" toml:"credit_line_currency" yaml:"credit_line_currency,omitempty"`
	MakerDate            null.Time         `boil:"maker_date" json:"maker_date,omitempty" toml:"maker_date" yaml:"maker_date,omitempty"`
	CheckerDate          null.Time         `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID              null.String       `boil:"maker_id" json:"maker_id,omitempty" toml:"maker_id" yaml:"maker_id,omitempty"`
	CheckerID            null.String       `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy           null.String       `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate         null.Time         `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *accountCreditLineR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L accountCreditLineL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AccountCreditLineColumns = struct {
	AccountCreditLineID  string
	AccountID            string
	BankID               string
	CreditLineType       string
	AccountCreditLinecol string
	CreditLineIncluded   string
	CreditLineAmount     string
	CreditLineCurrency   string
	MakerDate            string
	CheckerDate          string
	MakerID              string
	CheckerID            string
	ModifiedBy           string
	ModifiedDate         string
}{
	AccountCreditLineID:  "account_credit_line_id",
	AccountID:            "account_id",
	BankID:               "bank_id",
	CreditLineType:       "credit_line_type",
	AccountCreditLinecol: "AccountCreditLinecol",
	CreditLineIncluded:   "credit_line_included",
	CreditLineAmount:     "credit_line_amount",
	CreditLineCurrency:   "credit_line_currency",
	MakerDate:            "maker_date",
	CheckerDate:          "checker_date",
	MakerID:              "maker_id",
	CheckerID:            "checker_id",
	ModifiedBy:           "modified_by",
	ModifiedDate:         "modified_date",
}

// Generated where

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AccountCreditLineWhere = struct {
	AccountCreditLineID  whereHelperint
	AccountID            whereHelpernull_Int
	BankID               whereHelpernull_Int
	CreditLineType       whereHelpernull_String
	AccountCreditLinecol whereHelpernull_String
	CreditLineIncluded   whereHelpernull_String
	CreditLineAmount     whereHelpertypes_NullDecimal
	CreditLineCurrency   whereHelpernull_String
	MakerDate            whereHelpernull_Time
	CheckerDate          whereHelpernull_Time
	MakerID              whereHelpernull_String
	CheckerID            whereHelpernull_String
	ModifiedBy           whereHelpernull_String
	ModifiedDate         whereHelpernull_Time
}{
	AccountCreditLineID:  whereHelperint{field: `account_credit_line_id`},
	AccountID:            whereHelpernull_Int{field: `account_id`},
	BankID:               whereHelpernull_Int{field: `bank_id`},
	CreditLineType:       whereHelpernull_String{field: `credit_line_type`},
	AccountCreditLinecol: whereHelpernull_String{field: `AccountCreditLinecol`},
	CreditLineIncluded:   whereHelpernull_String{field: `credit_line_included`},
	CreditLineAmount:     whereHelpertypes_NullDecimal{field: `credit_line_amount`},
	CreditLineCurrency:   whereHelpernull_String{field: `credit_line_currency`},
	MakerDate:            whereHelpernull_Time{field: `maker_date`},
	CheckerDate:          whereHelpernull_Time{field: `checker_date`},
	MakerID:              whereHelpernull_String{field: `maker_id`},
	CheckerID:            whereHelpernull_String{field: `checker_id`},
	ModifiedBy:           whereHelpernull_String{field: `modified_by`},
	ModifiedDate:         whereHelpernull_Time{field: `modified_date`},
}

// AccountCreditLineRels is where relationship names are stored.
var AccountCreditLineRels = struct {
	Account string
}{
	Account: "Account",
}

// accountCreditLineR is where relationships are stored.
type accountCreditLineR struct {
	Account *Account
}

// NewStruct creates a new relationship struct
func (*accountCreditLineR) NewStruct() *accountCreditLineR {
	return &accountCreditLineR{}
}

// accountCreditLineL is where Load methods for each relationship are stored.
type accountCreditLineL struct{}

var (
	accountCreditLineColumns               = []string{"account_credit_line_id", "account_id", "bank_id", "credit_line_type", "AccountCreditLinecol", "credit_line_included", "credit_line_amount", "credit_line_currency", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	accountCreditLineColumnsWithoutDefault = []string{"account_credit_line_id", "account_id", "bank_id", "credit_line_type", "AccountCreditLinecol", "credit_line_included", "credit_line_amount", "credit_line_currency", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	accountCreditLineColumnsWithDefault    = []string{}
	accountCreditLinePrimaryKeyColumns     = []string{"account_credit_line_id"}
)

type (
	// AccountCreditLineSlice is an alias for a slice of pointers to AccountCreditLine.
	// This should generally be used opposed to []AccountCreditLine.
	AccountCreditLineSlice []*AccountCreditLine
	// AccountCreditLineHook is the signature for custom AccountCreditLine hook methods
	AccountCreditLineHook func(context.Context, boil.ContextExecutor, *AccountCreditLine) error

	accountCreditLineQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	accountCreditLineType                 = reflect.TypeOf(&AccountCreditLine{})
	accountCreditLineMapping              = queries.MakeStructMapping(accountCreditLineType)
	accountCreditLinePrimaryKeyMapping, _ = queries.BindMapping(accountCreditLineType, accountCreditLineMapping, accountCreditLinePrimaryKeyColumns)
	accountCreditLineInsertCacheMut       sync.RWMutex
	accountCreditLineInsertCache          = make(map[string]insertCache)
	accountCreditLineUpdateCacheMut       sync.RWMutex
	accountCreditLineUpdateCache          = make(map[string]updateCache)
	accountCreditLineUpsertCacheMut       sync.RWMutex
	accountCreditLineUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var accountCreditLineBeforeInsertHooks []AccountCreditLineHook
var accountCreditLineBeforeUpdateHooks []AccountCreditLineHook
var accountCreditLineBeforeDeleteHooks []AccountCreditLineHook
var accountCreditLineBeforeUpsertHooks []AccountCreditLineHook

var accountCreditLineAfterInsertHooks []AccountCreditLineHook
var accountCreditLineAfterSelectHooks []AccountCreditLineHook
var accountCreditLineAfterUpdateHooks []AccountCreditLineHook
var accountCreditLineAfterDeleteHooks []AccountCreditLineHook
var accountCreditLineAfterUpsertHooks []AccountCreditLineHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *AccountCreditLine) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountCreditLineBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *AccountCreditLine) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountCreditLineBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *AccountCreditLine) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountCreditLineBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *AccountCreditLine) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountCreditLineBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *AccountCreditLine) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountCreditLineAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *AccountCreditLine) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountCreditLineAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *AccountCreditLine) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountCreditLineAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *AccountCreditLine) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountCreditLineAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *AccountCreditLine) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range accountCreditLineAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddAccountCreditLineHook registers your hook function for all future operations.
func AddAccountCreditLineHook(hookPoint boil.HookPoint, accountCreditLineHook AccountCreditLineHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		accountCreditLineBeforeInsertHooks = append(accountCreditLineBeforeInsertHooks, accountCreditLineHook)
	case boil.BeforeUpdateHook:
		accountCreditLineBeforeUpdateHooks = append(accountCreditLineBeforeUpdateHooks, accountCreditLineHook)
	case boil.BeforeDeleteHook:
		accountCreditLineBeforeDeleteHooks = append(accountCreditLineBeforeDeleteHooks, accountCreditLineHook)
	case boil.BeforeUpsertHook:
		accountCreditLineBeforeUpsertHooks = append(accountCreditLineBeforeUpsertHooks, accountCreditLineHook)
	case boil.AfterInsertHook:
		accountCreditLineAfterInsertHooks = append(accountCreditLineAfterInsertHooks, accountCreditLineHook)
	case boil.AfterSelectHook:
		accountCreditLineAfterSelectHooks = append(accountCreditLineAfterSelectHooks, accountCreditLineHook)
	case boil.AfterUpdateHook:
		accountCreditLineAfterUpdateHooks = append(accountCreditLineAfterUpdateHooks, accountCreditLineHook)
	case boil.AfterDeleteHook:
		accountCreditLineAfterDeleteHooks = append(accountCreditLineAfterDeleteHooks, accountCreditLineHook)
	case boil.AfterUpsertHook:
		accountCreditLineAfterUpsertHooks = append(accountCreditLineAfterUpsertHooks, accountCreditLineHook)
	}
}

// One returns a single accountCreditLine record from the query.
func (q accountCreditLineQuery) One(ctx context.Context, exec boil.ContextExecutor) (*AccountCreditLine, error) {
	o := &AccountCreditLine{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for AccountCreditLine")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all AccountCreditLine records from the query.
func (q accountCreditLineQuery) All(ctx context.Context, exec boil.ContextExecutor) (AccountCreditLineSlice, error) {
	var o []*AccountCreditLine

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to AccountCreditLine slice")
	}

	if len(accountCreditLineAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all AccountCreditLine records in the query.
func (q accountCreditLineQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count AccountCreditLine rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q accountCreditLineQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if AccountCreditLine exists")
	}

	return count > 0, nil
}

// Account pointed to by the foreign key.
func (o *AccountCreditLine) Account(mods ...qm.QueryMod) accountQuery {
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
func (accountCreditLineL) LoadAccount(ctx context.Context, e boil.ContextExecutor, singular bool, maybeAccountCreditLine interface{}, mods queries.Applicator) error {
	var slice []*AccountCreditLine
	var object *AccountCreditLine

	if singular {
		object = maybeAccountCreditLine.(*AccountCreditLine)
	} else {
		slice = *maybeAccountCreditLine.(*[]*AccountCreditLine)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &accountCreditLineR{}
		}
		if !queries.IsNil(object.AccountID) {
			args = append(args, object.AccountID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &accountCreditLineR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.AccountID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.AccountID) {
				args = append(args, obj.AccountID)
			}

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

	if len(accountCreditLineAfterSelectHooks) != 0 {
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
		foreign.R.AccountAccountCreditLines = append(foreign.R.AccountAccountCreditLines, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.AccountID, foreign.AccountID) {
				local.R.Account = foreign
				if foreign.R == nil {
					foreign.R = &accountR{}
				}
				foreign.R.AccountAccountCreditLines = append(foreign.R.AccountAccountCreditLines, local)
				break
			}
		}
	}

	return nil
}

// SetAccount of the accountCreditLine to the related item.
// Sets o.R.Account to related.
// Adds o to related.R.AccountAccountCreditLines.
func (o *AccountCreditLine) SetAccount(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Account) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `AccountCreditLine` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"account_id"}),
		strmangle.WhereClause("`", "`", 0, accountCreditLinePrimaryKeyColumns),
	)
	values := []interface{}{related.AccountID, o.AccountCreditLineID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.AccountID, related.AccountID)
	if o.R == nil {
		o.R = &accountCreditLineR{
			Account: related,
		}
	} else {
		o.R.Account = related
	}

	if related.R == nil {
		related.R = &accountR{
			AccountAccountCreditLines: AccountCreditLineSlice{o},
		}
	} else {
		related.R.AccountAccountCreditLines = append(related.R.AccountAccountCreditLines, o)
	}

	return nil
}

// RemoveAccount relationship.
// Sets o.R.Account to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *AccountCreditLine) RemoveAccount(ctx context.Context, exec boil.ContextExecutor, related *Account) error {
	var err error

	queries.SetScanner(&o.AccountID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("account_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Account = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.AccountAccountCreditLines {
		if queries.Equal(o.AccountID, ri.AccountID) {
			continue
		}

		ln := len(related.R.AccountAccountCreditLines)
		if ln > 1 && i < ln-1 {
			related.R.AccountAccountCreditLines[i] = related.R.AccountAccountCreditLines[ln-1]
		}
		related.R.AccountAccountCreditLines = related.R.AccountAccountCreditLines[:ln-1]
		break
	}
	return nil
}

// AccountCreditLines retrieves all the records using an executor.
func AccountCreditLines(mods ...qm.QueryMod) accountCreditLineQuery {
	mods = append(mods, qm.From("`AccountCreditLine`"))
	return accountCreditLineQuery{NewQuery(mods...)}
}

// FindAccountCreditLine retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindAccountCreditLine(ctx context.Context, exec boil.ContextExecutor, accountCreditLineID int, selectCols ...string) (*AccountCreditLine, error) {
	accountCreditLineObj := &AccountCreditLine{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `AccountCreditLine` where `account_credit_line_id`=?", sel,
	)

	q := queries.Raw(query, accountCreditLineID)

	err := q.Bind(ctx, exec, accountCreditLineObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from AccountCreditLine")
	}

	return accountCreditLineObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *AccountCreditLine) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no AccountCreditLine provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(accountCreditLineColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	accountCreditLineInsertCacheMut.RLock()
	cache, cached := accountCreditLineInsertCache[key]
	accountCreditLineInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			accountCreditLineColumns,
			accountCreditLineColumnsWithDefault,
			accountCreditLineColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(accountCreditLineType, accountCreditLineMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(accountCreditLineType, accountCreditLineMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `AccountCreditLine` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `AccountCreditLine` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `AccountCreditLine` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, accountCreditLinePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into AccountCreditLine")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.AccountCreditLineID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for AccountCreditLine")
	}

CacheNoHooks:
	if !cached {
		accountCreditLineInsertCacheMut.Lock()
		accountCreditLineInsertCache[key] = cache
		accountCreditLineInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the AccountCreditLine.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *AccountCreditLine) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	accountCreditLineUpdateCacheMut.RLock()
	cache, cached := accountCreditLineUpdateCache[key]
	accountCreditLineUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			accountCreditLineColumns,
			accountCreditLinePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update AccountCreditLine, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `AccountCreditLine` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, accountCreditLinePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(accountCreditLineType, accountCreditLineMapping, append(wl, accountCreditLinePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update AccountCreditLine row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for AccountCreditLine")
	}

	if !cached {
		accountCreditLineUpdateCacheMut.Lock()
		accountCreditLineUpdateCache[key] = cache
		accountCreditLineUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q accountCreditLineQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for AccountCreditLine")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for AccountCreditLine")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AccountCreditLineSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountCreditLinePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `AccountCreditLine` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, accountCreditLinePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in accountCreditLine slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all accountCreditLine")
	}
	return rowsAff, nil
}

var mySQLAccountCreditLineUniqueColumns = []string{
	"account_credit_line_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *AccountCreditLine) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no AccountCreditLine provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(accountCreditLineColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLAccountCreditLineUniqueColumns, o)

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

	accountCreditLineUpsertCacheMut.RLock()
	cache, cached := accountCreditLineUpsertCache[key]
	accountCreditLineUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			accountCreditLineColumns,
			accountCreditLineColumnsWithDefault,
			accountCreditLineColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			accountCreditLineColumns,
			accountCreditLinePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert AccountCreditLine, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "AccountCreditLine", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `AccountCreditLine` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(accountCreditLineType, accountCreditLineMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(accountCreditLineType, accountCreditLineMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for AccountCreditLine")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(accountCreditLineType, accountCreditLineMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for AccountCreditLine")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for AccountCreditLine")
	}

CacheNoHooks:
	if !cached {
		accountCreditLineUpsertCacheMut.Lock()
		accountCreditLineUpsertCache[key] = cache
		accountCreditLineUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single AccountCreditLine record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *AccountCreditLine) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AccountCreditLine provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), accountCreditLinePrimaryKeyMapping)
	sql := "DELETE FROM `AccountCreditLine` WHERE `account_credit_line_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from AccountCreditLine")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for AccountCreditLine")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q accountCreditLineQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no accountCreditLineQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from AccountCreditLine")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for AccountCreditLine")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AccountCreditLineSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no AccountCreditLine slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(accountCreditLineBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountCreditLinePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `AccountCreditLine` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, accountCreditLinePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from accountCreditLine slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for AccountCreditLine")
	}

	if len(accountCreditLineAfterDeleteHooks) != 0 {
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
func (o *AccountCreditLine) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindAccountCreditLine(ctx, exec, o.AccountCreditLineID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AccountCreditLineSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AccountCreditLineSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), accountCreditLinePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `AccountCreditLine`.* FROM `AccountCreditLine` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, accountCreditLinePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in AccountCreditLineSlice")
	}

	*o = slice

	return nil
}

// AccountCreditLineExists checks if the AccountCreditLine row exists.
func AccountCreditLineExists(ctx context.Context, exec boil.ContextExecutor, accountCreditLineID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `AccountCreditLine` where `account_credit_line_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, accountCreditLineID)
	}

	row := exec.QueryRowContext(ctx, sql, accountCreditLineID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if AccountCreditLine exists")
	}

	return exists, nil
}
