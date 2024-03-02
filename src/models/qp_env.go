package models

import (
	"errors"
	"os"
	"strconv"
	"strings"

	whatsapp "github.com/nocodeleaks/quepasa/whatsapp"
	"google.golang.org/protobuf/proto"
)

const (
	ENV_WEBAPIPORT = "WEBAPIPORT"
	ENV_WEBAPIHOST = "WEBAPIHOST"

	ENV_DBDRIVER   = "DBDRIVER" // database driver, default sqlite3
	ENV_DBHOST     = "DBHOST"
	ENV_DBDATABASE = "DBDATABASE"
	ENV_DBPORT     = "DBPORT"
	ENV_DBUSER     = "DBUSER"
	ENV_DBPASSWORD = "DBPASSWORD"
	ENV_DBSSLMODE  = "DBSSLMODE"

	ENV_SIGNING_SECRET = "SIGNING_SECRET" // token for hash singing cookies

	ENV_WEBSOCKETSSL        = "WEBSOCKETSSL"      // use ssl for websocket qrcode
	ENV_ENVIRONMENT         = "APP_ENV"           // development | production
	ENV_MIGRATIONS          = "MIGRATIONS"        // enable migrations
	ENV_TITLE               = "APP_TITLE"         // application title for whatsapp id
	ENV_DEBUG_REQUESTS      = "DEBUGREQUESTS"     // debug api and form requests
	ENV_DEBUG_JSON_MESSAGES = "DEBUGJSONMESSAGES" // debug json messages
	ENV_REMOVEDIGIT9        = "REMOVEDIGIT9"
	ENV_SYNOPSISLENGTH      = "SYNOPSISLENGTH"
	ENV_CONVERT_WAVE_TO_OGG = "CONVERT_WAVE_TO_OGG"

	ENV_READRECEIPTS    = "READRECEIPTS"
	ENV_REJECTCALLS     = "REJECTCALLS"
	ENV_GROUPS          = "GROUPS"
	ENV_BROADCASTS      = "BROADCASTS"
	ENV_HISTORYSYNCDAYS = "HISTORYSYNCDAYS"

	ENV_LOGLEVEL            = "LOGLEVEL"
	ENV_WHATSMEOWLOGLEVEL   = "WHATSMEOW_LOGLEVEL"
	ENV_WHATSMEOWDBLOGLEVEL = "WHATSMEOW_DBLOGLEVEL"
)

type Environment struct{}

var ENV Environment

func (*Environment) ShouldConvertWaveToOgg() bool {
	environment, _ := GetEnvBool(ENV_CONVERT_WAVE_TO_OGG, proto.Bool(true))
	return *environment
}

func (*Environment) IsDevelopment() bool {
	environment, _ := GetEnvStr(ENV_ENVIRONMENT)
	if strings.ToLower(environment) == "development" {
		return true
	} else {
		return false
	}
}

// WEBSOCKETSSL => default false
func (*Environment) UseSSLForWebSocket() bool {
	migrations, _ := GetEnvStr(ENV_WEBSOCKETSSL)
	boolMigrations, err := strconv.ParseBool(migrations)
	if err == nil {
		return boolMigrations
	} else {
		return false
	}
}

// MIGRATIONS => Path to database migrations folder
func (*Environment) Migrate() bool {
	migrations, _ := GetEnvStr(ENV_MIGRATIONS)
	boolMigrations, err := strconv.ParseBool(migrations)
	if err == nil {
		return boolMigrations
	} else {
		return true
	}
}

// MIGRATIONS => Path to database migrations folder
func (*Environment) MigrationPath() string {
	migrations, _ := GetEnvStr(ENV_MIGRATIONS)
	_, err := strconv.ParseBool(migrations)
	if err != nil {
		return migrations
	} else {
		return "" // indicates that should use default path
	}
}

func (*Environment) AppTitle() string {
	result, _ := GetEnvStr(ENV_TITLE)
	return result
}

