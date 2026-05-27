package helper

import (
	"io"

	"github.com/hdt3213/rdb/model"
)

const crlf = "\r\n"

// CmdLine is alias for [][]byte, represents a command line
type CmdLine = [][]byte

// lexOrder traversal map in lex order to create
type lexOrder struct{}

func makeMultiBulkResp(args [][]byte) []byte { _ = "STUB: not implemented"; return nil }

var setCmd = []byte("SET")

func stringToCmd(obj *model.StringObject) CmdLine { _ = "STUB: not implemented"; return *new(CmdLine) }

var rPushAllCmd = []byte("RPUSH")

func listToCmd(obj *model.ListObject) CmdLine { _ = "STUB: not implemented"; return *new(CmdLine) }

var sAddCmd = []byte("SADD")

func setToCmd(obj *model.SetObject) CmdLine { _ = "STUB: not implemented"; return *new(CmdLine) }

var hMSetCmd = []byte("HMSET")
var hPExpireAtCmd = []byte("HPEXPIREAT") // redis 7.4.0+
var hPersistCmd = []byte("HPERSIST")     // redis 7.4.0+

func hashToCmd(obj *model.HashObject, useLexOrder bool) []CmdLine {
	_ = "STUB: not implemented"
	return nil
}

// HPEXPIRE key seconds FIELDS num FIELD...

// HPEXPIRE key seconds FIELDS num FIELD...

var zAddCmd = []byte("ZADD")

func zSetToCmd(obj *model.ZSetObject) CmdLine { _ = "STUB: not implemented"; return *new(CmdLine) }

var pExpireAtBytes = []byte("PEXPIREAT")

// MakeExpireCmd generates command line to set expiration for the given key
func makeExpireCmd(obj model.RedisObject) CmdLine { _ = "STUB: not implemented"; return *new(CmdLine) }

var (
	xaddCmd = []byte("XADD")
)

func formatStreamID(streamID *model.StreamId) string { _ = "STUB: not implemented"; return "" }

func streamToCmd(stream *model.StreamObject) []CmdLine { _ = "STUB: not implemented"; return nil }

// TODO: groups, consumers, pending

// ObjectToCmd convert redis object to redis command line
func ObjectToCmd(obj model.RedisObject, opts ...interface{}) []CmdLine {
	_ = "STUB: not implemented"
	return nil
}

// CmdLinesToResp convert []CmdLine to RESP bytes
func CmdLinesToResp(cmds []CmdLine) []byte { _ = "STUB: not implemented"; return nil }

// WriteObjectToResp convert object to resp and write
func WriteObjectToResp(w io.Writer, obj model.RedisObject) error {
	_ = "STUB: not implemented"
	return nil
}

// the arg may be a large key or value
