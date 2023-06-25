package phone

import (
	"regexp"
	"strings"
)

type Phone struct {
	Name string `json:"Name"`
	Code string `json:"Code"`
}

func Codes() []Phone {
	return []Phone{
		{Name: "ANDORRA", Code: "376"},
		{Name: "UNITED ARAB EMIRATES", Code: "971"},
		{Name: "AFGHANISTAN", Code: "93"},
		{Name: "ANDORRA", Code: "376"},
		{Name: "UNITED ARAB EMIRATES", Code: "971"},
		{Name: "AFGHANISTAN", Code: "93"},
		{Name: "ANTIGUA AND BARBUDA", Code: "1268"},
		{Name: "ANGUILLA", Code: "1264"},
		{Name: "ALBANIA", Code: "355"},
		{Name: "ARMENIA", Code: "374"},
		{Name: "NETHERLANDS ANTILLES", Code: "599"},
		{Name: "ANGOLA", Code: "244"},
		{Name: "ANTARCTICA", Code: "672"},
		{Name: "ARGENTINA", Code: "54"},
		{Name: "AMERICAN SAMOA", Code: "1684"},
		{Name: "AUSTRIA", Code: "43"},
		{Name: "AUSTRALIA", Code: "61"},
		{Name: "ARUBA", Code: "297"},
		{Name: "AZERBAIJAN", Code: "994"},
		{Name: "BOSNIA AND HERZEGOVINA", Code: "387"},
		{Name: "BARBADOS", Code: "1246"},
		{Name: "BANGLADESH", Code: "880"},
		{Name: "BELGIUM", Code: "32"},
		{Name: "BURKINA FASO", Code: "226"},
		{Name: "BULGARIA", Code: "359"},
		{Name: "BAHRAIN", Code: "973"},
		{Name: "BURUNDI", Code: "257"},
		{Name: "BENIN", Code: "229"},
		{Name: "SAINT BARTHELEMY", Code: "590"},
		{Name: "BERMUDA", Code: "1441"},
		{Name: "BRUNEI DARUSSALAM", Code: "673"},
		{Name: "BOLIVIA", Code: "591"},
		{Name: "BRAZIL", Code: "55"},
		{Name: "BAHAMAS", Code: "1242"},
		{Name: "BHUTAN", Code: "975"},
		{Name: "BOTSWANA", Code: "267"},
		{Name: "BELARUS", Code: "375"},
		{Name: "BELIZE", Code: "501"},
		{Name: "CANADA", Code: "1"},
		{Name: "COCOS (KEELING) ISLANDS", Code: "61"},
		{Name: "CONGO, THE DEMOCRATIC REPUBLIC OF THE", Code: "243"},
		{Name: "CENTRAL AFRICAN REPUBLIC", Code: "236"},
		{Name: "CONGO", Code: "242"},
		{Name: "SWITZERLAND", Code: "41"},
		{Name: "COTE D IVOIRE", Code: "225"},
		{Name: "COOK ISLANDS", Code: "682"},
		{Name: "CHILE", Code: "56"},
		{Name: "CAMEROON", Code: "237"},
		{Name: "CHINA", Code: "86"},
		{Name: "COLOMBIA", Code: "57"},
		{Name: "COSTA RICA", Code: "506"},
		{Name: "CUBA", Code: "53"},
		{Name: "CAPE VERDE", Code: "238"},
		{Name: "CHRISTMAS ISLAND", Code: "61"},
		{Name: "CYPRUS", Code: "357"},
		{Name: "CZECH REPUBLIC", Code: "420"},
		{Name: "GERMANY", Code: "49"},
		{Name: "DJIBOUTI", Code: "253"},
		{Name: "DENMARK", Code: "45"},
		{Name: "DOMINICA", Code: "1767"},
		{Name: "DOMINICAN REPUBLIC", Code: "1809"},
		{Name: "ALGERIA", Code: "213"},
		{Name: "ECUADOR", Code: "593"},
		{Name: "ESTONIA", Code: "372"},
		{Name: "EGYPT", Code: "20"},
		{Name: "ERITREA", Code: "291"},
		{Name: "SPAIN", Code: "34"},
		{Name: "ETHIOPIA", Code: "251"},
		{Name: "FINLAND", Code: "358"},
		{Name: "FIJI", Code: "679"},
		{Name: "FALKLAND ISLANDS (MALVINAS)", Code: "500"},
		{Name: "MICRONESIA, FEDERATED STATES OF", Code: "691"},
		{Name: "FAROE ISLANDS", Code: "298"},
		{Name: "FRANCE", Code: "33"},
		{Name: "GABON", Code: "241"},
		{Name: "UNITED KINGDOM", Code: "44"},
		{Name: "GRENADA", Code: "1473"},
		{Name: "GEORGIA", Code: "995"},
		{Name: "GHANA", Code: "233"},
		{Name: "GIBRALTAR", Code: "350"},
		{Name: "GREENLAND", Code: "299"},
		{Name: "GAMBIA", Code: "220"},
		{Name: "GUINEA", Code: "224"},
		{Name: "EQUATORIAL GUINEA", Code: "240"},
		{Name: "GREECE", Code: "30"},
		{Name: "GUATEMALA", Code: "502"},
		{Name: "GUAM", Code: "1671"},
		{Name: "GUINEA-BISSAU", Code: "245"},
		{Name: "GUYANA", Code: "592"},
		{Name: "HONG KONG", Code: "852"},
		{Name: "HONDURAS", Code: "504"},
		{Name: "CROATIA", Code: "385"},
		{Name: "HAITI", Code: "509"},
		{Name: "HUNGARY", Code: "36"},
		{Name: "INDONESIA", Code: "62"},
		{Name: "IRELAND", Code: "353"},
		{Name: "ISRAEL", Code: "972"},
		{Name: "ISLE OF MAN", Code: "44"},
		{Name: "INDIA", Code: "91"},
		{Name: "IRAQ", Code: "964"},
		{Name: "IRAN, ISLAMIC REPUBLIC OF", Code: "98"},
		{Name: "ICELAND", Code: "354"},
		{Name: "ITALY", Code: "39"},
		{Name: "JAMAICA", Code: "1876"},
		{Name: "JORDAN", Code: "962"},
		{Name: "JAPAN", Code: "81"},
		{Name: "KENYA", Code: "254"},
		{Name: "KYRGYZSTAN", Code: "996"},
		{Name: "CAMBODIA", Code: "855"},
		{Name: "KIRIBATI", Code: "686"},
		{Name: "COMOROS", Code: "269"},
		{Name: "SAINT KITTS AND NEVIS", Code: "1869"},
		{Name: "KOREA DEMOCRATIC PEOPLES REPUBLIC OF", Code: "850"},
		{Name: "KOREA REPUBLIC OF", Code: "82"},
		{Name: "KUWAIT", Code: "965"},
		{Name: "CAYMAN ISLANDS", Code: "1345"},
		{Name: "KAZAKSTAN", Code: "7"},
		{Name: "LAO PEOPLES DEMOCRATIC REPUBLIC", Code: "856"},
		{Name: "LEBANON", Code: "961"},
		{Name: "SAINT LUCIA", Code: "1758"},
		{Name: "LIECHTENSTEIN", Code: "423"},
		{Name: "SRI LANKA", Code: "94"},
		{Name: "LIBERIA", Code: "231"},
		{Name: "LESOTHO", Code: "266"},
		{Name: "LITHUANIA", Code: "370"},
		{Name: "LUXEMBOURG", Code: "352"},
		{Name: "LATVIA", Code: "371"},
		{Name: "LIBYAN ARAB JAMAHIRIYA", Code: "218"},
		{Name: "MOROCCO", Code: "212"},
		{Name: "MONACO", Code: "377"},
		{Name: "MOLDOVA, REPUBLIC OF", Code: "373"},
		{Name: "MONTENEGRO", Code: "382"},
		{Name: "SAINT MARTIN", Code: "1599"},
		{Name: "MADAGASCAR", Code: "261"},
		{Name: "MARSHALL ISLANDS", Code: "692"},
		{Name: "MACEDONIA, THE FORMER YUGOSLAV REPUBLIC OF", Code: "389"},
		{Name: "MALI", Code: "223"},
		{Name: "MYANMAR", Code: "95"},
		{Name: "MONGOLIA", Code: "976"},
		{Name: "MACAU", Code: "853"},
		{Name: "NORTHERN MARIANA ISLANDS", Code: "1670"},
		{Name: "MAURITANIA", Code: "222"},
		{Name: "MONTSERRAT", Code: "1664"},
		{Name: "MALTA", Code: "356"},
		{Name: "MAURITIUS", Code: "230"},
		{Name: "MALDIVES", Code: "960"},
		{Name: "MALAWI", Code: "265"},
		{Name: "MEXICO", Code: "52"},
		{Name: "MALAYSIA", Code: "60"},
		{Name: "MOZAMBIQUE", Code: "258"},
		{Name: "NAMIBIA", Code: "264"},
		{Name: "NEW CALEDONIA", Code: "687"},
		{Name: "NIGER", Code: "227"},
		{Name: "NIGERIA", Code: "234"},
		{Name: "NICARAGUA", Code: "505"},
		{Name: "NETHERLANDS", Code: "31"},
		{Name: "NORWAY", Code: "47"},
		{Name: "NEPAL", Code: "977"},
		{Name: "NAURU", Code: "674"},
		{Name: "NIUE", Code: "683"},
		{Name: "NEW ZEALAND", Code: "64"},
		{Name: "OMAN", Code: "968"},
		{Name: "PANAMA", Code: "507"},
		{Name: "PERU", Code: "51"},
		{Name: "FRENCH POLYNESIA", Code: "689"},
		{Name: "PAPUA NEW GUINEA", Code: "675"},
		{Name: "PHILIPPINES", Code: "63"},
		{Name: "PAKISTAN", Code: "92"},
		{Name: "POLAND", Code: "48"},
		{Name: "SAINT PIERRE AND MIQUELON", Code: "508"},
		{Name: "PITCAIRN", Code: "870"},
		{Name: "PUERTO RICO", Code: "1"},
		{Name: "PORTUGAL", Code: "351"},
		{Name: "PALAU", Code: "680"},
		{Name: "PARAGUAY", Code: "595"},
		{Name: "QATAR", Code: "974"},
		{Name: "ROMANIA", Code: "40"},
		{Name: "SERBIA", Code: "381"},
		{Name: "RUSSIAN FEDERATION", Code: "7"},
		{Name: "RWANDA", Code: "250"},
		{Name: "SAUDI ARABIA", Code: "966"},
		{Name: "SOLOMON ISLANDS", Code: "677"},
		{Name: "SEYCHELLES", Code: "248"},
		{Name: "SUDAN", Code: "249"},
		{Name: "SWEDEN", Code: "46"},
		{Name: "SINGAPORE", Code: "65"},
		{Name: "SAINT HELENA", Code: "290"},
		{Name: "SLOVENIA", Code: "386"},
		{Name: "SLOVAKIA", Code: "421"},
		{Name: "SIERRA LEONE", Code: "232"},
		{Name: "SAN MARINO", Code: "378"},
		{Name: "SENEGAL", Code: "221"},
		{Name: "SOMALIA", Code: "252"},
		{Name: "SURINAME", Code: "597"},
		{Name: "SAO TOME AND PRINCIPE", Code: "239"},
		{Name: "EL SALVADOR", Code: "503"},
		{Name: "SYRIAN ARAB REPUBLIC", Code: "963"},
		{Name: "SWAZILAND", Code: "268"},
		{Name: "TURKS AND CAICOS ISLANDS", Code: "1649"},
		{Name: "CHAD", Code: "235"},
		{Name: "TOGO", Code: "228"},
		{Name: "THAILAND", Code: "66"},
		{Name: "TAJIKISTAN", Code: "992"},
		{Name: "TOKELAU", Code: "690"},
		{Name: "TIMOR-LESTE", Code: "670"},
		{Name: "TURKMENISTAN", Code: "993"},
		{Name: "TUNISIA", Code: "216"},
		{Name: "TONGA", Code: "676"},
		{Name: "TURKEY", Code: "90"},
		{Name: "TRINIDAD AND TOBAGO", Code: "1868"},
		{Name: "TUVALU", Code: "688"},
		{Name: "TAIWAN, PROVINCE OF CHINA", Code: "886"},
		{Name: "TANZANIA, UNITED REPUBLIC OF", Code: "255"},
		{Name: "UKRAINE", Code: "380"},
		{Name: "UGANDA", Code: "256"},
		{Name: "UNITED STATES", Code: "1"},
		{Name: "URUGUAY", Code: "598"},
		{Name: "UZBEKISTAN", Code: "998"},
		{Name: "HOLY SEE (VATICAN CITY STATE)", Code: "39"},
		{Name: "SAINT VINCENT AND THE GRENADINES", Code: "1784"},
		{Name: "VENEZUELA", Code: "58"},
		{Name: "VIRGIN ISLANDS, BRITISH", Code: "1284"},
		{Name: "VIRGIN ISLANDS, U.S.", Code: "1340"},
		{Name: "VIET NAM", Code: "84"},
		{Name: "VANUATU", Code: "678"},
		{Name: "WALLIS AND FUTUNA", Code: "681"},
		{Name: "SAMOA", Code: "685"},
		{Name: "KOSOVO", Code: "381"},
		{Name: "YEMEN", Code: "967"},
		{Name: "MAYOTTE", Code: "262"},
		{Name: "SOUTH AFRICA", Code: "27"},
		{Name: "ZAMBIA", Code: "260"},
		{Name: "ZIMBABWE", Code: "263"},
	}
}

