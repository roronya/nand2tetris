package symboletable

import "testing"

func TestNew(t *testing.T) {
	st := New()
	st.addEntry("TEST", 0)
	address := st.getAddress("TEST")
	if address != 0 {
		t.Fatalf("fail")
	}
	address = st.getAddress("KBD")
	if address != 24567 {
		t.Fatalf("address is expected %d. got=%d", 24567, address)
	}
	address = st.getAddress("R7")
	if address != 7 {
		t.Fatalf("address is expected %d. got=%d", 7, address)
	}
}
