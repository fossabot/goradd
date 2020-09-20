package model

// Code generated by goradd. DO NOT EDIT.

import (
	"context"
	"fmt"

	"github.com/goradd/goradd/pkg/orm/broadcast"
	"github.com/goradd/goradd/pkg/orm/db"
	. "github.com/goradd/goradd/pkg/orm/op"
	"github.com/goradd/goradd/pkg/orm/query"
	"github.com/goradd/goradd/pkg/stringmap"
	"github.com/goradd/goradd/web/examples/gen/goradd/model/node"

	//"./node"
	"bytes"
	"encoding/gob"
	"encoding/json"
)

// addressBase is a base structure to be embedded in a "subclass" and provides the ORM access to the database.
// Do not directly access the internal variables, but rather use the accessor functions, since this class maintains internal state
// related to the variables.

type addressBase struct {
	id        string
	idIsValid bool
	idIsDirty bool

	personID        string
	personIDIsValid bool
	personIDIsDirty bool
	oPerson         *Person

	street        string
	streetIsValid bool
	streetIsDirty bool

	city        string
	cityIsNull  bool
	cityIsValid bool
	cityIsDirty bool

	// Custom aliases, if specified
	_aliases map[string]interface{}

	// Indicates whether this is a new object, or one loaded from the database. Used by Save to know whether to Insert or Update
	_restored bool

	// The original primary key for updates
	_originalPK string
}

const (
	AddressIDDefault       = ""
	AddressPersonIDDefault = ""
	AddressStreetDefault   = ""
	AddressCityDefault     = ""
)

const (
	Address_ID       = `ID`
	Address_PersonID = `PersonID`
	Address_Person   = `Person`
	Address_Street   = `Street`
	Address_City     = `City`
)

// Initialize or re-initialize a Address database object to default values.
func (o *addressBase) Initialize() {

	o.id = ""
	o.idIsValid = false
	o.idIsDirty = false

	o.personID = ""
	o.personIDIsValid = false
	o.personIDIsDirty = false

	o.street = ""
	o.streetIsValid = false
	o.streetIsDirty = false

	o.city = ""
	o.cityIsNull = true
	o.cityIsValid = true
	o.cityIsDirty = true

	o._restored = false
}

func (o *addressBase) PrimaryKey() string {
	return o.id
}

// ID returns the loaded value of ID.
func (o *addressBase) ID() string {
	return fmt.Sprint(o.id)
}

// IDIsValid returns true if the value was loaded from the database or has been set.
func (o *addressBase) IDIsValid() bool {
	return o._restored && o.idIsValid
}

// PersonID returns the loaded value of PersonID.
func (o *addressBase) PersonID() string {
	if o._restored && !o.personIDIsValid {
		panic("personID was not selected in the last query and has not been set, and so is not valid")
	}
	return o.personID
}

// PersonIDIsValid returns true if the value was loaded from the database or has been set.
func (o *addressBase) PersonIDIsValid() bool {
	return o.personIDIsValid
}

// Person returns the current value of the loaded Person, and nil if its not loaded.
func (o *addressBase) Person() *Person {
	return o.oPerson
}

// LoadPerson returns the related Person. If it is not already loaded,
// it will attempt to load it first.
func (o *addressBase) LoadPerson(ctx context.Context) *Person {
	if !o.personIDIsValid {
		return nil
	}

	if o.oPerson == nil {
		// Load and cache
		o.oPerson = LoadPerson(ctx, o.PersonID())
	}
	return o.oPerson
}

// SetPersonID sets the value of PersonID in the object, to be saved later using the Save() function.
func (o *addressBase) SetPersonID(v string) {
	o.personIDIsValid = true
	if o.personID != v || !o._restored {
		o.personID = v
		o.personIDIsDirty = true
		o.oPerson = nil
	}

}

// SetPerson sets the value of Person in the object, to be saved later using the Save() function.
func (o *addressBase) SetPerson(v *Person) {
	if v == nil {
		panic("Cannot set Person to a null value.")
	} else {
		o.oPerson = v
		o.personIDIsValid = true
		if o.personID != v.PrimaryKey() {
			o.personID = v.PrimaryKey()
			o.personIDIsDirty = true
		}
	}
}

