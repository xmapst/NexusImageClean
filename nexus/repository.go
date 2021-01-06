package nexus

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"
)

type RepositoryResponse struct {
	Tid    int    `json:"tid,omitempty"`
	Action string `json:"action,omitempty"`
	Method string `json:"method,omitempty"`
	Result struct {
		Message string `json:"message,omitempty"`
		Success bool   `json:"success,omitempty"`
		Data    []struct {
			ID          string `json:"id,omitempty"`
			Text        string `json:"text,omitempty"`
			Type        string `json:"type,omitempty"`
			Leaf        bool   `json:"leaf,omitempty"`
			ComponentID string `json:"componentId,omitempty"`
			AssetID     string `json:"assetId,omitempty"`
		} `json:"data,omitempty"`
	} `json:"result,omitempty"`
	Type string `json:"type,omitempty"`
}

func (s *DefaultClient) GetRepository(version string, direction ...string) (response *RepositoryResponse, err error) {
	tid, _ := rand.Int(rand.Reader, big.NewInt(9999))
	payloadStr := `{"action":"coreui_Browse","method":"read","data":[{"repositoryName":"%s","sort":[{"property":"leaf","direction":"%s"}],"node":"%s"}],"type":"rpc","tid":%d}`
	if len(direction) > 0 && direction[0] == "" {
		direction[0] = "ASC"
	}
	payload := strings.NewReader(fmt.Sprintf(payloadStr, s.Repository, direction, version, tid))
	bodyByte, req, err := s.Post(RpcEndpoint, payload)
	if err != nil {
		return
	}
	if req.StatusCode != 200 {
		err = fmt.Errorf("nexus is unable to serve requests")
		return
	}

	if err = json.Unmarshal(bodyByte, &response); err != nil {
		return
	}
	if !response.Result.Success {
		return nil, fmt.Errorf(response.Result.Message)
	}
	if len(response.Result.Data) == 0 {
		return nil, fmt.Errorf("no found data")
	}
	for _, _t := range response.Result.Data {
		if _t.Type != Folder {
			continue
		}
		_, err := s.GetRepositoryAllTags(_t.ID)
		if err == nil || err.Error() != NotFount {
			continue
		}
		_r, err := s.GetRepository(_t.ID)
		if err != nil {
			continue
		}
		for _, __t := range _r.Result.Data {
			if __t.Type != Folder {
				continue
			}
			response.Result.Data = append(response.Result.Data, __t)
		}
	}

	return
}

type RepositoryTagsResponse struct {
	Tid    int    `json:"tid,omitempty"`
	Action string `json:"action,omitempty"`
	Method string `json:"method,omitempty"`
	Result struct {
		Message string `json:"message,omitempty"`
		Success bool   `json:"success,omitempty"`
		Data    []struct {
			ID          string `json:"id,omitempty"`
			Text        string `json:"text,omitempty"`
			Type        string `json:"type,omitempty"`
			Leaf        bool   `json:"leaf,omitempty"`
			ComponentID string `json:"componentId,omitempty"`
			AssetID     string `json:"assetId,omitempty"`
		} `json:"data,omitempty"`
	} `json:"result,omitempty"`
	Type string `json:"type,omitempty"`
}

func (s *DefaultClient) GetRepositoryAllTags(nodeID string, direction ...string) (response *RepositoryTagsResponse, err error) {
	tid, _ := rand.Int(rand.Reader, big.NewInt(9999))
	payloadStr := `{"action":"coreui_Browse","method":"read","data":[{"repositoryName":"%s","sort":[{"property":"leaf","direction":"%s"}],"node":"%s/tags"}],"type":"rpc","tid":%d}`
	if len(direction) > 0 && direction[0] == "" {
		direction[0] = "ASC"
	}
	payload := strings.NewReader(fmt.Sprintf(payloadStr, s.Repository, direction, nodeID, tid))
	bodyByte, req, err := s.Post(RpcEndpoint, payload)
	if err != nil {
		return
	}
	if req.StatusCode != 200 {
		err = fmt.Errorf("nexus is unable to serve requests")
		return
	}

	if err = json.Unmarshal(bodyByte, &response); err != nil {
		return
	}
	if !response.Result.Success {
		return nil, fmt.Errorf(response.Result.Message)
	}
	if len(response.Result.Data) == 0 {
		return nil, fmt.Errorf("No found data")
	}
	return
}