func (*Environment) DEBUGRequests() bool {

	if ENV.IsDevelopment() {
		environment, err := GetEnvBool(ENV_DEBUG_REQUESTS, proto.Bool(true))
		if err == nil {
			return *environment
		}
	}

	return false
}

func (*Environment) DEBUGJsonMessages() bool {

	if ENV.IsDevelopment() {
		environment, err := GetEnvBool(ENV_DEBUG_JSON_MESSAGES, proto.Bool(true))
		if err == nil {
			return *environment
		}
	}

	return false
}

var ErrEnvVarEmpty = errors.New("getenv: environment variable empty")

func GetEnvBool(key string, value *bool) (*bool, error) {
	result := value
	s, err := GetEnvStr(key)
	if err == nil {
		trying, err := strconv.ParseBool(s)
		if err == nil {
			result = &trying
		}
	}
	return result, err
}

func GetEnvStr(key string) (string, error) {
	v := os.Getenv(key)
	if v == "" {
		return v, ErrEnvVarEmpty
	}
	return v, nil
}

func (*Environment) ShouldRemoveDigit9() bool {
	value, _ := GetEnvBool(ENV_REMOVEDIGIT9, proto.Bool(false))
	return *value
}

//#region WHATSAPP SERVICE OPTIONS - WHATSMEOW

func ParseWhatsappBoolean(v string) whatsapp.WhatsappOptionsBoolean {
	if len(v) > 0 {
		switch strings.ToLower(v) {
		case "1", "t", "true", "yes":
			return whatsapp.TrueBooleanType
		case "0", "f", "false", "no":
			return whatsapp.FalseBooleanType
		case "forcedfalse":
			return whatsapp.ForcedFalseBooleanType
		case "forcedtrue":
			return whatsapp.ForcedTrueBooleanType
		}
	}

	return whatsapp.UnknownBooleanType
}

func (*Environment) Broadcasts() whatsapp.WhatsappOptionsBoolean {
	v := os.Getenv(ENV_BROADCASTS)
	return ParseWhatsappBoolean(v)
}

func (*Environment) Groups() whatsapp.WhatsappOptionsBoolean {
	v := os.Getenv(ENV_GROUPS)
	return ParseWhatsappBoolean(v)
}

func (*Environment) ReadReceipts() whatsapp.WhatsappOptionsBoolean {
	v := os.Getenv(ENV_READRECEIPTS)
	return ParseWhatsappBoolean(v)
}

func (*Environment) RejectCalls() whatsapp.WhatsappOptionsBoolean {
	v := os.Getenv(ENV_REJECTCALLS)
	return ParseWhatsappBoolean(v)
}

// Force Default Log Level
func (*Environment) LogLevel() string {
	result, _ := GetEnvStr(ENV_LOGLEVEL)
	return result
}

// Force Default Whatsmeow Log Level
func (*Environment) WhatsmeowLogLevel() string {
	result, _ := GetEnvStr(ENV_WHATSMEOWLOGLEVEL)
	return result
}

// Force Default Whatsmeow DataBase Log Level
func (*Environment) WhatsmeowDBLogLevel() string {
	result, _ := GetEnvStr(ENV_WHATSMEOWDBLOGLEVEL)
	return result
}

// Get history sync days, environment whatsapp service global option
func (*Environment) HistorySync() *uint32 {
	stringValue, err := GetEnvStr(ENV_HISTORYSYNCDAYS)
	if err == nil {
		value, err := strconv.ParseUint(stringValue, 10, 32)
		if err == nil {
			return proto.Uint32(uint32(value))
		}
	}

	return nil
}

//#endregion

// MIGRATIONS => Path to database migrations folder
func (*Environment) SynopsisLength() uint64 {
	stringValue, err := GetEnvStr(ENV_SYNOPSISLENGTH)
	if err == nil {
		value, err := strconv.ParseUint(stringValue, 10, 32)
		if err == nil {
			return value
		}
	}

	return 50
}
