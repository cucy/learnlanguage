package handlers

import (
	"net/http"
	"log"
	"path/filepath"
	"os"
	"io"
	"image/png"
	"github.com/nfnt/resize"
	"strings"
)

type UploadImageForm struct {
	FieldNames []string
	Fields     map[string]string
	Errors     map[string]string
}

// 显示上传表单
func DisplayUploadImageForm(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {
	RenderTemplate(w, "./templates/uploadimageform.html", u)
}

// 处理图片

func ProcessUploadImage(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {

	// 获取文件,首部信息
	file, fileheader, err := r.FormFile("imagefile")

	if err != nil {
		log.Println("Encountered error when attempting to read uploaded file: ", err)
		return
	}

	// uuid生成
	randomFileName := GenerateUUID()

	if fileheader != nil {

		// 获取扩展名
		extension := filepath.Ext(fileheader.Filename)
		r.ParseMultipartForm(32 << 20) // int64

		defer file.Close()

		// 写入文件
		imageFilePathWithoutExtension := "./static/uploads/images/" + randomFileName

		// 打开文件
		f, err := os.OpenFile(imageFilePathWithoutExtension+extension, os.O_WRONLY|os.O_CREATE, 0666)

		if err != nil {
			log.Println(err)
			return
		}

		defer f.Close()
		// 拷贝到新文件
		io.Copy(f, file)

		// 生成新的文件名
		thumbImageFilePath := imageFilePathWithoutExtension + "_thumb.png"

		// 打开文件
		originalimagefile, err := os.Open(imageFilePathWithoutExtension + extension)

		if err != nil {
			log.Println(err)
			return
		}

		// 处理图像
		img, err := png.Decode(originalimagefile)

		if err != nil {
			log.Println("Encountered Error while decoding image file: ", err)
			return
		}

		thumbImage := resize.Resize(270, 0, img, resize.Lanczos3)
		thumbImageFile, err := os.Create(thumbImageFilePath)

		if err != nil {
			log.Println("Encountered error while resizing image:", err)
			return
		}

		defer thumbImageFile.Close()

		png.Encode(thumbImageFile, thumbImage)

		m := make(map[string]string)
		m["thumbnailPath"] = strings.TrimPrefix(imageFilePathWithoutExtension, ".") + "_thumb.png"
		m["imagePath"] = strings.TrimPrefix(imageFilePathWithoutExtension, ".") + ".png"

		RenderTemplate(w, "./templates/imagepreview.html", m)

	} else {
		w.Write([]byte("Failed to process uploaded file!"))
	}
}

func ValidateUploadImageForm(w http.ResponseWriter, r *http.Request, u *UploadImageForm) {

	ProcessUploadImage(w, r, u)

}

// 处理器
func UploadImageHandler(w http.ResponseWriter, r *http.Request) {
	u := UploadImageForm{}
	u.Fields = make(map[string]string)
	u.Errors = make(map[string]string)

	switch r.Method {

	case "GET":
		DisplayUploadImageForm(w, r, &u)
	case "POST":
		ValidateUploadImageForm(w, r, &u)
	default:
		DisplayUploadImageForm(w, r, &u)
	}
}
