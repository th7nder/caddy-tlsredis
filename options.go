package storageredis

const (

	// DefaultAESKey needs to be 32 bytes long
	DefaultAESKey = "redistls-01234567890-caddytls-32"

	// DefaultKeyPrefix defines the default prefix in KV store
	DefaultKeyPrefix = "caddytls"

	// DefaultValuePrefix sets a prefix to KV values to check validation
	DefaultValuePrefix = "caddy-storage-redis"

	// DefaultRedisHost define the Redis instance host
	DefaultRedisHost = "127.0.0.1"

	// DefaultRedisPort define the Redis instance port
	DefaultRedisPort = "6379"

	// DefaultRedisPassword define the Redis instance password, if any
	DefaultRedisPassword = ""

	// DefaultRedisTimeout define the Redis wait time in (s)
	DefaultRedisTimeout = 5
)

// Options is option to set plugin configuration
type Options struct {
	Host        string
	Port        string
	DB          int
	Password    string
	Timeout     int
	KeyPrefix   string
	ValuePrefix string
	AESKey      string
	TLSEnabled  bool
	TLSSecure   bool
}

// GetOptions gets default options
func GetOptions() *Options {
	return ParseOptions(&Options{})
}

// ParseOptions generate options from env or default
func ParseOptions(opt *Options) *Options {
	if opt.Host == "" {
		opt.Host = DefaultRedisHost
	}
	if opt.Port == "" {
		opt.Port = DefaultRedisPort
	}
	if opt.Password == "" {
		opt.Password = DefaultRedisPassword
	}
	if opt.Timeout == 0 {
		opt.Timeout = DefaultRedisTimeout
	}
	if opt.KeyPrefix == "" {
		opt.KeyPrefix = DefaultKeyPrefix
	}
	if opt.ValuePrefix == "" {
		opt.ValuePrefix = DefaultValuePrefix
	}
	if opt.AESKey == "" {
		opt.AESKey = DefaultAESKey
	}

	return opt
}

// GetAESKeyByte get aes key as byte
func (op *Options) GetAESKeyByte() []byte {
	return []byte(op.AESKey)
}
