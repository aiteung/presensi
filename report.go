package presensi

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aiteung/atmessage"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateReportCurrentMonth(MongoConn *mongo.Database, to types.JID, whatsapp *whatsmeow.Client) (resp whatsmeow.SendResponse, err error) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	res := GetPresensiCurrentMonth(MongoConn)
	fmt.Println(res)
	filename := "rekapbulanini.csv"
	file, err := os.Create(filename)
	if err != nil {
		log.Fatalln("failed to create file", err)
	}
	cw := csv.NewWriter(file)
	cw.Write([]string{"DateTime", "Location", "Phone_Number", "CheckIn", "Nama", "Jabatan"})
	for _, prn := range res {
		waktuutc := prn.Datetime.Time().In(location).String()
		cw.Write([]string{waktuutc, prn.Location, prn.Phone_number, prn.Checkin, prn.Biodata.Nama, prn.Biodata.Jabatan})
	}
	cw.Flush()
	filebyte, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	resp, err = atmessage.SendDocumentMessage(filebyte, to, whatsapp)
	return
}
