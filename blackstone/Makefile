proto:
	python3 -m pip install -r requirements.txt
	python3 -m grpc_tools.protoc \
	    -I../server/ --python_out=. --grpc_python_out=. \
	    ../server/bstone.proto