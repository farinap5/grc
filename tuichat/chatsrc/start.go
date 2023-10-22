package chatsrc

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/marcusolsson/tui-go"
)

var User string = "none"
var TargetId string = "none"

var History *tui.Box
var Ui tui.UI

func Cmds(cmd string) string {
	cmdS := strings.Split(cmd," ")
	if cmdS[0][1:] == "connect" {
		if len(cmdS) != 3 {
			return "Error! Type: !connect 0.0.0.0:8080 username"
		}
		return InitConnection(cmdS[1],
		strings.Split(cmd," ")[2])

	} else if string(cmd[1:]) == "disconnect" {
		CloseSocket()
		return "Connection Closed"

	} else if cmdS[0][1:] == "target" {
		if len(cmdS) != 2 {
			return "Error! Type: !target username"
		}
		TargetId = cmdS[1]
		return "Talking to "+TargetId
	} else {
		return "Not a command"
	}
}

func testReceive(history *tui.Box, ui tui.UI) {
	for i:=0; i < 10; i++ {
		time.Sleep(1000 *time.Millisecond)
		PrintMessage("any","test")
	}
}

func PrintMessage(who ,text string) {
	Ui.Update(func() {
		History.Append(tui.NewHBox(
			tui.NewLabel(time.Now().Format("15:04")),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", who))),
			tui.NewLabel(text),
			tui.NewSpacer(),
		))
	})
}

func createChatHystory() (*tui.Box,*tui.Box) {
	history := tui.NewVBox()

	historyScroll := tui.NewScrollArea(history)
	historyScroll.SetAutoscrollToBottom(true)
	historyBox := tui.NewVBox(historyScroll)
	historyBox.SetBorder(true)

	chat := tui.NewVBox(historyBox)
	chat.SetSizePolicy(tui.Expanding, tui.Expanding)
	return chat,history
}

func createInput(history *tui.Box) *tui.Box {
	input := tui.NewEntry()
	input.SetFocused(true)
	input.SetSizePolicy(tui.Expanding, tui.Maximum)

	input.OnSubmit(func(e *tui.Entry) {
		msg := ""
		if strings.HasPrefix(e.Text(),"!") {
			msg = Cmds(e.Text())
		} else {
			if TargetId != "none" {
				Send(TargetId,e.Text()) // Send via socket
				msg = e.Text()
			} else {
				msg = "Select a target"
			}
		}

		history.Append(tui.NewHBox(
			tui.NewLabel(time.Now().Format("15:04")),
			tui.NewPadder(1, 0, tui.NewLabel(fmt.Sprintf("<%s>", User))),
			tui.NewLabel(msg),
			tui.NewSpacer(),
		))
		input.SetText("")
	})

	inputBox := tui.NewHBox(input)
	inputBox.SetBorder(true)
	inputBox.SetSizePolicy(tui.Expanding, tui.Maximum)
	return inputBox
}

func Start() {
	chat,history := createChatHystory()
	input := createInput(history)
	root := tui.NewVBox(chat, input)

	ui, err := tui.New(root)
	if err != nil {
		log.Fatal(err)
	}
	Ui = ui
	History = history
	
	ui.SetKeybinding("Esc", func() { ui.Quit() })
	if err := ui.Run(); err != nil {
		log.Fatal(err)
	}
}
