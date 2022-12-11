package socket

import (
	"net"
	"reflect"
	"testing"

	mp "github.com/mackerelio/go-mackerel-plugin"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name    string
		want    *SocketQueue
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New()
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_checkEndian(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkEndian(); got != tt.want {
				t.Errorf("checkEndian() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSocketQueue_ParseTcpFile(t *testing.T) {
	type fields struct {
		Prefix      string
		ip          net.IP
		port        uint16
		state       uint8
		listenQueue uint64
		acceptQueue uint64
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "test001 pass",
			fields: fields{
				Prefix:      "",
				ip:          net.ParseIP("0.0.0.0"),
				port:        9000,
				state:       10,
				listenQueue: 0,
				acceptQueue: 0,
			},
			args: args{
				path: "./test/proc/001.txt",
			},
			wantErr: false,
		},
		{
			name: "test002 pass",
			fields: fields{
				Prefix:      "",
				ip:          net.ParseIP("192.168.1.182"),
				port:        9000,
				state:       10,
				listenQueue: 0,
				acceptQueue: 1,
			},
			args: args{
				path: "./test/proc/001.txt",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sq := &SocketQueue{
				Prefix:      tt.fields.Prefix,
				ip:          tt.fields.ip,
				port:        tt.fields.port,
				state:       tt.fields.state,
				listenQueue: tt.fields.listenQueue,
				acceptQueue: tt.fields.acceptQueue,
			}
			if err := sq.ParseTcpFile(tt.args.path); (err != nil) != tt.wantErr {
				t.Errorf("SocketQueue.ParseTcpFile() error = %v, wantErr %v", err, tt.wantErr)
			}
			if sq.acceptQueue != tt.fields.acceptQueue {
				t.Errorf("parse error %v: %v", sq.acceptQueue, tt.fields.acceptQueue)
			}
			if sq.listenQueue != tt.fields.listenQueue {
				t.Errorf("parse error %v: %v", sq.listenQueue, tt.fields.listenQueue)
			}
		})
	}
}

func TestSocketQueue_FetchMetrics(t *testing.T) {
	type fields struct {
		Prefix      string
		ip          net.IP
		port        uint16
		state       uint8
		listenQueue uint64
		acceptQueue uint64
	}
	tests := []struct {
		name    string
		fields  fields
		want    map[string]float64
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sq := SocketQueue{
				Prefix:      tt.fields.Prefix,
				ip:          tt.fields.ip,
				port:        tt.fields.port,
				state:       tt.fields.state,
				listenQueue: tt.fields.listenQueue,
				acceptQueue: tt.fields.acceptQueue,
			}
			got, err := sq.FetchMetrics()
			if (err != nil) != tt.wantErr {
				t.Errorf("SocketQueue.FetchMetrics() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SocketQueue.FetchMetrics() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_graphGen(t *testing.T) {
	type args struct {
		labelPrefix string
	}
	tests := []struct {
		name string
		args args
		want map[string]mp.Graphs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := graphGen(tt.args.labelPrefix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("graphGen() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSocketQueue_GraphDefinition(t *testing.T) {
	type fields struct {
		Prefix      string
		ip          net.IP
		port        uint16
		state       uint8
		listenQueue uint64
		acceptQueue uint64
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]mp.Graphs
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sq := SocketQueue{
				Prefix:      tt.fields.Prefix,
				ip:          tt.fields.ip,
				port:        tt.fields.port,
				state:       tt.fields.state,
				listenQueue: tt.fields.listenQueue,
				acceptQueue: tt.fields.acceptQueue,
			}
			if got := sq.GraphDefinition(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SocketQueue.GraphDefinition() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRun(t *testing.T) {
	tests := []struct {
		name    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Run(); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_parseArgs(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := parseArgs(tt.args.args); (err != nil) != tt.wantErr {
				t.Errorf("parseArgs() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestDo(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Do()
		})
	}
}
