package utility

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"
)

func PrintError(errorInfo interface{}) {
	errorJSON, err := json.MarshalIndent(errorInfo, "", "    ")
	if err != nil {
		log.Fatal("Error while Marshaling the Error Object:", err)
		return
	}
	fmt.Println(string(errorJSON))
}

func PrintInfo(infoObject interface{}) {
	infoJSON, err := json.MarshalIndent(infoObject, "", "    ")
	if err != nil {
		log.Fatal("Error while Marshaling the Object:", err)
		return
	}
	fmt.Println(string(infoJSON))
}

func PrintStarter() {
	log.Print("Use the following flags according to your requirements to fully utilize WhoisFreaks' services.")
	flag.VisitAll(func(f *flag.Flag) {
		fmt.Printf("	-%s: %s\n", f.Name, f.Usage)
	})
	log.Print("For Detailed Usage you can consult with README.md.")
	os.Exit(0)
}
