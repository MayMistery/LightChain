package cmd

type Config struct {
	Dir     string
	Host    string
	hostNet string
	Chain   string
	Port    int
	Server  bool
	Client  bool
	Local   bool
	Dep     bool
	Alive   bool
}
