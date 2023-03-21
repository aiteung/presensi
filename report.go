package presensi

import (
	"encoding/csv"
	"os"
	"strconv"
	"time"

	"github.com/aiteung/atmessage"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateReportCurrentMonth(MongoConn *mongo.Database, to types.JID, whatsapp *whatsmeow.Client) (resp whatsmeow.SendResponse, err error) {
	location, _ := time.LoadLocation("Asia/Jakarta")
	atmessage.SendMessage("Mohon tunggu sebentar, laporan sedang di buat ya kak", to, whatsapp)
	res := GetPresensiCurrentMonth(MongoConn)
	msg := "Data rekap sebanyak : " + strconv.Itoa(len(res)) + " baris"
	atmessage.SendMessage(msg, to, whatsapp)
	path, err := os.Getwd()
	if err != nil {
		atmessage.SendMessage(err.Error(), to, whatsapp)
	}
	filename := path + "/rekapbulanini.csv"
	atmessage.SendMessage("nama file : "+filename, to, whatsapp)
	file, err := os.Create(filename)
	if err != nil {
		atmessage.SendMessage("failed to create file "+err.Error(), to, whatsapp)
	}
	cw := csv.NewWriter(file)
	err = cw.Write([]string{"DateTime", "Location", "Phone_Number", "CheckIn", "Nama", "Jabatan"})
	if err != nil {
		atmessage.SendMessage("failed to write file", to, whatsapp)
	}
	for _, prn := range res {
		waktuutc := prn.Datetime.Time().In(location).String()
		err = cw.Write([]string{waktuutc, prn.Location, prn.Phone_number, prn.Checkin, prn.Biodata.Nama, prn.Biodata.Jabatan})
		if err != nil {
			atmessage.SendMessage("failed to write file", to, whatsapp)
		}

	}
	cw.Flush()
	filebyte, err := os.ReadFile(filename)
	if err != nil {
		atmessage.SendMessage("failed to Read file", to, whatsapp)
	}

	msg = "File dikirim ke server : " + filename
	atmessage.SendMessage(msg, to, whatsapp)
	resp, err = atmessage.SendDocumentMessage(filebyte, "rekap.csv", "Rekapitulasi Presensi Bulan Ini", "application/csv", to, whatsapp)
	return
}
