package crypto

import (
	"cloud.google.com/go/firestore"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"encoding/asn1"
	"encoding/hex"
	"encoding/pem"
	"github.com/covidlocaldev/common/firebase"
	"github.com/covidlocaldev/common/logging"
	"google.golang.org/api/iterator"
	"io"
)

var firestoreClient = firebase.NewFirestoreClient()

func GenerateChatKey(chatID string, participantRefs []*firestore.DocumentRef) {
	ctx := context.Background()

	reader := rand.Reader

	chatAESKey, err := generateNewAESKey()
	if err != nil {
		logging.Error("error generating new AES key: %s", err)
		return
	}

	for _, participantRef := range participantRefs {
		participantPublicKeys, err2 := getParticipantDevicesPublicKeys(participantRef.ID)
		if err2 != nil {
			continue
		}

		for _, participantDevicePublicKey := range participantPublicKeys {
			//encryptedChatAESKeyBytes, err2 := rsa.EncryptPKCS1v15(reader, &participantDevicePublicKey.PublicKey, chatAESKey)
			encryptedChatAESKeyBytes, err2 := rsa.EncryptOAEP(sha1.New(), reader, &participantDevicePublicKey.PublicKey, chatAESKey, nil)
			if err2 != nil {
				logging.Warn("error occurred encrypting chat %s private key for participant %s: %s", chatID, participantRef.ID, err2)
				continue
			}

			chatParticipantPrivateKey := ChatParticipantPrivateKeyData{
				ParticipantUID:             participantRef.ID,
				DeviceToken:                participantDevicePublicKey.DeviceToken,
				ChatID:                     chatID,
				EncodedEncryptedAESKey:     hex.EncodeToString(encryptedChatAESKeyBytes),
			}

			_, _, err2 = firestoreClient.Collection("encryptedPrivateKeys").Add(ctx, chatParticipantPrivateKey)
			if err2 != nil {
				logging.Warn("error occurred storing chat %s private key for participant %s: %s", chatID, participantRef.ID, err2)
				continue
			}
		}
	}
}

func getParticipantDevicesPublicKeys(participantUID string) (participantPublicKeys []DevicePublicKey, err error) {
	participantDeviceSnapshots := firestoreClient.Collection("users").
		Doc(participantUID).
		Collection("devices").
		Documents(context.Background())

	for {
		participantDeviceSnapshot, err2 := participantDeviceSnapshots.Next()
		if err2 == iterator.Done {
			return
		} else if err2 != nil {
			logging.Warn("error occurred getting participant %s device: %s", participantUID, err2)
			continue
		}

		var participantDeviceData UserDeviceData
		err2 = participantDeviceSnapshot.DataTo(&participantDeviceData)
		if err2 != nil {
			logging.Warn("error occurred getting participant %s public chatKey data: %s", participantUID, err2)
			continue
		}

		if !participantDeviceData.IsMasterDevice && participantDeviceData.AuthorizedBy == nil {
			continue
		}

		block, _ := pem.Decode([]byte(participantDeviceData.PublicKey))

		var participantDevicePublicKey rsa.PublicKey
		_, err2 = asn1.Unmarshal(block.Bytes, &participantDevicePublicKey)
		if err2 != nil {
			logging.Warn("error occurred parsing participant %s public key: %s", participantUID, err2)
			continue
		}

		devicePublicKey := DevicePublicKey{
			DeviceToken: participantDeviceData.DeviceToken,
			PublicKey:   participantDevicePublicKey,
		}

		participantPublicKeys = append(participantPublicKeys, devicePublicKey)
	}
}

func generateNewAESKey() ([]byte, error) {
	key := make([]byte, 32)

	if _, err := io.ReadFull(rand.Reader, key); err != nil {
		logging.Error(err.Error())
		return nil, err
	}

	return key, nil
}
