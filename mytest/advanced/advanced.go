package advanced

import "time"

// CLI represents a command-line application with subcommands and time types.
// :quickclop
type CLI struct {
	// Global options
	Verbose    bool          `clop:"-v; --verbose" usage:"Enable verbose output"`
	ConfigFile string        `clop:"-c; --config" usage:"Configuration file path"`
	Timeout    time.Duration `clop:"-t; --timeout" usage:"Operation timeout"`
	
	// Subcommands
	Server *ServerCmd `clop:"subcmd=server" description:"Start the server"`
	Client *ClientCmd `clop:"subcmd=client" description:"Run as client"`
}

// ServerCmd represents the server subcommand.
// :quickclop
type ServerCmd struct {
	Port     int           `clop:"-p; --port" usage:"Server port"`
	Host     string        `clop:"-h; --host" usage:"Server host"`
	CertFile string        `clop:"--cert" usage:"TLS certificate file"`
	KeyFile  string        `clop:"--key" usage:"TLS key file"`
	StartAt  *time.Time    `clop:"--start-at" usage:"Schedule server start time (RFC3339 format)"`
	Interval time.Duration `clop:"--interval" usage:"Health check interval"`
}

// ClientCmd represents the client subcommand.
// :quickclop
type ClientCmd struct {
	Server    string        `clop:"-s; --server" usage:"Server address"`
	Timeout   *time.Duration `clop:"-t; --timeout" usage:"Connection timeout"`
	Retries   int           `clop:"-r; --retries" usage:"Number of retries"`
	Operation string        `clop:"args=operation" usage:"Operation to perform"`
}
