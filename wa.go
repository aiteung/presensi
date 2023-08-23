package presensi

import (
	"fmt"
	"strings"
	"time"

	"github.com/aiteung/atmessage"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)

func MessageTidakMasukKerja(nama string, long, lat float64, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
	msg := "*Selamat Datang di Layanan Presensi Kak...*\n"
	msg = msg + "Hai kak " + nama + ", kakak belum berada pada lokasi presensi nih, ke lokasi presensi dulu ya kak. Atau barangkali ada perlu lain kak?\n"
	msg = msg + fmt.Sprintf("Lokasi kakak saat ini di koordinat : https://www.google.com/maps/@%f,%f,20z", lat, long)
	atmessage.SendMessage(msg, Info.Sender, whatsapp)
}

func ReplyMessageTidakMasukKerja(nama string, long, lat float64) string {
	msg := "*Selamat Datang di Layanan Presensi Kak...*\n"
	msg = msg + "Hai kak " + nama + ", kakak belum berada pada lokasi presensi nih, ke lokasi presensi dulu ya kak. Atau barangkali ada perlu lain kak?\n"
	msg = msg + fmt.Sprintf("Lokasi kakak saat ini di koordinat : https://www.google.com/maps/@%f,%f,20z", lat, long)
	return msg
}

func MessageMasukKerja(karyawan Karyawan, id interface{}, lokasi string, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
	msg := "*Masuk Kerja*\n"
	msg = msg + "Hai kak _" + karyawan.Nama + "_,\ndari bagian *" + karyawan.Jabatan + "*, \nmakasih ya sudah melakukan presensi masuk kerja\nLokasi : _*" + lokasi + "*_\nJangan lupa presensi pulangnya ya kak, caranya tinggal share live location lagi aja sama seperti presensi masuk tapi pada saat jam pulang ya kak. Makasi kak...\n"
	msg = msg + fmt.Sprintf("ID presensi masuk : %v", id)
	atmessage.SendMessage(msg, Info.Sender, whatsapp)
}

func ReplyMessageMasukKerja(karyawan Karyawan, id interface{}, lokasi string) string {
	msg := "*Masuk Kerja*\n"
	msg = msg + "Hai kak _" + karyawan.Nama + "_,\ndari bagian *" + karyawan.Jabatan + "*, \nmakasih ya sudah melakukan presensi masuk kerja\nLokasi : _*" + lokasi + "*_\nJangan lupa presensi pulangnya ya kak, caranya tinggal share live location lagi aja sama seperti presensi masuk tapi pada saat jam pulang ya kak. Makasi kak...\n"
	msg = msg + fmt.Sprintf("ID presensi masuk : %v", id)
	return msg
}

func MessageJamKerja(karyawan Karyawan, aktifjamkerja time.Duration, presensihariini Presensi, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
	msg := "*Keterangan Presensi Kerja*\n"
	msg = msg + fmt.Sprintf("yah kak, mohon maaf jam kerja nya belum %v jam. Sabar dulu ya..... nanti presensi kembali.\n", karyawan.Jam_kerja[0].Durasi)
	msg = msg + fmt.Sprintf("ID presensi masuk : %v", presensihariini.ID) + "\n" + "Durasi Kerja : " + strings.Replace(aktifjamkerja.String(), "h", " jam ", 1)
	atmessage.SendMessage(msg, Info.Sender, whatsapp)
}

func ReplyMessageJamKerja(karyawan Karyawan, aktifjamkerja time.Duration, presensihariini Presensi) string {
	msg := "*Keterangan Presensi Kerja*\n"
	msg = msg + fmt.Sprintf("yah kak, mohon maaf jam kerja nya belum %v jam. Sabar dulu ya..... nanti presensi kembali.\n", karyawan.Jam_kerja[0].Durasi)
	msg = msg + fmt.Sprintf("ID presensi masuk : %v", presensihariini.ID) + "\n" + "Durasi Kerja : " + strings.Replace(aktifjamkerja.String(), "h", " jam ", 1)
	return msg
}

func MessagePulangKerja(karyawan Karyawan, aktifjamkerja time.Duration, id interface{}, lokasi string, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
	msg := "*Pulang Kerja*\n"
	msg = msg + "Hai kak _" + karyawan.Nama + "_,\ndari bagian *" + karyawan.Jabatan + "*, \nmakasih ya sudah melakukan presensi pulang kerja\nLokasi : _*" + lokasi + "*_\n"
	msg = msg + fmt.Sprintf("\nID presensi pulang : %v", id) + "\n" + "Durasi Kerja : " + strings.Replace(aktifjamkerja.String(), "h", " jam ", 1)
	atmessage.SendMessage(msg, Info.Sender, whatsapp)
}

