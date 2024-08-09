// package publisher

// import (
// 	"encoding/json"
// 	"log"
// 	"net/http"
// 	"os"
// 	"time"

// 	"go-crud-nats/model"

// 	"github.com/nats-io/nats.go"
// )

// func PublishUserCreation(user model.User, w http.ResponseWriter) (map[string]interface{}, error) {
// 	nc, err := nats.Connect(os.Getenv("NATS_URL"))
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Fatalf("Failed to connect to NATS: %v", err)
// 		return nil, err
// 	}
// 	defer nc.Close()

// 	data, err := json.Marshal(user)
// 	log.Printf("Data: %s", string(data))
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Printf("Failed to marshal user JSON: %v", err)
// 		return nil, err
// 	}

// 	reply, err := nc.Request("user.create", data, 2*time.Second)
// 	if err != nil {
// 		http.Error(w, "Failed to create user", http.StatusInternalServerError)
// 		log.Printf("Failed to publish message to NATS: %v", err)
// 		return nil, err
// 	}

// 	var response map[string]interface{}
// 	err = json.Unmarshal(reply.Data, &response)
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Printf("Failed to unmarshal response JSON: %v", err)
// 		return nil, err
// 	}

// 	log.Printf("Response: %s", string(reply.Data))
// 	json.NewEncoder(w).Encode(response)
// 	return response, nil
// }

// func PublishGetUsers(w http.ResponseWriter) ([]model.User, error) {
// 	nc, err := nats.Connect(os.Getenv("NATS_URL"))
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Fatalf("Failed to connect to NATS: %v", err)
// 		return nil, err
// 	}
// 	defer nc.Close()

// 	msg, err := nc.Request("user.get_all", nil, nats.DefaultTimeout)
// 	if err != nil {
// 		http.Error(w, "Failed to get all users", http.StatusInternalServerError)
// 		log.Printf("Failed to publish message to NATS: %v", err)
// 		return nil, err
// 	}

// 	var users []model.User
// 	if err := json.Unmarshal(msg.Data, &users); err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

// func PublishGetUser(userID string, w http.ResponseWriter) (*model.User, error) {
// 	nc, err := nats.Connect(os.Getenv("NATS_URL"))
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Fatalf("Failed to connect to NATS: %v", err)
// 		return nil, err
// 	}
// 	defer nc.Close()

// 	msg, err := nc.Request("user.get", []byte(userID), 2*time.Second)
// 	if err != nil {
// 		http.Error(w, "Failed to get user", http.StatusInternalServerError)
// 		log.Printf("Failed to publish message to NATS: %v", err)
// 		return nil, err
// 	}

// 	var user model.User
// 	if err := json.Unmarshal(msg.Data, &user); err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Printf("Failed to unmarshal user JSON: %v", err)
// 		return nil, err
// 	}

// 	return &user, nil
// }

// func PublishUpdateUser(user model.User, userID string, w http.ResponseWriter) (map[string]interface{}, error) {
// 	nc, err := nats.Connect(os.Getenv("NATS_URL"))
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Fatalf("Failed to connect to NATS: %v", err)
// 		return nil, err
// 	}
// 	defer nc.Close()

// 	request := map[string]interface{}{
// 		"user": user,
// 		"id":   userID,
// 	}

// 	data, err := json.Marshal(request)
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Printf("Failed to marshal request JSON: %v", err)
// 		return nil, err
// 	}

// 	reply, err := nc.Request("user.update", data, 5*time.Second)
// 	if err != nil {
// 		http.Error(w, "Failed to update user", http.StatusInternalServerError)
// 		log.Printf("Failed to publish message to NATS: %v", err)
// 		return nil, err
// 	}

// 	var response map[string]interface{}
// 	err = json.Unmarshal(reply.Data, &response)
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Printf("Failed to unmarshal response JSON: %v", err)
// 		return nil, err
// 	}

// 	json.NewEncoder(w).Encode(response)
// 	return response, nil
// }

// func PublishDeleteUser(userID string, w http.ResponseWriter) (map[string]interface{}, error) {
// 	nc, err := nats.Connect(os.Getenv("NATS_URL"))
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Fatalf("Failed to connect to NATS: %v", err)
// 		return nil, err
// 	}
// 	defer nc.Close()

