package storage

import (
	"testing"
	"os"
	"strings"
	"bytes"
	"encoding/binary"
	"log"
	"io/ioutil"
)

func GetStorage(t *testing.T) *CephStorage{
	confPath := os.Args[len(os.Args)-1]
	if strings.Index(confPath, ".conf") == -1 {
		t.Fatalf("No Specified conf file. \nUsage:\n go test [-v] [ceph_test.go] CONF_FILE")
	}

	logger := log.New(os.Stdout, "[yig]", log.LstdFlags)
	return NewCephStorage(confPath, logger)
}

func TestCephStorage_GetUniqUploadName(t *testing.T) {
	c := GetStorage(t)
	if c == nil {
		t.Fatal("Failed.")
	}
	a := c.GetUniqUploadName()
	b := c.GetUniqUploadName()
	if a == b {
		t.Fatal("GetUniqUploadName Failed.")
	}
}

func TestCephStorage_GetUsedSpacePercent(t *testing.T) {
	c := GetStorage(t)
	if c == nil {
		t.Fatal("Failed.")
	}
	_, err := c.GetUsedSpacePercent()
	if err != nil {
		t.Fatal("GetUsedSpacePercent err:", err)
	}
}

func TestCephStorage_IO(t *testing.T) {
	PoolName := "tiger"
	Oid := "test"
	c := GetStorage(t)
	if c == nil {
		t.Fatal("Failed.")
	}
	//10MB
	buff := new(bytes.Buffer)
	for i:= 0 ; i < (10 << 20) ;i++ {
		binary.Write(buff, binary.BigEndian,byte(1))
	}

	size, err := c.Put(PoolName, Oid, buff)
	if err != nil {
		t.Fatal("Put err:", err)
	}
	if size != (10 << 20) {
		t.Fatal("upload size should be 10MB")
	}

	reader, err := c.getReader(PoolName, Oid, 0, size)
	if err != nil {
		t.Fatal("getReader err:", err)
	}
	defer reader.Close()

	data, err := ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal("ReadAll err:", err)
	}

	if len(data) != (10 << 20) {
		t.Fatal("ReadAll size should be 10MB. size:", len(data))
	}

	err = c.Remove(PoolName, Oid)
	if err != nil {
		t.Fatal("Remove err:", err)
	}
}



