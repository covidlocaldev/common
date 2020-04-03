package keys

var TwilioKeys = initTwilioKeys()
var InternalServiceKeys = initInternalServiceKeys()
var AdminKeys = initAdminKeys()
var GoogleOAuthKeys = initGoogleAuthKeys()
var GoogleApiKeys = initGoogleApiKeys()

func initTwilioKeys() TwilioKey {
	var keys TwilioKey
	Load(&keys, "stage")

	return keys
}

func initInternalServiceKeys() InternalServiceKey {
	var keys InternalServiceKey
	Load(&keys, "stage")

	return keys
}

func initAdminKeys() AdminKey {
	var keys AdminKey
	Load(&keys, "stage")

	return keys
}

func initGoogleAuthKeys() GoogleOAuthKey {
	var keys GoogleOAuthKey
	Load(&keys, "stage")

	return keys
}

func initGoogleApiKeys() GoogleApiKey {
	var keys GoogleApiKey
	Load(&keys, "stage")

	return keys
}
