package lmdatastore

import (
	"context"
	"strconv"

	grest "github.com/klummy/Grest"
	"google.golang.org/appengine/datastore"
)

// DatastoreCreateAll - Create a new Datastore entity
func DatastoreCreateAll(ctx context.Context, itemName string, item interface{}, ancestorKey *datastore.Key, customID string) (string, error) {
	if customID != "" {
		return DatastoreCreateWithID(ctx, itemName, item, ancestorKey, customID)
	}

	return DatastoreCreate(ctx, itemName, item, ancestorKey)
}

// DatastoreCreate - Simple create a new Datastore entity and return the object key
func DatastoreCreate(ctx context.Context, itemName string, item interface{}, ancestorKey *datastore.Key) (itemKey string, err error) {
	key := datastore.NewIncompleteKey(ctx, itemName, ancestorKey)

	generatedKey, generateErr := datastore.Put(ctx, key, item)

	if generateErr != nil {
		grest.ErrHandler(generateErr, "Generate error")
		return "", generateErr
	}

	itemKey = strconv.FormatInt(generatedKey.IntID(), 10)

	return itemKey, nil
}

// DatastoreCreateWithID - Create an entity by providing a custom ID
func DatastoreCreateWithID(ctx context.Context, itemName string, item interface{}, ancestorKey *datastore.Key, customID string) (itemKey string, err error) {
	// If a custom ID is passed (for example, firebase IDs), use that to create the entity
	if customID == "" {
		return "", grest.CustomError("No custom ID set")
	}

	customKey := datastore.NewKey(ctx, itemName, customID, 0, nil)

	// Check for an existing entity with the same ID to avoid it being overwritten
	query := datastore.NewQuery("UserData").Filter("__key__=", customKey).KeysOnly()
	results, err := query.Count(ctx)
	if err != nil {
		// ErrHandler(checkForExistingEntityErr, "Error getting data")
		return "", err
	} else if results > 0 {
		return "", grest.CustomError("Entity already exists")
	}

	customKey, customErr := datastore.Put(ctx, customKey, item)
	if customErr != nil {
		return "", customErr
	}

	return customID, nil
}

// DatastoreRead an item from the datastore and return the item
func DatastoreRead(ctx context.Context, itemName string, itemID string, itemStruct interface{}) (item interface{}, err error) {

	key, err := strconv.ParseInt(itemID, 10, 64)
	if err != nil {
		return nil, err
	}

	itemKey := datastore.NewKey(ctx, itemName, "", key, nil)
	if datastoreErr := datastore.Get(ctx, itemKey, itemStruct); datastoreErr != nil {
		return nil, datastoreErr
	}

	return itemStruct, nil
}

// DatastoreUpdateEntity an item from the datastore and return the item
func DatastoreUpdateEntity(ctx context.Context, itemName string, itemID string, itemStruct interface{}) (item interface{}, err error) {

	key, err := strconv.ParseInt(itemID, 10, 64)
	if err != nil {
		return nil, err
	}

	itemKey := datastore.NewKey(ctx, itemName, "", key, nil)
	if _, datastoreErr := datastore.Put(ctx, itemKey, itemStruct); datastoreErr != nil {
		return nil, datastoreErr
	}

	return itemStruct, nil
}