// Street returns the loaded value of Street.
func (o *addressBase) Street() string {
	if o._restored && !o.streetIsValid {
		panic("street was not selected in the last query and has not been set, and so is not valid")
	}
	return o.street
}

// StreetIsValid returns true if the value was loaded from the database or has been set.
func (o *addressBase) StreetIsValid() bool {
	return o.streetIsValid
}

// SetStreet sets the value of Street in the object, to be saved later using the Save() function.
func (o *addressBase) SetStreet(v string) {
	o.streetIsValid = true
	if o.street != v || !o._restored {
		o.street = v
		o.streetIsDirty = true
	}

}

// City returns the loaded value of City.
func (o *addressBase) City() string {
	if o._restored && !o.cityIsValid {
		panic("city was not selected in the last query and has not been set, and so is not valid")
	}
	return o.city
}

// CityIsValid returns true if the value was loaded from the database or has been set.
func (o *addressBase) CityIsValid() bool {
	return o.cityIsValid
}

// CityIsNull returns true if the related database value is null.
func (o *addressBase) CityIsNull() bool {
	return o.cityIsNull
}

// City_I returns the loaded value of City as an interface.
// If the value in the database is NULL, a nil interface is returned.
func (o *addressBase) City_I() interface{} {
	if o._restored && !o.cityIsValid {
		panic("city was not selected in the last query and has not been set, and so is not valid")
	} else if o.cityIsNull {
		return nil
	}
	return o.city
}

func (o *addressBase) SetCity(i interface{}) {
	o.cityIsValid = true
	if i == nil {
		if !o.cityIsNull {
			o.cityIsNull = true
			o.cityIsDirty = true
			o.city = ""
		}
	} else {
		v := i.(string)
		if o.cityIsNull ||
			!o._restored ||
			o.city != v {

			o.cityIsNull = false
			o.city = v
			o.cityIsDirty = true
		}
	}
}

// GetAlias returns the alias for the given key.
func (o *addressBase) GetAlias(key string) query.AliasValue {
	if a, ok := o._aliases[key]; ok {
		return query.NewAliasValue(a)
	} else {
		panic("Alias " + key + " not found.")
		return query.NewAliasValue([]byte{})
	}
}

// Load returns a Address from the database.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
func LoadAddress(ctx context.Context, primaryKey string, joinOrSelectNodes ...query.NodeI) *Address {
	return queryAddresses(ctx).Where(Equal(node.Address().ID(), primaryKey)).joinOrSelect(joinOrSelectNodes...).Get()
}

// LoadAddressByID queries for a single Address object by the given unique index values.
// joinOrSelectNodes lets you provide nodes for joining to other tables or selecting specific fields. Table nodes will
// be considered Join nodes, and column nodes will be Select nodes. See Join() and Select() for more info.
// If you need a more elaborate query, use QueryAddresses() to start a query builder.
func LoadAddressByID(ctx context.Context, id string, joinOrSelectNodes ...query.NodeI) *Address {
	q := queryAddresses(ctx)
	q = q.Where(Equal(node.Address().ID(), id))
	return q.
		joinOrSelect(joinOrSelectNodes...).
		Get()
}

// HasAddressByID returns true if the
// given unique index values exist in the database.
func HasAddressByID(ctx context.Context, id string) bool {
	q := queryAddresses(ctx)
	q = q.Where(Equal(node.Address().ID(), id))
	return q.Count(false) == 1
}

// The AddressesBuilder uses the QueryBuilderI interface from the database to build a query.
// All query operations go through this query builder.
// End a query by calling either Load, Count, or Delete
type AddressesBuilder struct {
	base                query.QueryBuilderI
	hasConditionalJoins bool
}

func newAddressBuilder(ctx context.Context) *AddressesBuilder {
	b := &AddressesBuilder{
		base: db.GetDatabase("goradd").NewBuilder(ctx),
	}
	return b.Join(node.Address())
}

