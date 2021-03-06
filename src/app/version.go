// simple verion
//
// @author darryl.west@ebay.com
// @created 2017-07-20 09:59:37

package app

import "fmt"

const (
	major = 18
	minor = 7
	patch = 5
)

var logo = `
 _______ __     __         ______             
|     __|__|.--|  |.-----.|      |.---.-.----.
|__     |  ||  _  ||  -__||   ---||  _  |   _|
|_______|__||_____||_____||______||___._|__|  
`

// Version - return the version number as a single string
func Version() string {
	return fmt.Sprintf("%d.%d.%d", major, minor, patch)
}

// AppLogo the application logo
func appLogo() string {
	return logo
}
