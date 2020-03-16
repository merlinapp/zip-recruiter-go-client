package jobs

type Client interface {
	Get(request ZipRequest) (ZipResponse, error)
}
