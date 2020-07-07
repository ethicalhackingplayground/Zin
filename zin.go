//
// Payload Injector
//
// Coder: zoid

package main

import (
"strings"
"sync"
"io/ioutil"
"flag"
"log"
"bufio"
"os"
"fmt"
"net/url"
"net/http"
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
	var c int
	flag.IntVar(&c,"c", 20, "the concurrency")
	var payload string
	flag.StringVar(&payload,"p", "", "the payload to be used")
	var payloads string
	flag.StringVar(&payloads,"pL", "", "the list of payloads to be used")
	var statusCode int
	flag.IntVar(&statusCode,"s", 200, "filter by status codes")
	var pattern string
	flag.StringVar(&pattern, "g", "", "match the response with some string")
	var paths string
	flag.StringVar(&paths,"paths", "false", "are they just hosts ending in '/'")
	// Parse the arguments
	flag.Parse()
	if (payload == "" && payloads == "") {
		flag.PrintDefaults()
		return	
	}else{
        

		if payloads != "" && payload == "" {
			fmt.Println(White + "[" + Blue + "~" + White + "] Searching for URL(s)")
      			fmt.Println(White + "[" + Green+ "~" + White + "]" + Red + " Multiple Payloads")
        		fmt.Println(White + "[" + Green+ "~" + White + "]" + Red + " Match: " + Green+pattern)
        		fmt.Println(White + "========================================================================================\n")

        		fmt.Println("\nStatus Code\tBytes\t\tURL")
        		fmt.Println("-----------\t-----\t\t---\n")

			 // Implement Concurrency
                	var wg sync.WaitGroup
               	 	for i := 0; i < c; i++ {
                        	wg.Add(1)
                        	go func() {
                                	// Run the scanner
                                	runWithMultiplePayload(payloads, statusCode, pattern, paths)
                                	wg.Done()
                        	}()
                        	wg.Wait()
                	}
		}else{

			fmt.Println(White + "[" + Blue + "~" + White + "] Searching for URL(s)")
        	 	fmt.Println(White + "[" + Green+ "~" + White + "]" + Red + " Payload: " + Cyan + payload)
        		fmt.Println(White + "[" + Green+ "~" + White + "]" + Red + " Match: " + Green+pattern)
        		fmt.Println(White + "========================================================================================\n")

        		fmt.Println("\nStatus Code\tBytes\t\tURL")
        		fmt.Println("-----------\t-----\t\t---\n")

			// Implement Concurrency
        		var wg sync.WaitGroup
        		for i := 0; i < c; i++ {
               	 		wg.Add(1)
                		go func() {
                        		// Run the scanner
                        		runWithSinglePayload(payload, statusCode, pattern, paths)
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
func runWithMultiplePayload(payloads string,  status int, grep string, paths string) {

 	
 	// Create the 'NewScanner' object and print each line in the file
 	scanner := bufio.NewScanner(os.Stdin)
 	file,err := os.Open(payloads)
 	client := http.Client{}
 	if err != nil {
 		log.Fatal(err)
 	}
 	for scanner.Scan() {

                pL := bufio.NewScanner(file)
                for pL.Scan() {
                        payload:=pL.Text()

 		        // Parse the URL
	  	        u,err := url.Parse(scanner.Text())
   		        if err != nil {
      			        continue
   		        }
   		        // Fetch the URL Values
		        qs := url.Values{}

			
			if paths == "true" {	

				// Create a new Request
               		 	req,err := http.NewRequest("GET", u.String()+"/"+payload, nil)
                		if err != nil {
                       			continue
                		}
	
       		         	resp,err:=client.Do(req)
               		 	if err != nil {
                       		 	continue
                		}

                		bytes,err := ioutil.ReadAll(resp.Body)
                		if err != nil {
                       			continue
                		}
	
	                	bodyStr:=string(bytes)
	                	if (strings.ContainsAny(bodyStr, grep)) {
       	              		   	if resp.StatusCode == status {

                               			// Print the values
                                		fmt.Printf("%s\t", resp.StatusCode)
                                		fmt.Printf("%d Bytes\t", len(bodyStr))
                                		fmt.Println(White + "[" + Green + "~" + White + "] " + White + u.String())
						continue
                	       	 	}
                		}			

			}else{


				for param,_ :=range  u.Query() {
					qs.Set(param, payload)
				}
				u.RawQuery=qs.Encode()

	
        			// Create a new Request
        			req,err := http.NewRequest("GET", u.String(), nil)
   				if err != nil {
					continue
	       			}		
		
				resp,err:=client.Do(req)
				if err != nil {
					continue
				}	

				bytes,err := ioutil.ReadAll(resp.Body)
				if err != nil {
					continue	
				}
			
				bodyStr:=string(bytes)
				if (strings.ContainsAny(bodyStr, grep)) {     	
					if resp.StatusCode == status {		
	
	          	      		        // Print the values
       		 		                fmt.Printf("%s\t", resp.StatusCode)
                	       	 		fmt.Printf("%d Bytes\t", len(bodyStr))
                        			fmt.Println(White + "[" + Green + "~" + White + "] " + White + u.String())
                			}
				}	
			}	
		}
 	}
}


// Read the file containing the urls from stdin
func runWithSinglePayload(payload string, status int, grep string, paths string) {


        // Create the 'NewScanner' object and print each line in the file
        scanner := bufio.NewScanner(os.Stdin)
        client := http.Client{}
        for scanner.Scan() {


                // Parse the URL
                u,err := url.Parse(scanner.Text())
                if err != nil {
                        continue
                }
                // Fetch the URL Values
                qs := url.Values{}


                if paths == "true" {

                                // Create a new Request
                                req,err := http.NewRequest("GET", u.String()+"/"+payload, nil)
                                if err != nil {
                                        continue
                                }

                                resp,err:=client.Do(req)
                                if err != nil {
                                        continue
                                }

                                bytes,err := ioutil.ReadAll(resp.Body)
                                if err != nil {
                                        continue
                                }

                                bodyStr:=string(bytes)
                                if (strings.ContainsAny(grep, bodyStr)) {
                                        if resp.StatusCode == status {

                                                // Print the values
                                                fmt.Printf("%s\t", resp.StatusCode)
                                                fmt.Printf("%d Bytes\t", len(bodyStr))
                                                fmt.Println(White + "[" + Green + "~" + White + "] " + White + u.String())
                                                continue
                                        }
                                }

                }else{


                                for param,_ :=range  u.Query() {
                                        qs.Set(param, payload)
                                }
                                u.RawQuery=qs.Encode()


                                // Create a new Request
                                req,err := http.NewRequest("GET", u.String(), nil)
                                if err != nil {
                                        continue
                                }

                                resp,err:=client.Do(req)
                                if err != nil {
                                        continue
                                }

                                bytes,err := ioutil.ReadAll(resp.Body)
                                if err != nil {
                                        continue
                                }

                                bodyStr:=string(bytes)
                                if (strings.ContainsAny(bodyStr, grep)) {
                                        if resp.StatusCode == status {

                                                // Print the values
                                                fmt.Printf("%s\t", resp.StatusCode)
                                                fmt.Printf("%d Bytes\t", len(bodyStr))
                                                fmt.Println(White + "[" + Green + "~" + White + "] " + White + u.String())
                                        }
				}
			
		}
	}
}

