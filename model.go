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

func (model *CatBoostModel) Close() {
	C.ModelCalcerDelete(model.Handler)
}

func LoadCatBoostModelFromFile(filename string) (*CatBoostModel, error) {
	model := &CatBoostModel{}
	model.Handler = C.ModelCalcerCreate()
	if !C.LoadFullModelFromFile(model.Handler, C.CString(filename)) {
		return nil, fmt.Errorf("Cannot open model")
	}
	return model, nil
}

func (model *CatBoostModel) Predict(floats [][]float32, floatLength int, cats [][]string, catLength int) ([]float64, error) {
	nSamples := len(floats)
	results := make([]float64, nSamples)
	floatsC := make([]*C.float, nSamples)
	catsC := make([]**C.char, nSamples)
	C.CalcModelPrediction(
		model.Handler,
		C.size_t(nSamples),
		(**C.float)(&floatsC[0]),
		C.size_t(floatLength),
		(***C.char)(&catsC[0]),
		C.size_t(catLength),
		(*C.double)(&results[0]),
		C.size_t(nSamples),
	)
	return results, nil
}
