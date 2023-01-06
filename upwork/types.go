package upwork

import "time"

type LoggedOutError struct {
	Error struct {
		Code    int    `json:"code"`
		Message string `json:"message"`
	}
}

type UpworkApiResponse struct {
	URL           string `json:"url"`
	SearchGUID    string `json:"searchGuid"`
	SearchResults struct {
		Q      string `json:"q"`
		Paging struct {
			Total       int `json:"total"`
			Offset      int `json:"offset"`
			Count       int `json:"count"`
			ResultSetTs int `json:"resultSetTs"`
		} `json:"paging"`
		Spellcheck struct {
			CorrectedQueries []interface{} `json:"corrected_queries"`
		} `json:"spellcheck"`
		Jobs []struct {
			Title           string        `json:"title"`
			CreatedOn       time.Time     `json:"createdOn"`
			Type            int           `json:"type"`
			Ciphertext      string        `json:"ciphertext"`
			Description     string        `json:"description"`
			Category2       interface{}   `json:"category2"`
			Subcategory2    interface{}   `json:"subcategory2"`
			Skills          []interface{} `json:"skills"`
			Duration        string        `json:"duration"`
			ShortDuration   string        `json:"shortDuration"`
			DurationLabel   string        `json:"durationLabel"`
			Engagement      string        `json:"engagement"`
			ShortEngagement string        `json:"shortEngagement"`
			Amount          struct {
				CurrencyCode string `json:"currencyCode"`
				Amount       int    `json:"amount"`
			} `json:"amount"`
			Recno  int    `json:"recno"`
			UID    string `json:"uid"`
			Client struct {
				PaymentVerificationStatus int `json:"paymentVerificationStatus"`
				Location                  struct {
					Country string `json:"country"`
				} `json:"location"`
				TotalSpent           float64     `json:"totalSpent"`
				TotalReviews         int         `json:"totalReviews"`
				TotalFeedback        float64     `json:"totalFeedback"`
				CompanyRid           int         `json:"companyRid"`
				CompanyName          interface{} `json:"companyName"`
				EdcUserID            int         `json:"edcUserId"`
				LastContractPlatform interface{} `json:"lastContractPlatform"`
				LastContractRid      int         `json:"lastContractRid"`
				LastContractTitle    interface{} `json:"lastContractTitle"`
				FeedbackText         string      `json:"feedbackText"`
				CompanyOrgUID        string      `json:"companyOrgUid"`
				HasFinancialPrivacy  bool        `json:"hasFinancialPrivacy"`
			} `json:"client"`
			FreelancersToHire               int           `json:"freelancersToHire"`
			RelevanceEncoded                string        `json:"relevanceEncoded"`
			EnterpriseJob                   bool          `json:"enterpriseJob"`
			TierText                        string        `json:"tierText"`
			Tier                            string        `json:"tier"`
			TierLabel                       string        `json:"tierLabel"`
			IsSaved                         interface{}   `json:"isSaved"`
			Feedback                        string        `json:"feedback"`
			ProposalsTier                   string        `json:"proposalsTier"`
			IsApplied                       bool          `json:"isApplied"`
			Sticky                          bool          `json:"sticky"`
			StickyLabel                     string        `json:"stickyLabel"`
			JobTs                           string        `json:"jobTs"`
			PrefFreelancerLocationMandatory bool          `json:"prefFreelancerLocationMandatory"`
			PrefFreelancerLocation          []interface{} `json:"prefFreelancerLocation"`
			Premium                         bool          `json:"premium"`
			PlusBadge                       interface{}   `json:"plusBadge"`
			PublishedOn                     time.Time     `json:"publishedOn"`
			RenewedOn                       time.Time     `json:"renewedOn"`
			SandsService                    interface{}   `json:"sandsService"`
			SandsSpec                       interface{}   `json:"sandsSpec"`
			SandsAttrs                      interface{}   `json:"sandsAttrs"`
			Occupation                      interface{}   `json:"occupation"`
			Attrs                           []struct {
				ParentSkillUID interface{} `json:"parentSkillUid"`
				FreeText       interface{} `json:"freeText"`
				SkillType      int         `json:"skillType"`
				UID            string      `json:"uid"`
				Highlighted    bool        `json:"highlighted"`
				PrettyName     string      `json:"prettyName"`
			} `json:"attrs"`
			IsLocal     bool          `json:"isLocal"`
			WorkType    interface{}   `json:"workType"`
			Locations   []interface{} `json:"locations"`
			Occupations struct {
				Category struct {
					UID       string `json:"uid"`
					PrefLabel string `json:"prefLabel"`
				} `json:"category"`
				Subcategories []struct {
					UID       string `json:"uid"`
					PrefLabel string `json:"prefLabel"`
				} `json:"subcategories"`
				Oservice struct {
					UID       string `json:"uid"`
					PrefLabel string `json:"prefLabel"`
				} `json:"oservice"`
			} `json:"occupations"`
			WeeklyBudget                       interface{}   `json:"weeklyBudget"`
			HourlyBudgetText                   interface{}   `json:"hourlyBudgetText"`
			Tags                               []interface{} `json:"tags"`
			ClientRelation                     interface{}   `json:"clientRelation"`
			TotalFreelancersToHire             int           `json:"totalFreelancersToHire"`
			TeamUID                            interface{}   `json:"teamUid"`
			MultipleFreelancersToHirePredicted interface{}   `json:"multipleFreelancersToHirePredicted"`
			ConnectPrice                       int           `json:"connectPrice"`
		} `json:"jobs"`
		Facets struct {
			JobType struct {
				Num0           int `json:"0"`
				Num1           int `json:"1"`
				WeeklyRetainer int `json:"weekly_retainer"`
			} `json:"jobType"`
			Workload struct {
				None     int `json:"none"`
				FullTime int `json:"full_time"`
				AsNeeded int `json:"as_needed"`
			} `json:"workload"`
			Duration struct {
				Ongoing  int `json:"ongoing"`
				Week     int `json:"week"`
				Month    int `json:"month"`
				Semester int `json:"semester"`
				Quarter  int `json:"quarter"`
			} `json:"duration"`
			ClientHires struct {
				Num0 int `json:"0"`
				One9 int `json:"1-9"`
				One0 int `json:"10-"`
			} `json:"clientHires"`
			Budget struct {
				Num0      int `json:"0"`
				Num50     int `json:"50"`
				Num100    int `json:"100"`
				Num250    int `json:"250"`
				Num500    int `json:"500"`
				Num1000   int `json:"1000"`
				Num2000   int `json:"2000"`
				Num5000   int `json:"5000"`
				Num10000  int `json:"10000"`
				Num20000  int `json:"20000"`
				Num100000 int `json:"100000"`
			} `json:"budget"`
			ContractorTier struct {
				Num1 int `json:"1"`
				Num2 int `json:"2"`
				Num3 int `json:"3"`
			} `json:"contractorTier"`
			Categories2 []struct {
				Category struct {
					Slug        string `json:"slug"`
					Name        string `json:"name"`
					ResultCount int    `json:"resultCount"`
				} `json:"category"`
				Subcategories []struct {
					Slug        string `json:"slug"`
					Name        string `json:"name"`
					ResultCount int    `json:"resultCount"`
				} `json:"subcategories"`
			} `json:"categories2"`
			PreviousClients struct {
				All int `json:"all"`
			} `json:"previousClients"`
			Proposals struct {
				Five9         int `json:"5-9"`
				Two049        int `json:"20-49"`
				Zero4         int `json:"0-4"`
				One014        int `json:"10-14"`
				NAMING_FAILED int `json:"*"`
				One519        int `json:"15-19"`
			} `json:"proposals"`
			Amount struct {
				FiveThousandPlus          int `json:"5000-"`
				ANY                       int `json:"*"`
				Hundred_to_Five_hundred   int `json:"100-499"`
				Thousand_to_five_thousand int `json:"1000-4999"`
				ZeroTo_Ninty_Nine         int `json:"0-99"`
				FiveHundred_to_thousand   int `json:"500-999"`
			} `json:"amount"`
			UserLocationMatch struct {
				Num1 int `json:"1"`
			} `json:"userLocationMatch"`
			Services   interface{} `json:"services"`
			Specs      interface{} `json:"specs"`
			Attrs      interface{} `json:"attrs"`
			DurationV2 struct {
				Week    int `json:"week"`
				Ongoing int `json:"ongoing"`
				Month   int `json:"month"`
			} `json:"durationV2"`
			Enterprise struct {
				True  int `json:"true"`
				False int `json:"false"`
			} `json:"enterprise"`
			UserDomesticJobsLastDays struct {
				Num7 int `json:"7"`
			} `json:"userDomesticJobsLastDays"`
			DurationV3 struct {
				Weeks    int `json:"weeks"`
				Months   int `json:"months"`
				Semester int `json:"semester"`
				Ongoing  int `json:"ongoing"`
			} `json:"durationV3"`
			Occupations []struct {
				TaxonomyLevel string `json:"taxonomyLevel"`
				UID           string `json:"uid"`
				Label         string `json:"label"`
				Count         int    `json:"count"`
				Occupations   []struct {
					TaxonomyLevel string `json:"taxonomyLevel"`
					UID           string `json:"uid"`
					Label         string `json:"label"`
					Count         int    `json:"count"`
					Occupations   []struct {
						TaxonomyLevel string      `json:"taxonomyLevel"`
						UID           string      `json:"uid"`
						Label         string      `json:"label"`
						Count         int         `json:"count"`
						Occupations   interface{} `json:"occupations"`
					} `json:"occupations"`
				} `json:"occupations"`
			} `json:"occupations"`
			EngagementDuration interface{} `json:"engagementDuration"`
			FreelancersNeeded  struct {
				Zero1 int `json:"0-1"`
				Six   int `json:"6-"`
				Two5  int `json:"2-5"`
				Two   int `json:"2-"`
				Zero  int `json:"0-"`
			} `json:"freelancersNeeded"`
			Location struct {
				SouthAmerica                     int `json:"South America"`
				PapuaNewGuinea                   int `json:"Papua New Guinea"`
				Cambodia                         int `json:"Cambodia"`
				Paraguay                         int `json:"Paraguay"`
				Kazakhstan                       int `json:"Kazakhstan"`
				Bahamas                          int `json:"Bahamas"`
				MarshallIslands                  int `json:"Marshall Islands"`
				Mali                             int `json:"Mali"`
				Panama                           int `json:"Panama"`
				Guadeloupe                       int `json:"Guadeloupe"`
				EasternAsia                      int `json:"Eastern Asia"`
				Laos                             int `json:"Laos"`
				Argentina                        int `json:"Argentina"`
				Seychelles                       int `json:"Seychelles"`
				Africa                           int `json:"Africa"`
				Zambia                           int `json:"Zambia"`
				Belize                           int `json:"Belize"`
				Bahrain                          int `json:"Bahrain"`
				Congo                            int `json:"Congo"`
				GuineaBissau                     int `json:"Guinea-Bissau"`
				SaintBarthelemy                  int `json:"Saint Barthelemy"`
				Namibia                          int `json:"Namibia"`
				PalestinianTerritories           int `json:"Palestinian Territories"`
				Finland                          int `json:"Finland"`
				FaroeIslands                     int `json:"Faroe Islands"`
				NetherlandsAntilles              int `json:"Netherlands Antilles"`
				Georgia                          int `json:"Georgia"`
				SaintKittsAndNevis               int `json:"Saint Kitts and Nevis"`
				Yemen                            int `json:"Yemen"`
				PuertoRico                       int `json:"Puerto Rico"`
				Madagascar                       int `json:"Madagascar"`
				Aruba                            int `json:"Aruba"`
				Sweden                           int `json:"Sweden"`
				Malawi                           int `json:"Malawi"`
				Andorra                          int `json:"Andorra"`
				Liechtenstein                    int `json:"Liechtenstein"`
				Poland                           int `json:"Poland"`
				UnitedStatesVirginIslands        int `json:"United States Virgin Islands"`
				Jordan                           int `json:"Jordan"`
				Bulgaria                         int `json:"Bulgaria"`
				Tunisia                          int `json:"Tunisia"`
				UnitedArabEmirates               int `json:"United Arab Emirates"`
				Kenya                            int `json:"Kenya"`
				FrenchPolynesia                  int `json:"French Polynesia"`
				Lebanon                          int `json:"Lebanon"`
				Djibouti                         int `json:"Djibouti"`
				Azerbaijan                       int `json:"Azerbaijan"`
				CzechRepublic                    int `json:"Czech Republic"`
				SaintLucia                       int `json:"Saint Lucia"`
				Mauritania                       int `json:"Mauritania"`
				Guernsey                         int `json:"Guernsey"`
				Mayotte                          int `json:"Mayotte"`
				SanMarino                        int `json:"San Marino"`
				Israel                           int `json:"Israel"`
				Australia                        int `json:"Australia"`
				Tajikistan                       int `json:"Tajikistan"`
				Myanmar                          int `json:"Myanmar"`
				CentralAmerica                   int `json:"Central America"`
				EasternAfrica                    int `json:"Eastern Africa"`
				Cameroon                         int `json:"Cameroon"`
				Gibraltar                        int `json:"Gibraltar"`
				Cyprus                           int `json:"Cyprus"`
				Malaysia                         int `json:"Malaysia"`
				Oman                             int `json:"Oman"`
				Iceland                          int `json:"Iceland"`
				Armenia                          int `json:"Armenia"`
				Gabon                            int `json:"Gabon"`
				WesternAsia                      int `json:"Western Asia"`
				Polynesia                        int `json:"Polynesia"`
				Luxembourg                       int `json:"Luxembourg"`
				Brazil                           int `json:"Brazil"`
				TurksAndCaicosIslands            int `json:"Turks and Caicos Islands"`
				Algeria                          int `json:"Algeria"`
				Slovenia                         int `json:"Slovenia"`
				Jersey                           int `json:"Jersey"`
				AntiguaAndBarbuda                int `json:"Antigua and Barbuda"`
				Ecuador                          int `json:"Ecuador"`
				Colombia                         int `json:"Colombia"`
				Moldova                          int `json:"Moldova"`
				Vanuatu                          int `json:"Vanuatu"`
				UnitedStatesMinorOutlyingIslands int `json:"United States Minor Outlying Islands"`
				Italy                            int `json:"Italy"`
				Honduras                         int `json:"Honduras"`
				Haiti                            int `json:"Haiti"`
				Burundi                          int `json:"Burundi"`
				Singapore                        int `json:"Singapore"`
				FrenchGuiana                     int `json:"French Guiana"`
				AmericanSamoa                    int `json:"American Samoa"`
				Russia                           int `json:"Russia"`
				Netherlands                      int `json:"Netherlands"`
				China                            int `json:"China"`
				Martinique                       int `json:"Martinique"`
				SaintPierreAndMiquelon           int `json:"Saint Pierre and Miquelon"`
				Kyrgyzstan                       int `json:"Kyrgyzstan"`
				Reunion                          int `json:"Reunion"`
				Bhutan                           int `json:"Bhutan"`
				Romania                          int `json:"Romania"`
				MiddleAfrica                     int `json:"Middle Africa"`
				Togo                             int `json:"Togo"`
				SouthernAsia                     int `json:"Southern Asia"`
				Philippines                      int `json:"Philippines"`
				CoteDIvoire                      int `json:"Cote d\'Ivoire"`
				Uzbekistan                       int `json:"Uzbekistan"`
				Asia                             int `json:"Asia"`
				BritishVirginIslands             int `json:"British Virgin Islands"`
				Zimbabwe                         int `json:"Zimbabwe"`
				BritishIndianOceanTerritory      int `json:"British Indian Ocean Territory"`
				Montenegro                       int `json:"Montenegro"`
				Indonesia                        int `json:"Indonesia"`
				Dominica                         int `json:"Dominica"`
				Benin                            int `json:"Benin"`
				Angola                           int `json:"Angola"`
				EasternEurope                    int `json:"Eastern Europe"`
				Portugal                         int `json:"Portugal"`
				BruneiDarussalam                 int `json:"Brunei Darussalam"`
				NewCaledonia                     int `json:"New Caledonia"`
				Grenada                          int `json:"Grenada"`
				Greece                           int `json:"Greece"`
				CaymanIslands                    int `json:"Cayman Islands"`
				Mongolia                         int `json:"Mongolia"`
				Latvia                           int `json:"Latvia"`
				Morocco                          int `json:"Morocco"`
				Guyana                           int `json:"Guyana"`
				Guatemala                        int `json:"Guatemala"`
				Chile                            int `json:"Chile"`
				Nepal                            int `json:"Nepal"`
				NorthernEurope                   int `json:"Northern Europe"`
				IsleOfMan                        int `json:"Isle of Man"`
				Ukraine                          int `json:"Ukraine"`
				Tanzania                         int `json:"Tanzania"`
				Ghana                            int `json:"Ghana"`
				HolySee                          int `json:"Holy See"`
				CentralAsia                      int `json:"Central Asia"`
				Anguilla                         int `json:"Anguilla"`
				SouthEasternAsia                 int `json:"South-Eastern Asia"`
				India                            int `json:"India"`
				Canada                           int `json:"Canada"`
				Maldives                         int `json:"Maldives"`
				Turkey                           int `json:"Turkey"`
				Belgium                          int `json:"Belgium"`
				Taiwan                           int `json:"Taiwan"`
				TrinidadAndTobago                int `json:"Trinidad and Tobago"`
				SouthernEurope                   int `json:"Southern Europe"`
				SouthAfrica                      int `json:"South Africa"`
				Bermuda                          int `json:"Bermuda"`
				AlandIslands                     int `json:"Aland Islands"`
				Jamaica                          int `json:"Jamaica"`
				Turkmenistan                     int `json:"Turkmenistan"`
				Peru                             int `json:"Peru"`
				Germany                          int `json:"Germany"`
				Americas                         int `json:"Americas"`
				Fiji                             int `json:"Fiji"`
				HongKong                         int `json:"Hong Kong"`
				UnitedStates                     int `json:"United States"`
				Guinea                           int `json:"Guinea"`
				MicronesiaFederatedStatesOf      int `json:"Micronesia, Federated States of"`
				Somalia                          int `json:"Somalia"`
				Chad                             int `json:"Chad"`
				Thailand                         int `json:"Thailand"`
				Kiribati                         int `json:"Kiribati"`
				EquatorialGuinea                 int `json:"Equatorial Guinea"`
				CostaRica                        int `json:"Costa Rica"`
				SaintMartinFrenchPart            int `json:"Saint Martin (French part)"`
				Vietnam                          int `json:"Vietnam"`
				Nigeria                          int `json:"Nigeria"`
				Kuwait                           int `json:"Kuwait"`
				Croatia                          int `json:"Croatia"`
				Uruguay                          int `json:"Uruguay"`
				SriLanka                         int `json:"Sri Lanka"`
				CookIslands                      int `json:"Cook Islands"`
				TimorLeste                       int `json:"Timor-Leste"`
				UnitedKingdom                    int `json:"United Kingdom"`
				Switzerland                      int `json:"Switzerland"`
				Samoa                            int `json:"Samoa"`
				Spain                            int `json:"Spain"`
				WesternAfrica                    int `json:"Western Africa"`
				Venezuela                        int `json:"Venezuela"`
				BurkinaFaso                      int `json:"Burkina Faso"`
				Swaziland                        int `json:"Swaziland"`
				Caribbean                        int `json:"Caribbean"`
				Estonia                          int `json:"Estonia"`
				Niue                             int `json:"Niue"`
				SouthKorea                       int `json:"South Korea"`
				Austria                          int `json:"Austria"`
				Mozambique                       int `json:"Mozambique"`
				ElSalvador                       int `json:"El Salvador"`
				Monaco                           int `json:"Monaco"`
				Guam                             int `json:"Guam"`
				Lesotho                          int `json:"Lesotho"`
				Tonga                            int `json:"Tonga"`
				NorthernAfrica                   int `json:"Northern Africa"`
				WesternSahara                    int `json:"Western Sahara"`
				Hungary                          int `json:"Hungary"`
				Japan                            int `json:"Japan"`
				Europe                           int `json:"Europe"`
				Curacao                          int `json:"Curacao"`
				Belarus                          int `json:"Belarus"`
				Mauritius                        int `json:"Mauritius"`
				WesternEurope                    int `json:"Western Europe"`
				Albania                          int `json:"Albania"`
				NewZealand                       int `json:"New Zealand"`
				NorthernAmerica                  int `json:"Northern America"`
				SintMaartenDutchPart             int `json:"Sint Maarten (Dutch part)"`
				Senegal                          int `json:"Senegal"`
				Macedonia                        int `json:"Macedonia"`
				Ethiopia                         int `json:"Ethiopia"`
				Egypt                            int `json:"Egypt"`
				SierraLeone                      int `json:"Sierra Leone"`
				Bolivia                          int `json:"Bolivia"`
				Oceania                          int `json:"Oceania"`
				Malta                            int `json:"Malta"`
				SaudiArabia                      int `json:"Saudi Arabia"`
				CapeVerde                        int `json:"Cape Verde"`
				Pakistan                         int `json:"Pakistan"`
				Gambia                           int `json:"Gambia"`
				Qatar                            int `json:"Qatar"`
				Ireland                          int `json:"Ireland"`
				Slovakia                         int `json:"Slovakia"`
				Serbia                           int `json:"Serbia"`
				Lithuania                        int `json:"Lithuania"`
				France                           int `json:"France"`
				BosniaAndHerzegovina             int `json:"Bosnia and Herzegovina"`
				AustraliaAndNewZealand           int `json:"Australia and New Zealand"`
				Rwanda                           int `json:"Rwanda"`
				Bangladesh                       int `json:"Bangladesh"`
				Nicaragua                        int `json:"Nicaragua"`
				Barbados                         int `json:"Barbados"`
				Norway                           int `json:"Norway"`
				SouthernAfrica                   int `json:"Southern Africa"`
				SaintVincentAndTheGrenadines     int `json:"Saint Vincent and the Grenadines"`
				Botswana                         int `json:"Botswana"`
				Melanesia                        int `json:"Melanesia"`
				Macao                            int `json:"Macao"`
				DominicanRepublic                int `json:"Dominican Republic"`
				Denmark                          int `json:"Denmark"`
				Uganda                           int `json:"Uganda"`
				Mexico                           int `json:"Mexico"`
				Suriname                         int `json:"Suriname"`
				Greenland                        int `json:"Greenland"`
			} `json:"location"`
			Timezone struct {
				int                        `json:""`
				AmericaSaoPaulo            int `json:"America/Sao_Paulo"`
				AsiaVladivostok            int `json:"Asia/Vladivostok"`
				EuropeBerlin               int `json:"Europe/Berlin"`
				AfricaCairo                int `json:"Africa/Cairo"`
				EuropeMoscow               int `json:"Europe/Moscow"`
				PacificHonolulu            int `json:"Pacific/Honolulu"`
				AustraliaHobart            int `json:"Australia/Hobart"`
				EuropeLondon               int `json:"Europe/London"`
				AsiaBaghdad                int `json:"Asia/Baghdad"`
				AsiaShanghai               int `json:"Asia/Shanghai"`
				AmericaTijuana             int `json:"America/Tijuana"`
				AmericaManagua             int `json:"America/Managua"`
				AsiaYerevan                int `json:"Asia/Yerevan"`
				AsiaKamchatka              int `json:"Asia/Kamchatka"`
				AfricaHarare               int `json:"Africa/Harare"`
				AmericaNome                int `json:"America/Nome"`
				AsiaYakutsk                int `json:"Asia/Yakutsk"`
				AmericaChicago             int `json:"America/Chicago"`
				AmericaHalifax             int `json:"America/Halifax"`
				AmericaIndianaIndianapolis int `json:"America/Indiana/Indianapolis"`
				EuropeParis                int `json:"Europe/Paris"`
				PacificFiji                int `json:"Pacific/Fiji"`
				AsiaTehran                 int `json:"Asia/Tehran"`
				AmericaLaPaz               int `json:"America/La_Paz"`
				AsiaTashkent               int `json:"Asia/Tashkent"`
				AsiaBangkok                int `json:"Asia/Bangkok"`
				PacificMidway              int `json:"Pacific/Midway"`
				AmericaRecife              int `json:"America/Recife"`
				AmericaBuenosAires         int `json:"America/Buenos_Aires"`
				AustraliaAdelaide          int `json:"Australia/Adelaide"`
				AsiaYangon                 int `json:"Asia/Yangon"`
				AsiaKatmandu               int `json:"Asia/Katmandu"`
				AsiaAlmaty                 int `json:"Asia/Almaty"`
				AmericaPhoenix             int `json:"America/Phoenix"`
				EuropePrague               int `json:"Europe/Prague"`
				AmericaMexicoCity          int `json:"America/Mexico_City"`
				AsiaTbilisi                int `json:"Asia/Tbilisi"`
				AsiaJerusalem              int `json:"Asia/Jerusalem"`
				EuropeLisbon               int `json:"Europe/Lisbon"`
				AtlanticSouthGeorgia       int `json:"Atlantic/South_Georgia"`
				AsiaKarachi                int `json:"Asia/Karachi"`
				AustraliaPerth             int `json:"Australia/Perth"`
				AustraliaDarwin            int `json:"Australia/Darwin"`
				AsiaCalcutta               int `json:"Asia/Calcutta"`
				AmericaBogota              int `json:"America/Bogota"`
				AsiaKabul                  int `json:"Asia/Kabul"`
				AmericaNewYork             int `json:"America/New_York"`
				AtlanticAzores             int `json:"Atlantic/Azores"`
				AsiaKrasnoyarsk            int `json:"Asia/Krasnoyarsk"`
				EET                        int `json:"EET"`
				PacificAuckland            int `json:"Pacific/Auckland"`
				EuropeMinsk                int `json:"Europe/Minsk"`
				AfricaCasablanca           int `json:"Africa/Casablanca"`
				AmericaCaracas             int `json:"America/Caracas"`
				EuropeKiev                 int `json:"Europe/Kiev"`
				AsiaMagadan                int `json:"Asia/Magadan"`
				AmericaRegina              int `json:"America/Regina"`
				PacificGuam                int `json:"Pacific/Guam"`
				AsiaIrkutsk                int `json:"Asia/Irkutsk"`
				AfricaAlgiers              int `json:"Africa/Algiers"`
				AmericaStJohns             int `json:"America/St_Johns"`
				AmericaFortaleza           int `json:"America/Fortaleza"`
				AmericaDenver              int `json:"America/Denver"`
				AmericaIndianaKnox         int `json:"America/Indiana/Knox"`
				EtcUTC                     int `json:"Etc/UTC"`
				AustraliaSydney            int `json:"Australia/Sydney"`
				PacificApia                int `json:"Pacific/Apia"`
				AsiaTokyo                  int `json:"Asia/Tokyo"`
				AsiaIstanbul               int `json:"Asia/Istanbul"`
				AsiaOmsk                   int `json:"Asia/Omsk"`
				AustraliaBrisbane          int `json:"Australia/Brisbane"`
				AmericaLosAngeles          int `json:"America/Los_Angeles"`
				AsiaYekaterinburg          int `json:"Asia/Yekaterinburg"`
				EuropeAthens               int `json:"Europe/Athens"`
			} `json:"timezone"`
			ConnectPrice struct {
				Num4  int `json:"4"`
				Num6  int `json:"6"`
				Zero2 int `json:"0-2"`
			} `json:"connectPrice"`
			ContractToHire struct {
				True  int `json:"true"`
				False int `json:"false"`
			} `json:"contractToHire"`
			Categories      []interface{} `json:"categories"`
			Subcategories   []interface{} `json:"subcategories"`
			PaymentVerified struct {
				Num1 int `json:"1"`
			} `json:"paymentVerified"`
		} `json:"facets"`
		IsSearchWithEmptyParams bool `json:"isSearchWithEmptyParams"`
		CurrentQuery            struct {
			Sort string `json:"sort"`
		} `json:"currentQuery"`
		QueryParsedParams struct {
			Sort   string `json:"sort"`
			Paging string `json:"paging"`
		} `json:"queryParsedParams"`
		PageTitle      string `json:"pageTitle"`
		JobSearchError bool   `json:"jobSearchError"`
		RssLink        string `json:"rssLink"`
		AtomLink       string `json:"atomLink"`
		SmartSearch    struct {
			DownloadTeamApplication bool `json:"downloadTeamApplication"`
		} `json:"smartSearch"`
	} `json:"searchResults"`
	Filters struct {
		Q      string `json:"q"`
		Sort   string `json:"sort"`
		Skills struct {
			Name    string        `json:"name"`
			Label   string        `json:"label"`
			Options []interface{} `json:"options"`
		} `json:"skills"`
		Categories []struct {
			URLName       string `json:"urlName"`
			Value         string `json:"value"`
			Label         string `json:"label"`
			ActiveLabel   string `json:"activeLabel"`
			Checked       bool   `json:"checked"`
			Count         int    `json:"count"`
			Subcategories []struct {
				URLName     string `json:"urlName"`
				Value       string `json:"value"`
				Label       string `json:"label"`
				ActiveLabel string `json:"activeLabel"`
				Checked     bool   `json:"checked"`
				Count       int    `json:"count"`
			} `json:"subcategories"`
		} `json:"categories"`
		JobType struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			LoggingLabel       string `json:"loggingLabel"`
			Options            []struct {
				Value   string `json:"value"`
				Label   string `json:"label"`
				Checked bool   `json:"checked"`
				Count   int    `json:"count"`
			} `json:"options"`
		} `json:"jobType"`
		ContractorTier struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			LoggingLabel       string `json:"loggingLabel"`
			Options            []struct {
				Value   string `json:"value"`
				Label   string `json:"label"`
				Checked bool   `json:"checked"`
				Count   int    `json:"count"`
			} `json:"options"`
		} `json:"contractorTier"`
		ClientHires struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			LoggingLabel       string `json:"loggingLabel"`
			Options            []struct {
				Value   string `json:"value"`
				Label   string `json:"label"`
				Checked bool   `json:"checked"`
				Count   int    `json:"count"`
			} `json:"options"`
		} `json:"clientHires"`
		Proposals struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			LoggingLabel       string `json:"loggingLabel"`
			Options            []struct {
				Value       string `json:"value"`
				Label       string `json:"label"`
				Checked     bool   `json:"checked"`
				Count       int    `json:"count"`
				ActiveLabel string `json:"activeLabel"`
			} `json:"options"`
		} `json:"proposals"`
		Amount struct {
			Name                 string `json:"name"`
			Label                string `json:"label"`
			LoggingSublocation   string `json:"loggingSublocation"`
			DisabledJobTypeValue string `json:"disabledJobTypeValue"`
			DisabledMessage      string `json:"disabledMessage"`
			Options              []struct {
				Value   string `json:"value"`
				Label   string `json:"label"`
				Checked bool   `json:"checked"`
				Count   int    `json:"count"`
			} `json:"options"`
			CustomAmount struct {
				Value     interface{} `json:"value"`
				Label     string      `json:"label"`
				Checked   bool        `json:"checked"`
				CustomMin interface{} `json:"customMin"`
				CustomMax interface{} `json:"customMax"`
			} `json:"customAmount"`
		} `json:"amount"`
		Workload struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			LoggingLabel       string `json:"loggingLabel"`
			Options            []struct {
				Value   string `json:"value"`
				Label   string `json:"label"`
				Checked bool   `json:"checked"`
				Count   int    `json:"count"`
			} `json:"options"`
			DisabledJobTypeValue string `json:"disabledJobTypeValue"`
			DisabledMessage      string `json:"disabledMessage"`
		} `json:"workload"`
		DurationV3 struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			LoggingLabel       string `json:"loggingLabel"`
			Options            []struct {
				Value   string `json:"value"`
				Label   string `json:"label"`
				Checked bool   `json:"checked"`
				Count   int    `json:"count"`
			} `json:"options"`
			DisabledJobTypeValue string `json:"disabledJobTypeValue"`
			DisabledMessage      string `json:"disabledMessage"`
		} `json:"duration_v3"`
		PreviousClients struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			Value              string `json:"value"`
			Checked            bool   `json:"checked"`
			Count              int    `json:"count"`
		} `json:"previousClients"`
		PaymentVerified struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			Value              string `json:"value"`
			Checked            bool   `json:"checked"`
			Count              int    `json:"count"`
		} `json:"paymentVerified"`
		TalentClouds struct {
			Name               string        `json:"name"`
			Label              string        `json:"label"`
			LoggingSublocation string        `json:"loggingSublocation"`
			Options            []interface{} `json:"options"`
		} `json:"talentClouds"`
		UserLocationMatch struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			Value              string `json:"value"`
			Checked            bool   `json:"checked"`
			Count              int    `json:"count"`
		} `json:"userLocationMatch"`
		Occupation struct {
			Name               string        `json:"name"`
			Label              string        `json:"label"`
			LoggingSublocation string        `json:"loggingSublocation"`
			Options            []interface{} `json:"options"`
		} `json:"occupation"`
		OntologySkills struct {
			Name               string        `json:"name"`
			Label              string        `json:"label"`
			LoggingSublocation string        `json:"loggingSublocation"`
			Options            []interface{} `json:"options"`
		} `json:"ontologySkills"`
		HourlyRate struct {
			Name               string        `json:"name"`
			Label              string        `json:"label"`
			LoggingSublocation string        `json:"loggingSublocation"`
			Options            []interface{} `json:"options"`
			CustomHourlyRate   struct {
				Value     interface{} `json:"value"`
				Label     string      `json:"label"`
				Checked   bool        `json:"checked"`
				CustomMin interface{} `json:"customMin"`
				CustomMax interface{} `json:"customMax"`
			} `json:"customHourlyRate"`
		} `json:"hourly_rate"`
		FreelancersNeeded struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			LoggingLabel       string `json:"loggingLabel"`
			Options            []struct {
				Value   string `json:"value"`
				Label   string `json:"label"`
				Checked bool   `json:"checked"`
				Count   int    `json:"count"`
			} `json:"options"`
		} `json:"freelancers_needed"`
		Location struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			Options            []struct {
				UID          string `json:"uid"`
				Label        string `json:"label"`
				SubRegionUID string `json:"subRegionUid,omitempty"`
				RegionUID    string `json:"regionUid,omitempty"`
				Value        string `json:"value"`
				Type         string `json:"type"`
				Checked      bool   `json:"checked"`
				Count        int    `json:"count"`
			} `json:"options"`
		} `json:"location"`
		Timezones struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			Options            []struct {
				Value   string `json:"value"`
				Label   string `json:"label"`
				Checked bool   `json:"checked"`
				Count   int    `json:"count"`
			} `json:"options"`
		} `json:"timezones"`
		Connects struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			LoggingLabel       string `json:"loggingLabel"`
			Options            []struct {
				Value   string `json:"value"`
				Label   string `json:"label"`
				Checked bool   `json:"checked"`
				Count   int    `json:"count"`
			} `json:"options"`
		} `json:"connects"`
		ContractToHire struct {
			Name               string `json:"name"`
			Label              string `json:"label"`
			LoggingSublocation string `json:"loggingSublocation"`
			Value              string `json:"value"`
			Checked            bool   `json:"checked"`
			Count              int    `json:"count"`
		} `json:"contractToHire"`
	} `json:"filters"`
}
