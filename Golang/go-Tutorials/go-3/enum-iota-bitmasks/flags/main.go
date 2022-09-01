package main

import "fmt"

type MyConf uint8

const (
	VERBOSE           MyConf = 1 << 0
	CONFIG_FROM_DISK  MyConf = 1 << 1
	DATABASE_REQUIRED MyConf = 1 << 2
	LOGGER_ACTIVATED  MyConf = 1 << 3
	DEBUG             MyConf = 1 << 4
	FLOAT_SUPPORT     MyConf = 1 << 5
	RECOVERY_MODE     MyConf = 1 << 6
	REBOOT_ON_FAILURE MyConf = 1 << 7
)

func main() {
	MyComplexFunction(VERBOSE|REBOOT_ON_FAILURE, "test")
}

func MyComplexFunction(conf MyConf, databaseDsn string) {
	fmt.Printf("conf : %08b\n", conf)
	test := conf & REBOOT_ON_FAILURE
	fmt.Printf("test : %08b\n", test)
}
