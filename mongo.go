package presensi

import (
	"context"
	"fmt"
	"time"

	"github.com/aiteung/module/model"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetNamaFromPhoneNumber(mongoconn *mongo.Database, phone_number string) (nama string) {
	karyawan := mongoconn.Collection("karyawan")
	filter := bson.M{"phone_number": phone_number}
	var staf Karyawan
	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("GetNamaFromPhoneNumber: %v\n", err)
	}
	return staf.Nama
}

func GetBiodataFromPhoneNumber(mongoconn *mongo.Database, phone_number string) (staf Karyawan) {
	karyawan := mongoconn.Collection("karyawan")
	filter := bson.M{"phone_number": phone_number}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("GetBiodataFromPhoneNumber: %v\n", err)
	}
	return staf
}

func GetKaryawanFromPhoneNumber(mongoconn *mongo.Database, phone_number string) (staf Karyawan) {
	karyawan := mongoconn.Collection("karyawan")
	filter := bson.M{"phone_number": phone_number}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getKaryawanFromPhoneNumber: %v\n", err)
	}
	return staf
}

func GetPresensiTodayFromPhoneNumber(mongoconn *mongo.Database, phone_number string) (presensi Presensi) {
	coll := mongoconn.Collection("presensi")
	today := bson.M{
		"$gte": primitive.NewDateTimeFromTime(time.Now().Truncate(24 * time.Hour).UTC()),
	}
	filter := bson.M{"phone_number": phone_number, "datetime": today}
	err := coll.FindOne(context.TODO(), filter).Decode(&presensi)
	if err != nil {
		fmt.Printf("getPresensiTodayFromPhoneNumber: %v\n", err)
	}
	return presensi
}

func GetPresensiCurrentMonth(mongoconn *mongo.Database) (allpresensi []Presensi) {
	startdate, enddate := GetFirstLastDateCurrentMonth()
	coll := mongoconn.Collection("presensi")
	today := bson.M{
		"$gte": primitive.NewDateTimeFromTime(startdate),
		"$lte": primitive.NewDateTimeFromTime(enddate),
	}
	filter := bson.M{"datetime": today}
	cursor, err := coll.Find(context.TODO(), filter)
	if err != nil {
		fmt.Printf("getPresensiTodayFromPhoneNumber: %v\n", err)
	}
	err = cursor.All(context.TODO(), &allpresensi)
	if err != nil {
		fmt.Println(err)
	}

	return
}

func GetLokasi(mongoconn *mongo.Database, long float64, lat float64) (namalokasi string) {
	lokasicollection := mongoconn.Collection("lokasi")
	filter := bson.M{
		"batas": bson.M{
			"$geoIntersects": bson.M{
				"$geometry": bson.M{
					"type":        "Point",
					"coordinates": []float64{long, lat},
				},
			},
		},
	}
	var lokasi Lokasi
	err := lokasicollection.FindOne(context.TODO(), filter).Decode(&lokasi)
	if err != nil {
		fmt.Printf("GetLokasi: %v\n", err)
	}
	return lokasi.Nama

}

func getKaryawanFromPhoneNumber(mongoconn *mongo.Database, phone_number string) (staf Karyawan) {
	karyawan := mongoconn.Collection("karyawan")
	filter := bson.M{"phone_number": phone_number}
	err := karyawan.FindOne(context.TODO(), filter).Decode(&staf)
	if err != nil {
		fmt.Printf("getKaryawanFromPhoneNumber: %v\n", err)
	}
	return staf
}

func getPresensiTodayFromPhoneNumber(mongoconn *mongo.Database, phone_number string) (presensi Presensi) {
	coll := mongoconn.Collection("presensi")
	today := bson.M{
		"$gte": primitive.NewDateTimeFromTime(time.Now().Truncate(24 * time.Hour).UTC()),
	}
	filter := bson.M{"phone_number": phone_number, "datetime": today}
	err := coll.FindOne(context.TODO(), filter).Decode(&presensi)
	if err != nil {
		fmt.Printf("getPresensiTodayFromPhoneNumber: %v\n", err)
	}
	return presensi
}

func InsertPresensi(Info *types.MessageInfo, Message *waProto.Message, Checkin string, mongoconn *mongo.Database) (InsertedID interface{}) {
	insertResult, err := mongoconn.Collection("presensi").InsertOne(context.TODO(), fillStructPresensi(Info, Message, Checkin, mongoconn))
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}

func insertPresensi(Pesan model.IteungMessage, Checkin string, mongoconn *mongo.Database) (InsertedID interface{}) {
	insertResult, err := mongoconn.Collection("presensi").InsertOne(context.TODO(), FillStructPresensi(Pesan, Checkin, mongoconn))
	if err != nil {
		fmt.Printf("InsertOneDoc: %v\n", err)
	}
	return insertResult.InsertedID
}
