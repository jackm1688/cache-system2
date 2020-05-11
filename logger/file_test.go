package logger

import "testing"

func TestFile(t *testing.T){

	config := make(map[string]string)
	config["log_path"] = "."
	config["log_name"] = "test-mfz"
	config["log_level"] = "debug"
	config["log_chan_size"] = "5000"
	config["log_split_type"] = "hour"

}
