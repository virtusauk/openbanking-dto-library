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

// Notification is an object representing the database table.
type Notification struct {
	MSGID           int         `boil:"msg_id" json:"msg_id" toml:"msg_id" yaml:"msg_id"`
	MessageType     string      `boil:"message_type" json:"message_type" toml:"message_type" yaml:"message_type"`
	Referencenumber string      `boil:"referencenumber" json:"referencenumber" toml:"referencenumber" yaml:"referencenumber"`
	DocumentID      null.Int    `boil:"document_id" json:"document_id,omitempty" toml:"document_id" yaml:"document_id,omitempty"`
	PartyID         null.Int    `boil:"party_id" json:"party_id,omitempty" toml:"party_id" yaml:"party_id,omitempty"`
	Message         null.String `boil:"message" json:"message,omitempty" toml:"message" yaml:"message,omitempty"`
	AlertSendDate   null.Time   `boil:"alert_send_date" json:"alert_send_date,omitempty" toml:"alert_send_date" yaml:"alert_send_date,omitempty"`

	R *notificationR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L notificationL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var NotificationColumns = struct {
	MSGID           string
	MessageType     string
	Referencenumber string
	DocumentID      string
	PartyID         string
	Message         string
	AlertSendDate   string
}{
	MSGID:           "msg_id",
	MessageType:     "message_type",
	Referencenumber: "referencenumber",
	DocumentID:      "document_id",
	PartyID:         "party_id",
	Message:         "message",
	AlertSendDate:   "alert_send_date",
}

// Generated where

var NotificationWhere = struct {
	MSGID           whereHelperint
	MessageType     whereHelperstring
	Referencenumber whereHelperstring
	DocumentID      whereHelpernull_Int
	PartyID         whereHelpernull_Int
	Message         whereHelpernull_String
	AlertSendDate   whereHelpernull_Time
}{
	MSGID:           whereHelperint{field: `msg_id`},
	MessageType:     whereHelperstring{field: `message_type`},
	Referencenumber: whereHelperstring{field: `referencenumber`},
	DocumentID:      whereHelpernull_Int{field: `document_id`},
	PartyID:         whereHelpernull_Int{field: `party_id`},
	Message:         whereHelpernull_String{field: `message`},
	AlertSendDate:   whereHelpernull_Time{field: `alert_send_date`},
}

// NotificationRels is where relationship names are stored.
var NotificationRels = struct {
	Document string
	Party    string
}{
	Document: "Document",
	Party:    "Party",
}

// notificationR is where relationships are stored.
type notificationR struct {
	Document *Document
	Party    *Party
}

// NewStruct creates a new relationship struct
func (*notificationR) NewStruct() *notificationR {
	return &notificationR{}
}

// notificationL is where Load methods for each relationship are stored.
type notificationL struct{}

var (
	notificationColumns               = []string{"msg_id", "message_type", "referencenumber", "document_id", "party_id", "message", "alert_send_date"}
	notificationColumnsWithoutDefault = []string{"msg_id", "message_type", "referencenumber", "document_id", "party_id", "message", "alert_send_date"}
	notificationColumnsWithDefault    = []string{}
	notificationPrimaryKeyColumns     = []string{"msg_id"}
)

type (
	// NotificationSlice is an alias for a slice of pointers to Notification.
	// This should generally be used opposed to []Notification.
	NotificationSlice []*Notification
	// NotificationHook is the signature for custom Notification hook methods
	NotificationHook func(context.Context, boil.ContextExecutor, *Notification) error

	notificationQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	notificationType                 = reflect.TypeOf(&Notification{})
	notificationMapping              = queries.MakeStructMapping(notificationType)
	notificationPrimaryKeyMapping, _ = queries.BindMapping(notificationType, notificationMapping, notificationPrimaryKeyColumns)
	notificationInsertCacheMut       sync.RWMutex
	notificationInsertCache          = make(map[string]insertCache)
	notificationUpdateCacheMut       sync.RWMutex
	notificationUpdateCache          = make(map[string]updateCache)
	notificationUpsertCacheMut       sync.RWMutex
	notificationUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var notificationBeforeInsertHooks []NotificationHook
var notificationBeforeUpdateHooks []NotificationHook
var notificationBeforeDeleteHooks []NotificationHook
var notificationBeforeUpsertHooks []NotificationHook

var notificationAfterInsertHooks []NotificationHook
var notificationAfterSelectHooks []NotificationHook
var notificationAfterUpdateHooks []NotificationHook
var notificationAfterDeleteHooks []NotificationHook
var notificationAfterUpsertHooks []NotificationHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Notification) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range notificationBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Notification) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range notificationBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Notification) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range notificationBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Notification) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range notificationBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Notification) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range notificationAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Notification) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range notificationAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Notification) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range notificationAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Notification) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range notificationAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Notification) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range notificationAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddNotificationHook registers your hook function for all future operations.
func AddNotificationHook(hookPoint boil.HookPoint, notificationHook NotificationHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		notificationBeforeInsertHooks = append(notificationBeforeInsertHooks, notificationHook)
	case boil.BeforeUpdateHook:
		notificationBeforeUpdateHooks = append(notificationBeforeUpdateHooks, notificationHook)
	case boil.BeforeDeleteHook:
		notificationBeforeDeleteHooks = append(notificationBeforeDeleteHooks, notificationHook)
	case boil.BeforeUpsertHook:
		notificationBeforeUpsertHooks = append(notificationBeforeUpsertHooks, notificationHook)
	case boil.AfterInsertHook:
		notificationAfterInsertHooks = append(notificationAfterInsertHooks, notificationHook)
	case boil.AfterSelectHook:
		notificationAfterSelectHooks = append(notificationAfterSelectHooks, notificationHook)
	case boil.AfterUpdateHook:
		notificationAfterUpdateHooks = append(notificationAfterUpdateHooks, notificationHook)
	case boil.AfterDeleteHook:
		notificationAfterDeleteHooks = append(notificationAfterDeleteHooks, notificationHook)
	case boil.AfterUpsertHook:
		notificationAfterUpsertHooks = append(notificationAfterUpsertHooks, notificationHook)
	}
}

