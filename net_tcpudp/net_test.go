package main

func TestNETServer_Request(t *testing.T) {
	tt := []struct{
		test string
		payload []byte
		want []byte
	}{
		{
			"Sending a simple request returns result",
			[]byte("hello\n")
			[]byte("Request received: hello")
		},
		{
			"Sending another request",
			[]byte("hello2\n")
			[]byte("Request received: hello2")
		},
	}

	for _, tc := range tt {
		t.Run(tc.test, func(t *testing.T) {
			conn, err := net.Dial("tcp", ":1234")
			if err != nil {
				t.Error("could not connect to a tcp server", err)
			}
			defer conn.Close()

			if _, err := conn.Write(tc.payload); err != nil {
				t.Error("could not write payload to a TCP server", err)
			}

			out := make([]byte, 1024)
			if _, err := conn.Read(out); err == nil {
				if bytes.Compare(out, tc.want) == 0 {
					t.Error("response did not match expected ouput")
				}
			} else {
				t.Error()
			}
		})
	}
}