package dictionaryReplacer

import "testing"
import "fmt"

	func TestEmptyStringAndDictionary (t *testing.T){
		myDictionary := map[string]string{}
		var replaced = replace("", myDictionary)
		if (replaced != ""){
			t.Errorf("Empty replacement was incorrect")
		} else{
			fmt.Println("OK: " + replaced)
		}
	}

	func TestDictionaryWithOneWord(t *testing.T){
		myDictionary := make(map[string]string)
		myDictionary["temp"]="temporary"

		var replaced = replace("$temp$", myDictionary)
		if (replaced != "temporary"){
			t.Errorf("One word replacement was incorrect")
		} else{
			fmt.Println("OK: " + replaced)
		}
	}

	func TestDictionaryWithALongSentence(t *testing.T){
		
		myDictionary := make(map[string]string)
		myDictionary["temp"]="temporary"
		myDictionary["name"]="John Doe"

		var replaced = replace("$temp$ here comes the name $name$", myDictionary)

		if (replaced != "temporary here comes the name John Doe"){
			t.Errorf("Long sentence replacement was incorrect")
		} else{
			fmt.Println("OK: " + replaced)
		}
	}