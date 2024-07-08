package files

import "testing"

func TestZipFileDir(t *testing.T) {
	path := "C:\\Users\\GYJ\\Downloads\\Furore"

	path, err := ZipFileDir(path)

	if err != nil {
		t.Error(err)
	}

	t.Log("success")
	t.Log(path)
}

func TestCheckAndCreatePath(t *testing.T) {
	path := "/usr/local/conf/ai_chat/rsa_pub.key"

	CheckAndCreatePath(path)

	t.Log("success")
}
