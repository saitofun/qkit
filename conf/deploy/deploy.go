package deploy

type Deployer interface {
	Name() string
	Write(filename string) error
}