func FormatPhone(phone string) string {
	if len(phone) > 13 {
		return strings.TrimSpace(simpleFormat(phone))
	}

	if phone == "" {
		return ""
	}

	if !ValidatePhoneNo(phone) {
		return strings.TrimSpace(simpleFormat(phone))
	}

	codes := Codes()

	if strings.HasPrefix(phone, "+") {
		for _, p := range codes {
			re := regexp.MustCompile(`(\+` + p.Code + `)[1-9][0-9]{2}[0-9]{3}[0-9]{3}$`)
			if re.MatchString(phone) {
				result := "+" + regexp.MustCompile(`^\+? ?(\d{3})?[ /]?(\d{3})[ /]?(\d{3})[ /]?(\d{3})$`).ReplaceAllString(phone, "$1 $2 $3 $4")
				return strings.TrimSpace(result)
			}
		}
	} else {
		if len(phone) > 9 {
			for _, p := range codes {
				re := regexp.MustCompile(`(` + p.Code + `)[1-9][0-9]{2}[0-9]{3}[0-9]{3}$`)
				if re.MatchString(phone) {
					result := "+" + regexp.MustCompile(`^(\d{3})[ /]?(\d{3})[ /]?(\d{3})[ /]?(\d{3})$`).ReplaceAllString(phone, "$1 $2 $3 $4")
					return strings.TrimSpace(result)
				}
			}
		}

		re := regexp.MustCompile(`^[0-9]{9}$`)
		if re.MatchString(phone) {
			result := "+420" + regexp.MustCompile(`^(\d{3})[ /]?(\d{3})[ /]?(\d{3})[ /]?(\d{3})$`).ReplaceAllString(phone, "$1 $2 $3 $4")
			return strings.TrimSpace(result)
		}
	}

	return strings.TrimSpace(simpleFormat(phone))
}

