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

// Application is an object representing the database table.
type Application struct {
	UserID                      null.String `boil:"user_id" json:"user_id,omitempty" toml:"user_id" yaml:"user_id,omitempty"`
	ClientName                  string      `boil:"client_name" json:"client_name" toml:"client_name" yaml:"client_name"`
	ClientID                    string      `boil:"client_id" json:"client_id" toml:"client_id" yaml:"client_id"`
	ClientSecret                null.String `boil:"client_secret" json:"client_secret,omitempty" toml:"client_secret" yaml:"client_secret,omitempty"`
	Description                 null.String `boil:"description" json:"description,omitempty" toml:"description" yaml:"description,omitempty"`
	Iss                         string      `boil:"iss" json:"iss" toml:"iss" yaml:"iss"`
	Aud                         string      `boil:"aud" json:"aud" toml:"aud" yaml:"aud"`
	Organization                null.String `boil:"organization" json:"organization,omitempty" toml:"organization" yaml:"organization,omitempty"`
	Scope                       string      `boil:"scope" json:"scope" toml:"scope" yaml:"scope"`
	TokenEndpointAuthMethod     string      `boil:"token_endpoint_auth_method" json:"token_endpoint_auth_method" toml:"token_endpoint_auth_method" yaml:"token_endpoint_auth_method"`
	GrantTypes                  string      `boil:"grant_types" json:"grant_types" toml:"grant_types" yaml:"grant_types"`
	ResponseTypes               string      `boil:"response_types" json:"response_types" toml:"response_types" yaml:"response_types"`
	RedirectUris                string      `boil:"redirect_uris" json:"redirect_uris" toml:"redirect_uris" yaml:"redirect_uris"`
	IDTokenSignedResponseAlg    string      `boil:"id_token_signed_response_alg" json:"id_token_signed_response_alg" toml:"id_token_signed_response_alg" yaml:"id_token_signed_response_alg"`
	RequestObjectSigningAlg     string      `boil:"request_object_signing_alg" json:"request_object_signing_alg" toml:"request_object_signing_alg" yaml:"request_object_signing_alg"`
	SoftwareID                  string      `boil:"software_id" json:"software_id" toml:"software_id" yaml:"software_id"`
	SoftwareStatement           string      `boil:"software_statement" json:"software_statement" toml:"software_statement" yaml:"software_statement"`
	Kid                         null.String `boil:"kid" json:"kid,omitempty" toml:"kid" yaml:"kid,omitempty"`
	JWTHeader                   null.String `boil:"jwt_header" json:"jwt_header,omitempty" toml:"jwt_header" yaml:"jwt_header,omitempty"`
	SoftwareJWKSEndpoint        string      `boil:"software_jwks_endpoint" json:"software_jwks_endpoint" toml:"software_jwks_endpoint" yaml:"software_jwks_endpoint"`
	SoftwareJWKSRevokedEndpoint string      `boil:"software_jwks_revoked_endpoint" json:"software_jwks_revoked_endpoint" toml:"software_jwks_revoked_endpoint" yaml:"software_jwks_revoked_endpoint"`
	CreatedDate                 time.Time   `boil:"created_date" json:"created_date" toml:"created_date" yaml:"created_date"`

	R *applicationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L applicationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ApplicationColumns = struct {
	UserID                      string
	ClientName                  string
	ClientID                    string
	ClientSecret                string
	Description                 string
	Iss                         string
	Aud                         string
	Organization                string
	Scope                       string
	TokenEndpointAuthMethod     string
	GrantTypes                  string
	ResponseTypes               string
	RedirectUris                string
	IDTokenSignedResponseAlg    string
	RequestObjectSigningAlg     string
	SoftwareID                  string
	SoftwareStatement           string
	Kid                         string
	JWTHeader                   string
	SoftwareJWKSEndpoint        string
	SoftwareJWKSRevokedEndpoint string
	CreatedDate                 string
}{
	UserID:                      "user_id",
	ClientName:                  "client_name",
	ClientID:                    "client_id",
	ClientSecret:                "client_secret",
	Description:                 "description",
	Iss:                         "iss",
	Aud:                         "aud",
	Organization:                "organization",
	Scope:                       "scope",
	TokenEndpointAuthMethod:     "token_endpoint_auth_method",
	GrantTypes:                  "grant_types",
	ResponseTypes:               "response_types",
	RedirectUris:                "redirect_uris",
	IDTokenSignedResponseAlg:    "id_token_signed_response_alg",
	RequestObjectSigningAlg:     "request_object_signing_alg",
	SoftwareID:                  "software_id",
	SoftwareStatement:           "software_statement",
	Kid:                         "kid",
	JWTHeader:                   "jwt_header",
	SoftwareJWKSEndpoint:        "software_jwks_endpoint",
	SoftwareJWKSRevokedEndpoint: "software_jwks_revoked_endpoint",
	CreatedDate:                 "created_date",
}

// Generated where

var ApplicationWhere = struct {
	UserID                      whereHelpernull_String
	ClientName                  whereHelperstring
	ClientID                    whereHelperstring
	ClientSecret                whereHelpernull_String
	Description                 whereHelpernull_String
	Iss                         whereHelperstring
	Aud                         whereHelperstring
	Organization                whereHelpernull_String
	Scope                       whereHelperstring
	TokenEndpointAuthMethod     whereHelperstring
	GrantTypes                  whereHelperstring
	ResponseTypes               whereHelperstring
	RedirectUris                whereHelperstring
	IDTokenSignedResponseAlg    whereHelperstring
	RequestObjectSigningAlg     whereHelperstring
	SoftwareID                  whereHelperstring
	SoftwareStatement           whereHelperstring
	Kid                         whereHelpernull_String
	JWTHeader                   whereHelpernull_String
	SoftwareJWKSEndpoint        whereHelperstring
	SoftwareJWKSRevokedEndpoint whereHelperstring
	CreatedDate                 whereHelpertime_Time
}{
	UserID:                      whereHelpernull_String{field: `user_id`},
	ClientName:                  whereHelperstring{field: `client_name`},
	ClientID:                    whereHelperstring{field: `client_id`},
	ClientSecret:                whereHelpernull_String{field: `client_secret`},
	Description:                 whereHelpernull_String{field: `description`},
	Iss:                         whereHelperstring{field: `iss`},
	Aud:                         whereHelperstring{field: `aud`},
	Organization:                whereHelpernull_String{field: `organization`},
	Scope:                       whereHelperstring{field: `scope`},
	TokenEndpointAuthMethod:     whereHelperstring{field: `token_endpoint_auth_method`},
	GrantTypes:                  whereHelperstring{field: `grant_types`},
	ResponseTypes:               whereHelperstring{field: `response_types`},
	RedirectUris:                whereHelperstring{field: `redirect_uris`},
	IDTokenSignedResponseAlg:    whereHelperstring{field: `id_token_signed_response_alg`},
	RequestObjectSigningAlg:     whereHelperstring{field: `request_object_signing_alg`},
	SoftwareID:                  whereHelperstring{field: `software_id`},
	SoftwareStatement:           whereHelperstring{field: `software_statement`},
	Kid:                         whereHelpernull_String{field: `kid`},
	JWTHeader:                   whereHelpernull_String{field: `jwt_header`},
	SoftwareJWKSEndpoint:        whereHelperstring{field: `software_jwks_endpoint`},
	SoftwareJWKSRevokedEndpoint: whereHelperstring{field: `software_jwks_revoked_endpoint`},
	CreatedDate:                 whereHelpertime_Time{field: `created_date`},
}

// ApplicationRels is where relationship names are stored.
var ApplicationRels = struct {
}{}

// applicationR is where relationships are stored.
type applicationR struct {
}

// NewStruct creates a new relationship struct
func (*applicationR) NewStruct() *applicationR {
	return &applicationR{}
}

// applicationL is where Load methods for each relationship are stored.
type applicationL struct{}

var (
	applicationColumns               = []string{"user_id", "client_name", "client_id", "client_secret", "description", "iss", "aud", "organization", "scope", "token_endpoint_auth_method", "grant_types", "response_types", "redirect_uris", "id_token_signed_response_alg", "request_object_signing_alg", "software_id", "software_statement", "kid", "jwt_header", "software_jwks_endpoint", "software_jwks_revoked_endpoint", "created_date"}
	applicationColumnsWithoutDefault = []string{"user_id", "client_name", "client_id", "client_secret", "description", "iss", "aud", "organization", "scope", "token_endpoint_auth_method", "grant_types", "response_types", "redirect_uris", "id_token_signed_response_alg", "request_object_signing_alg", "software_id", "software_statement", "kid", "jwt_header", "software_jwks_endpoint", "software_jwks_revoked_endpoint", "created_date"}
	applicationColumnsWithDefault    = []string{}
	applicationPrimaryKeyColumns     = []string{"client_id"}
)

type (
	// ApplicationSlice is an alias for a slice of pointers to Application.
	// This should generally be used opposed to []Application.
	ApplicationSlice []*Application
	// ApplicationHook is the signature for custom Application hook methods
	ApplicationHook func(context.Context, boil.ContextExecutor, *Application) error

	applicationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	applicationType                 = reflect.TypeOf(&Application{})
	applicationMapping              = queries.MakeStructMapping(applicationType)
	applicationPrimaryKeyMapping, _ = queries.BindMapping(applicationType, applicationMapping, applicationPrimaryKeyColumns)
	applicationInsertCacheMut       sync.RWMutex
	applicationInsertCache          = make(map[string]insertCache)
	applicationUpdateCacheMut       sync.RWMutex
	applicationUpdateCache          = make(map[string]updateCache)
	applicationUpsertCacheMut       sync.RWMutex
	applicationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var applicationBeforeInsertHooks []ApplicationHook
var applicationBeforeUpdateHooks []ApplicationHook
var applicationBeforeDeleteHooks []ApplicationHook
var applicationBeforeUpsertHooks []ApplicationHook

var applicationAfterInsertHooks []ApplicationHook
var applicationAfterSelectHooks []ApplicationHook
var applicationAfterUpdateHooks []ApplicationHook
var applicationAfterDeleteHooks []ApplicationHook
var applicationAfterUpsertHooks []ApplicationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Application) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Application) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Application) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Application) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Application) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Application) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Application) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Application) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Application) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range applicationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddApplicationHook registers your hook function for all future operations.
func AddApplicationHook(hookPoint boil.HookPoint, applicationHook ApplicationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		applicationBeforeInsertHooks = append(applicationBeforeInsertHooks, applicationHook)
	case boil.BeforeUpdateHook:
		applicationBeforeUpdateHooks = append(applicationBeforeUpdateHooks, applicationHook)
	case boil.BeforeDeleteHook:
		applicationBeforeDeleteHooks = append(applicationBeforeDeleteHooks, applicationHook)
	case boil.BeforeUpsertHook:
		applicationBeforeUpsertHooks = append(applicationBeforeUpsertHooks, applicationHook)
	case boil.AfterInsertHook:
		applicationAfterInsertHooks = append(applicationAfterInsertHooks, applicationHook)
	case boil.AfterSelectHook:
		applicationAfterSelectHooks = append(applicationAfterSelectHooks, applicationHook)
	case boil.AfterUpdateHook:
		applicationAfterUpdateHooks = append(applicationAfterUpdateHooks, applicationHook)
	case boil.AfterDeleteHook:
		applicationAfterDeleteHooks = append(applicationAfterDeleteHooks, applicationHook)
	case boil.AfterUpsertHook:
		applicationAfterUpsertHooks = append(applicationAfterUpsertHooks, applicationHook)
	}
}

