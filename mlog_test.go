package mlog

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestMlog(t *testing.T) {
	tests := []struct {
		name string
		run  func(Logger)
		want string
	}{
		{
			name: "basic",
			run:  testAllMlogMethods,
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"e","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"w","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"we","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"i","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"ie","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"d","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"de","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"t","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"te","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"all","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"always","panda":2}
`,
		},
		{
			name: "with values",
			run: func(l Logger) {
				testAllMlogMethods(l.WithValues("hi", 42))
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"e","hi":42,"panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"w","hi":42,"warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"we","hi":42,"warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"i","hi":42,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"ie","hi":42,"error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"d","hi":42,"panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"de","hi":42,"error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"t","hi":42,"panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"te","hi":42,"error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"all","hi":42,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"always","hi":42,"panda":2}
`,
		},
		{
			name: "with values conflict", // duplicate key is included twice ...
			run: func(l Logger) {
				testAllMlogMethods(l.WithValues("panda", false))
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"e","panda":false,"panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"w","panda":false,"warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"we","panda":false,"warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"i","panda":false,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"ie","panda":false,"error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"d","panda":false,"panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"de","panda":false,"error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"t","panda":false,"panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"te","panda":false,"error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"all","panda":false,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"always","panda":false,"panda":2}
`,
		},
		{
			name: "with values nested",
			run: func(l Logger) {
				testAllMlogMethods(l.WithValues("hi", 42).WithValues("not", time.Hour))
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"e","hi":42,"not":"1h0m0s","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"w","hi":42,"not":"1h0m0s","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"we","hi":42,"not":"1h0m0s","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"i","hi":42,"not":"1h0m0s","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"ie","hi":42,"not":"1h0m0s","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"d","hi":42,"not":"1h0m0s","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"de","hi":42,"not":"1h0m0s","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"t","hi":42,"not":"1h0m0s","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"te","hi":42,"not":"1h0m0s","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"all","hi":42,"not":"1h0m0s","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"always","hi":42,"not":"1h0m0s","panda":2}
`,
		},
		{
			name: "with name",
			run: func(l Logger) {
				testAllMlogMethods(l.WithName("yoyo"))
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"e","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"w","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"we","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"i","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"ie","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"d","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"de","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"t","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"te","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"all","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"always","panda":2}
`,
		},
		{
			name: "with name nested",
			run: func(l Logger) {
				testAllMlogMethods(l.WithName("yoyo").WithName("gold"))
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo.gold","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"e","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo.gold","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"w","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo.gold","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"we","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo.gold","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"i","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo.gold","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"ie","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo.gold","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"d","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo.gold","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"de","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo.gold","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"t","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo.gold","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"te","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo.gold","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"all","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","logger":"yoyo.gold","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"always","panda":2}
`,
		},
		{
			name: "depth 3",
			run: func(l Logger) {
				testAllMlogMethods(l.withDepth(3))
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"testing/testing.go:<line>$testing.tRunner","message":"e","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"testing/testing.go:<line>$testing.tRunner","message":"w","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"testing/testing.go:<line>$testing.tRunner","message":"we","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"testing/testing.go:<line>$testing.tRunner","message":"i","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"testing/testing.go:<line>$testing.tRunner","message":"ie","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"testing/testing.go:<line>$testing.tRunner","message":"d","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"testing/testing.go:<line>$testing.tRunner","message":"de","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"testing/testing.go:<line>$testing.tRunner","message":"t","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"testing/testing.go:<line>$testing.tRunner","message":"te","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"testing/testing.go:<line>$testing.tRunner","message":"all","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"testing/testing.go:<line>$testing.tRunner","message":"always","panda":2}
`,
		},
		{
			name: "depth 2",
			run: func(l Logger) {
				testAllMlogMethods(l.withDepth(2))
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func15","message":"e","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func15","message":"w","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func15","message":"we","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func15","message":"i","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func15","message":"ie","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func15","message":"d","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func15","message":"de","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func15","message":"t","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func15","message":"te","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func15","message":"all","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func15","message":"always","panda":2}
`,
		},
		{
			name: "depth 1",
			run: func(l Logger) {
				testAllMlogMethods(l.withDepth(1))
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func8","message":"e","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func8","message":"w","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func8","message":"we","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func8","message":"i","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func8","message":"ie","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func8","message":"d","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func8","message":"de","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func8","message":"t","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func8","message":"te","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func8","message":"all","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func8","message":"always","panda":2}
`,
		},
		{
			name: "depth 0",
			run: func(l Logger) {
				testAllMlogMethods(l.withDepth(0))
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"e","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"w","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"we","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"i","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"ie","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"d","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"de","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"t","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"te","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"all","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.testAllMlogMethods","message":"always","panda":2}
`,
		},
		{
			name: "depth -1",
			run: func(l Logger) {
				testAllMlogMethods(l.withDepth(-1))
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.Error","message":"e","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.Warning","message":"w","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.WarningErr","message":"we","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.Info","message":"i","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.InfoErr","message":"ie","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.Debug","message":"d","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.DebugErr","message":"de","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.Trace","message":"t","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.TraceErr","message":"te","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.All","message":"all","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.Always","message":"always","panda":2}
`,
		},
		{
			name: "depth -2",
			run: func(l Logger) {
				testAllMlogMethods(l.withDepth(-2))
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"logr@v1.2.3/logr.go:<line>$logr.Logger.Error","message":"e","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.warningDepth","message":"w","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.warningDepth","message":"we","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.infoDepth","message":"i","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.infoDepth","message":"ie","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.debugDepth","message":"d","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.debugDepth","message":"de","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.traceDepth","message":"t","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.traceDepth","message":"te","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"logr@v1.2.3/logr.go:<line>$logr.Logger.Info","message":"all","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"logr@v1.2.3/logr.go:<line>$logr.Logger.Info","message":"always","panda":2}
`,
		},
		{
			name: "depth -3",
			run: func(l Logger) {
				testAllMlogMethods(l.withDepth(-3))
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"zapr@v1.2.3/zapr.go:<line>$zapr.(*zapLogger).Error","message":"e","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"logr@v1.2.3/logr.go:<line>$logr.Logger.Info","message":"w","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"logr@v1.2.3/logr.go:<line>$logr.Logger.Info","message":"we","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"logr@v1.2.3/logr.go:<line>$logr.Logger.Info","message":"i","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"logr@v1.2.3/logr.go:<line>$logr.Logger.Info","message":"ie","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"logr@v1.2.3/logr.go:<line>$logr.Logger.Info","message":"d","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"logr@v1.2.3/logr.go:<line>$logr.Logger.Info","message":"de","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"logr@v1.2.3/logr.go:<line>$logr.Logger.Info","message":"t","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"logr@v1.2.3/logr.go:<line>$logr.Logger.Info","message":"te","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"zapr@v1.2.3/zapr.go:<line>$zapr.(*zapLogger).Info","message":"all","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"zapr@v1.2.3/zapr.go:<line>$zapr.(*zapLogger).Info","message":"always","panda":2}
`,
		},
		{
			name: "closure",
			run: func(l Logger) {
				func() {
					func() {
						testErr := fmt.Errorf("some err")

						l.Error("e", testErr, "panda", 2)
						l.Warning("w", "panda", 2)
						l.WarningErr("we", testErr, "panda", 2)
						l.Info("i", "panda", 2)
						l.InfoErr("ie", testErr, "panda", 2)
						l.Debug("d", "panda", 2)
						l.DebugErr("de", testErr, "panda", 2)
						l.Trace("t", "panda", 2)
						l.TraceErr("te", testErr, "panda", 2)
						l.All("all", "panda", 2)
						l.Always("always", "panda", 2)
					}()
				}()
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func13.1.1","message":"e","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func13.1.1","message":"w","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func13.1.1","message":"we","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func13.1.1","message":"i","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func13.1.1","message":"ie","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func13.1.1","message":"d","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func13.1.1","message":"de","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func13.1.1","message":"t","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func13.1.1","message":"te","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func13.1.1","message":"all","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog_test.go:<line>$mlog.TestMlog.func13.1.1","message":"always","panda":2}
`,
		},
		{
			name: "closure depth -1",
			run: func(l Logger) {
				func() {
					func() {
						testErr := fmt.Errorf("some err")

						l = l.withDepth(-1)
						l.Error("e", testErr, "panda", 2)
						l.Warning("w", "panda", 2)
						l.WarningErr("we", testErr, "panda", 2)
						l.Info("i", "panda", 2)
						l.InfoErr("ie", testErr, "panda", 2)
						l.Debug("d", "panda", 2)
						l.DebugErr("de", testErr, "panda", 2)
						l.Trace("t", "panda", 2)
						l.TraceErr("te", testErr, "panda", 2)
						l.All("all", "panda", 2)
						l.Always("always", "panda", 2)
					}()
				}()
			},
			want: `
{"level":"error","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.Error","message":"e","panda":2,"error":"some err"}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.Warning","message":"w","warning":true,"panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.WarningErr","message":"we","warning":true,"error":"some err","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.Info","message":"i","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.InfoErr","message":"ie","error":"some err","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.Debug","message":"d","panda":2}
{"level":"debug","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.DebugErr","message":"de","error":"some err","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.Trace","message":"t","panda":2}
{"level":"trace","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.TraceErr","message":"te","error":"some err","panda":2}
{"level":"all","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.All","message":"all","panda":2}
{"level":"info","timestamp":"2099-08-08T13:57:36.123456Z","caller":"mlog/mlog.go:<line>$mlog.mLogger.Always","message":"always","panda":2}
`,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			var log bytes.Buffer
			tt.run(TestLogger(t, &log))

			require.Equal(t, strings.TrimSpace(tt.want), strings.TrimSpace(log.String()))
		})
	}
}

func testAllMlogMethods(l Logger) {
	testErr := fmt.Errorf("some err")

	l.Error("e", testErr, "panda", 2)
	l.Warning("w", "panda", 2)
	l.WarningErr("we", testErr, "panda", 2)
	l.Info("i", "panda", 2)
	l.InfoErr("ie", testErr, "panda", 2)
	l.Debug("d", "panda", 2)
	l.DebugErr("de", testErr, "panda", 2)
	l.Trace("t", "panda", 2)
	l.TraceErr("te", testErr, "panda", 2)
	l.All("all", "panda", 2)
	l.Always("always", "panda", 2)
}
