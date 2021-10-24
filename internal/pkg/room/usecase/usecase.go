package product

import (
	"context"
	"encoding/json"
	"fmt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
	"io/ioutil"
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/models"
	"github.com/aridae/web-dreamit-api-based-labs/internal/pkg/room"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
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


type RoomUseCase struct {
	RoomRepo room.Repository
}

func (r RoomUseCase) DeleteRoomBooking(userId uint64, eventId int64) error {
	return r.RoomRepo.DeleteRoomBooking(userId, eventId)
}

func (r RoomUseCase) MyRoomBooking(userId uint64) ([]models.Booking, error) {
	return r.RoomRepo.MyRoomBooking(userId)
}

func (r RoomUseCase) AddRoomBooking(roomId int64, event models.Event) (int64, error) {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		fmt.Println(f.Name())
	}

	result, err := r.RoomRepo.AddRoomBooking(roomId, event)
	if err != nil {
		return result, err
	}
	ctx := context.Background()
	b, err := ioutil.ReadFile("credentials.json")
	if err != nil {
		log.Fatalf("Unable to read client secret file: %v", err)
	}

	// If modifying these scopes, delete your previously saved token.json.
	config, err := google.ConfigFromJSON(b, calendar.CalendarScope)
	if err != nil {
		log.Fatalf("Unable to parse client secret file to config: %v", err)
	}
	client := getClient(config)

	srv, err := calendar.NewService(ctx, option.WithHTTPClient(client))
	if err != nil {
		log.Fatalf("Unable to retrieve Calendar client: %v", err)
	}

	calendarEvent := &calendar.Event{
		Summary: event.Title,
		Start: &calendar.EventDateTime{
			DateTime: event.Start,
			TimeZone: "Europe/Moscow",
		},
		End: &calendar.EventDateTime{
			DateTime: event.End,
			TimeZone: "Europe/Moscow",
		},
		Creator: &calendar.EventCreator{
			DisplayName: strconv.Itoa(int(event.Author)),
		},
	}

	calendarId := "primary"
	calendarEvent, err = srv.Events.Insert(calendarId, calendarEvent).Do()
	if err != nil {
		log.Fatalf("Unable to create event. %v\n", err)
	}
	fmt.Printf("Event created: %s\n", calendarEvent.HtmlLink)

	return result, nil
}

func (r RoomUseCase) GetAllRooms() ([]models.Room, error) {
	return r.RoomRepo.GetAllRooms()
}

func (r RoomUseCase) GetRoomCalendar(roomId int64) ([]models.Event, error) {
	return r.RoomRepo.GetRoomCalendarById(roomId)
}

func (r RoomUseCase) GetRoomSchedule(roomId int64, date time.Time) (*models.Schedule, error) {
	panic("implement me")
}

func (r RoomUseCase) UpdateRoomBooking(roomId int64, event models.Event) error {
	return r.RoomRepo.UpdateRoomBookingById(roomId, event)
}

func NewUseCase(roomRepo room.Repository) room.UseCase {
	return &RoomUseCase{
		RoomRepo: roomRepo,
	}
}
