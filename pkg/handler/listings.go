package handler

import (
	"github.com/Horronyt/marketplace"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) createListing(c *gin.Context) {
	id, err := getUserId(c)
	if err != nil {
		return
	}

	var input marketplace.Listing
	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err = h.services.Listing.Create(id, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	c.JSON(http.StatusOK, gin.H{"id": id})
}

type GetAllListingsResponse struct {
	Data []marketplace.ListingOutputFormat `json:"data"`
}

type GetAllListingsResponseAnon struct {
	Data []marketplace.ListingOutputFormatAnon `json:"data"`
}

func (h *Handler) getListings(c *gin.Context) {
	if userId, err := getUserId(c); err != nil {
		listings, err := h.services.Listing.GetAllAnonymously()
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, GetAllListingsResponseAnon{Data: listings})
	} else {
		listings, err := h.services.Listing.GetAll(userId)
		if err != nil {
			newErrorResponse(c, http.StatusBadRequest, err.Error())
			return
		}

		c.JSON(http.StatusOK, GetAllListingsResponse{Data: listings})
	}
}
