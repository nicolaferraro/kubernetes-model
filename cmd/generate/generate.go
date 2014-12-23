package main

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"time"

	kapi "github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1beta2"
	kutil "github.com/GoogleCloudPlatform/kubernetes/pkg/util"

	buildapi "github.com/openshift/origin/pkg/build/api"
	deployapi "github.com/openshift/origin/pkg/deploy/api"
	imageapi "github.com/openshift/origin/pkg/image/api"
	routeapi "github.com/openshift/origin/pkg/route/api"

	"github.com/csrwng/origin-schema-generator/pkg/schemagen"
)

type Schema struct {
	PodList                   kapi.PodList
	ReplicationControllerList kapi.ReplicationControllerList
	ServiceList               kapi.ServiceList
	BuildList                 buildapi.BuildList
	BuildConfigList           buildapi.BuildConfigList
	ImageList                 imageapi.ImageList
	ImageRepositoryList       imageapi.ImageRepositoryList
	DeploymentList            deployapi.DeploymentList
	DeploymentConfigList      deployapi.DeploymentConfigList
	RouteList                 routeapi.RouteList
	ContainerStatus           kapi.ContainerStatus
}

func main() {
	packages := []schemagen.PackageDescriptor{
		{"github.com/GoogleCloudPlatform/kubernetes/pkg/api/v1beta2", "io.fabric8.kubernetes.api.model", "kubernetes_"},
		{"github.com/GoogleCloudPlatform/kubernetes/pkg/api", "io.fabric8.kubernetes.api.model", "kubernetes_"},
		{"github.com/GoogleCloudPlatform/kubernetes/pkg/util", "io.fabric8.kubernetes.api.model", "kubernetes_util_"},
		{"github.com/fsouza/go-dockerclient", "io.fabric8.docker.api.model", "docker_"},
		{"github.com/openshift/origin/pkg/build/api", "io.fabric8.openshift.api.model", "os_build_"},
		{"github.com/openshift/origin/pkg/deploy/api", "io.fabric8.openshift.api.model", "os_deploy_"},
		{"github.com/openshift/origin/pkg/image/api", "io.fabric8.openshift.api.model", "os_image_"},
		{"github.com/openshift/origin/pkg/route/api", "io.fabric8.openshift.api.model", "os_route_"},
	}

	typeMap := map[reflect.Type]reflect.Type{
		reflect.TypeOf(kutil.Time{}): reflect.TypeOf(""),
		reflect.TypeOf(time.Time{}):  reflect.TypeOf(""),
		reflect.TypeOf(struct{}{}):   reflect.TypeOf(""),
	}
	schema, err := schemagen.GenerateSchema(reflect.TypeOf(Schema{}), packages, typeMap)
	if err != nil {
		fmt.Errorf("An error occurred: %v", err)
		return
	}
	b, _ := json.Marshal(&schema)
	result := string(b)
	result = strings.Replace(result, "\"additionalProperty\":", "\"additionalProperties\":", -1)
	fmt.Println(result)
}
