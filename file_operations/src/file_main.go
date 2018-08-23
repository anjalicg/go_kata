package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

func main() {
	data, err := ioutil.ReadFile("testfile.txt")
	if err == nil {
		data_str := string(data)
		fmt.Println(string(data_str))
		fmt.Println(len(data_str))
		fmt.Println(reflect.TypeOf(data_str))
		fmt.Printf("Type of data_str= %T\n", data_str)
		fmt.Println(reflect.TypeOf(data_str))
		fmt.Println(reflect.TypeOf(data))
		fmt.Println(len(data))
		fmt.Printf("Type of data= %T\n", data)
	} else {
		fmt.Println("Error!!", err)
	}
	f, err := os.Open("testfile.txt")
	defer f.Close()
	b1 := make([]byte, 20)
	if err == nil {
		fmt.Printf("%T\n", f)

		n1, err := f.Read(b1)
		if err == nil {
			fmt.Println(n1, len(b1), string(b1))
		}
		o2, err := f.Seek(6, 0)
		if err == nil {
			fmt.Println("seek output", o2)
		}
		n1, err = f.Read(b1)
		if err == nil {
			fmt.Println(n1, len(b1), string(b1))
		}
	} else {
		fmt.Println("Error!!", err)
	}

	err = ioutil.WriteFile("testfile_write.txt", b1, 0644)
	if err != nil {
		fmt.Println("Error happened", err)

	} else {
		fmt.Println("ERROR is NIL")
	}

	f, err = os.OpenFile("testfile_write.txt", os.O_APPEND|os.O_WRONLY, 0644)
	defer f.Close()
	d2 := []byte("Dummy stufff for writing")
	f.Write(d2)

	o2, err := f.Seek(2, 0)
	if err == nil {
		fmt.Println("seek output", o2)
	}
	n1, err := f.Write([]byte("THIS IS ALL CAPS"))
	if err == nil {
		fmt.Println("after write", n1, len(b1))
	} else {
		fmt.Println("ERROR in writing all caps")
	}
}
