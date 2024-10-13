package metrics

import (
	"regexp"
	"strconv"

	"time"

	"github.com/gin-gonic/gin"
)

// path иногда может быть /api/v1/getPost/1, а иногда /api/v1/getPost/2. Поэтому в ручках
// нужно сделать не так api.GET("/getPost/:id", h.GetPostById) а через ?, как в
// api.GET("/getPosts", h.GetPosts)

// Если это не сделать то в графане будет графки на каждый /api/v1/getPost/1 /2 /3 /4

var pathPattern = regexp.MustCompile(`^/api/v1/getPost/\d+$`)

func TrackMetrics() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		startTime := time.Now()

		ctx.Next()

		status := ctx.Writer.Status()

		path := ctx.Request.URL.Path
		method := ctx.Request.Method

		// Обработка динамических частей пути
		if pathPattern.MatchString(path) {
			path = "/api/v1/getPost"
		}

		if path == "/metrics" {
			// Если путь /metrics, увеличиваем только HttpRequestPath
			HttpRequestPath.WithLabelValues(path).Inc()
		} else {
			// Для всех остальных путей увеличиваем HttpRequestCount и HttpRequestPath
			HttpRequestCount.Inc()
			HttpRequestPath.WithLabelValues(path).Inc()

			// Запись времени выполнения запроса в гистограмму
			HttpRequestDuration.WithLabelValues(path).Observe(time.Since(startTime).Seconds())
		}

		if status >= 400 {
			HttpErrorCount.WithLabelValues(path, method, strconv.Itoa(status)).Inc()
		}
	}
}
