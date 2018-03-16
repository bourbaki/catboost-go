package catboost

/*
#cgo linux LDFLAGS: -lcatboostmodel
#include <stdlib.h>
#include <stdbool.h>
#include <model_calcer_wrapper.h>
static char** makeCharArray(int size) {
        return calloc(sizeof(char*), size);
}

static void setArrayString(char **a, char *s, int n) {
        a[n] = s;
}

static void freeCharArray(char **a, int size) {
        int i;
        for (i = 0; i < size; i++)
                free(a[i]);
        free(a);
}
*/
import "C"

import (
	"fmt"
	"math"
	"unsafe"
)

func makeCStringArrayPointer(array []string) **C.char {
	cargs := C.makeCharArray(C.int(len(array)))
	for i, s := range array {
		C.setArrayString(cargs, C.CString(s), C.int(i))
	}
	return cargs
}

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
	for i, v := range floats {
		floatsC[i] = (*C.float)(&v[0])
	}

	catsC := make([]**C.char, nSamples)
	for i, v := range cats {
		pointer := makeCStringArrayPointer(v)
		defer C.freeCharArray(pointer, C.int(len(v)))
		catsC[i] = pointer
	}

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

	for i, _ := range results {
		result[i] = 1.0 / (1.0 + math.Exp(-results[i]))
	}
	return results, nil
}
