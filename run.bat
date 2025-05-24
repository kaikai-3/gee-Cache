@echo off
chcp 65001 >nul
title GeeCache Test

REM Build project
echo Building project...
go build -o server.exe
if %ERRORLEVEL% NEQ 0 (
    echo Build failed, please check the error message
    pause
    exit /b 1
)

echo Starting servers...
REM Start three server instances
start "Cache Server 1" cmd /c "server.exe -port=8001"
start "Cache Server 2" cmd /c "server.exe -port=8002"
start "API Server" cmd /c "server.exe -port=8003 -api=1"

REM Wait for servers to start
timeout /t 2 /nobreak >nul

echo.
echo ===== Starting Test =====
echo.

REM Send test requests
echo Sending test request 1...
curl "http://localhost:9999/api?key=Tom"
echo.

echo Sending test request 2...
curl "http://localhost:9999/api?key=Tom"
echo.

echo Sending test request 3...
curl "http://localhost:9999/api?key=Tom"
echo.

echo.
echo ===== Test Complete =====
echo Press any key to exit...
pause >nul

REM Clean up processes
taskkill /F /IM server.exe >nul 2>&1
exit