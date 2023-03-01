package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/fajarcandraaa/go-mux-crud/entity"
	"github.com/fajarcandraaa/go-mux-crud/entity/userentity"
	"github.com/fajarcandraaa/go-mux-crud/helpers"
	"github.com/fajarcandraaa/go-mux-crud/src/user"
	"github.com/gorilla/mux"
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

func (ud *UserHandler) FindUserByUserID(w http.ResponseWriter, r *http.Request) {
	var (
		userID    = mux.Vars(r)["id"]
		responder = helpers.NewHTTPResponse("registerNewUser")
		ctx       = r.Context()
	)

	findUser, err := ud.service.FindUser(ctx, userID)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrUserNotExist:
			responder.ErrorJSON(w, http.StatusNotFound, "user not found")
			return
		default:
			responder.FailureJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	responder.SuccessJSON(w, findUser, http.StatusOK, "User found")
	return
}

func (uh *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	var (
		param        = r.URL.Query()
		paramPage    = param.Get("page")
		paramPerPage = param.Get("per_page")
		paramOrderBy = param.Get("order_by")
		paramSortBy  = param.Get("sort_by")
		responder    = helpers.NewHTTPResponse("registerNewUser")
		ctx          = r.Context()
	)

	paginationParam, err := helpers.SetDefaultPginationParam(paramPage, paramPerPage, paramOrderBy, paramSortBy)
	if err != nil {
		responder.FieldErrors(w, err, http.StatusUnprocessableEntity, "value of query parameters has diferent type")
		return
	}
	sortBy := paginationParam.SortBy
	orderBy := paginationParam.OrderBy
	perPage := paginationParam.PerPage
	page, _ := strconv.Atoi(paginationParam.Page)

	users, total, err := uh.service.GetListUsers(ctx, sortBy, orderBy, int(perPage), page)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrUserNotExist:
			responder.ErrorJSON(w, http.StatusNotFound, "users list not found")
			return
		default:
			responder.FailureJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	pagination, err := helpers.GetPagination(helpers.PaginationParams{
		Path:        "list.users",
		Page:        strconv.Itoa(page),
		TotalRows:   int32(total),
		PerPage:     int32(perPage),
		OrderBy:     orderBy,
		SortBy:      sortBy,
		CurrentPage: int32(page),
	})
	if err != nil {
		responder.ErrorJSON(w, http.StatusConflict, "error pagination")
		return
	}

	responder.SuccessWithMeta(w, users, pagination, http.StatusOK, "uses list")
	return

}

func (uh *UserHandler) UpdateDataUsers(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("updateDataUser")
		ctx       = r.Context()
	)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	var payload userentity.UserData
	err = json.Unmarshal(body, &payload)
	if err != nil {
		responder.ErrorWithStatusCode(w, http.StatusUnprocessableEntity, fmt.Sprint(err))
		return
	}

	updatedUser, err := uh.service.UpdateUser(ctx, &payload)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrUserNotExist:
			responder.FieldErrors(w, err, http.StatusNotExtended, err.Error())
			return
		default:
			responder.FieldErrors(w, err, http.StatusInternalServerError, fmt.Sprint(err))
			return
		}
	}
	responder.SuccessJSON(w, updatedUser, http.StatusCreated, "Succes to update data user")
	return
}

func (uh *UserHandler) UserDelete(w http.ResponseWriter, r *http.Request) {
	var (
		responder = helpers.NewHTTPResponse("updateDataUser")
		ctx       = r.Context()
		userID    = mux.Vars(r)["id"]
	)

	err := uh.service.DeleteDataUser(ctx, userID)
	if err != nil {
		causer := errors.Cause(err)
		switch causer {
		case entity.ErrUserNotExist:
			responder.ErrorJSON(w, http.StatusNotFound, "user not found")
			return
		default:
			responder.FailureJSON(w, err, http.StatusInternalServerError)
			return
		}
	}

	responder.SuccessWithoutData(w, http.StatusOK, "successfully to delete user")
	return
}
