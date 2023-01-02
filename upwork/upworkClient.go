package upwork

import (
	"fmt"
	"scrapers/network"
)

type Upwork struct {
	UpworkHttpClient *network.Client
}

type UrlArgs struct {
	Page     int
	Per_Page int
	Query    string
}

func (u Upwork) ConstructUrl(args UrlArgs) string {
	url := "https://www.upwork.com/ab/jobs/search/url?q=%s&per_page=%d&sort=recency&payment_verified=1&page=%d"
	return fmt.Sprintf(url, args.Query, args.Per_Page, args.Page)
}

func (u Upwork) SendRequest(url string) (string, error) {
	Upclient := u.UpworkHttpClient
	resp, err := Upclient.GetRequest(url)

	if err != nil {
		return "", err
	}
	return resp.Body, nil

}

func InitUpwork() *Upwork {
	headers := map[string]string{
		"authority":       "www.upwork.com",
		"accept":          "application/json, text/plain",
		"accept-language": "en",
		"cache-control":   "no-cache",
		"pragma":          "no-cache",
		"referer":         "https://www.upwork.com/search/jobs/url?per_page=10&sort=recency",
		"sec-fetch-site":  "same-origin",
		"sec-gpc":         "1",

		"vnd-eo-parent-span-id": "b8f3fe84-2aa4-4b9e-b750-6dc0e6e169a9",
		"vnd-eo-span-id":        "0d5bed38-342c-4b44-a885-b9c3de71a32a",
		"x-odesk-user-agent":    "oDesk LM",
		"x-requested-with":      "XMLHttpRequest",
		"authorization":         "Bearer oauth2v2_029104aca6380322c31d5ff1beb45bdc",
		"cookie":                "visitor_id=84.239.14.170.1672593945152000; lang=en; visitor_gql_token=oauth2v2_73062637ed25bfaab8a975d46e74a2a6; cookie_prefix=; cookie_domain=.upwork.com; __cf_bm=CeIOv6_PAxdKYRcfxgS.0i5S6T7H_FtSQJIRJ7TfxZE-1672593945-0-AWkqgfz2NFzVJjyMCtqW9XH5v601Yw3Faj7Rnd/lalOGPrlRTLVeEMwwxlSd8rxVIx7Gxmeil0zA3eoueK2DafA=; __cfruid=ea76c3d0a524b002bd34afa80c598fda56e6ec51-1672593945; lang=en; _sp_ses.2a16=*; G_ENABLED_IDPS=google; _cq_duid=1.1672593953.libVfRjYZI9g1VVb; _cq_suid=1.1672593953.iSaHfjf7kYLkR6pW; OptanonAlertBoxClosed=2023-01-01T17:26:03.771Z; odesk_signup.referer.raw=https://www.upwork.com/ab/account-security/login; device_view=full; recognized=66e1fe97; console_user=66e1fe97; user_uid=1561687738217156608; master_access_token=fff79f08.oauth2v2_96a3e8bf07ca7af5426f97ae5f3462a8; oauth2_global_js_token=oauth2v2_029104aca6380322c31d5ff1beb45bdc; SZ=8d599fc6dcae93ee0a7d31376c071dd5c53432dd15d764e8bcc69ecfa8c138aa; company_last_accessed=d1020106921; current_organization_uid=1561687738217156609; user_oauth2_slave_access_token=fff79f08.oauth2v2_96a3e8bf07ca7af5426f97ae5f3462a8:1561687738217156608.oauth2v2_11b83ba7d6c21a48610deaa621b1763b; visitor_account_security_token=oauth2v2_f14381d2a80fa116031ad0a28ea995dd; XSRF-TOKEN=3d96f0d3ccf4b54113e78176144f442b; jsTSize_1561687738217156608=2; channel=other; visitor_ui_gql_token_jobsearch=oauth2v2_492e535ef89cfea762ce623120e9f3ce; forterToken=964af2adf2e946bfa97715b691a68da4_1672594152809_30_UAS_14ck; OptanonConsent=isGpcEnabled=1&datestamp=Sun+Jan+01+2023+22:29:15+GMT+0500+(Pakistan+Standard+Time)&version=6.39.0&isIABGlobal=false&hosts=&consentId=50406b2b-45ad-45e3-a842-f3a0feed1af9&interactionCount=1&landingPath=NotLandingPage&groups=C0001:1,C0002:0,C0003:0,C0004:0&geolocation=RO;&AwaitingReconsent=false; _sp_id.2a16=afb5f089-5283-4f0c-a133-b22247a07924.1672593950.1.1672594164..0587ef56-a7dc-4f9e-bf69-87dc9b20091c..473139f6-344e-4278-8d49-d8d489e57cfb.1672594154654.8; AWSALB=hklmGZU3Xl6e/0vh9R+DEs8vb5000Uvr68FRLtAMSg8Rg6zJNTEB6uCIalp51mRkZoEgxVAUWfVK0aEtAotogkeLDPWDCv6VKFvJ/PlsbWBYTuNzvfDgC071eRXa; AWSALBCORS=hklmGZU3Xl6e/0vh9R+DEs8vb5000Uvr68FRLtAMSg8Rg6zJNTEB6uCIalp51mRkZoEgxVAUWfVK0aEtAotogkeLDPWDCv6VKFvJ/PlsbWBYTuNzvfDgC071eRXa; spt=37d89b33-2cbd-43c3-a090-aa585d1f4300; enabled_ff=CI9570Air2Dot5,OTBnrOn,!MP16400Air3Migration,!air2Dot76Qt,air2Dot76,!CI10270Air2Dot5QTAllocations,!CI10857Air3Dot0,!CI12577UniversalSearch,CI11132Air2Dot75,SSINav",
	}
	client := network.InitClient(headers)
	upwork := Upwork{
		UpworkHttpClient: client,
	}

	return &upwork

}
