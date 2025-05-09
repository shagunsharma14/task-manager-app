package models

type Task struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Title       string  `gorm:"not null" json:"title" binding:"required"`
	Description string  `gorm:"type:text" json:"description" binding:"required"`
	Status      string  `gorm:"type:varchar(20);default:'Pending'" json:"status" binding:"required,oneof=Pending In-Progress Completed"`
	DueDate     *string `gorm:"type:datetime" json:"due_date"` // nullable, can be nil
}
