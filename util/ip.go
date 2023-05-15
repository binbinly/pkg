package util

import (
	"net"
	"net/http"
	"strings"
	"time"
)

// InternalIPv4 获取内网地址 (IPv4)
func InternalIPv4() string {
	return InternalIP("", "udp4")
}

// InternalIPv6 获取内网地址 (临时 IPv6 地址)
func InternalIPv6() string {
	return InternalIP("[2001:4860:4860::8888]:53", "udp6")
}

// InternalIP 获取内网地址 (出口本地地址)
func InternalIP(dstAddr, network string) string {
	if dstAddr == "" {
		dstAddr = "8.8.8.8:53"
	}
	if network == "" {
		network = "udp"
	}

	conn, err := net.DialTimeout(network, dstAddr, time.Second)
	if err != nil {
		return ""
	}

	defer func() {
		_ = conn.Close()
	}()

	addr := conn.LocalAddr().String()
	ip := net.ParseIP(addr).String()
	if ip == "<nil>" {
		ip, _, _ = net.SplitHostPort(addr)
	}

	return ip
}

// LocalIP 获取本地地址 (第一个)
func LocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok &&
				!ipnet.IP.IsLinkLocalUnicast() && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}

	return ""
}

// LocalIPv4s 获取所有本地地址 IPv4
func LocalIPv4s() (ips []string) {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok &&
				!ipnet.IP.IsLinkLocalUnicast() && !ipnet.IP.IsLoopback() && ipnet.IP.To4() != nil {
				ips = append(ips, ipnet.IP.String())
			}
		}
	}

	return
}

// InterfaceAddrs 获取所有带 IP 的接口和对应的所有 IP
// 排除本地链路地址和环回地址
func InterfaceAddrs(v ...string) (map[string][]net.IP, error) {
	ifAddrs := make(map[string][]net.IP)
	ifaces, err := net.Interfaces()
	if err != nil {
		return ifAddrs, err
	}

	var (
		ip net.IP
		t  string
	)
	if len(v) > 0 {
		t = strings.ToLower(v[0])
	}

	for _, i := range ifaces {
		addrs, err := i.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			default:
				ip = net.IPv4zero
			}
			if ip.IsLinkLocalUnicast() || ip.IsLoopback() {
				continue
			}
			switch t {
			case "ipv6":
				if ip.To4() != nil {
					continue
				}
			case "ipv4":
				if ip.To4() == nil {
					continue
				}
			}
			ifAddrs[i.Name] = append(ifAddrs[i.Name], ip)
		}
	}
	return ifAddrs, nil
}

// GetRealIP get user real ip
func GetRealIP(r *http.Request) (ip string) {
	var header = r.Header
	var index int
	if ip = header.Get("X-Forwarded-For"); ip != "" {
		index = strings.IndexByte(ip, ',')
		if index < 0 {
			return ip
		}
		if ip = ip[:index]; ip != "" {
			return ip
		}
	}
	if ip = header.Get("X-Real-Ip"); ip != "" {
		index = strings.IndexByte(ip, ',')
		if index < 0 {
			return ip
		}
		if ip = ip[:index]; ip != "" {
			return ip
		}
	}
	if ip = header.Get("Proxy-Forwarded-For"); ip != "" {
		index = strings.IndexByte(ip, ',')
		if index < 0 {
			return ip
		}
		if ip = ip[:index]; ip != "" {
			return ip
		}
	}
	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		return ""
	}

	return ip
}
