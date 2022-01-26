package cloud

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"

	"cloud.google.com/go/storage"
	"github.com/gin-gonic/gin"
	"google.golang.org/api/option"
	"google.golang.org/appengine"
)

var (
	storageClient *storage.Client
)

func HandleFileUploadToBucket(c *gin.Context) {
	bucket := "image_services_golang" //your bucket name

	var err error

	ctx := appengine.NewContext(c.Request)

	storageClient, err = storage.NewClient(ctx, option.WithCredentialsFile("key.json"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	f, uploadedFile, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	defer f.Close()

	sw := storageClient.Bucket(bucket).Object(uploadedFile.Filename).NewWriter(ctx)

    // fmt.Print(sw.Attrs().Name)
	if _, err := io.Copy(sw, f); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}

	if err := sw.Close(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"error":   true,
		})
		return
	}
var urlLink = "https://storage.googleapis.com"
	u, err := url.Parse(urlLink + "/" + bucket + "/" + sw.Attrs().Name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
			"Error":   true,
		})
		return
	}


	c.JSON(http.StatusOK, gin.H{
		"message":  "file uploaded successfully",
		"pathname": urlLink+u.EscapedPath(),
		
	})

}
func GetUrlFile(c *gin.Context) {
	bucket := "image_services_golang"
	filename := "product3.png"
	method := "PUT"
	expires := time.Now().Add(time.Second * 60)

	url, err := storage.SignedURL(bucket, filename, &storage.SignedURLOptions{
		GoogleAccessID: "image-bucket-services@product-image-services.iam.gserviceaccount.com",
		PrivateKey:     []byte("-----BEGIN PRIVATE KEY-----\nMIIEvgIBADANBgkqhkiG9w0BAQEFAASCBKgwggSkAgEAAoIBAQDd0sH6B2Xszsm7\nLZRZtFKAIxxB1xdRk/tdPTPcKHGU/SJBL/e1I5t0UM1qletxQME7o+lxvMl29Q87\n8Q1vHhEvLlhnmE1vn3OtI7HnxQ/AQyPPACCYF69IKh0mFutnOfER2Rdx9yHPidLi\nRchcH70VnfSLRsmiW++A4k/ggl4DsR96wE2kHlBq6Y5wMXzkW33HPahpnElHEJMG\nrH6HiNMIFVMbHwtPCl00NsU1q+6t9RxQBwRoCSEbF5krC1adx9W02hX5oyzi8cP6\nK0KABhAx4RQzz4Dc70XVcPlpn5GK71SGjdKyGAX6qJ+lBXx0rtzdI06Ebzvk1MLP\nR0pBYIGhAgMBAAECggEAEIhy9/xZbNAgUcTEyfdST1XelysthCFqEqHbRclXZt4D\nkbV+KkhyP5X+q3crVlR32oa94WHrZ6P9+88ce9fTOs+i2+zKd7t8Ew9mJQLHaOJw\nLq9hHojkYfXumH5MkxPq23RRZqd/ZAEvIDmIhJJChPQSDBftcx4UUKg9gRuiX3au\nbIlqUVnnPrRoosUSwMKlJoFVLKtgHsIGg1RQpisVhphUEEIdovQXBnI65DLax0aI\nwlN1od6o0B7EiWA2xvGEsUAGzDHXo9ERbFT+uSSWbX62QITGEQjmzg7dn6mOVjBV\n16wkNjphR55xV6KAqR1dtMMsZ6zxmJIfUqzOyISO+QKBgQD9489UxhgiUPb9dgxV\nHslfdiOzgjowF0y06nU5yyBqnYFsrbFTeDjED96kF+Fse7q2K9pMDGrdIKuhd5Cd\n+Lb6/xZyq1H4rye7MCiol3XxrAxXjaumGhHUBbxQlN17jDOZIHXV9eSHwQkKvtl/\nVRgGTAC2t8zl37wBN/hs54jbuQKBgQDfqridDu/KKgd1tVciDgTY34L2ZA5M4aQv\nYOaa54PyA6PsKWernGxAtLLv41kpVfzXp3fiRk/Hjo8gQ9Pp+S/foQSulP2MMHNH\nGj10Obc00z7R9FOvy6o3XtipZ/dx+vNqUjY1EaZXWA4TXzTMNniWc1IN0D7FLesv\ncET0daJZKQKBgQD9e61M3lrKKDvw4yN8+Lfk73bFioe97AVRu6Q+h2deCtNlRiV9\nSNKkLZQEETOntAC+URoqQ1uOW0gAdfeVQPSvtG8dHZ9Bwt7QLUzqxg2jtDq+T1vJ\nAs45+WACtB5Nc7UwdRAxBsecIkZ8y/8q+jJ6Vvd/dhLEj5SNQuxtDt29QQKBgG+5\n1uhVisCAyCMrR3Aycodm9wNfLamH2Tz1eZwNY+KjoOGaOTgHNigIW43rEiHM2zVa\naU81ciqr8qDaYOPyXtClnTIcKJ87oIn2+JWzMuoHT80O8DLTWJ66GR5eWcOs6KTG\nll9iBqaAzN8uYrBT0V7OEkHmMUTL0DxtJ3S5wjQpAoGBAJtqOewf28v0/n9CFlzv\nTctN/TWTrCgxcGWXvaOaggf2ccoNtFwzpmjdDr+3fLOpeJigwIh1KipPkbzo5lVo\nBvwfZYSL3FWNNkor2e9nDYRJl/fFHXaVMKEQ3rx1JFrY6jxO+2cnmMNHauoLqRqo\nQQ6puCIK6OSoLAHMGY0NGxSK\n-----END PRIVATE KEY-----\n"),
		Method:         method,
		Expires:        expires,
	})
	if err != nil {
		fmt.Println("Error " + err.Error())
	}
	fmt.Println("URL = " + url)

}

// part 1 https://medium.com/wesionary-team/golang-image-upload-with-google-cloud-storage-and-gin-part-1-e5e668c1a5e2
// part 2 https://medium.com/wesionary-team/golang-image-upload-with-google-cloud-storage-and-gin-part-2-99f4a642e06a
