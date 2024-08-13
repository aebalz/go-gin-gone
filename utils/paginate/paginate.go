package paginate

// TODO move to gcl utils
import (
	"math"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type PaginateMeta struct {
	LastPage    int   `json:"last_page"`
	CurrentPage int   `json:"current_page"`
	Limit       int   `json:"limit"`
	Total       int64 `json:"total"`
}

type Param struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

func GetPaginateParam(c *gin.Context) *Param {
	page, _ := strconv.Atoi(c.Query("page"))
	if page <= 0 {
		page = 1
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	if limit <= 0 {
		limit = 15
	}

	return &Param{
		Page:  page,
		Limit: limit,
	}
}

func ORMScope(param *Param) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if param.Page == 0 {
			param.Page = 1
		}

		offset := (param.Page - 1) * param.Limit

		return db.Offset(offset).Limit(param.Limit)
	}
}

func CalculateLastPage(total int64, limit int) int {
	lastPage := float64(total) / float64(limit)
	return int(math.Ceil(lastPage))
}
