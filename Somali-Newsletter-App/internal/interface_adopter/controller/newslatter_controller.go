package controller

import (
	"Somali-Newsletter-App/internal/entity"
	"Somali-Newsletter-App/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
)


type NewsletterController struct {
	newslatterService service.NewsletterService
}

func NewNewsletterController(newslatterService service.NewsletterService) *NewsletterController {
	return &NewsletterController{
		newslatterService: newslatterService,
	}
}

func (Nc *NewsletterController) CreateNewslatter(c *gin.Context) {
	var newsLatter entity.Newsletter
	if err := c.BindJSON(&newsLatter); err!= nil  {
		c.JSON(400, gin.H{"error": err.Error()} )
		return
	 }
// call service
   createdNewsLatter, err := Nc.newslatterService.CreateNewslatter(newsLatter.Title, newsLatter.Content, newsLatter.CreatorID, newsLatter.Published)
 if err != nil {
	c.JSON(500, gin.H{"error": err.Error()})
	return
 }

 c.JSON(201, createdNewsLatter)
}

func (Nc *NewsletterController) UpdateNewslatter(c *gin.Context) {
	var newslatter entity.Newsletter
		// param
		newslatterParam := c.Param("id")
		newslatterID, err := uuid.FromString(newslatterParam)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
	// bind 
	if err := c.BindJSON(&newslatter); err != nil {
		c.JSON(400, gin.H{"error":err.Error()})
		return
	}

	newslatter.ID = newslatterID
// call service
  if err := Nc.newslatterService.UpdateNewslatter(&newslatter); err != nil {
	c.JSON(500, gin.H{"error": err.Error()})
	return
 }

 // return
 c.JSON(200, gin.H{"message": "Newslatter updated sucessfully "})

}

func  (Nc *NewsletterController) GetAllNewsletter(c *gin.Context) {
  newslatter, err := Nc.newslatterService.GetAllNewsletter()
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
// return
c.JSON(200, newslatter)
}

func (Nc *NewsletterController)	GetNewsletterByID(c *gin.Context) {
	newslatterParam := c.Param("id")
		newslatterID, err := uuid.FromString(newslatterParam)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		// call service
	newslatter, err := Nc.newslatterService.GetNewsletterByID(newslatterID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return
	c.JSON(200, newslatter)
}

func (Nc *NewsletterController) DeleteNewsletter(c *gin.Context) {
	newslatterParam := c.Param("id")
	newslatterID, err := uuid.FromString(newslatterParam)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	// call
	if err := Nc.newslatterService.DeleteNewsletter(newslatterID); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// return
	c.JSON(200, gin.H{"message": " Newslatter Sucessfully deleted!" })

}

