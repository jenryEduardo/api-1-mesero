package application

import "api-main/pedido/domain"

type UpdateStatusPedido struct {
	repo domain.Ipedido
}

func NewUpdateStatusPedido(repo domain.Ipedido) *UpdateStatusPedido {
	return &UpdateStatusPedido{repo: repo}
}

func (c *UpdateStatusPedido) Execute(id int, status string) error {
	return c.repo.UpdateStatus(id, status)
}