// One returns a single application record from the query.
func (q applicationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Application, error) {
	o := &Application{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for Application")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Application records from the query.
func (q applicationQuery) All(ctx context.Context, exec boil.ContextExecutor) (ApplicationSlice, error) {
	var o []*Application

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Application slice")
	}

	if len(applicationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Application records in the query.
func (q applicationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count Application rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q applicationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if Application exists")
	}

	return count > 0, nil
}

// Applications retrieves all the records using an executor.
func Applications(mods ...qm.QueryMod) applicationQuery {
	mods = append(mods, qm.From("`Application`"))
	return applicationQuery{NewQuery(mods...)}
}

// FindApplication retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindApplication(ctx context.Context, exec boil.ContextExecutor, clientID string, selectCols ...string) (*Application, error) {
	applicationObj := &Application{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `Application` where `client_id`=?", sel,
	)

	q := queries.Raw(query, clientID)

	err := q.Bind(ctx, exec, applicationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from Application")
	}

	return applicationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Application) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Application provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(applicationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	applicationInsertCacheMut.RLock()
	cache, cached := applicationInsertCache[key]
	applicationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			applicationColumns,
			applicationColumnsWithDefault,
			applicationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(applicationType, applicationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(applicationType, applicationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `Application` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `Application` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `Application` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, applicationPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into Application")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ClientID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for Application")
	}

CacheNoHooks:
	if !cached {
		applicationInsertCacheMut.Lock()
		applicationInsertCache[key] = cache
		applicationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Application.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Application) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	applicationUpdateCacheMut.RLock()
	cache, cached := applicationUpdateCache[key]
	applicationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			applicationColumns,
			applicationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update Application, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `Application` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, applicationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(applicationType, applicationMapping, append(wl, applicationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update Application row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for Application")
	}

	if !cached {
		applicationUpdateCacheMut.Lock()
		applicationUpdateCache[key] = cache
		applicationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q applicationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for Application")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for Application")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ApplicationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), applicationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `Application` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, applicationPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in application slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all application")
	}
	return rowsAff, nil
}

