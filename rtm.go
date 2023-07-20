package rtmpackage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertOneDoc(db *mongo.Database, collection string, doc interface{}) (insertedID interface{}) {
	insertResult, err := db.Collection(collection).InsertOne(context.TODO(), doc)
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func InsertDataListJobdesk(db *mongo.Database, jbtitle_LJ, deskripsi_lj, deadline_lj string, priority_lj string) (InsertedID interface{}) {
	var listjbdesk List_Jbdesk
	listjbdesk.Jbtitle_LJ = jbtitle_LJ
	listjbdesk.Deskripsi_LJ = deskripsi_lj
	listjbdesk.Deadline_LJ = deadline_lj
	listjbdesk.Priority_LJ = priority_lj
	return InsertOneDoc(db, "ListJD", listjbdesk)
}

func GetDataListJobdeskDeskripsi(deskripsi_lj string, db *mongo.Database, col string) (data List_Jbdesk) {
	user := db.Collection(col)
	filter := bson.M{"deskripsi_lj": deskripsi_lj}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdeskripsi: %v\n", err)
	}
	return data
}

func GetDataListJobdeskDeadline(deadline_lj string, db *mongo.Database, col string) (data List_Jbdesk) {
	user := db.Collection(col)
	filter := bson.M{"deadline_lj": deadline_lj}
	err := user.FindOne(context.TODO(), filter).Decode(&data)
	if err != nil {
		fmt.Printf("getdeadline: %v\n", err)
	}
	return data
}

func DeleteDataListJobdeskDeskripsi(deskripsi_lj string, db *mongo.Database, col string) (data List_Jbdesk) {
	jbdk := db.Collection(col)
	filter := bson.M{"deskripsi_lj": deskripsi_lj}
	err, _ := jbdk.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Deletedatalistdeksripsi : %v\n", err)
	}
	fmt.Println("Success Delete your data")
	return data
}

func DeleteDataListJobdeskDeadline(deadline_lj string, db *mongo.Database, col string) (data List_Jbdesk) {
	jbdl := db.Collection(col)
	filter := bson.M{"deadline_lj": deadline_lj}
	err, _ := jbdl.DeleteOne(context.TODO(), filter)
	if err != nil {
		fmt.Printf("Deletedatalistdeadline : %v\n", err)
	}
	fmt.Println("Success Delete your data")
	return data
}
