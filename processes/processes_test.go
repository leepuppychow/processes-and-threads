package processes

// func TestCreateProcessList(t *testing.T) {
// 	stdout := `
// 		UID   PID  PPID   C STIME   TTY           TIME CMD               WQ
// 		0     1     0   0  5:01AM ??         0:28.66 /sbin/launchd       3
// 		0    35     1   0  5:01AM ??         0:04.83 /usr/libexec/Use    4
// 		0    36     1   0  5:01AM ??         0:00.74 /usr/sbin/syslog    -
// 	`
// 	proc1 := Process{
// 		ProcessId:   1,
// 		ThreadCount: 3,
// 		Children:    nil,
// 		Command:     "/sbin/launchd",
// 		Duration:    "0:28.66",
// 	}
// 	proc2 := Process{
// 		ProcessId:   35,
// 		ThreadCount: 4,
// 		Children:    nil,
// 		Command:     "/usr/libexec/Use",
// 		Duration:    "0:04.83",
// 	}
// 	proc3 := Process{
// 		ProcessId:   36,
// 		ThreadCount: 0,
// 		Children:    nil,
// 		Command:     "/usr/sbin/syslog",
// 		Duration:    "0:00.74",
// 	}
// 	expected := []Process{
// 		proc1,
// 		proc2,
// 		proc3,
// 	}
// 	actual := CreateProcessList(stdout)

// 	for i, _ := range actual {
// 		if expected[i].ProcessId != actual[i].ProcessId {
// 			t.Errorf("Create Process List failed, processId does not match")
// 		}
// 		if expected[i].ThreadCount != actual[i].ThreadCount {
// 			t.Errorf("Create Process List failed, threadCount does not match")
// 		}
// 		if expected[i].Command != actual[i].Command {
// 			t.Errorf("Create Process List failed, Command does not match")
// 		}
// 		if expected[i].Duration != actual[i].Duration {
// 			t.Errorf("Create Process List failed, Duration does not match")
// 		}
// 	}
// }
