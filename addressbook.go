package i2pcontrol

import "strconv"

// AddAddressBookEntry adds an entry to the address book. bookType is one of "local", "router", "published", "private"
func AddAddressBookEntry(hostname, destination, bookType string) (string, error) {
	retpre, err := Call("AddressBook", map[string]interface{}{
		"Hostname":    hostname,
		"Destination": destination,
		"Type":        bookType,
	})
	if err != nil {
		return "", err
	}
	result := retpre["message"].(string)
	return result, nil
}

// RemoveAddressBookEntry removes an entry from the address book. bookType is one of "local", "router", "published", "private"
func RemoveAddressBookEntry(hostname, destination, bookType string) (string, error) {
	retpre, err := Call("AddressBook", map[string]interface{}{
		"Hostname":    hostname,
		"Destination": destination,
		"Type":        bookType,
		"Delete":      "",
	})
	if err != nil {
		return "", err
	}
	result := retpre["message"].(string)
	return result, nil
}

// EditAddressBookSubscription edits the subscriptions for the address book.
// Hostnames is a list of hostnames to subscribe to. It will replace the existing subscriptions.
func EditAddressBookSubscription(hostnames []string) (string, error) {
	retpre, err := Call("AddressBook", map[string]interface{}{
		"SetSubscriptions": hostnames,
	})
	if err != nil {
		return "", err
	}
	result := retpre["message"].(string)

	return result, nil
}

// EditConfigFile edits the config file.
func EditConfigFile(etags, lastFetched, lastModified, localAddressbook, logPath, namingService, privateAddressbook, proxyHost string, proxyPort int, publishedAddressbook, routerAddressbook string, shouldPublish bool, subscriptions string, updateDelay int, updateDirect bool) (string, error) {
	retpre, err := Call("AddressBook", map[string]interface{}{
		"SetConfig": map[string]interface{}{
			"etags": etags, "last_fetched": lastFetched, "last_modified": lastModified,
			"local_addressbook": localAddressbook, "log": logPath, "naming_service": namingService,
			"private_addressbook": privateAddressbook, "proxy_host": proxyHost, "proxy_port": strconv.Itoa(proxyPort),
			"published_addressbook": publishedAddressbook, "router_addressbook": routerAddressbook,
			"should_publish": strconv.FormatBool(shouldPublish), "subscriptions": subscriptions,
			"update_delay": strconv.Itoa(updateDelay), "update_direct": strconv.FormatBool(updateDirect),
		},
	})
	if err != nil {
		return "", err
	}
	result := retpre["message"].(string)
	return result, nil
}
