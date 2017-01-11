package gnip

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/eapache/channels"
)

const (
	streamURL      = "https://gnip-stream.twitter.com/stream/powertrack/accounts/%s/publishers/twitter/%s.json"
	rulesURL       = "https://gnip-api.twitter.com/rules/powertrack/accounts/%s/publishers/twitter/%s.json"
	replayURL      = "https://stream.gnip.com:443/accounts/%s/publishers/twitter/replay/track/%s.json"
	replayRulesURL = "https://api.gnip.com:443/accounts/%s/publishers/twitter/replay/track/%s/rules.json"
	fullArchiveURL = "https://data-api.twitter.com/search/fullarchive/accounts/%s/%s"
)

/*
Client ...
*/
type Client struct {
	client   *http.Client
	user     string
	password string
	account  string
	ch       *channels.InfiniteChannel
	stop     chan bool
	active   bool
}

/*
NewClient ...
*/
func NewClient(user string, password string, account string) *Client {
	c := &Client{
		&http.Client{
			Transport: &http.Transport{
				Dial:            (&net.Dialer{}).Dial,
				IdleConnTimeout: 30 * time.Second,
			},
		},
		user,
		password,
		account,
		channels.NewInfiniteChannel(),
		make(chan bool),
		false,
	}
	return c
}

/*
ConnectPowertrack ...
*/
func (client *Client) ConnectPowertrack(streamLabel string) (<-chan interface{}, error) {
	powertrackStream := fmt.Sprintf(streamURL, client.account, streamLabel)
	req, err := http.NewRequest("GET", powertrackStream, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Connection", "keep-alive")
	req.SetBasicAuth(client.user, client.password)

	go func() {
		resp, _ := client.client.Do(req)
		defer resp.Body.Close()
		client.active = true
		reader := bufio.NewReader(resp.Body)
		for {
			select {
			case <-client.stop:
				log.Println("STOP RECEIVED")
				client.active = false
				client.ch.Close()
				break
			default:
				if client.active {
					line, _ := reader.ReadBytes('\r')
					line = bytes.TrimSpace(line)

					if line != nil && string(line[:]) != "" {
						client.ch.In() <- line
					}
				}
			}
		}
	}()

	return client.ch.Out(), nil
}

/*
StopPowertrack ...
*/
func (client *Client) StopPowertrack() {
	client.stop <- true
}

/*
IsActive ...
*/
func (client *Client) IsActive() bool {
	return client.active
}
