package oauth2client

import (
	"io/ioutil"
	"os"
	"testing"
)

const testJson = `{
		"client_config": {
			"client_id": "testclientid",
			"client_secret": "testclientsecret",
			"endpoint": {
				"auth_url": "https://provider.com/o/oauth2/auth",
				"token_url": "https://provider.com/o/oauth2/token"
			},
			"redirect_uri": "https://example.com/callback",
			"scopes": ["SCOPE1", "SCOPE2"]
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
		{
			Test:     "test endpoint",
			Result:   config.ClientConfig.ClientSecret,
			Expected: "testclientsecret",
		},
		{
			Test:     "test redirect_uri",
			Result:   config.ClientConfig.RedirectURL,
			Expected: "https://example.com/callback",
		},
		{
			Test:     "test scope1",
			Result:   config.ClientConfig.Scopes[0],
			Expected: "SCOPE1",
		},
		{
			Test:     "test scope2",
			Result:   config.ClientConfig.Scopes[1],
			Expected: "SCOPE2",
		},
		{
			Test:     "test auth_url",
			Result:   config.ClientConfig.Endpoint.AuthURL,
			Expected: "https://provider.com/o/oauth2/auth",
		},
		{
			Test:     "test token_url",
			Result:   config.ClientConfig.Endpoint.TokenURL,
			Expected: "https://provider.com/o/oauth2/token",
		},
	}

	for _, test := range testCases {
		if test.Expected != test.Result {
			t.Errorf("test:%v, expected:%v, result:%v", test.Test, test.Expected, test.Result)
		}
	}

}
