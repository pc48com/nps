type Connection struct {
    ID          string    `json:"id"`
    ClientID    int       `json:"client_id"`
    RemoteAddr  string    `json:"remote_addr"`  // 格式：IP:Port
    LocalAddr   string    `json:"local_addr"`
    StartTime   time.Time `json:"start_time"`
    BandwidthUp int64     `json:"bandwidth_up"`
    BandwidthDown int64   `json:"bandwidth_down"`
    Status      string    `json:"status"`       // connected/disconnected
}

// 新增防火墙模型
type FirewallIP struct {
    IP          string    `json:"ip"`
    AccessCount int       `json:"access_count"`
    TotalBytes  int64     `json:"total_bytes"`
    FirstSeen   time.Time `json:"first_seen"`
    LastSeen    time.Time `json:"last_seen"`
}

var (
    ConnectionMap = make(map[string]*Connection)
    IPAccessMap   = make(map[string]*FirewallIP)
    IPBlacklist   = make(map[string]bool)
)
