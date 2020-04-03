package oauth

import (
	"cloud.google.com/go/firestore"
	"context"
	"fmt"
	"github.com/covidlocaldev/common/constants"
	"github.com/covidlocaldev/common/firebase"
	"github.com/covidlocaldev/common/keys"
	"github.com/covidlocaldev/common/logging"
	"golang.org/x/oauth2"
	"google.golang.org/api/iterator"
	"log"
	"net/http"
)

var (
	firestoreClient               = firebase.NewFirestoreClient()
	googleOAuthPractitionerConfig = &oauth2.Config{
		ClientID:     keys.GoogleOAuthKeys.DoctorClientID,
		ClientSecret: keys.GoogleOAuthKeys.DoctorClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.google.com/o/oauth2/auth",
			TokenURL:  "https://oauth2.googleapis.com/token",
			AuthStyle: 0,
		},
	}
	googleOAuthPatientConfig = &oauth2.Config{
		ClientID:     keys.GoogleOAuthKeys.PatientClientID,
		ClientSecret: keys.GoogleOAuthKeys.PatientClientSecret,
		Endpoint: oauth2.Endpoint{
			AuthURL:   "https://accounts.google.com/o/oauth2/auth",
			TokenURL:  "https://oauth2.googleapis.com/token",
			AuthStyle: 0,
		},
	}
)

type storedOAuthToken struct {
	Token    oauth2.Token
	Provider string
}

func RegisterUserOAuthToken(userID string, userOAuthCode string) (err error) {
	ctx := context.Background()

	oauthConfig := getUserOAuthConfig(userID)

	token, err := oauthConfig.Exchange(ctx, userOAuthCode, oauth2.AccessTypeOffline)
	if err != nil {
		log.Printf("error exchanging oauth token: %s", err)
		return
	}

	// save oauth token to user collection
	userTokensRef := firestoreClient.Collection(fmt.Sprintf("users/%s/oauthTokens", userID))

	storedTokenSnapshot, err := userTokensRef.
		Where("provider", "==", "google").
		Documents(ctx).
		Next()

	if err != nil {
		// no stored token
		oauthToken := storedOAuthToken{
			Token:    *token,
			Provider: "google",
		}

		_, _, err = userTokensRef.Add(ctx, oauthToken)
		if err != nil {
			log.Printf("error storing google oauth token: %s", err)
		}
	} else {
		// update stored token
		update := firestore.Update{Path: "Token", Value: *token}

		_, err = storedTokenSnapshot.Ref.Update(ctx, []firestore.Update{update})
		if err != nil {
			log.Printf("error updating stored google oauth token: %s", err)
		}
	}

	return
}

func GetGoogleUserOAuthClient(userID string) (client *http.Client, err error) {
	ctx := context.Background()

	oauthConfig := getUserOAuthConfig(userID)

	if oauthConfig == nil {
		return nil, nil
	}

	userOAuthCollectionRef := firestoreClient.Collection(fmt.Sprintf("users/%s/oauthTokens", userID))

	var oauthToken storedOAuthToken

	userOAuth, err := userOAuthCollectionRef.
		Where("Provider", "==", "google").
		Documents(ctx).
		Next()

	if err == iterator.Done {
		// no stored token was found

		return nil, nil

	} else if err != nil {
		logging.Warn("an error occurred getting user %s oauth: %s", userID, err)
		return
	} else {
		// stored token was found

		err2 := userOAuth.DataTo(&oauthToken)
		if err2 != nil {
			logging.Warn("error parsing user oauth: %s", err)
			return
		}
	}

	return oauth2.NewClient(ctx, oauthConfig.TokenSource(ctx, &oauthToken.Token)), nil
}

func getUserOAuthConfig(userID string) (config *oauth2.Config) {
	userSnapshot, err := firestoreClient.Doc(fmt.Sprintf("users/%s", userID)).Get(context.Background())
	if err != nil {
		log.Printf("error occurred getting user snapshot: %s", err)
		return
	}

	role, _ := userSnapshot.DataAt("role")

	if role == nil {
		logging.Warn("no role found for user %s", userID)
		return nil
	} else if role.(string) == constants.RolePractitioner {
		return googleOAuthPractitionerConfig
	} else {
		return googleOAuthPatientConfig
	}
}