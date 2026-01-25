package handlers

import (
	"mime/multipart"
	"strconv"
	"time"

	"github.com/plenya/api/internal/services"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// ArticleHandler gerencia as requisições HTTP de artigos
type ArticleHandler struct {
	service *services.ArticleService
}

// NewArticleHandler cria uma nova instância do handler
func NewArticleHandler(service *services.ArticleService) *ArticleHandler {
	return &ArticleHandler{service: service}
}

// CreateArticle godoc
// @Summary Criar artigo
// @Description Cria um novo artigo científico manualmente
// @Tags articles
// @Accept json
// @Produce json
// @Param article body services.CreateArticleDTO true "Dados do artigo"
// @Success 201 {object} models.Article
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles [post]
func (h *ArticleHandler) CreateArticle(c *fiber.Ctx) error {
	var dto services.CreateArticleDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados inválidos",
		})
	}

	// Obter user ID do contexto (injetado pelo middleware de auth)
	userID := c.Locals("userID").(uuid.UUID)

	article, err := h.service.CreateArticle(&dto, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao criar artigo: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(article)
}

// GetArticle godoc
// @Summary Buscar artigo por ID
// @Description Retorna os detalhes de um artigo específico
// @Tags articles
// @Produce json
// @Param id path string true "ID do artigo"
// @Success 200 {object} models.Article
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /articles/{id} [get]
func (h *ArticleHandler) GetArticle(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	article, err := h.service.GetArticleByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Artigo não encontrado",
		})
	}

	// Atualizar timestamp de último acesso
	go h.service.UpdateLastAccessed(id)

	return c.JSON(article)
}

// ListArticles godoc
// @Summary Listar artigos
// @Description Lista artigos com paginação e filtros
// @Tags articles
// @Produce json
// @Param page query int false "Página" default(1)
// @Param pageSize query int false "Itens por página" default(20)
// @Param journal query string false "Filtrar por revista"
// @Param specialty query string false "Filtrar por especialidade"
// @Param articleType query string false "Filtrar por tipo"
// @Param favorite query bool false "Filtrar favoritos"
// @Param rating query int false "Filtrar por rating"
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles [get]
func (h *ArticleHandler) ListArticles(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "20"))

	// Construir filtros
	filters := make(map[string]interface{})

	if journal := c.Query("journal"); journal != "" {
		filters["journal"] = journal
	}
	if specialty := c.Query("specialty"); specialty != "" {
		filters["specialty"] = specialty
	}
	if articleType := c.Query("articleType"); articleType != "" {
		filters["articleType"] = articleType
	}
	if favorite := c.Query("favorite"); favorite != "" {
		filters["favorite"] = favorite == "true"
	}
	if rating := c.Query("rating"); rating != "" {
		r, _ := strconv.Atoi(rating)
		filters["rating"] = r
	}

	// Filtros de data
	if publishedAfter := c.Query("publishedAfter"); publishedAfter != "" {
		if t, err := time.Parse("2006-01-02", publishedAfter); err == nil {
			filters["publishedAfter"] = t
		}
	}
	if publishedBefore := c.Query("publishedBefore"); publishedBefore != "" {
		if t, err := time.Parse("2006-01-02", publishedBefore); err == nil {
			filters["publishedBefore"] = t
		}
	}

	articles, total, err := h.service.ListArticles(page, pageSize, filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao listar artigos: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"articles": articles,
		"pagination": fiber.Map{
			"page":      page,
			"pageSize":  pageSize,
			"total":     total,
			"totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

// UpdateArticle godoc
// @Summary Atualizar artigo
// @Description Atualiza os dados de um artigo
// @Tags articles
// @Accept json
// @Produce json
// @Param id path string true "ID do artigo"
// @Param article body services.UpdateArticleDTO true "Dados a atualizar"
// @Success 200 {object} models.Article
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/{id} [put]
func (h *ArticleHandler) UpdateArticle(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	var dto services.UpdateArticleDTO
	if err := c.BodyParser(&dto); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados inválidos",
		})
	}

	userID := c.Locals("userID").(uuid.UUID)

	article, err := h.service.UpdateArticle(id, &dto, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao atualizar artigo: " + err.Error(),
		})
	}

	return c.JSON(article)
}

// DeleteArticle godoc
// @Summary Deletar artigo
// @Description Remove um artigo do sistema (soft delete)
// @Tags articles
// @Param id path string true "ID do artigo"
// @Success 204
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/{id} [delete]
func (h *ArticleHandler) DeleteArticle(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	if err := h.service.DeleteArticle(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao deletar artigo: " + err.Error(),
		})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

