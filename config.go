package logger

// user config set through a config file (logger section)
type Config struct {
	// log level to print to stderr, don't print any message with log level below this one
	// for example by setting log level to Info5 (int value 55) every message
	// with log level value bigger than 55, for example Verbose3 (int value 73)
	// will not be printed to stderr
	// lines are always saved to file if file is enabled (to look at log lines in detail)
	LogLevel LogLevel `json:"log_level,omitempty"`
	// specify a log directory if you want to duplicate all logs into a file by default
	// in your programs provide a way to override this with --log-dir flag
	LogDir string `json:"log_dir,omitempty"`
	// environment variable name that will allow us to identify program instance
	// for example HOSTNAME can be used inside container to get container id
	// if this variable is set and os.Getenv(ContainerIdVar) is not empty then
	// LodDir = path.Join(LogDir, os.Getenv(ContainerIdVar))
	ContainerIdVarName string `json:"container_id_var_name,omitempty"` // switch to NONE to not put log files to a separate directory
	// if we want to print goroutine id with each log message
	UseTid *bool `json:"use_tid,omitempty"`
	// time format to use
	TimeFormat string `json:"time_format,omitempty"`
	// log file format
	LogFileFormat string `json:"log_file_format,omitempty"`

	// colorizer for the timestamp. Not JSON-serializable; runtime-only.
	LogTimeColor Colorizer `json:"-"`
}

var Cfg Config = defaultConfig() // this one we use to access config values from anywhere

func defaultConfig() Config {
	useTid := false
	return Config{
		LogLevel:           99,
		LogDir:             "",
		ContainerIdVarName: "HOSTNAME",
		UseTid:             &useTid,
		TimeFormat:         "2006/Jan/02 15:04:05",
		LogFileFormat:      "02_Jan_2006_15_04_05.jsonl",
		LogTimeColor:       DimText, // soft “dim white/gray”
	}
}

func InitializeConfig(userConfig *Config) {
	// If not provided - just use defaultConfig
	if userConfig == nil {
		Log(Info, Color, "%s config is %s, keeping %s", "logger", "not provided", "default logger config")
		return
	}
	Log(Info, Color, "%s config was %s, using %s", "logger", "provided", "user config")
	// If local Config is provided - use it
	Cfg = *userConfig

	// apply defaults for every missing value
	defaultConfig := defaultConfig()
	ApplyDefaults(&Cfg, defaultConfig, func(field string, defVal any) {
		Log(
			Info, Color,
			"%s field is %s in %s configuration. Using default value: %v",
			field, "missing", "logger", PrettyValue(defVal),
		)
	})

	LogPrettyJSON(Cfg, "effective config")

	if Cfg.LogDir != "" {
		// this function will change Cfg.LoggerFilePath and Cfg.LoggerFile
		err, errMsg := OpenLoggerFile(Cfg.LogDir)
		QuitIfErrorLoggerIndependent(err, errMsg)
	}
}
