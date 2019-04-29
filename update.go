package repository

import (
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

// UpdateAndFind does an update and returns the updated model.
func UpdateAndFind(r Repository, id interface{}, dst interface{}, object interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		_, err := db.C(r.GetCollectionName()).Find(bson.M{"_id": id}).Apply(mgo.Change{
			Update: bson.M{
				"$set": object,
			},
			ReturnNew: true,
		}, dst)
		return err
	})
}

// Update does a complete update of a model.
func Update(r Repository, id interface{}, object interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		return db.C(r.GetCollectionName()).Update(bson.D{
			bson.DocElem{
				Name:  "_id",
				Value: id,
			}}, bson.M{
			"$set": object,
		})
	})
}

// UpdateRaw does a complete update of a model.
func UpdateRaw(r Repository, id interface{}, object interface{}) error {
	return r.GetQueryRunner().RunWithDB(func(db *mgo.Database) error {
		return db.C(r.GetCollectionName()).Update(
			bson.D{
				bson.DocElem{
					Name:  "_id",
					Value: id,
				},
			}, object,
		)
	})
}
