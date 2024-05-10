package response

import (
	"fmt"
	"reflect"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type PageResp struct {
	PageNo   int
	PageSize int
	Count    int64
	Data     interface{}
}

type PageReq struct {
	PageNo   int
	PageSize int
}

func CheckErr(err error, msg string) error {
	if err == nil {
		return nil
	}
	logrus.Error(msg)
	return err
}

func ErrRecordNotFound(err error, msg string) error {
	if err == gorm.ErrRecordNotFound || err == nil {
		return nil
	}
	logrus.Error(msg)
	return err
}

func CopyStruct(dst, src interface{}) error {
	// Ensure that src and dst are pointers
	srcVal := reflect.ValueOf(src)
	dstVal := reflect.ValueOf(dst)
	if srcVal.Kind() != reflect.Ptr || dstVal.Kind() != reflect.Ptr {
		return fmt.Errorf("both source and destination must be pointers to structs")
	}

	// Dereference pointers to get the actual struct objects
	srcVal = srcVal.Elem()
	dstVal = dstVal.Elem()

	// Check that we're dealing with structs
	if srcVal.Kind() != reflect.Struct || dstVal.Kind() != reflect.Struct {
		return fmt.Errorf("both source and destination must be structs")
	}

	// Copy data from src to dst
	for i := 0; i < srcVal.NumField(); i++ {
		srcField := srcVal.Field(i)
		dstField := dstVal.FieldByName(srcVal.Type().Field(i).Name)

		if dstField.IsValid() && dstField.CanSet() {
			// Check if the destination field can accept the source field type
			if dstField.Type() == srcField.Type() {
				dstField.Set(srcField)
			}
		}
	}

	return nil
}
