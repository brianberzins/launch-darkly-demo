package main

import (
	"time"
	"os"
	"fmt"

	launchDarkly "gopkg.in/launchdarkly/go-client.v4"
)

const launchDarklySdkEnvironmentVariable = "LAUNCH_DARKLY_SDK_KEY"

func createUser() launchDarkly.User {
	key := "brian.berzins@murasaki.com"
	anonymous := true

	return launchDarkly.User{
		Key: &key,
		Anonymous: &anonymous,
	}
}

func main() {
	// grab our sdk key from the environment
	sdkKey, present := os.LookupEnv(launchDarklySdkEnvironmentVariable)
	if !present {
		panic(fmt.Sprintf("required environment variable %s is not specified", launchDarklySdkEnvironmentVariable))
	}

	// create the launch darkly client
	client, err := launchDarkly.MakeClient(sdkKey, time.Second)
	defer client.Close()
	if err != nil {
		panic(err)
	}

	// create user
	user := createUser()
	displayHostname, err := client.BoolVariation("display-hostname", user, true)

	if displayHostname {
		// application code to show the feature
		fmt.Println("Showing your feature to " + *user.Key)
	} else {
		// the code to run if the feature is off
		fmt.Println("Not showing your feature to " + *user.Key)
	}

}
