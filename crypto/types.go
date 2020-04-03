package crypto

import (
	"cloud.google.com/go/firestore"
	"crypto/rsa"
)

type UserDeviceData struct {
	DeviceToken    string                 `firestore:"token"`
	PublicKey      string                 `firestore:"publicKey"`
	IsMasterDevice bool                   `firestore:"isMasterDevice,omitempty"`
	AuthorizedBy   *firestore.DocumentRef `firestore:"authorizedBy,omitempty"`
}

type ChatPublicKeyData struct {
	ChatID    string `firestore:"chatID"`
	PublicKey string `firestore:"publicKey"`
}

type ChatParticipantPrivateKeyData struct {
	ParticipantUID             string `firestore:"participantUID"`
	DeviceToken                string `firestore:"deviceToken"`
	ChatID                     string `firestore:"chatID"`
	ChatPublicKey              string `firestore:"chatPublicKey"`
	EncodedEncryptedPrivateKey string `firestore:"encodedEncryptedPrivateKey"`
	EncodedEncryptedAESKey     string `firestore:"encodedEncryptedAESKey"`
	EncodedAESIV               string `firestore:"encodedAESIV"`
}

type DevicePublicKey struct {
	DeviceToken string
	PublicKey   rsa.PublicKey
}
