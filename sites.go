package sites

// ParseSites - parse sites to find proxy.
// logDebug to show debug information
// logError to show error messages
func ParseSites(logDebug, logError bool) []string {
	type parser struct {
		name string
		ips  []string
	}
	var ips []string
	useDebug = logDebug
	useError = logError
	debugmsg("start parse sites")
	ch := make(chan parser)
	go func() {
		data := parser{name: "cnProxyCom", ips: cnProxyCom()}
		ch <- data
	}()
	go func() {
		data := parser{name: "coolProxyNet", ips: coolProxyNet()}
		ch <- data
	}()
	go func() {
		data := parser{name: "cyberSyndromeNet", ips: cyberSyndromeNet()}
		ch <- data
	}()
	go func() {
		data := parser{name: "fakeMyIPInfo", ips: fakeMyIPInfo()}
		ch <- data
	}()
	go func() {
		data := parser{name: "freeProxyList", ips: freeProxyList()}
		ch <- data
	}()
	go func() {
		data := parser{name: "freeProxyListCom", ips: freeProxyListCom()}
		ch <- data
	}()
	go func() {
		data := parser{name: "gatherProxyCom", ips: gatherProxyCom()}
		ch <- data
	}()
	go func() {
		data := parser{name: "idcloakCom", ips: idcloakCom()}
		ch <- data
	}()
	go func() {
		data := parser{name: "ipaddressCom", ips: ipaddressCom()}
		ch <- data
	}()
	go func() {
		data := parser{name: "kuaidaili", ips: kuaidaili()}
		ch <- data
	}()
	go func() {
		data := parser{name: "mrHinkyDinkCom", ips: mrHinkyDinkCom()}
		ch <- data
	}()
	go func() {
		data := parser{name: "nnTimeCom", ips: nnTimeCom()}
		ch <- data
	}()
	go func() {
		data := parser{name: "proxyListOrg", ips: proxyListOrg()}
		ch <- data
	}()
	go func() {
		data := parser{name: "proxyServerList24Top", ips: proxyServerList24Top()}
		ch <- data
	}()
	go func() {
		data := parser{name: "rawList", ips: rawList()}
		ch <- data
	}()
	go func() {
		data := parser{name: "webanetLabs", ips: webanetLabs()}
		ch <- data
	}()
	go func() {
		data := parser{name: "xicidailiCom", ips: xicidailiCom()}
		ch <- data
	}()

	for i := 0; i < 17; i++ {
		data := <-ch
		debugmsg("get", len(data.ips), "from", data.name)
		ips = append(ips, fixURI(data.ips)...)
	}
	debugmsg("end parse sites, found", len(ips), "proxy")
	return ips
}
