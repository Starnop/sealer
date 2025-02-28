// Copyright © 2021 Alibaba Group Holding Ltd.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package image

//
//import (
//	"github.com/sealerio/sealer/pkg/image/store"
//	v1 "github.com/sealerio/sealer/types/api/v1"
//	"net"
//)
//
//// MetadataService is the interface for providing image metadata service
//type MetadataService interface {
//	Tag(imageName, tarImageName string) error
//	List() (store.ImageMetadataMap, error)
//	GetImage(imageName string, platform *v1.Platform) (*v1.Image, error)
//	GetRemoteImage(imageName string, platform *v1.Platform) (v1.Image, error)
//	DeleteImage(imageName string, platform *v1.Platform) error
//}
//
//// FileService is the interface for file operations
//type FileService interface {
//	Load(imageSrc string) error
//	Save(imageName, imageTar string, platforms []*v1.Platform) error
//	Merge(image *v1.Image) error
//}
//
//// Service is image service
//type Service interface {
//	Pull(imageName string, platform []*v1.Platform) error
//	PullIfNotExist(imageName string, platform []*v1.Platform) error
//	Push(imageName string) error
//	Delete(imageName string, platforms []*v1.Platform) error
//	Login(RegistryURL, RegistryUsername, RegistryPasswd string) error
//	//TODO
//	Mount(imageName string, hosts []net.IP) error
//	UMount(imageName string, hosts []net.IP) error
//	CacheBuilder
//}
//
//type LayerService interface {
//	LayerStore() store.LayerStore
//}
