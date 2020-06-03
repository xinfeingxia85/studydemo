package main

import (
	"fmt"
	"gopkg.in/ldap.v3"
	"log"
	"os"
)

var conn *ldap.Conn

func main() {
	var err error

	//connection
	conn, err = ldap.Dial("tcp", "jinan.eden.3jk.ink:38389")
	if err != nil {
		log.Println("Connect error: ", err)
		os.Exit(1)
	}

	//认证
	err = conn.Bind("cn=admin,dc=sanjieke,dc=com", "HjoG6Oymsdw2")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println("Authentication sucess!!")
	}

	//搜索用户
	//searchUser(conn, "fangkaiming")
	searchGroups(conn)
}

func searchUser(conn *ldap.Conn, userName string) {
	//搜索
	searchRequest := ldap.NewSearchRequest(
		"ou=users,dc=sanjieke,dc=com",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		"(&(uid="+userName+"))",
		[]string{"cn", "displayName", "uid", "mail", "mobile"},
		nil,
	)

	searchResults, _ := conn.Search(searchRequest)
	fmt.Println(searchResults.Entries[0].DN)
	fmt.Println(searchResults.Entries[0].GetAttributeValue("cn"))
	fmt.Println(searchResults.Entries[0].GetAttributeValue("displayName"))
	fmt.Println(searchResults.Entries[0].GetAttributeValue("uid"))
	fmt.Println(searchResults.Entries[0].GetAttributeValue("mail"))
	fmt.Println(searchResults.Entries[0].GetAttributeValue("mobile"))
}

func searchGroups(conn *ldap.Conn) {
	//搜索
	searchRequest := ldap.NewSearchRequest(
		"ou=三节课,ou=users,dc=sanjieke,dc=com",
		ldap.ScopeWholeSubtree,
		ldap.NeverDerefAliases,
		0,
		0,
		false,
		//"(&(objectClass=inetOrgPerson))",
		"(&(objectClass=organizationalUnit))",
		[]string{"ou", "objectClass"},
		nil,
	)

	searchResults, _ := conn.Search(searchRequest)
	if len(searchResults.Entries) > 0 {
		for _, item := range searchResults.Entries {
			//member := item.GetAttributeValues("uniqueMember")
			//cn := item.GetAttributeValues("cn")
			//if len(member) > 0 {
			//	fmt.Println(cn)
			//}
			fmt.Println(item.DN)

		}
	}
	//fmt.Println(searchResults.Entries[0].DN)
	//fmt.Println(searchResults.Entries[0].GetAttributeValue("cn"))
	//fmt.Println(searchResults.Entries[0].GetAttributeValue("displayName"))
	//fmt.Println(searchResults.Entries[0].GetAttributeValue("uid"))
	//fmt.Println(searchResults.Entries[0].GetAttributeValue("mail"))
	//fmt.Println(searchResults.Entries[0].GetAttributeValue("mobile"))
}
