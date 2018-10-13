rd /s/q demo-win
md demo-win
go build -o demo.exe
COPY demo.exe demo-win\
COPY app.conf demo-win\
XCOPY view\*.* demo-win\view\  /s /e