// One returns a single notification record from the query.
func (q notificationQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Notification, error) {
	o := &Notification{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for Notifications")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Notification records from the query.
func (q notificationQuery) All(ctx context.Context, exec boil.ContextExecutor) (NotificationSlice, error) {
	var o []*Notification

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Notification slice")
	}

	if len(notificationAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Notification records in the query.
func (q notificationQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count Notifications rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q notificationQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if Notifications exists")
	}

	return count > 0, nil
}

// Document pointed to by the foreign key.
func (o *Notification) Document(mods ...qm.QueryMod) documentQuery {
	queryMods := []qm.QueryMod{
		qm.Where("document_id=?", o.DocumentID),
	}

	queryMods = append(queryMods, mods...)

	query := Documents(queryMods...)
	queries.SetFrom(query.Query, "`Document`")

	return query
}

// Party pointed to by the foreign key.
func (o *Notification) Party(mods ...qm.QueryMod) partyQuery {
	queryMods := []qm.QueryMod{
		qm.Where("party_id=?", o.PartyID),
	}

	queryMods = append(queryMods, mods...)

	query := Parties(queryMods...)
	queries.SetFrom(query.Query, "`Parties`")

	return query
}

// LoadDocument allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (notificationL) LoadDocument(ctx context.Context, e boil.ContextExecutor, singular bool, maybeNotification interface{}, mods queries.Applicator) error {
	var slice []*Notification
	var object *Notification

	if singular {
		object = maybeNotification.(*Notification)
	} else {
		slice = *maybeNotification.(*[]*Notification)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &notificationR{}
		}
		if !queries.IsNil(object.DocumentID) {
			args = append(args, object.DocumentID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &notificationR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.DocumentID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.DocumentID) {
				args = append(args, obj.DocumentID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`Document`), qm.WhereIn(`document_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Document")
	}

	var resultSlice []*Document
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Document")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for Document")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for Document")
	}

	if len(notificationAfterSelectHooks) != 0 {
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
		object.R.Document = foreign
		if foreign.R == nil {
			foreign.R = &documentR{}
		}
		foreign.R.DocumentNotifications = append(foreign.R.DocumentNotifications, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.DocumentID, foreign.DocumentID) {
				local.R.Document = foreign
				if foreign.R == nil {
					foreign.R = &documentR{}
				}
				foreign.R.DocumentNotifications = append(foreign.R.DocumentNotifications, local)
				break
			}
		}
	}

	return nil
}

// LoadParty allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (notificationL) LoadParty(ctx context.Context, e boil.ContextExecutor, singular bool, maybeNotification interface{}, mods queries.Applicator) error {
	var slice []*Notification
	var object *Notification

	if singular {
		object = maybeNotification.(*Notification)
	} else {
		slice = *maybeNotification.(*[]*Notification)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &notificationR{}
		}
		if !queries.IsNil(object.PartyID) {
			args = append(args, object.PartyID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &notificationR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.PartyID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.PartyID) {
				args = append(args, obj.PartyID)
			}

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

	if len(notificationAfterSelectHooks) != 0 {
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
		foreign.R.PartyNotifications = append(foreign.R.PartyNotifications, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.PartyID, foreign.PartyID) {
				local.R.Party = foreign
				if foreign.R == nil {
					foreign.R = &partyR{}
				}
				foreign.R.PartyNotifications = append(foreign.R.PartyNotifications, local)
				break
			}
		}
	}

	return nil
}

// SetDocument of the notification to the related item.
// Sets o.R.Document to related.
// Adds o to related.R.DocumentNotifications.
func (o *Notification) SetDocument(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Document) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `Notifications` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"document_id"}),
		strmangle.WhereClause("`", "`", 0, notificationPrimaryKeyColumns),
	)
	values := []interface{}{related.DocumentID, o.MSGID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.DocumentID, related.DocumentID)
	if o.R == nil {
		o.R = &notificationR{
			Document: related,
		}
	} else {
		o.R.Document = related
	}

	if related.R == nil {
		related.R = &documentR{
			DocumentNotifications: NotificationSlice{o},
		}
	} else {
		related.R.DocumentNotifications = append(related.R.DocumentNotifications, o)
	}

	return nil
}

// RemoveDocument relationship.
// Sets o.R.Document to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Notification) RemoveDocument(ctx context.Context, exec boil.ContextExecutor, related *Document) error {
	var err error

	queries.SetScanner(&o.DocumentID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("document_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Document = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.DocumentNotifications {
		if queries.Equal(o.DocumentID, ri.DocumentID) {
			continue
		}

		ln := len(related.R.DocumentNotifications)
		if ln > 1 && i < ln-1 {
			related.R.DocumentNotifications[i] = related.R.DocumentNotifications[ln-1]
		}
		related.R.DocumentNotifications = related.R.DocumentNotifications[:ln-1]
		break
	}
	return nil
}

// SetParty of the notification to the related item.
// Sets o.R.Party to related.
// Adds o to related.R.PartyNotifications.
func (o *Notification) SetParty(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Party) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `Notifications` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"party_id"}),
		strmangle.WhereClause("`", "`", 0, notificationPrimaryKeyColumns),
	)
	values := []interface{}{related.PartyID, o.MSGID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.PartyID, related.PartyID)
	if o.R == nil {
		o.R = &notificationR{
			Party: related,
		}
	} else {
		o.R.Party = related
	}

	if related.R == nil {
		related.R = &partyR{
			PartyNotifications: NotificationSlice{o},
		}
	} else {
		related.R.PartyNotifications = append(related.R.PartyNotifications, o)
	}

	return nil
}

// RemoveParty relationship.
// Sets o.R.Party to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Notification) RemoveParty(ctx context.Context, exec boil.ContextExecutor, related *Party) error {
	var err error

	queries.SetScanner(&o.PartyID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("party_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.Party = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.PartyNotifications {
		if queries.Equal(o.PartyID, ri.PartyID) {
			continue
		}

		ln := len(related.R.PartyNotifications)
		if ln > 1 && i < ln-1 {
			related.R.PartyNotifications[i] = related.R.PartyNotifications[ln-1]
		}
		related.R.PartyNotifications = related.R.PartyNotifications[:ln-1]
		break
	}
	return nil
}

// Notifications retrieves all the records using an executor.
func Notifications(mods ...qm.QueryMod) notificationQuery {
	mods = append(mods, qm.From("`Notifications`"))
	return notificationQuery{NewQuery(mods...)}
}

// FindNotification retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindNotification(ctx context.Context, exec boil.ContextExecutor, mSGID int, selectCols ...string) (*Notification, error) {
	notificationObj := &Notification{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `Notifications` where `msg_id`=?", sel,
	)

	q := queries.Raw(query, mSGID)

	err := q.Bind(ctx, exec, notificationObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from Notifications")
	}

	return notificationObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Notification) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Notifications provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(notificationColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	notificationInsertCacheMut.RLock()
	cache, cached := notificationInsertCache[key]
	notificationInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			notificationColumns,
			notificationColumnsWithDefault,
			notificationColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(notificationType, notificationMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(notificationType, notificationMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `Notifications` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `Notifications` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `Notifications` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, notificationPrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into Notifications")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.MSGID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for Notifications")
	}

CacheNoHooks:
	if !cached {
		notificationInsertCacheMut.Lock()
		notificationInsertCache[key] = cache
		notificationInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Notification.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Notification) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	notificationUpdateCacheMut.RLock()
	cache, cached := notificationUpdateCache[key]
	notificationUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			notificationColumns,
			notificationPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update Notifications, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `Notifications` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, notificationPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(notificationType, notificationMapping, append(wl, notificationPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update Notifications row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for Notifications")
	}

	if !cached {
		notificationUpdateCacheMut.Lock()
		notificationUpdateCache[key] = cache
		notificationUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q notificationQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for Notifications")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for Notifications")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o NotificationSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), notificationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `Notifications` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, notificationPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in notification slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all notification")
	}
	return rowsAff, nil
}

var mySQLNotificationUniqueColumns = []string{
	"msg_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Notification) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Notifications provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(notificationColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLNotificationUniqueColumns, o)

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

	notificationUpsertCacheMut.RLock()
	cache, cached := notificationUpsertCache[key]
	notificationUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			notificationColumns,
			notificationColumnsWithDefault,
			notificationColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			notificationColumns,
			notificationPrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert Notifications, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "Notifications", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `Notifications` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(notificationType, notificationMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(notificationType, notificationMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for Notifications")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(notificationType, notificationMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for Notifications")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for Notifications")
	}

CacheNoHooks:
	if !cached {
		notificationUpsertCacheMut.Lock()
		notificationUpsertCache[key] = cache
		notificationUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Notification record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Notification) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Notification provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), notificationPrimaryKeyMapping)
	sql := "DELETE FROM `Notifications` WHERE `msg_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from Notifications")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for Notifications")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q notificationQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no notificationQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from Notifications")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Notifications")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o NotificationSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Notification slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(notificationBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), notificationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `Notifications` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, notificationPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from notification slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Notifications")
	}

	if len(notificationAfterDeleteHooks) != 0 {
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
func (o *Notification) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindNotification(ctx, exec, o.MSGID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *NotificationSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := NotificationSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), notificationPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `Notifications`.* FROM `Notifications` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, notificationPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in NotificationSlice")
	}

	*o = slice

	return nil
}

// NotificationExists checks if the Notification row exists.
func NotificationExists(ctx context.Context, exec boil.ContextExecutor, mSGID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `Notifications` where `msg_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, mSGID)
	}

	row := exec.QueryRowContext(ctx, sql, mSGID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if Notifications exists")
	}

	return exists, nil
}
