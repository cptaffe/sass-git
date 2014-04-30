package main

import(
	"fmt"
	"os"
	"os/exec"
	"log"
)

const(
	// colors (linux)
	BLUE="\033[36m"
	RED="\033[31m"
	ENDC="\033[0m"
	BEGIN=BLUE+"# "+ENDC
	CMDOUT=RED+"> "+ENDC
)

type SassGit struct{
	file string
	msg string
	push bool
}

func ParseArgs()(*SassGit){
	s := new(SassGit)
	if len(os.Args) < 2{
		fmt.Printf(BEGIN+"Sorry, first argument("+RED+"!"+ENDC+")\n")
	} else {
		for i := 1; i < len(os.Args); i++ {
			if os.Args[i][0] == '-'{
				// flag
				if os.Args[i] == "-p" {
					s.push = true
				}
				if os.Args[i] == "-m" {
					// begin quoted string
					i++
					s.msg = os.Args[i]
				}
			}else {
				s.file = os.Args[i]
			}
		}
	}
	return s
}

func main(){
	s := ParseArgs()
	fmt.Printf("%s\n", s);
	
	// sass compile to file
	fmt.Printf(BEGIN+"Compiling %[1]s.scss %[1]s.css\n", s.file)
	cmd := exec.Command("sass", s.file+".scss", s.file+".css")
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	
	// git add
	fmt.Printf(BEGIN+"Adding %[1]s.scss %[1]s.css\n", s.file)
	err = exec.Command("git", "add", s.file+".scss", s.file+".css").Run()
	if err != nil {
		log.Fatal(err)
	}
	
	// git commit
	fmt.Printf(BEGIN+"Committing \"%s\"\n", s.msg)
	err = exec.Command("git", "commit", "-m", s.msg).Run()
	if err != nil {
		log.Fatal(err)
	}
	
	if s.push == true {
		// git push
		fmt.Printf(BEGIN+"Pushing\n")
		err = exec.Command("git", "push").Run()
		if err != nil {
			log.Fatal(err)
		}
	}
}
