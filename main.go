package main

import (
	eth "Safu/eth"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

var CA []string
var GroupID []int64
var Activated []string
var ethlib eth.EthLib

func main() {
	bot, err := tgbotapi.NewBotAPI("6185386343:AAGWt3VqV5h4yZqA6YWzc6uOnOZ9_1rmdJw")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates, err := bot.GetUpdatesChan(u)
	//offsetMsg := regexp.MustCompile("(?i)0x[a-fA-F0-9]{40}") https://bsc-dataseed1.binance.org/https://rpc.ankr.com/bsc
	ethlib = eth.InitEthlib(context.Background(), "https://bsc-dataseed1.binance.org/", "https://mainnet.infura.io/v3/56bb53b84c2e439fa277c9e6522044fe", "royal base fatigue this cry donate shoot impulse region square field toilet")
	for update := range updates {
		if len(CA) > 30 {
			CA = []string{}
			GroupID = []int64{}
			fmt.Println("Reseting CA Store", CA, "Group ID", GroupID)
		}
		if update.Message == nil { // ignore non-Message updates
			continue
		}
		if update.Message.Chat.Type == "supergroup" || update.Message.Chat.Type == "group" {
			done := make(chan string)
			go func(update tgbotapi.Update, donX chan string) {
				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				var command string
				command = update.Message.Text
				var arg string
				arg = ""
				fmt.Println("leng of arry", len(strings.Split(command, " ")))
				if len(strings.Split(command, " ")) > 1 {
					arg = strings.Split(command, " ")[1]
					command = strings.Split(command, " ")[0]

				}

				replyText, id := processCommandGroup(command, update.Message.From.UserName, arg, update.Message.Chat.ID, update.Message.MessageID, *bot)
				fmt.Println("Bot Chat ID:-", update.Message.Chat.ID)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
				msg.ParseMode = tgbotapi.ModeHTML

				if id != 0 {
					msg.ReplyToMessageID = id
				}

				bot.Send(msg)
				donX <- "Done"
			}(update, done)
			go func() {
				var Xnm int = 0
				for Xnm < 1 {

					select {
					case <-done:
						Xnm = 5

					}
					if Xnm > 1 {
						fmt.Println("========Number of go rountin Completer", runtime.NumGoroutine())
					}
				}
			}()
			//	<-done
			//	close(done)
			//continue
		} else {
			done := make(chan string)
			go func(update tgbotapi.Update, donX chan string) {

				log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)
				var command string
				command = update.Message.Text
				var arg string
				arg = ""
				fmt.Println("====================>leng of arry", len(strings.Split(command, " ")))
				if len(strings.Split(command, " ")) > 1 {
					arg = strings.Split(command, " ")[1]
					command = strings.Split(command, " ")[0]

				}

				replyText, id := processCommand(command, update.Message.From.UserName, arg, update.Message.Chat.ID, update.Message.MessageID, *bot)
				fmt.Println("Bot Chat ID:-", update.Message.Chat.ID)
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, replyText)
				msg.ParseMode = tgbotapi.ModeHTML

				if id != 0 {
					msg.ReplyToMessageID = id
				}

				bot.Send(msg)
				donX <- "Done"
			}(update, done)
			go func() {
				var Xnm int = 0
				for Xnm < 1 {

					select {
					case <-done:
						Xnm = 5

					}
					if Xnm > 1 {
						fmt.Println("========Number of go rountin Completer", runtime.NumGoroutine())
					}
				}
			}()
			//<-done
			//close(done)
			//Wg.Wait()
			fmt.Println("========End P", runtime.NumGoroutine())
			//tgbotapi.Message()
		}
	}
}

type Data struct {
	UserID     string `json:"userID"`
	AccessID   int    `json:"accessID"`
	AccessHash string `json:"accessHash"`
	BOTID      int64  `json:"BotID"`
}

func AddToArry(ca string, id int64) {
	notavailable := true
	for i := 0; i < len(CA); i++ {
		if CA[i] == ca {
			notavailable = false
			break
		}
	}
	if notavailable {
		CA = append(CA, ca)
		GroupID = append(GroupID, id)

		//fmt.Println("The Cas", CA, "The Groups", GroupID)
	} else {
		//fmt.Println("Added")
	}
}

