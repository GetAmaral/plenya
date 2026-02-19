package handlers

import (
	"errors"
	"fmt"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"strings"
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

// UploadFile godoc
// @Summary Upload de artigo ou livro
// @Description Faz upload de um arquivo (PDF, EPUB, TXT ou MD). Use ?type=book para importar como livro com capítulos.
// @Tags articles
// @Accept multipart/form-data
// @Produce json
// @Param file formData file true "Arquivo PDF, EPUB, TXT ou MD"
// @Param type query string false "Tipo de importação: 'book' para dividir em capítulos"
// @Success 201 {object} models.Article
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/upload [post]
func (h *ArticleHandler) UploadFile(c *fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Arquivo não encontrado",
		})
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	allowedExts := map[string]bool{
		".pdf":  true,
		".epub": true,
		".txt":  true,
		".md":   true,
	}
	if !allowedExts[ext] {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Formato não suportado. Formatos aceitos: PDF, EPUB, TXT, MD",
		})
	}

	// Detectar se é importação como livro (?type=book)
	isBook := c.Query("type") == "book"

	// Limites de tamanho: 50MB para artigos, 200MB para livros
	maxSize := int64(50 * 1024 * 1024)
	if isBook {
		maxSize = int64(200 * 1024 * 1024)
	}
	if file.Size > maxSize {
		if isBook {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("Arquivo muito grande (máximo 200MB para livros, atual: %s)", formatBytes(file.Size)),
			})
		}
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": fmt.Sprintf("Arquivo muito grande (máximo 50MB, atual: %s)", formatBytes(file.Size)),
		})
	}

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

	if isBook {
		bookArticle, bookErr := h.service.UploadBook(fileReader, file.Filename, userID)
		if bookErr != nil {
			var dupErr *services.DuplicateFileError
			if errors.As(bookErr, &dupErr) {
				return c.Status(fiber.StatusConflict).JSON(fiber.Map{
					"error":         "DUPLICATE_FILE",
					"message":       "Este arquivo já foi importado anteriormente",
					"existingTitle": dupErr.ExistingTitle,
					"existingId":    dupErr.ExistingID.String(),
				})
			}
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Erro ao processar livro: " + bookErr.Error(),
			})
		}
		return c.Status(fiber.StatusCreated).JSON(bookArticle)
	}

	regularArticle, err := h.service.UploadFile(fileReader, file.Filename, userID)
	if err != nil {
		var dupErr *services.DuplicateFileError
		if errors.As(err, &dupErr) {
			return c.Status(fiber.StatusConflict).JSON(fiber.Map{
				"error":         "DUPLICATE_FILE",
				"message":       "Este arquivo já foi importado anteriormente",
				"existingTitle": dupErr.ExistingTitle,
				"existingId":    dupErr.ExistingID.String(),
			})
		}
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao processar arquivo: " + err.Error(),
		})
	}

	return c.Status(fiber.StatusCreated).JSON(regularArticle)
}

// formatBytes formata bytes para exibição legível
func formatBytes(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}

// GetChapters godoc
// @Summary Listar capítulos de um livro
// @Description Retorna a lista ordenada de capítulos de um livro (source_type=book)
// @Tags articles
// @Produce json
// @Param id path string true "ID do livro"
// @Success 200 {object} map[string]interface{}
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/{id}/chapters [get]
func (h *ArticleHandler) GetChapters(c *fiber.Ctx) error {
	idParam := c.Params("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	chapters, err := h.service.GetChaptersByBookID(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar capítulos: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"chapters": chapters,
		"total":    len(chapters),
	})
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

// AddScoreItemsToArticle godoc
// @Summary Adicionar itens de escore a um artigo
// @Description Cria associação many-to-many entre artigo e itens de escore
// @Tags articles
// @Accept json
// @Produce json
// @Param id path string true "ID do artigo"
// @Param scoreItemIds body []string true "Array de IDs dos itens de escore"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/{id}/score-items [post]
func (h *ArticleHandler) AddScoreItemsToArticle(c *fiber.Ctx) error {
	idParam := c.Params("id")
	articleID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	var request struct {
		ScoreItemIDs []string `json:"scoreItemIds" validate:"required,min=1"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados inválidos",
		})
	}

	// Converter strings para UUIDs
	scoreItemIDs := make([]uuid.UUID, len(request.ScoreItemIDs))
	for i, idStr := range request.ScoreItemIDs {
		id, err := uuid.Parse(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("ID inválido: %s", idStr),
			})
		}
		scoreItemIDs[i] = id
	}

	if err := h.service.AddScoreItemsToArticle(articleID, scoreItemIDs); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao adicionar itens de escore: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Itens de escore adicionados com sucesso",
	})
}

// RemoveScoreItemsFromArticle godoc
// @Summary Remover itens de escore de um artigo
// @Description Remove associação many-to-many entre artigo e itens de escore
// @Tags articles
// @Accept json
// @Produce json
// @Param id path string true "ID do artigo"
// @Param scoreItemIds body []string true "Array de IDs dos itens de escore"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/{id}/score-items [delete]
func (h *ArticleHandler) RemoveScoreItemsFromArticle(c *fiber.Ctx) error {
	idParam := c.Params("id")
	articleID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	var request struct {
		ScoreItemIDs []string `json:"scoreItemIds" validate:"required,min=1"`
	}

	if err := c.BodyParser(&request); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Dados inválidos",
		})
	}

	// Converter strings para UUIDs
	scoreItemIDs := make([]uuid.UUID, len(request.ScoreItemIDs))
	for i, idStr := range request.ScoreItemIDs {
		id, err := uuid.Parse(idStr)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": fmt.Sprintf("ID inválido: %s", idStr),
			})
		}
		scoreItemIDs[i] = id
	}

	if err := h.service.RemoveScoreItemsFromArticle(articleID, scoreItemIDs); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao remover itens de escore: " + err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "Itens de escore removidos com sucesso",
	})
}

// GetScoreItemsForArticle godoc
// @Summary Listar itens de escore de um artigo
// @Description Retorna todos os itens de escore associados a um artigo
// @Tags articles
// @Produce json
// @Param id path string true "ID do artigo"
// @Success 200 {array} models.ScoreItem
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /articles/{id}/score-items [get]
func (h *ArticleHandler) GetScoreItemsForArticle(c *fiber.Ctx) error {
	idParam := c.Params("id")
	articleID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	scoreItems, err := h.service.GetScoreItemsForArticle(articleID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar itens de escore: " + err.Error(),
		})
	}

	return c.JSON(scoreItems)
}

// GetArticlesForScoreItem godoc
// @Summary Listar artigos de um item de escore
// @Description Retorna todos os artigos associados a um item de escore
// @Tags score-items
// @Produce json
// @Param id path string true "ID do item de escore"
// @Success 200 {array} models.Article
// @Failure 400 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security BearerAuth
// @Router /score-items/{id}/articles [get]
func (h *ArticleHandler) GetArticlesForScoreItem(c *fiber.Ctx) error {
	idParam := c.Params("id")
	scoreItemID, err := uuid.Parse(idParam)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "ID inválido",
		})
	}

	articles, err := h.service.GetArticlesForScoreItem(scoreItemID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Erro ao buscar artigos: " + err.Error(),
		})
	}

	return c.JSON(articles)
}