func ReplyMessagePulangKerja(karyawan Karyawan, aktifjamkerja time.Duration, id interface{}, lokasi string) string {
	msg := "*Pulang Kerja*\n"
	msg = msg + "Hai kak _" + karyawan.Nama + "_,\ndari bagian *" + karyawan.Jabatan + "*, \nmakasih ya sudah melakukan presensi pulang kerja\nLokasi : _*" + lokasi + "*_\n"
	msg = msg + fmt.Sprintf("\nID presensi pulang : %v", id) + "\n" + "Durasi Kerja : " + strings.Replace(aktifjamkerja.String(), "h", " jam ", 1)
	return msg
}

func ButtonMessageJamKerja(karyawan Karyawan, aktifjamkerja time.Duration, presensihariini Presensi, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
	var btnmsg atmessage.ButtonsMessage
	btnmsg.Message.HeaderText = "Keterangan Presensi Kerja"
	btnmsg.Message.ContentText = fmt.Sprintf("yah kak, mohon maaf jam kerja nya belum %v jam. Sabar dulu ya..... nanti presensi kembali.", karyawan.Jam_kerja[0].Durasi)
	btnmsg.Message.FooterText = fmt.Sprintf("ID presensi masuk : %v", presensihariini.ID) + "\n" + "Durasi Kerja : " + strings.Replace(aktifjamkerja.String(), "h", " jam ", 1)
	btnmsg.Buttons = []atmessage.WaButton{{
		ButtonId:    "adorable|ijin|wekwek",
		DisplayText: "Ijin Keluar",
	},
		{
			ButtonId:    "adorable|sakit|lalala",
			DisplayText: "Lagi Sakit",
		},
		{
			ButtonId:    "adorable|dinas|kopkop",
			DisplayText: "Dinas Luar",
		},
	}
	atmessage.SendButtonMessage(btnmsg, Info.Sender, whatsapp)
}

func ButtonMessagePulangKerja(karyawan Karyawan, aktifjamkerja time.Duration, id interface{}, lokasi string, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
	var btnmsg atmessage.ButtonsMessage
	btnmsg.Message.HeaderText = "Pulang Kerja"

	btnmsg.Message.FooterText = fmt.Sprintf("ID presensi pulang : %v", id) + "\n" + "Durasi Kerja : " + strings.Replace(aktifjamkerja.String(), "h", " jam ", 1)
	btnmsg.Message.ContentText = "Hai kak _" + karyawan.Nama + "_,\ndari bagian *" + karyawan.Jabatan + "*, \nmakasih ya sudah melakukan presensi pulang kerja\nLokasi : _*" + lokasi + "*_"
	btnmsg.Buttons = []atmessage.WaButton{{
		ButtonId:    "adorable|pulang|wekwek",
		DisplayText: "Langsung Pulang",
	}, {
		ButtonId:    "adorable|lembur|wekwek",
		DisplayText: "Lanjut Lembur",
	},
	}
	atmessage.SendButtonMessage(btnmsg, Info.Sender, whatsapp)
}

func ButtonMessageMasukKerja(karyawan Karyawan, id interface{}, lokasi string, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
	var btnmsg atmessage.ButtonsMessage
	btnmsg.Message.HeaderText = "Masuk Kerja"
	btnmsg.Message.ContentText = "Hai kak _" + karyawan.Nama + "_,\ndari bagian *" + karyawan.Jabatan + "*, \nmakasih ya sudah melakukan presensi masuk kerja\nLokasi : _*" + lokasi + "*_\nJangan lupa presensi pulangnya ya kak, caranya tinggal share live location lagi aja sama seperti presensi masuk tapi pada saat jam pulang ya kak. Makasi kak..."
	btnmsg.Message.FooterText = fmt.Sprintf("ID presensi masuk : %v", id)
	btnmsg.Buttons = []atmessage.WaButton{{
		ButtonId:    "adorable|ijin|wekwek",
		DisplayText: "Ijin Keluar",
	},
		{
			ButtonId:    "adorable|sakit|lalala",
			DisplayText: "Lagi Sakit",
		},
		{
			ButtonId:    "adorable|dinas|kopkop",
			DisplayText: "Dinas Luar",
		},
	}
	atmessage.SendButtonMessage(btnmsg, Info.Sender, whatsapp)

}

