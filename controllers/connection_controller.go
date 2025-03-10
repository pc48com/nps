// TCP连接详情接口
func GetTCPConnections(c *gin.Context) {
    list := make([]models.Connection, 0)
    for _, conn := range models.ConnectionMap {
        if conn.Status == "connected" {
            list = append(list, *conn)
        }
    }
    c.JSON(200, gin.H{"code": 200, "data": list})
}

// 防火墙相关接口
func GetRecentIPs(c *gin.Context) {
    list := make([]models.FirewallIP, 0)
    now := time.Now()
    for _, ip := range models.IPAccessMap {
        if now.Sub(ip.FirstSeen) <= 24*time.Hour {
            list = append(list, *ip)
        }
    }
    c.JSON(200, gin.H{"code": 200, "data": list})
}

func BlockIP(c *gin.Context) {
    ip := c.Query("ip")
    models.IPBlacklist[ip] = true
    c.JSON(200, gin.H{"code": 200, "msg": "封禁成功"})
}

func UnblockIP(c *gin.Context) {
    ip := c.Query("ip")
    delete(models.IPBlacklist, ip)
    c.JSON(200, gin.H{"code": 200, "msg": "解封成功"})
}

func GetBlacklist(c *gin.Context) {
    list := make([]string, 0)
    for ip := range models.IPBlacklist {
        list = append(list, ip)
    }
    c.JSON(200, gin.H{"code": 200, "data": list})
}
