package crmenforcer

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
)

func check(s session, resource string, action string) (ok bool) {
	tenant, user := s.Destructure()
	r := rbac{
		Tenant:   tenant,
		User:     user,
		Resource: resource,
		Action:   action,
	}

	status, err := exec(r, s)
	if err != nil {
		ok = false
		return
	}

	ok = status == http.StatusOK
	return
}

func parse(i interface{}) (s session) {
	data, _ := json.Marshal(i)
	json.Unmarshal(data, &s)
	return
}

func ClientsRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "clients", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ClientsWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "clients", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ContactsRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "contacts", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ContactsWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "contacts", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func HoldingsRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "holdings", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func HoldingsWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "holdings", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func AreasRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "areas", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func AreasWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "areas", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func CategoriesRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "categories", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func CategoriesWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "categories", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func FunnelsRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "funnels", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func FunnelsWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "funnels", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ProjectsRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "projects", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ProjectsWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "projects", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ProjectsTasksRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "projects_tasks", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ProjectsTasksWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "projects_tasks", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ProjectsCommentsRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "projects_comments", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ProjectsCommentsWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "projects_comments", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ProjectsDocumentsRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "projects_documents", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ProjectsDocumentsWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "projects_documents", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ProjectsHistoryRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "projects_history", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ProjectsHistoryWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "projects_history", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ProjectsPropertiesRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "projects_properties", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ProjectsPropertiesWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "projects_properties", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ReportsStatusGeneralRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "reports_status_general", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ReportsStatusGeneralWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "reports_status_general", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ReportsStatusPersonalRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "reports_status_personal", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ReportsStatusPersonalWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "reports_status_personal", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ReportsStatusUserRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "reports_status_user", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ReportsStatusUserWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "reports_status_user", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ReportsStatusFunnelRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "reports_status_funnel", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func ReportsStatusFunnelWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "reports_status_funnel", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func UsersRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "users", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func UsersWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "users", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func RolesRead() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "roles", "read") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}

func RolesWrite() gin.HandlerFunc {
	return func(c *gin.Context) {
		s := parse(c.MustGet("session"))
		if !check(s, "roles", "write") {
			c.AbortWithStatus(http.StatusForbidden)
			return
		}
		c.Next()
	}
}
