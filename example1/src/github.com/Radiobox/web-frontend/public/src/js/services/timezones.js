define(['./module'], function (services) {
    'use strict';
    var countries =     {
        "Aaland Islands": {
            "Europe/Mariehamn": "Europe/Mariehamn"
        },
        "Afghanistan": {
            "Asia/Kabul": "Asia/Kabul"
        },
        "Albania": {
            "Europe/Tirane": "Europe/Tirane"
        },
        "Algeria": {
            "Africa/Algiers": "Africa/Algiers"
        },
        "Andorra": {
            "Europe/Andorra": "Europe/Andorra"
        },
        "Angola": {
            "Africa/Luanda": "Africa/Luanda"
        },
        "Anguilla": {
            "America/Anguilla": "America/Anguilla"
        },
        "Antarctica": {
            "Antarctica/Casey": "Antarctica/Casey",
            "Antarctica/Davis": "Antarctica/Davis",
            "Antarctica/DumontDUrville": "Antarctica/DumontDUrville",
            "Antarctica/Mawson": "Antarctica/Mawson",
            "Antarctica/McMurdo": "Antarctica/McMurdo",
            "Antarctica/Palmer": "Antarctica/Palmer",
            "Antarctica/Rothera": "Antarctica/Rothera",
            "Antarctica/South_Pole": "Antarctica/South Pole",
            "Antarctica/Syowa": "Antarctica/Syowa",
            "Antarctica/Vostok": "Antarctica/Vostok"
        },
        "Antigua & Barbuda": {
            "America/Antigua": "America/Antigua"
        },
        "Argentina": {
            "America/Argentina/Buenos_Aires": "America/Argentina/Buenos Aires",
            "America/Argentina/Catamarca": "America/Argentina/Catamarca",
            "America/Argentina/Cordoba": "America/Argentina/Cordoba",
            "America/Argentina/Jujuy": "America/Argentina/Jujuy",
            "America/Argentina/La_Rioja": "America/Argentina/La Rioja",
            "America/Argentina/Mendoza": "America/Argentina/Mendoza",
            "America/Argentina/Rio_Gallegos": "America/Argentina/Rio Gallegos",
            "America/Argentina/Salta": "America/Argentina/Salta",
            "America/Argentina/San_Juan": "America/Argentina/San Juan",
            "America/Argentina/San_Luis": "America/Argentina/San Luis",
            "America/Argentina/Tucuman": "America/Argentina/Tucuman",
            "America/Argentina/Ushuaia": "America/Argentina/Ushuaia"
        },
        "Armenia": {
            "Asia/Yerevan": "Asia/Yerevan"
        },
        "Aruba": {
            "America/Aruba": "America/Aruba"
        },
        "Australia": {
            "Antarctica/Macquarie": "Antarctica/Macquarie",
            "Australia/Adelaide": "Australia/Adelaide",
            "Australia/Brisbane": "Australia/Brisbane",
            "Australia/Broken_Hill": "Australia/Broken Hill",
            "Australia/Currie": "Australia/Currie",
            "Australia/Darwin": "Australia/Darwin",
            "Australia/Eucla": "Australia/Eucla",
            "Australia/Hobart": "Australia/Hobart",
            "Australia/Lindeman": "Australia/Lindeman",
            "Australia/Lord_Howe": "Australia/Lord Howe",
            "Australia/Melbourne": "Australia/Melbourne",
            "Australia/Perth": "Australia/Perth",
            "Australia/Sydney": "Australia/Sydney"
        },
        "Austria": {
            "Europe/Vienna": "Europe/Vienna"
        },
        "Azerbaijan": {
            "Asia/Baku": "Asia/Baku"
        },
        "Bahamas": {
            "America/Nassau": "America/Nassau"
        },
        "Bahrain": {
            "Asia/Bahrain": "Asia/Bahrain"
        },
        "Bangladesh": {
            "Asia/Dhaka": "Asia/Dhaka"
        },
        "Barbados": {
            "America/Barbados": "America/Barbados"
        },
        "Belarus": {
            "Europe/Minsk": "Europe/Minsk"
        },
        "Belgium": {
            "Europe/Brussels": "Europe/Brussels"
        },
        "Belize": {
            "America/Belize": "America/Belize"
        },
        "Benin": {
            "Africa/Porto": "Africa/Porto-Novo"
        },
        "Bermuda": {
            "Atlantic/Bermuda": "Atlantic/Bermuda"
        },
        "Bhutan": {
            "Asia/Thimphu": "Asia/Thimphu"
        },
        "Bolivia": {
            "America/La_Paz": "America/La Paz"
        },
        "Bonaire, St Eustatius & Saba": {
            "America/Kralendijk": "America/Kralendijk"
        },
        "Bosnia & Herzegovina": {
            "Europe/Sarajevo": "Europe/Sarajevo"
        },
        "Botswana": {
            "Africa/Gaborone": "Africa/Gaborone"
        },
        "Brazil": {
            "America/Araguaina": "America/Araguaina",
            "America/Bahia": "America/Bahia",
            "America/Belem": "America/Belem",
            "America/Boa_Vista": "America/Boa Vista",
            "America/Campo_Grande": "America/Campo Grande",
            "America/Cuiaba": "America/Cuiaba",
            "America/Eirunepe": "America/Eirunepe",
            "America/Fortaleza": "America/Fortaleza",
            "America/Maceio": "America/Maceio",
            "America/Manaus": "America/Manaus",
            "America/Noronha": "America/Noronha",
            "America/Porto_Velho": "America/Porto Velho",
            "America/Recife": "America/Recife",
            "America/Rio_Branco": "America/Rio Branco",
            "America/Santarem": "America/Santarem",
            "America/Sao_Paulo": "America/Sao Paulo"
        },
        "Britain (UK)": {
            "Europe/London": "Europe/London"
        },
        "British Indian Ocean Territory": {
            "Indian/Chagos": "Indian/Chagos"
        },
        "Brunei": {
            "Asia/Brunei": "Asia/Brunei"
        },
        "Bulgaria": {
            "Europe/Sofia": "Europe/Sofia"
        },
        "Burkina Faso": {
            "Africa/Ouagadougou": "Africa/Ouagadougou"
        },
        "Burundi": {
            "Africa/Bujumbura": "Africa/Bujumbura"
        },
        "Cambodia": {
            "Asia/Phnom_Penh": "Asia/Phnom Penh"
        },
        "Cameroon": {
            "Africa/Douala": "Africa/Douala"
        },
        "Canada": {
            "America/Atikokan": "America/Atikokan",
            "America/Blanc": "America/Blanc-Sablon",
            "America/Cambridge_Bay": "America/Cambridge Bay",
            "America/Creston": "America/Creston",
            "America/Dawson": "America/Dawson",
            "America/Dawson_Creek": "America/Dawson Creek",
            "America/Edmonton": "America/Edmonton",
            "America/Glace_Bay": "America/Glace Bay",
            "America/Goose_Bay": "America/Goose Bay",
            "America/Halifax": "America/Halifax",
            "America/Inuvik": "America/Inuvik",
            "America/Iqaluit": "America/Iqaluit",
            "America/Moncton": "America/Moncton",
            "America/Montreal": "America/Montreal",
            "America/Nipigon": "America/Nipigon",
            "America/Pangnirtung": "America/Pangnirtung",
            "America/Rainy_River": "America/Rainy River",
            "America/Rankin_Inlet": "America/Rankin Inlet",
            "America/Regina": "America/Regina",
            "America/Resolute": "America/Resolute",
            "America/St_Johns": "America/St Johns",
            "America/Swift_Current": "America/Swift Current",
            "America/Thunder_Bay": "America/Thunder Bay",
            "America/Toronto": "America/Toronto",
            "America/Vancouver": "America/Vancouver",
            "America/Whitehorse": "America/Whitehorse",
            "America/Winnipeg": "America/Winnipeg",
            "America/Yellowknife": "America/Yellowknife"
        },
        "Cape Verde": {
            "Atlantic/Cape_Verde": "Atlantic/Cape Verde"
        },
        "Cayman Islands": {
            "America/Cayman": "America/Cayman"
        },
        "Central African Rep.": {
            "Africa/Bangui": "Africa/Bangui"
        },
        "Chad": {
            "Africa/Ndjamena": "Africa/Ndjamena"
        },
        "Chile": {
            "America/Santiago": "America/Santiago",
            "Pacific/Easter": "Pacific/Easter"
        },
        "China": {
            "Asia/Chongqing": "Asia/Chongqing",
            "Asia/Harbin": "Asia/Harbin",
            "Asia/Kashgar": "Asia/Kashgar",
            "Asia/Shanghai": "Asia/Shanghai",
            "Asia/Urumqi": "Asia/Urumqi"
        },
        "Christmas Island": {
            "Indian/Christmas": "Indian/Christmas"
        },
        "Cocos (Keeling) Islands": {
            "Indian/Cocos": "Indian/Cocos"
        },
        "Colombia": {
            "America/Bogota": "America/Bogota"
        },
        "Comoros": {
            "Indian/Comoro": "Indian/Comoro"
        },
        "Congo (Dem. Rep.)": {
            "Africa/Kinshasa": "Africa/Kinshasa",
            "Africa/Lubumbashi": "Africa/Lubumbashi"
        },
        "Congo (Rep.)": {
            "Africa/Brazzaville": "Africa/Brazzaville"
        },
        "Cook Islands": {
            "Pacific/Rarotonga": "Pacific/Rarotonga"
        },
        "Costa Rica": {
            "America/Costa_Rica": "America/Costa Rica"
        },
        "Cote d'Ivoire": {
            "Africa/Abidjan": "Africa/Abidjan"
        },
        "Croatia": {
            "Europe/Zagreb": "Europe/Zagreb"
        },
        "Cuba": {
            "America/Havana": "America/Havana"
        },
        "Curacao": {
            "America/Curacao": "America/Curacao"
        },
        "Cyprus": {
            "Asia/Nicosia": "Asia/Nicosia"
        },
        "Czech Republic": {
            "Europe/Prague": "Europe/Prague"
        },
        "Denmark": {
            "Europe/Copenhagen": "Europe/Copenhagen"
        },
        "Djibouti": {
            "Africa/Djibouti": "Africa/Djibouti"
        },
        "Dominica": {
            "America/Dominica": "America/Dominica"
        },
        "Dominican Republic": {
            "America/Santo_Domingo": "America/Santo Domingo"
        },
        "East Timor": {
            "Asia/Dili": "Asia/Dili"
        },
        "Ecuador": {
            "America/Guayaquil": "America/Guayaquil",
            "Pacific/Galapagos": "Pacific/Galapagos"
        },
        "Egypt": {
            "Africa/Cairo": "Africa/Cairo"
        },
        "El Salvador": {
            "America/El_Salvador": "America/El Salvador"
        },
        "Equatorial Guinea": {
            "Africa/Malabo": "Africa/Malabo"
        },
        "Eritrea": {
            "Africa/Asmara": "Africa/Asmara"
        },
        "Estonia": {
            "Europe/Tallinn": "Europe/Tallinn"
        },
        "Ethiopia": {
            "Africa/Addis_Ababa": "Africa/Addis Ababa"
        },
        "Falkland Islands": {
            "Atlantic/Stanley": "Atlantic/Stanley"
        },
        "Faroe Islands": {
            "Atlantic/Faroe": "Atlantic/Faroe"
        },
        "Fiji": {
            "Pacific/Fiji": "Pacific/Fiji"
        },
        "Finland": {
            "Europe/Helsinki": "Europe/Helsinki"
        },
        "France": {
            "Europe/Paris": "Europe/Paris"
        },
        "French Guiana": {
            "America/Cayenne": "America/Cayenne"
        },
        "French Polynesia": {
            "Pacific/Gambier": "Pacific/Gambier",
            "Pacific/Marquesas": "Pacific/Marquesas",
            "Pacific/Tahiti": "Pacific/Tahiti"
        },
        "French Southern & Antarctic Lands": {
            "Indian/Kerguelen": "Indian/Kerguelen"
        },
        "Gabon": {
            "Africa/Libreville": "Africa/Libreville"
        },
        "Gambia": {
            "Africa/Banjul": "Africa/Banjul"
        },
        "Georgia": {
            "Asia/Tbilisi": "Asia/Tbilisi"
        },
        "Germany": {
            "Europe/Berlin": "Europe/Berlin",
            "Europe/Busingen": "Europe/Busingen"
        },
        "Ghana": {
            "Africa/Accra": "Africa/Accra"
        },
        "Gibraltar": {
            "Europe/Gibraltar": "Europe/Gibraltar"
        },
        "Greece": {
            "Europe/Athens": "Europe/Athens"
        },
        "Greenland": {
            "America/Danmarkshavn": "America/Danmarkshavn",
            "America/Godthab": "America/Godthab",
            "America/Scoresbysund": "America/Scoresbysund",
            "America/Thule": "America/Thule"
        },
        "Grenada": {
            "America/Grenada": "America/Grenada"
        },
        "Guadeloupe": {
            "America/Guadeloupe": "America/Guadeloupe"
        },
        "Guam": {
            "Pacific/Guam": "Pacific/Guam"
        },
        "Guatemala": {
            "America/Guatemala": "America/Guatemala"
        },
        "Guernsey": {
            "Europe/Guernsey": "Europe/Guernsey"
        },
        "Guinea": {
            "Africa/Conakry": "Africa/Conakry"
        },
        "Guinea-Bissau": {
            "Bissau": "Africa/Bissau"
        },
        "Guyana": {
            "America/Guyana": "America/Guyana"
        },
        "Haiti": {
            "America/Port": "America/Port-au-Prince"
        },
        "Honduras": {
            "America/Tegucigalpa": "America/Tegucigalpa"
        },
        "Hong Kong": {
            "Asia/Hong_Kong": "Asia/Hong Kong"
        },
        "Hungary": {
            "Europe/Budapest": "Europe/Budapest"
        },
        "Iceland": {
            "Atlantic/Reykjavik": "Atlantic/Reykjavik"
        },
        "India": {
            "Asia/Kolkata": "Asia/Kolkata"
        },
        "Indonesia": {
            "Asia/Jakarta": "Asia/Jakarta",
            "Asia/Jayapura": "Asia/Jayapura",
            "Asia/Makassar": "Asia/Makassar",
            "Asia/Pontianak": "Asia/Pontianak"
        },
        "Iran": {
            "Asia/Tehran": "Asia/Tehran"
        },
        "Iraq": {
            "Asia/Baghdad": "Asia/Baghdad"
        },
        "Ireland": {
            "Europe/Dublin": "Europe/Dublin"
        },
        "Isle of Man": {
            "Europe/Isle_of_Man": "Europe/Isle of Man"
        },
        "Israel": {
            "Asia/Jerusalem": "Asia/Jerusalem"
        },
        "Italy": {
            "Europe/Rome": "Europe/Rome"
        },
        "Jamaica": {
            "America/Jamaica": "America/Jamaica"
        },
        "Japan": {
            "Asia/Tokyo": "Asia/Tokyo"
        },
        "Jersey": {
            "Europe/Jersey": "Europe/Jersey"
        },
        "Jordan": {
            "Asia/Amman": "Asia/Amman"
        },
        "Kazakhstan": {
            "Asia/Almaty": "Asia/Almaty",
            "Asia/Aqtau": "Asia/Aqtau",
            "Asia/Aqtobe": "Asia/Aqtobe",
            "Asia/Oral": "Asia/Oral",
            "Asia/Qyzylorda": "Asia/Qyzylorda"
        },
        "Kenya": {
            "Africa/Nairobi": "Africa/Nairobi"
        },
        "Kiribati": {
            "Pacific/Enderbury": "Pacific/Enderbury",
            "Pacific/Kiritimati": "Pacific/Kiritimati",
            "Pacific/Tarawa": "Pacific/Tarawa"
        },
        "Korea (North)": {
            "Asia/Pyongyang": "Asia/Pyongyang"
        },
        "Korea (South)": {
            "Asia/Seoul": "Asia/Seoul"
        },
        "Kuwait": {
            "Asia/Kuwait": "Asia/Kuwait"
        },
        "Kyrgyzstan": {
            "Asia/Bishkek": "Asia/Bishkek"
        },
        "Laos": {
            "Asia/Vientiane": "Asia/Vientiane"
        },
        "Latvia": {
            "Europe/Riga": "Europe/Riga"
        },
        "Lebanon": {
            "Asia/Beirut": "Asia/Beirut"
        },
        "Lesotho": {
            "Africa/Maseru": "Africa/Maseru"
        },
        "Liberia": {
            "Africa/Monrovia": "Africa/Monrovia"
        },
        "Libya": {
            "Africa/Tripoli": "Africa/Tripoli"
        },
        "Liechtenstein": {
            "Europe/Vaduz": "Europe/Vaduz"
        },
        "Lithuania": {
            "Europe/Vilnius": "Europe/Vilnius"
        },
        "Luxembourg": {
            "Europe/Luxembourg": "Europe/Luxembourg"
        },
        "Macau": {
            "Asia/Macau": "Asia/Macau"
        },
        "Macedonia": {
            "Europe/Skopje": "Europe/Skopje"
        },
        "Madagascar": {
            "Indian/Antananarivo": "Indian/Antananarivo"
        },
        "Malawi": {
            "Africa/Blantyre": "Africa/Blantyre"
        },
        "Malaysia": {
            "Asia/Kuala_Lumpur": "Asia/Kuala Lumpur",
            "Asia/Kuching": "Asia/Kuching"
        },
        "Maldives": {
            "Indian/Maldives": "Indian/Maldives"
        },
        "Mali": {
            "Africa/Bamako": "Africa/Bamako"
        },
        "Malta": {
            "Europe/Malta": "Europe/Malta"
        },
        "Marshall Islands": {
            "Pacific/Kwajalein": "Pacific/Kwajalein",
            "Pacific/Majuro": "Pacific/Majuro"
        },
        "Martinique": {
            "America/Martinique": "America/Martinique"
        },
        "Mauritania": {
            "Africa/Nouakchott": "Africa/Nouakchott"
        },
        "Mauritius": {
            "Indian/Mauritius": "Indian/Mauritius"
        },
        "Mayotte": {
            "Indian/Mayotte": "Indian/Mayotte"
        },
        "Mexico": {
            "America/Bahia_Banderas": "America/Bahia Banderas",
            "America/Cancun": "America/Cancun",
            "America/Chihuahua": "America/Chihuahua",
            "America/Hermosillo": "America/Hermosillo",
            "America/Matamoros": "America/Matamoros",
            "America/Mazatlan": "America/Mazatlan",
            "America/Merida": "America/Merida",
            "America/Mexico_City": "America/Mexico City",
            "America/Monterrey": "America/Monterrey",
            "America/Ojinaga": "America/Ojinaga",
            "America/Santa_Isabel": "America/Santa Isabel",
            "America/Tijuana": "America/Tijuana"
        },
        "Micronesia": {
            "Pacific/Chuuk": "Pacific/Chuuk",
            "Pacific/Kosrae": "Pacific/Kosrae",
            "Pacific/Pohnpei": "Pacific/Pohnpei"
        },
        "Moldova": {
            "Europe/Chisinau": "Europe/Chisinau"
        },
        "Monaco": {
            "Europe/Monaco": "Europe/Monaco"
        },
        "Mongolia": {
            "Asia/Choibalsan": "Asia/Choibalsan",
            "Asia/Hovd": "Asia/Hovd",
            "Asia/Ulaanbaatar": "Asia/Ulaanbaatar"
        },
        "Montenegro": {
            "Europe/Podgorica": "Europe/Podgorica"
        },
        "Montserrat": {
            "America/Montserrat": "America/Montserrat"
        },
        "Morocco": {
            "Africa/Casablanca": "Africa/Casablanca"
        },
        "Mozambique": {
            "Africa/Maputo": "Africa/Maputo"
        },
        "Myanmar (Burma)": {
            "Asia/Rangoon": "Asia/Rangoon"
        },
        "Namibia": {
            "Africa/Windhoek": "Africa/Windhoek"
        },
        "Nauru": {
            "Pacific/Nauru": "Pacific/Nauru"
        },
        "Nepal": {
            "Asia/Kathmandu": "Asia/Kathmandu"
        },
        "Netherlands": {
            "Europe/Amsterdam": "Europe/Amsterdam"
        },
        "New Caledonia": {
            "Pacific/Noumea": "Pacific/Noumea"
        },
        "New Zealand": {
            "Pacific/Auckland": "Pacific/Auckland",
            "Pacific/Chatham": "Pacific/Chatham"
        },
        "Nicaragua": {
            "America/Managua": "America/Managua"
        },
        "Niger": {
            "Africa/Niamey": "Africa/Niamey"
        },
        "Nigeria": {
            "Africa/Lagos": "Africa/Lagos"
        },
        "Niue": {
            "Pacific/Niue": "Pacific/Niue"
        },
        "Norfolk Island": {
            "Pacific/Norfolk": "Pacific/Norfolk"
        },
        "Northern Mariana Islands": {
            "Pacific/Saipan": "Pacific/Saipan"
        },
        "Norway": {
            "Europe/Oslo": "Europe/Oslo"
        },
        "Oman": {
            "Asia/Muscat": "Asia/Muscat"
        },
        "Pakistan": {
            "Asia/Karachi": "Asia/Karachi"
        },
        "Palau": {
            "Pacific/Palau": "Pacific/Palau"
        },
        "Palestine": {
            "Asia/Gaza": "Asia/Gaza",
            "Asia/Hebron": "Asia/Hebron"
        },
        "Panama": {
            "America/Panama": "America/Panama"
        },
        "Papua New Guinea": {
            "Pacific/Port_Moresby": "Pacific/Port Moresby"
        },
        "Paraguay": {
            "America/Asuncion": "America/Asuncion"
        },
        "Peru": {
            "America/Lima": "America/Lima"
        },
        "Philippines": {
            "Asia/Manila": "Asia/Manila"
        },
        "Pitcairn": {
            "Pacific/Pitcairn": "Pacific/Pitcairn"
        },
        "Poland": {
            "Europe/Warsaw": "Europe/Warsaw"
        },
        "Portugal": {
            "Atlantic/Azores": "Atlantic/Azores",
            "Atlantic/Madeira": "Atlantic/Madeira",
            "Europe/Lisbon": "Europe/Lisbon"
        },
        "Puerto Rico": {
            "America/Puerto_Rico": "America/Puerto Rico"
        },
        "Qatar": {
            "Asia/Qatar": "Asia/Qatar"
        },
        "Reunion": {
            "Indian/Reunion": "Indian/Reunion"
        },
        "Romania": {
            "Europe/Bucharest": "Europe/Bucharest"
        },
        "Russia": {
            "Asia/Anadyr": "Asia/Anadyr",
            "Asia/Irkutsk": "Asia/Irkutsk",
            "Asia/Kamchatka": "Asia/Kamchatka",
            "Asia/Khandyga": "Asia/Khandyga",
            "Asia/Krasnoyarsk": "Asia/Krasnoyarsk",
            "Asia/Magadan": "Asia/Magadan",
            "Asia/Novokuznetsk": "Asia/Novokuznetsk",
            "Asia/Novosibirsk": "Asia/Novosibirsk",
            "Asia/Omsk": "Asia/Omsk",
            "Asia/Sakhalin": "Asia/Sakhalin",
            "Asia/Ust": "Asia/Ust-Nera",
            "Asia/Vladivostok": "Asia/Vladivostok",
            "Asia/Yakutsk": "Asia/Yakutsk",
            "Asia/Yekaterinburg": "Asia/Yekaterinburg",
            "Europe/Kaliningrad": "Europe/Kaliningrad",
            "Europe/Moscow": "Europe/Moscow",
            "Europe/Samara": "Europe/Samara",
            "Europe/Volgograd": "Europe/Volgograd"
        },
        "Rwanda": {
            "Africa/Kigali": "Africa/Kigali"
        },
        "Samoa (American)": {
            "Pacific/Pago_Pago": "Pacific/Pago Pago"
        },
        "Samoa (western)": {
            "Pacific/Apia": "Pacific/Apia"
        },
        "San Marino": {
            "Europe/San_Marino": "Europe/San Marino"
        },
        "Sao Tome & Principe": {
            "Africa/Sao_Tome": "Africa/Sao Tome"
        },
        "Saudi Arabia": {
            "Asia/Riyadh": "Asia/Riyadh"
        },
        "Senegal": {
            "Africa/Dakar": "Africa/Dakar"
        },
        "Serbia": {
            "Europe/Belgrade": "Europe/Belgrade"
        },
        "Seychelles": {
            "Indian/Mahe": "Indian/Mahe"
        },
        "Sierra Leone": {
            "Africa/Freetown": "Africa/Freetown"
        },
        "Singapore": {
            "Asia/Singapore": "Asia/Singapore"
        },
        "Slovakia": {
            "Europe/Bratislava": "Europe/Bratislava"
        },
        "Slovenia": {
            "Europe/Ljubljana": "Europe/Ljubljana"
        },
        "Solomon Islands": {
            "Pacific/Guadalcanal": "Pacific/Guadalcanal"
        },
        "Somalia": {
            "Africa/Mogadishu": "Africa/Mogadishu"
        },
        "South Africa": {
            "Africa/Johannesburg": "Africa/Johannesburg"
        },
        "South Georgia & the South Sandwich Islands": {
            "Atlantic/South_Georgia": "Atlantic/South Georgia"
        },
        "South Sudan": {
            "Africa/Juba": "Africa/Juba"
        },
        "Spain": {
            "Africa/Ceuta": "Africa/Ceuta",
            "Atlantic/Canary": "Atlantic/Canary",
            "Europe/Madrid": "Europe/Madrid"
        },
        "Sri Lanka": {
            "Asia/Colombo": "Asia/Colombo"
        },
        "St Barthelemy": {
            "America/St_Barthelemy": "America/St Barthelemy"
        },
        "St Helena": {
            "Atlantic/St_Helena": "Atlantic/St Helena"
        },
        "St Kitts & Nevis": {
            "America/St_Kitts": "America/St Kitts"
        },
        "St Lucia": {
            "America/St_Lucia": "America/St Lucia"
        },
        "St Maarten (Dutch part)": {
            "America/Lower_Princes": "America/Lower Princes"
        },
        "St Martin (French part)": {
            "America/Marigot": "America/Marigot"
        },
        "St Pierre & Miquelon": {
            "America/Miquelon": "America/Miquelon"
        },
        "St Vincent": {
            "America/St_Vincent": "America/St Vincent"
        },
        "Sudan": {
            "Africa/Khartoum": "Africa/Khartoum"
        },
        "Suriname": {
            "America/Paramaribo": "America/Paramaribo"
        },
        "Svalbard & Jan Mayen": {
            "Arctic/Longyearbyen": "Arctic/Longyearbyen"
        },
        "Swaziland": {
            "Africa/Mbabane": "Africa/Mbabane"
        },
        "Sweden": {
            "Europe/Stockholm": "Europe/Stockholm"
        },
        "Switzerland": {
            "Europe/Zurich": "Europe/Zurich"
        },
        "Syria": {
            "Asia/Damascus": "Asia/Damascus"
        },
        "Taiwan": {
            "Asia/Taipei": "Asia/Taipei"
        },
        "Tajikistan": {
            "Asia/Dushanbe": "Asia/Dushanbe"
        },
        "Tanzania": {
            "Africa/Dar_es_Salaam": "Africa/Dar es Salaam"
        },
        "Thailand": {
            "Asia/Bangkok": "Asia/Bangkok"
        },
        "Togo": {
            "Africa/Lome": "Africa/Lome"
        },
        "Tokelau": {
            "Pacific/Fakaofo": "Pacific/Fakaofo"
        },
        "Tonga": {
            "Pacific/Tongatapu": "Pacific/Tongatapu"
        },
        "Trinidad & Tobago": {
            "America/Port_of_Spain": "America/Port of Spain"
        },
        "Tunisia": {
            "Africa/Tunis": "Africa/Tunis"
        },
        "Turkey": {
            "Europe/Istanbul": "Europe/Istanbul"
        },
        "Turkmenistan": {
            "Asia/Ashgabat": "Asia/Ashgabat"
        },
        "Turks & Caicos Is": {
            "America/Grand_Turk": "America/Grand Turk"
        },
        "Tuvalu": {
            "Pacific/Funafuti": "Pacific/Funafuti"
        },
        "US minor outlying islands": {
            "Pacific/Johnston": "Pacific/Johnston",
            "Pacific/Midway": "Pacific/Midway",
            "Pacific/Wake": "Pacific/Wake"
        },
        "Uganda": {
            "Africa/Kampala": "Africa/Kampala"
        },
        "Ukraine": {
            "Europe/Kiev": "Europe/Kiev",
            "Europe/Simferopol": "Europe/Simferopol",
            "Europe/Uzhgorod": "Europe/Uzhgorod",
            "Europe/Zaporozhye": "Europe/Zaporozhye"
        },
        "United Arab Emirates": {
            "Asia/Dubai": "Asia/Dubai"
        },
        "United States": {
            "America/Adak": "America/Adak",
            "America/Anchorage": "America/Anchorage",
            "America/Boise": "America/Boise",
            "America/Chicago": "America/Chicago",
            "America/Denver": "America/Denver",
            "America/Detroit": "America/Detroit",
            "America/Indiana/Indianapolis": "America/Indiana/Indianapolis",
            "America/Indiana/Knox": "America/Indiana/Knox",
            "America/Indiana/Marengo": "America/Indiana/Marengo",
            "America/Indiana/Petersburg": "America/Indiana/Petersburg",
            "America/Indiana/Tell_City": "America/Indiana/Tell City",
            "America/Indiana/Vevay": "America/Indiana/Vevay",
            "America/Indiana/Vincennes": "America/Indiana/Vincennes",
            "America/Indiana/Winamac": "America/Indiana/Winamac",
            "America/Juneau": "America/Juneau",
            "America/Kentucky/Louisville": "America/Kentucky/Louisville",
            "America/Kentucky/Monticello": "America/Kentucky/Monticello",
            "America/Los_Angeles": "America/Los Angeles",
            "America/Menominee": "America/Menominee",
            "America/Metlakatla": "America/Metlakatla",
            "America/New_York": "America/New York",
            "America/Nome": "America/Nome",
            "America/North_Dakota/Beulah": "America/North Dakota/Beulah",
            "America/North_Dakota/Center": "America/North Dakota/Center",
            "America/North_Dakota/New_Salem": "America/North Dakota/New Salem",
            "America/Phoenix": "America/Phoenix",
            "America/Shiprock": "America/Shiprock",
            "America/Sitka": "America/Sitka",
            "America/Yakutat": "America/Yakutat",
            "Pacific/Honolulu": "Pacific/Honolulu"
        },
        "Uruguay": {
            "America/Montevideo": "America/Montevideo"
        },
        "Uzbekistan": {
            "Asia/Samarkand": "Asia/Samarkand",
            "Asia/Tashkent": "Asia/Tashkent"
        },
        "Vanuatu": {
            "Pacific/Efate": "Pacific/Efate"
        },
        "Vatican City": {
            "Europe/Vatican": "Europe/Vatican"
        },
        "Venezuela": {
            "America/Caracas": "America/Caracas"
        },
        "Vietnam": {
            "Asia/Ho_Chi_Minh": "Asia/Ho Chi Minh"
        },
        "Virgin Islands (UK)": {
            "America/Tortola": "America/Tortola"
        },
        "Virgin Islands (US)": {
            "America/St_Thomas": "America/St Thomas"
        },
        "Wallis & Futuna": {
            "Pacific/Wallis": "Pacific/Wallis"
        },
        "Western Sahara": {
            "Africa/El_Aaiun": "Africa/El Aaiun"
        },
        "Yemen": {
            "Asia/Aden": "Asia/Aden"
        },
        "Zambia": {
            "Africa/Lusaka": "Africa/Lusaka"
        },
        "Zimbabwe": {
            "Africa/Harare": "Africa/Harare"
        }
    }
    
    var timezones = {};
    
    angular.forEach(countries, function(country){
        angular.forEach(country, function(zone, k){
            timezones[k] = zone;
        });
    });
    
    services.provider('$timezones', function () {
        
        this.$get = function($interval) {
            
            
            
            var wrappedService = {
                countries: countries,
                timezones: timezones
            }
            return wrappedService;
        }

    });

});
