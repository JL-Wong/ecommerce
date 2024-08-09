package subscriber

import (
	"database/sql"
	"encoding/json"
	"log"
	"sync"

	"go-crud-nats/model"

	//"github.com/jmoiron/sqlx"
	"github.com/nats-io/nats.go"
)

const numWorkers = 250

func CreateUserSubscriber(db *sql.DB, nc *nats.Conn) {
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go createUserWorker(db, nc, &wg)
	}

	log.Println("CreateUserSubscriber running with worker pool")
	wg.Wait()
}

func createUserWorker(db *sql.DB, nc *nats.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := nc.QueueSubscribe("user.create", "user_create_group", func(msg *nats.Msg) {
		var user model.User
		json.Unmarshal(msg.Data, &user)

		_, err := db.Exec("INSERT INTO users (user_id, name, password, role) VALUES ($1, $2, $3, $4)", user.User_ID, user.Name, user.Password, user.Role)
		if err != nil {
			log.Println("Error creating user:", err)
			msg.Respond([]byte(`{"status":"error", "message":"Failed to create user"}`))
			return
		}
		msg.Respond([]byte(`{"status":"success", "user_id":"` + user.User_ID + `"}`))
	})

	if err != nil {
		log.Println("Error subscribing to user.create:", err)
	}
}

func GetAllUserSubscriber(db *sql.DB, nc *nats.Conn) {
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go getAllUserWorker(db, nc, &wg)
	}

	log.Println("GetAllUserSubscriber running with worker pool")
	wg.Wait()
}

func getAllUserWorker(db *sql.DB, nc *nats.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := nc.QueueSubscribe("user.get_all", "user_get_all_group", func(msg *nats.Msg) {
		rows, err := db.Query("SELECT user_id, name, password, role FROM users")
		if err != nil {
			log.Println("Error querying users:", err)
			msg.Respond([]byte(`{"status":"error", "message":"Failed to get users"}`))
			return
		}
		defer rows.Close()

		users := []model.User{}
		for rows.Next() {
			var u model.User
			if err := rows.Scan(&u.User_ID, &u.Name, &u.Password, &u.Role); err != nil {
				log.Println("Error scanning user:", err)
				msg.Respond([]byte(`{"status":"error", "message":"Failed to get users"}`))
				return
			}
			users = append(users, u)
		}
		if err := rows.Err(); err != nil {
			log.Println("Error with rows:", err)
			msg.Respond([]byte(`{"status":"error", "message":"Failed to get users"}`))
			return
		}

		response, _ := json.Marshal(users)
		msg.Respond(response)
	})

	if err != nil {
		log.Println("Error subscribing to user.get_all:", err)
	}
}

func GetUserSubscriber(db *sql.DB, nc *nats.Conn) {
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go getUserWorker(db, nc, &wg)
	}

	log.Println("GetUserSubscriber running with worker pool")
	wg.Wait()
}

func getUserWorker(db *sql.DB, nc *nats.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := nc.QueueSubscribe("user.get", "user_get_group", func(msg *nats.Msg) {
		userID := string(msg.Data)

		var u model.User
		err := db.QueryRow("SELECT user_id, name, password, role FROM users WHERE user_id = $1", userID).Scan(&u.User_ID, &u.Name, &u.Password, &u.Role)
		if err != nil {
			if err == sql.ErrNoRows {
				log.Println("User not found")
				msg.Respond([]byte(`{"status":"error", "message":"User not found"}`))
				return
			}
			log.Println("Error querying user:", err)
			msg.Respond([]byte(`{"status":"error", "message":"Internal server error"}`))
			return
		}

		response, _ := json.Marshal(u)
		msg.Respond(response)
	})

	if err != nil {
		log.Println("Error subscribing to user.get:", err)
	}
}

func UpdateUserSubscriber(db *sql.DB, nc *nats.Conn) {
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go updateUserWorker(db, nc, &wg)
	}

	log.Println("UpdateUserSubscriber running with worker pool")
	wg.Wait()
}

func updateUserWorker(db *sql.DB, nc *nats.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := nc.QueueSubscribe("user.update", "user_update_group", func(msg *nats.Msg) {
		var request map[string]interface{}
		err := json.Unmarshal(msg.Data, &request)
		if err != nil {
			log.Printf("Failed to unmarshal request JSON: %v", err)
			msg.Respond([]byte(`{"status":"error", "message":"Invalid request"}`))
			return
		}

		userMap := request["user"].(map[string]interface{})
		userID := request["id"].(string)

		var u model.User
		u.User_ID = userID
		u.Name = userMap["name"].(string)
		u.Password = userMap["password"].(string)
		u.Role = userMap["role"].(string)

		_, err = db.Exec("UPDATE users SET name = $1, password = $2, role = $3 WHERE user_id = $4", u.Name, u.Password, u.Role, u.User_ID)
		if err != nil {
			log.Printf("Error updating user: %v", err)
			msg.Respond([]byte(`{"status":"error", "message":"Internal server error"}`))
			return
		}

		var updatedUser model.User
		err = db.QueryRow("SELECT user_id, name, password, role FROM users WHERE user_id = $1", u.User_ID).Scan(&updatedUser.User_ID, &updatedUser.Name, &updatedUser.Password, &updatedUser.Role)
		if err != nil {
			log.Printf("Error querying updated user: %v", err)
			msg.Respond([]byte(`{"status":"error", "message":"Internal server error"}`))
			return
		}

		response, _ := json.Marshal(map[string]interface{}{
			"status": "success",
			"user":   updatedUser,
		})
		msg.Respond(response)
	})

	if err != nil {
		log.Println("Error subscribing to user.update:", err)
	}
}

