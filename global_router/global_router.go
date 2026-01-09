package global_router

import "net/http"

func HandleGlobalRouter(mux *http.ServeMux) http.Handler {
	globalRouter := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.Header().Set("Content-Type", "application/json")
		
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return 
		}

		mux.ServeHTTP(w, r)
		
	}

	return http.HandlerFunc(globalRouter)
}