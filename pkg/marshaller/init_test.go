package marshaller

import (
	"context"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/configuration"
	"github.com/SENERGY-Platform/mgw-external-task-worker/pkg/devicerepo"
	"github.com/SENERGY-Platform/service-commons/pkg/cache"
	"github.com/SENERGY-Platform/service-commons/pkg/cache/fallback"
	"log"
	"testing"
)

func TestSNRGY3713(t *testing.T) {
	t.Skip("used to test SNRGY-3713; set config.AuthUserName and config.AuthPassword to enable this test")
	config, err := configuration.Load("../../config.json")
	if err != nil {
		t.Error(err)
		return
	}
	config.AuthEndpoint = "https://auth.senergy.infai.org"
	config.AuthUserName = ""
	config.AuthPassword = ""
	config.FallbackFile = t.TempDir() + "/fallback.json"
	config.DeviceRepoUrl = "https://api.senergy.infai.org/device-repository"
	c, err := cache.New(cache.Config{FallbackProvider: fallback.NewProvider(config.FallbackFile)})
	if err != nil {
		log.Fatal(err)
		return
	}
	iotProvider := &devicerepo.Provider{Config: config, Cache: c}
	factory := Factory{Config: config, DeviceRepo: iotProvider.GetImpl()}
	_ = factory.New(context.Background(), "")
}