// Load terminates the query builder, performs the query, and returns a slice of Address objects. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice
func (b *AddressesBuilder) Load() (addressSlice []*Address) {
	results := b.base.Load()
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Address)
		o.load(item, o, nil, "")
		addressSlice = append(addressSlice, o)
	}
	return addressSlice
}

// LoadI terminates the query builder, performs the query, and returns a slice of interfaces. If there are
// any errors, they are returned in the context object. If no results come back from the query, it will return
// an empty slice.
func (b *AddressesBuilder) LoadI() (addressSlice []interface{}) {
	results := b.base.Load()
	if results == nil {
		return
	}
	for _, item := range results {
		o := new(Address)
		o.load(item, o, nil, "")
		addressSlice = append(addressSlice, o)
	}
	return addressSlice
}

// Get is a convenience method to return only the first item found in a query.
// The entire query is performed, so you should generally use this only if you know
// you are selecting on one or very few items.
func (b *AddressesBuilder) Get() *Address {
	results := b.Load()
	if results != nil && len(results) > 0 {
		obj := results[0]
		return obj
	} else {
		return nil
	}
}

// Expand expands an array type node so that it will produce individual rows instead of an array of items
func (b *AddressesBuilder) Expand(n query.NodeI) *AddressesBuilder {
	b.base.Expand(n)
	return b
}

// Join adds a node to the node tree so that its fields will appear in the query. Optionally add conditions to filter
// what gets included. The conditions will be AND'd with the basic condition matching the primary keys of the join.
func (b *AddressesBuilder) Join(n query.NodeI, conditions ...query.NodeI) *AddressesBuilder {
	var condition query.NodeI
	if len(conditions) > 1 {
		condition = And(conditions)
	} else if len(conditions) == 1 {
		condition = conditions[0]
	}
	b.base.Join(n, condition)
	if condition != nil {
		b.hasConditionalJoins = true
	}
	return b
}

// Where adds a condition to filter what gets selected.
func (b *AddressesBuilder) Where(c query.NodeI) *AddressesBuilder {
	b.base.Condition(c)
	return b
}

// OrderBy specifies how the resulting data should be sorted.
func (b *AddressesBuilder) OrderBy(nodes ...query.NodeI) *AddressesBuilder {
	b.base.OrderBy(nodes...)
	return b
}

// Limit will return a subset of the data, limited to the offset and number of rows specified
func (b *AddressesBuilder) Limit(maxRowCount int, offset int) *AddressesBuilder {
	b.base.Limit(maxRowCount, offset)
	return b
}

// Select optimizes the query to only return the specified fields. Once you put a Select in your query, you must
// specify all the fields that you will eventually read out. Be careful when selecting fields in joined tables, as joined
// tables will also contain pointers back to the parent table, and so the parent node should have the same field selected
// as the child node if you are querying those fields.
func (b *AddressesBuilder) Select(nodes ...query.NodeI) *AddressesBuilder {
	b.base.Select(nodes...)
	return b
}

// Alias lets you add a node with a custom name. After the query, you can read out the data using GetAlias() on a
// returned object. Alias is useful for adding calculations or subqueries to the query.
func (b *AddressesBuilder) Alias(name string, n query.NodeI) *AddressesBuilder {
	b.base.Alias(name, n)
	return b
}

// Distinct removes duplicates from the results of the query. Adding a Select() may help you get to the data you want, although
// using Distinct with joined tables is often not effective, since we force joined tables to include primary keys in the query, and this
// often ruins the effect of Distinct.
func (b *AddressesBuilder) Distinct() *AddressesBuilder {
	b.base.Distinct()
	return b
}

// GroupBy controls how results are grouped when using aggregate functions in an Alias() call.
func (b *AddressesBuilder) GroupBy(nodes ...query.NodeI) *AddressesBuilder {
	b.base.GroupBy(nodes...)
	return b
}

// Having does additional filtering on the results of the query.
func (b *AddressesBuilder) Having(node query.NodeI) *AddressesBuilder {
	b.base.Having(node)
	return b
}

