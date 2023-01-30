package presensi

import (
	"context"
	"fmt"
	"time"

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

func GetPresensiMonth(mongoconn *mongo.Database, month_number int) (presensi Presensi) {
	coll := mongoconn.Collection("presensi")
	today := bson.M{
		"$eq": []interface{}{bson.M{"$month": "datetime"}, month_number},
	}
	filter := bson.M{"$expr": today}
	err := coll.FindOne(context.TODO(), filter).Decode(&presensi)
	if err != nil {
		fmt.Printf("getPresensiTodayFromPhoneNumber: %v\n", err)
	}
	return presensi
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
