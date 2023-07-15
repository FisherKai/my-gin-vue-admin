package system

import (
	"context"
	"database/sql"
	"demo-gin-1/global"
	"demo-gin-1/model/system/request"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"sort"
)

const (
	Mysql           = "mysql"
	InitSuccess     = "[%v]-->初始化✅\n"
	InitDataExist   = "[%v]-->初始化数据已存在\n"
	InitDataFailed  = "[%v]-->初始化数据失败❌️！\n"
	InitDataSuccess = "[%v]--> 初始化数据✅\n"
)

const (
	InitOrderSystem   = 10
	InitOrderInternal = 1000
	InitOrderExternal = 100000
)

var (
	ErrMissingDBContext        = errors.New("missing db in context")
	ErrMissingDependentContext = errors.New("missing dependent value in context")
	ErrDBTypeMismatch          = errors.New("db type mismatch")
)

type SubInitiaLizer interface {
	InitializerName() string
	MigrateTable(ctx context.Context) (next context.Context, err error)
	InitializeData(ctx context.Context) (next context.Context, err error)
	TableCreated(ctx context.Context) bool
	DataInserted(ctx context.Context) bool
}

type orderedInitializer struct {
	order int
	SubInitiaLizer
}

type initSlice []*orderedInitializer

func (i initSlice) Len() int {
	return len(i)
}

func (a initSlice) Less(i, j int) bool {
	return a[i].order < a[j].order
}

func (a initSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

var (
	initializers initSlice
	cache        map[string]*orderedInitializer
)

type InitDBService struct{}

type TypedDBInitHandler interface {
	EnsureDB(ctx context.Context, conf *request.InitDB) (context.Context, error)
	WriteConfig(ctx context.Context) error
	InitTables(ctx context.Context, inits initSlice) error
	InitData(ctx context.Context, inits initSlice) error
}

func (initDBService *InitDBService) InitDB(conf request.InitDB) (err error) {
	ctx := context.TODO()
	if len(initializers) == 0 {
		return errors.New("无可用初始化过程，请检查初始化是否已执行完成")
	}
	sort.Sort(&initializers)
	var initHandler TypedDBInitHandler
	switch conf.DBType {
	case "mysql":
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "mysql")
	default:
		initHandler = NewMysqlInitHandler()
		ctx = context.WithValue(ctx, "dbtype", "mysql")
	}

	ctx, err = initHandler.EnsureDB(ctx, &conf)
	if err != nil {
		return err
	}

	db := ctx.Value("db").(*gorm.DB)
	global.GVA_DB = db

	if err = initHandler.InitTables(ctx, initializers); err != nil {
		return err
	}
	if err = initHandler.InitData(ctx, initializers); err != nil {
		return err
	}

	if err = initHandler.WriteConfig(ctx); err != nil {
		return err
	}
	initializers = initSlice{}
	cache = map[string]*orderedInitializer{}
	return nil
}

func createTables(ctx context.Context, inits initSlice) error {
	next, cancel := context.WithCancel(ctx)
	defer func(c func()) { c() }(cancel)
	for _, init := range inits {
		if init.TableCreated(next) {
			continue
		}
		if n, err := init.MigrateTable(next); err != nil {
			return err
		} else {
			next = n
		}
	}
	return nil
}

// 创建数据库 （EnsureDB中调用）
func createDatabase(dsn string, driver string, createSql string) error {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		return err
	}
	defer func(db *sql.DB) {
		err = db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
	if err = db.Ping(); err != nil {
		return err
	}
	_, err = db.Exec(createSql)
	return err
}
