package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os/user"
	"path"
	"sync"
	"time"

	"github.com/davecgh/go-spew/spew"
	"github.com/saitofun/qkit/base/cmd"
	"github.com/saitofun/qkit/base/types"
	"github.com/saitofun/qkit/conf/app"
	"github.com/saitofun/qkit/conf/deploy/supervisor"
	"github.com/saitofun/qkit/conf/section_config"
	"github.com/spf13/cobra"
)

// PgConf postgresql config
type PgConf struct {
	Endpoint types.Endpoint `env:""`
	*sql.DB  `env:"-"`
}

func (c *PgConf) SetDefault() {
	if c.Endpoint.Scheme == "" {
		c.Endpoint.Scheme = "postgres"
	}
	if c.Endpoint.Hostname == "" {
		c.Endpoint.Hostname = "127.0.0.1"
	}
	if c.Endpoint.Port == 0 {
		c.Endpoint.Port = 5432
	}
	if c.Endpoint.Username == "" {
		c.Endpoint.Username = "postgres_username"
	}
	if c.Endpoint.Base == "" {
		c.Endpoint.Base = "postgres"
	}
}

// RedisConf redis config
type RedisConf struct {
	Endpoint types.Endpoint `env:""`
}

// ServerConf http server config
type ServerConf struct {
	Port int
}

func (s *ServerConf) Addr() string { return fmt.Sprintf(":%d", s.Port) }

// DemoConf app config
type DemoConf struct {
	SomeBusinessRangeMin int64
	SomeBusinessRangeMax int64
	SomeCronTaskInterval types.Second
	SomeDefaultKey       string
	SomeAssertCertFile   types.Address
}

var (
	Database = &PgConf{}
	Server   = &ServerConf{}
	Cache    = &RedisConf{}
	Config   = &DemoConf{}
	program  = &supervisor.Program{
		Section:               section_config.Section{Name: "program", Value: "demo"},
		AutoStart:             true,
		AutoRestart:           true,
		Command:               cmd.Command{Name: "demo"},
		User:                  "demo",
		Priority:              500,
		ExitCodes:             types.CommaSplitInts{0},
		StopSignal:            types.SIGQUIT,
		Envs:                  types.CommaSplitStrings{},
		StdoutLogFile:         "/a/b/c/demo.log",
		StdoutLogFileMaxBytes: 5,
		StdoutLogFileBackups:  3,
	}
)

func init() {
	app := app.New(
		app.WithName("srv-demo"),
		app.WithRoot("."),
		app.WithVersion("0.0.1"),
		app.WithDeployer(program),
	)
	app.Conf(Server, Database, Cache, Config)

	app.AddCommand("version", func(...string) {
		fmt.Println(app.String())
	}, func(cmd *cobra.Command) {
		cmd.Short = "print app version"
	})

	app.AddCommand("gen", func(...string) {
		if err := app.MarshalDefault(); err != nil {
			panic(err)
		}
	}, func(cmd *cobra.Command) {
		cmd.Short = "generate app default config"
	})

	app.AddCommand("supervisor", func(...string) {
		program.SetDefault()
		filename := path.Join(app.Root(), "supervisor.conf")
		if err := program.Write(filename); err != nil {
			panic(fmt.Errorf("init `supervisor.conf` error: %v", err))
		}
		data, _ := program.Bytes()
		fmt.Println(string(data))
	}, func(cmd *cobra.Command) {
		cmd.Short = "generate app supervisor config"
	})

	app.AddCommand("docker", func(...string) {
		fmt.Println("todo: generate app docker config")
	}, func(cmd *cobra.Command) {
		cmd.Short = "generate app docker config"
	})

	app.AddCommand("config", func(...string) {
		spew.Dump(Server)
		spew.Dump(Database)
		spew.Dump(Cache)
		spew.Dump(Config)
	}, func(cmd *cobra.Command) {
		cmd.Short = "dump app config value"
	})

	app.AddCommand("user", func(...string) {
		u, _ := user.Current()
		if u.Username != "root" {
			panic("pls run this with root")
		}
		if u, err := user.Lookup(program.User); err != nil {
			fmt.Printf("%s user is not exist\n", program.User)
		} else {
			fmt.Printf("%s:%s:%v created\n", u.Name, u.Username, u.Gid)
		}
		fmt.Println("todo: create system user for app permission")
	}, func(cmd *cobra.Command) {
		cmd.Short = "create system user for app permission"
	})

	app.Execute(func(...string) {
		BatchRun(
			StartServer,
			PushLog,
			Proxy,
		)
	})
}

func StartServer() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = w.Write([]byte("Hello Demo"))
	})
	log.Printf("StartServer: http server started at %s", Server.Addr())
	err := http.ListenAndServe(Server.Addr(), nil)
	if err != nil {
		log.Fatalf("server: %v", err)
	}
}

func Proxy() {
	log.Printf("Proxy: streaming proxy started")
}

func PushLog() {
	log.Printf("PushLog: log pushing started")
}

func BatchRun(tasks ...func()) {
	wg := &sync.WaitGroup{}

	for i := range tasks {
		fn := tasks[i]
		wg.Add(1)

		go func() {
			defer wg.Done()
			fn()
			time.Sleep(200 * time.Millisecond)
		}()
	}
	wg.Wait()
}

func main() {}
