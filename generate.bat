@echo off
setlocal

:: Установка переменных окружения
set PROTO_PATH=cmd\api\loms\v1
set PROTO_FILE=%PROTO_PATH%\loms.proto
set OUT_DIR=loms\pkg\api\loms\v1

:: Создание необходимых директорий
if not exist %PROTO_PATH% mkdir %PROTO_PATH%
if not exist %OUT_DIR% mkdir %OUT_DIR%
if not exist vendor-proto mkdir vendor-proto

:: Запуск protoc
C:\protoc\bin\protoc.exe ^
-I %PROTO_PATH% ^
-I vendor-proto ^
--go_out=%OUT_DIR% ^
--go_opt=paths=source_relative ^
--go-grpc_out=%OUT_DIR% ^
--go-grpc_opt=paths=source_relative ^
%PROTO_FILE%

:: Проверка на ошибки
if %ERRORLEVEL% neq 0 (
    echo Error generating protobuf files
    exit /b %ERRORLEVEL%
)

go mod tidy

echo Protocol buffer files generated successfully
