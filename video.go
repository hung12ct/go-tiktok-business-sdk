package tiktokbiz

import (
	"context"
	"net/http"
)

const (
	videoGetPATH          = "/file/video/ad/info/"
	videoUploadPATH       = "/file/video/ad/upload/"
	videoSearchPATH       = "/file/video/ad/search/"
	updateVideoNamePATH   = "/file/video/ad/update/"
	suggestCoverVideoPATH = "file/video/suggestcover/"
)

// Upload method
type UploadType_ string

const (
	UPLOAD_BY_FILE     UploadType_ = "UPLOAD_BY_FILE"
	UPLOAD_BY_URL      UploadType_ = "UPLOAD_BY_URL"
	UPLOAD_BY_FILE_ID  UploadType_ = "UPLOAD_BY_FILE_ID"
	UPLOAD_BY_VIDEO_ID UploadType_ = "UPLOAD_BY_VIDEO_ID"
)

type VideoAdRequest struct {
	// Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Video name.
	FileName string `json:"file_name,omitempty"`
	// Image upload method.
	UploadType UploadType_ `json:"upload_type,omitempty"`
	// Video file.
	// VideoFile *UploadField `json:"video_file,omitempty"`
	// Video MD5 (used for server verification).; Required when upload_type is UPLOAD_BY_FILE
	VideoSignature string `json:"video_signature,omitempty"`
	// Video URL address; Required when upload_type is UPLOAD_BY_URL
	VideoURL string `json:"video_url,omitempty"`
	// Id of the file that you want to upload; Required when upload_type is UPLOAD_BY_FILE_ID
	FileId string `json:"file_id,omitempty"`
	// Video id; Required when upload_type is UPLOAD_BY_VIDEO_ID
	VideoID string `json:"video_id,omitempty"`
	// The video is third party or not
	IsThirdParty bool `json:"is_third_party,omitempty"`
	// Whether to automatically detect an issue in your video
	FlawDetect bool `json:"flaw_detect,omitempty"`
	// Whether to automatically fix the detected issue.
	AutoFixEnabled bool `json:"auto_fix_enabled,omitempty"`
	// Whether to automatically upload the fixed video to your creative library.
	AutoBindEnabled bool `json:"auto_bind_enabled,omitempty"`
}

type VideoData struct {
	// Temporary URL for video cover, valid for six hours and needs to be re-acquired after expiration
	VideoCoverURL string `json:"video_cover_url,omitempty"`
	// Video format.
	Format string `json:"format,omitempty"`
	// Video preview link, valid for six hours and needs to be re-acquired after expiration.
	PreviewUrl string `json:"preview_url,omitempty"`
	// The expiration time of the video preview link, in the format of YYYY-MM-DD HH:MM:SS (UTC+0)
	PreviewUrlExpireTime string `json:"preview_url_expire_time,omitempty"`
	// Video name.
	FileName string `json:"file_name,omitempty"`
	// Video name.
	Displayable bool `json:"displayable,omitempty"`
	// Video height.
	Height int `json:"height,omitempty"`
	// Video width.
	Width int `json:"width,omitempty"`
	// Bit rate in bps
	BitRate int64 `json:"bit_rate,omitempty"`
	// Creation time. UTC time. Format: 2020-06-10T07:39:14Z.
	CreateTime string `json:"create_time,omitempty"`
	// Modification time. UTC time. Format: 2020-06-10T07:39:14Z.
	ModifyTime string `json:"modify_time,omitempty"`
	// Video file MD5.
	Signature string `json:"signature,omitempty"`
	// Video duration, in seconds.
	Duration float64 `json:"duration,omitempty"`
	// Video ID, can be used to create ad in ad delivery.
	VideoID string `json:"video_id,omitempty"`
	// Video size, in bytes.
	Size uint64 `json:"size,omitempty"`
	// Material ID
	MaterialID string `json:"material_id,omitempty"`
	// Available placements
	AllowedPlacements []string `json:"allowed_placements,omitempty"`
	// Whether the video is downloadable
	AllowDownload bool `json:"allow_download,omitempty"`
	// Fix task ID
	FixTaskId string `json:"fix_task_id,omitempty"`
	// Video issue types
	FlawTypes string `json:"flaw_types,omitempty"`
}

// VideoUpload: upload a video
func (c *Client) VideoUpload(ctx context.Context, request *VideoAdRequest) (
	respData []*VideoData,
	err error,
) {
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(videoUploadPATH), withBody(request))
	if err != nil {
		return nil, err
	}

	var responseData []*VideoData
	if err := c.sendAndUnmarshal(req, &responseData); err != nil {
		return nil, err
	}

	return responseData, nil
}

