package handler

import (
	"github.com/fastenhealth/fastenhealth-onprem/backend/pkg/database"
	"github.com/fastenhealth/fastenhealth-onprem/backend/pkg/hub"
	"github.com/fastenhealth/fastenhealth-onprem/backend/pkg/models"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func CreateSource(c *gin.Context) {
	logger := c.MustGet("LOGGER").(*logrus.Entry)
	databaseRepo := c.MustGet("REPOSITORY").(database.DatabaseRepository)

	providerCred := models.Source{}
	if err := c.ShouldBindJSON(&providerCred); err != nil {
		logger.Errorln("An error occurred while parsing posted provider credential", err)
		c.JSON(http.StatusBadRequest, gin.H{"success": false})
		return
	}

	logger.Infof("Parsed Create Provider Credentials Payload: %v", providerCred)

	err := databaseRepo.CreateSource(c, &providerCred)
	if err != nil {
		logger.Errorln("An error occurred while storing provider credential", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	// after creating the source, we should do a bulk import
	sourceClient, err := hub.NewClient(providerCred.ProviderId, nil, logger, providerCred)
	if err != nil {
		logger.Errorln("An error occurred while initializing hub client using source credential", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}
	err = sourceClient.SyncAll(databaseRepo)
	if err != nil {
		logger.Errorln("An error occurred while bulk import of resources from source", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}

	c.JSON(http.StatusOK, gin.H{"success": true, "data": providerCred})
}

func ListSource(c *gin.Context) {
	logger := c.MustGet("LOGGER").(*logrus.Entry)
	databaseRepo := c.MustGet("REPOSITORY").(database.DatabaseRepository)

	providerCredentials, err := databaseRepo.GetSources(c)
	if err != nil {
		logger.Errorln("An error occurred while storing provider credential", err)
		c.JSON(http.StatusInternalServerError, gin.H{"success": false})
		return
	}
	c.JSON(http.StatusOK, gin.H{"success": true, "data": providerCredentials})
}