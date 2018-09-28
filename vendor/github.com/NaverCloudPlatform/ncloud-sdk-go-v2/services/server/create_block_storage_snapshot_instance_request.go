/*
 * server
 *
 * <br/>https://ncloud.apigw.ntruss.com/server/v2
 *
 * API version: 2018-09-21T05:59:54Z
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package server

type CreateBlockStorageSnapshotInstanceRequest struct {

	// 블록스토리지인스턴스번호
BlockStorageInstanceNo *string `json:"blockStorageInstanceNo"`

	// 블록스토리지스냅샷이름
BlockStorageSnapshotName *string `json:"blockStorageSnapshotName,omitempty"`

	// 블록스토리지스냅샷설명
BlockStorageSnapshotDescription *string `json:"blockStorageSnapshotDescription,omitempty"`
}
