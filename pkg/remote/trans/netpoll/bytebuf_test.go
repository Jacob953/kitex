/*
 * Copyright 2021 CloudWeGo Authors
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package netpoll

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/cloudwego/kitex/internal/test"
	"github.com/cloudwego/kitex/pkg/remote"
	"github.com/cloudwego/netpoll"
)

func TestByteBuffer(t *testing.T) {
	msg := "hello world"
	reader := bufio.NewReader(strings.NewReader(msg + msg + msg + msg + msg))
	writer := bufio.NewWriter(bytes.NewBufferString(msg + msg + msg + msg + msg))
	bufRW := bufio.NewReadWriter(reader, writer)
	rw := netpoll.NewReadWriter(bufRW)

	buf1 := NewReaderWriterByteBuffer(rw)
	checkWritable(t, buf1)
	checkReadable(t, buf1)

	buf2 := buf1.NewBuffer()
	checkWritable(t, buf2)
	checkUnreadable(t, buf2)

	wi := bytes.NewBufferString(msg + msg + msg + msg + msg)
	buf3 := NewReaderWriterByteBuffer(netpoll.NewReadWriter(wi))
	checkWritable(t, buf3)

	err := buf1.AppendBuffer(buf2)
	test.Assert(t, err == nil)
}

func TestWriterBuffer(t *testing.T) {
	msg := "hello world"
	wi := bytes.NewBufferString(msg + msg + msg + msg + msg)
	w := netpoll.NewWriter(wi)
	buf := NewWriterByteBuffer(w)
	checkWritable(t, buf)
	checkUnreadable(t, buf)
	err := w.WriteDirect([]byte(msg), 11)
	test.Assert(t, err == nil, err)
}

func TestReaderBuffer(t *testing.T) {
	msg := "hello world"
	ri := strings.NewReader(msg + msg + msg + msg + msg)
	r := netpoll.NewReader(ri)
	buf := NewReaderByteBuffer(r)
	checkUnwritable(t, buf)
	checkReadable(t, buf)
}

func checkWritable(t *testing.T, buf remote.ByteBuffer) {
	msg := "hello world"

	p, err := buf.Malloc(len(msg))
	test.Assert(t, err == nil, err)
	test.Assert(t, len(p) == len(msg))
	copy(p, msg)
	l := buf.MallocLen()
	test.Assert(t, l == len(msg))
	l, err = buf.WriteString(msg)
	test.Assert(t, err == nil, err)
	test.Assert(t, l == len(msg))
	l, err = buf.WriteBinary([]byte(msg))
	test.Assert(t, err == nil, err)
	test.Assert(t, l == len(msg))
	err = buf.Flush()
	test.Assert(t, err == nil, err)
}

func checkReadable(t *testing.T, buf remote.ByteBuffer) {
	msg := "hello world"

	p, err := buf.Peek(len(msg))
	test.Assert(t, err == nil, err)
	test.Assert(t, string(p) == msg)
	err = buf.Skip(1 + len(msg))
	test.Assert(t, err == nil, err)
	p, err = buf.Next(len(msg) - 1)
	test.Assert(t, err == nil, err)
	test.Assert(t, string(p) == msg[1:])
	n := buf.ReadableLen()
	test.Assert(t, n == 3*len(msg), n)
	n = buf.ReadLen()
	test.Assert(t, n == len(msg)-1, n)

	var s string
	s, err = buf.ReadString(len(msg))
	test.Assert(t, err == nil, err)
	test.Assert(t, s == msg)
	p, err = buf.ReadBinary(len(msg))
	test.Assert(t, err == nil, err)
	test.Assert(t, string(p) == msg)
}

func checkUnwritable(t *testing.T, buf remote.ByteBuffer) {
	msg := "hello world"
	_, err := buf.Malloc(len(msg))
	test.Assert(t, err != nil)
	l := buf.MallocLen()
	test.Assert(t, l == -1, l)
	_, err = buf.WriteString(msg)
	test.Assert(t, err != nil)
	_, err = buf.WriteBinary([]byte(msg))
	test.Assert(t, err != nil)
	err = buf.Flush()
	test.Assert(t, err != nil)
	var n int
	n, err = buf.Write([]byte(msg))
	test.Assert(t, err != nil)
	test.Assert(t, n == -1, n)
	_, err = buf.Write(nil)
	test.Assert(t, err != nil)
}

func checkUnreadable(t *testing.T, buf remote.ByteBuffer) {
	msg := "hello world"
	_, err := buf.Peek(len(msg))
	test.Assert(t, err != nil)
	err = buf.Skip(1)
	test.Assert(t, err != nil)
	_, err = buf.Next(len(msg) - 1)
	test.Assert(t, err != nil)
	n := buf.ReadableLen()
	test.Assert(t, n == -1)
	n = buf.ReadLen()
	test.Assert(t, n == 0)
	_, err = buf.ReadString(len(msg))
	test.Assert(t, err != nil)
	_, err = buf.ReadBinary(len(msg))
	test.Assert(t, err != nil)
	p := make([]byte, len(msg))
	n, err = buf.Read(p)
	test.Assert(t, err != nil)
	test.Assert(t, n == -1, n)
	_, err = buf.Read(nil)
	test.Assert(t, err != nil)
}
