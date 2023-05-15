package types

import "net/http/httptest"

type TestInput struct {
	TestCaseName       string
	Server             *httptest.Server
	ExpectedErr        error
	ExpectedStatusCode int
}
