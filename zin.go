//
// Payload Injector
//
// Coder: zoid

package main

import (
"sync"
"flag"
"log"
"bufio"
"os"
"fmt"
"net/url"
"net/http"
"net/http/httputil"
"runtime"
)


var Reset   = "\033[0m"
var Red     = "\033[31m"
var Green   = "\033[32m"
var Yellow  = "\033[33m"
var Blue    = "\033[34m"
var Purple  = "\033[35m"
var Cyan    = "\033[36m"
var Gray    = "\033[37m"
var White   = "\033[97m"

func init() {
	if runtime.GOOS == "windows" {
		Reset   = ""
		Red     = ""
		Green   = ""
		Yellow  = ""
		Blue    = ""
		Purple  = ""
		Cyan    = ""
		Gray    = ""
		White   = ""
	}
}


func main() {
	
	// Banner
	banner()
	
	// Payload to be used.
	concurrPtr:= flag.Int("c", 20, "the concurrency")
	payloadPtr := flag.String("p", "", "the payload to be used")
	payloadsPtr := flag.String("pL", "", "the list of payloads to be used")

	// Parse the arguments
	flag.Parse()

	if (*payloadPtr == "" && *payloadsPtr == "") {
		flag.PrintDefaults()
		return	
	}else{
        

		// Create the output directory
 		_,err := os.Stat("output")
 		if os.IsNotExist(err) {
   			errDir := os.Mkdir("output", 0755)
   			if errDir != nil {
      				log.Fatal(err)
   			}
		}
		if *payloadsPtr != "" && *payloadPtr == "" {
			 // Implement Concurrency
                	var wg sync.WaitGroup
               	 	for i := 0; i < *concurrPtr/2; i++ {
                        	wg.Add(1)
                        	go func() {
                                	// Run the scanner
                                	runWithMultiplePayload(*payloadsPtr)
                                	wg.Done()
                        	}()
                        	wg.Wait()
                	}
		}else{

			// Implement Concurrency
        		var wg sync.WaitGroup
        		for i := 0; i < *concurrPtr/2; i++ {
               	 		wg.Add(1)
                		go func() {
                        		// Run the scanner
                        		runWithSinglePayload(*payloadPtr)
                        		wg.Done()
               			}()
                		wg.Wait()
        		}
	        }
        }
}

// Print the banner
func banner() {
	

	m1 := `	
      _       
     (_)      
  _____ _ __  
 |_  / | '_ \ 
  / /| | | | |
 /___|_|_| |_|

	`
	m2 := `
May the bounties come

	`
	
	fmt.Println(Red + m1 + Cyan + m2) 
}


// Read the file containing the urls from stdin
func runWithMultiplePayload(payloads string) {

 fmt.Println(White + "[" + Blue + "~" + White + "] Searching for URL(s)")
 fmt.Println(White + "[" + Green+ "~" + White + "]" + Red + " Multiple Payloads")
 fmt.Println(White + "========================================================================================\n")

 fmt.Println("Status Code\tBytes\t\tURL")
 fmt.Println("-----------\t-----\t\t---\n")

 	
 // Create the 'NewScanner' object and print each line in the file
 scanner := bufio.NewScanner(os.Stdin)
 for scanner.Scan() {


 	// Parse the URL
  	u,err := url.Parse(scanner.Text())
   	if err != nil{
      		log.Fatal(err)
   	}
   	// Fetch the URL Values
	qs := url.Values{}

	file,err := os.Open(payloads)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()


	// Create the response file
        f,err := os.Create("output/responses.txt")
        if err != nil {
        	log.Fatal(err)
        }

        // Dump the response
        resp,err := http.Get(scanner.Text())
   	if err != nil {
      		log.Fatal(err)
       	}

   	dump,err := httputil.DumpResponse(resp, true)
   	if err != nil {
      		log.Fatal(err)
   	}
  	l,err := f.WriteString(string(dump))
   	if err != nil {
        	log.Fatal(err)
        }

	defer resp.Body.Close()	

	pL := bufio.NewScanner(file)

        for pL.Scan() {


                // Get the url paraemters and set the newvalue (payload)
                for param,vv := range u.Query() {
                        qs.Set(param, vv[0]+pL.Text())
                }

		// Url encoding the url
        	u.RawQuery = qs.Encode()
        	// Print the values
        	fmt.Printf("%s\t", resp.StatusCode)
        	fmt.Printf("%d Bytes\t", l)
        	fmt.Println(White + "[" + Green + "~" + White + "] " + White + u.String())
        }
        if err := pL.Err(); err != nil {
                log.Fatal(err)
        }

 }
 if err := scanner.Err(); err != nil {
   log.Fatal(err)
 }
}


// Read the file containing the urls from stdin
func runWithSinglePayload(payload string) {

	fmt.Println(White + "[" + Blue + "~" + White + "] Searching for URL(s)")
	fmt.Println(White + "[" + Green+ "~" + White + "] Payload: " + payload)
	fmt.Println(White + "========================================================================================\n")
		
	fmt.Println("Status Code\tBytes\t\tURL")
	fmt.Println("-----------\t-----\t\t---\n")

	// Create the 'NewScanner' object and print each line in the file
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {


		// Parse the URL
		u,err := url.Parse(scanner.Text())
		if err != nil{
			log.Fatal(err)
		}
		// Fetch the URL Values
		qs := url.Values{}
		

		// Get the url paraemters and set the newvalue (payload)
		for param,vv := range u.Query() {
			qs.Set(param, vv[0]+payload)
		}

		// Url encoding the url
		u.RawQuery = qs.Encode()
		
		// Dump the response
		resp,err := http.Get(scanner.Text())
		if err != nil {
                        log.Fatal(err)
                }

		// Create the response file
                f,err := os.Create("output/responses.txt")
                if err != nil {
                        log.Fatal(err)
                }

		dump,err := httputil.DumpResponse(resp, true)
		if err != nil {
			log.Fatal(err)
		}
		l,err := f.WriteString(string(dump))
		if err != nil {
                        log.Fatal(err)
                }

		defer resp.Body.Close()

		// Print the values
		fmt.Printf("%s\t", resp.StatusCode)
		fmt.Printf("%d Bytes\t", l)
		fmt.Println(White + "[" + Green + "~" + White + "] " + White + u.String()) 
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
