package main

import (
	"fmt"
	"os"
	"os/exec"
	"text/template"
)

//type CreateEks []string

type KubeadmControlPlane struct {
	Replicas int
	Version  string
}

type AWSMachineTemplate struct {
	InstanceType string
	SshKeyName   string
}

type AWSMachinePool struct {
	InstanceType string
	SshKeyName   string
	MaxSize      int
	MinSize      int
}

type MachinePool struct {
	Replicas int
	Version  string
}

type AwsEksParams struct {
	Namespace   string
	ClusterName string
	KubeadmControlPlane
	AWSMachineTemplate
	AWSMachinePool
	MachinePool
}

//GetAwsEksParams
func RenderTemplateToFile(awsEksParams AwsEksParams, fileName string) string {
	// Read from this file
	tplFileName := "./cluster-machinePoll.tpl"
	tplEks, err := template.ParseFiles(tplFileName)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	// Output to this file
	outFileName, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Create file: ", err)
		os.Exit(1)
	}

	//K8sTemplate := ioutil.ReadFile(fileName)
	// Render tpl file
	err = tplEks.Execute(outFileName, awsEksParams)
	if err != nil {
		panic(err)
		os.Exit(1)
	}
	return fileName
}

func CreateAwsEks(awsEksParams AwsEksParams, fileName string) {
	manifests := RenderTemplateToFile(awsEksParams, fileName)
	apply := exec.Command("/usr/bin/kubectl apply -f %s", manifests)
	if _, err := apply.Output(); err != nil {
		fmt.Println("apply error: ", err)
		os.Exit(1)
	}

}
