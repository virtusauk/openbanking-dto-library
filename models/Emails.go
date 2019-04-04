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

// Email is an object representing the database table.
type Email struct {
	PartyID      int         `boil:"party_id" json:"party_id" toml:"party_id" yaml:"party_id"`
	BankID       int         `boil:"bank_id" json:"bank_id" toml:"bank_id" yaml:"bank_id"`
	EmailType    string      `boil:"email_type" json:"email_type" toml:"email_type" yaml:"email_type"`
	EmailID      string      `boil:"email_id" json:"email_id" toml:"email_id" yaml:"email_id"`
	Active       string      `boil:"active" json:"active" toml:"active" yaml:"active"`
	MakerDate    time.Time   `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate  null.Time   `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID      string      `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID    null.String `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy   null.String `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate null.Time   `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *emailR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L emailL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var EmailColumns = struct {
	PartyID      string
	BankID       string
	EmailType    string
	EmailID      string
	Active       string
	MakerDate    string
	CheckerDate  string
	MakerID      string
	CheckerID    string
	ModifiedBy   string
	ModifiedDate string
}{
	PartyID:      "party_id",
	BankID:       "bank_id",
	EmailType:    "email_type",
	EmailID:      "email_id",
	Active:       "active",
	MakerDate:    "maker_date",
	CheckerDate:  "checker_date",
	MakerID:      "maker_id",
	CheckerID:    "checker_id",
	ModifiedBy:   "modified_by",
	ModifiedDate: "modified_date",
}

// Generated where

var EmailWhere = struct {
	PartyID      whereHelperint
	BankID       whereHelperint
	EmailType    whereHelperstring
	EmailID      whereHelperstring
	Active       whereHelperstring
	MakerDate    whereHelpertime_Time
	CheckerDate  whereHelpernull_Time
	MakerID      whereHelperstring
	CheckerID    whereHelpernull_String
	ModifiedBy   whereHelpernull_String
	ModifiedDate whereHelpernull_Time
}{
	PartyID:      whereHelperint{field: `party_id`},
	BankID:       whereHelperint{field: `bank_id`},
	EmailType:    whereHelperstring{field: `email_type`},
	EmailID:      whereHelperstring{field: `email_id`},
	Active:       whereHelperstring{field: `active`},
	MakerDate:    whereHelpertime_Time{field: `maker_date`},
	CheckerDate:  whereHelpernull_Time{field: `checker_date`},
	MakerID:      whereHelperstring{field: `maker_id`},
	CheckerID:    whereHelpernull_String{field: `checker_id`},
	ModifiedBy:   whereHelpernull_String{field: `modified_by`},
	ModifiedDate: whereHelpernull_Time{field: `modified_date`},
}

// EmailRels is where relationship names are stored.
var EmailRels = struct {
	Party string
	Bank  string
}{
	Party: "Party",
	Bank:  "Bank",
}

// emailR is where relationships are stored.
type emailR struct {
	Party *Party
	Bank  *Bank
}

// NewStruct creates a new relationship struct
func (*emailR) NewStruct() *emailR {
	return &emailR{}
}

// emailL is where Load methods for each relationship are stored.
type emailL struct{}

var (
	emailColumns               = []string{"party_id", "bank_id", "email_type", "email_id", "active", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	emailColumnsWithoutDefault = []string{"party_id", "bank_id", "email_id", "active", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	emailColumnsWithDefault    = []string{"email_type"}
	emailPrimaryKeyColumns     = []string{"party_id", "email_type"}
)

type (
	// EmailSlice is an alias for a slice of pointers to Email.
	// This should generally be used opposed to []Email.
	EmailSlice []*Email
	// EmailHook is the signature for custom Email hook methods
	EmailHook func(context.Context, boil.ContextExecutor, *Email) error

	emailQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	emailType                 = reflect.TypeOf(&Email{})
	emailMapping              = queries.MakeStructMapping(emailType)
	emailPrimaryKeyMapping, _ = queries.BindMapping(emailType, emailMapping, emailPrimaryKeyColumns)
	emailInsertCacheMut       sync.RWMutex
	emailInsertCache          = make(map[string]insertCache)
	emailUpdateCacheMut       sync.RWMutex
	emailUpdateCache          = make(map[string]updateCache)
	emailUpsertCacheMut       sync.RWMutex
	emailUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var emailBeforeInsertHooks []EmailHook
var emailBeforeUpdateHooks []EmailHook
var emailBeforeDeleteHooks []EmailHook
var emailBeforeUpsertHooks []EmailHook

var emailAfterInsertHooks []EmailHook
var emailAfterSelectHooks []EmailHook
var emailAfterUpdateHooks []EmailHook
var emailAfterDeleteHooks []EmailHook
var emailAfterUpsertHooks []EmailHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Email) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Email) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Email) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Email) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Email) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Email) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Email) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Email) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Email) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range emailAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddEmailHook registers your hook function for all future operations.
func AddEmailHook(hookPoint boil.HookPoint, emailHook EmailHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		emailBeforeInsertHooks = append(emailBeforeInsertHooks, emailHook)
	case boil.BeforeUpdateHook:
		emailBeforeUpdateHooks = append(emailBeforeUpdateHooks, emailHook)
	case boil.BeforeDeleteHook:
		emailBeforeDeleteHooks = append(emailBeforeDeleteHooks, emailHook)
	case boil.BeforeUpsertHook:
		emailBeforeUpsertHooks = append(emailBeforeUpsertHooks, emailHook)
	case boil.AfterInsertHook:
		emailAfterInsertHooks = append(emailAfterInsertHooks, emailHook)
	case boil.AfterSelectHook:
		emailAfterSelectHooks = append(emailAfterSelectHooks, emailHook)
	case boil.AfterUpdateHook:
		emailAfterUpdateHooks = append(emailAfterUpdateHooks, emailHook)
	case boil.AfterDeleteHook:
		emailAfterDeleteHooks = append(emailAfterDeleteHooks, emailHook)
	case boil.AfterUpsertHook:
		emailAfterUpsertHooks = append(emailAfterUpsertHooks, emailHook)
	}
}

