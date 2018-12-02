package monga

import (
	"fmt"
	"log"
	"strings"
	"time"

	"gopkg.in/mgo.v2"
)

type Config struct {
	Name  string
	Hosts []string
}

type Client struct {
	*mgo.Session
}

func (c *Client) C(name string) *mgo.Collection {
	return c.Session.DB("").C(name)
}

func MustNew(cfg Config) *Client {
	return &Client{
		Session: dial(cfg),
	}
}

func dial(cfg Config) *mgo.Session {
	hosts := strings.Join(cfg.Hosts, ",")
	url := fmt.Sprintf("%s/%s", hosts, cfg.Name)

	s, err := mgo.Dial(url)
	if err != nil {
		log.Fatalf("monga: failed to dial %q: %v", url, err)
	}

	s.SetSocketTimeout(5 * time.Second)
	s.SetSafe(&mgo.Safe{})

	return s

}
