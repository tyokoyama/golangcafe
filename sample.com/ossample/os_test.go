package ossample

import (
	"os"
	"testing"
)

func openWorkDirectory() (file *os.File, err error) {
	// パスで~は使えない！
	homedir := os.Getenv("HOME")

	file, err = os.Open(homedir + "/golang/work")
	return	
}

func openCurrentWorkDirectory() (file *os.File, err error) {
	homedir := os.Getenv("HOME")
	file, err = os.Open(homedir + "/work")
	return
}

func TestChdir(t *testing.T) {
	_, err := openWorkDirectory()
	if os.IsNotExist(err) {
		t.Logf("Directory Not Found\n")
		t.Skip()
	}

	homedir := os.Getenv("HOME")
	if err := os.Chdir(homedir + "/golang/work"); err != nil {
		t.Fatalf("Chdir Error %v", err)
	}

}

func TestMkdirAll(t *testing.T) {
	homedir := os.Getenv("HOME")
	_, err := openWorkDirectory()
	if os.IsExist(err) {
		// RemoveAll()は指定したパスとその配下のファイルを全て削除する。
		if err := os.RemoveAll(homedir + "/golang/work"); err != nil {
			t.Fatalf("os.RemoveAll Error %v", err)
		}
	}

	// 複数のディレクトリ階層はMkdirAll()でないと作れない。
	if err := os.MkdirAll(homedir + "/golang/work", 0755); err != nil {
		t.Fatalf("MkdirAll Error %v", err)
	}
}

func TestMkdir(t *testing.T) {
	homedir := os.Getenv("HOME")
	_, err := openCurrentWorkDirectory()
	if os.IsExist(err) {
		if err := os.Remove(homedir + "/work2"); err != nil {
			t.Fatalf("os.RemoveAll Error %v", err)
		}
	}

	if err2 := os.Mkdir(homedir + "/work2", 0755); err2 != nil {
		t.Fatalf("Mkdir Error %v", err2)
	}
}

func TestEnviron(t *testing.T) {
	// 環境変数一覧
	for pos, str := range os.Environ() {
		t.Logf("Environ[%d] = %s", pos, str)
	}
}

func TestTempDir(t *testing.T) {
	// 
}