func ListMessageMasukKerja(karyawan Karyawan, id interface{}, lokasi string, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
	var lmsg atmessage.ListMessage
	lmsg.Title = "Masuk Kerja"
	lmsg.Description = "Hai kak _" + karyawan.Nama + "_,\ndari bagian *" + karyawan.Jabatan + "*, \nmakasih ya sudah melakukan presensi masuk kerja\nLokasi : _*" + lokasi + "*_\nJangan lupa presensi pulangnya ya kak, caranya tinggal share live location lagi aja sama seperti presensi masuk tapi pada saat jam pulang ya kak. Makasi kak..."
	lmsg.FooterText = fmt.Sprintf("ID presensi masuk : %v", id)

	lmsg.ButtonText = "Keterangan"
	var listrow []atmessage.WaListRow
	var poin atmessage.WaListRow

	poin.Title = "Ijin Keluar"
	poin.Description = "Konfirmasi Atasan"
	poin.RowId = "adorable|ijin|wekwek"
	listrow = append(listrow, poin)

	poin.Title = "Lagi Sakit"
	poin.Description = "Konfirmasi Atasan"
	poin.RowId = "adorable|sakit|wekwek"
	listrow = append(listrow, poin)

	poin.Title = "Dinas Keluar"
	poin.Description = "Konfirmasi Atasan"
	poin.RowId = "adorable|dinas|wekwek"
	listrow = append(listrow, poin)

	var sec atmessage.WaListSection
	sec.Title = "Jika Tidak Masuk Kerja"
	sec.Rows = listrow
	var secs []atmessage.WaListSection
	secs = append(secs, sec)
	lmsg.Sections = secs
	atmessage.SendListMessage(lmsg, Info.Sender, whatsapp)

}

func ListMessagePulangKerja(karyawan Karyawan, aktifjamkerja time.Duration, id interface{}, lokasi string, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
	var lmsg atmessage.ListMessage
	lmsg.Title = "Pulang Kerja"
	lmsg.FooterText = fmt.Sprintf("ID presensi pulang : %v", id) + "\n" + "Durasi Kerja : " + strings.Replace(aktifjamkerja.String(), "h", " jam ", 1)
	lmsg.Description = "Hai kak _" + karyawan.Nama + "_,\ndari bagian *" + karyawan.Jabatan + "*, \nmakasih ya sudah melakukan presensi pulang kerja\nLokasi : _*" + lokasi + "*_"

	lmsg.ButtonText = "Keterangan"
	var listrow []atmessage.WaListRow
	var poin atmessage.WaListRow

	poin.Title = "Langsung Pulang"
	poin.Description = "Terima Kasih atas kontribusinya hari ini"
	poin.RowId = "adorable|pulang|wekwek"
	listrow = append(listrow, poin)

	poin.Title = "Lanjut Lembur"
	poin.Description = "Untuk melanjutkan lembur"
	poin.RowId = "adorable|lembur|wekwek"
	listrow = append(listrow, poin)

	var sec atmessage.WaListSection
	sec.Title = "Keterangan"
	sec.Rows = listrow
	var secs []atmessage.WaListSection
	secs = append(secs, sec)
	lmsg.Sections = secs
	atmessage.SendListMessage(lmsg, Info.Sender, whatsapp)

}

func ListMessageJamKerja(karyawan Karyawan, aktifjamkerja time.Duration, presensihariini Presensi, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
	var lmsg atmessage.ListMessage
	lmsg.Title = "Keterangan Presensi Kerja"
	lmsg.Description = fmt.Sprintf("yah kak, mohon maaf jam kerja nya belum %v jam. Sabar dulu ya..... nanti presensi kembali.", karyawan.Jam_kerja[0].Durasi)
	lmsg.FooterText = fmt.Sprintf("ID presensi masuk : %v", presensihariini.ID) + "\n" + "Durasi Kerja : " + strings.Replace(aktifjamkerja.String(), "h", " jam ", 1)

	lmsg.ButtonText = "Keterangan"
	var listrow []atmessage.WaListRow
	var poin atmessage.WaListRow

	poin.Title = "Ijin Keluar"
	poin.Description = "Konfirmasi Atasan"
	poin.RowId = "adorable|ijin|wekwek"
	listrow = append(listrow, poin)

	poin.Title = "Lagi Sakit"
	poin.Description = "Konfirmasi Atasan"
	poin.RowId = "adorable|sakit|wekwek"
	listrow = append(listrow, poin)

	poin.Title = "Dinas Keluar"
	poin.Description = "Konfirmasi Atasan"
	poin.RowId = "adorable|dinas|wekwek"
	listrow = append(listrow, poin)

	var sec atmessage.WaListSection
	sec.Title = "Jika Berhalangan Kerja"
	sec.Rows = listrow
	var secs []atmessage.WaListSection
	secs = append(secs, sec)
	lmsg.Sections = secs
	atmessage.SendListMessage(lmsg, Info.Sender, whatsapp)

}
