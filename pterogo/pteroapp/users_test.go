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

func TestUserAppGetUserList(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	fmt.Println(env.Application.endpoint)

	env.Mux.HandleFunc("/api/application/users", func(w http.ResponseWriter, r *http.Request) {
		testUserOne := User{}
		testUserOne.Attributes.ID = 1

		testUserTwo := User{}
		testUserTwo.Attributes.ID = 2

		testUserList := UsersList{}
		testUserList.Users = append(testUserList.Users, testUserOne)
		testUserList.Users = append(testUserList.Users, testUserTwo)

		json.NewEncoder(w).Encode(testUserList)
	})

	ctx := context.Background()

	userList, _, err := env.Application.Users.GetList(ctx)

	if err != nil {
		t.Fatal(err)
	}

	if len(userList.Users) != 2 {
		t.Errorf("unexpected number of users in: %v", userList)
	}
}

func TestUserAppCreate(t *testing.T) {
	env := newTestEnv()
	defer env.Teardown()

	env.Mux.HandleFunc("/api/application/users", func(w http.ResponseWriter, r *http.Request) {
		var reqBody UserCreateOpts
		if err := json.NewDecoder(r.Body).Decode(&reqBody); err != nil {
			t.Fatal(err)
		}

		if reqBody.Email != "pterogo@elasthi.cc" {
			t.Errorf("unexpected email: %s", reqBody.Email)
		}
		testUser := User{}
		testUser.Attributes.ID = 1
		json.NewEncoder(w).Encode(testUser)
	})

	userOpts := UserCreateOpts{
		Email:     "pterogo@elasthi.cc",
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
