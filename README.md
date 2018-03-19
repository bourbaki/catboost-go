# CatBoost Wrapper for Go
Simple wrapper of [CatBoost C library](https://tech.yandex.com/catboost/doc/dg/concepts/c-plus-plus-api_dynamic-c-pluplus-wrapper-docpage/) for prediction

## Installation
CatBoost library is assumed to be installed and all its includes and library files are assumed to be found in corresponding paths. One way to do it is using environment variables:
```sh
git clone https://github.com/catboost/catboost.git
cd catboost/catboost/libs/model_interface && ../../ya make -r .
export CATBOOST_DIR=$(pwd)
export C_INCLUDE_PATH=$CATBOOST_DIR:$C_INCLUDE_PATH
export LIBRARY_PATH=$CATBOOST_DIR:$LIBRARY_PATH
export LD_LIBRARY_PATH=$CATBOOST_DIR:$LD_LIBRARY_PATH
```
The other way is to put compiled library files and include files to default search diretories (`/usr/local/lib`, `/usr/local/include`).
If everything above is properly configured then a simple `go get` command will do the trick:
```
go get -u github.com/ma3axaka/catboost-go
```
