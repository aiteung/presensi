package presensi

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoInfo = atdb.DBInfo{
	DBString: os.Getenv("MONGOSTRING"),
	DBName:   "adorable",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestWacipher(t *testing.T) {
	hasil := GetPresensiMonth(MongoConn, 1)
	fmt.Println(hasil)
}
