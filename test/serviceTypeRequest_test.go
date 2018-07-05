package test

import (
	"testing"

	"github.com/nbio/st"
)

func TestServiceTypeRequest_ok(t *testing.T) {
	testRequest("serviceTypeRequest_answOk.txt")
	data, err := zApi.ServiceTypeRequest()

	st.Expect(t, err, nil)
	st.Expect(t, data[0].ServiceCode, "1")
}
