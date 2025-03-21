/* package main

import (
    "github.com/gin-gonic/gin"
    "github.com/gin-contrib/cors"
)

func main() {
    r := gin.Default()

    // تمكين CORS للسماح لأي جهة بالوصول إلى API
    r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"}, // يمكنك تحديد نطاق معين بدلاً من "*" للأمان
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Accept"},
        AllowCredentials: true,
    }))

    r.GET("/sura/:id", func(c *gin.Context) {
        id := c.Param("id")
        // افترض أن لديك دالة getSurahText(id) تجلب نص السورة من قاعدة البيانات
        surahText := getSurahText(id)
        if surahText == "" {
            c.JSON(404, gin.H{"error": "لم يتم العثور على السورة"})
            return
        }
        c.JSON(200, gin.H{"sura": surahText})
    })

    r.Run(":8080")
}

// محاكاة دالة جلب النص من قاعدة البيانات
func getSurahText(id string) string {
    if id == "2" {
        return "الم * ذلك الكتاب لا ريب فيه هدى للمتقين..."
    }
    return ""
}*/