package presensi

import (
	"fmt"
	"strings"
	"time"

	"github.com/aiteung/atmessage"
	"go.mau.fi/whatsmeow"
	"go.mau.fi/whatsmeow/types"
)

func MessageJamKerja(karyawan Karyawan, aktifjamkerja time.Duration, presensihariini Presensi, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
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

func MessagePulangKerja(karyawan Karyawan, aktifjamkerja time.Duration, id interface{}, lokasi string, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
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

func MessageMasukKerja(karyawan Karyawan, id interface{}, lokasi string, Info *types.MessageInfo, whatsapp *whatsmeow.Client) {
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
