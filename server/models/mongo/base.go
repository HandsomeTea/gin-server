package models

import (
	"context"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

type BaseModel struct {
	Colllection *mongo.Collection
}

func TransformId(_id string) bson.ObjectID {
	var id bson.ObjectID

	if len(_id) == 24 {
		id, _ = bson.ObjectIDFromHex(_id)
	} else {
		panic("Invalid mongodb _id")
	}

	return id
}

func dealFilter(filter bson.M) bson.M {
	if filter["_id"] != nil {
		strId, isString := filter["_id"].(string)

		if isString {
			filter["_id"] = TransformId(strId)
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

func (b *BaseModel) DeleteOne(filter bson.M) {
	b.Colllection.DeleteOne(context.TODO(), dealFilter(filter))
}

func (b *BaseModel) DeleteById(_id string) {
	b.Colllection.DeleteOne(context.TODO(), bson.M{"_id": TransformId(_id)})
}

func (b *BaseModel) FindOneAndDelete(filter bson.M) {
	b.Colllection.FindOneAndDelete(context.TODO(), dealFilter(filter))
}

func (b *BaseModel) DeleteMany(filter bson.M) {
	b.Colllection.DeleteMany(context.TODO(), dealFilter(filter))
}

func (b *BaseModel) UpdateOne(filter bson.M, update map[string]any) {
	b.Colllection.UpdateOne(context.TODO(), dealFilter(filter), update)
}

func (b *BaseModel) UpdateById(_id string, update map[string]any) {
	b.Colllection.UpdateOne(context.TODO(), bson.M{"_id": TransformId(_id)}, update)
}

func (b *BaseModel) UpdateMany(filter bson.M, update map[string]any) {
	b.Colllection.UpdateMany(context.TODO(), dealFilter(filter), update)
}

func (b *BaseModel) FindOne(filter bson.M) bson.M {
	var result bson.M

	b.Colllection.FindOne(context.TODO(), dealFilter(filter)).Decode(&result)

	return result
}

func (b *BaseModel) FindById(_id string) bson.M {
	var result bson.M

	b.Colllection.FindOne(context.TODO(), bson.M{"_id": TransformId(_id)}).Decode(&result)

	return result
}

func (b *BaseModel) Find(filter bson.M) []bson.M {
	var result []bson.M

	find, _ := b.Colllection.Find(context.TODO(), dealFilter(filter))
	find.All(context.TODO(), &result)

	return result
}

/* 返回修改后的数据 */
func (b *BaseModel) FindOneAndUpdate(filter bson.M, update map[string]any) bson.M {
	var result bson.M

	b.Colllection.FindOneAndUpdate(context.TODO(), dealFilter(filter), update, options.FindOneAndUpdate().SetReturnDocument(options.After)).Decode(&result)

	return result
}
