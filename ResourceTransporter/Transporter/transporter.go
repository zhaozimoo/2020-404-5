package Transporter

import (
	"golang.org/x/crypto/ssh"
	"fmt"
	"net"
	"log"
	"os"
	"ResourceTransporter/config"
	"Logger/logger"
)

func Connect(user, password, host string, port int) (*ssh.Session, error) {
	var (
		auth    []ssh.AuthMethod
		addr    string
		client  *ssh.Client
		session *ssh.Session
		err     error
	)
	// get auth method
	auth = make([]ssh.AuthMethod, 0)
	auth = append(auth, ssh.Password(password))

	// connet to ssh
	addr = fmt.Sprintf("%s:%d", host, port)

	client, err = ssh.Dial("tcp", addr, &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(password),
		},
		//需要验证服务端，不做验证返回nil
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
	})

	// create session
	if session, err = client.NewSession(); err != nil {
		return nil, err
	}

	return session, nil
}


func RunShell(env string,projectname string) (status string,err error){
	user:=config.GetConfig().ServerInfo.User
	passwd:=config.GetConfig().ServerInfo.Passwd
	host:=config.GetConfig().ServerInfo.Host
	port:=config.GetConfig().ServerInfo.Port
	targetip:=config.GetConfig().ServerInfo.TargetIp
	targetpwd:=config.GetConfig().ServerInfo.TargetPasswd
	localpath:=config.GetConfig().ServerInfo.LocalPath
	localdevpath:=config.GetConfig().ServerInfo.LocalDevPath
	remotepath:=config.GetConfig().ServerInfo.RemotePath
	remotedevpath:=config.GetConfig().ServerInfo.RemoteDevPath
	session, err := Connect(user, passwd, host, port)
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	session.Stdout = os.Stdout
	session.Stderr = os.Stderr
//	session.Run("cat > /test/jsonfile<<EOF"+`
//                   `+fmt.Sprint(installscript)+`
//`+"EOF") //sh 命令路径
	if env=="dev"{
		const status  = "********发布测试环境*********"
		session.Run("/home/restrans/runscp.sh "+targetip+` `+remotedevpath+fmt.Sprint(projectname)+".tar.gz"+` `+localdevpath+` `+targetpwd)
		logger.Debugf("********发布测试环境*********")
		return status,err
	}else if env=="prod"{
		const status  = "********发布正式环境*********"
		session.Run("/home/restrans/runscp.sh "+targetip+` `+remotepath+fmt.Sprint(projectname)+".tar.gz"+` `+localpath+` `+targetpwd)
		logger.Debugf("********发布正式环境*********")
		return status,err
	}else{
		const status  = "********环境错误*********"
		logger.Debugf("********环境错误*********")
		return status,err
	}

}