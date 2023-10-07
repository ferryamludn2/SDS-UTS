package route

import (
	"uts/database"
	"uts/models"

	"github.com/gofiber/fiber/v2"
)

// buatlah endpoint Insert Data sesuai dengan Colection Postman
func InsertData(c *fiber.Ctx) error {

	//Tulis Jawaban Code di Sini :))
	var DataUser map[string]string
	if err := c.BodyParser(&DataUser); err != nil {
		return err
	}
	User := models.User{
		Nama:     DataUser["nama"],
		Email:    DataUser["email"],
		Username: DataUser["username"],
		Password: DataUser["password"],
	}
	database.DB.Create(&User)

	return c.JSON(fiber.Map{
		"Pesan": "Data telah berhasil di tambahkan",
	})
}

// Lengkapi Code Berikut untuk untuk Mengambil data untuk semua user - user
func GetAllData(c *fiber.Ctx) error {
	DataUsers := []models.User{}
	database.DB.Find(&DataUsers)

	return c.JSON(fiber.Map{
		"data": DataUsers,
	})

}

//Lengkapi Code berikut Untuk Mengambil data dari id_user berdasarkan Parameter

func GetUserByid(c *fiber.Ctx) error {
	var DataUser models.User
	ID_User := c.Params("id_user")

	database.DB.Where("Id_user = ?", ID_User).Find(&DataUser)

	return c.JSON(fiber.Map{
		"data": DataUser,
	})
}

func Delete(c *fiber.Ctx) error {

	var user models.User

	id_user := c.Params("id_user")

	database.DB.Where("id_user = ?", id_user).Delete(user)

	return c.JSON(fiber.Map{
		"Pesan": "Data telah di hapus",
	})
}

// mengupdate data user berdasarkan id_user yang di tempatkan di parameter
func Update(c *fiber.Ctx) error {

	id_user := c.Params("id_user")

	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}
	var users models.User
	database.DB.Find(&users)
	//data yang di ubah
	//membuat variable user berdasarkan model user
	var user models.User

	update := models.User{
		Nama:     data["nama"],
		Email:    data["email"],
		Password: data["password"],
	}
	//mengambil database untuk di update

	database.DB.Model(&user).Where("id_user = ?", id_user).Updates(update)

	return c.JSON(fiber.Map{
		"Pesan": "Data User telah di Update",
	})
}
