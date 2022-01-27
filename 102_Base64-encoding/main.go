package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	s := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Proin ornare ut metus ut sagittis. Integer eget orci vestibulum, ultricies orci at, porttitor tortor. Nunc id porttitor turpis. Donec varius felis eu orci iaculis imperdiet ac nec ante. Cras quis vestibulum mauris. Donec eu cursus arcu, nec interdum leo. Vivamus tincidunt semper sapien et interdum. In lacinia risus eget lectus consequat, feugiat bibendum tortor tristique. Etiam varius mi et nunc dictum faucibus. Nunc feugiat, velit quis vehicula gravida, turpis tellus viverra velit, vitae aliquam felis elit sit amet nisl. Cras tincidunt elit et urna condimentum, et iaculis orci lobortis."

	std := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+_"
	s64 := base64.NewEncoding(std).EncodeToString([]byte(s))

	bs, err := base64.StdEncoding.DecodeString(s64)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(len(s))
	fmt.Println(len(s64))
	fmt.Println(s64)
	fmt.Println(string(bs))
}