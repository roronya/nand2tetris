package symboletable

import "testing"

func TestNew(t *testing.T) {
	commands := []string{"@Xxx", "(label1)", "(label2)"}
	st := New(commands)
	st.AddEntry("TEST", 0)
	address := st.GetAddress("TEST")
	if address != 0 {
		t.Fatalf("fail")
	}
	address = st.GetAddress("KBD")
	if address != 24567 {
		t.Fatalf("address is expected %d. got=%d", 24567, address)
	}
	address = st.GetAddress("R1")
	if address != 1 {
		t.Fatalf("address is expected %d. got=%d", 1, address)
	}
	address = st.GetAddress("R7")
	if address != 7 {
		t.Fatalf("address is expected %d. got=%d", 7, address)
	}
	address = st.GetAddress("label1")
	if address != 1 {
		t.Fatalf("address is expected %d. got=%d", 1, address)
	}
	address = st.GetAddress("label2")
	if address != 2 {
		t.Fatalf("address is expected %d. got=%d", 2, address)
	}
}
