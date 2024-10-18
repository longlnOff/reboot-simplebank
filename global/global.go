package global

import (
	"database/sql"

	"github.com/longln/reboot-simplebank/pkg/logger"
	"github.com/longln/reboot-simplebank/pkg/setting"
)


var (
	Config setting.GlobalConfig
	Logger *logger.LoggerZap
	Database *sql.DB
)