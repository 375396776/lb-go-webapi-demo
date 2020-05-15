package invoice

//通用传入实体
type Common struct {
	PlatformCode string `json:"platformCode" binding:"required"`
	SignType     string `json:"signType" binding:"required"`
	Sign         string `json:"sign" binding:"required"`
	Format       string `json:"format" binding:"required"`
	Timestamp    string `json:"timestamp" binding:"required"`
	Version      string `json:"version" binding:"required"`
	SerialNo     string `json:"serialNo" binding:"required"`
}
