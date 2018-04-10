# zabbix-agent-extension-mysql

zabbix-agent-extension-mysql - this extension for monitoring mysql.

### Supported features

This extension obtains stats of three types:

#### Global stats

- [x] Aborted clients 
- [x] Aborted connects
- [x] Bytes received
- [x] Bytes sent
- [x] Com begin
- [x] Com commit
- [x] Com delete
- [x] Com insert
- [x] Com rollback
- [x] Com select
- [x] Com update
- [x] Innodb rows deleted
- [x] Innodb rows inserted
- [x] Innodb rows read
- [x] Innodb rows updated
- [x] Queries
- [x] Slow queries
- [x] Threads running
- [x] Uptime

#### Process stats

- [x] Processlist count
- [x] Max query time

#### Galera stats

- [x] Galera cluster size
- [x] Galera cluster state uuid
- [x] Galera cluster status
- [x] Galera connected
- [x] Galera internal state EVS Protocol
- [x] Galera cluster gcom uuid
- [x] Galera local state
- [x] Galera local state comment
- [x] Galera local state uuid
- [x] Galera protocol version
- [x] Galera provider name
- [x] Galera ready

#### Slave stats

- [x] Mysql io running
- [x] Mysql sql running
- [x] Mysql Seconds behind master
- [x] Mysql master host
- [x] Mysql master port

### Installation

#### Manual build

```sh
# Building
git clone https://github.com/zarplata/zabbix-agent-extension-mysql.git
cd zabbix-agent-extension-mysql
make

#Installing
make install

# By default, binary installs into /usr/bin/ and zabbix config in /etc/zabbix/zabbix_agentd.conf.d/ but,
# you may manually copy binary to your executable path and zabbix config to specific include directory
```

#### Arch Linux package
```sh
# Building
git clone https://github.com/zarplata/zabbix-agent-extension-mysql.git
git checkout pkgbuild

makepkg

#Installing
pacman -U *.tar.xz
```

### Requirements

Add mysql users:
```sh
GRANT REPLICATION CLIENT,PROCESS ON *.* TO 'zabbix'@'127.0.0.1'  IDENTIFIED BY 'zabbix';
```

### Dependencies

zabbix-agent-extension-elasticsearch requires [zabbix-agent](http://www.zabbix.com/download) v3.4+ to run.

### Zabbix configuration
In order to start getting metrics, it is enough to import template and attach it to monitored node.

`WARNING:` You must define macro with name - `{$ZABBIX_SERVER_IP}` in global or local (template) scope with IP address of  zabbix server.
