package sys

import (
        "os/user"        
    
        "inspektra/db"       

        "github.com/spf13/pflag"
)

// 
type Config struct {
    db.PgParams
}


func getOSUser() string {
    currentUser, err := user.Current()
    if err != nil {
        return ""
    }
    return currentUser.Username
}

func LoadParams() *Config {
    conf := &Config{}
    
    hHost := "Host address of the postgres database."
    hPort := "Port number at which the postgres instance is listening.\n"
    hPort = hPort + "Default 5432."
    hDatabase := "Database name to connect to."
    hUser := "Username to connect to the postgres database."

    pflag.StringVarP(&conf.Host, "host", "h", "", hHost)
    pflag.IntVarP(&conf.Port, "port", "p", 5432, hPort)
    pflag.StringVarP(&conf.Database, "database", "d", "postgres", hDatabase)
    pflag.StringVarP(&conf.User, "username", "U", getOSUser(), hUser)

    pflag.Parse()

    return conf    
}

