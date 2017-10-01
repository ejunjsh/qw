package main

import "fmt"

const (
	color_red = uint8(iota + 91)
	color_green
	color_yellow
	color_blue
	color_magenta
)

const format  = "\x1b[%dm%s\x1b[0m"

func red(s string, args ...interface{}) {
	fmt.Println(color(s,color_red,args...))
}

func green(s string, args ...interface{})  {
	fmt.Println(color(s,color_green,args...))
}

func yellow(s string, args ...interface{})  {
	fmt.Println(color(s,color_yellow,args...))
}

func blue(s string, args ...interface{})  {
	fmt.Println(color(s,color_blue,args...))
}

func magenta(s string, args ...interface{})  {
	fmt.Println(color(s,color_magenta,args...))
}

func color(s string, c uint8,args ...interface{}) string{
	var ss string
	if len(args)>0{
		ss=fmt.Sprintf(s,args...)
		return fmt.Sprintf(format, c, ss)
	}
	return fmt.Sprintf(format, c, s)
}