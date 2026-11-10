package advertisements

import (
	"encoding/json"
	"io"
	"testing"

	"github.com/stretchr/testify/require"

	"api-tests-template/internal/client/http/advertisements"
	"api-tests-template/internal/managers/advertisements/models"
)

// GetAdvertisementsList возвращает список объявлений, найденных по подстроке в названии
func GetAdvertisementsList(t *testing.T, token, text string) models.AdvertisementsList {
	advertisementsRawResponse := advertisements.HttpGetAdvertisements(t, token, text)
	defer advertisementsRawResponse.Body.Close()

	respArray, errResponse := io.ReadAll(advertisementsRawResponse.Body)
	require.NoError(t, errResponse)

	advertisementsResponse := models.AdvertisementsList{}
	err := json.Unmarshal(respArray, &advertisementsResponse)
	require.NoError(t, err)

	return advertisementsResponse
}
