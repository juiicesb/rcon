package rcon

import "log"

type Squad struct {
	conn Conn
}

func (s *Squad) ShowCurrentMap() {
	resp, err := s.conn.Execute("ShowCurrentMap")
	if err != nil {
		log.Fatalln("Command failed", err)
	}
	log.Println(resp)
}
