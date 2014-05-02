package main

import(
	"fmt"
	"os"
	"os/exec"
	"errors"
)

const(
	// colors (linux)
	BLUE="\033[36m"
	//BLUE="\033[38;5;220m" -- 256 colors
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

// Parses arguments to create struct
func ParseArgs()(*SassGit, error){
	s := new(SassGit)
	if len(os.Args) < 2{
		return s, errors.New("not enough args")
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
		if s.msg == ""{
			s.msg = fmt.Sprintf("Compiled %[1]s.scss to %[1]s.css", s.file)
		}
		return s, nil
	}
}

func (s *SassGit) CmdExec() error {
	// sass compile to file
	PrintMsg(fmt.Sprintf("Compiling %[1]s.scss %[1]s.css\n", s.file))
	cmd := exec.Command("sass", s.file+".scss", s.file+".css")
	str, err := cmd.Output()
	if err != nil {
		return err
	} else if string(str) != "" {
		PrintExec(str)
	}
	
	
	// git add
	PrintMsg(fmt.Sprintf("Adding %[1]s.scss %[1]s.css\n", s.file))
	str, err = exec.Command("git", "add", s.file+".scss", s.file+".css").Output()
	if err != nil {
		return err
	} else if string(str) != "" {
		PrintExec(str)
	}
	
	// git commit
	PrintMsg(fmt.Sprintf("Committing \"%s\"\n", s.msg))
	str, err = exec.Command("git", "commit", "-m", s.msg).Output()
	if err != nil {
		return err
	} else if string(str) != "" {
		PrintExec(str)
	}
	
	if s.push == true {
		// git push
		PrintMsg("Pushing\n")
		str, err = exec.Command("git", "push").Output()
		if err != nil {
			return err
		} else if string(str) != "" {
			PrintExec(str)
		}
	}
	return nil;
}

func PrintError(err error){
	PrintMsg(fmt.Sprintf("Sorry, %s ("+RED+"!"+ENDC+")\n", err))
}

func PrintMsg(s string){
	fmt.Printf(BEGIN+"%s", s)
}

func PrintExec(b []byte){
	s := string(b)
	sub := ""
	for i := 0; i < len(s); i++ {
		if s[i] != '\n'{
			sub += string(s[i])
		} else {
			fmt.Printf(CMDOUT+"%s\n", sub)
			sub = ""
		}
	}
}

// calls functions
func main(){
	s, err := ParseArgs()
	if err != nil{
		PrintError(err)
	}else{
		err = s.CmdExec()
		if err != nil{
			PrintError(err)
		}
	}
}
