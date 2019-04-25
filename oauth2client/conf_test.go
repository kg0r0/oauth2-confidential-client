package oauth2client

import (
	"io/ioutil"
	"os"
	"testing"
)

const testJson = `{
		"client_config": {
			"client_id": "testclientid",
			"client_secret": "testclientsecret"
		}
	}`

func TestUnmarshalConfig(t *testing.T) {
	type testCase struct {
		Test     string
		Result   string
		Expected string
	}

	tmpDir, _ := ioutil.TempDir("", "test")
	defer os.RemoveAll(tmpDir)
	ioutil.WriteFile(tmpDir+"/test_config.json", []byte(testJson), 0644)
	testConfigFilePath := tmpDir + "/test_config.json"
	config, err := NewConfig(testConfigFilePath)
	if err != nil {
		t.Error(err)
	}

	testCases := []testCase{
		{
			Test:     "test client_id",
			Result:   config.ClientConfig.ClientID,
			Expected: "testclientid",
		},
		{
			Test:     "test client_secret",
			Result:   config.ClientConfig.ClientSecret,
			Expected: "testclientsecret",
		},
	}

	for _, test := range testCases {
		if test.Expected != test.Result {
			t.Errorf("test:%v, expected:%v, result:%v", test.Test, test.Expected, test.Result)
		}
	}

}