type RepositoryTagInfoResponse struct {
	Tid    int    `json:"tid,omitempty"`
	Action string `json:"action,omitempty"`
	Method string `json:"method,omitempty"`
	Result struct {
		Message string `json:"message,omitempty"`
		Success bool   `json:"success,omitempty"`
		Data    struct {
			ID                       string    `json:"id,omitempty" yaml:"id"`
			Name                     string    `json:"name,omitempty" yaml:"name"`
			Format                   string    `json:"format,omitempty" yaml:"format"`
			ContentType              string    `json:"contentType,omitempty" yaml:"contentType"`
			Size                     int       `json:"size,omitempty" yaml:"size"`
			RepositoryName           string    `json:"repositoryName,omitempty" yaml:"repositoryName"`
			ContainingRepositoryName string    `json:"containingRepositoryName,omitempty" yaml:"containingRepositoryName"`
			BlobCreated              time.Time `json:"blobCreated,omitempty" yaml:"blobCreated"`
			BlobUpdated              time.Time `json:"blobUpdated,omitempty" yaml:"blobUpdated"`
			LastDownloaded           time.Time `json:"lastDownloaded,omitempty" yaml:"lastDownloaded"`
			DownloadCount            int       `json:"downloadCount,omitempty" yaml:"downloadCount"`
			BlobRef                  string    `json:"blobRef,omitempty" yaml:"blobRef"`
			ComponentID              string    `json:"componentId,omitempty" yaml:"componentId"`
			CreatedBy                string    `json:"createdBy,omitempty" yaml:"createdBy"`
			CreatedByIP              string    `json:"createdByIp,omitempty" yaml:"createdByIp"`
			Attributes               struct {
				Checksum struct {
					Sha1   string `json:"sha1,omitempty" yaml:"sha1"`
					Sha256 string `json:"sha256,omitempty" yaml:"sha256"`
				} `json:"checksum,omitempty" yaml:"checksum"`
				Cache struct {
				} `json:"cache,omitempty" yaml:"cache"`
				Provenance struct {
					HashesNotVerified bool `json:"hashes_not_verified,omitempty" yaml:"hashes_not_verified"`
				} `json:"provenance,omitempty" yaml:"provenance"`
				Content struct {
					LastModified time.Time `json:"last_modified,omitempty" yaml:"last_modified"`
				} `json:"content" yaml:"content"`
				Docker struct {
					AssetKind     string `json:"asset_kind,omitempty" yaml:"asset_kind"`
					ContentDigest string `json:"content_digest,omitempty" yaml:"content_digest"`
				} `json:"docker,omitempty" yaml:"docker"`
			} `json:"attributes,omitempty" yaml:"attributes"`
		} `json:"data,omitempty"`
	} `json:"result,omitempty"`
	Type string `json:"type,omitempty"`
}

func (s *DefaultClient) GetRepositoryTagInfo(tagID string) (response *RepositoryTagInfoResponse, err error) {
	tid, _ := rand.Int(rand.Reader, big.NewInt(9999))
	payloadStr := `{"action":"coreui_Component","method":"readAsset","data":["%s","%s"],"type":"rpc","tid":%d}`
	payload := strings.NewReader(fmt.Sprintf(payloadStr, tagID, s.Repository, tid))
	bodyByte, req, err := s.Post(RpcEndpoint, payload)
	if err != nil {
		return
	}
	if req.StatusCode != 200 {
		err = fmt.Errorf("nexus is unable to serve requests")
		return
	}

	if err = json.Unmarshal(bodyByte, &response); err != nil {
		return
	}
	if !response.Result.Success {
		return nil, fmt.Errorf(response.Result.Message)
	}
	return
}

type DeleteImageTagResponse struct {
	Tid    int    `json:"tid,omitempty"`
	Action string `json:"action,omitempty"`
	Method string `json:"method,omitempty"`
	Result struct {
		Success bool        `json:"success,omitempty"`
		Message string      `json:"message,omitempty"`
		Data    interface{} `json:"data,omitempty"`
	} `json:"result,omitempty"`
	Type string `json:"type,omitempty"`
}

func (s *DefaultClient) DeleteImageTag(tagID string) (err error) {
	var response DeleteImageTagResponse
	tid, _ := rand.Int(rand.Reader, big.NewInt(9999))
	payloadStr := `{"action":"coreui_Component","method":"deleteAsset","data":["%s","%s"],"type":"rpc","tid":%d}`
	payload := strings.NewReader(fmt.Sprintf(payloadStr, tagID, s.Repository, tid))
	bodyByte, req, err := s.Post(RpcEndpoint, payload)
	if err != nil {
		return
	}
	if req.StatusCode != 200 {
		err = fmt.Errorf("nexus is unable to serve requests")
		return
	}

	if err = json.Unmarshal(bodyByte, &response); err != nil {
		return
	}
	if !response.Result.Success {
		return fmt.Errorf(response.Result.Message)
	}
	return
}
