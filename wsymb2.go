package smhi

var Wsymb2 = map[int][2]string{
	1:  {"Klart", "Det är molnfritt eller bara några obetydliga molntussar. Det innebär soligt väder på dagen och stjärnklart på natten."},
	2:  {"Lätt molnighet", "En stor del av himlen kan vara täckt av molnslöjor, men de är i så fall så tunna att de inte skymmer solen nämnvärt. På natten kan man se åtminstone de ljusaste stjärnorna genom molnslöjorna. Om det är fråga om tjockare moln så täcker de högst en fjärdedel av himlen."},
	3:  {"Halvklart", "Omkring halva himlen är täckt men det kan vara mer än så om molnen bara är tunna. På sommarhalvåret varierar molnigheten en hel del så att solens skyms till och ifrån. Då handlar det ofta om stackmoln och det kan ofta beskrivas som växlande molnighet."},
	4:  {"Molnigt", "Det är mycket moln på himlen och solen kanske bara skymtar fram i någon glugg. Symbolen kan innebära att stackmolnen håller på att växa till under dagen och det kan leda till någon regnskur så småningom. Symbolen förekommer ibland även när ett molntäcke lättar eller när molnigheten ökar."},
	5:  {"Mycket moln", "Himlen är heltäckt men det är ändå ganska tunna moln som släpper igenom en del sol."},
	6:  {"Mulet", "Himlen är helt täckt med tjocka moln och ofta är himlen helt grå. Det kan möjligen komma någon droppe regn eller på vintern någon snöflinga."},
	7:  {"Dimma", "Dimmsymbolen visar när det är dålig sikt under 1000 meter och det kan vara i samband med både mulna och klara vädersituationer. I det senare fallet handlar det om dimbankar. Dimmsymbolen förekommer enbart när det är uppehåll."},
	8:  {"Lätt regnskur", "Mycket moln och lätt regn. Med lätt nederbörd menas regnmängder på 0,5 mm eller mindre på en timme.  Vanligen beskriver symbolen skurar och på sommarhalvåret innebär det skurmoln/bymoln som bildats ur stackmoln som växt och tornat upp sig. Symbolen kan också beskriva ett förlopp med sol följt av regn eller regn som slutar och följs av sol."},
	9:  {"Regnskur", "Mycket moln och regn eller regnskurar. Symbolen avser måttlig nederbörd och med det menas regnmängder på 0,5-4 mm på en timme.  Vanligen beskriver symbolen skurar under sommarhalvåret och det innebär skurmoln/bymoln som bildats ur stackmoln som växt och tornat upp sig. Symbolen kan också beskriva ett förlopp med sol följt av regn eller regn som slutar och följs av sol."},
	10: {"Krafig regnskur", "Mycket moln och regn eller regnskurar. Med kraftig nederbörd menas regnmängder på mer än 4 mm på en timme.  Vanligen beskriver symbolen skurar under sommarhalvåret och det innebär skurmoln/bymoln som bildats ur stackmoln som växt och tornat upp sig. Åska kan då även förekomma. Symbolen kan också beskriva ett förlopp med sol följt av regn eller regn som slutar och följs av sol."},
	11: {"Åskskur", "Mycket moln och regnskurar med stor risk för åska. I samband med åska är ofta regnskurarna kraftiga och det kan förekomma hagel. Symbolen beskriver också ett förlopp med sol följt av åskregn eller åskregn som slutar och följs av sol."},
	12: {"Lätt by av regn och snö", "Mycket moln och byar av regn eller snö blandat. Med lätt nederbörd menas mängder i smält form på 0,5 mm eller mindre på en timme.  Symbolen beskriver också ett förlopp med sol följt av snöblandat regn eller snöblandat regn som slutar och följs av sol."},
	13: {"By av regn och snö", "Mycket moln och byar av regn eller snö blandat. Symbolen avser måttlig nederbörd och med det menas mängder i smält form på 0,5-4 mm på en timme. Symbolen beskriver också ett förlopp med sol följt av snöblandat regn eller snöblandat regn som slutar och följs av sol."},
	14: {"Kraftig by av regn och snö", "Mycket moln och byar av regn eller snö blandat. Med kraftig nederbörd menas mängder i smält form på mer än 4 mm på en timme. Symbolen beskriver också ett förlopp med sol följt av snöblandat regn eller snöblandat regn som slutar och följs av sol."},
	15: {"Lätt snöby", "Mycket moln och snöbyar eller tidvis snöfall. Med lätt nederbörd menas mängder i smält form på 0,5 mm eller mindre på en timme.  Symbolen beskriver också ett förlopp med sol följt av snöfall eller snöfall som slutar och följs av sol."},
	16: {"Snöby", "Mycket moln och snöbyar eller tidvis snöfall. Symbolen avser måttlig nederbörd och med det menas mängder i smält form på 0,5-4 mm på en timme. Symbolen beskriver också ett förlopp med sol följt av snöfall eller snöfall som slutar och följs av sol."},
	17: {"Kraftig snöby", "Mycket moln och snöbyar. Med kraftig nederbörd menas mängder i smält form på mer än 4 mm på en timme.  Symbolen beskriver också ett förlopp med sol följt av snöfall eller snöfall som slutar och följs av sol."},
	18: {"Lätt regn", "Mulet och lätt regn, ihållande eller tidvis. Med lätt nederbörd menas regnmängder på 0,5 mm eller mindre på en timme."},
	19: {"Regn", "Mulet och regn eller regnskurar. Symbolen avser måttlig nederbörd och med det menas regnmängder på 0,5-4 mm på en timme."},
	20: {"Kraftigt regn", "Mulet och regn eller regnskurar. Med kraftig nederbörd menas regnmängder på mer än 4 mm på en timme.  På sommaren kan det i samband med kraftigt regn även vara en viss åskrisk."},
	21: {"Åska", "Himlen är helt täckt av moln och det kommer regn eller skurar med stor risk för åska. I samband med åska är ofta regnskurarna kraftiga och det kan förekomma hagel"},
	22: {"Lätt snöblandat regn", "Himlen är helt täckt av moln och det kommer regn eller snö blandat, ihållande eller med avbrott. Med lätt nederbörd menas mängder i smält form på 0,5 mm eller mindre på en timme."},
	23: {"Snöblandat regn", "Himlen är helt täckt av moln och det kommer regn eller snö blandat, ihållande eller med avbrott. Symbolen avser måttlig nederbörd och med det menas mängder i smält form på 0,5-4 mm på en timme."},
	24: {"Kraftigt snöblandat regn", "Himlen är helt täckt av moln och det kommer regn eller snö blandat, ihållande eller med avbrott. Med kraftig nederbörd menas mängder i smält form på mer än 4 mm på en timme."},
	25: {"Lätt snöfall", "Himlen är helt täckt av moln och det kommer snö. Snöfallet kan vara ihållande eller med avbrott.  Med lätt nederbörd menas mängder i smält form på 0,5 mm eller mindre på en timme. Det motsvarar mindre än 1 cm snö"},
	26: {"Snöfall", "Himlen är helt täckt av moln och det kommer snö. Snöfallet kan vara ihållande eller med avbrott. Symbolen avser måttlig nederbörd och med det menas mängder i smält form som motsvarar 0,5-4 mm på en timme. En generell tumregel är att varje mm motsvarar 1 cm snö, men det kan vara mer än så om det är riktigt kallt."},
	27: {"Ymnigt snöfall", "Himlen är helt täckt av moln och det kommer snö. Snöfallet kan vara ihållande eller med avbrott.  Med ymnigt snöfall menas mängder som i smält form motsvarar mer än 4 mm på en timme.  En generell tumregel är att varje mm motsvarar 1 cm snö, men det kan vara mer än så om det är riktigt kallt."},
}

var Pcat = map[int]string{
	0: "Ingen nederbörd", // No precipitation"
	1: "Snö",             //Snow
	2: "Snö och reg",     //Snow and rain
	3: "Regn",            // Rain
	4: "Duggregn",        // Drizzle
	5: "Hagel",           // Freezing rain
	6: "Hagel-regn",      //Freezing drizzle
}
