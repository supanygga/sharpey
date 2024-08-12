package testhelpers

const (
	defaultImage    = "postgres:16-alpine"
	defaultUser     = "user"
	defaultPassword = "pass"
	defaultDatabase = "database"
)

// options.
type options struct {
	image    string
	user     string
	password string
	database string
}

// defaultOptions
func defaultOptions() *options {
	return &options{
		image:    defaultImage,
		user:     defaultUser,
		password: defaultPassword,
		database: defaultDatabase,
	}
}

// apply.
func (o *options) apply(opts ...WithFunc) {
	for _, opt := range opts {
		opt(o)
	}
}

// WithFunc.
type WithFunc func(o *options)

// WithImage.
func WithImage(image string) WithFunc {
	return func(o *options) {
		o.image = image
	}
}

// WithUser.
func WithUser(user string) WithFunc {
	return func(o *options) {
		o.user = user
	}
}

// WithPassword.
func WithPassword(password string) WithFunc {
	return func(o *options) {
		o.password = password
	}
}

// WithDatabase.
func WithDatabase(database string) WithFunc {
	return func(o *options) {
		o.database = database
	}
}
