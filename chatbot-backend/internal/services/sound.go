package services

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Play() {
	script := `
		live_loop :example do
			sample :loop_amen
			sleep 1
		end
		`
	scriptFilePath := "./temp_sonic_pi_script.rb"

	err := os.WriteFile(scriptFilePath, []byte(script), 0644)
	if err != nil {
		log.Fatal("Error writing Sonic Pi script:", err)
	}

	cmd0 := exec.Command("ls", ".")
	// Capture the output of the command
	out, err := cmd0.CombinedOutput()
	if err != nil {
		fmt.Println("Error executing ls command:", err)
	}
	// Print the output
	fmt.Println(string(out))

	// Execute Sonic Pi with the script file
	cmd := exec.Command("sonic-pi-tool", "run-file", scriptFilePath)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err = cmd.Run()
	if err != nil {
		fmt.Println("Error executing Sonic Pi script:", err)
	}

	// Clean up temporary script file
	err = os.Remove(scriptFilePath)
	if err != nil {
		fmt.Println("Error cleaning up:", err)
	}
}
