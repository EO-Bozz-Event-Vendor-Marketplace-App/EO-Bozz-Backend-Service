package thirdparty

import (
	"capstone-alta1/utils/helper"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

// Retrieve a token, saves the token, then returns the generated client.
func getClient(config *oauth2.Config) *http.Client {
	// The file token.json stores the user's access and refresh tokens, and is
	// created automatically when the authorization flow completes for the first
	// time.
	tokFile := "token.json"
	tok, err := tokenFromFile(tokFile)
	if err != nil {
		tok = getTokenFromWeb(config)
		saveToken(tokFile, tok)
	}
	return config.Client(context.Background(), tok)
}

// Request a token from the web, then returns the retrieved token.
func getTokenFromWeb(config *oauth2.Config) *oauth2.Token {
	authURL := config.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, err := config.Exchange(context.TODO(), authCode)
	if err != nil {
		log.Fatalf("Unable to retrieve token from web: %v", err)
	}
	return tok
}

// Retrieves a token from a local file.
func tokenFromFile(file string) (*oauth2.Token, error) {
	f, err := os.Open(file)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	tok := &oauth2.Token{}
	err = json.NewDecoder(f).Decode(tok)
	return tok, err
}

// Saves a token to a file path.
func saveToken(path string, token *oauth2.Token) {
	fmt.Printf("Saving credential file to: %s\n", path)
	f, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		log.Fatalf("Unable to cache oauth token: %v", err)
	}
	defer f.Close()
	json.NewEncoder(f).Encode(token)
}

func Calendar(email, eventName, startDate, endDate, address string) (string, error) {
	event := &calendar.Event{
		Summary:     eventName,
		Location:    address,
		Description: "Contact Admin: eo.bozz01@gmail.com",
		Start: &calendar.EventDateTime{
			Date: startDate,

			TimeZone: "Asia/Jakarta",
		},
		End: &calendar.EventDateTime{
			Date:     endDate,
			TimeZone: "Asia/Jakarta",
		},
		//Recurrence: []string{"RRULE:FREQ=DAILY;COUNT=2"},
		Attendees: []*calendar.EventAttendee{
			{Email: email},
		},
	}
	ctx := context.Background()
	// b, err := ioutil.ReadFile("credentials.json")
	// if err != nil {
	// 	log.Fatalf("Unable to read client secret file: %v", err)
	// }

	client_id := os.Getenv("GOOGLE_OAUTH_CLIENT_ID")
	project := os.Getenv("GOOGLE_PROJECT_ID")
	secret := os.Getenv("GOOGLE_OAUTH_CLIENT_SECRET")
	b := `{"web":{"client_id":"` + client_id + `","project_id":"` + project + `","auth_uri":"https://accounts.google.com/o/oauth2/auth","token_uri":"https://oauth2.googleapis.com/token","auth_provider_x509_cert_url":"https://www.googleapis.com/oauth2/v1/certs","client_secret":"` + secret + `","redirect_uris":["http://localhost"]}}`

	bt := []byte(b)

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(bt, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	// calendarId := "primary"
	// event_notification := srv.Events.Insert(calendarId, event).SendUpdates("all")
	// event, err = event_notification.Do()
	// if err != nil {
	// 	helper.LogDebug("Thirdpary Calendar Err Create Event = ", err)
	// 	log.Fatal("Unable to create event")
	// }

	calendarId := "primary"
	event, err = srv.Events.Insert(calendarId, event).SendUpdates("all").Do()
	if err != nil {
		helper.LogDebug("Thirdpary Calendar Err Create Event = ", err)
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", event.HtmlLink)

	return event.HtmlLink, nil
}
