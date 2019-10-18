// Code generated by ObjectBox; DO NOT EDIT.
// Learn more about defining entities and generating this file - visit https://golang.objectbox.io/entity-annotations

package model

import (
	"github.com/google/flatbuffers/go"
	"github.com/objectbox/objectbox-go/objectbox"
	"github.com/objectbox/objectbox-go/objectbox/fbutils"
)

type entityByValue_EntityInfo struct {
	objectbox.Entity
	Uid uint64
}

var EntityByValueBinding = entityByValue_EntityInfo{
	Entity: objectbox.Entity{
		Id: 3,
	},
	Uid: 2793387980842421409,
}

// EntityByValue_ contains type-based Property helpers to facilitate some common operations such as Queries.
var EntityByValue_ = struct {
	Id   *objectbox.PropertyUint64
	Text *objectbox.PropertyString
}{
	Id: &objectbox.PropertyUint64{
		BaseProperty: &objectbox.BaseProperty{
			Id:     1,
			Entity: &EntityByValueBinding.Entity,
		},
	},
	Text: &objectbox.PropertyString{
		BaseProperty: &objectbox.BaseProperty{
			Id:     2,
			Entity: &EntityByValueBinding.Entity,
		},
	},
}

// GeneratorVersion is called by ObjectBox to verify the compatibility of the generator used to generate this code
func (entityByValue_EntityInfo) GeneratorVersion() int {
	return 4
}

// AddToModel is called by ObjectBox during model build
func (entityByValue_EntityInfo) AddToModel(model *objectbox.Model) {
	model.Entity("EntityByValue", 3, 2793387980842421409)
	model.Property("Id", 6, 1, 8853550994304785841)
	model.PropertyFlags(1)
	model.Property("Text", 9, 2, 6704507893545428268)
	model.EntityLastPropertyId(2, 6704507893545428268)
}

// GetId is called by ObjectBox during Put operations to check for existing ID on an object
func (entityByValue_EntityInfo) GetId(object interface{}) (uint64, error) {
	if obj, ok := object.(*EntityByValue); ok {
		return obj.Id, nil
	} else {
		return object.(EntityByValue).Id, nil
	}
}

// SetId is called by ObjectBox during Put to update an ID on an object that has just been inserted
func (entityByValue_EntityInfo) SetId(object interface{}, id uint64) error {
	if obj, ok := object.(*EntityByValue); ok {
		obj.Id = id
		return nil
	} else {
		// NOTE while this can't update, it will at least behave consistently (panic in case of a wrong type)
		_ = object.(EntityByValue).Id
		return nil
	}
}

// PutRelated is called by ObjectBox to put related entities before the object itself is flattened and put
func (entityByValue_EntityInfo) PutRelated(ob *objectbox.ObjectBox, object interface{}, id uint64) error {
	return nil
}

// Flatten is called by ObjectBox to transform an object to a FlatBuffer
func (entityByValue_EntityInfo) Flatten(object interface{}, fbb *flatbuffers.Builder, id uint64) error {
	var obj *EntityByValue
	if objPtr, ok := object.(*EntityByValue); ok {
		obj = objPtr
	} else {
		objVal := object.(EntityByValue)
		obj = &objVal
	}

	var offsetText = fbutils.CreateStringOffset(fbb, obj.Text)

	// build the FlatBuffers object
	fbb.StartObject(2)
	fbutils.SetUint64Slot(fbb, 0, id)
	fbutils.SetUOffsetTSlot(fbb, 1, offsetText)
	return nil
}

// Load is called by ObjectBox to load an object from a FlatBuffer
func (entityByValue_EntityInfo) Load(ob *objectbox.ObjectBox, bytes []byte) (interface{}, error) {
	var table = &flatbuffers.Table{
		Bytes: bytes,
		Pos:   flatbuffers.GetUOffsetT(bytes),
	}

	var propId = table.GetUint64Slot(4, 0)

	return &EntityByValue{
		Id:   propId,
		Text: fbutils.GetStringSlot(table, 6),
	}, nil
}

// MakeSlice is called by ObjectBox to construct a new slice to hold the read objects
func (entityByValue_EntityInfo) MakeSlice(capacity int) interface{} {
	return make([]EntityByValue, 0, capacity)
}

// AppendToSlice is called by ObjectBox to fill the slice of the read objects
func (entityByValue_EntityInfo) AppendToSlice(slice interface{}, object interface{}) interface{} {
	return append(slice.([]EntityByValue), *object.(*EntityByValue))
}

