package main

import (
	"encryptString/utils"
	"fmt"
)

// AES Keys should be of length 16, 24, 32
func main() {
	key := "111023043350789514532147"
	message := "jiI am A message"
	fmt.Println("Original mgs: ", message)
	encryptedString := utils.EncryptString(key, message)
	fmt.Println("Encrypted msg: ", encryptedString)
	decryptedString := utils.DecryptString(key, encryptedString)
	fmt.Println("Decrypted msg", decryptedString)

}
