package presensi

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/aiteung/atmessage"
	"github.com/aiteung/musik"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
)

const Keyword string = "adorable"

func Handler(Info *types.MessageInfo, Message *waProto.Message, whatsapp *whatsmeow.Client, mongoconn *mongo.Database) {
	if Message.LiveLocationMessage != nil {
		LiveLocationMessage(Info, Message, whatsapp, mongoconn)
	} else if Message.ButtonsResponseMessage != nil {
		ButtonMessage(Info, Message, whatsapp)
	} else if IsMultiKey(Info, Message) {
		GenerateReportCurrentMonth(mongoconn, Info.Chat, whatsapp)
	}
}

func IsMultiKey(Info *types.MessageInfo, Message *waProto.Message) bool {
	m := musik.NormalizeString(Message.GetConversation())
	if (strings.Contains(m, "teung") && Info.Chat.Server == "g.us") || (Info.Chat.Server == "s.whatsapp.net") {
		complete, match := musik.IsMatch(m, "adorable", "rekap", "presen", "absen", "hrd", "sdm", "excel", "data", "bulan")
		fmt.Println(complete)
		if match >= 2 {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func ButtonMessage(Info *types.MessageInfo, Message *waProto.Message, whatsapp *whatsmeow.Client) {
	var msg string
	switch *Message.ButtonsResponseMessage.SelectedButtonId {
	case "adorable|ijin|wekwek":
		msg = "mau ijin kemana kak? yuk ingetin c obos buat set nomornya dulu biar bisa ijin ini jalan."
	case "adorable|sakit|lalala":
		msg = "Semoga lekas sembuh kak. yuk ingetin c obos buat set nomornya dulu biar bisa ijin sakit"
	case "adorable|dinas|kopkop":
		msg = "Ciee cieee yang lagi dinas diluar. yuk ingetin c obos buat set nomornya dulu biar bisa ijin sakit"
	case "adorable|lembur|wekwek":
		msg = "semangat 45 kak kejar setoran. kasih tau c obos belum set nomor nya biar bisa approve lembur"
	case "adorable|pulang|wekwek":
		msg = "hati hati di jalan ya kak, lihat lurus kedepan jangan ke lain hati kak nanti kesasar, sakit rasanya."
	default:
		msg = "Selamat datang di modul presensi, saat ini anda mengakses modul presensi."
	}

	atmessage.SendMessage(msg, Info.Sender, whatsapp)
}

func LiveLocationMessage(Info *types.MessageInfo, Message *waProto.Message, whatsapp *whatsmeow.Client, mongoconn *mongo.Database) {
	lokasi := GetLokasi(mongoconn, *Message.LiveLocationMessage.DegreesLongitude, *Message.LiveLocationMessage.DegreesLatitude)
	if lokasi != "" {
		hadirHandler(Info, Message, lokasi, whatsapp, mongoconn)
	} else {
		tidakhadirHandler(Info, Message, whatsapp, mongoconn)
	}

}

func tidakhadirHandler(Info *types.MessageInfo, Message *waProto.Message, whatsapp *whatsmeow.Client, mongoconn *mongo.Database) {
	lat, long := atmessage.GetLiveLoc(Message)
	nama := GetNamaFromPhoneNumber(mongoconn, Info.Sender.User)
	MessageTidakMasukKerja(nama, long, lat, Info, whatsapp)
}

func hadirHandler(Info *types.MessageInfo, Message *waProto.Message, lokasi string, whatsapp *whatsmeow.Client, mongoconn *mongo.Database) {
	presensihariini := getPresensiTodayFromPhoneNumber(mongoconn, Info.Sender.User)
	karyawan := getKaryawanFromPhoneNumber(mongoconn, Info.Sender.User)
	fmt.Println(karyawan.Jam_kerja[0].Durasi)
	if !reflect.ValueOf(presensihariini).IsZero() {
		fmt.Println(presensihariini)
		aktifjamkerja := time.Now().UTC().Sub(presensihariini.ID.Timestamp().UTC())
		fmt.Println(aktifjamkerja)
		if int(aktifjamkerja.Hours()) >= karyawan.Jam_kerja[0].Durasi {
			id := InsertPresensi(Info, Message, "pulang", mongoconn)
			MessagePulangKerja(karyawan, aktifjamkerja, id, lokasi, Info, whatsapp)
		} else {
			MessageJamKerja(karyawan, aktifjamkerja, presensihariini, Info, whatsapp)
		}
	} else {
		id := InsertPresensi(Info, Message, "masuk", mongoconn)
		MessageMasukKerja(karyawan, id, lokasi, Info, whatsapp)
	}
}

func fillStructPresensi(Info *types.MessageInfo, Message *waProto.Message, Checkin string, mongoconn *mongo.Database) (presensi Presensi) {
	presensi.Latitude, presensi.Longitude = atmessage.GetLiveLoc(Message)
	presensi.Location = GetLokasi(mongoconn, *Message.LiveLocationMessage.DegreesLongitude, *Message.LiveLocationMessage.DegreesLatitude)
	presensi.Phone_number = Info.Sender.User
	presensi.Datetime = primitive.NewDateTimeFromTime(time.Now().UTC())
	presensi.Checkin = Checkin
	presensi.Biodata = GetBiodataFromPhoneNumber(mongoconn, Info.Sender.User)
	return presensi
}

func Member(Info *types.MessageInfo, Message *waProto.Message, mongoconn *mongo.Database) (status bool) {
	if GetNamaFromPhoneNumber(mongoconn, Info.Sender.User) != "" && Info.Chat.Server != "g.us" && (Message.LiveLocationMessage != nil || Message.ButtonsResponseMessage != nil || IsMultiKey(Info, Message)) {
		if Message.ButtonsResponseMessage != nil {
			if strings.Contains(*Message.ButtonsResponseMessage.SelectedButtonId, Keyword) {
				status = true
			}
		} else {
			status = true
		}
	}
	return

}
