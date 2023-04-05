// Package util contains utility functions
package util

import (
	"os"
	"path"

	"github.com/rs/zerolog"

	"github.com/fsnotify/fsnotify"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

const pathToLogConfig = "configuration/livesettings.json"
const pathToConfig = "configuration/constsettings.json"
const logLevel = "Logging.LogLevel.Default"

// SetConstValues gets constant values from the file and injects them
func SetConstValues() {
	currentPath, _ := os.Getwd()
	fullPath := path.Join(currentPath, pathToConfig)
	viper.SetConfigFile(fullPath)
	viper.SetConfigType("json")
	err := viper.ReadInConfig() // Find and read the config file
	// just use the default value(s) if the config file was not found
	if _, ok := err.(*os.PathError); ok {
		log.Warn().Msgf("No config file '%s' not found. Using default values", fullPath)
	} else if err != nil { // Handle other errors that occurred while reading the config file
		log.Err(err).Msgf("error while reading the config file")
	}
	viper.SetDefault("CannotReadPayloadMsg", "Cannot read payload")
	viper.SetDefault("PayloadMissingMsg", "Payload is missing")
	viper.SetDefault("CannotParsePayloadMsg", "Cannot parse payload")
	viper.SetDefault("UsersCollection", "users")
	viper.SetDefault("ConfirmationsCollection", "confirmations")
	viper.SetDefault("GeographyCollection", "geography")
	viper.SetDefault("TripsCollection", "trips")
	viper.SetDefault("TripDetailsView", "tripdetail")
	viper.SetDefault("ConversationsCollection", "conversations")
	viper.SetDefault("RequestCollection", "requests")
	viper.SetDefault("ConformationCodeSubject", "Carpool üyeliğiniz için onay kodu")
	viper.SetDefault("ConfirmationCodeMessage", "Merhaba %s,<br>Bi' Dünya Oy'a Hoşgeldiniz.<br>Üyeliğinizi onaylamak için gerekli onay kodunuz: <font size= \"5\" weight=\"bold\">%s</font><br>Üyeliğinizi onaylamak için <a href=\"http://bidunyaoy.com/user/confirm/%s\">buraya</a> tıklayarak onay kodunuzu girebilirsiniz.<br><a href=\"mailto:info@bidunyaoy.com\">")
	viper.SetDefault("InvitationMessage", "Değerli üyemiz, %s kullanımız sizin bu yolculukla ilgilenebileceğinizi düşündü")
	viper.SetDefault("InvitationSubject", "Yolculuk talebiniz var")
	viper.SetDefault("ApprovalSubject", "Yolculuk talebi onay durumu güncellemesi")
	viper.SetDefault("ApprovalMessagePositive", "Değerli üyemiz, %s kullanıcımız görüşmekte olduğunuz yolculuk talebini onayladı")
	viper.SetDefault("ApprovalMessageNegative", "Değerli üyemiz, %s kullanıcımız görüşmekte olduğunuz yolculuk talebini reddetti")
	viper.SetDefault("RequestStates", []string{"Aktif", "Görüşmede", "Tamamlandı"})
}

// SetLogLevels gets configuration values from the file and injects them
func SetLogLevels() {
	currentPath, _ := os.Getwd()
	fullPath := path.Join(currentPath, pathToLogConfig)
	viper.SetConfigFile(fullPath)
	viper.SetConfigType("json")
	err := viper.ReadInConfig() // Find and read the config file
	// just use the default value(s) if the config file was not found
	if _, ok := err.(*os.PathError); ok {
		log.Warn().Msgf("No config file '%s' not found. Using default values", fullPath)
	} else if err != nil { // Handle other errors that occurred while reading the config file
		log.Err(err).Msgf("error while reading the config file")
	} else {
		log.Info().Msgf("Log Level from config: %s", viper.GetString(logLevel))
		setLogLevel(viper.GetString(logLevel))
		// monitor the changes in the config file
		viper.WatchConfig()
		viper.OnConfigChange(func(e fsnotify.Event) {
			log.Info().Msgf("Log Level from config: %s", viper.GetString(logLevel))
			setLogLevel(viper.GetString(logLevel))
		})
	}
}

func setLogLevel(level string) {
	switch level {
	case "Debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		break
	case "Info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		break
	case "Warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		break
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		break
	default:
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	}
}