func retriveUserD(username string) (int, string, int64) {
	filez, err := os.Open("usersDetails.json")
	if err != nil {
		fmt.Println("Error opening file:", err)

	}
	defer filez.Close()

	// Read existing data from the file
	bytes, err := ioutil.ReadAll(filez)
	if err != nil {
		fmt.Println("Error reading file:", err)

	}

	var users []Data
	err = json.Unmarshal(bytes, &users)
	if err != nil {
		fmt.Println("Error unmarshaling JSON:", err)

	}

	// Loop through the users and check if the username is available
	//username := "johndoe"
	var accessID int
	var accessHash string
	var botID int64
	accessID = 0
	botID = 0
	accessHash = ""

	for _, user := range users {
		//fmt.Println(user.UserID)
		if user.UserID == username {
			accessHash = user.AccessHash
			accessID = user.AccessID
			botID = user.BOTID
			break
		}
	}
	return accessID, accessHash, botID
}

type Message struct {
	Start  string `xml:"start"`
	Report string `xml:"report"`
}

func getMessage(from string) string {

	// Open the XML file
	file, err := os.Open("defaultMsg.xml")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return ""
	}
	defer file.Close()

	// Read the XML data into a Person struct
	var message Message
	decoder := xml.NewDecoder(file)
	err = decoder.Decode(&message)
	if err != nil {
		fmt.Println("Error decoding XML:", err)
		return ""
	}

	switch strings.ToLower(from) {
	case "start":
		return message.Start
	case "report":
		return message.Report
	default:
		return ""
	}
	// Print the data
	//fmt.Printf("Name: %s\n", message.start)

}

