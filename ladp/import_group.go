package main

import (
	"fmt"
	"gopkg.in/ldap.v3"
	"log"
	"strings"
	"time"
)

type LDAP_CONFIG struct {
	baseDN          string
	basePWD         string
	host            string
	groupDN         string
	groupFilter     string
	groupAttributes []string
}

var ldapConfig = LDAP_CONFIG{
	"cn=admin,dc=sanjieke,dc=com",
	"HjoG6Oymsdw2",
	"jinan.eden.3jk.ink:38389",
	"ou=三节课,ou=users,dc=sanjieke,dc=com",
	"(&(objectClass=organizationalUnit))",
	//"(&(objectClass=organizationalUnit))",
	[]string{"ou"},
}

//var groups map[string]interface{}

func main() {
	conn := connect()
	fetchGroup(conn, ldapConfig.groupDN)
}

func connect() *ldap.Conn {
	ldap.DefaultTimeout = 15 * time.Second
	conn, err := ldap.Dial("tcp", ldapConfig.host)
	if err != nil {
		log.Fatal(err)
	}

	//authenticate user and password
	err = conn.Bind(ldapConfig.baseDN, ldapConfig.basePWD)
	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func fetchGroup(conn *ldap.Conn, groupDN string) {
	searchRequest := ldap.NewSearchRequest(
		groupDN,
		ldap.ScopeSingleLevel,
		ldap.DerefFindingBaseObj,
		0,
		0,
		false,
		ldapConfig.groupFilter,
		ldapConfig.groupAttributes,
		nil,
	)

	searchResults, err := conn.Search(searchRequest)
	if err != nil {
		log.Println(err)
	}

	if len(searchResults.Entries) > 0 {
		for _, entry := range searchResults.Entries {
			ou := entry.GetAttributeValue("ou")
			fmt.Println(strings.Repeat("-", strings.Count(groupDN, "ou")*2), ou)
			fetchGroup(conn, "ou="+ou+","+groupDN)
		}
	}
}