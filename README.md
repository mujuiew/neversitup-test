# neversitup-test

![[Neversitup] Candidate Test](https://docs.google.com/document/d/1N3owA8z3MrvnPMqUcfEsi-Q8cCIRyeaWpTXE1el0cME/edit)


📖 neversitup-test 📖 

All the programs this project owns belongs inside the cmd/ folder. The folders under cmd/ are always named for each program that will be built.

| Module  | Is module is in service? |
| ------------- | ------------- |
| API  | :ok:  |

## Tools

- Go 1.20 or lastest
- Docker and Docker-Compose

## Lasted projec of init project structure and defined convention
```
/myproject
  /_cicd
    /build
      /dockers
  /cmd  
    /servapi
      main.go
  /internal
    /controller
      /pkg1
        pkg1.go
      /pkg2
        pkg2.go
    /core
      /pkg1
        pkg1.go
      /pkg2
        pkg2.go
  go.mod
  go.sum
  makefile
  README.md
```
  /cmd  : จุดเริ่มต้นสำหรับโครงการของคุณ แต่ละไดเร็กทอรีย่อยภายใต้ สำหรับไฟล์ main
  /internal/controller : ไว้สำหรับงานในส่วนของ controller
  /internal/core : core process
  makefile : คำสั่งโดยย่อให้ง่ายต่อการใช้งาน