// Box provides CRUD access to EntityByValue objects
type EntityByValueBox struct {
	*objectbox.Box
}

// BoxForEntityByValue opens a box of EntityByValue objects
func BoxForEntityByValue(ob *objectbox.ObjectBox) *EntityByValueBox {
	return &EntityByValueBox{
		Box: ob.InternalBox(3),
	}
}

// Put synchronously inserts/updates a single object.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the EntityByValue.Id property on the passed object will be assigned the new ID as well.
func (box *EntityByValueBox) Put(object *EntityByValue) (uint64, error) {
	return box.Box.Put(object)
}

// Insert synchronously inserts a single object. As opposed to Put, Insert will fail if given an ID that already exists.
// In case the Id is not specified, it would be assigned automatically (auto-increment).
// When inserting, the EntityByValue.Id property on the passed object will be assigned the new ID as well.
func (box *EntityByValueBox) Insert(object *EntityByValue) (uint64, error) {
	return box.Box.Insert(object)
}

// Update synchronously updates a single object.
// As opposed to Put, Update will fail if an object with the same ID is not found in the database.
func (box *EntityByValueBox) Update(object *EntityByValue) error {
	return box.Box.Update(object)
}

// PutAsync asynchronously inserts/updates a single object.
// Deprecated: use box.Async().Put() instead
func (box *EntityByValueBox) PutAsync(object *EntityByValue) (uint64, error) {
	return box.Box.PutAsync(object)
}

// PutMany inserts multiple objects in single transaction.
// In case Ids are not set on the objects, they would be assigned automatically (auto-increment).
//
// Returns: IDs of the put objects (in the same order).
// When inserting, the EntityByValue.Id property on the objects in the slice will be assigned the new IDs as well.
//
// Note: In case an error occurs during the transaction, some of the objects may already have the EntityByValue.Id assigned
// even though the transaction has been rolled back and the objects are not stored under those IDs.
//
// Note: The slice may be empty or even nil; in both cases, an empty IDs slice and no error is returned.
func (box *EntityByValueBox) PutMany(objects []EntityByValue) ([]uint64, error) {
	return box.Box.PutMany(objects)
}

// Get reads a single object.
//
// Returns nil (and no error) in case the object with the given ID doesn't exist.
func (box *EntityByValueBox) Get(id uint64) (*EntityByValue, error) {
	object, err := box.Box.Get(id)
	if err != nil {
		return nil, err
	} else if object == nil {
		return nil, nil
	}
	return object.(*EntityByValue), nil
}

// GetMany reads multiple objects at once.
// If any of the objects doesn't exist, its position in the return slice is an empty object
func (box *EntityByValueBox) GetMany(ids ...uint64) ([]EntityByValue, error) {
	objects, err := box.Box.GetMany(ids...)
	if err != nil {
		return nil, err
	}
	return objects.([]EntityByValue), nil
}

// GetAll reads all stored objects
func (box *EntityByValueBox) GetAll() ([]EntityByValue, error) {
	objects, err := box.Box.GetAll()
	if err != nil {
		return nil, err
	}
	return objects.([]EntityByValue), nil
}

// Remove deletes a single object
func (box *EntityByValueBox) Remove(object *EntityByValue) error {
	return box.Box.Remove(object)
}

// RemoveMany deletes multiple objects at once.
// Returns the number of deleted object or error on failure.
// Note that this method will not fail if an object is not found (e.g. already removed).
// In case you need to strictly check whether all of the objects exist before removing them,
// you can execute multiple box.Contains() and box.Remove() inside a single write transaction.
func (box *EntityByValueBox) RemoveMany(objects ...*EntityByValue) (uint64, error) {
	var ids = make([]uint64, len(objects))
	for k, object := range objects {
		ids[k] = object.Id
	}
	return box.Box.RemoveIds(ids...)
}

// Creates a query with the given conditions. Use the fields of the EntityByValue_ struct to create conditions.
// Keep the *EntityByValueQuery if you intend to execute the query multiple times.
// Note: this function panics if you try to create illegal queries; e.g. use properties of an alien type.
// This is typically a programming error. Use QueryOrError instead if you want the explicit error check.
func (box *EntityByValueBox) Query(conditions ...objectbox.Condition) *EntityByValueQuery {
	return &EntityByValueQuery{
		box.Box.Query(conditions...),
	}
}

