package handlers

import (
	"fmt"
	"time"
	"github.com/jlaffaye/ftp"
)

func openFTP() {
	//connect to ftp server
	c, err := ftp.DialTimeout("ftp.avinnovz.com:21",5*time.Second)
	if err == nil {
		//login to ftp server
		err = c.Login("admin@avinnovz.com", "avinnovz@1234")
		if err == nil {
			//get directories listing
			directories , err := c.List("/TBoxStations")
			if err == nil {
				for _,d := range directories {
					if d.Type == 1 && d.Name != "." && d.Name != ".." {
						
					}
				}
			} else {
				fmt.Println("failed to retrieved listing");
			}

			//retrieve csv file
			// r, err := c.Retr("/TBoxStations")
			// if err == nil {
			// 	//read csv file
			// 	buf, err := ioutil.ReadAll(r)
			// 	if err == nil {
			// 		fmt.Println("reading : ", string(buf))
			// 	} else {
			// 		panic(fmt.Sprintf("failed to read file ---> %s",err))
			// 	}
			// 	r.Close()
			// } else {
			// 	panic(fmt.Sprintf("failed to retrieve file ---> %s",err))
			// }
			// panic(fmt.Sprintf("failed to store file ---> %s",err))
		} else {
			panic(fmt.Sprintf("failed to login ---> %s",err))
		}
	} else {
		panic(fmt.Sprintf("failed to connect in ftp server ---> %s",err))
	}
}