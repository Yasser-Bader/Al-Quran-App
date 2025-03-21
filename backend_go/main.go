package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/quran-app/config"
	"github.com/quran-app/models"
)

var db, err = config.ConnectDB()

func main() {
	//watcher.Fun_watcher()

	// الاتصال بقاعدة البيانات

	if err != nil {
		log.Fatal("❌ فشل الاتصال بقاعدة البيانات:", err)
	}

	/*/ التأكد من وجود الجدول
	err = db.AutoMigrate(&models.Quran_text{})
	if err != nil {
		log.Fatal("❌ فشل إنشاء الجدول:", err)
	}*/

	fmt.Println("✅ تم الاتصال بقاعدة البيانات بنجاح!")
	// إنشاء راوتر باستخدام Gin
	r := gin.Default()

	// تمكين CORS للسماح لأي جهة بالوصول إلى API
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"}, // يمكنك تحديد نطاق معين بدلاً من "*" للأمان
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
		AllowCredentials: true,
	}))

	// الاتصال بقاعدة البيانات
	if err != nil {
		log.Fatal("فشل في الاتصال بقاعدة البيانات:", err)
	}

	// نقطة نهاية لاسترجاع جميع السور
	r.GET("/sura/:suraNumber", func(c *gin.Context) {
		var ayat []models.Quran_texts
		suraNumber := c.Param("suraNumber")

		result := db.Where("sura = ?", suraNumber).Find(&ayat)
		if result.Error != nil {
			c.JSON(404, gin.H{"error": result.Error.Error()})
			return
		}

		c.JSON(200, gin.H{
			"view": ayat,
		})
	})

	r.GET("/all-suars", func(c *gin.Context) {
		var suars []models.Quran_texts
		db.Find(&suars)
		c.JSON(200, gin.H{
			"view": suars,
		})
	})
	r.GET("/search", searchAyah)

	// تشغيل السيرفر
	r.Run(":8080")
}

/*func removeDiacritics(text string) string {
	var result strings.Builder
	for _, r := range text {
		if unicode.Is(unicode.Mn, r) {
			continue
		}
		result.WriteRune(r)
	}
	return result.String()
}*/

func searchAyah(c *gin.Context) {
	query := strings.TrimSpace(c.Query("q"))
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"يرجى ادخال كلمة البحث": "error"})
		return
	}
	//query = removeDiacritics(query)

	var results []models.Quran_texts
	result := db.Where("text LIKE ?", "%"+query+"%").Find(&results)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if len(results) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"لم يتم العثور على نتائج": "message"})
		return
	}
	c.JSON(200, gin.H{
		"view": results,
	})
}

// محاكاة دالة جلب النص من قاعدة البيانات
func getSurahText(id string) string {
	if id == "2" {
		return "الم * ذلك الكتاب لا ريب فيه هدى للمتقين..."
	}
	return ""
}
