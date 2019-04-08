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

// Resource is an object representing the database table.
type Resource struct {
	ResourceID     int         `boil:"resource_id" json:"resource_id" toml:"resource_id" yaml:"resource_id"`
	ResourceName   null.String `boil:"resource_name" json:"resource_name,omitempty" toml:"resource_name" yaml:"resource_name,omitempty"`
	ResourceTypeID null.Int    `boil:"resource_type_id" json:"resource_type_id,omitempty" toml:"resource_type_id" yaml:"resource_type_id,omitempty"`
	ActionType     null.String `boil:"action_type" json:"action_type,omitempty" toml:"action_type" yaml:"action_type,omitempty"`
	Enabled        null.String `boil:"enabled" json:"enabled,omitempty" toml:"enabled" yaml:"enabled,omitempty"`
	MakerDate      time.Time   `boil:"maker_date" json:"maker_date" toml:"maker_date" yaml:"maker_date"`
	CheckerDate    null.Time   `boil:"checker_date" json:"checker_date,omitempty" toml:"checker_date" yaml:"checker_date,omitempty"`
	MakerID        string      `boil:"maker_id" json:"maker_id" toml:"maker_id" yaml:"maker_id"`
	CheckerID      null.String `boil:"checker_id" json:"checker_id,omitempty" toml:"checker_id" yaml:"checker_id,omitempty"`
	ModifiedBy     null.String `boil:"modified_by" json:"modified_by,omitempty" toml:"modified_by" yaml:"modified_by,omitempty"`
	ModifiedDate   null.Time   `boil:"modified_date" json:"modified_date,omitempty" toml:"modified_date" yaml:"modified_date,omitempty"`

	R *resourceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L resourceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var ResourceColumns = struct {
	ResourceID     string
	ResourceName   string
	ResourceTypeID string
	ActionType     string
	Enabled        string
	MakerDate      string
	CheckerDate    string
	MakerID        string
	CheckerID      string
	ModifiedBy     string
	ModifiedDate   string
}{
	ResourceID:     "resource_id",
	ResourceName:   "resource_name",
	ResourceTypeID: "resource_type_id",
	ActionType:     "action_type",
	Enabled:        "enabled",
	MakerDate:      "maker_date",
	CheckerDate:    "checker_date",
	MakerID:        "maker_id",
	CheckerID:      "checker_id",
	ModifiedBy:     "modified_by",
	ModifiedDate:   "modified_date",
}

// Generated where

var ResourceWhere = struct {
	ResourceID     whereHelperint
	ResourceName   whereHelpernull_String
	ResourceTypeID whereHelpernull_Int
	ActionType     whereHelpernull_String
	Enabled        whereHelpernull_String
	MakerDate      whereHelpertime_Time
	CheckerDate    whereHelpernull_Time
	MakerID        whereHelperstring
	CheckerID      whereHelpernull_String
	ModifiedBy     whereHelpernull_String
	ModifiedDate   whereHelpernull_Time
}{
	ResourceID:     whereHelperint{field: `resource_id`},
	ResourceName:   whereHelpernull_String{field: `resource_name`},
	ResourceTypeID: whereHelpernull_Int{field: `resource_type_id`},
	ActionType:     whereHelpernull_String{field: `action_type`},
	Enabled:        whereHelpernull_String{field: `enabled`},
	MakerDate:      whereHelpertime_Time{field: `maker_date`},
	CheckerDate:    whereHelpernull_Time{field: `checker_date`},
	MakerID:        whereHelperstring{field: `maker_id`},
	CheckerID:      whereHelpernull_String{field: `checker_id`},
	ModifiedBy:     whereHelpernull_String{field: `modified_by`},
	ModifiedDate:   whereHelpernull_Time{field: `modified_date`},
}

// ResourceRels is where relationship names are stored.
var ResourceRels = struct {
	ResourceType          string
	ResourceRoleResources string
}{
	ResourceType:          "ResourceType",
	ResourceRoleResources: "ResourceRoleResources",
}

// resourceR is where relationships are stored.
type resourceR struct {
	ResourceType          *ResourcesType
	ResourceRoleResources RoleResourceSlice
}

// NewStruct creates a new relationship struct
func (*resourceR) NewStruct() *resourceR {
	return &resourceR{}
}

// resourceL is where Load methods for each relationship are stored.
type resourceL struct{}

var (
	resourceColumns               = []string{"resource_id", "resource_name", "resource_type_id", "action_type", "enabled", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	resourceColumnsWithoutDefault = []string{"resource_id", "resource_name", "resource_type_id", "action_type", "enabled", "maker_date", "checker_date", "maker_id", "checker_id", "modified_by", "modified_date"}
	resourceColumnsWithDefault    = []string{}
	resourcePrimaryKeyColumns     = []string{"resource_id"}
)

type (
	// ResourceSlice is an alias for a slice of pointers to Resource.
	// This should generally be used opposed to []Resource.
	ResourceSlice []*Resource
	// ResourceHook is the signature for custom Resource hook methods
	ResourceHook func(context.Context, boil.ContextExecutor, *Resource) error

	resourceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	resourceType                 = reflect.TypeOf(&Resource{})
	resourceMapping              = queries.MakeStructMapping(resourceType)
	resourcePrimaryKeyMapping, _ = queries.BindMapping(resourceType, resourceMapping, resourcePrimaryKeyColumns)
	resourceInsertCacheMut       sync.RWMutex
	resourceInsertCache          = make(map[string]insertCache)
	resourceUpdateCacheMut       sync.RWMutex
	resourceUpdateCache          = make(map[string]updateCache)
	resourceUpsertCacheMut       sync.RWMutex
	resourceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

var resourceBeforeInsertHooks []ResourceHook
var resourceBeforeUpdateHooks []ResourceHook
var resourceBeforeDeleteHooks []ResourceHook
var resourceBeforeUpsertHooks []ResourceHook

var resourceAfterInsertHooks []ResourceHook
var resourceAfterSelectHooks []ResourceHook
var resourceAfterUpdateHooks []ResourceHook
var resourceAfterDeleteHooks []ResourceHook
var resourceAfterUpsertHooks []ResourceHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Resource) doBeforeInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range resourceBeforeInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Resource) doBeforeUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range resourceBeforeUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Resource) doBeforeDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range resourceBeforeDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Resource) doBeforeUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range resourceBeforeUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Resource) doAfterInsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range resourceAfterInsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Resource) doAfterSelectHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range resourceAfterSelectHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Resource) doAfterUpdateHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range resourceAfterUpdateHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Resource) doAfterDeleteHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range resourceAfterDeleteHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Resource) doAfterUpsertHooks(ctx context.Context, exec boil.ContextExecutor) (err error) {
	if boil.HooksAreSkipped(ctx) {
		return nil
	}

	for _, hook := range resourceAfterUpsertHooks {
		if err := hook(ctx, exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddResourceHook registers your hook function for all future operations.
func AddResourceHook(hookPoint boil.HookPoint, resourceHook ResourceHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		resourceBeforeInsertHooks = append(resourceBeforeInsertHooks, resourceHook)
	case boil.BeforeUpdateHook:
		resourceBeforeUpdateHooks = append(resourceBeforeUpdateHooks, resourceHook)
	case boil.BeforeDeleteHook:
		resourceBeforeDeleteHooks = append(resourceBeforeDeleteHooks, resourceHook)
	case boil.BeforeUpsertHook:
		resourceBeforeUpsertHooks = append(resourceBeforeUpsertHooks, resourceHook)
	case boil.AfterInsertHook:
		resourceAfterInsertHooks = append(resourceAfterInsertHooks, resourceHook)
	case boil.AfterSelectHook:
		resourceAfterSelectHooks = append(resourceAfterSelectHooks, resourceHook)
	case boil.AfterUpdateHook:
		resourceAfterUpdateHooks = append(resourceAfterUpdateHooks, resourceHook)
	case boil.AfterDeleteHook:
		resourceAfterDeleteHooks = append(resourceAfterDeleteHooks, resourceHook)
	case boil.AfterUpsertHook:
		resourceAfterUpsertHooks = append(resourceAfterUpsertHooks, resourceHook)
	}
}

// One returns a single resource record from the query.
func (q resourceQuery) One(ctx context.Context, exec boil.ContextExecutor) (*Resource, error) {
	o := &Resource{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for Resources")
	}

	if err := o.doAfterSelectHooks(ctx, exec); err != nil {
		return o, err
	}

	return o, nil
}

// All returns all Resource records from the query.
func (q resourceQuery) All(ctx context.Context, exec boil.ContextExecutor) (ResourceSlice, error) {
	var o []*Resource

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Resource slice")
	}

	if len(resourceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(ctx, exec); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// Count returns the count of all Resource records in the query.
func (q resourceQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count Resources rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q resourceQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if Resources exists")
	}

	return count > 0, nil
}

// ResourceType pointed to by the foreign key.
func (o *Resource) ResourceType(mods ...qm.QueryMod) resourcesTypeQuery {
	queryMods := []qm.QueryMod{
		qm.Where("Resource_type_id=?", o.ResourceTypeID),
	}

	queryMods = append(queryMods, mods...)

	query := ResourcesTypes(queryMods...)
	queries.SetFrom(query.Query, "`ResourcesType`")

	return query
}

// ResourceRoleResources retrieves all the RoleResource's RoleResources with an executor via resource_id column.
func (o *Resource) ResourceRoleResources(mods ...qm.QueryMod) roleResourceQuery {
	var queryMods []qm.QueryMod
	if len(mods) != 0 {
		queryMods = append(queryMods, mods...)
	}

	queryMods = append(queryMods,
		qm.Where("`RoleResources`.`resource_id`=?", o.ResourceID),
	)

	query := RoleResources(queryMods...)
	queries.SetFrom(query.Query, "`RoleResources`")

	if len(queries.GetSelect(query.Query)) == 0 {
		queries.SetSelect(query.Query, []string{"`RoleResources`.*"})
	}

	return query
}

// LoadResourceType allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (resourceL) LoadResourceType(ctx context.Context, e boil.ContextExecutor, singular bool, maybeResource interface{}, mods queries.Applicator) error {
	var slice []*Resource
	var object *Resource

	if singular {
		object = maybeResource.(*Resource)
	} else {
		slice = *maybeResource.(*[]*Resource)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &resourceR{}
		}
		if !queries.IsNil(object.ResourceTypeID) {
			args = append(args, object.ResourceTypeID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &resourceR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.ResourceTypeID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.ResourceTypeID) {
				args = append(args, obj.ResourceTypeID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`ResourcesType`), qm.WhereIn(`Resource_type_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load ResourcesType")
	}

	var resultSlice []*ResourcesType
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice ResourcesType")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for ResourcesType")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for ResourcesType")
	}

	if len(resourceAfterSelectHooks) != 0 {
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
		object.R.ResourceType = foreign
		if foreign.R == nil {
			foreign.R = &resourcesTypeR{}
		}
		foreign.R.ResourceTypeResources = append(foreign.R.ResourceTypeResources, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.ResourceTypeID, foreign.ResourceTypeID) {
				local.R.ResourceType = foreign
				if foreign.R == nil {
					foreign.R = &resourcesTypeR{}
				}
				foreign.R.ResourceTypeResources = append(foreign.R.ResourceTypeResources, local)
				break
			}
		}
	}

	return nil
}

// LoadResourceRoleResources allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for a 1-M or N-M relationship.
func (resourceL) LoadResourceRoleResources(ctx context.Context, e boil.ContextExecutor, singular bool, maybeResource interface{}, mods queries.Applicator) error {
	var slice []*Resource
	var object *Resource

	if singular {
		object = maybeResource.(*Resource)
	} else {
		slice = *maybeResource.(*[]*Resource)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &resourceR{}
		}
		args = append(args, object.ResourceID)
	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &resourceR{}
			}

			for _, a := range args {
				if a == obj.ResourceID {
					continue Outer
				}
			}

			args = append(args, obj.ResourceID)
		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(qm.From(`RoleResources`), qm.WhereIn(`resource_id in ?`, args...))
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load RoleResources")
	}

	var resultSlice []*RoleResource
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice RoleResources")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results in eager load on RoleResources")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for RoleResources")
	}

	if len(roleResourceAfterSelectHooks) != 0 {
		for _, obj := range resultSlice {
			if err := obj.doAfterSelectHooks(ctx, e); err != nil {
				return err
			}
		}
	}
	if singular {
		object.R.ResourceRoleResources = resultSlice
		for _, foreign := range resultSlice {
			if foreign.R == nil {
				foreign.R = &roleResourceR{}
			}
			foreign.R.Resource = object
		}
		return nil
	}

	for _, foreign := range resultSlice {
		for _, local := range slice {
			if local.ResourceID == foreign.ResourceID {
				local.R.ResourceRoleResources = append(local.R.ResourceRoleResources, foreign)
				if foreign.R == nil {
					foreign.R = &roleResourceR{}
				}
				foreign.R.Resource = local
				break
			}
		}
	}

	return nil
}

// SetResourceType of the resource to the related item.
// Sets o.R.ResourceType to related.
// Adds o to related.R.ResourceTypeResources.
func (o *Resource) SetResourceType(ctx context.Context, exec boil.ContextExecutor, insert bool, related *ResourcesType) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE `Resources` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, []string{"resource_type_id"}),
		strmangle.WhereClause("`", "`", 0, resourcePrimaryKeyColumns),
	)
	values := []interface{}{related.ResourceTypeID, o.ResourceID}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, updateQuery)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.ResourceTypeID, related.ResourceTypeID)
	if o.R == nil {
		o.R = &resourceR{
			ResourceType: related,
		}
	} else {
		o.R.ResourceType = related
	}

	if related.R == nil {
		related.R = &resourcesTypeR{
			ResourceTypeResources: ResourceSlice{o},
		}
	} else {
		related.R.ResourceTypeResources = append(related.R.ResourceTypeResources, o)
	}

	return nil
}

// RemoveResourceType relationship.
// Sets o.R.ResourceType to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *Resource) RemoveResourceType(ctx context.Context, exec boil.ContextExecutor, related *ResourcesType) error {
	var err error

	queries.SetScanner(&o.ResourceTypeID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("resource_type_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.R.ResourceType = nil
	if related == nil || related.R == nil {
		return nil
	}

	for i, ri := range related.R.ResourceTypeResources {
		if queries.Equal(o.ResourceTypeID, ri.ResourceTypeID) {
			continue
		}

		ln := len(related.R.ResourceTypeResources)
		if ln > 1 && i < ln-1 {
			related.R.ResourceTypeResources[i] = related.R.ResourceTypeResources[ln-1]
		}
		related.R.ResourceTypeResources = related.R.ResourceTypeResources[:ln-1]
		break
	}
	return nil
}

// AddResourceRoleResources adds the given related objects to the existing relationships
// of the Resource, optionally inserting them as new records.
// Appends related to o.R.ResourceRoleResources.
// Sets related.R.Resource appropriately.
func (o *Resource) AddResourceRoleResources(ctx context.Context, exec boil.ContextExecutor, insert bool, related ...*RoleResource) error {
	var err error
	for _, rel := range related {
		if insert {
			rel.ResourceID = o.ResourceID
			if err = rel.Insert(ctx, exec, boil.Infer()); err != nil {
				return errors.Wrap(err, "failed to insert into foreign table")
			}
		} else {
			updateQuery := fmt.Sprintf(
				"UPDATE `RoleResources` SET %s WHERE %s",
				strmangle.SetParamNames("`", "`", 0, []string{"resource_id"}),
				strmangle.WhereClause("`", "`", 0, roleResourcePrimaryKeyColumns),
			)
			values := []interface{}{o.ResourceID, rel.RoleResourceID}

			if boil.DebugMode {
				fmt.Fprintln(boil.DebugWriter, updateQuery)
				fmt.Fprintln(boil.DebugWriter, values)
			}

			if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
				return errors.Wrap(err, "failed to update foreign table")
			}

			rel.ResourceID = o.ResourceID
		}
	}

	if o.R == nil {
		o.R = &resourceR{
			ResourceRoleResources: related,
		}
	} else {
		o.R.ResourceRoleResources = append(o.R.ResourceRoleResources, related...)
	}

	for _, rel := range related {
		if rel.R == nil {
			rel.R = &roleResourceR{
				Resource: o,
			}
		} else {
			rel.R.Resource = o
		}
	}
	return nil
}

// Resources retrieves all the records using an executor.
func Resources(mods ...qm.QueryMod) resourceQuery {
	mods = append(mods, qm.From("`Resources`"))
	return resourceQuery{NewQuery(mods...)}
}

// FindResource retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindResource(ctx context.Context, exec boil.ContextExecutor, resourceID int, selectCols ...string) (*Resource, error) {
	resourceObj := &Resource{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from `Resources` where `resource_id`=?", sel,
	)

	q := queries.Raw(query, resourceID)

	err := q.Bind(ctx, exec, resourceObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from Resources")
	}

	return resourceObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *Resource) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Resources provided for insertion")
	}

	var err error

	if err := o.doBeforeInsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(resourceColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	resourceInsertCacheMut.RLock()
	cache, cached := resourceInsertCache[key]
	resourceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			resourceColumns,
			resourceColumnsWithDefault,
			resourceColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(resourceType, resourceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(resourceType, resourceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO `Resources` (`%s`) %%sVALUES (%s)%%s", strings.Join(wl, "`,`"), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO `Resources` () VALUES ()%s%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			cache.retQuery = fmt.Sprintf("SELECT `%s` FROM `Resources` WHERE %s", strings.Join(returnColumns, "`,`"), strmangle.WhereClause("`", "`", 0, resourcePrimaryKeyColumns))
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
		return errors.Wrap(err, "models: unable to insert into Resources")
	}

	var identifierCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	identifierCols = []interface{}{
		o.ResourceID,
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, identifierCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, identifierCols...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for Resources")
	}

CacheNoHooks:
	if !cached {
		resourceInsertCacheMut.Lock()
		resourceInsertCache[key] = cache
		resourceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(ctx, exec)
}

// Update uses an executor to update the Resource.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *Resource) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	if err = o.doBeforeUpdateHooks(ctx, exec); err != nil {
		return 0, err
	}
	key := makeCacheKey(columns, nil)
	resourceUpdateCacheMut.RLock()
	cache, cached := resourceUpdateCache[key]
	resourceUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			resourceColumns,
			resourcePrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("models: unable to update Resources, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE `Resources` SET %s WHERE %s",
			strmangle.SetParamNames("`", "`", 0, wl),
			strmangle.WhereClause("`", "`", 0, resourcePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(resourceType, resourceMapping, append(wl, resourcePrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "models: unable to update Resources row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for Resources")
	}

	if !cached {
		resourceUpdateCacheMut.Lock()
		resourceUpdateCache[key] = cache
		resourceUpdateCacheMut.Unlock()
	}

	return rowsAff, o.doAfterUpdateHooks(ctx, exec)
}

// UpdateAll updates all rows with the specified column values.
func (q resourceQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for Resources")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for Resources")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ResourceSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), resourcePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE `Resources` SET %s WHERE %s",
		strmangle.SetParamNames("`", "`", 0, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, resourcePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in resource slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all resource")
	}
	return rowsAff, nil
}

var mySQLResourceUniqueColumns = []string{
	"resource_id",
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *Resource) Upsert(ctx context.Context, exec boil.ContextExecutor, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no Resources provided for upsert")
	}

	if err := o.doBeforeUpsertHooks(ctx, exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(resourceColumnsWithDefault, o)
	nzUniques := queries.NonZeroDefaultSet(mySQLResourceUniqueColumns, o)

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

	resourceUpsertCacheMut.RLock()
	cache, cached := resourceUpsertCache[key]
	resourceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			resourceColumns,
			resourceColumnsWithDefault,
			resourceColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			resourceColumns,
			resourcePrimaryKeyColumns,
		)

		if len(update) == 0 {
			return errors.New("models: unable to upsert Resources, could not build update column list")
		}

		ret = strmangle.SetComplement(ret, nzUniques)
		cache.query = buildUpsertQueryMySQL(dialect, "Resources", update, insert)
		cache.retQuery = fmt.Sprintf(
			"SELECT %s FROM `Resources` WHERE %s",
			strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, ret), ","),
			strmangle.WhereClause("`", "`", 0, nzUniques),
		)

		cache.valueMapping, err = queries.BindMapping(resourceType, resourceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(resourceType, resourceMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert for Resources")
	}

	var uniqueMap []uint64
	var nzUniqueCols []interface{}

	if len(cache.retMapping) == 0 {
		goto CacheNoHooks
	}

	uniqueMap, err = queries.BindMapping(resourceType, resourceMapping, nzUniques)
	if err != nil {
		return errors.Wrap(err, "models: unable to retrieve unique values for Resources")
	}
	nzUniqueCols = queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), uniqueMap)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.retQuery)
		fmt.Fprintln(boil.DebugWriter, nzUniqueCols...)
	}

	err = exec.QueryRowContext(ctx, cache.retQuery, nzUniqueCols...).Scan(returns...)
	if err != nil {
		return errors.Wrap(err, "models: unable to populate default values for Resources")
	}

CacheNoHooks:
	if !cached {
		resourceUpsertCacheMut.Lock()
		resourceUpsertCache[key] = cache
		resourceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(ctx, exec)
}

// Delete deletes a single Resource record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Resource) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Resource provided for delete")
	}

	if err := o.doBeforeDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), resourcePrimaryKeyMapping)
	sql := "DELETE FROM `Resources` WHERE `resource_id`=?"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from Resources")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for Resources")
	}

	if err := o.doAfterDeleteHooks(ctx, exec); err != nil {
		return 0, err
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q resourceQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no resourceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from Resources")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Resources")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ResourceSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no Resource slice provided for delete all")
	}

	if len(o) == 0 {
		return 0, nil
	}

	if len(resourceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(ctx, exec); err != nil {
				return 0, err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), resourcePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM `Resources` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, resourcePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from resource slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for Resources")
	}

	if len(resourceAfterDeleteHooks) != 0 {
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
func (o *Resource) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindResource(ctx, exec, o.ResourceID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ResourceSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := ResourceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), resourcePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT `Resources`.* FROM `Resources` WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 0, resourcePrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ResourceSlice")
	}

	*o = slice

	return nil
}

// ResourceExists checks if the Resource row exists.
func ResourceExists(ctx context.Context, exec boil.ContextExecutor, resourceID int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from `Resources` where `resource_id`=? limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, resourceID)
	}

	row := exec.QueryRowContext(ctx, sql, resourceID)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if Resources exists")
	}

	return exists, nil
}
