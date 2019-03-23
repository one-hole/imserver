package connections

/*
按照我们的实际情况来看、这个地方需要先存在有 Manager
所以最终路由应该如右侧: DELETE /managers/:manager_id/connections/:id
*/
// func loadConnection(c *gin.Context) *sockets.Client {
// 	managerID, err := strconv.Atoi(c.Param("manager_id"))
// 	connID, err := strconv.Atoi(c.Param("id"))

// 	if err != nil {
// 		panic(err)
// 	}
// 	manager := managers.ManagerByIndex(managerID)
// 	conn := connByManagerAndID(manager, connID)
// 	return conn
// }

// func connByManagerAndID(m *sockets.ClientManager, index int) *sockets.Client {
// 	return nil
// }
