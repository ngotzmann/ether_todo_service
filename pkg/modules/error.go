package modules

import (
"fmt"
"log"
"strconv"

"github.com/spf13/viper"
)

const (
	InternalServerError = "InternalServerError"
	ValidationError     = "ValidationError"
	DatabaseError       = "DatabaseError"
	FileNotFoundError   = "FileNotFoundError"
	AuthenticationError = "AuthenticationError"
)

var Gorrors []Gorror

type Gorror struct {
	ErrorTitle           string `json:"X-Error-Titel"`
	ErrorCode            int    `json:"X-Error-Code"`
	ErrorMessage         string `json:"X-Error-Message"`
	SpecificErrorMessage string `json:"X-Specific-Error-Message"`
	RequestId            string `json:"X-Request-Id"`
}

func (g *Gorror) Error() string {
	return fmt.Sprintf(
		"Title: " + g.ErrorTitle +
			"; Code: " + strconv.Itoa(g.ErrorCode) +
			"; Message: " + g.ErrorMessage +
			"; SpecificMessage: " + g.SpecificErrorMessage +
			"; RequestId: " + g.RequestId)
}

// Initialise default errors and errors of the errors.yaml
// Default value of path is `config/errors.yaml`, it is not necessary to set
// The method will load default errors and errors from yaml in in Gorrors array
func InitErrors(cfg Config) {

	if cfg.GorrorFilePath != "" {
		readYaml(cfg.GorrorFilePath)
	}
	Gorrors = append(Gorrors, *getDefaultGorrors()...)
}

func readYaml(path string) {
	viper.SetConfigName("errors")
	viper.AddConfigPath(path)
	err := viper.ReadInConfig()
	if err != nil {
		log.Panic(err)
	}
	err = viper.UnmarshalKey("gorrors", &Gorrors)
	if err != nil {
		log.Panic(err)
	}
}

//Check if the given error is from type Gorror and return Gorror struct
//If the error is not from type Gorror return nil
func CastErrorToGorror(err error) *Gorror {
	gorr, ok := err.(*Gorror)
	if ok {
		return gorr
	}
	return nil
}

// Return error interface as Gorror type
// From the title the error will be created with the values from yaml or defaults
// If there is no error with the title, an `ErrorNotFound` error return
func CreateError(title string, specErrMess string) error {
	var rst Gorror
	for _, gs := range Gorrors {
		if title == gs.ErrorTitle {
			rst.ErrorTitle = gs.ErrorTitle
			rst.ErrorCode = gs.ErrorCode
			rst.ErrorMessage = gs.ErrorMessage
			rst.RequestId = gs.RequestId
		}
	}
	rst.SpecificErrorMessage = specErrMess

	if rst.ErrorTitle == "" {
		rst.ErrorTitle = "ErrorNotFound"
		rst.ErrorCode = 404
		rst.ErrorMessage = "Error with given 'ErrorTitle' was not found"
		rst.SpecificErrorMessage = title
		log.Println(rst)
	}
	return &rst
}

func getDefaultGorrors() *[]Gorror {
	return &[]Gorror{
		{
			ErrorTitle:   InternalServerError,
			ErrorCode:    500,
			ErrorMessage: "The server encountered an internal error and was unable to complete your request",
		},
		{
			ErrorTitle:   ValidationError,
			ErrorCode:    409,
			ErrorMessage: "Please check your input data",
		},
		{
			ErrorTitle:   DatabaseError,
			ErrorCode:    500,
			ErrorMessage: "",
		},
		{
			ErrorTitle:   FileNotFoundError,
			ErrorCode:    404,
			ErrorMessage: "File was not found maybe the path is incorrect, the file is missing, there are wrong access permissions or something else",
		},
		{
			ErrorTitle:   AuthenticationError,
			ErrorCode:    401,
			ErrorMessage: "Authentication was not successful, check your given data",
		},
	}
}
