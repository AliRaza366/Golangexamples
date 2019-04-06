package golangexamples


import "github.com/ehteshamz/greetings"

func ConcatSlice(sliceToConcat []byte) string{
	Sliced := ""
	length:=len(sliceToConcat)
	for i:=0;  i < length; i++{

		if string(sliceToConcat[i])!=" " {
			//fmt.Println(sliceToConcat[i])
			Sliced += string(sliceToConcat[i])
			if i!=length-1{
				Sliced+="-"
			}
		}
	}
	return Sliced
}

func Encrypt(sliceToEncrypt	[]byte,	ceaserCount	int){
	EncryptedString:=""
	length:=len(sliceToEncrypt)
	for i:=0; i<length; i++{
		EncryptedString+=string(int(sliceToEncrypt[i])+ceaserCount)
	}
	fmt.Println("Encrypted String is: ")
	fmt.Println(EncryptedString)
}

func EZGreetings(name string) string{

		temp := greetings.PrintGreetings(name)
		return temp
	
	}
	
	