func ValidatePhoneNo(phone string) bool {
	codes := Codes()

	if strings.HasPrefix(phone, "+") && len(phone) == 13 {
		for _, p := range codes {
			re := regexp.MustCompile(`(\+` + p.Code + `)[1-9][0-9]{2}[0-9]{3}[0-9]{3}$`)
			if re.MatchString(phone) {
				return true
			}
		}
	} else {
		if len(phone) > 9 {
			for _, p := range codes {
				re := regexp.MustCompile(`(` + p.Code + `)[1-9][0-9]{2}[0-9]{3}[0-9]{3}$`)
				if re.MatchString(phone) {
					return true
				}
			}
		}

		re := regexp.MustCompile(`^[0-9]{9}$`)
		if re.MatchString(phone) {
			return true
		}
	}

	return false
}

func simpleFormat(phone string) string {
	if strings.Contains(phone, "+") {
		phone = strings.ReplaceAll(phone, "+", "")
		return "+" + regexp.MustCompile(`(.+?)(.+?)(.+?)`).ReplaceAllString(phone, "$1$2$3 ")
	}

	return regexp.MustCompile(`(.+?)(.+?)(.+?)`).ReplaceAllString(phone, "$1$2$3 ")
}

func ValidatePhonePrefix(prefix string) bool {
	codes := Codes()
	for _, code := range codes {
		if prefix == "+"+code.Code {
			return true
		}
	}
	return false
}
