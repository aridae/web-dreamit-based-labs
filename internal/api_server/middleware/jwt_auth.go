package middleware

import (
	"net/http"
	"strings"

	"github.com/aridae/web-dreamit-api-based-labs/internal/api_server/apimodels"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	"github.com/aridae/web-dreamit-api-based-labs/internal/domain"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/tools/http_utils"
)

type JWTHandler struct {
	SessionController *sessioncont.SessionController
}

// мотивация: у нас бизнес-клиент и БОЛЬШАЯ часть методов
// требует авторизации за некоторыми исключениями: собственно логин и регистрация
// и пара методов доступа к коллекциям просто, чтобы показать разницу с jwt/без jwt
// поэтому прописать авторизацию как мидлвар будет проще, чем следить за всеми методами
// потому что я уже запуталась где я проверяю токен, а где забыла
func (h *JWTHandler) JWTAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// возможно у сваггера должен быть тоже отдельный мультиплексор
		// потому что он вообще не относится к апи как таковому
		if strings.Contains(r.URL.Path, "swagger") || strings.Contains(r.URL.Path, "login") || strings.Contains(r.URL.Path, "signup") {
			next.ServeHTTP(w, r)
			return
		}

		// список допустимых строк в конфиг

		// проверяем, есть ли в заголовочнике токен и валидный ли он
		_, err := h.SessionController.ExtractAccessTokenMetadata(r)
		if err != nil {
			http_utils.SetJSONResponse(w,
				apimodels.MessageResponse{
					Message: domain.FAILED_TO_EXTACT_TOKEN,
				}, http.StatusUnauthorized)
			return
		}

		// все ок -- пропускаем дальше
		next.ServeHTTP(w, r)
	})
}
