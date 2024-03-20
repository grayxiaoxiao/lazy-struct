
package business
type Customer struct {
      
      
      
      
      
    Id uint `json:"id" gorm:"comment:id"`
    SerialNumber string `json:"serial_number" gorm:"comment:serial_number"`
    Name string `json:"name" gorm:"comment:name"`
    State int `json:"state" gorm:"comment:state"`
    Access bool `json:"access" gorm:"comment:access"`
}