// UploadPDF godoc
// @Summary Upload de PDF
// @Description Faz upload de um PDF e extrai metadados automaticamente
// @Tags articles
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Arquivo PDF"
// @Success 201 {object} models.Article
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/upload [post]
func (h *ArticleHandler) UploadPDF(c *fiber.Ctx) error {
	// Parse multipart form
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Arquivo não encontrado",
		})
	}

	// Validar tipo de arquivo
	if file.Header.Get("Content-Type") != "application/pdf" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Apenas arquivos PDF são permitidos",
		})
	}

	// Validar tamanho (max 50MB)
	maxSize := int64(50 * 1024 * 1024)
	if file.Size > maxSize {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Arquivo muito grande (máximo 50MB)",
		})
	}

	// Abrir arquivo
	fileReader, err := file.Open()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao abrir arquivo",
		})
	}
	defer func(fileReader multipart.File) {
		_ = fileReader.Close()
	}(fileReader)

	userID := c.Locals("userID").(uuid.UUID)

	// Upload e extrair metadados
	article, err := h.service.UploadPDF(fileReader, file.Filename, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao processar PDF: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(article)
}

// SearchArticles godoc
// @Summary Buscar artigos
// @Description Busca artigos por texto (título, autores, abstract, revista)
// @Tags articles
// @Produce json
// @Param q query string true "Texto de busca"
// @Param page query int false "Página" default(1)
// @Param pageSize query int false "Itens por página" default(20)
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/search [get]
func (h *ArticleHandler) SearchArticles(c *fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Parâmetro 'q' é obrigatório",
		})
	}

	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "20"))

	articles, total, err := h.service.SearchArticles(query, page, pageSize)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar artigos: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"articles": articles,
		"pagination": fiber.Map{
			"page":       page,
			"pageSize":   pageSize,
			"total":      total,
			"totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

// GetFavorites godoc
// @Summary Listar favoritos
// @Description Lista artigos marcados como favoritos
// @Tags articles
// @Produce json
// @Param page query int false "Página" default(1)
// @Param pageSize query int false "Itens por página" default(20)
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/favorites [get]
func (h *ArticleHandler) GetFavorites(c *fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("pageSize", "20"))

	filters := map[string]interface{}{
		"favorite": true,
	}

	articles, total, err := h.service.ListArticles(page, pageSize, filters)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao listar favoritos: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"articles": articles,
		"pagination": fiber.Map{
			"page":       page,
			"pageSize":   pageSize,
			"total":      total,
			"totalPages": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}

// ToggleFavorite godoc
// @Summary Marcar/desmarcar favorito
// @Description Alterna o status de favorito de um artigo
// @Tags articles
// @Produce json
// @Param id path string true "ID do artigo"
// @Success 200 {object} models.Article
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/{id}/favorite [patch]
func (h *ArticleHandler) ToggleFavorite(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	article, err := h.service.GetArticleByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Artigo não encontrado",
		})
	}

	// Toggle favorite
	newFavorite := !article.Favorite
	userID := c.Locals("userID").(uuid.UUID)

	updated, err := h.service.UpdateArticle(id, &services.UpdateArticleDTO{
		Favorite: &newFavorite,
	}, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao atualizar favorito: " + err.Error(),
		})
	}

	return c.JSON(updated)
}

// SetRating godoc
// @Summary Definir rating
// @Description Define a avaliação (0-5 estrelas) de um artigo
// @Tags articles
// @Accept json
// @Produce json
// @Param id path string true "ID do artigo"
// @Param body body map[string]int true "Rating (0-5)"
// @Success 200 {object} models.Article
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/{id}/rating [patch]
func (h *ArticleHandler) SetRating(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	var body struct {
		Rating int `json:"rating"`
	}
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados inválidos",
		})
	}

	if body.Rating < 0 || body.Rating > 5 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Rating deve estar entre 0 e 5",
		})
	}

	userID := c.Locals("userID").(uuid.UUID)

	updated, err := h.service.UpdateArticle(id, &services.UpdateArticleDTO{
		Rating: &body.Rating,
	}, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao atualizar rating: " + err.Error(),
		})
	}

	return c.JSON(updated)
}

// GetStatistics godoc
// @Summary Estatísticas
// @Description Retorna estatísticas gerais sobre artigos
// @Tags articles
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/stats [get]
func (h *ArticleHandler) GetStatistics(c *fiber.Ctx) error {
	// Esta funcionalidade precisa ser implementada no repository
	return c.JSON(fiber.Map{
		"message": "Estatísticas não implementadas ainda",
	})
}

// DownloadPDF godoc
// @Summary Download de PDF
// @Description Faz download do arquivo PDF de um artigo
// @Tags articles
// @Produce application/pdf
// @Param id path string true "ID do artigo"
// @Success 200 {file} binary
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Security BearerAuth
// @Router /articles/{id}/download [get]
func (h *ArticleHandler) DownloadPDF(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	article, err := h.service.GetArticleByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "Artigo não encontrado",
		})
	}

	if article.InternalLink == nil || *article.InternalLink == "" {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": "PDF não disponível",
		})
	}

	// Retornar arquivo
	return c.SendFile(*article.InternalLink)
}