// One returns a single email record from the query.
func (q emailQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Email, error) {
	o := &Email{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for Emails")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Email records from the query.
func (q emailQuery) All(ctx context.Context, exec boil.ContextExecutor) (EmailSlice, error) {
	var o []*Email

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Email slice")
	}

	if len(emailAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Email records in the query.
func (q emailQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count Emails rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q emailQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if Emails exists")
	}

	return count > 0, nil
}

// Party pointed to by the foreign key.
func (o *Email) Party(mods ...qm.QueryMod) partyQuery {
	queryMods := []qm.QueryMod{
		qm.Where("party_id=?", o.PartyID),
	}

	queryMods = append(queryMods, mods...)

	query := Parties(queryMods...)
	queries.SetFrom(query.Query, "`Parties`")

	return query
}

// Bank pointed to by the foreign key.
func (o *Email) Bank(mods ...qm.QueryMod) bankQuery {
	queryMods := []qm.QueryMod{
		qm.Where("bank_id=?", o.BankID),
	}

	queryMods = append(queryMods, mods...)

	query := Banks(queryMods...)
	queries.SetFrom(query.Query, "`Banks`")

	return query
}

// LoadParty allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (emailL) LoadParty(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEmail interface{}, mods queries.Applicator) error {
	var slice []*Email
	var object *Email

	if singular {
		object = maybeEmail.(*Email)
	} else {
		slice = *maybeEmail.(*[]*Email)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &emailR{}
		}
		args = append(args, object.PartyID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &emailR{}
			}

			for _, a := range args {
				if a == obj.PartyID {
					continue Outer
				}
			}

			args = append(args, obj.PartyID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`Parties`), qm.WhereIn(`party_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Party")
	}

	var resultSlice []*Party
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Party")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Parties")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Parties")
	}

	if len(emailAfterSelectHooks) != 0 {
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
		object.R.Party = foreign
		if foreign.R == nil {
			foreign.R = &partyR{}
		}
		foreign.R.PartyEmails = append(foreign.R.PartyEmails, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.PartyID == foreign.PartyID {
				local.R.Party = foreign
				if foreign.R == nil {
					foreign.R = &partyR{}
				}
				foreign.R.PartyEmails = append(foreign.R.PartyEmails, local)
				break
			}
		}
	}

	return nil
}

// LoadBank allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (emailL) LoadBank(ctx context.Context, e boil.ContextExecutor, singular bool, maybeEmail interface{}, mods queries.Applicator) error {
	var slice []*Email
	var object *Email

	if singular {
		object = maybeEmail.(*Email)
	} else {
		slice = *maybeEmail.(*[]*Email)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &emailR{}
		}
		args = append(args, object.BankID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &emailR{}
			}

			for _, a := range args {
				if a == obj.BankID {
					continue Outer
				}
			}

			args = append(args, obj.BankID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`Banks`), qm.WhereIn(`bank_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Bank")
	}

	var resultSlice []*Bank
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Bank")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Banks")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Banks")
	}

	if len(emailAfterSelectHooks) != 0 {
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
		object.R.Bank = foreign
		if foreign.R == nil {
			foreign.R = &bankR{}
		}
		foreign.R.BankEmails = append(foreign.R.BankEmails, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.BankID == foreign.BankID {
				local.R.Bank = foreign
				if foreign.R == nil {
					foreign.R = &bankR{}
				}
				foreign.R.BankEmails = append(foreign.R.BankEmails, local)
				break
			}
		}
	}

	return nil
}

