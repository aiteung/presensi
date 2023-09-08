package presensi

import (
	"encoding/base64"
	"encoding/csv"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/module"
	"github.com/aiteung/module/model"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/mongo"
)

func GenerateReportCurrentMonth(MongoConn *mongo.Database, im model.IteungMessage, ApiWa string, ApiWaDoc string) atmessage.Response {
	msg := model.GowaNotif{
		User:   im.Chat_number,
		Server: im.Chat_server,
	}
	location, _ := time.LoadLocation("Asia/Jakarta")
	msg.Messages = "Mohon tunggu sebentar, laporan sedang di buat ya kak"
	module.SendToGoWAAPI(msg, ApiWa)
	res := GetPresensiCurrentMonth(MongoConn)
	msg.Messages = "Data rekap sebanyak : " + strconv.Itoa(len(res)) + " baris"
	module.SendToGoWAAPI(msg, ApiWa)
	path, err := os.Getwd()
	if err != nil {
		msg.Messages = err.Error()
		module.SendToGoWAAPI(msg, ApiWa)
	}
	id := uuid.New()
	filename := path + "/" + id.String() + ".csv"
	msg.Messages = "nama file : " + filename
	module.SendToGoWAAPI(msg, ApiWa)
	file, err := os.Create(filename)
	if err != nil {
		msg.Messages = "failed to create file " + err.Error()
		module.SendToGoWAAPI(msg, ApiWa)
	}
	cw := csv.NewWriter(file)
	cw.Comma = ','
	err = cw.Write([]string{"DateTime", "Location", "Phone_Number", "CheckIn", "Nama", "Jabatan"})
	if err != nil {
		msg.Messages = "failed to write file " + err.Error()
		module.SendToGoWAAPI(msg, ApiWa)
	}
	for _, prn := range res {
		waktuutc := prn.Datetime.Time().In(location).String()
		err = cw.Write([]string{waktuutc, prn.Location, prn.Phone_number, prn.Checkin, prn.Biodata.Nama, prn.Biodata.Jabatan})
		if err != nil {
			msg.Messages = "failed to write file " + err.Error()
			module.SendToGoWAAPI(msg, ApiWa)
		}

	}
	cw.Flush()
	filebyte, err := os.ReadFile(filename)
	fileBase64Str := base64.StdEncoding.EncodeToString(filebyte)
	if err != nil {
		msg.Messages = "failed to Read file" + err.Error()
		module.SendToGoWAAPI(msg, ApiWa)
	}
	im.Filedata = fileBase64Str
	im.Filename = filepath.Base(filename)
	msg.Messages = "File dikirim ke server : " + filename
	os.Remove(filename)
	resp, _ := module.DocumentSendToGoWAAPI(im, ApiWaDoc)
	return resp
}
