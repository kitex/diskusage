## Clone this project 
```git clone git clone https://github.com/kitex/diskusage.git
cd diskusage```
## This is utility application for showing disk usage by partition 💬

## Initialize
```go mod init main
get golang.org/x/sys/unix```

##  Set build architecture and run build command 👋
```set GOARCH=amd64
set GOOS=linux
go build
```
# Example command
```sudo ./main -mount="/" -sort="diskusage"```

# Here are few application parameters
`-mount` 

 - argument is for mount point
 
`-sort` 

 - argument for sorting by diskusage or mount point.
 
* The default is sort by mount point *

Sometime files cannot be read due to system file or transient files. These files which are not read (specially in / mount point) are written in log file called execution.log

## I'm a Developer 💻,Motivator 📸, and Avid Reader 


## Running the application: (Need to run as sudo user because root directory has permission issue)

```
sudo ./main -mount="/" -sort="diskusage"
```

### 🤝 Connect with me: https://www.linkedin.com/in/sugandha-amatya/


## 🔭 I'm currently working on

- React Project
- Clickhouse
- My Health
- My Linux Skill / Azure / DevOps
- My SRE skill

## 🌱 I'm currently learning

- 📱 React Native
- Clickhouse
- Prometheus
- Reliability
- GoLang


