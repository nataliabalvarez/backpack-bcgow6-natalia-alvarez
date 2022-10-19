package store
import(
	"os"
	"encoding/json"
)

type Store interface {
	Read(data interface{}) error
	Write(data interface{}) error
}

type Type string

const (
	FileType Type = "file"
	MonoType Type = "mongo"
)

func NewStore(store Type, fileName string) Store {
	switch store {
	case FileType: 
	 return &fileStore{fileName}
	}
	return nil
}

type fileStore struct {
	FilePath string
}

func (fs *fileStore) Read(data interface{}) error {
	file, err := os.ReadFile(fs.FilePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(file, &data) //stores json into data
	//reads json, returns nil or unsmarshals error
}

func (fs *fileStore) Write(data interface{}) error {
	file, err := json.MarshalIndent(data, "", "   ")
	if err != nil {
		return err
	}

	return os.WriteFile(fs.FilePath, file, 0644)
}