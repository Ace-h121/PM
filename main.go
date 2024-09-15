package main

import (
	"crypto/aes"
	"encoding/hex"
	"errors"
	"fmt"
	"os"

	"Github.com/Ace-h121/PM/tree"
	"github.com/thanhpk/randstr"
)

func main(){
	// cipher key
	args := os.Args
	if(len(args) < 2){
		printHelp()
		os.Exit(1)
	}
	dir, err := os.UserHomeDir()
	if err!= nil{
		fmt.Println("Error: Could not find the users home directory, please ensure it exists")
		os.Exit(1)
	}

	switch(args[1]){
		case "setup":
			if _, err := os.Stat(dir + "/.config/PM.conf"); errors.Is(err, os.ErrNotExist){
				fmt.Println(err)
				if err := os.WriteFile(dir +".config/PM.conf", []byte(randstr.String(16)), 0666); err != nil{
					fmt.Println(err)
				}
				fmt.Println("Finished Setup")
			} else {
				fmt.Println("Setup is already complete")
			}
			err :=os.Mkdir(dir + "/PM/", 0777 )
		if err!=nil{
			fmt.Println(err)
		}
		case "generate":
			if len(os.Args) < 3{
				fmt.Println("Error: Missing arguments. Please provide a username. Usage: generate <username>")
				os.Exit(1)
			}
			username := os.Args[2]
			fmt.Printf("Generating new password for %s \n", username)
			password := randstr.String(32)
			fmt.Printf("Password for %s is %s", username, password)
		case "list": 
			tree.Run()
	
		case "save":
			if len(os.Args) < 4{
				fmt.Println("Error: Missing arguments. Please provide both a filename and a password. Usage: save <filename> <password>")
				os.Exit(1)
			}
			err := os.Chdir(dir + "/PM/")
			key, err:= os.ReadFile(dir + "/.config/PM.conf" )
			if err != nil {
				fmt.Println("Error: Encryption key not found. Please run 'setup' to initialize the password manager.")
				os.Exit(1)
			}
			fmt.Println(string(key))
			encryptedPass := EncryptAES(key, os.Args[3])
			os.WriteFile(os.Args[2], []byte(encryptedPass), 0777)
			fmt.Println("Password saved successfully.")
		case "show":
			if len(os.Args) <3 {
				fmt.Println("Error: Missing arguments. Please provide a filename. Usage: show <filename>")
				os.Exit(1)
			}

			encrypedPass, err := os.ReadFile(dir + "/PM/" + os.Args[2])
			if err != nil {
				fmt.Println("Error: Unable to find the specified file. Ensure the file exists in the 'PM' directory.")
				os.Exit(1)
			}

			key, err:= os.ReadFile(dir + "/.config/PM.conf" )
			if err != nil {
				fmt.Println("Key not found, please run Setup")
				os.Exit(1)
			}
			
			decryptedPass := DecryptAES(key, string(encrypedPass))
			
			fmt.Printf("The password for %s is %s", os.Args[2], decryptedPass)
		case "help":
			printHelp()


		default:
			fmt.Printf("Error: '%s' is not a valid command. Use 'help' to see available commands.\n", os.Args[1])			

	}
}
	



func EncryptAES(key []byte, plaintext string) string {
 
    c, err := aes.NewCipher(key)
    if (err!=nil){
		panic(err)
	}
 
    out := make([]byte, len(plaintext))
 
    c.Encrypt(out, []byte(plaintext))
 
    return hex.EncodeToString(out)
}

func DecryptAES(key []byte, ct string) string {
    ciphertext, _ := hex.DecodeString(ct)
 
    c, err := aes.NewCipher(key)
	if (err != nil){
		panic(err)
	}
 
    pt := make([]byte, len(ciphertext))
    c.Decrypt(pt, ciphertext)
 
    s := string(pt[:])
    return s
}

func printHelp() {
	fmt.Println("Usage:")
	fmt.Println("  <command> [arguments]")
	fmt.Println("\nCommands:")
	fmt.Println("  setup                  Initialize the password manager by generating a config file.")
	fmt.Println("  generate <username>     Generate a new 32-character password for the specified username.")
	fmt.Println("  list                   List all saved passwords.")
	fmt.Println("  save <filename> <pass>  Save the specified password in the given file, encrypted with the setup key.")
	fmt.Println("  show <filename>         Show the decrypted password for the specified file.")
	fmt.Println("\nExamples:")
	fmt.Println("  setup                  Run this first to initialize the config.")
	fmt.Println("  generate alice          Generate a new password for 'alice'.")
	fmt.Println("  list                   List all stored passwords.")
	fmt.Println("  save mypass.txt abc123  Save the password 'abc123' into 'mypass.txt'.")
	fmt.Println("  show mypass.txt         Decrypt and show the password stored in 'mypass.txt'.")
	fmt.Println("\nRun 'help' to display this message again.")
}

