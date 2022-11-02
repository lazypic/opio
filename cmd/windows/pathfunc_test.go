package main

import (
	"testing"
)

func Test_Lin2win(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{{
		in:   "/lustre2/Digitalidea_source/flib/ai/14",
		want: "\\\\10.0.200.100\\lustre_Digitalidea_source\\flib\\ai\\14",
	}, {
		in:   "/lustre/Digitalidea_source/flib/ai/14",
		want: "\\\\10.0.200.100\\lustre_Digitalidea_source\\flib\\ai\\14",
	}, {
		in:   "/show/ghost/seq",
		want: "\\\\10.0.200.100\\show_ghost\\seq",
	}, {
		in:   "/lustre/show/ghost/seq",
		want: "\\\\10.0.200.100\\show_ghost\\seq",
	}, {
		in:   "/lustre4/show/thesea2/seq",
		want: "\\\\10.0.200.100\\show_thesea2\\seq",
	}, {
		in:   "/lustre2/Marketing/2015Brochure/Creature/0911_confirm",
		want: "/lustre2/Marketing/2015Brochure/Creature/0911_confirm",
	}, {
		in:   "/MMHUB_nas01/show/test",
		want: "Z:\\show\\test",
	}}
	for _, c := range cases {
		got := Lin2win(c.in)
		if got != c.want {
			t.Fatalf("Win2lin(%v): 얻은 값 %v, 원하는 값 %v", c.in, got, c.want)
		}
	}
}

func Test_Win2lin(t *testing.T) {
	cases := []struct {
		in   string
		want string
	}{{
		in:   "\\\\10.0.200.100\\lustre_Digitalidea_source\\flib\\ai\\14",
		want: "/lustre2/Digitalidea_source/flib/ai/14",
	}, {
		in:   "\\\\10.0.200.100\\show_ghost\\seq",
		want: "/show/ghost/seq",
	}, {
		in:   "/lustre2/Marketing/2015Brochure/Creature/0911_confirm",
		want: "/lustre2/Marketing/2015Brochure/Creature/0911_confirm",
	}, {
		in:   "/lustre3/show/TEMP/tmp",
		want: "/lustre3/show/TEMP/tmp",
	}}
	for _, c := range cases {
		got := Win2lin(c.in)
		if got != c.want {
			t.Fatalf("Win2lin(%v): 얻은 값 %v, 원하는 값 %v", c.in, got, c.want)
		}
	}
}
