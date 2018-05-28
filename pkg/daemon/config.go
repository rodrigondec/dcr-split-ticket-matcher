package daemon

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	flags "github.com/btcsuite/go-flags"
	"github.com/decred/dcrd/dcrutil"
	"github.com/op/go-logging"
	"github.com/pkg/errors"
)

// ErrHelpRequested is the error returned when the help option was requested
// on the command line
var ErrHelpRequested = fmt.Errorf("help requested")

// Config stores the config needed to run an instance of the dcr split ticket
// matcher daemon
type Config struct {
	ConfigFile string `short:"C" long:"configfile" description:"Path to config file"`

	Port             int `long:"port" description:"Port to run the service on"`
	LogLevel         logging.Level
	LogLevelName     string `long:"loglevel" description:"Log Level (CRITICAL, ERROR, WARNING, NOTICE, INFO, DEBUG)"`
	LogDir           string
	KeyFile          string
	CertFile         string
	SplitPoolSignKey string `long:"splitpoolsignkey" description:"WIF private key for signing the split -> ticket intermediate pool fee txo"`
	DataDir          string `long:"datadir" description:"Dir where session and other data will be saved"`

	TestNet  bool   `long:"testnet" description:"Whether this is connecting to a testnet wallet/matcher service"`
	DcrdHost string `long:"dcrdhost" description:"Address of the dcrd daemon"`
	DcrdUser string `long:"dcrduser" description:"Username of the rpc connection to dcrd"`
	DcrdPass string `long:"dcrdpass" description:"Password of the rpc connection to dcrd"`
	DcrdCert string `long:"dcrdcert" description:"Location of the rpc.cert file of dcrd"`

	DcrwHost string `long:"dcrwhost" description:"Address of the dcrwallet daemon"`
	DcrwUser string `long:"dcrwuser" description:"Username of the rpc connection to dcrwallet"`
	DcrwPass string `long:"dcrwpass" description:"Password of the rpc connection to dcrwallet"`
	DcrwCert string `long:"dcrwcert" description:"Location of the rpc.cert file of dcrwallet"`

	MinAmount                   float64       `long:"minamount" description:"Minimum amount to participate on a split ticket (in DCR)"`
	MaxSessionDuration          time.Duration `long:"maxsessionduration" description:"Maximum number of seconds a session may take before being automatically closed"`
	StakeDiffChangeStopWindow   int32         `long:"stakediffchangestopwindow" description:"Stop the matching service when the the stake change is closer than this number of blocks"`
	PublishTransactions         bool          `long:"publishtransactions" description:"Whether to actually publish transactions of successful sessions"`
	ValidateVoteAddressOnWallet bool          `long:"validatevoteaddressonwallet" description:"Whether to validate the vote addresses of participants on the wallet"`
	PoolSubsidyWalletMasterPub  string        `long:"poolsubsidywalletmasterpub" description:"MasterPubKey for deriving addresses where the pool fee is payed to. If empty, pool fee addresses are not validated. Append a :[index] to generate addresses up to the provided index (default: 10000)."`

	AllowPublicSession bool `long:"allowpublicsession" description:"Whether to allow sessions with an empty name (public sessions) in the matcher."`
}

var (
	defaultDataDir     = dcrutil.AppDataDir("dcrstmd", false)
	defaultCfgFilePath = filepath.Join(defaultDataDir, "dcrstmd.conf")
)

func LoadConfig() (*Config, error) {
	var err error

	preCfg := &Config{
		ConfigFile: defaultCfgFilePath,
	}
	preParser := flags.NewParser(preCfg, flags.Default)
	_, err = preParser.Parse()
	if err != nil {
		e, ok := err.(*flags.Error)
		if ok && e.Type == flags.ErrHelp {
			return nil, ErrHelpRequested
		}
		preParser.WriteHelp(os.Stderr)
		return nil, errors.Wrapf(err, "error parsing arguments")
	}

	configFilePath := preCfg.ConfigFile

	cfg := &Config{
		Port:         8475,
		MinAmount:    2.0,
		LogLevel:     logging.INFO,
		LogLevelName: logging.INFO.String(),
		LogDir:       filepath.Join(defaultDataDir, "logs"),
		DataDir:      defaultDataDir,

		KeyFile:  filepath.Join(defaultDataDir, "rpc.key"),
		CertFile: filepath.Join(defaultDataDir, "rpc.cert"),

		DcrdHost: "localhost:19109",
		DcrdUser: "USER",
		DcrdPass: "PASSWORD",
		DcrdCert: filepath.Join(dcrutil.AppDataDir("dcrd", false), "rpc.cert"),

		DcrwHost: "localhost:19110",
		DcrwUser: "USER",
		DcrwPass: "PASSWORD",
		DcrwCert: filepath.Join(dcrutil.AppDataDir("dcrwallet", false), "rpc.cert"),

		MaxSessionDuration:            30,
		StakeDiffChangeStopWindow:     5,
		PublishTransactions:           false,
		AllowPublicSession:            false,
		ValidateVoteAddressOnWallet: false,
		PoolSubsidyWalletMasterPub: "",
	}

	parser := flags.NewParser(cfg, flags.Default)
	err = flags.NewIniParser(parser).ParseFile(configFilePath)
	if err != nil {
		return nil, err
	}

	_, err = parser.Parse()
	if err != nil {
		return nil, err
	}

	logLvl, err := logging.LogLevel(cfg.LogLevelName)
	if err != nil {
		return nil, err
	}
	cfg.LogLevel = logLvl

	return cfg, nil
}
