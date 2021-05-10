package configs

var (
	// Server config
	Server *server

	// MongoDB Config
	MongoDB *mongodb

	// Constant config
	Constant *constant

	// Auth config
	Auth *auth
)

func init() {
	Server = setupServer()
	MongoDB = setupMongoDB()
	Constant = setupConstant()
	Auth = setupAuth()
}