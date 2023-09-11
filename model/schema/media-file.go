package schema

/*
	MediaFile:
     description: This object contains a url to a media file.
     type: object
     properties:
       mimetype:
         description: 'indicates the nature and format of the document, file, or assortment of bytes. MIME types are defined and standardized in IETF''s RFC 6838'
         type: string
       url:
         description: The URL of the file
         type: string
         format: uri
       signature:
         description: The digital signature of the file signed by the sender
         type: string
       dsa:
         description: The signing algorithm used by the sender
         type: string
*/

type MediaFile struct {
	Mimetype  string `json:"mimetype"`  // indicates the nature and format of the document, file, or assortment of bytes. MIME types are defined and standardized in IETF's RFC 6838
	Url       string `json:"url"`       // The URL of the file
	Signature string `json:"signature"` // The digital signature of the file signed by the sender
	Dsa       string `json:"dsa"`       // The signing algorithm used by the sender
}
