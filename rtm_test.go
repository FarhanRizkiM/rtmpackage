package rtmpackage

import (
	"fmt"
	"os"
	"testing"

	"github.com/aiteung/atdb"
)

var MongoString string = os.Getenv("MONGOSTRING")

var MongoInfo = atdb.DBInfo{
	DBString: MongoString,
	DBName:   "List_Jobdesk",
}

var MongoConn = atdb.MongoConnect(MongoInfo)

func TestInsertDataListJobdesk(t *testing.T) {
	Jbtitle_LJ := "Manager"
	Deskripsi_LJ := "Mengembangkan dan menerapkan strategi pemasaran"
	Deadline_LJ := "31/07/2023"
	Priority_LJ := "High"
	hasil := InsertDataListJobdesk(MongoConn, Jbtitle_LJ, Deskripsi_LJ, Deadline_LJ, Priority_LJ)
	fmt.Println(hasil)

}

func TestGetDataListJobdeskDeskripsi(t *testing.T) {
	Deskripsi_LJ := "Mengembangkan dan menerapkan strategi pemasaran"
	hasil := GetDataListJobdeskDeskripsi(Deskripsi_LJ, MongoConn, "data_listjobdesk")
	fmt.Println(hasil)

}

func TestDeleteDataListJobdeskDeadline(t *testing.T) {
	Deadline_LJ := "High"
	hasil := DeleteDataListJobdeskDeadline(Deadline_LJ, MongoConn, "data_listjobdesk")
	fmt.Println(hasil)

}
