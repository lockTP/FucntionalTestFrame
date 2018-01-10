package common

import (
	"database/sql"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"

	log "github.com/Sirupsen/logrus"
	"github.com/jinzhu/gorm"

	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
)

var (
	Database, User, Password, Host, Port string
	GDB                                  *gorm.DB
	DB                                   *sql.DB
)

type Tables struct {
	Tables_in_testdWarehouse []string
}

func SetUPDB() {
	isDevelopment := viper.GetBool("isDevelopment")
	dialect := viper.GetString("db.dialect")
	User := viper.GetString("db.user")
	Password := viper.GetString("db.password")
	Database := viper.GetString("db.database")
	Host := viper.GetString("db.host")
	Port := viper.GetString("db.port")

	url := User + ":" + Password + "@tcp(" + Host + ":" + Port + ")/" + Database + "?charset=utf8&parseTime=True&loc=Local"
	//fmt.Println(url)
	//url = "web:123456@tcp(192.168.1.204:3306)/dTicket_new?charset=utf8"
	db, err := gorm.Open(dialect, url)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect database")
	}

	db.LogMode(isDevelopment)
	//db.LogMode()
	//db.SetLogger(initLog())
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(10)

	GDB = db


	//连接mysql
	db2, err := sql.Open(dialect, url)
	if(err != nil) {
		fmt.Println("error: ", err)
	}

	//用于保持连接数
	db2.SetMaxIdleConns(10)
	//用于设置最大打开的连接数，默认值为0表示不限制。
	db2.SetMaxOpenConns(10)
	connectErr := db2.Ping()      //测试连接
	if connectErr != nil {
		panic("failed to connect database:" + connectErr.Error())
	}
	DB = db2
}

//CreateAndOpen database
func SetupNewSchema() error {

	User = viper.GetString("db.user")
	Password = viper.GetString("db.password")
	Database = viper.GetString("db.database")
	Host = viper.GetString("db.host")
	Port = viper.GetString("db.port")

	DB, err := sql.Open("mysql", User+":"+Password+"@tcp("+Host+":"+Port+")/")
	// fmt.Println(User + ":" + Password + "@tcp(" + Host + ":" + Port + ")/")
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Cannot open the database URL : %s ", err)
	}
	defer DB.Close()

	_, err = DB.Exec("CREATE DATABASE IF NOT EXISTS " + Database)
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Cannot open the database URL : %s ", err)
	}
	log.Info("create database schema successful.")
	DB.Close()
	return nil
}

func readFromPath(path string) string {
	fi, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer fi.Close()
	fd, err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}

//CreateTables initial the database table structure
func CreateTables() error {

	return ExecuteSQLScriptFile("./../../config/dbstruct.sql")

}

func InitMetadata() error {

	return ExecuteSQLScriptFile("./../../config/exported8.sql")

}
func InitMetatables() error {

	return ExecuteSQLScriptFile("./../../config/metatable.sql")

}

//CreateTables initial the database table structure
func ExecuteSQLScriptFile(filePath string) error {
	log.Infof("Starting execute sql file " + filePath)

	fmt.Println(User + ":" + Password + "@tcp(" + Host + ":" + Port + ")/" + Database)
	DB, err := sql.Open("mysql", User+":"+Password+"@tcp("+Host+":"+Port+")/"+Database)
	if err != nil {
		log.Fatal(err)
	}

	sqlscript := readFromPath(filePath)
	tx, err := DB.Begin()
	if err != nil {
		log.Fatal(err)
		return errors.New("Cannot begin the transaction.")
	}
	defer tx.Rollback()
	// log.Debugln(sqlscript)
	log.Debugln("***********************")
	sqlstatments := strings.Split(sqlscript, ";")
	for _, ele := range sqlstatments {
		log.Debugln("++++++++++++++++++++++++++++++++++++start")
		preparedString := strings.Trim(ele, " ") + ";"

		log.Debugf("[Prepared Statements]: %s", preparedString)

		reg := regexp.MustCompile(`DROP.*\;`)

		if preparedString == ";" {
			continue
		}
		if strings.Contains(preparedString, "--") && !strings.Contains(preparedString, "INSERT INTO") {
			droptablestring := reg.FindString(preparedString)
			if droptablestring == "" {
				continue
			} else {
				preparedString = droptablestring
			}

		}
		if strings.Contains(preparedString, "/*") {
			continue
		}
		rs, err := tx.Exec(preparedString)
		if err != nil {
			log.Warnf("SQL script error: %s, %s", err, preparedString)
			return fmt.Errorf("SQL script error: %s", err)
		}
		rowAffected, err := rs.RowsAffected()
		if err != nil {
			log.Warnf("RowsAffected error: %s", err)
			return fmt.Errorf("RowsAffected error: %s", err)
		}
		log.Debugf("rowsAffected: %d", rowAffected)

		log.Debugln("++++++++++++++++++++++++++++++++++++ends")
	}

	err = tx.Commit()
	if err != nil {
		log.Warnf("commit error: %s", err)
		return fmt.Errorf("commit error: %s", err)
	}
	DB.Close()
	return nil
}

func Showtables() ([]string, error) {
	var result []string
	var Tables_in_testdWarehouse string
	rows, err := GDB.Raw("show tables;").Rows()
	if err != nil {
		log.Info("Get tables list error.")
		return nil, fmt.Errorf("get tables list error")
	}
	for rows.Next() {
		rows.Scan(&Tables_in_testdWarehouse)
		// log.Printf("get data, id: %s", Tables_in_testdWarehouse)
		result = append(result, Tables_in_testdWarehouse)
	}
	return result, nil
}

//CreateAndOpen database
func DestroySchema() error {
	// log.Infof("Starting destroy schema.")
	User = viper.GetString("db.user")
	Password = viper.GetString("db.password")
	Database = viper.GetString("db.database")
	Host = viper.GetString("db.host")
	Port = viper.GetString("db.port")

	DB, err := sql.Open("mysql", User+":"+Password+"@tcp("+Host+":"+Port+")/")
	// fmt.Println(User + ":" + Password + "@tcp(" + Host + ":" + Port + ")/")
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Cannot open the database URL : %s ", err)
	}
	defer DB.Close()

	_, err = DB.Exec("DROP DATABASE IF EXISTS " + Database + " ;")
	if err != nil {
		log.Fatal(err)
		return fmt.Errorf("Cannot destroy the database URL : %s ", err)
	}
	log.Info("Drop database schema Successful. ")
	DB.Close()
	return nil
}


func Init() {
	User = viper.GetString("db.user")
	Password = viper.GetString("db.password")
	Database = viper.GetString("db.database")
	Host = viper.GetString("db.host")
	Port = viper.GetString("db.port")
}