// 0x5589feefe8058936f7a50ff743dadfe99ef89f3b 0xB133E2F95E37ece88dD354434ba6f14aC322C195 0xB133E2F95E37ece88dD354434ba6f14aC322C195
func processCommand(command string, username string, arg string, botid int64, msgsID int, bot tgbotapi.BotAPI) (string, int) {

	offsetMsg := regexp.MustCompile("0x[a-fA-F0-9]{40}")
	evmAdd := offsetMsg.FindString(command)
	//fmt.Println("Evm Address", evmAdd)
	switch strings.ToLower(command) {
	case "/start":

		msg := fmt.Sprintf(getMessage("start"), username)
		return msg, 0
	case strings.ToLower(evmAdd):

		var msg string
		ntK := make(chan string)
		ntKTx := make(chan string)
		var wg sync.WaitGroup
		wg.Add(1)
		fmt.Println("========Number of go rountin Before Start", runtime.NumGoroutine())
		go ethlib.GetChain(evmAdd, &wg, ntK, ntKTx)
		fmt.Println("========Number of go rountin During Start", runtime.NumGoroutine())
		networkz := ""
		network := "None" //<-ntK
		report := make(chan string)
		funcReport := make(chan string)
		sentReport := make(chan string)
		nnkT := make(chan string)
		wg.Add(1)
		go func(nnkT chan string, wgx *sync.WaitGroup) {
			defer wgx.Done()
			var rptIst string = ""
			var rptScnd string = ""
			var inDex int = 0
			for inDex < 1 { //rptIst == "" || rptScnd == "" {
				select {
				case netWek := <-ntK:
					networkz = netWek
					if networkz != "None" {
						nnn := strings.Split(networkz, " ")
						fmt.Println("========Number of go rountin During 2 Start", runtime.NumGoroutine())
						if len(nnn) > 5 {
							msg = fmt.Sprintf(getMessage("report"), "<b>"+nnn[8]+"</b>")
							fmt.Println("======>>Network", networkz, msg, rptScnd, rptIst)
						}
						if rptIst != "" {

							break
						} else {
							rptIst = networkz
							msgx := tgbotapi.NewMessage(botid, msg)
							msgx.ParseMode = tgbotapi.ModeHTML

							bot.Send(msgx)
						}
					} else {
						rptIst = "None"
					}
					// close(ntK)
				case networkFul := <-ntKTx:
					//network = networkFul

					for rptIst == "" {

					}
					inDex = 5
					nnkT <- networkFul
					close(ntKTx)

					fmt.Println("========Number of go rountin During 3 Start", runtime.NumGoroutine())
					return

				}
			}
		}(nnkT, &wg)
		network = <-nnkT
		close(nnkT)

		wg.Wait()
		fmt.Println("========Number of go rountin End Start", runtime.NumGoroutine())
		if network != "None" {
			var innerWg sync.WaitGroup
			innerWg.Add(1)

			go ethlib.Simulate(network, evmAdd, report, funcReport, sentReport, &innerWg)
			//innerWg.Wait()

			nnn := strings.Split(network, " ")
			var MsgchatID int = -1
			//go ethlib.Simulate(network, evmAdd, report, funcReport)
			innerWg.Add(1)
			lstMsg := make(chan string)
			go func(network string, address string, botID int64, lstMsg chan string, wWg *sync.WaitGroup) {
				defer wWg.Done()
				xaddress := address
				var complete int = 0

				var firstMsg string
				var SecondMsg string = ""
				var indz int = 0
				for indz < 1 {
					select {
					case x := <-report:
						//return x, 0
						fmt.Println("========Number of go rountin3", runtime.NumGoroutine())
						firstMsg = x
						msg := tgbotapi.NewMessage(botID, x)

						msg.ParseMode = tgbotapi.ModeHTML
						//MsgchatID = int64(msg.ReplyToMessageID)
						response, err := bot.Send(msg)
						if err != nil {
							log.Fatal(err)
						}
						//bot.Send(msg)
						//break
						MsgchatID = response.MessageID
						if complete > 1 {
							break
						} else {
							complete += 1
						}

					case y := <-funcReport:
						fmt.Println("FuncRPT")
						for MsgchatID < 0 {

						}
						fmt.Println("========Number of go rountin3", runtime.NumGoroutine())
						SecondMsg = y
						var scanNetwork string
						var DexChart string
						switch network {
						case "BSC":
							scanNetwork = "https://bscscan.com/address/" + xaddress
							DexChart = "https://poocoin.app/tokens/" + xaddress
						case "ETH":
							scanNetwork = "https://etherscan.io//address/" + xaddress
							DexChart = "https://www.dextools.io/app/en/ether/pair-explorer/" + xaddress
						}

						inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
							// Create the first row with two buttons
							tgbotapi.NewInlineKeyboardRow(
								tgbotapi.NewInlineKeyboardButtonURL("🔎 "+network+" Scan", scanNetwork),
								tgbotapi.NewInlineKeyboardButtonURL("📈 Chart", DexChart),
								tgbotapi.NewInlineKeyboardButtonURL("🥷 Channel", "https://t.me/stealthLaunchE"),
							),
						)
						msg := tgbotapi.NewEditMessageText(botID, MsgchatID, firstMsg+y)
						msg.ReplyMarkup = &inlineKeyboard
						msg.ParseMode = tgbotapi.ModeHTML

						response, err := bot.Send(msg)
						if err != nil {
							log.Fatal(err)
						}
						MsgchatID = response.MessageID
						if complete > 1 {
							break
						} else {
							complete += 1
						}

					case z := <-sentReport:

						for SecondMsg == "" {

						}
						indz = 5
						fmt.Println("SentimentX", indz)
						SecondMsg = strings.Replace(SecondMsg, "<code><b>Always DYOR. Auto rugcheckers cannot detect all scams.</b></code>", "", -1)
						SecondMsg = strings.Replace(SecondMsg, "\n\n<code><b>⏱ Performing Sentiment Analysis. Please Wait...(2 mins)</b></code>", "", -1)
						//SecondMsg = SecondMsg[:len(SecondMsg)-2]
						fmt.Println("Sentiment", SecondMsg[:len(SecondMsg)-2])
						msgx := firstMsg + SecondMsg + z
						fmt.Println("Final Msg", msg)
						fmt.Println("========Number of go rountin3", runtime.NumGoroutine())

						var scanNetwork string
						var DexChart string
						switch network {
						case "BSC":
							scanNetwork = "https://bscscan.com/address/" + xaddress
							DexChart = "https://poocoin.app/tokens/" + xaddress
						case "ETH":
							scanNetwork = "https://etherscan.io//address/" + xaddress
							DexChart = "https://www.dextools.io/app/en/ether/pair-explorer/" + xaddress
						}

						inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
							// Create the first row with two buttons
							tgbotapi.NewInlineKeyboardRow(
								tgbotapi.NewInlineKeyboardButtonURL("🔎 "+network+" Scan", scanNetwork),
								tgbotapi.NewInlineKeyboardButtonURL("📈 Chart", DexChart),
								tgbotapi.NewInlineKeyboardButtonURL("🥷 Channel", "https://t.me/stealthLaunchE"),
							),
						)
						msg := tgbotapi.NewEditMessageText(botID, MsgchatID, msgx)
						msg.ReplyMarkup = &inlineKeyboard
						msg.ParseMode = tgbotapi.ModeHTML

						_, err := bot.Send(msg)
						if err != nil {
							log.Fatal(err)
						}
						lstMsg <- "DOne"
						break

					}
					if indz > 1 {
						fmt.Println("Fmt Indxz", indz)
					}

				}
				fmt.Println("+++Process Ended")
				close(sentReport)
				close(funcReport)
				close(report)

			}(nnn[8], evmAdd, botid, lstMsg, &innerWg)
			lastMsg := <-lstMsg
			close(lstMsg)

			innerWg.Wait()

			fmt.Println("========Gourotin COmplent Alpha", runtime.NumGoroutine(), "MMSS", lastMsg)
			return "", 0

		} else {

			// Loop through the updates to find the message with the text you're interested in
			// replace with the text you're interested in
			msgID := msgsID
			fmt.Println(msgID)

			msg = "<i>❗️ This address might be outside our network range</i>"
			return msg, msgID
		}
		//}()
	case "/hash":
		accessID := 5486942816
		person := Data{UserID: username, AccessID: accessID, AccessHash: arg, BOTID: botid}
		filez, err := os.Open("usersDetails.json")
		if err != nil {
			fmt.Println("Error opening file:", err)

		}
		defer filez.Close()

		// Read existing data from the file
		bytes, err := ioutil.ReadAll(filez)
		if err != nil {
			fmt.Println("Error reading file:", err)

		}

		var users []Data
		err = json.Unmarshal(bytes, &users)
		if err != nil {
			fmt.Println("Error unmarshaling JSON:", err)

		}

		// Loop through the users and check if the username is available
		//username := "johndoe"
		notavailable := true
		for _, user := range users {
			fmt.Println(user.UserID)
			if user.UserID == username {
				notavailable = false
				break
			}
		}
		if notavailable {
			file, err := os.Open("usersDetails.json")
			if err != nil {
				fmt.Println("Error opening file:", err)

			}
			defer file.Close()
			var data []Data
			decoder := json.NewDecoder(file)
			err = decoder.Decode(&data)
			if err != nil {
				fmt.Println("Error decoding JSON:", err)

			}
			//fmt.Print(decoder)
			data = append(data, person)

			// Open the file for writing
			file, err = os.OpenFile("usersDetails.json", os.O_WRONLY, 0644)
			if err != nil {
				fmt.Println("Error opening file:", err)

			}
			defer file.Close()

			// Encode the combined data into a JSON-encoded byte slice
			dataz, err := json.Marshal(data)
			if err != nil {
				fmt.Println("Error marshaling JSON:", err)

			}

			// Write the encoded data to the file

			_, err = file.Write(dataz)
			if err != nil {
				fmt.Println("Error writing to file:", err)

			}

			fmt.Println("Data appended to file successfully.")
			//arg := ""

			if arg == "" {
				return "🚫 Please Enter a valid Access Hash", 0
			}
			if username == "" {
				return "🚫 Please Set a username before using this option", 0
			}

			msg := fmt.Sprintf(`
	           🔹 SAVED 🔹
		____________________________
		🔹 AccessID:     %d
		🔹Access Hash:  %s
		🔹 UserID:      %s
		____________________________
		`, accessID, arg, username)
			return msg, 0
		} else {
			msg := fmt.Sprintf(`
	           🔹 Available 🔹
		____________________________
		🔹 The AccessID:     %d
		🔹The Access Hash:  %s
		🔹 Your UserID:      %s
		____________________________
		`, accessID, arg, username)
			return msg, 0
		}
	default:
		//fmt.Println("Evm Address", evmAdd, strings.ToLower(command))
		msgID := msgsID
		return "<i>Sorry, I didn't understand that command. Type /help to see a list of available commands.</i>", msgID
	}

}
func processCommandGroup(command string, username string, arg string, botid int64, msgsID int, bot tgbotapi.BotAPI) (string, int) {

	fmt.Println("Evm Address", strings.ToLower(command))
	switch strings.ToLower(command) {
	case "/safree":

		offsetMsg := regexp.MustCompile("0x[a-fA-F0-9]{40}")
		evmAdd := offsetMsg.FindString(arg)
		fmt.Println("Evm Address", strings.ToLower(arg), strings.ToLower(evmAdd))
		switch strings.ToLower(arg) {
		case strings.ToLower(evmAdd):

			var msg string
			ntK := make(chan string)
			ntKTx := make(chan string)
			var wg sync.WaitGroup
			wg.Add(1)
			fmt.Println("========Number of go rountin Before Start", runtime.NumGoroutine())
			go ethlib.GetChain(evmAdd, &wg, ntK, ntKTx)
			fmt.Println("========Number of go rountin During Start", runtime.NumGoroutine())
			networkz := ""
			network := "None" //<-ntK
			report := make(chan string)
			funcReport := make(chan string)
			sentReport := make(chan string)
			nnkT := make(chan string)
			wg.Add(1)
			go func(nnkT chan string, wgx *sync.WaitGroup) {
				defer wgx.Done()
				var rptIst string = ""
				var rptScnd string = ""
				var inDex int = 0
				for inDex < 1 { //rptIst == "" || rptScnd == "" {
					select {
					case netWek := <-ntK:
						networkz = netWek

						nnn := strings.Split(networkz, " ")
						fmt.Println("========Number of go rountin During 2 Start", runtime.NumGoroutine())
						if len(nnn) > 5 {
							msg = fmt.Sprintf(getMessage("report"), "<b>"+nnn[8]+"</b>")
							fmt.Println("======>>Network", networkz, msg, rptScnd, rptIst)
						}
						if rptIst != "" {

							break
						} else {
							rptIst = networkz
							msgx := tgbotapi.NewMessage(botid, msg)
							msgx.ParseMode = tgbotapi.ModeHTML

							bot.Send(msgx)
						}
						// close(ntK)
					case networkFul := <-ntKTx:
						//network = networkFul

						for rptIst == "" {

						}
						inDex = 5
						nnkT <- networkFul
						close(ntKTx)

						fmt.Println("========Number of go rountin During 3 Start", runtime.NumGoroutine())
						return

					}
				}
			}(nnkT, &wg)
			network = <-nnkT
			close(nnkT)

			wg.Wait()
			fmt.Println("========Number of go rountin End Start", runtime.NumGoroutine())
			if network != "None" {
				var innerWg sync.WaitGroup
				innerWg.Add(1)

				go ethlib.Simulate(network, evmAdd, report, funcReport, sentReport, &innerWg)
				//innerWg.Wait()

				nnn := strings.Split(network, " ")
				var MsgchatID int = -1
				//go ethlib.Simulate(network, evmAdd, report, funcReport)
				innerWg.Add(1)
				lstMsg := make(chan string)
				go func(network string, address string, botID int64, lstMsg chan string, wWg *sync.WaitGroup) {
					defer wWg.Done()
					xaddress := address
					var complete int = 0

					var firstMsg string
					var SecondMsg string = ""
					var indz int = 0
					for indz < 1 {
						select {
						case x := <-report:
							//return x, 0
							fmt.Println("========Number of go rountin3", runtime.NumGoroutine())
							firstMsg = x
							msg := tgbotapi.NewMessage(botID, x)

							msg.ParseMode = tgbotapi.ModeHTML
							//MsgchatID = int64(msg.ReplyToMessageID)
							response, err := bot.Send(msg)
							if err != nil {
								log.Fatal(err)
							}
							//bot.Send(msg)
							//break
							MsgchatID = response.MessageID
							if complete > 1 {
								break
							} else {
								complete += 1
							}

						case y := <-funcReport:
							fmt.Println("FuncRPT")
							for MsgchatID < 0 {

							}
							fmt.Println("========Number of go rountin3", runtime.NumGoroutine())
							SecondMsg = y
							var scanNetwork string
							var DexChart string
							switch network {
							case "BSC":
								scanNetwork = "https://bscscan.com/address/" + xaddress
								DexChart = "https://poocoin.app/tokens/" + xaddress
							case "ETH":
								scanNetwork = "https://etherscan.io//address/" + xaddress
								DexChart = "https://www.dextools.io/app/en/ether/pair-explorer/" + xaddress
							}

							inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
								// Create the first row with two buttons
								tgbotapi.NewInlineKeyboardRow(
									tgbotapi.NewInlineKeyboardButtonURL("🔎 "+network+" Scan", scanNetwork),
									tgbotapi.NewInlineKeyboardButtonURL("📈 Chart", DexChart),
									tgbotapi.NewInlineKeyboardButtonURL("🥷 Channel", "https://t.me/stealthLaunchE"),
								),
							)
							msg := tgbotapi.NewEditMessageText(botID, MsgchatID, firstMsg+y)
							msg.ReplyMarkup = &inlineKeyboard
							msg.ParseMode = tgbotapi.ModeHTML

							response, err := bot.Send(msg)
							if err != nil {
								log.Fatal(err)
							}
							MsgchatID = response.MessageID
							if complete > 1 {
								break
							} else {
								complete += 1
							}

						case z := <-sentReport:

							for SecondMsg == "" {

							}
							indz = 5
							fmt.Println("SentimentX", indz)
							SecondMsg = strings.Replace(SecondMsg, "<code><b>Always DYOR. Auto rugcheckers cannot detect all scams.</b></code>", "", -1)
							SecondMsg = strings.Replace(SecondMsg, "\n\n<code><b>⏱ Performing Sentiment Analysis. Please Wait...(2 mins)</b></code>", "", -1)
							//SecondMsg = SecondMsg[:len(SecondMsg)-2]
							fmt.Println("Sentiment", SecondMsg[:len(SecondMsg)-2])
							msgx := firstMsg + SecondMsg + z
							fmt.Println("Final Msg", msg)
							fmt.Println("========Number of go rountin3", runtime.NumGoroutine())

							var scanNetwork string
							var DexChart string
							switch network {
							case "BSC":
								scanNetwork = "https://bscscan.com/address/" + xaddress
								DexChart = "https://poocoin.app/tokens/" + xaddress
							case "ETH":
								scanNetwork = "https://etherscan.io//address/" + xaddress
								DexChart = "https://www.dextools.io/app/en/ether/pair-explorer/" + xaddress
							}

							inlineKeyboard := tgbotapi.NewInlineKeyboardMarkup(
								// Create the first row with two buttons
								tgbotapi.NewInlineKeyboardRow(
									tgbotapi.NewInlineKeyboardButtonURL("🔎 "+network+" Scan", scanNetwork),
									tgbotapi.NewInlineKeyboardButtonURL("📈 Chart", DexChart),
									tgbotapi.NewInlineKeyboardButtonURL("🥷 Channel", "https://t.me/stealthLaunchE"),
								),
							)
							msg := tgbotapi.NewEditMessageText(botID, MsgchatID, msgx)
							msg.ReplyMarkup = &inlineKeyboard
							msg.ParseMode = tgbotapi.ModeHTML

							_, err := bot.Send(msg)
							if err != nil {
								log.Fatal(err)
							}
							lstMsg <- "DOne"
							break

						}
						if indz > 1 {
							fmt.Println("Fmt Indxz", indz)
						}

					}
					fmt.Println("+++Process Ended")
					close(sentReport)
					close(funcReport)
					close(report)

				}(nnn[8], evmAdd, botid, lstMsg, &innerWg)
				lastMsg := <-lstMsg
				close(lstMsg)

				innerWg.Wait()

				fmt.Println("========Gourotin COmplent Alpha", runtime.NumGoroutine(), "MMSS", lastMsg)
				return "", 0

			} else {

				// Loop through the updates to find the message with the text you're interested in
				// replace with the text you're interested in
				msgID := msgsID
				fmt.Println(msgID)

				msg = "<i>❗️ This address might be outside our network range</i>"
				return msg, msgID
			}
		//}()
		default:

			msgID := msgsID
			return "<i>❗️ Unknow Argument Please retype</i>", msgID
		}
		//}()

	default:
		//fmt.Println("Evm Address", evmAdd, strings.ToLower(command))
		msgID := msgsID
		return "<i>🔍 Let me scan Your token for free !!!. I do a great job at that</i>", msgID
	}
}
