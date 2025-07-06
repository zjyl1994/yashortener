package startup

import (
	"os"
	"strconv"

	_ "github.com/joho/godotenv/autoload"
	gorm_logrus "github.com/onrik/gorm-logrus"
	"github.com/sirupsen/logrus"
	"github.com/zjyl1994/yashortener/infra/model"
	"github.com/zjyl1994/yashortener/infra/utils"
	"github.com/zjyl1994/yashortener/infra/vars"
	"github.com/zjyl1994/yashortener/server"
	"github.com/zjyl1994/yashortener/web"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func Start() (err error) {
	vars.ListenAddr = utils.Coalesce(os.Getenv("YASHORT_LISTEN"), ":19278")
	vars.BaseURL = os.Getenv("YASHORT_BASE_URL")
	vars.DBPath = utils.Coalesce(os.Getenv("YASHORT_DB_PATH"), "./yashortener.db")
	vars.AdminUser = utils.Coalesce(os.Getenv("YASHORT_ADMIN_USER"), "admin")
	vars.AdminPass = utils.Coalesce(os.Getenv("YASHORT_ADMIN_PASS"), "pass")
	vars.DebugMode, _ = strconv.ParseBool(os.Getenv("YASHORT_DEBUG"))
	if vars.DebugMode {
		logrus.SetLevel(logrus.DebugLevel)
	}
	vars.AnonymousCreate, _ = strconv.ParseBool(os.Getenv("YASHORT_ANONYMOUS_CREATE"))

	_, err = os.Stat("./web")
	if err != nil {
		if os.IsNotExist(err) {
			err = utils.ExtractDataTo(web.EMFS, "./web")
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	vars.DB, err = gorm.Open(sqlite.Open(vars.DBPath), &gorm.Config{
		Logger:         gorm_logrus.New(),
		TranslateError: true,
	})
	if err != nil {
		return err
	}
	err = vars.DB.Exec("PRAGMA journal_mode=WAL;").Error
	if err != nil {
		return err
	}
	err = vars.DB.AutoMigrate(&model.Link{}, &model.Access{})
	if err != nil {
		return err
	}

	logrus.Infoln("YASHORTENER running on", vars.ListenAddr)
	return server.Run(vars.ListenAddr)
}
