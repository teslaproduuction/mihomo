package convert

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// https://v2.hysteria.network/zh/docs/developers/URI-Scheme/
func TestConvertsV2Ray_normal(t *testing.T) {
	hy2test := "hysteria2://letmein@example.com:8443/?insecure=1&obfs=salamander&obfs-password=gawrgura&pinSHA256=deadbeef&sni=real.example.com&up=114&down=514&alpn=h3,h4#hy2test"

	expected := []map[string]interface{}{
		{
			"name":             "hy2test",
			"type":             "hysteria2",
			"server":           "example.com",
			"port":             "8443",
			"sni":              "real.example.com",
			"obfs":             "salamander",
			"obfs-password":    "gawrgura",
			"alpn":             []string{"h3", "h4"},
			"password":         "letmein",
			"up":               "114",
			"down":             "514",
			"skip-cert-verify": true,
			"fingerprint":      "deadbeef",
		},
	}

	proxies, err := ConvertsV2Ray([]byte(hy2test))

	assert.Nil(t, err)
	assert.Equal(t, expected, proxies)
}

func TestConvertsV2Ray_mieru(t *testing.T) {
	tcpLink := "mieru://2ff704c4-4c3b-439d-90cd-071218edec22:h@212.118.56.186/?handshake-mode=HANDSHAKE_NO_WAIT&mtu=1400&multiplexing=MULTIPLEXING_HIGH&port=15695-15698&protocol=TCP#212.118.56.186%20MieruTCP"
	udpLink := "mieru://2ff704c4-4c3b-439d-90cd-071218edec22:h@212.118.56.186/?handshake-mode=HANDSHAKE_NO_WAIT&mtu=1400&multiplexing=MULTIPLEXING_HIGH&port=42507-42510&protocol=UDP#212.118.56.186%20MieruUDP"

	expected := []map[string]interface{}{
		{
			"name":           "212.118.56.186 MieruTCP",
			"type":           "mieru",
			"server":         "212.118.56.186",
			"username":       "2ff704c4-4c3b-439d-90cd-071218edec22",
			"password":       "h",
			"transport":      "TCP",
			"port-range":     "15695-15698",
			"multiplexing":   "MULTIPLEXING_HIGH",
			"handshake-mode": "HANDSHAKE_NO_WAIT",
		},
		{
			"name":           "212.118.56.186 MieruUDP",
			"type":           "mieru",
			"server":         "212.118.56.186",
			"username":       "2ff704c4-4c3b-439d-90cd-071218edec22",
			"password":       "h",
			"transport":      "UDP",
			"port-range":     "42507-42510",
			"multiplexing":   "MULTIPLEXING_HIGH",
			"handshake-mode": "HANDSHAKE_NO_WAIT",
		},
	}

	proxies, err := ConvertsV2Ray([]byte(tcpLink + "\n" + udpLink))

	assert.Nil(t, err)
	assert.Equal(t, expected, proxies)
}
