package utilities

import (
  "sort"
  "go-domains/models"
)

type OrderServersBy func(server_one, server_two models.Server) bool

func (order_servers_by OrderServersBy) Sort(servers []*models.Server) {
	server_sorter := &serverSorter{
    servers:           servers,
    order_servers_by:  order_servers_by,
  }
	sort.Sort(server_sorter)
}

type serverSorter struct {
	servers           []*models.Server
	order_servers_by  func(server_one, server_two models.Server) bool
}

func (server_sorter *serverSorter) Len() int {
	return len(server_sorter.servers)
}

func (server_sorter *serverSorter) Swap(i, j int) {
	server_sorter.servers[i],
  server_sorter.servers[j] = server_sorter.servers[j],
  server_sorter.servers[i]
}

func (server_sorter *serverSorter) Less(i, j int) bool {
  return server_sorter.order_servers_by(
      *server_sorter.servers[i],
      *server_sorter.servers[j],
  )
}