// SetParty of the email to the related item.
// Sets o.R.Party to related.
// Adds o to related.R.PartyEmails.
func (o *Email) SetParty(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Party) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `Emails` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"party_id"}),
		strmangle.WhereClause("`", "`", 0, emailPrimaryKeyColumns),
	)
	values := []interface{}{related.PartyID, o.PartyID, o.EmailType}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.PartyID = related.PartyID
	if o.R == nil {
		o.R = &emailR{
			Party: related,
		}
	} else {
		o.R.Party = related
	}

	if related.R == nil {
		related.R = &partyR{
			PartyEmails: EmailSlice{o},
		}
	} else {
		related.R.PartyEmails = append(related.R.PartyEmails, o)
	}

	return nil
}

// SetBank of the email to the related item.
// Sets o.R.Bank to related.
// Adds o to related.R.BankEmails.
func (o *Email) SetBank(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Bank) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `Emails` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"bank_id"}),
		strmangle.WhereClause("`", "`", 0, emailPrimaryKeyColumns),
	)
	values := []interface{}{related.BankID, o.PartyID, o.EmailType}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.BankID = related.BankID
	if o.R == nil {
		o.R = &emailR{
			Bank: related,
		}
	} else {
		o.R.Bank = related
	}

	if related.R == nil {
		related.R = &bankR{
			BankEmails: EmailSlice{o},
		}
	} else {
		related.R.BankEmails = append(related.R.BankEmails, o)
	}

	return nil
}

// Emails retrieves all the records using an executor.
func Emails(mods ...qm.QueryMod) emailQuery {
	mods = append(mods, qm.From("`Emails`"))
	return emailQuery{NewQuery(mods...)}
}

// FindEmail retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindEmail(ctx context.Context, exec boil.ContextExecutor, partyID int, emailType string, selectCols ...string) (*Email, error) {
	emailObj := &Email{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `Emails` where `party_id`=? AND `email_type`=?", sel,
	)

	q := queries.Raw(query, partyID, emailType)

	err := q.Bind(ctx, exec, emailObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from Emails")
	}

	return emailObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Email) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Emails provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(emailColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	emailInsertCacheMut.RLock()
	cache, cached := emailInsertCache[key]
	emailInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			emailColumns,
			emailColumnsWithDefault,
			emailColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(emailType, emailMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(emailType, emailMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `Emails` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `Emails` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `Emails` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, emailPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into Emails")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.PartyID,
		o.EmailType,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for Emails")
	}

CacheNoHooks:
	if !cached {
		emailInsertCacheMut.Lock()
		emailInsertCache[key] = cache
		emailInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Email.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Email) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	emailUpdateCacheMut.RLock()
	cache, cached := emailUpdateCache[key]
	emailUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			emailColumns,
			emailPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update Emails, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `Emails` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, emailPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(emailType, emailMapping, append(wl, emailPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update Emails row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for Emails")
	}

	if !cached {
		emailUpdateCacheMut.Lock()
		emailUpdateCache[key] = cache
		emailUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q emailQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for Emails")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for Emails")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o EmailSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `Emails` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, emailPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in email slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all email")
	}
	return rowsAff, nil
}

var mySQLEmailUniqueColumns = []string{}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Email) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Emails provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(emailColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLEmailUniqueColumns, o)

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

	emailUpsertCacheMut.RLock()
	cache, cached := emailUpsertCache[key]
	emailUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			emailColumns,
			emailColumnsWithDefault,
			emailColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			emailColumns,
			emailPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert Emails, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "Emails", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `Emails` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(emailType, emailMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(emailType, emailMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for Emails")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(emailType, emailMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for Emails")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for Emails")
	}

CacheNoHooks:
	if !cached {
		emailUpsertCacheMut.Lock()
		emailUpsertCache[key] = cache
		emailUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Email record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Email) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Email provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), emailPrimaryKeyMapping)
	sql := "DELETE FROM `Emails` WHERE `party_id`=? AND `email_type`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from Emails")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for Emails")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q emailQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no emailQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from Emails")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Emails")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o EmailSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Email slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(emailBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `Emails` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, emailPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from email slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Emails")
	}

	if len(emailAfterDeleteHooks) != 0 {
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
func (o *Email) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindEmail(ctx, exec, o.PartyID, o.EmailType)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *EmailSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := EmailSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), emailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `Emails`.* FROM `Emails` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, emailPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in EmailSlice")
	}

	*o = slice

	return nil
}

// EmailExists checks if the Email row exists.
func EmailExists(ctx context.Context, exec boil.ContextExecutor, partyID int, emailType string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `Emails` where `party_id`=? AND `email_type`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, partyID, emailType)
	}

	row := exec.QueryRowContext(ctx, sql, partyID, emailType)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if Emails exists")
	}

	return exists, nil
}