// 	data, err := json.Marshal(map[string]string{"id": userID})
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Printf("Failed to marshal request JSON: %v", err)
// 		return nil, err
// 	}

// 	reply, err := nc.Request("user.delete", data, 2*time.Second)
// 	if err != nil {
// 		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
// 		log.Printf("Failed to publish message to NATS: %v", err)
// 		return nil, err
// 	}

// 	var response map[string]interface{}
// 	err = json.Unmarshal(reply.Data, &response)
// 	if err != nil {
// 		http.Error(w, "Internal server error", http.StatusInternalServerError)
// 		log.Printf("Failed to unmarshal response JSON: %v", err)
// 		return nil, err
// 	}

// 	json.NewEncoder(w).Encode(response)
// 	return response, nil
// }

package publisher

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"time"

	"go-crud-nats/model"

	"github.com/nats-io/nats.go"
)

var nc *nats.Conn

func init() {
	var err error
	nc, err = nats.Connect(os.Getenv("NATS_URL"))
	if err != nil {
		log.Fatalf("Failed to connect to NATS: %v", err)
	}
}

func PublishUserCreation(user model.User, w http.ResponseWriter) (map[string]interface{}, error) {
	data, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Failed to marshal user JSON: %v", err)
		return nil, err
	}

	reply, err := nc.Request("user.create", data, 2*time.Second)
	if err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		log.Printf("Failed to publish message to NATS: %v", err)
		return nil, err
	}

	var response map[string]interface{}
	err = json.Unmarshal(reply.Data, &response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Failed to unmarshal response JSON: %v", err)
		return nil, err
	}

	json.NewEncoder(w).Encode(response)
	return response, nil
}

func PublishGetUsers(w http.ResponseWriter) ([]model.User, error) {
	msg, err := nc.Request("user.get_all", nil, nats.DefaultTimeout)
	if err != nil {
		http.Error(w, "Failed to get all users", http.StatusInternalServerError)
		log.Printf("Failed to publish message to NATS: %v", err)
		return nil, err
	}

	var users []model.User
	if err := json.Unmarshal(msg.Data, &users); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Failed to unmarshal response JSON: %v", err)
		return nil, err
	}

	return users, nil
}

func PublishGetUser(userID string, w http.ResponseWriter) (*model.User, error) {
	msg, err := nc.Request("user.get", []byte(userID), 2*time.Second)
	if err != nil {
		http.Error(w, "Failed to get user", http.StatusInternalServerError)
		log.Printf("Failed to publish message to NATS: %v", err)
		return nil, err
	}

	var user model.User
	if err := json.Unmarshal(msg.Data, &user); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Failed to unmarshal user JSON: %v", err)
		return nil, err
	}

	return &user, nil
}

func PublishUpdateUser(user model.User, userID string, w http.ResponseWriter) (map[string]interface{}, error) {
	request := map[string]interface{}{
		"user": user,
		"id":   userID,
	}

	data, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Failed to marshal request JSON: %v", err)
		return nil, err
	}

	reply, err := nc.Request("user.update", data, 5*time.Second)
	if err != nil {
		http.Error(w, "Failed to update user", http.StatusInternalServerError)
		log.Printf("Failed to publish message to NATS: %v", err)
		return nil, err
	}

	var response map[string]interface{}
	err = json.Unmarshal(reply.Data, &response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Failed to unmarshal response JSON: %v", err)
		return nil, err
	}

	json.NewEncoder(w).Encode(response)
	return response, nil
}

func PublishDeleteUser(userID string, w http.ResponseWriter) (map[string]interface{}, error) {
	data, err := json.Marshal(map[string]string{"id": userID})
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Failed to marshal request JSON: %v", err)
		return nil, err
	}

	reply, err := nc.Request("user.delete", data, 2*time.Second)
	if err != nil {
		http.Error(w, "Failed to delete user", http.StatusInternalServerError)
		log.Printf("Failed to publish message to NATS: %v", err)
		return nil, err
	}

	var response map[string]interface{}
	err = json.Unmarshal(reply.Data, &response)
	if err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		log.Printf("Failed to unmarshal response JSON: %v", err)
		return nil, err
	}

	json.NewEncoder(w).Encode(response)
	return response, nil
}