func DeleteUserSubscriber(db *sql.DB, nc *nats.Conn) {
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go deleteUserWorker(db, nc, &wg)
	}

	log.Println("DeleteUserSubscriber running with worker pool")
	wg.Wait()
}

func deleteUserWorker(db *sql.DB, nc *nats.Conn, wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := nc.QueueSubscribe("user.delete", "user_delete_group", func(msg *nats.Msg) {
		var request map[string]string
		err := json.Unmarshal(msg.Data, &request)
		if err != nil {
			log.Printf("Failed to unmarshal request JSON: %v", err)
			msg.Respond([]byte(`{"status":"error", "message":"Invalid request"}`))
			return
		}

		userID := request["id"]

		_, err = db.Exec("DELETE FROM users WHERE user_id = $1", userID)
		if err != nil {
			log.Printf("Error deleting user: %v", err)
			msg.Respond([]byte(`{"status":"error", "message":"Internal server error"}`))
			return
		}

		response, _ := json.Marshal(map[string]interface{}{
			"status":  "success",
			"message": "User deleted",
		})
		msg.Respond(response)
	})

	if err != nil {
		log.Println("Error subscribing to user.delete:", err)
	}
}

// //

// package subscriber

// import (
//     "database/sql"
//     "encoding/json"
//     "log"
//     "sync"

//     "github.com/nats-io/nats.go"
//     "go-crud-nats/model"
// )

// const numWorkers = 10

// func CreateUserSubscriber(db *sql.DB, nc *nats.Conn) {
//     jobs := make(chan *nats.Msg, 100)
//     var wg sync.WaitGroup

//     for i := 0; i < numWorkers; i++ {
//         wg.Add(1)
//         go createWorker(db, jobs, &wg)
//     }

//     nc.Subscribe("user.create", func(msg *nats.Msg) {
//         jobs <- msg
//     })

//     log.Println("CreateUserSubscriber running with worker pool")
//     wg.Wait()
// }

// func createWorker(db *sql.DB, jobs <-chan *nats.Msg, wg *sync.WaitGroup) {
//     defer wg.Done()
//     for msg := range jobs {
//         var user model.User
//         json.Unmarshal(msg.Data, &user)

//         _, err := db.Exec("INSERT INTO users (user_id, name, password, role) VALUES ($1, $2, $3, $4)", user.User_ID, user.Name, user.Password, user.Role)
//         if err != nil {
//             log.Println("Error creating user:", err)
//             msg.Respond([]byte(`{"status":"error", "message":"Failed to create user"}`))
//             continue
//         }
//         msg.Respond([]byte(`{"status":"success", "user_id":"` + user.User_ID + `"}`))
//     }
// }

// func GetAllUserSubscriber(db *sql.DB, nc *nats.Conn) {
//     jobs := make(chan *nats.Msg, 100)
//     var wg sync.WaitGroup

//     for i := 0; i < numWorkers; i++ {
//         wg.Add(1)
//         go getAllWorker(db, jobs, &wg)
//     }

//     nc.Subscribe("user.get_all", func(msg *nats.Msg) {
//         jobs <- msg
//     })

//     log.Println("GetAllUserSubscriber running with worker pool")
//     wg.Wait()
// }

// func getAllWorker(db *sql.DB, jobs <-chan *nats.Msg, wg *sync.WaitGroup) {
//     defer wg.Done()
//     for msg := range jobs {
//         rows, err := db.Query("SELECT user_id, name, password, role FROM users")
//         if err != nil {
//             log.Println("Error querying users:", err)
//             msg.Respond([]byte(`{"status":"error", "message":"Failed to get users"}`))
//             continue
//         }
//         defer rows.Close()

//         users := []model.User{}
//         for rows.Next() {
//             var u model.User
//             if err := rows.Scan(&u.User_ID, &u.Name, &u.Password, &u.Role); err != nil {
//                 log.Println("Error scanning user:", err)
//                 msg.Respond([]byte(`{"status":"error", "message":"Failed to get users"}`))
//                 continue
//             }
//             users = append(users, u)
//         }
//         if err := rows.Err(); err != nil {
//             log.Println("Error with rows:", err)
//             msg.Respond([]byte(`{"status":"error", "message":"Failed to get users"}`))
//             continue
//         }

//         response, _ := json.Marshal(users)
//         msg.Respond(response)
//     }
// }

