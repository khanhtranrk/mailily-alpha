package account

type Account struct {
    Host string `yaml:"host"`
    Port int32 `yaml:"port"`
    Username string `yaml:"username"`
    Password string `yaml:"password"`
}
