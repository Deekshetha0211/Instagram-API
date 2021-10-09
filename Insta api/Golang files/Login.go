import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
  
unauthenticatedApi := &instagram.Api{
  ClientId: "my-client-id",
}
authenticatedApi := &instagram.Api{
  AccessToken: "my-access-token",
}
- if enforceSigned false
anotherAuthenticatedApi := instagram.New("", "", "my-access-token", false)

- if enforceSigned true
anotherAuthenticatedApi := instagram.New("client_id", "client_secret", "my-access-token", false)