// Count terminates a query and returns just the number of items selected.
//
// distinct wll count the number of distinct items, ignoring duplicates.
//
// nodes will select individual fields, and should be accompanied by a GroupBy.
func (b *AddressesBuilder) Count(distinct bool, nodes ...query.NodeI) uint {
	return b.base.Count(distinct, nodes...)
}

// Delete uses the query builder to delete a group of records that match the criteria
func (b *AddressesBuilder) Delete() {
	b.base.Delete()
	broadcast.BulkChange(b.base.Context(), "goradd", "address")
}

// Subquery uses the query builder to define a subquery within a larger query. You MUST include what
// you are selecting by adding Alias or Select functions on the subquery builder. Generally you would use
// this as a node to an Alias function on the surrounding query builder.
func (b *AddressesBuilder) Subquery() *query.SubqueryNode {
	return b.base.Subquery()
}

// joinOrSelect is a private helper function for the Load* functions
func (b *AddressesBuilder) joinOrSelect(nodes ...query.NodeI) *AddressesBuilder {
	for _, n := range nodes {
		switch n.(type) {
		case query.TableNodeI:
			b.base.Join(n, nil)
		case *query.ColumnNode:
			b.Select(n)
		}
	}
	return b
}

func CountAddressByID(ctx context.Context, id string) uint {
	return queryAddresses(ctx).Where(Equal(node.Address().ID(), id)).Count(false)
}

func CountAddressByPersonID(ctx context.Context, personID string) uint {
	return queryAddresses(ctx).Where(Equal(node.Address().PersonID(), personID)).Count(false)
}

func CountAddressByStreet(ctx context.Context, street string) uint {
	return queryAddresses(ctx).Where(Equal(node.Address().Street(), street)).Count(false)
}

func CountAddressByCity(ctx context.Context, city string) uint {
	return queryAddresses(ctx).Where(Equal(node.Address().City(), city)).Count(false)
}

// load is the private loader that transforms data coming from the database into a tree structure reflecting the relationships
// between the object chain requested by the user in the query.
// Care must be taken in the query, as Select clauses might not be honored if the child object has fields selected which the parent object does not have.
func (o *addressBase) load(m map[string]interface{}, objThis *Address, objParent interface{}, parentKey string) {
	if v, ok := m["id"]; ok && v != nil {
		if o.id, ok = v.(string); ok {
			o.idIsValid = true
			o.idIsDirty = false
			o._originalPK = o.id
		} else {
			panic("Wrong type found for id.")
		}
	} else {
		o.idIsValid = false
		o.id = ""
	}

	if v, ok := m["person_id"]; ok && v != nil {
		if o.personID, ok = v.(string); ok {
			o.personIDIsValid = true
			o.personIDIsDirty = false
		} else {
			panic("Wrong type found for person_id.")
		}
	} else {
		o.personIDIsValid = false
		o.personID = ""
	}

	if v, ok := m["Person"]; ok {
		if oPerson, ok2 := v.(map[string]interface{}); ok2 {
			o.oPerson = new(Person)
			o.oPerson.load(oPerson, o.oPerson, objThis, "Addresses")
			o.personIDIsValid = true
			o.personIDIsDirty = false
		} else {
			panic("Wrong type found for oPerson object.")
		}
	} else {
		o.oPerson = nil
	}

	if v, ok := m["street"]; ok && v != nil {
		if o.street, ok = v.(string); ok {
			o.streetIsValid = true
			o.streetIsDirty = false
		} else {
			panic("Wrong type found for street.")
		}
	} else {
		o.streetIsValid = false
		o.street = ""
	}

	if v, ok := m["city"]; ok {
		if v == nil {
			o.city = ""
			o.cityIsNull = true
			o.cityIsValid = true
			o.cityIsDirty = false
		} else if o.city, ok = v.(string); ok {
			o.cityIsNull = false
			o.cityIsValid = true
			o.cityIsDirty = false
		} else {
			panic("Wrong type found for city.")
		}
	} else {
		o.cityIsValid = false
		o.cityIsNull = true
		o.city = ""
	}

	if v, ok := m["aliases_"]; ok {
		o._aliases = map[string]interface{}(v.(db.ValueMap))
	}
	o._restored = true
}

