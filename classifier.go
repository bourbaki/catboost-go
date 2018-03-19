package catboost

import "math"

// BinaryClassifer is wrapper over model object that add methods for binary classification
type BinaryClassifer struct {
	Model *Model
}

func sigmoid(probit float64) float64 {
	return 1.0 / (1.0 + math.Exp(-probit))
}

// LoadBinaryClassifierFromFile loads binary classifier from file
func LoadBinaryClassifierFromFile(filename string) (*BinaryClassifer, error) {
	model, err := LoadFullModelFromFile(filename)
	if err != nil {
		return nil, err
	}
	return &BinaryClassifer{Model: model}, nil
}

// PredictProba returns sigmoid scores which could be interpreted like probability
func (bc *BinaryClassifer) PredictProba(floats [][]float32, floatLength int, cats [][]string, catLength int) ([]float64, error) {
	results, err := bc.Model.CalcModelPrediction(floats, floatLength, cats, catLength)
	if err != nil {
		return nil, err
	}
	for i := range results {
		results[i] = sigmoid(results[i])
	}
	return results, nil
}

// Close deletes model handler
func (bc *BinaryClassifer) Close() {
	bc.Model.Close()
}
