// Please, note that you will need to add a folder called libs here and put
//  the bouncycastle file bcprov-jdk15on-156.jar in it.
// Otherwise the wrapper and the makefile won't work.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	arg := []string{"-cp", "./:./examples/:./examples/libs/bcprov-jdk15on-156.jar:./libs/bcprov-jdk15on-156.jar", "dsa_sha256_java"}
	arg = append(arg, os.Args[1:]...)

	cmd := exec.Command("java", arg...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Println(strings.TrimSpace(string(out)))
		log.Fatal(err)
	}
	fmt.Println(strings.TrimSpace(string(out)))
}