// Save will update or insert the object, depending on the state of the object.
// If it has any auto-generated ids, those will be updated.
func (o *addressBase) Save(ctx context.Context) {
	if o._restored {
		o.update(ctx)
	} else {
		o.insert(ctx)
	}
}

// update will update the values in the database, saving any changed values.
func (o *addressBase) update(ctx context.Context) {
	var modifiedFields map[string]interface{}
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {

		if o.oPerson != nil {
			o.oPerson.Save(ctx)
			id := o.oPerson.PrimaryKey()
			o.SetPersonID(id)
		}

		if !o._restored {
			panic("Cannot update a record that was not originally read from the database.")
		}

		modifiedFields = o.getModifiedFields()
		if len(modifiedFields) != 0 {
			d.Update(ctx, "address", modifiedFields, "id", o._originalPK)
		}

	}) // transaction
	o.resetDirtyStatus()
	if len(modifiedFields) != 0 {
		broadcast.Update(ctx, "goradd", "address", o._originalPK, stringmap.SortedKeys(modifiedFields)...)
	}
}

// insert will insert the item into the database. Related items will be saved.
func (o *addressBase) insert(ctx context.Context) {
	d := Database()
	db.ExecuteTransaction(ctx, d, func() {
		if o.oPerson != nil {
			o.oPerson.Save(ctx)
			o.SetPerson(o.oPerson)
		}

		if !o.personIDIsValid {
			panic("a value for PersonID is required, and there is no default value. Call SetPersonID() before inserting the record.")
		}

		if !o.streetIsValid {
			panic("a value for Street is required, and there is no default value. Call SetStreet() before inserting the record.")
		}
		m := o.getValidFields()

		id := d.Insert(ctx, "address", m)
		o.id = id
		o._originalPK = id

	}) // transaction
	o.resetDirtyStatus()
	o._restored = true
	broadcast.Insert(ctx, "goradd", "address", o.PrimaryKey())
}

func (o *addressBase) getModifiedFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.idIsDirty {
		fields["id"] = o.id
	}
	if o.personIDIsDirty {
		fields["person_id"] = o.personID
	}
	if o.streetIsDirty {
		fields["street"] = o.street
	}
	if o.cityIsDirty {
		if o.cityIsNull {
			fields["city"] = nil
		} else {
			fields["city"] = o.city
		}
	}
	return
}

func (o *addressBase) getValidFields() (fields map[string]interface{}) {
	fields = map[string]interface{}{}
	if o.personIDIsValid {
		fields["person_id"] = o.personID
	}
	if o.streetIsValid {
		fields["street"] = o.street
	}
	if o.cityIsValid {
		if o.cityIsNull {
			fields["city"] = nil
		} else {
			fields["city"] = o.city
		}
	}
	return
}

// Delete deletes the associated record from the database.
func (o *addressBase) Delete(ctx context.Context) {
	if !o._restored {
		panic("Cannot delete a record that has no primary key value.")
	}
	d := Database()
	d.Delete(ctx, "address", "id", o.id)
	broadcast.Delete(ctx, "goradd", "address", fmt.Sprint(o.id))
}

// deleteAddress deletes the associated record from the database.
func deleteAddress(ctx context.Context, pk string) {
	d := db.GetDatabase("goradd")
	d.Delete(ctx, "address", "id", pk)
	broadcast.Delete(ctx, "goradd", "address", fmt.Sprint(pk))
}

func (o *addressBase) resetDirtyStatus() {
	o.idIsDirty = false
	o.personIDIsDirty = false
	o.streetIsDirty = false
	o.cityIsDirty = false

}

func (o *addressBase) IsDirty() bool {
	return o.idIsDirty ||
		o.personIDIsDirty || (o.oPerson != nil && o.oPerson.IsDirty()) ||
		o.streetIsDirty ||
		o.cityIsDirty
}

