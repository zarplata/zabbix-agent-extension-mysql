package main

import (
	"fmt"
	"os"
	"strconv"

	zsend "github.com/blacked/go-zabbix"
	docopt "github.com/docopt/docopt-go"
	hierr "github.com/reconquest/hierr-go"
)

var (
	stats        []map[string]string
	version      = "[manual build]"
	queryGlobal  = "SHOW GLOBAL STATUS"
	queryGalera  = "SHOW GLOBAL STATUS LIKE 'wsrep_%'"
	queryProcess = "SHOW PROCESSLIST"
	querySlave   = "SHOW SLAVE STATUS"
)

func main() {
	usage := `zabbix-agent-extension-mysql

Usage:
  zabbix-agent-extension-mysql [options] [<zabbix options>] [<mysql options>]

Options:
  --discovery                 Run low-level discovery for detect galera/slave.
  --type <type>               Type of statistic global/process/slave/galera
                                send to zabbix [default: global].

Zabbix options:
  -z --zabbix <zabbix>        Hostname or IP address of zabbix server
                                [default: 127.0.0.1].
  -p --port <port>            Port of zabbix server [default: 10051].
  --zabbix-prefix <prefix>    Add part of your prefix for key [default: None].

MySQL options:
  -n --network <net>          Network type unix or tcp [default: tcp].
  -m --mysql <dsn>            MySQL DSN socket or host:port
                                [default: localhost:3306].
  --user <user>               MySQL user [default: zabbix].
  --pass <password>           MySQL password [default: zabbix].

Misc options:
  --version                   Show version.
  -h --help                   Show this screen.
`
	args, err := docopt.Parse(usage, nil, true, version, false)
	if err != nil {
		fmt.Println(hierr.Errorf(err, "can't parse docopt"))
		os.Exit(1)
	}

	zabbix := args["--zabbix"].(string)
	port, err := strconv.Atoi(args["--port"].(string))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	zabbixPrefix := args["--zabbix-prefix"].(string)
	if zabbixPrefix == "None" {
		zabbixPrefix = "mysql"
	} else {
		zabbixPrefix = fmt.Sprintf("%s.%s", zabbixPrefix, "mysql")
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	network := args["--network"].(string)

	mysqlDSN := args["--mysql"].(string)

	user := args["--user"].(string)
	password := args["--pass"].(string)

	statsType := args["--type"].(string)

	dsn := fmt.Sprintf("%s:%s@%s(%s)/", user, password, network, mysqlDSN)

	if args["--discovery"].(bool) {
		err = discovery(dsn)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		os.Exit(0)
	}

	var (
		metrics       []*zsend.Metric
		filterMetrics []string
	)

	switch statsType {
	case "galera":
		stats, err = getGlobalStats(queryGalera, dsn)
		filterMetrics = galeraMetrics
	case "process":
		stats, err = getStats(queryProcess, dsn)
		stats = calcProcessStats(stats)
		filterMetrics = processMetrics
	case "slave":
		stats, err = getStats(querySlave, dsn)
		filterMetrics = slaveMetrics
	default:
		stats, err = getGlobalStats(queryGlobal, dsn)
		filterMetrics = globalMetrics
	}
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	metrics = createMetrics(
		hostname,
		stats[0],
		statsType,
		filterMetrics,
		metrics,
		zabbixPrefix,
	)

	packet := zsend.NewPacket(metrics)
	sender := zsend.NewSender(
		zabbix,
		port,
	)
	sender.Send(packet)
	fmt.Println("OK")
}
