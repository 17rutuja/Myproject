1. First mkdir and go to the particular folder

2. Run the command and it will automatically generate "go.mode" file (This is similar to package.json file in nodejs) 
    => go mod init golang-restaurant-management
   
3. Some package installation
    => go get "github.com/gin-gonic/gin"
    => go get "go.mongodb.org/mongo-driver/mongo"
    => go mod "github.com/dgrijalva/jwt-go"
    
4. GO TIDY is used for all uninstall packages to install 
    => go mod tidy

Commands: go build && ./APP_NAME