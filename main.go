package main

import (
	"fmt"
	"strings"
)

var ad = `123 Main Street St. Louisville OH 43071, 432 Main Long Road St. Louisville OH 43071,786 High Street Pollocksville NY 56432,
  54 Holy Grail Street Niagara Town ZP 32908, 3200 Main Rd. Bern AE 56210,1 Gordon St. Atlanta RE 13000,
  10 Pussy Cat Rd. Chicago EX 34342, 10 Gordon St. Atlanta RE 13000, 58 Gordon Road Atlanta RE 13000,
  22 Tokyo Av. Tedmondville SW 43098, 674 Paris bd. Abbeville AA 45521, 10 Surta Alley Goodtown GG 30654,
  45 Holy Grail Al. Niagara Town ZP 32908, 320 Main Al. Bern AE 56210, 14 Gordon Park Atlanta RE 13000,
  100 Pussy Cat Rd. Chicago EX 34342, 2 Gordon St. Atlanta RE 13000, 5 Gordon Road Atlanta RE 13000,
  2200 Tokyo Av. Tedmondville SW 43098, 67 Paris St. Abbeville AA 45521, 11 Surta Avenue Goodtown GG 30654,
  45 Holy Grail Al. Niagara Town ZP 32918, 320 Main Al. Bern AE 56215, 14 Gordon Park Atlanta RE 13200,
  100 Pussy Cat Rd. Chicago EX 34345, 2 Gordon St. Atlanta RE 13222, 5 Gordon Road Atlanta RE 13001,
  2200 Tokyo Av. Tedmondville SW 43198, 67 Paris St. Abbeville AA 45522, 11 Surta Avenue Goodville GG 30655,
  2222 Tokyo Av. Tedmondville SW 43198, 670 Paris St. Abbeville AA 45522, 114 Surta Avenue Goodville GG 30655,
  2 Holy Grail Street Niagara Town ZP 32908, 3 Main Rd. Bern AE 56210, 77 Gordon St. Atlanta RE 13000,
  100 Pussy Cat Rd. Chicago OH 13201`

func main() {
	fmt.Println(Travel(ad, "AA 45522"))
	fmt.Println(Travel(ad, "OH 430"))
	fmt.Println(Travel(ad, "NY 56432"))
	fmt.Println(Travel(ad, "ZP 32908"))
	fmt.Println(Travel(ad, "EX 34345"))
	fmt.Println(Travel(ad, "OH 43071"))
	fmt.Println(Travel(ad, "RE 13000"))
	fmt.Println(Travel(ad, "SW 43098"))
	fmt.Println(Travel(ad, "GG 30654"))
	fmt.Println(Travel(ad, "AE 56210"))
	fmt.Println(Travel(ad, "ZP 32918"))
	fmt.Println(Travel(ad, "SW 43198"))
	fmt.Println(Travel(ad, "GG 30655"))
	fmt.Println(Travel(ad, "YY 45098"))
	fmt.Println(Travel(ad, "NY 56432"))

}
func Travel(r, zipcode string) string {
	list := strings.Split(r, ",")
	resultList := make([]string, 0)
	var salesman string
	var finalstreet string
	var finalHouseNumber string
	var zipCodeArr []string
	var zipcodeInAd string
	for _, e := range list {
		if strings.Contains(e, zipcode) {
			zipCodeArr = strings.Split(e, " ")
			zipcodeInAd = zipCodeArr[len(zipCodeArr)-2] + " " + zipCodeArr[len(zipCodeArr)-1]
			if strings.EqualFold(zipcodeInAd, zipcode) {
				resultList = append(resultList, e)
			} else {
				return zipcode + ":/"
			}
		}
	}
	if len(resultList) > 1 {
		for _, address := range resultList {
			streetAndTown := address[0:strings.Index(address, zipCodeArr[len(zipCodeArr)-2])]
			streetAndTown = strings.TrimSpace(streetAndTown)

			indexOfHouseNumber := strings.Index(streetAndTown, " ")
			finalHouseNumber += streetAndTown[0:indexOfHouseNumber] + ","
			finalstreet += strings.TrimSpace(streetAndTown[indexOfHouseNumber:len(streetAndTown)]) + ","

		}
	} else if len(resultList) == 1 {
		str := strings.Index(resultList[0], zipCodeArr[len(zipCodeArr)-1])
		streetAndTown := strings.TrimSpace(resultList[0][0:str])
		streetAndTown = strings.Replace(streetAndTown, zipCodeArr[len(zipCodeArr)-2], "", 1)
		streetAndTown = strings.Replace(streetAndTown, zipCodeArr[len(zipCodeArr)-1], "", 1)

		indexOfHouseNumber := strings.Index(streetAndTown, " ")
		finalHouseNumber += streetAndTown[0:indexOfHouseNumber] + ","
		finalstreet += strings.TrimSpace(streetAndTown[indexOfHouseNumber:len(streetAndTown)-1]) + ","
	} else {
		return zipcode + ":/"
	}
	finalHouseNumber = finalHouseNumber[0 : len(finalHouseNumber)-1]
	salesman = zipcode + ":" + strings.TrimSpace(finalstreet[0:len(finalstreet)-1]) + "/" + finalHouseNumber[0:len(finalHouseNumber)]
	return salesman
}
