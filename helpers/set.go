package helpers

import (
	"fmt"

	"github.com/byuoitav/common/nerr"
)

//Turn the light on
func LightOn(ipaddress, light string) (string, *nerr.E) {
	Conn, err := MakeConnection(ipaddress)
	if err != nil {
		return "", nerr.Create(fmt.Sprintf("%v", err), "")
	}
	defer Conn.Close()
	msg := fmt.Sprintf("{\"SetCh%s\":\"ON\",\"SetCh%s\":100}", light, light)
	fmt.Printf("Message: %s", msg)
	buf := []byte(msg)
	_, err = Conn.Write(buf)
	if err != nil {
		fmt.Println(msg, err)
	}
	return "", nil
}

//Turn the light off
func LightOff(ipaddress, light string) (string, *nerr.E) {
	Conn, err := MakeConnection(ipaddress)
	if err != nil {
		return "", nerr.Create(fmt.Sprintf("%v", err), "")
	}
	defer Conn.Close()
	msg := fmt.Sprintf("{\"SetCh%s\":\"OFF\",\"SetCh%s\":100}", light, light)
	fmt.Printf("Message: %s", msg)
	buf := []byte(msg)
	_, err = Conn.Write(buf)
	if err != nil {
		fmt.Println(msg, err)
	}
	return "", nil
}

//This is the Dim
func Dim(ipaddress, light, dimlevel string) (string, *nerr.E) {
	Conn, err := MakeConnection(ipaddress)
	if err != nil {
		return "", nerr.Create(fmt.Sprintf("%v", err), "")
	}
	defer Conn.Close()
	msg := fmt.Sprintf("{\"SetCh%s\":\"ON\",\"SetCh%s\":%s}", light, light, dimlevel)
	fmt.Printf("Message: %s", msg)
	buf := []byte(msg)
	_, err = Conn.Write(buf)
	if err != nil {
		fmt.Println(msg, err)
	}
	return "", nil
}
