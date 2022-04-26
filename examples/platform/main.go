package main

import (
	"fmt"
	//enva "my_platform/env"
)

func main() {
	awsEksParams := GetAwsEksParams()
	fmt.Println(awsEksParams)
	//aws := AwsEksParams
	//aws.CreateAwsEks(awsEksParams, "cluster-machinePoll.yaml")
	t := test{
		name: "qiongjie",
	}
	t.print()
	//aws.CreateAwsEks(awsEksParams, "cluster-machinePoll.yaml")
	//	if _, err := CreateAwsEks(awsEksParams, "cluster-machinePoll.yaml"); err != nil {
	//		fmt.Println("Create AWS EKS got Error: ", err)
	//		os.Exit(1)
	//	}

}

func GetAwsEksParams() AwsEksParams {
	//myaws := new(env.AwsEksParams)
	myaws := AwsEksParams{
		Namespace:   "mycluster",
		ClusterName: "capa-test",
		KubeadmControlPlane: KubeadmControlPlane{
			Replicas: 2,
			Version:  "v1.21.1",
		},
		AWSMachineTemplate: AWSMachineTemplate{
			InstanceType: "t2.medium",
			SshKeyName:   "bs-kevin-test",
		},
		AWSMachinePool: AWSMachinePool{
			InstanceType: "t2.medium",
			SshKeyName:   "bs-kevin-test",
			MaxSize:      3,
			MinSize:      1,
		},
		MachinePool: MachinePool{
			Replicas: 2,
			Version:  "v1.21.1",
		},
	}

	//	awsEksParam := AwsEksParams{
	//		aws.Namespace:   "mycluster",
	//		ClusterName: "capa-test",
	//		KubeadmControlPlane: KubeadmControlPlane{
	//			Replicas: 2,
	//			Version:  "v1.21.1",
	//		},
	//		AWSMachineTemplate: AWSMachineTemplate{
	//			InstanceType: "t2.medium",
	//			SshKeyName:   "bs-kevin-test",
	//		},
	//		AWSMachinePool: AWSMachinePool{
	//			InstanceType: "t2.medium",
	//			SshKeyName:   "bs-kevin-test",
	//			MaxSize:      3,
	//			MinSize:      1,
	//		},
	//		MachinePool: MachinePool{
	//			Replicas: 2,
	//			Version:  "v1.21.1",
	//		},
	//	}
	myaws.CreateAwsEks("cluster-machinePoll.yaml")
	return myaws
}
