package presensi

import (
	"os"

	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: os.Getenv("MONGOSTRING"),
	DBName:   "adorable",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

//func TestGetPresensiThisMonth(t *testing.T) {
//	GenerateReportCurrentMonth(MongoConn)

//}
