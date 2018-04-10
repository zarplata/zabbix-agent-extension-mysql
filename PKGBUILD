pkgname=zabbix-agent-extension-mysql
pkgver=20180216.4_3e034da
pkgrel=1
pkgdesc="Zabbix agent for MySQL stats."
arch=('any')
license=('GPL')
makedepends=('go')
depends=()
install='install.sh'
source=("git+http://github.com/zarplata/zabbix-agent-extension-mysql.git#branch=master")
md5sums=('SKIP')

pkgver() {
    cd "$srcdir/$pkgname"

    make ver
}
    
build() {
    cd "$srcdir/$pkgname"

    make
}

package() {
    cd "$srcdir/$pkgname"

    install -Dm 0755 .out/"${pkgname}" "${pkgdir}/usr/bin/${pkgname}"
    install -Dm 0644 "${pkgname}.conf" "${pkgdir}/etc/zabbix/zabbix_agentd.conf.d/${pkgname}.conf"
}