var mySQLApplicationUniqueColumns = []string{
	"client_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Application) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Application provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(applicationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLApplicationUniqueColumns, o)

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

	applicationUpsertCacheMut.RLock()
	cache, cached := applicationUpsertCache[key]
	applicationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			applicationColumns,
			applicationColumnsWithDefault,
			applicationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			applicationColumns,
			applicationPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert Application, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "Application", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `Application` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(applicationType, applicationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(applicationType, applicationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for Application")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(applicationType, applicationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for Application")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for Application")
	}

CacheNoHooks:
	if !cached {
		applicationUpsertCacheMut.Lock()
		applicationUpsertCache[key] = cache
		applicationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Application record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Application) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Application provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), applicationPrimaryKeyMapping)
	sql := "DELETE FROM `Application` WHERE `client_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from Application")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for Application")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q applicationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no applicationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from Application")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Application")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ApplicationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Application slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(applicationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), applicationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `Application` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, applicationPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from application slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Application")
	}

	if len(applicationAfterDeleteHooks) != 0 {
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
func (o *Application) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindApplication(ctx, exec, o.ClientID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ApplicationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ApplicationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), applicationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `Application`.* FROM `Application` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, applicationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ApplicationSlice")
	}

	*o = slice

	return nil
}

// ApplicationExists checks if the Application row exists.
func ApplicationExists(ctx context.Context, exec boil.ContextExecutor, clientID string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `Application` where `client_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, clientID)
	}

	row := exec.QueryRowContext(ctx, sql, clientID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if Application exists")
	}

	return exists, nil
}
