all: Test.cs test.pb.validate.cs

Test.cs: test.proto
	protoc --csharp_out=. -I . -I ../.. test.proto

test.pb.validate.cs: test.proto
	protoc --validate_out="lang=csharp:." -I . -I ../.. test.proto

clean:
	rm -rf Test.cs test.pb.validate.cs