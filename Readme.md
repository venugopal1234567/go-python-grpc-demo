## Create python env 
virtualenv -p python3 env

# Activate python env. windows user 
 source ./env/Scripts/activate

# Install python dependencies
pip install -r requirements.txt

## To run server
go run main.go

## To run client
python client/main.py