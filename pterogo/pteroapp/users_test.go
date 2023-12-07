package pteroapp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"
)

func TestUserAppGetByID(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	fmt.Println(env.Application.endpoint)

	env.Mux.HandleFunc("/api/application/users/1", func(w http.ResponseWriter, r *http.Request) {
		testUser := User{}
		testUser.Attributes.ID = 1
		json.NewEncoder(w).Encode(testUser)
	})

	ctx := context.Background()

	user, _, err := env.Application.Users.GetByID(ctx, 1)

	if err != nil {
		t.Fatal(err)
	}

	if user.Attributes.ID != 1 {
		t.Errorf("unexpected user ID: %v", user.Attributes.ID)
	}
}

func TestUserAppCreate(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	fmt.Println(env.Application.endpoint)

	env.Mux.HandleFunc("/api/application/users", func(w http.ResponseWriter, r *http.Request) {
		testUser := User{}
		testUser.Attributes.ID = 1
		json.NewEncoder(w).Encode(testUser)
	})

	userOpts := UserCreateOpts{
		Email:     "pterogo@example.com",
		Username:  "pterogo",
		FirstName: "ptero",
		LastName:  "go",
	}
	user, _, err := env.Application.Users.Create(context.Background(), userOpts)
	if err != nil {
		t.Fatal(err)
	}

	if user.Attributes.ID != 1 {
		t.Errorf("unexpected user ID: %v", user.Attributes.ID)
	}
}
