package clients

type MidtransClient struct {
	ServerKey    string
	IsProduction bool
}

type IMidtransClient interface {
	CreatePaymentLink()
}
