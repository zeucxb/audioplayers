package main

import (
	"github.com/go-flutter-desktop/go-flutter"
  "github.com/go-flutter-desktop/plugins/shared_preferences"
  "github.com/go-flutter-desktop/plugins/path_provider"
  "github.com/go-flutter-desktop/plugins/audioplayers"
)

var options = []flutter.Option{
	flutter.WindowInitialDimensions(640, 420),

  flutter.AddPlugin(&shared_preferences.SharedPreferencesPlugin{
    VendorName: "Fireslime",
    ApplicationName: "Audioplayers",
  }),

  flutter.AddPlugin(&path_provider.PathProviderPlugin{
    VendorName: "Fireslime",
    ApplicationName: "Audioplayers",
  }),

  flutter.AddPlugin(&audioplayers.AudioplayersPlugin{}),
}


