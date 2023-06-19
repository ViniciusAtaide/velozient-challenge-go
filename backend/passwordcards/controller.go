package passwordcards

import (
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/viniciusataide/velozient-challenge-go/domain"
	"github.com/viniciusataide/velozient-challenge-go/server"
	"github.com/viniciusataide/velozient-challenge-go/utils"
)

type Controller struct {
	Service *Service
	server  *server.Server
}

func ProvideController(service *Service, server *server.Server) *Controller {
	group := server.Group("/password-cards")

	ctrl := Controller{service, server}

	group.Get("/", ctrl.List)
	group.Post("/", bodyValidator[domain.PasswordCardCreate], ctrl.Create)
	group.Patch("/:id", pathValidator[domain.IdPathParam], bodyValidator[domain.PasswordCardUpdate], ctrl.Update)
	group.Delete("/:id", pathValidator[domain.IdPathParam], ctrl.Delete)
	group.Get("/:id", pathValidator[domain.IdPathParam], ctrl.Get)

	return &ctrl
}

// @Success 200 {array} domain.PasswordCard
// @Tags PasswordCards
// @Param limit query string false "Pagination limit"
// @Param offset query string false "Pagination offset"
// @Param name query string false "Query by paginationcard name"
// @Router /password-cards [get]
func (ctrl *Controller) List(c *fiber.Ctx) error {
	limit := c.QueryInt("limit", 10)
	offset := c.QueryInt("offset")
	filter := c.Query("name")

	pagination := &utils.Pagination{
		Size:   limit,
		Offset: offset,
	}

	cards, err := ctrl.Service.List(pagination, filter)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err})
	}

	// cards by default when empty come as nil, this is to make it marshall []
	if cards == nil {
		cards = make([]domain.PasswordCard, 0)
	}

	return c.JSON(cards)
}

// @Success 201 {object} domain.PasswordCard
// @Failure 400 {array} domain.ErrorResponse
// @Tags PasswordCards
// @Param body body domain.PasswordCardDto true "Pagination card"
// @Router /password-cards [post]
func (ctrl *Controller) Create(c *fiber.Ctx) error {
	dto := new(domain.PasswordCardCreate)

	if err := c.BodyParser(dto); err != nil {
		return err
	}

	card, err := ctrl.Service.Create(*dto)

	if err != nil {
		var status int
		if strings.Contains(err.Error(), "Constraint") {
			status = fiber.StatusConflict
		} else {
			status = fiber.StatusInternalServerError
		}
		return c.Status(status).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(card.Id.String())
}

// @Success 200 {object} domain.PasswordCard
// @Failure 400 {array} domain.ErrorResponse
// @Tags PasswordCards
// @Param id path string true "Pagination card id"
// @Router /password-cards/{id} [get]
func (ctrl *Controller) Get(c *fiber.Ctx) error {
	id := c.Params("id")
	pc := ctrl.Service.Get(id)

	if pc == nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.Status(fiber.StatusOK).JSON(pc)
}

// @Success 200
// @Failure 400 {array} domain.ErrorResponse
// @Failure 404
// @Tags PasswordCards
// @Param id path string true "Pagination card id"
// @Router /password-cards/{id} [delete]
func (ctrl *Controller) Delete(c *fiber.Ctx) error {
	id := c.Params("id")

	err := ctrl.Service.Delete(id)

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusOK)
}

// @Success 200
// @Failure 400 {array} domain.ErrorResponse
// @Failure 404
// @Tags PasswordCards
// @Param id path string true "Pagination card id"
// @Param body body domain.PasswordCardUpdateDto true "Pagination card"
// @Router /password-cards/{id} [patch]
func (ctrl *Controller) Update(c *fiber.Ctx) error {
	id := c.Params("id")
	dto := new(domain.PasswordCardUpdate)

	if err := c.BodyParser(dto); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := ctrl.Service.Update(id, dto)

	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.SendStatus(fiber.StatusOK)
}
