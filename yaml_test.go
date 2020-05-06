package conf

import (
	"testing"
)

func TestConfig(t *testing.T) {
	t.Log("mysql")
	t.Log(Conf.ZyjyMysql)
}