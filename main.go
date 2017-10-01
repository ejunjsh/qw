package main

import "os"

const logo  = `
   ██████╗ ██╗    ██╗
  ██╔═══██╗██║    ██║
  ██║   ██║██║ █╗ ██║
  ██║▄▄ ██║██║███╗██║
  ╚██████╔╝╚███╔███╔╝
   ╚══▀▀═╝  ╚══╝╚══╝
`

func usage() {
	green(logo)
	red("Usage:")
	blue("qw <word to query>        Query the word")
}

func main(){
	if len(os.Args) == 2 {
		query(os.Args[1])
	}else {
		usage()
		os.Exit(0)
	}
}