type UpdateVideoNameRequest struct {
	// Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Video name.
	FileName string `json:"file_name,omitempty"`
	// Video id
	VideoID string `json:"video_id,omitempty"`
}

type UpdateVideoNameData struct {
	// Returned data
	VideoData interface{} `json:"data,omitempty"`
}

// VideoNameUpdate: update the name of a video
func (c *Client) VideoNameUpdate(ctx context.Context, request *UpdateVideoNameRequest) (
	respData *UpdateVideoNameData,
	err error,
) {
	req, err := c.newRequest(ctx, http.MethodPost, c.fullURL(updateVideoNamePATH), withBody(request))
	if err != nil {
		return nil, err
	}

	var responseData UpdateVideoNameData
	if err := c.sendAndUnmarshal(req, &responseData); err != nil {
		return nil, err
	}

	return &responseData, nil
}

type VideoAdGetRequest struct {
	// Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Video ID list. Up to 60 IDs per request.
	VideoIDs []string `json:"video_ids,omitempty"`
}

type VideoGetResponseData struct {
	// List contains the list of videos.
	List []VideoData `json:"list,omitempty"`
}

// VideoGet: get info about videos
func (c *Client) VideoGet(ctx context.Context, request *VideoAdGetRequest) (
	respData *VideoGetResponseData,
	err error,
) {
	req, err := c.newRequest(ctx, http.MethodGet, c.fullURL(videoGetPATH), withBody(request))
	if err != nil {
		return nil, err
	}

	var responseData VideoGetResponseData
	if err := c.sendAndUnmarshal(req, &responseData); err != nil {
		return nil, err
	}

	return &responseData, nil
}

type VideoGetFilter struct {
	// Video height.
	Height int `json:"height,omitempty"`
	// Video aspect ratio
	Ratio []float64 `json:"ratio,omitempty"`
	// A list of video IDs. At most 100 IDs can be included in the list.
	VideoIDs []string `json:"video_ids,omitempty"`
	// A list of material IDs. At most 100 IDs can be included in the list.
	MaterialIDs []string `json:"material_ids,omitempty"`
	// Video width.
	Width int `json:"width,omitempty"`
	// Video name.
	Displayable bool `json:"displayable,omitempty"`
}

type VideoAdSearchRequest struct {
	// Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Filters on the data.
	Filtering *VideoGetFilter `json:"filtering,omitempty"`
	// Current page number. Default value: 1
	Page int `json:"page,omitempty"`
	// Page size. Default value: 20. Value range: 1-100
	PageSize int `json:"page_size,omitempty"`
}

type VideoAdSearchResponseData struct {
	// A list of video information
	List []VideoData `json:"list,omitempty"`
	// Pagination information.
	PageInfo *PageInfo `json:"page_info,omitempty"`
}

// VideoSearch: search for videos
func (c *Client) VideoSearch(ctx context.Context, request *VideoAdSearchRequest) (
	respData *VideoAdSearchResponseData,
	err error,
) {
	req, err := c.newRequest(ctx, http.MethodGet, c.fullURL(videoSearchPATH), withBody(request))
	if err != nil {
		return nil, err
	}

	var responseData VideoAdSearchResponseData
	if err := c.sendAndUnmarshal(req, &responseData); err != nil {
		return nil, err
	}

	return &responseData, nil
}

type SuggestedVideoThumbnailsRequest struct {
	// Advertiser ID
	AdvertiserID string `json:"advertiser_id,omitempty"`
	// Video id
	VideoID string `json:"video_id,omitempty"`
	// Number of cover candidates you want to get. Range: 1-10. Default value: 10 .
	PosterNumber int `json:"poster_number,omitempty"`
}

type SuggestedThumbnail struct {
	// Image width..
	Width int `json:"width,omitempty"`
	// Image height.
	Height int `json:"height,omitempty"`
	// Image ID.
	ID string `json:"id,omitempty"`
	// Picture preview address, valid for an hour and needs to be re-acquired after expiration.
	URL string `json:"url,omitempty"`
}

type SuggestedVideoThumbnailsData struct {
	// A list of image information..
	List []SuggestedThumbnail `json:"list,omitempty"`
}

// VideoThumbnailsSuggested: get suggested thumbnails for a video
func (c *Client) VideoThumbnailsSuggested(ctx context.Context, request *SuggestedVideoThumbnailsRequest) (
	respData *SuggestedVideoThumbnailsData,
	err error,
) {
	req, err := c.newRequest(ctx, http.MethodGet, c.fullURL(suggestCoverVideoPATH), withBody(request))
	if err != nil {
		return nil, err
	}

	var responseData SuggestedVideoThumbnailsData
	if err := c.sendAndUnmarshal(req, &responseData); err != nil {
		return nil, err
	}

	return &responseData, nil
}
