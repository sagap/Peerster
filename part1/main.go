package main

import (
"flag"
"github.com/sagap/Peerster/part1/gossiper2_backend"
)

func main(){
	uiPort := flag.Int("UIPort",12345,"user interface port(for clients)")
	gossipPort := flag.String("gossipAddr","127.0.0.1:10002","gossip protocol port")
	nodeName := flag.String("name","nodeA", "name for node")
	peers := flag.String("peers","","underscore separated peers")
	rtimer := flag.Int("rtimer",60,"rtimer flag")
	guiPort := flag.Int("GUIPort",0,"Gui Port")
	flag.Parse()
	gossiper := gossiper2_backend.NewGossiper(*gossipPort, *nodeName, *peers, *rtimer, *guiPort)
	gossiper.ServeClients(*uiPort)
}

