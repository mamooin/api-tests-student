package advertisementSearch

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"

	"api-tests-template/internal/assertions"
	advList "api-tests-template/internal/managers/advertisements"
	"api-tests-template/internal/managers/auth"
	"api-tests-template/internal/managers/auth/models"
	"api-tests-template/internal/utils"

	base "api-tests-template/tests"
)

// Объявление должно существовать на стенде до запуска тестов.
const (
	existingAdTitle = "Продам велосипед горный"
	searchBeginning = "Продам"
	searchMiddle    = "велосипед"
	searchEnd       = "горный"
)

type TestSuite struct {
	suite.Suite
	loginData models.LoginOkResponse
}

func TestSuiteRun(t *testing.T) {
	suite.Run(t, &TestSuite{})
}

func (s *TestSuite) SetupSuite() {
	base.SetupSuite()

	base.Precondition("Проверяем, что константы для поиска объявления заданы")
	require.NotEmpty(s.T(), existingAdTitle, "Задайте existingAdTitle: полное название существующего объявления")
	require.NotEmpty(s.T(), searchBeginning, "Задайте searchBeginning: подстроку из начала названия")
	require.NotEmpty(s.T(), searchMiddle, "Задайте searchMiddle: подстроку из середины названия")
	require.NotEmpty(s.T(), searchEnd, "Задайте searchEnd: подстроку из конца названия")

	base.Precondition("Авторизация пользователя с кредами из переменных окружения и получение его параметров")
	s.loginData = auth.Login(s.T(), os.Getenv("TEST_LOGIN"), os.Getenv("TEST_PASSWORD"))
}

func (s *TestSuite) TearDownSuite() {
	base.TearDownSuite()
}

// Positive: Проверяем поиск существующего объявления по подстроке в названии
func (s *TestSuite) TestSearchByTitle() {
	list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, searchMiddle)

	s.Run("В ответе есть хотя бы одно объявление", func() {
		require.NotEmpty(s.T(), list.Items)
	})

	s.Run("Среди найденных есть наше объявление", func() {
		assertions.ContainsAdvertisementByTitle(s.T(), list.Items, existingAdTitle)
	})
}

// Positive: Проверяем поиск по началу, середине, концу названия и в разных регистрах
func (s *TestSuite) TestSearchByTitle_Substring() {
	cases := []struct {
		name   string
		search string
	}{
		{"Начало", searchBeginning},
		{"Середина", searchMiddle},
		{"Конец", searchEnd},
		{"Нижний регистр", strings.ToLower(searchMiddle)},
		{"Верхний регистр", strings.ToUpper(searchMiddle)},
		{"Смешанный регистр", utils.MixCase(searchMiddle)},
	}

	for _, tc := range cases {
		s.Run(tc.name, func() {
			list := advList.GetAdvertisementsList(s.T(), s.loginData.Token, tc.search)
			require.NotEmpty(s.T(), list.Items)
			assertions.ContainsAdvertisementByTitle(s.T(), list.Items, existingAdTitle)
		})
	}
}
