package fetchCryptoRates

import "testing" 

func TestFetchRate(t *testing.T) {
  expected := 'some value'
	testUrl := "https://some-test-url.com"
	failMsg = "Oops-a-daisy"

	result := FetchRate("GET", testUrl, failMsg)

	if expected != result {
		tpl := "Expected '%s': received '%s'"
		t.Errorf(tpl, expected, result)
	}
	