// Creates a query with the given conditions. Use the fields of the EntityByValue_ struct to create conditions.
// Keep the *EntityByValueQuery if you intend to execute the query multiple times.
func (box *EntityByValueBox) QueryOrError(conditions ...objectbox.Condition) (*EntityByValueQuery, error) {
	if query, err := box.Box.QueryOrError(conditions...); err != nil {
		return nil, err
	} else {
		return &EntityByValueQuery{query}, nil
	}
}

// Async provides access to the default Async Box for asynchronous operations. See EntityByValueAsyncBox for more information.
func (box *EntityByValueBox) Async() *EntityByValueAsyncBox {
	return &EntityByValueAsyncBox{AsyncBox: box.Box.Async()}
}

// EntityByValueAsyncBox provides asynchronous operations on EntityByValue objects.
//
// Asynchronous operations are executed on a separate internal thread for better performance.
//
// There are two main use cases:
//
// 1) "execute & forget:" you gain faster put/remove operations as you don't have to wait for the transaction to finish.
//
// 2) Many small transactions: if your write load is typically a lot of individual puts that happen in parallel,
// this will merge small transactions into bigger ones. This results in a significant gain in overall throughput.
//
// In situations with (extremely) high async load, an async method may be throttled (~1ms) or delayed up to 1 second.
// In the unlikely event that the object could still not be enqueued (full queue), an error will be returned.
//
// Note that async methods do not give you hard durability guarantees like the synchronous Box provides.
// There is a small time window in which the data may not have been committed durably yet.
type EntityByValueAsyncBox struct {
	*objectbox.AsyncBox
}

// AsyncBoxForEntityByValue creates a new async box with the given operation timeout in case an async queue is full.
// The returned struct must be freed explicitly using the Close() method.
// It's usually preferable to use EntityByValueBox::Async() which takes care of resource management and doesn't require closing.
func AsyncBoxForEntityByValue(ob *objectbox.ObjectBox, timeoutMs uint64) *EntityByValueAsyncBox {
	var async, err = objectbox.NewAsyncBox(ob, 3, timeoutMs)
	if err != nil {
		panic("Could not create async box for entity ID 3: %s" + err.Error())
	}
	return &EntityByValueAsyncBox{AsyncBox: async}
}

// Put inserts/updates a single object asynchronously.
// When inserting a new object, the Id property on the passed object will be assigned the new ID the entity would hold
// if the insert is ultimately successful. The newly assigned ID may not become valid if the insert fails.
func (asyncBox *EntityByValueAsyncBox) Put(object *EntityByValue) (uint64, error) {
	return asyncBox.AsyncBox.Put(object)
}

// Insert a single object asynchronously.
// The Id property on the passed object will be assigned the new ID the entity would hold if the insert is ultimately
// successful. The newly assigned ID may not become valid if the insert fails.
// Fails silently if an object with the same ID already exists (this error is not returned).
func (asyncBox *EntityByValueAsyncBox) Insert(object *EntityByValue) (id uint64, err error) {
	return asyncBox.AsyncBox.Insert(object)
}

// Update a single object asynchronously.
// The object must already exists or the update fails silently (without an error returned).
func (asyncBox *EntityByValueAsyncBox) Update(object *EntityByValue) error {
	return asyncBox.AsyncBox.Update(object)
}

// Remove deletes a single object asynchronously.
func (asyncBox *EntityByValueAsyncBox) Remove(object *EntityByValue) error {
	return asyncBox.AsyncBox.Remove(object)
}

// Query provides a way to search stored objects
//
// For example, you can find all EntityByValue which Id is either 42 or 47:
// 		box.Query(EntityByValue_.Id.In(42, 47)).Find()
type EntityByValueQuery struct {
	*objectbox.Query
}

// Find returns all objects matching the query
func (query *EntityByValueQuery) Find() ([]EntityByValue, error) {
	objects, err := query.Query.Find()
	if err != nil {
		return nil, err
	}
	return objects.([]EntityByValue), nil
}

// Offset defines the index of the first object to process (how many objects to skip)
func (query *EntityByValueQuery) Offset(offset uint64) *EntityByValueQuery {
	query.Query.Offset(offset)
	return query
}

// Limit sets the number of elements to process by the query
func (query *EntityByValueQuery) Limit(limit uint64) *EntityByValueQuery {
	query.Query.Limit(limit)
	return query
}
