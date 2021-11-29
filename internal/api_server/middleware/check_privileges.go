package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/aridae/web-dreamit-api-based-labs/internal/api_server/apimodels"
	sessioncont "github.com/aridae/web-dreamit-api-based-labs/internal/controllers/session_controller"
	"github.com/aridae/web-dreamit-api-based-labs/internal/server/tools/http_utils"
)

// я пока сделала и енва - вообще так делать нехорошо
// просто для контейнера я смогу в компознике установить енв
// а конфиг файл подменить не смогу, поэтому пока так
const (
	ENV_IS_MASTER_VAR = "isMaster"
	NOT_MASTER_ERROR  = "Access forbidden"
)

var (
	isMaster = getIsMasterVar()
)

type AccessHandler struct {
	SessionController *sessioncont.SessionController
}

func getIsMasterVar() bool {
	isMaster, exists := os.LookupEnv(ENV_IS_MASTER_VAR)
	if !exists {
		fmt.Println("env var not found")
		return false
	}

	isMaster_f, err := strconv.ParseBool(isMaster)
	if err != nil {
		fmt.Println("invalid isMaster flag")
		return false
	}

	return isMaster_f
}

// читаем из конфига параметр -- мастер или не мастер
// если мастер пропускаем все запросы
// а если не мастер -- только гет, иначе 503
func (h *AccessHandler) CheckAccess(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if isMaster {
			// все ок -- пропускаем дальше -- мастер может все
			next.ServeHTTP(w, r)
			return
		}

		if r.Method == "GET" {
			// все ок -- пропускаем дальше -- слейв может только отхватить
			next.ServeHTTP(w, r)
			return
		}

		// иначе кидаем 503 - сервис не доступен
		_, err := h.SessionController.ExtractAccessTokenMetadata(r)
		if err != nil {
			http_utils.SetJSONResponse(w,
				apimodels.MessageResponse{
					Message: NOT_MASTER_ERROR,
				}, http.StatusServiceUnavailable)
			return
		}
	})
}
