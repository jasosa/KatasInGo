package dictionaryReplacer

//import "fmt"
import "strings"

func replace (toReplace string, dictionary map[string]string) (string) {
	token := "$"
	for k, v := range dictionary {
	    toReplace = strings.Replace(toReplace, token+k+token, v, -1)
  }
	
	return toReplace
}
