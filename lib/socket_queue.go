package socket

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"reflect"
	"strconv"
	"strings"
	"unsafe"

	"github.com/jessevdk/go-flags"
	mp "github.com/mackerelio/go-mackerel-plugin"
)

type options struct {
	IP     string `short:"i" long:"ip" description:"process name" required:"true"`
	PORT   uint16 `short:"p" long:"port" description:"release directory" required:"true"`
	Prefix string `long:"prefix" description:"process port name" required:"false"`
}

var opts options

type SocketQueue struct {
	Prefix      string
	ip          net.IP
	port        uint16
	state       uint8
	listenQueue uint64
	acceptQueue uint64
}

/*
enum {
	TCP_ESTABLISHED = 1,
	TCP_SYN_SENT,
	TCP_SYN_RECV,
	TCP_FIN_WAIT1,
	TCP_FIN_WAIT2,
	TCP_TIME_WAIT,
	TCP_CLOSE,
	TCP_CLOSE_WAIT,
	TCP_LAST_ACK,
	TCP_LISTEN,
	TCP_CLOSING,
	TCP_NEW_SYN_RECV,
	TCP_MAX_STATE
};
*/
func New() (*SocketQueue, error) {
	addr := net.ParseIP(opts.IP)
	if addr == nil {
		return &SocketQueue{}, fmt.Errorf("parse ip error")
	}
	if opts.PORT < 1 || opts.PORT > 65535 {
		return &SocketQueue{}, fmt.Errorf("port error")
	}

	var prefix string

	if opts.Prefix == "" {
		prefix = ""
	} else {
		prefix = fmt.Sprintf("%s-", opts.Prefix)
	}

	return &SocketQueue{
		ip:     addr,
		port:   opts.PORT,
		state:  10, // TCP_LISTEN
		Prefix: prefix,
	}, nil

}

// エンディアンの確認を実行する
func checkEndian() string {
	var i int = 0x0100
	ptr := unsafe.Pointer(&i)
	if 0x01 == *(*byte)(ptr) {
		return "big"
	} else if 0x00 == *(*byte)(ptr) {
		return "little"
	}

	return "other"
}

/*
   46: 010310AC:9C4C 030310AC:1770 01
   |      |      |      |      |   |--> connection state
   |      |      |      |      |------> remote TCP port number
   |      |      |      |-------------> remote IPv4 address
   |      |      |--------------------> local TCP port number
   |      |---------------------------> local IPv4 address
   |----------------------------------> number of entry

   00000150:00000000 01:00000019 00000000
      |        |     |     |       |--> number of unrecovered RTO timeouts
      |        |     |     |----------> number of jiffies until timer expires
      |        |     |----------------> timer_active (see below)
      |        |----------------------> receive-queue
      |-------------------------------> transmit-queue
*/
func (sq *SocketQueue) ParseTcpFile(path string) error {
	b, err := os.Open(path)
	if err != nil {
		return err
	}
	s := bufio.NewScanner(b)

	for s.Scan() {
		line := s.Text()
		var l []string
		for _, v := range strings.Split(line, " ") {
			if len(v) != 0 {
				l = append(l, v)
			}
		}

		// ステータスがListenのときのみ
		if l[3] == "0A" {
			var sl, d []int
			for _, v := range strings.Split(fmt.Sprintf("%s", sq.ip), ".") {
				s, err := strconv.Atoi(v)
				if err != nil {
					return fmt.Errorf("failed atoi: %s", v)
				}
				sl = append(sl, s)
			}
			dt := strings.Split(l[1], ":")[0]

			// リトルエンディアンの場合のみ対応
			if checkEndian() == "little" {
				for _, v := range []string{dt[6:8], dt[4:6], dt[2:4], dt[0:2]} {
					s, err := strconv.ParseInt(v, 16, 32)
					if err != nil {
						return fmt.Errorf("failed parseint: %s", v)
					}
					d = append(d, int(s))

				}

				// オプションで指定されたIPとprocfsのIPが一致しているかどうかを確認
				if reflect.DeepEqual(sl, d) {
					p, err := strconv.ParseUint(strings.Split(l[1], ":")[1], 16, 16)
					if err != nil {
						return fmt.Errorf("failed paserUint: %s", l[1])
					}
					if uint16(p) == sq.port {
						sq.listenQueue, err = strconv.ParseUint(strings.Split(l[4], ":")[0], 16, 64)
						if err != nil {
							return fmt.Errorf("failed paserUint: %s", l[4])
						}
						sq.acceptQueue, err = strconv.ParseUint(strings.Split(l[4], ":")[1], 16, 64)
						if err != nil {
							return fmt.Errorf("failed paserUint: %s", l[4])
						}
					}
				}
			}
		}
	}

	return nil
}

func (sq SocketQueue) FetchMetrics() (map[string]float64, error) {
	m := make(map[string]float64)

	m["syn-queue"] = float64(sq.listenQueue)
	m["accept-queue"] = float64(sq.acceptQueue)

	return m, nil
}

func graphGen(labelPrefix string) map[string]mp.Graphs {
	return map[string]mp.Graphs{
		fmt.Sprintf("%s-socket-queue", labelPrefix): {
			Label: labelPrefix,
			Unit:  mp.UnitInteger,
			Metrics: []mp.Metrics{
				{Name: "syn-queue", Label: "syn-queue", Diff: false},
				{Name: "accept-queue", Label: "accept-queue", Diff: false},
			},
		},
	}
}

func (sq SocketQueue) GraphDefinition() map[string]mp.Graphs {
	return graphGen(opts.Prefix)
}

func Run() error {
	sq, err := New()
	if err != nil {
		return err
	}
	err = sq.ParseTcpFile("/proc/net/tcp")
	if err != nil {
		return err
	}

	plugin := mp.NewMackerelPlugin(sq)
	plugin.Run()

	return nil
}

func parseArgs(args []string) error {
	_, err := flags.ParseArgs(&opts, os.Args)

	if opts.IP == "localhost" {
		opts.IP = "127.0.0.1"
	}

	if err != nil {
		return err
	}

	return nil
}

func Do() {
	if parseArgs(os.Args[1:]) != nil {
		os.Exit(1)
	}

	if Run() != nil {
		os.Exit(1)
	}
}
