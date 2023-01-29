package presensi

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Karyawan struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty"`
	Jabatan      string             `bson:"jabatan,omitempty"`
	Jam_kerja    []JamKerja         `bson:"jam_kerja,omitempty"`
	Hari_kerja   []string           `bson:"hari_kerja,omitempty"`
}

type JamKerja struct {
	Durasi     int      `bson:"durasi,omitempty"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty"`
	Jam_keluar string   `bson:"jam_keluar,omitempty"`
	Gmt        int      `bson:"gmt,omitempty"`
	Hari       []string `bson:"hari,omitempty"`
	Shift      int      `bson:"shift,omitempty"`
	Piket_tim  string   `bson:"piket_tim,omitempty"`
}

type Presensi struct {
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty"`
	Latitude     float64            `bson:"latitude,omitempty"`
	Location     string             `bson:"location,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty"`
	Checkin      string             `bson:"checkin,omitempty"`
	Biodata      Karyawan           `bson:"biodata,omitempty"`
}

type Lokasi struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty"`
	Batas    Geometry           `bson:"batas,omitempty"`
	Kategori string             `bson:"kategori,omitempty"`
}

type Geometry struct {
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}
