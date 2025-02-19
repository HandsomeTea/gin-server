package models

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type BaseModel struct {
	Colllection *mongo.Collection
}

func dealFilterId(filter map[string]any) map[string]any {
	if filter["_id"] != nil {
		strId, isString := filter["_id"].(string)

		if isString && len(strId) == 24 {
			filter["_id"], _ = bson.ObjectIDFromHex(strId)
		}
	}
	return filter
}

func (b *BaseModel) InsertOne(document interface{}) string {
	result, _ := b.Colllection.InsertOne(context.TODO(), document)

	return result.InsertedID.(bson.ObjectID).Hex()
}

func (b *BaseModel) InsertMany(documents []interface{}) []string {
	result, _ := b.Colllection.InsertMany(context.TODO(), documents)

	var ids []string

	for _, id := range result.InsertedIDs {
		ids = append(ids, id.(bson.ObjectID).Hex())
	}

	return ids
}

func (b *BaseModel) DeleteOne(filter map[string]any) {
	b.Colllection.DeleteOne(context.TODO(), dealFilterId(filter))
}

func (b *BaseModel) DeleteById(filter map[string]any) {
	//
}

func (b *BaseModel) DeleteMany(filter map[string]any) {
	b.Colllection.DeleteMany(context.TODO(), dealFilterId(filter))
}
func (b *BaseModel) UpdateOne(filter map[string]any, update map[string]any) {
	// b.Colllection.UpdateOne(context.TODO(), dealFilterId(filter), update)
}

func (b *BaseModel) UpdateById(_id string, update map[string]any) {
	// b.Colllection.UpdateOne(context.TODO(), bson.M{"_id": bson.ObjectIDFromHex(_id)}, update)
}

func (b *BaseModel) UpdateMany(filter map[string]any, update map[string]any) {
	// b.Colllection.UpdateMany(context.TODO(), dealFilterId(filter), update)
}

func (b *BaseModel) FindOne(filter map[string]any) bson.M {
	var result bson.M

	b.Colllection.FindOne(context.TODO(), dealFilterId(filter)).Decode(&result)

	return result
}

func (b *BaseModel) FindById(_id string) bson.M {
	var id bson.ObjectID

	if len(_id) == 24 {
		id, _ = bson.ObjectIDFromHex(_id)
	}
	var result bson.M

	b.Colllection.FindOne(context.TODO(), bson.M{"_id": id}).Decode(&result)

	return result
}

func (b *BaseModel) Find(filter map[string]any) []bson.M {
	var result []bson.M

	find, _ := b.Colllection.Find(context.TODO(), dealFilterId(filter))
	find.All(context.TODO(), &result)

	return result
}
