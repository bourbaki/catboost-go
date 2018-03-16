package catboost
/*
#cgo linux LDFLAGS: -lcatboostmodel
#include <stdbool.h>
#include <model_calcer_wrapper.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type CatBoostModel struct {
	Handler unsafe.Pointer
}

func (model *CatBoostModel) GetFloatFeaturesCount() int {
	return int(C.GetFloatFeaturesCount(model.Handler))
}

func (model *CatBoostModel) GetCatFeaturesCount() int {
	return int(C.GetCatFeaturesCount(model.Handler))
}

func LoadCatBoostModelFromFile(filename string) (*CatBoostModel, error) {
	model := &CatBoostModel{}
	model.Handler = C.ModelCalcerCreate() 
	if(!C.LoadFullModelFromFile(model.Handler, C.CString(filename))) {
		return nil, fmt.Errorf("Cannot open model")
	}
	return model, nil
}

func (model *CatBoostModel) Predict(floats []float32, cats []string) (float64, error) {
	return 0.0, nil
}