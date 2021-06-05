package config

import (
	"errors"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	tests := []struct {
		name   string
		setEnv func()
		want   Config
		err    error
	}{
		{
			name: "should return a config when set env values",
			setEnv: func() {
				setEnv("KAFKA_CONNECT_URL", "kafka:9092")
				setEnv("CONNECTORS_FILES_DIR", "/tmp/")
				setEnv("MAX_RETRY", "10")
				setEnv("WAIT_START_TIME", "10s")
			},
			want: Config{
				ConnectUrl:    "kafka:9092",
				ConnectorsDir: "/tmp/",
				MaxRetry:      10,
				WaitStartTime: 10 * time.Second,
			},
			err: nil,
		},
		{
			name: "should not return error when files_dir env is filled",
			setEnv: func() {
				setEnv("CONNECTORS_FILES_DIR", "/tmp/")
			},
			want: Config{
				ConnectUrl:    DefaultUrl,
				ConnectorsDir: "/tmp/",
				MaxRetry:      DefaultMaxRetry,
				WaitStartTime: DefaultWaitStartTime,
			},
			err: nil,
		},
		{
			name:   "should return error when files_dir is empty",
			setEnv: func() {},
			want:   Config{},
			err:    errors.New("CONNECTORS_FILES_DIR empty"),
		},
		{
			name: "should return error when max retry is invalid",
			setEnv: func() {
				setEnv("CONNECTORS_FILES_DIR", "/tmp/")
				setEnv("MAX_RETRY", "one")
			},
			want: Config{},
			err:  errors.New("strconv.Atoi: parsing \"one\": invalid syntax"),
		},
		{
			name: "should return error when wait start time is invalid",
			setEnv: func() {
				setEnv("CONNECTORS_FILES_DIR", "/tmp/")
				setEnv("WAIT_START_TIME", "midnight")
			},
			want: Config{},
			err:  errors.New("time: invalid duration \"midnight\""),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Clearenv()
			tt.setEnv()

			got, err := New()
			if tt.err == nil {
				if err != nil {
					t.Errorf("New() error = %v, wantErr %v", err, tt.err)
				}
			} else {
				if err == nil || err.Error() != tt.err.Error() {
					t.Errorf("New() error = %v, wantErr %v", err, tt.err)
				}
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func setEnv(key, value string) {
	if err := os.Setenv(key, value); err != nil {
		panic(err)
	}
}
