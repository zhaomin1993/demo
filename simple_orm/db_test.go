package simple_orm

import (
	"fmt"
	"testing"
	"time"
)

func TestMysql_db_OrmQuery(t *testing.T) {
	type MyFdCache struct {
		Id            int64     `json:"id" db:"id"`
		Dep           string    `json:"dep" db:"dep"`
		Arr           string    `json:"arr" db:"arr"`
		FlightNo      string    `json:"flightNo" db:"flightNo"`
		FlightTime    string    `json:"flightTime" db:"flightTime"`
		IsCodeShare   bool      `json:"isCodeShare" db:"isCodeShare"`
		Tax           int       `json:"tax" db:"tax"`
		Yq            int       `json:"yq" db:"yq"`
		IbePrice      int       `json:"ibePrice" db:"ibe_price"`
		CtripPrice    int       `json:"ctripPrice" db:"ctrip_price"`
		OfficialPrice int       `json:"officialPrice" db:"official_price"`
		Cabin         string    `json:"cabin" db:"cabin"`
		FlightDate    time.Time `json:"flightDate" db:"flightDate"`
		Uptime        time.Time `json:"uptime" db:"uptime"`
	}
	db := &mysql_db{}
	db.mysql_open()
	defer db.mysql_close()

	var res []MyFdCache
	err := db.OrmQuery(&res, `select * from mf_fd_cache`)
	fmt.Println(err)
	fmt.Println(res)
}
