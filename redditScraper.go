package main

import (
	"fmt"
	//"errors"
	"strings"
	"strconv"
	"github.com/go-rod/rod"
	"os"
     "time"
	"github.com/TwiN/go-color"
	//"github.com/go-rod/rod/lib/input"
	"github.com/go-rod/rod/lib/launcher"


)
func redditInfo(s1 string, s2 string, s5 string){

	fmt.Print("\033[H\033[2J")
	//function for info
	println(color.Ize(color.Green, "Reddit Info"))
	println(color.Ize(color.Green, "------------------"))
	println(color.Ize(color.Green, "Total Post: " + s1))
	println(color.Ize(color.Green, "Total Members: " + s2))
	println(color.Ize(color.Green, "Online Members: " + s5))

}
func printFilters(){
	fmt.Print("\033[H\033[2J")
	println(color.Ize(color.Green, "Filter Menu"))
	println(color.Ize(color.Green, "------------------"))
	println(color.Ize(color.Green, "1.  Info"))
	println(color.Ize(color.Green, "2.  Fan Art"))
	println(color.Ize(color.Green, "3.  Merch Concepts"))
	println(color.Ize(color.Green, "4.  WEED"))
	println(color.Ize(color.Green, "5.  Talents"))
	println(color.Ize(color.Green, "6.  Memes"))
	println(color.Ize(color.Green, "7.  Bong Rips"))
	println(color.Ize(color.Green, "8.  Crypto"))
	println(color.Ize(color.Green, "9.  Quit"))

}
func goBack(){

			var input string
			println(color.Ize(color.Green, "Press b to go back!"))
			fmt.Scanln(&input)
			if(input=="b"){
			fmt.Print("\033[H\033[2J")
			redditSearch()
			}



}
func redditSearch(){
	url := launcher.New().
		Headless(true).
		Devtools(true).
		Delete("use-mock-keychain"). 
		MustLaunch()
	
	browser := rod.New().
		ControlURL(url).
		Trace(false).
		SlowMotion(1 * time.Second).
		MustConnect()
	
	defer browser.MustClose()
	page := 	browser.MustPage("https://www.reddit.com/r/StonerApeClub/")
	time.Sleep(time.Millisecond*1000)
	println(color.Ize(color.Green, "Reddit Search Engine"))
	println(color.Ize(color.Green, "------------------"))
	println(color.Ize(color.Green, "1. Search by Keywords"))
	println(color.Ize(color.Green, "2. Search by Filters"))
	println(color.Ize(color.Green, "3. Go Back"))
	println(color.Ize(color.Green, "4. Quit"))
	var userSelection string
	//scan user input
	fmt.Scanln(&userSelection)
	switch userSelection {
		case "1":
		fmt.Print("\033[H\033[2J")
		var userInput string
		print(color.Ize(color.Green, "Keywords: "))
		fmt.Scanln(&userInput)
		//search by keyword, click searchbar
	
//input user input

	page.MustElement(`input[type=search]`).MustWaitVisible().MustInput(userInput).MustPress('\r')
	time.Sleep(time.Millisecond*2000)
	totalPost := len(page.MustElements(`div[data-testid=post-container]`))
	s1 := strconv.Itoa(totalPost)
	switch {
		case totalPost<1:
		println(color.Ize(color.Green, "Found 0 Results, Retrying..."))
	page.MustElement(`input[type=search]`).MustWaitVisible().MustInput(userInput).MustPress('\r')
totalPost1 := len(page.MustElements(`div[data-testid=post-container]`))
s2 := strconv.Itoa(totalPost1)
		time.Sleep(time.Millisecond*2000)
		//page.MustElement(`div[class="_2-rGW3UtrT-pD45pofU3tx"]`).MustWaitVisible().MustClick()		
		println(color.Ize(color.Green, "Found " + s2 + " Total Post for " + userInput))
		case totalPost>=1:
	println(color.Ize(color.Green, "Found " + s1 + " Total Post for " + userInput))
			goBack()			
	}
		case "2":
			//search by filters
			printFilters()
			var userSelection2 string
			fmt.Scanln(&userSelection2)
			switch userSelection2 {
			case "1":
				//info
			page.MustElement(`a[href="/r/StonerApeClub/?f=flair_name%3A%22Info%22"]`).MustWaitVisible().MustClick()
time.Sleep(time.Millisecond*2000)
			totalInfo := len(page.MustElements(`div[data-testid=post-container]`))
			totalInfo1 := strconv.Itoa(totalInfo)
			println(color.Ize(color.Green, "Found " + totalInfo1 + " Posts with the filter Info!"))
			goBack()
			case "2":
				//Fan Art
			page.MustElement(`a[href=="/r/StonerApeClub/?f=flair_name%3A%22Fan%20Art%22"]`).MustWaitVisible().MustClick()
	time.Sleep(time.Millisecond*2000)
totalArt := len(page.MustElements(`div[data-testid=post-container]`))
			totalArt1 := strconv.Itoa(totalArt)
			println(color.Ize(color.Green, "Found " + totalArt1 + " Posts with the filter Fan-Art!"))
			goBack()
			case "3":
				//Merch Concepts
	page.MustElement(`a[href="/r/StonerApeClub/?f=flair_name%3A%22Merch%20Concepts%22"]`).MustWaitVisible().MustClick()
	time.Sleep(time.Millisecond*2000)
totalMerch := len(page.MustElements(`div[data-testid=post-container]`))
			totalMerch1 := strconv.Itoa(totalMerch)
			println(color.Ize(color.Green, "Found " + totalMerch1 + " Posts with the filter Merch Concepts!"))
			goBack()
			case "4":
				//Weed
	page.MustElement(`a[href="/r/StonerApeClub/?f=flair_name%3A%22WEED%22"]`).MustWaitVisible().MustClick()
	time.Sleep(time.Millisecond*2000)
totalWeed := len(page.MustElements(`div[data-testid=post-container]`))
			totalWeed1 := strconv.Itoa(totalWeed)
			println(color.Ize(color.Green, "Found " + totalWeed1 + " Posts with the filter WEED!"))
			goBack()
			case "5":
				//Talents
	page.MustElement(`a[href="/r/StonerApeClub/?f=flair_name%3A%22Talents%22"]`).MustWaitVisible().MustClick()
	time.Sleep(time.Millisecond*2000)
totalTalents := len(page.MustElements(`div[data-testid=post-container]`))
			totalTalents1 := strconv.Itoa(totalTalents)
			println(color.Ize(color.Green, "Found " + totalTalents1 + " Posts with the filter Talents!"))
			goBack()
			case "6":
				//Memes
	page.MustElement(`a[href="/r/StonerApeClub/?f=flair_name%3A%22Memes%22"]`).MustWaitVisible().MustClick()
	time.Sleep(time.Millisecond*2000)
totalMemes := len(page.MustElements(`div[data-testid=post-container]`))
			totalMemes1 := strconv.Itoa(totalMemes)
			println(color.Ize(color.Green, "Found " + totalMemes1 + " Posts with the filter Memes!"))
			goBack()
			case "7":
				//Bong Rips
	page.MustElement(`a[href="/r/StonerApeClub/?f=flair_name%3A%22Massive%20Rips%22"]`).MustWaitVisible().MustClick()
	time.Sleep(time.Millisecond*2000)
totalRips := len(page.MustElements(`div[data-testid=post-container]`))
			totalRips1 := strconv.Itoa(totalRips)
			println(color.Ize(color.Green, "Found " + totalRips1 + " Posts with the filter Bong Rips!"))
			goBack()
			case "8":
				//crypto
	page.MustElement(`a[href="/r/StonerApeClub/?f=flair_name%3A%22Crypto%22"]`).MustWaitVisible().MustClick()
	time.Sleep(time.Millisecond*2000)
totalCrypto := len(page.MustElements(`div[data-testid=post-container]`))
			totalCrypto1 := strconv.Itoa(totalCrypto)
			println(color.Ize(color.Green, "Found " + totalCrypto1 + " Posts with the filter Crypto!"))
			goBack()
			case "9":
				os.Exit(3)
			}
		case "3":
			//go back to reddit menu
			redditScraper()
		case "4":
			//quit
			os.Exit(3)
	}




}
func redditScraper() {
fmt.Print("\033[H\033[2J")
//implement go-rod
url := launcher.New().
		Headless(true).
		Devtools(true).
		Delete("use-mock-keychain"). 
		MustLaunch()

browser := rod.New().
		ControlURL(url).
		Trace(false).
		//SlowMotion(1 * time.Second).
		MustConnect()
//wait for page to be visible..
defer browser.MustClose()
page := browser.MustPage("https://www.reddit.com/r/StonerApeClub/")
page.Timeout(100*time.Second)
//page.WaitLoad()
time.Sleep(time.Millisecond*1000)


println(color.Ize(color.Green, "Accessing Reddit.."))
//first lets click menu
page.MustElement(`button[id="LayoutSwitch--picker"]`).MustWaitVisible().MustClick()
//now click compact
time.Sleep(time.Millisecond*500)
mouse := page.Mouse
mouse.Move(750.0,350.0,20)
time.Sleep(time.Millisecond*100)
mouse.MustDown("left")
time.Sleep(100*time.Millisecond)
mouse.MustUp("left")
time.Sleep(time.Millisecond*1000)
mouse.Scroll(500,5000,5)
time.Sleep(time.Millisecond*1000)
mouse.Scroll(500,10000,8)
time.Sleep(time.Millisecond*1000)
mouse.Scroll(500,10000,8)
time.Sleep(time.Millisecond*1000)
mouse.Scroll(500,10000,8)
time.Sleep(time.Millisecond*1000)
mouse.Scroll(500,10000,8)
//<div class="_3XFx6CfPlg-4Usgxm0gK8R">696</div>
var s2 string
var s3 string
s2= page.MustElement(`div[class="_3XFx6CfPlg-4Usgxm0gK8R"]`).MustWaitVisible().MustText()
//"_3_HlHJ56dAfStT19Jgl1bF
s3=page.MustElement(`div[class="_3_HlHJ56dAfStT19Jgl1bF"]`).MustWaitVisible().MustText()
//fmt.Println("Found", len(page.Timeout(5000 * time.Second).MustElements(`div[data-testid=post-container]`)), "Total Reddit Post!")
totalPost := len(page.MustElements(`div[data-testid=post-container]`))
//convert
s1 := strconv.Itoa(totalPost)
time.Sleep(time.Millisecond*1000)
println(color.Ize(color.Green, "Welcome to the Reddit Menu!"))
println(color.Ize(color.Green, "1.   Random Reddit Post!"))
println(color.Ize(color.Green, "2.   Search the Reddit!"))
//call search function, allow applying filters
println(color.Ize(color.Green, "3.   Reddit Info!"))
println(color.Ize(color.Green, "4.   Quit!"))
var redditSelection string
fmt.Scanln(&redditSelection)
fmt.Print("\033[H\033[2J")
//replace white space in s3
//s4:=strings.ReplaceAll(s3, " ", "")
s4:=strings.Replace(s3, "\n", ".", -1)
s5:=strings.Replace(s4, ".", " ", -1)

switch redditSelection {
case "1":
	redditScraper()
case "2":
	//reddit search engine, pull up menu first
	redditSearch()
case "3":
	//call reddit info function
	redditInfo(s1,s2,s5)
case "4":
	//close
	os.Exit(3)
}

}
func printMenu() {
fmt.Print("\033[H\033[2J")
fmt.Println("Choose one of the options below!")
fmt.Println("1.  Search the SAC Reddit!")
fmt.Println("2.  Pull up SAC Twitter")
fmt.Println("3.  Random SAC Sneak Peek!")
fmt.Println("4.  Close!")
var userSelection string
fmt.Scanln(&userSelection)
fmt.Print("\033[H\033[2J")
//switch case for menu selection
switch userSelection {
case "1":
	redditScraper()
case "4":
	//close
	os.Exit(3)
}
}

func main() {
var userInput string
fmt.Println("Welcome!")
fmt.Println("Press 1 for menu or q for quit!")
fmt.Scanln(&userInput)
fmt.Print("\033[H\033[2J")
if(userInput=="1"){
//print menu
printMenu()
}
if(userInput=="q"){
//quit
fmt.Println("See ya!")
os.Exit(3)
}
}