package main

import "./car"

func main() {
	assembly := car.NewBuilder()

	familyCar := assembly.Wheels(car.SportsWheels).TopSpeed(50 * car.MPH).Build()
	familyCar.Drive()

	sportsCar := assembly.Wheels(car.SteelWheels).TopSpeed(150 * car.MPH).Build()
	sportsCar.Drive()
}

// gorm 的链式结构就是 builder 模式的实践
// err := db.WithContext(ctx).Debug().Table("abc").Where("a = ?", 3).Find(&ret).Error
// stmt := db.WithContext(ctx).Debug().Table("abc").Updates(&dao)
// n, err := stmt.RowAffected, stmt.Error
