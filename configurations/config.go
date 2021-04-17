package configurations

var (
	Server *server
	MySQL *mysql
)

func init() {
	MySQL = setupMySQL()
	Server = setupServer()
}
