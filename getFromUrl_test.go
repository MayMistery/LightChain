package main

import "testing"

func Test_getFromUrl(t *testing.T) {
	type args struct {
		strURL   string
		filename string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			"test1",
			args{
				"https://www.ucg.ac.me/skladiste/blog_44233/objava_64433/fajlovi/Computer%20Networking%20_%20A%20Top%20Down%20Approach,%207th,%20converted.pdf",
				"", // ""表示采用url中名字
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getFromUrl(tt.args.strURL, tt.args.filename)
		})
	}
}
