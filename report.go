package presensi

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateReportCurrentMonth(MongoConn *mongo.Database) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	res := GetPresensiCurrentMonth(MongoConn)
	fmt.Println(res)
	file, err := os.Create("rekap.csv")
	if err != nil {
		log.Fatalln("failed to open file", err)
	}
	cw := csv.NewWriter(file)
	cw.Write([]string{"DateTime", "Location", "Phone_Number", "CheckIn", "Nama", "Jabatan"})
	for _, prn := range res {
		waktuutc := prn.Datetime.Time().In(location).String()
		cw.Write([]string{waktuutc, prn.Location, prn.Phone_number, prn.Checkin, prn.Biodata.Nama, prn.Biodata.Jabatan})
	}
	cw.Flush()
}
