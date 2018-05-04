package rfc5424

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSetTimestamp(t *testing.T) {
	m := &SyslogMessage{}

	assert.Equal(t, time.Date(2003, 10, 11, 22, 14, 15, 0, time.UTC), *m.SetTimestamp("2003-10-11T22:14:15Z").Timestamp)
	assert.Equal(t, time.Date(2003, 10, 11, 22, 14, 15, 3000, time.UTC), *m.SetTimestamp("2003-10-11T22:14:15.000003Z").Timestamp)
}

func TestSetNilTimestamp(t *testing.T) {
	m := &SyslogMessage{}
	assert.Nil(t, m.SetTimestamp("-").Timestamp)
}

func TestSetIncompleteTimestamp(t *testing.T) {
	m := &SyslogMessage{}
	date := []byte("2003-11-02T23:12:46.012345")
	prev := make([]byte, 0, len(date))
	for _, d := range date {
		prev = append(prev, d)
		assert.Nil(t, m.SetTimestamp(string(prev)).Timestamp)
	}
}

func TestSetSyntacticallyCompleteButIncorrectTimestamp(t *testing.T) {
	m := &SyslogMessage{}
	assert.Nil(t, m.SetTimestamp("2003-42-42T22:14:15Z").Timestamp)
}

func TestSetImpossibleButSyntacticallyCorrectTimestamp(t *testing.T) {
	m := &SyslogMessage{}
	assert.Nil(t, m.SetTimestamp("2003-09-31T22:14:15Z").Timestamp)
}

func TestSetTooLongHostname(t *testing.T) {
	m := &SyslogMessage{}
	m.SetHostname("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcX")
	assert.Nil(t, m.Hostname)
}

func TestSetNilOrEmptyHostname(t *testing.T) {
	m := &SyslogMessage{}
	assert.Nil(t, m.SetHostname("-").Hostname)
	assert.Nil(t, m.SetHostname("").Hostname)
}

func TestSetHostname(t *testing.T) {
	m := &SyslogMessage{}

	maxlen := []byte("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabc")

	prev := make([]byte, 0, len(maxlen))
	for _, input := range maxlen {
		prev = append(prev, input)
		str := string(prev)
		assert.Equal(t, str, *m.SetHostname(str).Hostname)
	}
}

func TestSetAppname(t *testing.T) {
	m := &SyslogMessage{}

	maxlen := []byte("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdef")

	prev := make([]byte, 0, len(maxlen))
	for _, input := range maxlen {
		prev = append(prev, input)
		str := string(prev)
		assert.Equal(t, str, *m.SetAppname(str).Appname)
	}
}

func TestSetProcID(t *testing.T) {
	m := &SyslogMessage{}

	maxlen := []byte("abcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzabcdefghilmnopqrstuvzab")

	prev := make([]byte, 0, len(maxlen))
	for _, input := range maxlen {
		prev = append(prev, input)
		str := string(prev)
		assert.Equal(t, str, *m.SetProcID(str).ProcID)
	}
}

func TestSetMsgID(t *testing.T) {
	m := &SyslogMessage{}

	maxlen := []byte("abcdefghilmnopqrstuvzabcdefghilm")

	prev := make([]byte, 0, len(maxlen))
	for _, input := range maxlen {
		prev = append(prev, input)
		str := string(prev)
		assert.Equal(t, str, *m.SetMsgID(str).MsgID)
	}
}

func TestSetSyntacticallyWrongHostnameAppnameProcIDMsgID(t *testing.T) {
	m := &SyslogMessage{}
	assert.Nil(t, m.SetHostname("white space not possible").Hostname)
	assert.Nil(t, m.SetHostname(string([]byte{0x0})).Hostname)
	assert.Nil(t, m.SetAppname("white space not possible").Appname)
	assert.Nil(t, m.SetAppname(string([]byte{0x0})).Appname)
	assert.Nil(t, m.SetProcID("white space not possible").ProcID)
	assert.Nil(t, m.SetProcID(string([]byte{0x0})).ProcID)
	assert.Nil(t, m.SetMsgID("white space not possible").MsgID)
	assert.Nil(t, m.SetMsgID(string([]byte{0x0})).MsgID)
}

func TestSetMessage(t *testing.T) {
	m := &SyslogMessage{}
	greek := "κόσμε"
	assert.Equal(t, greek, *m.SetMessage(greek).Message)
}

func TestSetEmptyMessage(t *testing.T) {
	m := &SyslogMessage{}
	m.SetMessage("")
	assert.Nil(t, m.Message)
}

func TestSetWrongUTF8Message(t *testing.T) {}

func TestSetMessageWithBOM(t *testing.T) {}

func TestSetOutOfRangeVersion(t *testing.T) {
	m := &SyslogMessage{}
	m.SetVersion(1000)
	assert.Equal(t, m.Version, uint16(0)) // 0 signals nil for version
	m.SetVersion(0)
	assert.Equal(t, m.Version, uint16(0)) // 0 signals nil for version
}

func TestSetOutOfRangePriority(t *testing.T) {
	m := &SyslogMessage{}
	m.SetPriority(192)
	assert.Nil(t, m.Priority)
}

