package middleware

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"runtime/debug"
	"slices"
	"strings"

	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/umefy/godash/jsonkit"
	"github.com/umefy/godash/logger"
)

func Recover(logger *logger.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if rec := recover(); rec != nil {
					if rec == http.ErrAbortHandler {
						// we don't recover http.ErrAbortHandler so the response
						// to the client is aborted, this should not be logged
						panic(rec)
					}

					// Get and format stack trace
					chiMiddleware.PrintPrettyStack(rec)
					logger.ErrorContext(r.Context(), "Recover Panic",
						slog.String("error", fmt.Sprintf("%v", rec)),
						slog.String("panic_line", parseStack(debug.Stack())),
					)

					w.WriteHeader(http.StatusInternalServerError)
					jsonkit.JSONResponse(w, http.StatusInternalServerError, map[string]string{"error": "Internal Server Error"})
				}
			}()
			next.ServeHTTP(w, r)
		})
	}
}

func parseStack(debugStack []byte) string {
	// process debug stack info
	stack := strings.Split(string(debugStack), "\n")
	lines := []string{}

	// locate panic line, as we may have nested panics
	for i := len(stack) - 1; i > 0; i-- {
		lines = append(lines, stack[i])
		if strings.HasPrefix(stack[i], "panic(") {
			lines = lines[0 : len(lines)-2] // remove boilerplate
			break
		}
	}

	// reverse
	slices.Reverse(lines)

	if len(lines) < 2 {
		return ""
	}

	sourceLine := strings.TrimSpace(lines[1])

	idx := strings.LastIndex(sourceLine, ".go:")
	if idx < 0 {
		return ""
	}

	buf := &bytes.Buffer{}
	path := sourceLine[0 : idx+3]
	lineno := sourceLine[idx+3:]

	idx = strings.LastIndex(path, string(os.PathSeparator))
	dir := path[0 : idx+1]
	file := path[idx+1:]

	idx = strings.Index(lineno, " ")
	if idx > 0 {
		lineno = lineno[0:idx]
	}
	buf.WriteString(dir)
	buf.WriteString(file)
	buf.WriteString(lineno)

	return buf.String()
}