// func GetUserSubscriber(db *sql.DB, nc *nats.Conn) {
//     jobs := make(chan *nats.Msg, 100)
//     var wg sync.WaitGroup

//     for i := 0; i < numWorkers; i++ {
//         wg.Add(1)
//         go getUserWorker(db, jobs, &wg)
//     }

//     nc.Subscribe("user.get", func(msg *nats.Msg) {
//         jobs <- msg
//     })

//     log.Println("GetUserSubscriber running with worker pool")
//     wg.Wait()
// }

// func getUserWorker(db *sql.DB, jobs <-chan *nats.Msg, wg *sync.WaitGroup) {
//     defer wg.Done()
//     for msg := range jobs {
//         userID := string(msg.Data)

//         var u model.User
//         err := db.QueryRow("SELECT user_id, name, password, role FROM users WHERE user_id = $1", userID).Scan(&u.User_ID, &u.Name, &u.Password, &u.Role)
//         if err != nil {
//             if err == sql.ErrNoRows {
//                 log.Println("User not found")
//                 msg.Respond([]byte(`{"status":"error", "message":"User not found"}`))
//                 continue
//             }
//             log.Println("Error querying user:", err)
//             msg.Respond([]byte(`{"status":"error", "message":"Internal server error"}`))
//             continue
//         }

//         response, _ := json.Marshal(u)
//         msg.Respond(response)
//     }
// }

// func UpdateUserSubscriber(db *sql.DB, nc *nats.Conn) {
//     jobs := make(chan *nats.Msg, 100)
//     var wg sync.WaitGroup

//     for i := 0; i < numWorkers; i++ {
//         wg.Add(1)
//         go updateUserWorker(db, jobs, &wg)
//     }

//     nc.Subscribe("user.update", func(msg *nats.Msg) {
//         jobs <- msg
//     })

//     log.Println("UpdateUserSubscriber running with worker pool")
//     wg.Wait()
// }

// func updateUserWorker(db *sql.DB, jobs <-chan *nats.Msg, wg *sync.WaitGroup) {
//     defer wg.Done()
//     for msg := range jobs {
//         var request map[string]interface{}
//         err := json.Unmarshal(msg.Data, &request)
//         if err != nil {
//             log.Printf("Failed to unmarshal request JSON: %v", err)
//             msg.Respond([]byte(`{"status":"error", "message":"Invalid request"}`))
//             continue
//         }

//         userMap := request["user"].(map[string]interface{})
//         userID := request["id"].(string)

//         var u model.User
//         u.User_ID = userID
//         u.Name = userMap["name"].(string)
//         u.Password = userMap["password"].(string)
//         u.Role = userMap["role"].(string)

//         _, err = db.Exec("UPDATE users SET name = $1, password = $2, role = $3 WHERE user_id = $4", u.Name, u.Password, u.Role, u.User_ID)
//         if err != nil {
//             log.Printf("Error updating user: %v", err)
//             msg.Respond([]byte(`{"status":"error", "message":"Internal server error"}`))
//             continue
//         }

//         var updatedUser model.User
//         err = db.QueryRow("SELECT user_id, name, password, role FROM users WHERE user_id = $1", u.User_ID).Scan(&updatedUser.User_ID, &updatedUser.Name, &updatedUser.Password, &updatedUser.Role)
//         if err != nil {
//             log.Printf("Error querying updated user: %v", err)
//             msg.Respond([]byte(`{"status":"error", "message":"Internal server error"}`))
//             continue
//         }

//         response, _ := json.Marshal(map[string]interface{}{
//             "status": "success",
//             "user":   updatedUser,
//         })
//         msg.Respond(response)
//     }
// }

// func DeleteUserSubscriber(db *sql.DB, nc *nats.Conn) {
//     jobs := make(chan *nats.Msg, 100)
//     var wg sync.WaitGroup

//     for i := 0; i < numWorkers; i++ {
//         wg.Add(1)
//         go deleteUserWorker(db, jobs, &wg)
//     }

//     nc.Subscribe("user.delete", func(msg *nats.Msg) {
//         jobs <- msg
//     })

//     log.Println("DeleteUserSubscriber running with worker pool")
//     wg.Wait()
// }

// func deleteUserWorker(db *sql.DB, jobs <-chan *nats.Msg, wg *sync.WaitGroup) {
//     defer wg.Done()
//     for msg := range jobs {
//         var request map[string]string
//         err := json.Unmarshal(msg.Data, &request)
//         if err != nil {
//             log.Printf("Failed to unmarshal request JSON: %v", err)
//             msg.Respond([]byte(`{"status":"error", "message":"Invalid request"}`))
//             continue
//         }

//         userID := request["id"]

//         _, err = db.Exec("DELETE FROM users WHERE user_id = $1", userID)
//         if err != nil {
//             log.Printf("Error deleting user: %v", err)
//             msg.Respond([]byte(`{"status":"error", "message":"Internal server error"}`))
//             continue
//         }

//         response, _ := json.Marshal(map[string]interface{}{
//             "status":  "success",
//             "message": "User deleted",
//         })
//         msg.Respond(response)
//     }
// }
