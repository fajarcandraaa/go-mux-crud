package handler

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/fajarcandraaa/go-mux-crud/entity"
	"github.com/fajarcandraaa/go-mux-crud/entity/userentity"
	"github.com/fajarcandraaa/go-mux-crud/helpers"
	"github.com/fajarcandraaa/go-mux-crud/src/user"
	"github.com/pkg/errors"
)

type UserHandler struct {
	service user.Service
}

func NewUserHandler(service user.Service) *UserHandler {
	return &UserHandler{
		service: service,
	}
}

// RegisterNewUser is func to Handle user registration
func (uh *UserHandler) RegisterNewUser(w http.ResponseWriter, r *http.Request) {
	responder := helpers.NewHTTPResponse("registerNewUser")
	ctx := r.Context()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	var payload userentity.UserRequest
	err = json.Unmarshal(body, &payload)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	newUser, err := uh.service.InsertNewUser(ctx, &payload) //uh.service.InsertNewUser(&payload)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrUserAlreadyExist:
			responder.FieldErrors(w, err, http.StatusNotAcceptable, err.Error())
			return
		default:
			responder.FieldErrors(w, err, http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
	}
	responder.SuccessJSON(w, newUser, http.StatusCreated, "Succes to register new user")
	return
}
