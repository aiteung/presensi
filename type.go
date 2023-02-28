package presensi

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Karyawan struct { //data karwayan unik
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Nama         string             `bson:"nama,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty"`
	Jabatan      string             `bson:"jabatan,omitempty"`
	Jam_kerja    []JamKerja         `bson:"jam_kerja,omitempty"`
	Hari_kerja   []string           `bson:"hari_kerja,omitempty"`
}

type JamKerja struct { //info tambahan dari karyawan
	Durasi     int      `bson:"durasi,omitempty"`
	Jam_masuk  string   `bson:"jam_masuk,omitempty"`
	Jam_keluar string   `bson:"jam_keluar,omitempty"`
	Gmt        int      `bson:"gmt,omitempty"`
	Hari       []string `bson:"hari,omitempty"`
	Shift      int      `bson:"shift,omitempty"`
	Piket_tim  string   `bson:"piket_tim,omitempty"`
}

type Presensi struct { // input presensi, dimana pulang adalaha kewajiban 8 jam
	ID           primitive.ObjectID `bson:"_id,omitempty"`
	Longitude    float64            `bson:"longitude,omitempty"`
	Latitude     float64            `bson:"latitude,omitempty"`
	Location     string             `bson:"location,omitempty"`
	Phone_number string             `bson:"phone_number,omitempty"`
	Datetime     primitive.DateTime `bson:"datetime,omitempty"`
	Checkin      string             `bson:"checkin,omitempty"`
	Biodata      Karyawan           `bson:"biodata,omitempty"`
}

type RekapPresensi struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	In            Presensi           `bson:"in,omitempty"`
	Out           Presensi           `bson:"out,omitempty"`
	Lembur        Presensi           `bson:"lembur,omitempty"`
	Keterangan    string             `bson:"keterangan,omitempty"`
	TotalJamKerja primitive.DateTime `bson:"totaljamkerja,omitempty"`
	Late          primitive.DateTime `bson:"late,omitempty"`
}

type Lokasi struct { //lokasi yang bisa melakukan presensi
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Nama     string             `bson:"nama,omitempty"`
	Batas    Geometry           `bson:"batas,omitempty"`
	Kategori string             `bson:"kategori,omitempty"`
}

type Geometry struct { //data geometry untuk lokasi presensi
	Type        string      `json:"type" bson:"type"`
	Coordinates interface{} `json:"coordinates" bson:"coordinates"`
}
