package simple_orm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"reflect"
	"time"

	"github.com/go-sql-driver/mysql"
)

const (
	dbhostsip  = "127.0.0.1:3309"
	dbusername = "root"
	dbpassword = "123456"
	dbname     = "mytest"
)

type mysql_db struct {
	db *sql.DB
}

func (f *mysql_db) mysql_open() {
	db, err := sql.Open("mysql", dbusername+":"+dbpassword+"@tcp("+dbhostsip+")/"+dbname+"?charset=utf8")
	db.SetConnMaxLifetime(time.Minute * 5)
	db.SetMaxOpenConns(200)
	db.SetMaxIdleConns(20)
	checkErr(err)
	p("链接数据库成功...........已经打开")
	f.db = db

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func p(a ...interface{}) {
	//fmt.Println(i, "=========================", i)
	fmt.Println(a...)
	//i++
}
func (f *mysql_db) mysql_close() {
	f.db.Close()
	p("链接数据库成功...........已经关闭")
}

// OrmQuery 根据sql和参数将查询结果通过反射映射到切片中
func (m mysql_db) OrmQuery(res interface{}, sqlStr string, args ...interface{}) error {
	rType := reflect.TypeOf(res)
	if rType.Kind() != reflect.Ptr {
		return errors.New("res must be *[]Struct")
	}
	if rType.Elem().Kind() != reflect.Slice {
		return errors.New("res must be *[]Struct")
	}
	if rType.Elem().Elem().Kind() != reflect.Struct {
		return errors.New("res must be *[]Struct")
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	rows, err := m.db.QueryContext(ctx, sqlStr, args...)
	defer func() {
		if rows != nil {
			_ = rows.Close()
		}
	}()
	if err != nil {
		return err
	}

	columns, err := rows.ColumnTypes()
	if err != nil {
		return err
	}
	scans := make([]interface{}, len(columns))
	names := make([]string, len(columns))
	for i := range columns {
		scans[i] = reflect.New(columns[i].ScanType()).Interface()
		names[i] = columns[i].Name()
	}
	model := rType.Elem().Elem()
	var fieldIndexMap = make(map[string]int, model.NumField())
	for j := 0; j < model.NumField(); j++ {
		tag, ok := model.Field(j).Tag.Lookup("db")
		if ok {
			fieldIndexMap[tag] = j
		}
	}
	valueList := make([]reflect.Value, 0)
	for rows.Next() {
		m := reflect.New(model).Elem()
		if err = rows.Scan(scans...); err != nil {
			return fmt.Errorf("rows scan error:%s", err.Error())
		}
		for i := range scans {
			columnName := names[i]
			index, ok := fieldIndexMap[columnName]
			if !ok {
				continue
			}
			fieldValue := m.Field(index)
			switch (scans[i]).(type) {
			case *int8:
				val := *(scans[i].(*int8))
				// 需要什么类型的数据可自己手动增加
				switch fieldValue.Type().Kind() {
				case reflect.Bool:
					var value bool
					if val != 0 {
						value = true
					}
					fieldValue.SetBool(value)
				case reflect.Int8:
					fieldValue.Set(reflect.ValueOf(val))
				case reflect.Int:
					fieldValue.Set(reflect.ValueOf(int(val)))
				case reflect.Int32:
					fieldValue.Set(reflect.ValueOf(int32(val)))
				case reflect.Int64:
					fieldValue.SetInt(int64(val))
				}
			case *int32:
				val := *(scans[i].(*int32))
				// 需要什么类型的数据可自己手动增加
				switch fieldValue.Type().Kind() {
				case reflect.Int8:
					fieldValue.Set(reflect.ValueOf(int8(val)))
				case reflect.Int:
					fieldValue.Set(reflect.ValueOf(int(val)))
				case reflect.Int32:
					fieldValue.Set(reflect.ValueOf(val))
				case reflect.Int64:
					fieldValue.SetInt(int64(val))
				}
			case *int64:
				val := *(scans[i].(*int64))
				// 需要什么类型的数据可自己手动增加
				switch fieldValue.Type().Kind() {
				case reflect.Int8:
					fieldValue.Set(reflect.ValueOf(int8(val)))
				case reflect.Int:
					fieldValue.Set(reflect.ValueOf(int(val)))
				case reflect.Int32:
					fieldValue.Set(reflect.ValueOf(int32(val)))
				case reflect.Int64:
					fieldValue.SetInt(val)
				}
			case *sql.RawBytes:
				val := *(scans[i].(*sql.RawBytes))
				switch fieldValue.Type().Kind() {
				case reflect.String:
					fieldValue.SetString(string(val))
				}
			case *mysql.NullTime:
				val := scans[i].(*mysql.NullTime)
				fieldValue.Set(reflect.ValueOf(val.Time))
			}
		}
		valueList = append(valueList, m)
	}
	if len(valueList) > 0 {
		rValue := reflect.ValueOf(res).Elem()
		rValue.Set(reflect.Append(rValue, valueList...))
	}
	if err = rows.Err(); err != nil {
		return err
	}
	if err = rows.Close(); err != nil {
		return err
	}
	return nil
}
