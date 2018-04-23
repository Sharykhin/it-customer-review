package handler

import (
	"net/http"

	"github.com/Sharykhin/it-customer-review/api/entity"
	"github.com/Sharykhin/it-customer-review/api/grpc"
	"github.com/Sharykhin/it-customer-review/api/util"
)

// Index returns a list of reviews with some criteria
func Index(w http.ResponseWriter, r *http.Request) {

	limit, err := queryParamInt(r, "limit", 10)
	if err != nil {
		util.JSONBadRequest(err, w)
		return
	}

	offset, err := queryParamInt(r, "offset", 0)
	if err != nil {
		util.JSONBadRequest(err, w)
		return
	}

	var c []entity.Criteria
	if cr := queryCriteria(r, "category"); cr != nil {
		c = append(c, *cr)
	}

	if cr := queryCriteria(r, "published"); cr != nil {
		c = append(c, *cr)
	}

	rs, t, err := grpc.ReviewService.Index(r.Context(), c, limit, offset)

	if err != nil {
		util.JSON(util.Response{
			Success: false,
			Data:    nil,
			Error:   util.ErrorField{Err: err},
			Meta:    nil,
		}, w, http.StatusInternalServerError)
	}

	util.JSON(util.Response{
		Success: true,
		Data:    rs,
		Error:   util.ErrorField{Err: nil},
		Meta: map[string]int64{
			"count":  int64(len(rs)),
			"total":  t,
			"limit":  limit,
			"offset": offset,
		},
	}, w, http.StatusOK)
}
