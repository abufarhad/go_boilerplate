package conn

import (
	"core/infra/config"
	"core/infra/logger"
	"core/model"
	"database/sql"
	"fmt"
	gomysql "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
	"io/ioutil"
	"os"
	"time"
)

var db *gorm.DB

type DbErrors struct {
	*gomysql.MySQLError
}

func getdsn() string {
	envErr := godotenv.Load()
	if envErr != nil {
		logger.Error("Error loading .env file", envErr)
	}

	var dsn = fmt.Sprintf("port=%s host=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_PORT"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"))

	logger.Info(dsn)
	return dsn
}

func ConnectDb() {
	envErr := godotenv.Load()
	if envErr != nil {
		logger.Error("Error loading .env file", envErr)
	}
	conf := config.Db()

	logger.Info("connecting to mysql at " + conf.Host + ":" + conf.Port + "...")

	logMode := gormlogger.Silent
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", conf.User, conf.Pass, conf.Host, conf.Port, conf.Schema)
	logger.Info(dsn)

	dB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		PrepareStmt: true,
		Logger:      gormlogger.Default.LogMode(logMode),
	})

	//dB, err := gorm.Open(postgres.Open(getdsn()), &gorm.Config{
	//	PrepareStmt: true,
	//	Logger:      gormlogger.Default.LogMode(logMode),
	//})

	if err != nil {
		logger.Error("mysql connection error ", err)
		panic(err)
	}

	sqlDb, err := dB.DB()
	if err != nil {
		panic(err)
	}

	if conf.MaxIdleConn != 0 {
		sqlDb.SetMaxIdleConns(conf.MaxIdleConn)
	}
	if conf.MaxOpenConn != 0 {
		sqlDb.SetMaxOpenConns(conf.MaxOpenConn)
	}
	if conf.MaxConnLifetime != 0 {
		sqlDb.SetConnMaxLifetime(conf.MaxConnLifetime * time.Second)
	}

	db = dB
	populateDbModel(db)

	logger.Info("mysql connection successful...")
	//logger.Info("postgres connection successful...")
}

func Db() *gorm.DB {
	return db
}

func Migrate() {
	database, e := sql.Open("postgres", getdsn())
	if e != nil {
		logger.Error("gooes connection error ", e)
		panic(e)
	}

	logger.Info("postgres connection successful...")
	logger.Info("Data migration starting ...")
	if err := goose.Run("up", database, "/", ""); err != nil {
		panic(err)
	}
	logger.Info("Data migration Success")

	err := database.Close()
	if err != nil {
		panic(err)
	}
}

func populateDbModel(db *gorm.DB) {
	db.AutoMigrate(
		&model.AboutUs{},
	)
}

type Seed struct {
	Name string
	Run  func(db *gorm.DB, truncate bool) error
}

func SeedAll() []Seed {
	//return []Seed{
	//	{
	//		Name: "CreateRoles",
	//		Run: func(db *gorm.DB, truncate bool) error {
	//			if err := seedRoles(db, "/infra/seed/roles.json", truncate); err != nil {
	//				return err
	//			}
	//			return nil
	//		},
	//	},
	//	{
	//		Name: "CreatePermissions",
	//		Run: func(db *gorm.DB, truncate bool) error {
	//			if err := seedPermissions(db, "/infra/seed/permissions.json"); err != nil {
	//				return err
	//			}
	//			return nil
	//		},
	//	},
	//}

	return []Seed{}
}

//func seedRoles(db *gorm.DB, jsonfilPath string, truncate bool) error {
//	file, _ := readSeedFile(jsonfilPath)
//	roles := []domain.Role{}
//
//	_ = json.Unmarshal([]byte(file), &roles)
//
//	if truncate {
//		db.Exec("TRUNCATE TABLE core_user.role_permissions;")
//		db.Exec("TRUNCATE TABLE core_user.permissions;")
//		db.Exec("TRUNCATE TABLE core_user.roles;")
//	}
//
//	var count int64
//
//	db.Model(&domain.Role{}).Count(&count)
//	if count == 0 {
//		db.AboutContent(&roles)
//	}
//
//	return nil
//}

func readSeedFile(jsonfilPath string) ([]byte, error) {
	BaseDir, _ := os.Getwd()
	seedFile := BaseDir + jsonfilPath
	if BaseDir == "/" {
		seedFile = jsonfilPath
	}
	fmt.Println("seed folder: ", seedFile)

	return ioutil.ReadFile(seedFile)
}
