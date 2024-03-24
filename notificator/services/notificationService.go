package services

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io"
	"net/http"
	"net/url"
	"notificator/events"
	"notificator/models"
	"notificator/repositories"
	"os"
	"strconv"
	"strings"
)

func mapMsgToEntity(body []byte) events.UserCreated {
	var event events.UserCreated
	err := json.Unmarshal(body, &event)
	if err != nil {
		log.Errorf("Error unmarshalling notification: %v", err)
	}
	return event
}

func getNotificationProfile(userId uint) models.NotificationProfile {
	notificationProfile, err := repositories.GetNotificationProfileByUserId(userId)
	if err != nil {
		notificationProfile, err = repositories.CreateNotificationProfile(models.NotificationProfile{UserId: userId})
		if err != nil {
			log.Errorf("Error creating notification profile: %v", err)
			return models.NotificationProfile{}
		}
	}
	return notificationProfile

}

func ProcessNotification(event []byte) {
	message := mapMsgToEntity(event)
	notificationProfile := getNotificationProfile(message.UserId)
	sendNotification(notificationProfile)
}

func sendNotification(profile models.NotificationProfile) {
	log.Infof("Sending notification to user %v", profile.UserId)
	user, err := getUserData(profile.UserId)
	if err != nil {
		log.Errorf("Error fetching user data: %v", err)
		return
	}
	if profile.DefaultNotificationChannel == models.SMS {
		sendSms(user.PhoneNumber, "Hello from notificator!")
	} else if profile.DefaultNotificationChannel == models.EMAIL {
		sendEmail("test", "Hello")
	} else if profile.DefaultNotificationChannel == models.PUSH {
		sendPush("test", "Hello")
	}

}

func sendSms(phoneNumber string, message string) {
	log.Infof("Sending SMS to %v", phoneNumber)
	data := url.Values{}
	data.Set("key", os.Getenv("SMS_KEY"))
	data.Set("password", os.Getenv("SMS_PASSWORD"))
	data.Set("from", os.Getenv("SMS_SENDER"))
	data.Set("to", phoneNumber)
	data.Set("msg", message)

	client := &http.Client{}
	apiUrl := os.Getenv("SMS_API_URL") + "/sms"
	req, err := http.NewRequest("POST", apiUrl, strings.NewReader(data.Encode()))
	if err != nil {
		log.Errorf("Error creating new request: %v", err)
		return
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("Error while making the request %v", err)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Errorf("Error while sending SMS: %v", resp.Status)
	}
}

func sendEmail(email string, message string) {
	log.Infof("Sending email to %v", email)
}

func sendPush(pushToken string, message string) {
	log.Infof("Sending push to %v", pushToken)
}

func getUserData(userId uint) (models.User, error) {
	var client = http.Client{}
	url := os.Getenv("USER_SERVICE_URL") + "/user/" + strconv.Itoa(int(userId))
	log.Info("Fetching user data from: ", url)
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Errorf("Error creating new request: %v", err)
		return models.User{}, err
	}

	response, err := client.Do(request)
	if err != nil {
		log.Errorf("Error while making the request %v", err)
		return models.User{}, err
	}

	if response.StatusCode != 200 {
		log.Errorf("Error while fetching user: %v , ID: %v", response.Status, userId)
		return models.User{}, err
	}

	var user models.User

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Errorf("Error closing response body: %v", err)
		}
	}(response.Body)
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Errorf("Error reading response body: %v", err)
		return models.User{}, err
	}

	err = json.Unmarshal(body, &user)
	if err != nil {
		log.Errorf("Error while unmarshalling response body: %v", err)
		return models.User{}, err
	}
	return user, nil
}
