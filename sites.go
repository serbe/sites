package sites

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
		data := parser{name: "freeproxylist", ips: freeproxylist()}
		ch <- data
	}()
	go func() {
		data := parser{name: "freeproxylistcom", ips: freeproxylistcom()}
		ch <- data
	}()
	go func() {
		data := parser{name: "gatherproxycom", ips: gatherproxycom()}
		ch <- data
	}()
	go func() {
		data := parser{name: "kuaidaili", ips: kuaidaili()}
		ch <- data
	}()
	go func() {
		data := parser{name: "proxylistorg", ips: proxylistorg()}
		ch <- data
	}()
	go func() {
		data := parser{name: "proxyserverlist24top", ips: proxyserverlist24top()}
		ch <- data
	}()
	go func() {
		data := parser{name: "rawlist", ips: rawlist()}
		ch <- data
	}()
	go func() {
		data := parser{name: "webanetlabs", ips: webanetlabs()}
		ch <- data
	}()
	for i := 0; i < 8; i++ {
		data := <-ch
		ips = append(ips, data.ips...)
	}
	debugmsg("end parse sites, found", len(ips), "proxy")
	return ips
}
