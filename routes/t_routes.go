package routes

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/Ccj-pro/admin_server/model"
	"github.com/gin-gonic/gin"
)

func InitTestRoutes(r *gin.RouterGroup) gin.IRoutes {
	role := r.Group("/role")

	{
		role.GET("/list", func(c *gin.Context) {
			var users []model.Usert

			db, err := sql.Open("mysql", "root:csj123@tcp(127.0.0.1:3306)/car_test")
			if err != nil {
				log.Fatal(err)
			}
			defer db.Close()

			rows, err := db.Query("SELECT id, username, nickname FROM userts")
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}
			defer rows.Close()

			for rows.Next() {
				var user model.Usert
				if err := rows.Scan(&user.ID, &user.Username, &user.Nickname); err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				users = append(users, user)
			}

			if err := rows.Err(); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
				return
			}

			c.JSON(http.StatusOK, users)
		})
	}
	return r
}
