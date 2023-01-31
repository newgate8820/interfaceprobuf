
## 协议文件生成
### protoc 镜像构建
```shell
docker build -t protoc:laster .
```
### protoc 生成 proto grpc 文件 e.g script_path要生成的脚本路径 目录下放  文件名.proto
### cd ~/message
```shell
script_path=updates
script_name=updates
docker run -it --rm -v "${PWD}":/src  protoc:laster \
protoc -I "${script_path}"  \
-I ./ \
--go_opt=paths=source_relative \
--go-grpc_opt=paths=source_relative \
--go_out="${script_path}" --go-grpc_out=:"${script_path}" "${script_path}"/${script_name}.proto  
```