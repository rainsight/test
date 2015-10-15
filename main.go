package main

import (
	"fmt"

	// "github.com/gogap/wechat/mp"
	// "github.com/gogap/wechat/mp/jssdk"
	"github.com/chanxuehong/wechat/mp"
	"github.com/chanxuehong/wechat/mp/jssdk"
)

func main() {
	var accessTokenServer = mp.NewDefaultAccessTokenServer("wx8dd685b36012ecb6", "8660ab2e951c442658467f2d0f2da440", nil)
	var mpClient = mp.NewClient(accessTokenServer, nil)
	var ticketServer = jssdk.NewDefaultTicketServer(mpClient)

	for i := 0; i < 2; i++ {
		ticket, err := ticketServer.Ticket()
		fmt.Println("ticket:", ticket, err)
	}

	for i := 0; i < 2; i++ {
		ticket, err := ticketServer.TicketRefresh()
		fmt.Println("ticket refresh:", ticket, err)
	}
}
