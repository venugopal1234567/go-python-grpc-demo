# To generate GoLang code run this command
protoc --proto_path=. --go-grpc_out=go  --go_out=go  score.proto

#To generate Python code run the following command
python -m grpc_tools.protoc -I. --python_out=python --grpc_python_out=python score.proto