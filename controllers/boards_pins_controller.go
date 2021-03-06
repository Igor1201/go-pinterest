package controllers

import (
	"github.com/BrandonRomano/wrecker"
	"github.com/carrot/go-pinterest/models"
)

// BoardsPinsController is the controller that is responsible for all
// /v1/boards/<board_spec:board>/pins/ endpoints in the Pinterest API.
type BoardsPinsController struct {
	wreckerClient *wrecker.Wrecker
}

// newBoardsPinsController instantiates a new BoardsPinsController.
func newBoardsPinsController(wc *wrecker.Wrecker) *BoardsPinsController {
	return &BoardsPinsController{
		wreckerClient: wc,
	}
}

// BoardsPinsFetchOptionals is a struct that represents the optional parameters
// that can be passed to the Fetch endpoint
type BoardsPinsFetchOptionals struct {
	Cursor string
}

// Fetch loads a board from the board_spec (username/board-slug)
// Endpoint: [GET] /v1/boards/<board_spec:board>/pins/
func (bpc *BoardsPinsController) Fetch(boardSpec string, optionals *BoardsPinsFetchOptionals) (*[]models.Pin, error) {
	// Make request
	response := new(models.Response)
	response.Data = &[]models.Pin{}
	resp, err := bpc.wreckerClient.Get("/boards/"+boardSpec+"/pins/").
		URLParam("fields", "id,link,note,url,attribution,color,board,counts,created_at,creator,image,media,metadata,original_link").
		Into(response).
		Execute()

	// Error from Wrecker
	if err != nil {
		return nil, err
	}

	// Status code
	if !(resp.StatusCode >= 200 && resp.StatusCode < 300) {
		return nil, &models.PinterestError{
			StatusCode: resp.StatusCode,
			Message:    response.Message,
		}
	}

	// OK
	return response.Data.(*[]models.Pin), nil
}
