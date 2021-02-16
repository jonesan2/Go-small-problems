package main

import (
	"fmt"
	"strconv"
	s "strings"
)

func validateIPv4Byte(ipByte string) bool {
	var intByte, _ = strconv.Atoi(ipByte)
	if intByte == 0 && ipByte == "0" {
		return true
	} else if intByte >= 1 && intByte <= 255 && string(ipByte[0]) != "0" {
		return true
	} else {
		return false
	}
}

func validateIPv6Group(ipByte string) bool {
	var validChars = "0123456789abcdefABCDEF"
	if len(ipByte) < 1 || len(ipByte) > 4 {
		return false
	}
	for _, char := range ipByte {
		if !s.Contains(validChars, string(char)) {
			return false
		}
	}
	return true
}

func validateIPv4(ip string) bool {
	var ipBytes = s.Split(ip, ".")
	if len(ipBytes) != 4 {
		return false
	}
	for _, ipByte := range ipBytes {
		if !validateIPv4Byte(ipByte) {
			return false
		}
	}
	return true
}

func validateIPv6(ip string) bool {
	var ipBytes = s.Split(ip, ":")
	if len(ipBytes) != 8 {
		return false
	}
	for _, ipByte := range ipBytes {
		if !validateIPv6Group(ipByte) {
			return false
		}
	}
	return true
}

func validateIP(ip string) string {
	if validateIPv4(ip) {
		return "IPv4"
	} else if validateIPv6(ip) {
		return "IPv6"
	} else {
		return "Neither"
	}
}

func main() {
	fmt.Println(validateIP("192.168.0.1"))
	fmt.Println(validateIP("192.168.0"))
	fmt.Println(validateIP("192.168.0.01"))
	fmt.Println(validateIP("192.168.00.1"))
	fmt.Println(validateIP("2001:0db8:85a3:0000:0000:8a2e:0370:7334"))
	fmt.Println(validateIP("20F1:0db8:85a3:0000:00B0:8a2e:0370:733A"))
	fmt.Println(validateIP("20F1:0db8:85g3:0000:00B0:8a2e:0370:733A"))
	fmt.Println(validateIP("0"))
}
