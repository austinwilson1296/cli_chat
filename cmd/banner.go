package cmd

import (
	"fmt"
)

func PrintBanner() {
	asciiArt := `
 ## ##    ## ##    ## ##   ####               ## ##   ###  ##    ##     #### ##  
##   ##  ##   ##  ##   ##   ##               ##   ##   ##  ##     ##    # ## ##  
##       ##   ##  ##   ##   ##               ##        ##  ##   ## ##     ##     
##       ##   ##  ##   ##   ##               ##        ## ###   ##  ##    ##     
##       ##   ##  ##   ##   ##               ##        ##  ##   ## ###    ##     
##   ##  ##   ##  ##   ##   ##  ##           ##   ##   ##  ##   ##  ##    ##     
 ## ##    ## ##    ## ##   ### ###            ## ##   ###  ##  ###  ##   ####    `
	fmt.Println(asciiArt)
}