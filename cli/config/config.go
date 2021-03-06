// Package config holds CLI configuration.
package config

import (
	"time"

	"github.com/neilotoole/sq/drivers/userdriver"
	"github.com/neilotoole/sq/libsq/core/errz"
	"github.com/neilotoole/sq/libsq/core/stringz"
	"github.com/neilotoole/sq/libsq/source"
)

// Config holds application config/session data.
type Config struct {
	// Version is the config version. This will allow sq to
	// upgrade config files if needed.
	Version string `yaml:"version" json:"version"`

	// Defaults contains default settings, such as output format.
	Defaults Defaults `yaml:"defaults" json:"defaults"`

	// Sources is the set of data sources.
	Sources *source.Set `yaml:"sources" json:"sources"`

	// Ext holds sq config extensions, such as user driver config.
	Ext Ext `yaml:"-" json:"-"`
}

func (c *Config) String() string {
	return stringz.SprintJSON(c)
}

// Ext holds additional config (extensions) loaded from other
// config files, e.g. ~/.config/sq/ext/*.sq.yml
type Ext struct {
	UserDrivers []*userdriver.DriverDef `yaml:"user_drivers" json:"user_drivers"`
}

// Defaults contains default config values.
type Defaults struct {
	// Format is the default output format: json, table, etc.
	Format Format `yaml:"output_format" json:"output_format"`

	// Header determines if a header should be printed (if relevant
	// for the output format).
	Header bool `yaml:"output_header" json:"output_header"`

	// PingTimeout is the allowed time for a ping.
	PingTimeout time.Duration `yaml:"ping_timeout" json:"ping_timeout"`

	// ShellCompletionTimeout is the time allowed for the shell
	// completion callback to execute.
	ShellCompletionTimeout time.Duration `yaml:"shell_completion_timeout" json:"shell_completion_timeout"`
}

// New returns a config instance with default options set.
func New() *Config {
	cfg := &Config{}

	// By default, we want header to be true; this is
	// ugly wrt initCfg, as the zero value of a bool
	// is false, but we actually want it to be true for Header.
	cfg.Defaults.Header = true

	initCfg(cfg)
	return cfg
}

// initCfg checks if required values are present, and if not, sets them.
func initCfg(cfg *Config) {
	if cfg.Sources == nil {
		cfg.Sources = &source.Set{}
	}

	if cfg.Defaults.Format == "" {
		cfg.Defaults.Format = FormatTable
	}

	if cfg.Defaults.PingTimeout == 0 {
		// Probably should be setting this in the New function,
		// but we haven't yet defined cli's behavior wrt
		// a zero timeout. Does it mean no timeout?
		cfg.Defaults.PingTimeout = 10 * time.Second
	}

	if cfg.Defaults.ShellCompletionTimeout == 0 {
		cfg.Defaults.ShellCompletionTimeout = time.Millisecond * 500
	}
}

// Format is a sq output format such as json or xml.
type Format string

// UnmarshalText implements encoding.TextUnmarshaler.
func (f *Format) UnmarshalText(text []byte) error {
	switch Format(text) {
	default:
		return errz.Errorf("unknown output format %q", string(text))
	case FormatJSON, FormatJSONA, FormatJSONL, FormatTable, FormatRaw,
		FormatHTML, FormatMarkdown, FormatXLSX, FormatXML, FormatCSV, FormatTSV:
	}

	*f = Format(text)
	return nil
}

// Output format values.
const (
	FormatJSON     Format = "json"
	FormatJSONL    Format = "jsonl"
	FormatJSONA    Format = "jsona"
	FormatTable    Format = "table"
	FormatRaw      Format = "raw"
	FormatHTML     Format = "html"
	FormatMarkdown Format = "markdown"
	FormatXLSX     Format = "xlsx"
	FormatXML      Format = "xml"
	FormatCSV      Format = "csv"
	FormatTSV      Format = "tsv"
)
