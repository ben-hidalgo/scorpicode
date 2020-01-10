(cd backend/cmd/roxie/  && go run main.go) &
(cd backend/cmd/site/   && go run main.go) &
(cd backend/cmd/hats/   && go run main.go) &
wait
echo