// Get returns the value of a field in the object based on the field's name.
// It will also get related objects if they are loaded.
// Invalid fields and objects are returned as nil
func (o *addressBase) Get(key string) interface{} {

	switch key {
	case "ID":
		if !o.idIsValid {
			return nil
		}
		return o.id

	case "PersonID":
		if !o.personIDIsValid {
			return nil
		}
		return o.personID

	case "Person":
		return o.Person()

	case "Street":
		if !o.streetIsValid {
			return nil
		}
		return o.street

	case "City":
		if !o.cityIsValid {
			return nil
		}
		return o.city

	}
	return nil
}

// MarshalBinary serializes the object into a buffer that is deserializable using UnmarshalBinary.
// It should be used for transmitting database objects over the wire, or for temporary storage. It does not send
// a version number, so if the data format changes, its up to you to invalidate the old stored objects.
// The framework uses this to serialize the object when it is stored in a control.
func (o *addressBase) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)
	encoder := gob.NewEncoder(buf)

	if err := encoder.Encode(o.id); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.idIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.idIsDirty); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.personID); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.personIDIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.personIDIsDirty); err != nil {
		return nil, err
	}

	if o.oPerson == nil {
		if err := encoder.Encode(false); err != nil {
			return nil, err
		}
	} else {
		if err := encoder.Encode(true); err != nil {
			return nil, err
		}
		if err := encoder.Encode(o.oPerson); err != nil {
			return nil, err
		}
	}
	if err := encoder.Encode(o.street); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.streetIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.streetIsDirty); err != nil {
		return nil, err
	}

	if err := encoder.Encode(o.city); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.cityIsNull); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.cityIsValid); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o.cityIsDirty); err != nil {
		return nil, err
	}

	if o._aliases == nil {
		if err := encoder.Encode(false); err != nil {
			return nil, err
		}
	} else {
		if err := encoder.Encode(true); err != nil {
			return nil, err
		}
		if err := encoder.Encode(o._aliases); err != nil {
			return nil, err
		}
	}

	if err := encoder.Encode(o._restored); err != nil {
		return nil, err
	}
	if err := encoder.Encode(o._originalPK); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (o *addressBase) UnmarshalBinary(data []byte) (err error) {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	var isPtr bool

	_ = isPtr

	if err = dec.Decode(&o.id); err != nil {
		return
	}
	if err = dec.Decode(&o.idIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.idIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.personID); err != nil {
		return
	}
	if err = dec.Decode(&o.personIDIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.personIDIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&isPtr); err != nil {
		return
	}
	if isPtr {
		if err = dec.Decode(&o.oPerson); err != nil {
			return
		}
	}
	if err = dec.Decode(&o.street); err != nil {
		return
	}
	if err = dec.Decode(&o.streetIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.streetIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&o.city); err != nil {
		return
	}
	if err = dec.Decode(&o.cityIsNull); err != nil {
		return
	}
	if err = dec.Decode(&o.cityIsValid); err != nil {
		return
	}
	if err = dec.Decode(&o.cityIsDirty); err != nil {
		return
	}

	if err = dec.Decode(&isPtr); err != nil {
		return
	}
	if isPtr {
		if err = dec.Decode(&o._aliases); err != nil {
			return
		}
	}

	if err = dec.Decode(&o._restored); err != nil {
		return
	}
	if err = dec.Decode(&o._originalPK); err != nil {
		return
	}

	return
}

// MarshalJSON serializes the object into a JSON object.
// Only valid data will be serialized, meaning, you can control what gets serialized by using Select to
// select only the fields you want when you query for the object.
func (o *addressBase) MarshalJSON() (data []byte, err error) {
	v := make(map[string]interface{})

	if o.idIsValid {
		v["id"] = o.id
	}

	if o.personIDIsValid {
		v["personID"] = o.personID
	}

	if val := o.Person(); val != nil {
		v["person"] = val
	}
	if o.streetIsValid {
		v["street"] = o.street
	}

	if o.cityIsValid {
		if o.cityIsNull {
			v["city"] = nil
		} else {
			v["city"] = o.city
		}
	}

	for _k, _v := range o._aliases {
		v[_k] = _v
	}
	return json.Marshal(v)
}

// Custom functions. See goradd/codegen/templates/orm/modelBase.
