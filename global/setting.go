package global

import (
	"github.com/mwqnice/oh-admin/pkg/logger"
	"github.com/mwqnice/oh-admin/pkg/setting"
)

var (
	ServerSetting   *setting.ServerSettingS
	AppSetting      *setting.AppSettingS
	EmailSetting    *setting.EmailSettingS
	JWTSetting      *setting.JWTSettingS
	DatabaseSetting *setting.DatabaseSettingS
	Logger          *logger.Logger
	TracerSetting   *setting.TracerSettingS
)
