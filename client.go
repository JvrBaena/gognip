package gnip

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"encoding/json"

	"io/ioutil"

	"github.com/JvrBaena/gognip/types"
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

	go client.processResponse(req)

	return client.ch.Out(), nil
}

func (client *Client) processResponse(req *http.Request) {
	resp, err := client.client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		log.Println("Error in Request")
		client.active = false
		client.ch.Close()
		return
	}

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

/*
AddRule ...
*/
func (client *Client) AddRule(streamLabel string, rule *types.Rule) (*types.RuleRequestResponse, error) {
	rulesEndpoint := fmt.Sprintf(rulesURL, client.account, streamLabel)
	return client.postRules(rulesEndpoint, []*types.Rule{rule})
}

/*
AddRules ...
*/
func (client *Client) AddRules(streamLabel string, rules []*types.Rule) (*types.RuleRequestResponse, error) {
	rulesEndpoint := fmt.Sprintf(rulesURL, client.account, streamLabel)
	return client.postRules(rulesEndpoint, rules)
}

/*
RemoveRule ...
*/
func (client *Client) RemoveRule(streamLabel string, rule *types.Rule) (*types.RuleRequestResponse, error) {
	rulesEndpoint := fmt.Sprintf(rulesURL, client.account, streamLabel)
	rulesEndpoint += "?_method=delete"
	return client.postRules(rulesEndpoint, []*types.Rule{rule})
}

/*
RemoveRules ...
*/
func (client *Client) RemoveRules(streamLabel string, rules []*types.Rule) (*types.RuleRequestResponse, error) {
	rulesEndpoint := fmt.Sprintf(rulesURL, client.account, streamLabel)
	rulesEndpoint += "?_method=delete"
	return client.postRules(rulesEndpoint, rules)
}

/*
GetRule ...
*/
func (client *Client) GetRule(streamLabel string, rule *types.Rule) (*types.RuleRequestResponse, error) {
	rulesEndpoint := fmt.Sprintf(rulesURL, client.account, streamLabel)
	rulesEndpoint += "?_method=get"
	return client.postRules(rulesEndpoint, []*types.Rule{rule})
}

func (client *Client) postRules(endpoint string, rules []*types.Rule) (*types.RuleRequestResponse, error) {
	body := &types.RuleRequest{
		Rules: rules,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "json")
	req.SetBasicAuth(client.user, client.password)

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}

	reqResponse := types.RuleRequestResponse{}
	jsonResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		gnipError := types.APIRequestError{}
		json.Unmarshal(jsonResponse, &gnipError)
		return nil, &gnipError
	}

	err = json.Unmarshal(jsonResponse, &reqResponse)
	if err != nil {
		return nil, err
	}

	return &reqResponse, nil
}

/*
GetRules ...
*/
func (client *Client) GetRules(streamLabel string) (*types.RuleRequestResponse, error) {
	rulesEndpoint := fmt.Sprintf(rulesURL, client.account, streamLabel)
	req, err := http.NewRequest("GET", rulesEndpoint, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "json")
	req.SetBasicAuth(client.user, client.password)

	resp, err := client.client.Do(req)
	if err != nil {
		return nil, err
	}

	reqResponse := types.RuleRequestResponse{}
	jsonResponse, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		gnipError := types.APIRequestError{}
		json.Unmarshal(jsonResponse, &gnipError)
		return nil, &gnipError
	}

	err = json.Unmarshal(jsonResponse, &reqResponse)
	if err != nil {
		return nil, err
	}

	return &reqResponse, nil
}
