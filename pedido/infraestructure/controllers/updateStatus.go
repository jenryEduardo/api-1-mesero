package controllers

import (
	"api-main/pedido/application"
	"api-main/pedido/infraestructure"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func UpdateStatusPedido(c *gin.Context) {

	idString := c.Param("idPedido")
	id, err := strconv.Atoi(idString)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "no se pudo encontrar un id verifique su solicitud"})
		return
	}

	var status string

	if err := c.ShouldBindJSON(&status); err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": "no se pudo decodificar el archivo json"})
		return
	}

	repo := infraestructure.NewMySQLRepository()
	useCase := application.NewUpdateStatusPedido(repo)

	if err := useCase.Execute(id, status); err != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "no se pudo realizar la solicitud"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ok": "se actualizo el pedido correctamente"})
}