func TestSetVersion(t *testing.T) {
	m := &SyslogMessage{}
	m.SetVersion(1)
	assert.Equal(t, m.Version, uint16(1))
	m.SetVersion(999)
	assert.Equal(t, m.Version, uint16(999))
}

func TestSetPriority(t *testing.T) {
	m := &SyslogMessage{}
	m.SetPriority(0)
	assert.Equal(t, *m.Priority, uint8(0))
	m.SetPriority(1)
	assert.Equal(t, *m.Priority, uint8(1))
	m.SetPriority(191)
	assert.Equal(t, *m.Priority, uint8(191))
}

func TestSetSDID(t *testing.T) {
	identifier := "one"
	m := &SyslogMessage{}
	assert.Nil(t, m.StructuredData)
	m.SetElementID(identifier)
	sd := m.StructuredData
	assert.NotNil(t, sd)
	assert.IsType(t, (*map[string]map[string]string)(nil), sd)
	assert.NotNil(t, (*sd)[identifier])
	assert.IsType(t, map[string]string{}, (*sd)[identifier])
	m.SetElementID(identifier)
	assert.Len(t, *sd, 1)
}

func TestSetAllLenghtsSDID(t *testing.T) {
	m := &SyslogMessage{}

	maxlen := []byte("abcdefghilmnopqrstuvzabcdefghilm")

	prev := make([]byte, 0, len(maxlen))
	for i, input := range maxlen {
		prev = append(prev, input)
		id := string(prev)
		m.SetElementID(id)
		assert.Len(t, *m.StructuredData, i+1)
		assert.IsType(t, map[string]string{}, (*m.StructuredData)[id])
	}
}

func TestSetTooLongSDID(t *testing.T) {
	m := &SyslogMessage{}
	m.SetElementID("abcdefghilmnopqrstuvzabcdefghilmX")
	assert.Nil(t, m.StructuredData)
}

func TestSetSyntacticallyWrongSDID(t *testing.T) {
	m := &SyslogMessage{}
	m.SetElementID("no whitespaces")
	assert.Nil(t, m.StructuredData)
	m.SetElementID(" ")
	assert.Nil(t, m.StructuredData)
	m.SetElementID(`"`)
	assert.Nil(t, m.StructuredData)
	m.SetElementID(`no"`)
	assert.Nil(t, m.StructuredData)
	m.SetElementID(`"no`)
	assert.Nil(t, m.StructuredData)
	m.SetElementID("]")
	assert.Nil(t, m.StructuredData)
	m.SetElementID("no]")
	assert.Nil(t, m.StructuredData)
	m.SetElementID("]no")
	assert.Nil(t, m.StructuredData)
}

func TestSetEmptySDID(t *testing.T) {
	m := &SyslogMessage{}
	m.SetElementID("")
	assert.Nil(t, m.StructuredData)
}

func TestSetSDParam(t *testing.T) {
	id := "one"
	pn := "pname"
	pv := "pvalue"
	m := &SyslogMessage{}
	m.SetParameter(id, pn, pv)
	sd := m.StructuredData
	assert.NotNil(t, sd)
	assert.IsType(t, (*map[string]map[string]string)(nil), sd)
	assert.NotNil(t, (*sd)[id])
	assert.IsType(t, map[string]string{}, (*sd)[id])
	assert.Len(t, *sd, 1)
	assert.Len(t, (*sd)[id], 1)
	assert.Equal(t, pv, (*sd)[id][pn])

	pn1 := "pname1"
	pv1 := "pvalue1"
	m.SetParameter(id, pn1, pv1)
	assert.Len(t, (*sd)[id], 2)
	assert.Equal(t, pv1, (*sd)[id][pn1])

	id1 := "another"
	m.SetParameter(id1, pn1, pv1).SetParameter(id1, pn, pv)
	assert.Len(t, *sd, 2)
	assert.Len(t, (*sd)[id1], 2)
	assert.Equal(t, pv1, (*sd)[id1][pn1])
	assert.Equal(t, pv, (*sd)[id1][pn])

	id2 := "tre"
	pn2 := "meta"
	m.SetParameter(id2, pn, `is\\valid`).SetParameter(id2, pn1, `is\]valid`).SetParameter(id2, pn2, `is\"valid`)
	assert.Len(t, *sd, 3)
	assert.Len(t, (*sd)[id2], 3)
	assert.Equal(t, `is\valid`, (*sd)[id2][pn])
	assert.Equal(t, `is]valid`, (*sd)[id2][pn1])
	assert.Equal(t, `is"valid`, (*sd)[id2][pn2])
	// Cannot contain \, ], " unless escaped
	m.SetParameter(id2, pn, `is\valid`).SetParameter(id2, pn1, `is]valid`).SetParameter(id2, pn2, `is"valid`)
	assert.Len(t, (*sd)[id2], 3)
}

func TestSetEmptySDParam(t *testing.T) {
	id := "id"
	pn := "pn"
	m := &SyslogMessage{}
	m.SetParameter(id, pn, "")
	sd := m.StructuredData
	assert.Len(t, *sd, 1)
	assert.Len(t, (*sd)[id], 1)
	assert.Equal(t, "", (*sd)[id][pn])
}
